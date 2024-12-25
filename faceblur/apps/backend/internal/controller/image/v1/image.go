package v1

import (
	"context"

	"backend/interfaces"
	imagepb "backend/proto/image/v1"
)

type ImageController struct {
	imagepb.UnimplementedImageServiceServer

	loggerService interfaces.LoggerService
}

func NewImageController(loggerService interfaces.LoggerService) *ImageController {
	return &ImageController{
		loggerService: loggerService,
	}
}

func (c *ImageController) ListImages(ctx context.Context, request *imagepb.ListImagesRequest) (*imagepb.ListImagesResponse, error) {
	c.loggerService.Info(ctx, "i am here")

	return &imagepb.ListImagesResponse{}, nil
}
