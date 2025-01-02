package v1

import (
	"context"

	"backend/interfaces"
	"backend/internal/util/syserr"
	imagepb "backend/proto/image/v1"
)

type ImageController struct {
	imagepb.UnimplementedImageServiceServer

	loggerService interfaces.LoggerService
	imageService  interfaces.ImageService
}

func NewImageController(loggerService interfaces.LoggerService, imageService interfaces.ImageService) *ImageController {
	return &ImageController{
		loggerService: loggerService,
		imageService:  imageService,
	}
}

func (c *ImageController) SubmitImage(ctx context.Context, request *imagepb.SubmitImageRequest) (*imagepb.SubmitImageResponse, error) {
	err := c.imageService.SubmitImageForProcessing(ctx, nil, request.Image.Url)
	if err != nil {
		return nil, syserr.Wrap(err, "could not submit image")
	}

	return &imagepb.SubmitImageResponse{
		Version: "v1",
	}, nil
}

func (c *ImageController) ListImages(ctx context.Context, request *imagepb.ListImagesRequest) (*imagepb.ListImagesResponse, error) {
	// c.loggerService.Info(ctx, "i am here")

	return nil, syserr.NewBadInput("very bad input", syserr.F("------foo------", "bar"))

	return &imagepb.ListImagesResponse{}, nil
}
