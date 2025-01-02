package image

import (
	"context"

	"github.com/google/uuid"

	"backend/interfaces"
	"backend/internal/database"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/syserr"
)

type Service struct {
	sessionManager                 interfaces.SessionManager
	imageRepository                interfaces.ImageRepository
	imageProcessingQueueRepository interfaces.ImageProcessingQueueRepository
	loggerService                  interfaces.LoggerService
}

func NewImageService(sessionManager interfaces.SessionManager, imageRepository interfaces.ImageRepository, imageProcessingQueueRepository interfaces.ImageProcessingQueueRepository) *Service {
	return &Service{
		sessionManager:                 sessionManager,
		imageRepository:                imageRepository,
		imageProcessingQueueRepository: imageProcessingQueueRepository,
	}
}

func (s *Service) SubmitImageForProcessing(ctx context.Context, handle interfaces.SessionHandle, url string) error {
	user := ctxUtil.GetUser(ctx)
	if user == nil {
		return syserr.NewInternal("user is missing in the context")
	}

	handle, err := s.sessionManager.Begin(handle)
	if err != nil {
		return syserr.Wrap(err, "could not start transaction")
	}
	defer func() {
		err = s.sessionManager.RollbackUnlessCommitted(handle)
		if err != nil {
			s.loggerService.LogError(ctx, syserr.Wrap(err, "could not rollback transaction"))
		}
	}()

	imageID := uuid.New()
	err = s.imageRepository.Create(ctx, handle.GetTx(), &database.Image{
		ID:          imageID,
		CreatedBy:   user.ID,
		OriginalURL: url,
		IsProcessed: false,
	})
	if err != nil {
		return syserr.Wrap(err, "could not create image")
	}

	err = s.imageProcessingQueueRepository.Create(ctx, handle.GetTx(), &database.ImageProcessingQueue{
		CreatedBy: user.ID,
		ImageID:   imageID,
		IsFailed:  false,
	})
	if err != nil {
		return syserr.Wrap(err, "could not create image processing queue element")
	}

	err = s.sessionManager.Commit(handle)
	if err != nil {
		return syserr.Wrap(err, "could not commit transaction")
	}

	// create message queue event

	return nil
}
