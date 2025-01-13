package image

import (
	"context"

	"github.com/google/uuid"

	"backend/interfaces"
	"backend/internal/database"
	"backend/internal/domain"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/syserr"
)

type Service struct {
	sessionManager                 interfaces.SessionManager
	imageRepository                interfaces.ImageRepository
	imageProcessingQueueRepository interfaces.ImageProcessingQueueRepository
	loggerService                  interfaces.LoggerService
	storageService                 interfaces.StorageService
	configService                  interfaces.ConfigService
}

func NewImageService(
	sessionManager interfaces.SessionManager,
	imageRepository interfaces.ImageRepository,
	imageProcessingQueueRepository interfaces.ImageProcessingQueueRepository,
	storageService interfaces.StorageService,
	configService interfaces.ConfigService,
) *Service {
	return &Service{
		sessionManager:                 sessionManager,
		imageRepository:                imageRepository,
		imageProcessingQueueRepository: imageProcessingQueueRepository,
		storageService:                 storageService,
		configService:                  configService,
	}
}

func (s *Service) SubmitImageForProcessing(ctx context.Context, handle interfaces.SessionHandle, objectName string) error {
	user := ctxUtil.GetUser(ctx)
	if user == nil {
		return syserr.NewInternal("user is missing in the context")
	}

	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not load config")
	}

	handle, err = s.sessionManager.Begin(handle)
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
		OriginalURL: s.storageService.GetPublicURL(config.Storage.ImageBucketName, objectName),
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

	// todo: create message queue event

	return nil
}

func (s *Service) ListImages(ctx context.Context, _ interfaces.SessionHandle, request *domain.ListImagesRequest) (*domain.ListImagesResponse, error) {
	user := ctxUtil.GetUser(ctx)
	if user == nil {
		return nil, syserr.NewInternal("user is missing in the context")
	}

	filter := &database.ImageFilter{
		CreatedBy: &user.ID,
	}

	response := &domain.ListImagesResponse{}

	count, err := s.imageRepository.Count(ctx, nil, database.ImageCountParameters{
		Filter: filter,
	})
	if err != nil {
		return nil, syserr.Wrap(err, "could not get image count")
	}

	response.PageNavigation = *domain.NewPageNavigationResponseFromRequest(&request.PageNavigation, count)

	if count > 0 {
		images, err := s.imageRepository.List(ctx, nil, database.ImageListParameters{
			Filter: filter,
		})
		if err != nil {
			return nil, syserr.Wrap(err, "could not get image list")
		}

		for _, image := range images {
			domainImage, err := image.ToDomain()
			if err != nil {
				return nil, syserr.Wrap(err, "could not convert image to domain", syserr.F("image_id", image.ID))
			}

			response.Images = append(response.Images, *domainImage)
		}
	}

	return response, nil
}
