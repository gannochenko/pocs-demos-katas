package v1

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	v1 "backend/proto/common/page_navigation/v1"
)

const (
	MaxPageSize     = 200
	DefaultPageSize = 50
)

func ValidatePageNavigationRequest(value interface{}) error {
	pageNav, ok := value.(*v1.PageNavigationRequest)
	if !ok {
		return fmt.Errorf("invalid PageNavigationRequest type")
	}

	return validation.ValidateStruct(pageNav,
		validation.Field(&pageNav.PageSize, validation.Required, validation.Min(1), validation.Max(MaxPageSize)),
		validation.Field(&pageNav.PageNumber, validation.Required, validation.Min(1)),
	)
}
