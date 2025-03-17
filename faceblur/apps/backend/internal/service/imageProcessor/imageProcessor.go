package imageProcessor

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"backend/interfaces"
	"backend/internal/database"
	"backend/internal/domain"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"

	ctxUtil "backend/internal/util/ctx"
	imageUtil "backend/internal/util/image"
	typeUtil "backend/internal/util/types"

	otelMetric "go.opentelemetry.io/otel/metric"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

const (
	taskBufferThreshold = 5
	imageProcessingQueueBatchSize = 30
	meterName = "image_processor"
)

type Service struct {
	configService interfaces.ConfigService
	eventBusService interfaces.EventBusService
	loggerService   interfaces.LoggerService
	imageQueueRepository interfaces.ImageProcessingQueueRepository
	imageRepository interfaces.ImageRepository
	faceDetectionService interfaces.FaceDetectionService
	storageService interfaces.StorageService
	monitoringService interfaces.MonitoringService

	hasNewTasks atomic.Bool
	taskBuffer sync.Map

	channel chan database.ImageProcessingQueue
}

func NewImageProcessor(
	configService interfaces.ConfigService,
	eventBusService interfaces.EventBusService,
	loggerService interfaces.LoggerService,
	imageQueueRepository interfaces.ImageProcessingQueueRepository,
	imageRepository interfaces.ImageRepository,
	faceDetectionService interfaces.FaceDetectionService,
	storageService interfaces.StorageService,
	monitoringService interfaces.MonitoringService,
) *Service {
	result := &Service{
		configService: configService,
		eventBusService: eventBusService,
		loggerService:   loggerService,
		imageQueueRepository: imageQueueRepository,
		imageRepository: imageRepository,
		channel: make(chan database.ImageProcessingQueue),
		faceDetectionService: faceDetectionService,
		storageService: storageService,
		monitoringService: monitoringService,
	}

	result.hasNewTasks.Store(true) // upon startup check if there are some tasks

	return result
}

func (s *Service) Start(ctx context.Context) error {
	var wg sync.WaitGroup

	err := s.init(ctx, &wg)
	if err != nil {
		return syserr.Wrap(err, "could not initialize")
	}

	callback := func(event *domain.EventBusEvent) {
		s.loggerService.Info(ctx, "worker: new event received", logger.F("event", event))
		s.hasNewTasks.Store(true)
	}

	err = s.eventBusService.AddEventListener(domain.EventBusEventTypeImageCreated, callback)
	if err != nil {
		return syserr.Wrap(err, "could not start listening to events")
	}

	wg.Add(1)
	go func(){
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				err = syserr.Wrap(ctx.Err(), "context is done")
				return
			default:
				if s.hasNewTasks.Swap(false) {
					err = s.ProcessImages(ctx)
					if err != nil {
						s.loggerService.LogError(ctx, syserr.Wrap(err, "could not process images"))
					}
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	wg.Wait()

	return err
}

func (s *Service) ProcessImages(ctx context.Context) error {
	s.loggerService.Info(ctx, "processing images")

	for {
		select {
		case <-ctx.Done():
			return syserr.Wrap(ctx.Err(), "context is done")
		default:
			if typeUtil.GetSyncMapSize(&s.taskBuffer) < taskBufferThreshold {
				// the buffer is getting empty, let's add some items
				res, err := s.imageQueueRepository.List(ctx, nil, database.ImageProcessingQueueListParameters{
					Filter: &database.ImageProcessingQueueFilter{
						IsFailed: lo.ToPtr(false),
						IsCompleted: lo.ToPtr(false),
					},
					Pagination: &database.Pagination{
						PageNumber: 1,
						PageSize: imageProcessingQueueBatchSize,
					},
				})
				if err != nil {
					return syserr.Wrap(err, "could not list images")
				}

				s.loggerService.Info(ctx, "found images", logger.F("count", len(res)))

				if len(res) == 0 {
					// nothing left to do
					return nil
				}

				wasAdded := false
				for _, task := range res {
					if _, ok := s.taskBuffer.Load(task.ID); !ok {
						s.taskBuffer.Store(task.ID, task)
						wasAdded = true

						s.channel <- task
					}
				}

				if !wasAdded {
					// all tasks are already in the buffer, exiting
					return nil
				}
			}

			time.Sleep(time.Second)
		}
	}
}

func (s *Service) Stop() error {
	// todo: implement this
	// s.eventBusService.RemoveEventListener(domain.EventBusEventTypeImageCreated, callback)

	return nil
}

func (s *Service) init(ctx context.Context, wg *sync.WaitGroup) error {
	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not extract config")
	}

	workerPoolSize := config.Backend.Worker.ThreadCount
	if workerPoolSize < 1 {
		return syserr.NewInternal("invalid worker pool size", syserr.F("value", workerPoolSize))
	}

	for i := 0;	i < workerPoolSize; i++ {
		wg.Add(1)
		go s.processImages(ctx, i, wg)
	}

	return nil
}

func (s *Service) processImages(ctx context.Context, workerId int, wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		s.loggerService.Info(ctx, fmt.Sprintf("worker %d exited", workerId), logger.F("workerId", workerId))
	}()

	s.loggerService.Info(ctx, fmt.Sprintf("worker %d started", workerId), logger.F("workerId", workerId))

	for {
		select {
		case <-ctx.Done():
			return
		case task := <-s.channel:
			operationID := uuid.New().String()
			processCtx := ctxUtil.WithOperationID(ctx, operationID)

			err := s.processTask(processCtx, task)
			if err != nil {
				s.loggerService.LogError(processCtx, syserr.Wrap(err, "could not process task"))
				err = s.markTaskFailed(processCtx, task, operationID, err.Error())
				if err != nil {
					s.loggerService.LogError(processCtx, syserr.Wrap(err, "could not update image processing queue"))
				}
				err = s.markImageProcessed(processCtx, task.ImageID, true, nil)
				if err != nil {
					s.loggerService.LogError(processCtx, syserr.Wrap(err, "could not mark image processed"))
				}

				s.monitoringService.AddInt64Counter(processCtx, meterName, "error", 1, "", "")
			}

			s.taskBuffer.Delete(task.ID)
		}
	}
}

func (s *Service) processTask(processCtx context.Context, task database.ImageProcessingQueue) error {
	var err error
	startTime := time.Now()

	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not extract config")
	}

	s.loggerService.Info(processCtx, "processing image", logger.F("imageId", task.ID))

	operationID := ctxUtil.GetOperationID(processCtx)

	taskCtx, cancelTaskCtx := context.WithTimeout(processCtx, time.Second * 15)
	defer cancelTaskCtx()

	var detections []*domain.BoundingBox

	imageElement, err := s.imageRepository.GetByID(taskCtx, nil, task.ImageID)
	if err != nil {
		return syserr.Wrap(err, "could not get image")
	}

	if imageElement == nil {
		return syserr.NewInternal("image not found", syserr.F("id", task.ID))
	}

	image, err := imageUtil.DownloadImage(imageElement.OriginalURL)
	if err != nil {
		return syserr.Wrap(err, "could not download image")
	}

	detections, err = s.faceDetectionService.Detect(taskCtx, image)
	if err != nil {
		return syserr.Wrap(err, "could not detect faces")
	}

	if ctxUtil.IsTimeouted(taskCtx) {
		return syserr.Wrap(err, "context is done")
	}

	image, err = imageUtil.BlurBoxes(image, detections, 9.0)
	if err != nil {
		return syserr.Wrap(err, "could not blur faces")
	}

	buffer, err := imageUtil.EncodeImage(image, "jpg", 90)
	if err != nil {
		return syserr.Wrap(err, "could not encode image")
	}

	writer, err := s.storageService.GetWriter(taskCtx, config.Storage.ImageBucketName, operationID)
	if err != nil {
		return syserr.Wrap(err, "could not get writer")
	}
	defer writer.Close()

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return syserr.Wrap(err, "could not write image")
	}

	if ctxUtil.IsTimeouted(taskCtx) {
		return syserr.Wrap(err, "context is done")
	}

	err = s.markImageProcessed(processCtx, task.ImageID, false, lo.ToPtr(s.storageService.GetPublicURL(config.Storage.ImageBucketName, operationID)))
	if err != nil {
		return syserr.Wrap(err, "could not update image")
	}

	err = s.markTaskSucessful(processCtx, task, operationID)
	if err != nil {
		return syserr.Wrap(err, "could not update image processing queue")
	}

	err = s.eventBusService.TriggerEvent(&domain.EventBusEvent{
		Type: domain.EventBusEventTypeImageProcessed,
		Payload: &domain.EventBusEventPayloadImageProcessed{
			ImageID: task.ID,
			CreatorID: task.CreatedBy.String(),
		},
	})
	if err != nil {
		return syserr.Wrap(err, "could not trigger event bus event")
	}

	s.loggerService.Info(processCtx, "image was processed", logger.F("imageId", task.ID))
	s.monitoringService.AddInt64Counter(processCtx, meterName, "processed_images", 1, "", "")

	endTime := time.Since(startTime)
	s.monitoringService.RecordInt64Histogram(processCtx, meterName, "image_processing_duration", endTime.Milliseconds(), "", "", otelMetric.WithExplicitBucketBoundaries(histogramBoundaries...))

	return nil
}

func (s *Service) markTaskSucessful(ctx context.Context, task database.ImageProcessingQueue, operationID string) error {
	return s.imageQueueRepository.Update(ctx, nil, &database.ImageProcessingQueueUpdate{
		ID: task.ID,
		OperationID: &database.FieldValue[*string]{Value: &operationID},
		IsCompleted: &database.FieldValue[*bool]{Value: lo.ToPtr(true)},
		CompletedAt: &database.FieldValue[*time.Time]{Value: lo.ToPtr(time.Now().UTC())},
	})
}

func (s *Service) markTaskFailed(ctx context.Context, task database.ImageProcessingQueue, operationID string, reason string) error {
	return s.imageQueueRepository.Update(ctx, nil, &database.ImageProcessingQueueUpdate{
		ID: task.ID,
		OperationID: &database.FieldValue[*string]{Value: &operationID},
		IsCompleted: &database.FieldValue[*bool]{Value: lo.ToPtr(false)},
		IsFailed: &database.FieldValue[*bool]{Value: lo.ToPtr(true)},
		CompletedAt: &database.FieldValue[*time.Time]{Value: nil},
		FailureReason: &database.FieldValue[*string]{Value: &reason},
	})
}

func (s *Service) markImageProcessed(ctx context.Context, imageID uuid.UUID, failed bool, url *string) error {
	return s.imageRepository.Update(ctx, nil, &database.ImageUpdate{
		ID: imageID,
		IsProcessed: &database.FieldValue[*bool]{Value: lo.ToPtr(true)},
		IsFailed: &database.FieldValue[*bool]{Value: &failed},
		URL: &database.FieldValue[*string]{Value: url},
	})
}

var histogramBoundaries = []float64{
	500.0, 1000.0, 1500.0, 2000.0, 2500.0, 3000.0, 3500.0, 4000.0, 4500.0, 5000.0,
	5500.0, 6000.0, 6500.0, 7000.0, 7500.0, 8000.0, 8500.0, 9000.0, 9500.0, 10000.0,
	10500.0, 11000.0, 11500.0, 12000.0, 12500.0, 13000.0, 13500.0, 14000.0, 14500.0,
	15000.0, 15500.0, 16000.0, 16500.0, 17000.0, 17500.0, 18000.0, 18500.0, 19000.0,
	19500.0, 20000.0, 20500.0, 21000.0, 21500.0, 22000.0, 22500.0, 23000.0, 23500.0,
	24000.0, 24500.0, 25000.0, 25500.0, 26000.0, 26500.0, 27000.0, 27500.0, 28000.0,
	28500.0, 29000.0, 29500.0, 30000.0,
}
