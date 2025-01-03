package v1

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"backend/internal/domain"
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

func ConvertRequestToDomain(request *v1.PageNavigationRequest) *domain.PageNavigationRequest {
	return &domain.PageNavigationRequest{
		PageNumber: request.PageNumber,
		PageSize:   request.PageSize,
	}
}

func ConvertResponseToProto(response *domain.PageNavigationResponse) *v1.PageNavigationResponse {
	return &v1.PageNavigationResponse{
		PageNumber: response.PageNumber,
		PageSize:   response.PageSize,
		PageCount:  response.PageCount,
		Total:      response.Total,
	}
}
