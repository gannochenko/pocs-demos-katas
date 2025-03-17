package domain

import "github.com/google/uuid"

type EventBusEventType string

const (
	EventBusEventTypeImageCreated EventBusEventType = "image_created"
	EventBusEventTypeImageProcessed EventBusEventType = "image_processed"
)

type EventBusEventPayloadImageCreated struct {
	ImageID uuid.UUID
}

type EventBusEventPayloadImageProcessed struct {
	ImageID uuid.UUID
	Failed bool
	CreatorID string
}

type EventBusEvent struct {
	Type    EventBusEventType
	Payload any
}
