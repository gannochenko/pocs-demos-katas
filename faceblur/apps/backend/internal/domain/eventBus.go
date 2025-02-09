package domain

import "github.com/google/uuid"

type EventBusEventType string

const (
	EventBusEventTypeImageCreated EventBusEventType = "image_created"
)

type EventBusEventPayloadImageCreated struct {
	ImageID uuid.UUID
}

type EventBusEvent struct {
	Type    EventBusEventType
	Payload any
}
