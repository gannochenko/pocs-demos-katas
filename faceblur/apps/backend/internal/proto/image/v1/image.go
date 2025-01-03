package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"backend/internal/domain"
	imageV1 "backend/internal/proto/common/image/v1"
	pageNavigationV1 "backend/internal/proto/common/page_navigation/v1"
	imagebp "backend/proto/common/image/v1"
	pageNavigationbp "backend/proto/common/page_navigation/v1"
	imageServicepb "backend/proto/image/v1"
)

func ValidateSubmitImageRequest(req *imageServicepb.SubmitImageRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Image, validation.Required, validation.By(imageV1.ValidateCreateImage)),
	)
}

func ValidateListImagesRequest(req *imageServicepb.ListImagesRequest) error {
	if req.PageNavigation == nil {
		req.PageNavigation = &pageNavigationbp.PageNavigationRequest{
			PageNumber: 1,
			PageSize:   pageNavigationV1.DefaultPageSize,
		}
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.PageNavigation, validation.Required, validation.By(pageNavigationV1.ValidatePageNavigationRequest)),
	)
}

func ConvertListImagesRequestToDomain(request *imageServicepb.ListImagesRequest) *domain.ListImagesRequest {
	return &domain.ListImagesRequest{
		PageNavigation: *pageNavigationV1.ConvertRequestToDomain(request.PageNavigation),
	}
}

func ConvertListImagesResponseToProto(response *domain.ListImagesResponse) *imageServicepb.ListImagesResponse {
	var images []*imagebp.Image
	for _, image := range response.Images {
		images = append(images, imageV1.ConvertImageToProto(&image))
	}

	return &imageServicepb.ListImagesResponse{
		Version:        "v1",
		Images:         images,
		PageNavigation: pageNavigationV1.ConvertResponseToProto(&response.PageNavigation),
	}
}
