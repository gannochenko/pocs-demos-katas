package v1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/samber/lo"
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
		validation.Field(&image.Url, validation.Required, is.URL),
	)
}

func ConvertImageToProto(image *domain.Image) *imagebp.Image {
	return &imagebp.Image{
		Id:          image.ID.String(),
		Url:         lo.Ternary(image.URL == nil, image.OriginalURL, *image.URL),
		IsProcessed: image.IsProcessed,
		IsFailed:    false, // todo: ??
		CreatedAt:   timestamppb.New(image.CreatedAt),
		UpdatedAt:   timestamppb.New(image.UpdatedAt),
	}
}
