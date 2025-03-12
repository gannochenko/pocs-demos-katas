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
	typeUtil "backend/internal/util/types"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

const (
	taskBufferThreshold = 5
	imageProcessingQueueBatchSize = 30
)

type Service struct {
	configService interfaces.ConfigService
	eventBusService interfaces.EventBusService
	loggerService   interfaces.LoggerService
	imageQueueRepository interfaces.ImageProcessingQueueRepository

	hasNewTasks atomic.Bool
	taskBuffer sync.Map

	channel chan database.ImageProcessingQueue
}

func NewImageProcessor(
	configService interfaces.ConfigService,
	eventBusService interfaces.EventBusService,
	loggerService interfaces.LoggerService,
	imageQueueRepository interfaces.ImageProcessingQueueRepository,
) *Service {
	result := &Service{
		configService: configService,
		eventBusService: eventBusService,
		loggerService:   loggerService,
		imageQueueRepository: imageQueueRepository,
		channel: make(chan database.ImageProcessingQueue),
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
		s.loggerService.Info(ctx, "new event received", logger.F("event", event))
		s.hasNewTasks.Store(true)
	}

	err = s.eventBusService.AddEventListener(domain.EventBusEventTypeImageCreated, callback)
	if err != nil {
		return syserr.Wrap(err, "could not start listening to events")
	}
	defer func(){
		// todo: this will not work
		s.eventBusService.RemoveEventListener(domain.EventBusEventTypeImageCreated, callback)
	}()

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
		go s.processImage(ctx, i, wg)
	}

	return nil
}

func (s *Service) processImage(ctx context.Context, workerId int, wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		s.loggerService.Info(ctx, fmt.Sprintf("worker %d exited", workerId), logger.F("workerId", workerId))
	}()

	operationID := uuid.New().String()
	processCtx := ctxUtil.WithOperationID(ctx, operationID)

	s.loggerService.Info(ctx, fmt.Sprintf("worker %d started", workerId), logger.F("workerId", workerId))

	for {
		select {
		case <-processCtx.Done():
			return
		case task := <-s.channel:
			// todo: process here, with a timeout

			s.loggerService.Info(ctx, "processing image", logger.F("imageId", task.ID))

			err := s.markTaskSucessful(processCtx, task, operationID)
			if err != nil {
				s.loggerService.Error(ctx, "could not update image processing queue", logger.F("error", err))
				continue
			}

			err = s.eventBusService.TriggerEvent(&domain.EventBusEvent{
				Type: domain.EventBusEventTypeImageProcessed,
				Payload: &domain.EventBusEventPayloadImageProcessed{
					ImageID: task.ID,
				},
			})
			if err != nil {
				s.loggerService.Error(ctx, "could not trigger event bus event", logger.F("error", err))
				continue
			}

			s.taskBuffer.Delete(task.ID)
		}
	}
}

func (s *Service) markTaskSucessful(ctx context.Context, task database.ImageProcessingQueue, operationID string) error {
	return s.imageQueueRepository.Update(ctx, nil, &database.ImageProcessingQueueUpdate{
		ID: task.ID,
		OperationID: &database.FieldValue[*string]{Value: &operationID},
		IsCompleted: &database.FieldValue[*bool]{Value: lo.ToPtr(true)},
		CompletedAt: &database.FieldValue[*time.Time]{Value: lo.ToPtr(time.Now().UTC())},
	})
}