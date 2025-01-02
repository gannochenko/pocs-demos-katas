package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"backend/internal/util/syserr"
	v1 "backend/proto/common/image/v1"
)

func ValidateCreateImage(value interface{}) error {
	image, ok := value.(*v1.CreateImage)
	if !ok {
		return syserr.NewBadInput("invalid image type")
	}

	return validation.ValidateStruct(image,
		validation.Field(&image.Url, validation.Required, is.URL),
	)
}
