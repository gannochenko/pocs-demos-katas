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

func (s *Service) SubmitImageForProcessing(ctx context.Context, handle interfaces.SessionHandle, objectName string) (*domain.Image, error) {
	user := ctxUtil.GetUser(ctx)
	if user == nil {
		return nil, syserr.NewInternal("user is missing in the context")
	}

	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not load config")
	}

	handle, err = s.sessionManager.Begin(handle)
	if err != nil {
		return nil, syserr.Wrap(err, "could not start transaction")
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
		return nil, syserr.Wrap(err, "could not create image")
	}

	err = s.imageProcessingQueueRepository.Create(ctx, handle.GetTx(), &database.ImageProcessingQueue{
		CreatedBy: user.ID,
		ImageID:   imageID,
		IsFailed:  false,
	})
	if err != nil {
		return nil, syserr.Wrap(err, "could not create image processing queue element")
	}

	err = s.sessionManager.Commit(handle)
	if err != nil {
		return nil, syserr.Wrap(err, "could not commit transaction")
	}

	// todo: create message queue event

	images, err := s.imageRepository.List(ctx, nil, database.ImageListParameters{
		Filter: &database.ImageFilter{
			ID: &imageID,
		},
	})
	if err != nil {
		return nil, syserr.Wrap(err, "could not get the image", syserr.F("id", imageID.String()))
	}

	if len(images) == 0 {
		return nil, syserr.Wrap(err, "image not found", syserr.F("id", imageID.String()))
	}

	domainImage, err := images[0].ToDomain()
	if err != nil {
		return nil, syserr.Wrap(err, "could not convert image to domain", syserr.F("id", imageID.String()))
	}

	return domainImage, nil
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
