package v1

import (
	"backend/internal/domain"
	v1 "backend/proto/entity/image/event/v1"
)

func ConvertEventToDomain(event *v1.Event) *domain.EventBusEvent {
	return &domain.EventBusEvent{}
}

func ConvertEventFromProto(event *domain.EventBusEvent) *v1.Event {
	return &v1.Event{}
}
