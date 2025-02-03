package imageProcessor

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"backend/interfaces"
	"backend/internal/domain"
	"backend/internal/util/logger"
	"backend/internal/util/syserr"
)

type Service struct {
	eventBusService interfaces.EventBusService
	loggerService   interfaces.LoggerService

	newMessages atomic.Bool
}

func NewImageProcessor(eventBusService interfaces.EventBusService, loggerService interfaces.LoggerService) *Service {
	result := &Service{
		eventBusService: eventBusService,
		loggerService:   loggerService,
	}

	result.newMessages.Store(false)

	return result
}

func (s *Service) Start(ctx context.Context) error {
	err := s.eventBusService.AddEventListener(domain.EventBusEventTypeImageCreated, func(event *domain.EventBusEvent) {
		s.loggerService.Info(ctx, "new event received", logger.F("event", event))
		s.newMessages.Store(true)
	})
	if err != nil {
		return syserr.Wrap(err, "could not start listening to events")
	}

	select {
	case <-ctx.Done():
		return syserr.Wrap(ctx.Err(), "context is done")
	default:
		if s.newMessages.Swap(false) {
			fmt.Println("Flag was true, reacting now!")
			// get the new unprocessed messages
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (s *Service) Stop() error {
	return nil
}
