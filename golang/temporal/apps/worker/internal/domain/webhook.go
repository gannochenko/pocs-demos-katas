package domain

type WebhookEvent struct {
	EventID        string
	EventTimestamp string
	EventType      string
	Payload        map[string]interface{}
}
