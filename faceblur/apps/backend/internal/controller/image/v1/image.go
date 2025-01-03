package v1

import (
	"context"

	"backend/interfaces"
	v1 "backend/internal/proto/image/v1"
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

func (c *ImageController) GetUploadURL(ctx context.Context, _ *imagepb.GetUploadURLRequest) (*imagepb.GetUploadURLResponse, error) {
	return &imagepb.GetUploadURLResponse{
		Version: "v1",
	}, nil
}

func (c *ImageController) SubmitImage(ctx context.Context, request *imagepb.SubmitImageRequest) (*imagepb.SubmitImageResponse, error) {
	err := v1.ValidateSubmitImageRequest(request)
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "incorrect input")
	}

	err = c.imageService.SubmitImageForProcessing(ctx, nil, request.Image.Url)
	if err != nil {
		return nil, syserr.Wrap(err, "could not submit image")
	}

	return &imagepb.SubmitImageResponse{
		Version: "v1",
	}, nil
}

func (c *ImageController) ListImages(ctx context.Context, request *imagepb.ListImagesRequest) (*imagepb.ListImagesResponse, error) {
	err := v1.ValidateListImagesRequest(request)
	if err != nil {
		return nil, syserr.WrapAs(err, syserr.BadInputCode, "incorrect input")
	}

	return &imagepb.ListImagesResponse{}, nil
}
