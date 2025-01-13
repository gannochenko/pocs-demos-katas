package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/protobuf/types/known/timestamppb"

	"backend/internal/domain"
	"backend/internal/util/syserr"
	imagebp "backend/proto/common/image/v1"
	v1 "backend/proto/common/image/v1"
)

func ValidateCreateImage(value interface{}) error {
	image, ok := value.(*v1.CreateImage)
	if !ok {
		return syserr.NewBadInput("invalid image type")
	}

	return validation.ValidateStruct(image,
		validation.Field(&image.ObjectName, validation.Required, is.UUID),
	)
}

func ConvertImageToProto(image *domain.Image) *imagebp.Image {
	url := image.OriginalURL
	if image.URL != nil {
		url = *image.URL
	}

	return &imagebp.Image{
		Id:          image.ID.String(),
		Url:         url,
		IsProcessed: image.IsProcessed,
		IsFailed:    image.IsFailed,
		CreatedAt:   timestamppb.New(image.CreatedAt),
		UpdatedAt:   timestamppb.New(image.UpdatedAt),
	}
}
