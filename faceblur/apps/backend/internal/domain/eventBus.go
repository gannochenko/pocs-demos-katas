package domain

type EventBusEventType string

const (
	EventBusEventTypeImageCreated EventBusEventType = "image_created"
)

type EventBusEvent struct {
	Type    EventBusEventType
	Payload string
}
