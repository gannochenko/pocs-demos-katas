package imageProcessor

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"backend/interfaces"
	"backend/internal/database"
	"backend/internal/domain"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"

	"github.com/samber/lo"
)

type Service struct {
	configService interfaces.ConfigService
	eventBusService interfaces.EventBusService
	loggerService   interfaces.LoggerService
	imageQueueRepository interfaces.ImageProcessingQueueRepository

	hasNewMessages atomic.Bool
	bufferSize atomic.Int64

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

	result.hasNewMessages.Store(false)

	return result
}

func (s *Service) Start(ctx context.Context) error {
	// todo: implement the worker pool here
	err := s.init(ctx)
	if err != nil {
		return syserr.Wrap(err, "could not initialize")
	}

	err = s.eventBusService.AddEventListener(domain.EventBusEventTypeImageCreated, func(event *domain.EventBusEvent) {
		s.loggerService.Info(ctx, "new event received", logger.F("event", event))
		s.hasNewMessages.Store(true)
	})
	if err != nil {
		return syserr.Wrap(err, "could not start listening to events")
	}

	// running one cycle, no waitgroup is needed
	for {
		select {
		case <-ctx.Done():
			return syserr.Wrap(ctx.Err(), "context is done")
		default:
			if s.hasNewMessages.Swap(false) {
				err = s.ProcessImages(ctx)
				if err != nil {
					s.loggerService.LogError(ctx, syserr.Wrap(err, "could not process images"))
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Service) ProcessImages(ctx context.Context) error {
	fmt.Println("processing images")
	for {
		select {
		case <-ctx.Done():
			return syserr.Wrap(ctx.Err(), "context is done")
		default:
			res, err := s.imageQueueRepository.List(ctx, nil, database.ImageProcessingQueueListParameters{
				Filter: &database.ImageProcessingQueueFilter{
					IsFailed: lo.ToPtr(true),
					IsCompleted: lo.ToPtr(false),
				},
			})
			if err != nil {
				return syserr.Wrap(err, "could not list images")
			}
	
			s.bufferSize.Add(int64(len(res)))
	
			time.Sleep(time.Second)
		}
	}

	// todo: 
	// 1. get N images, save the buffer size
	// 2. feed the images to the workers
	// 3. when a worker is done, it decreases the buffer size
	// 4. here run an endless cycle and keep adding new items to the buffer if it drops below a certain threshold
	// 5. if no new images were found, exit the cycle
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) init(ctx context.Context) error {
	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not extract config")
	}

	workerPoolSize := config.Backend.Worker.ThreadCount
	if workerPoolSize < 1 {
		return syserr.NewInternal("invalid worker pool size", syserr.F("value", workerPoolSize))
	}

	for i := 0;	i < workerPoolSize; i++ {
		go s.processImage(ctx, i)
	}

	return nil
}

func (s *Service) processImage(ctx context.Context, workerId int) {

}