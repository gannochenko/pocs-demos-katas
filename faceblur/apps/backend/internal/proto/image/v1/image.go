package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	imageV1 "backend/internal/proto/common/image/v1"
	pageNavigationV1 "backend/internal/proto/common/page_navigation/v1"
	v1 "backend/proto/common/page_navigation/v1"
	imagepb "backend/proto/image/v1"
)

func ValidateSubmitImageRequest(req *imagepb.SubmitImageRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Image, validation.Required, validation.By(imageV1.ValidateCreateImage)),
	)
}

func ValidateListImagesRequest(req *imagepb.ListImagesRequest) error {
	if req.PageNavigation == nil {
		req.PageNavigation = &v1.PageNavigationRequest{
			PageNumber: 1,
			PageSize:   pageNavigationV1.DefaultPageSize,
		}
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.PageNavigation, validation.Required, validation.By(pageNavigationV1.ValidatePageNavigationRequest)),
	)
}
