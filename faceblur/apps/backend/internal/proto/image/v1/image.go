package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	v1 "backend/internal/proto/common/image/v1"
	imagepb "backend/proto/image/v1"
)

func ValidateSubmitImageRequest(req *imagepb.SubmitImageRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Image, validation.Required, validation.By(v1.ValidateCreateImage)),
	)
}
