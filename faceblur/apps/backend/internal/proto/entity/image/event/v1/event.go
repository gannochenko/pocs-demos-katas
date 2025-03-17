package v1

import (
	"github.com/google/uuid"

	"backend/internal/domain"
	"backend/internal/util/syserr"
	protoEventV1 "backend/proto/entity/image/event/v1"
	protoPayloadV1 "backend/proto/entity/image/payload/v1"
)

func ConvertEventToDomain(event *protoEventV1.Event) (*domain.EventBusEvent, error) {
	payloadImageCreated := event.GetImageCreated()
	if payloadImageCreated != nil {
		imageID, err := uuid.Parse(payloadImageCreated.ImageId)
		if err != nil {
			return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not convert image id to uuid")
		}

		return &domain.EventBusEvent{
			Type: domain.EventBusEventTypeImageCreated,
			Payload: &domain.EventBusEventPayloadImageCreated{
				ImageID: imageID,
			},
		}, nil
	}

	payloadImageProcessed := event.GetImageProcessed()
	if payloadImageProcessed != nil {
		imageID, err := uuid.Parse(payloadImageProcessed.ImageId)
		if err != nil {
			return nil, syserr.WrapAs(err, syserr.BadInputCode, "could not convert image id to uuid")
		}

		return &domain.EventBusEvent{
			Type: domain.EventBusEventTypeImageProcessed,
			Payload: &domain.EventBusEventPayloadImageProcessed{
				ImageID: imageID,
				Failed:  payloadImageProcessed.Failed,
				CreatorID:  payloadImageProcessed.CreatorId,
			},
		}, nil
	}

	return nil, syserr.NewBadInput("unknown message received")
}

func ConvertEventToProto(event *domain.EventBusEvent) (*protoEventV1.Event, error) {
	protoEvent := &protoEventV1.Event{}

	switch event.Type {
	case domain.EventBusEventTypeImageCreated:
		domainPayload, ok := event.Payload.(*domain.EventBusEventPayloadImageCreated)
		if !ok {
			return nil, syserr.NewInternal("payload is not of type EventBusEventPayloadImageCreated")
		}

		protoEvent.Payload = &protoEventV1.Event_ImageCreated{
			ImageCreated: &protoPayloadV1.ImageCreatedPayload{
				ImageId: domainPayload.ImageID.String(),
			},
		}
	case domain.EventBusEventTypeImageProcessed:
		domainPayload, ok := event.Payload.(*domain.EventBusEventPayloadImageProcessed)
		if !ok {
			return nil, syserr.NewInternal("payload is not of type EventBusEventPayloadImageProcessed")
		}

		protoEvent.Payload = &protoEventV1.Event_ImageProcessed{
			ImageProcessed: &protoPayloadV1.ImageProcessedPayload{
				ImageId: domainPayload.ImageID.String(),
				Failed:  domainPayload.Failed,
				CreatorId: domainPayload.CreatorID,
			},
		}
	}

	return protoEvent, nil
}
