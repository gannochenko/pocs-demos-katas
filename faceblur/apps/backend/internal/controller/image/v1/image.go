package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"backend/interfaces"
	imageV1 "backend/internal/proto/common/image/v1"
	v1 "backend/internal/proto/image/v1"
	"backend/internal/util/syserr"
	imagepb "backend/proto/image/v1"
)

const (
	Version      = "v1"
	SignedURLTTL = time.Hour
)

type ImageController struct {
	imagepb.UnimplementedImageServiceServer

	loggerService  interfaces.LoggerService
	imageService   interfaces.ImageService
	storageService interfaces.StorageService
	configService  interfaces.ConfigService
}

func NewImageController(
	loggerService interfaces.LoggerService,
	imageService interfaces.ImageService,
	storageService interfaces.StorageService,
	configService interfaces.ConfigService,
) *ImageController {
	return &ImageController{
		loggerService:  loggerService,
		imageService:   imageService,
		storageService: storageService,
		configService:  configService,
	}
}

// GetUploadURL implements POST /v1/image/upload-url/get
func (c *ImageController) GetUploadURL(ctx context.Context, _ *imagepb.GetUploadURLRequest) (*imagepb.GetUploadURLResponse, error) {
	config, err := c.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not load config")
	}

	objectName := uuid.New().String()

	fileURL, err := c.storageService.PrepareSignedURL(ctx, config.Storage.ImageBucketName, objectName, SignedURLTTL, http.MethodPut, "application/octet-stream")
	if err != nil {
		return nil, syserr.Wrap(err, "could not create signed url")
	}

	return &imagepb.GetUploadURLResponse{
		Version:    Version,
		Url:        fileURL,
		ObjectName: objectName,
	}, nil
}

// SubmitImage implements POST /v1/image/submit
func (c *ImageController) SubmitImage(ctx context.Context, request *imagepb.SubmitImageRequest) (*imagepb.SubmitImageResponse, error) {
	err := v1.ValidateSubmitImageRequest(request)
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "incorrect input")
	}

	image, err := c.imageService.SubmitImageForProcessing(ctx, nil, request.Image.ObjectName)
	if err != nil {
		return nil, syserr.Wrap(err, "could not submit image")
	}

	return &imagepb.SubmitImageResponse{
		Version: Version,
		Image:   imageV1.ConvertImageToProto(image),
	}, nil
}

// ListImages implements POST /v1/image/list
func (c *ImageController) ListImages(ctx context.Context, request *imagepb.ListImagesRequest) (*imagepb.ListImagesResponse, error) {
	err := v1.ValidateListImagesRequest(request)
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "incorrect input")
	}

	response, err := c.imageService.ListImages(ctx, nil, v1.ConvertListImagesRequestToDomain(request))
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not get list if images")
	}

	return v1.ConvertListImagesResponseToProto(response), nil
}
