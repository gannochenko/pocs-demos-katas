package eventBus

import (
	"context"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	"backend/interfaces"
	"backend/internal/domain"
	protoEventConverterV1 "backend/internal/proto/entity/image/event/v1"
	"backend/internal/util/syserr"
	protoEventV1 "backend/proto/entity/image/event/v1"
)

type EventHandler = func(event *domain.EventBusEvent)

type Service struct {
	configService  interfaces.ConfigService
	loggerService  interfaces.LoggerService
	connection     *amqp091.Connection
	channel        *amqp091.Channel
	config         *domain.Config
	eventListeners map[domain.EventBusEventType][]EventHandler // todo: make this map thread-safe
}

func NewEventBusService(configService interfaces.ConfigService, loggerService interfaces.LoggerService) *Service {
	return &Service{
		configService:  configService,
		loggerService:  loggerService,
		eventListeners: make(map[domain.EventBusEventType][]EventHandler),
	}
}

func (s *Service) Start(ctx context.Context) error {
	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not extract config")
	}

	s.config = config

	conn, err := amqp091.Dial(fmt.Sprintf("amqp://guest:guest@%s:%d/", config.RabbitMq.Host, config.RabbitMq.Port))
	if err != nil {
		return syserr.Wrap(err, "could not open connection")
	}

	s.connection = conn

	s.channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	err = s.consumeMessages(ctx, s.channel)
	if err != nil {
		return syserr.Wrap(err, "could not start consuming messages")
	}

	return nil
}

func (s *Service) Stop() error {
	if s.connection != nil {
		err := s.connection.Close()
		if err != nil {
			return syserr.Wrap(err, "could not close connection")
		}
	}

	if s.channel != nil {
		err := s.channel.Close()
		if err != nil {
			return syserr.Wrap(err, "could not close channel")
		}
	}

	return nil
}

func (s *Service) AddEventListener(eventType domain.EventBusEventType, cb EventHandler) error {
	_, ok := s.eventListeners[eventType]
	if !ok {
		s.eventListeners[eventType] = make([]EventHandler, 0)
	}

	s.eventListeners[eventType] = append(s.eventListeners[eventType], cb)

	return nil
}

func (s *Service) TriggerEvent(event *domain.EventBusEvent) error {
	if s.channel == nil {
		return syserr.NewInternal("client disconnected")
	}

	headers := amqp091.Table{}
	headers["eventType"] = string(event.Type)

	protoEvent, err := protoEventConverterV1.ConvertEventToProto(event)
	if err != nil {
		return syserr.Wrap(err, "could not convert payload to proto")
	}

	msgBytes, err := proto.Marshal(protoEvent)
	if err != nil {
		return syserr.Wrap(err, "could not marshal an event payload")
	}

	err = s.channel.Publish(
		s.config.RabbitMq.EventBus.ExchangeName,
		s.config.RabbitMq.EventBus.RoutingKey,
		false,
		false,
		amqp091.Publishing{
			Headers:     headers,
			ContentType: "application/octet-stream",
			Body:        msgBytes,
		},
	)
	if err != nil {
		return syserr.Wrap(err, "could not publish an event")
	}

	return nil
}

func (s *Service) consumeMessages(ctx context.Context, ch *amqp091.Channel) error {
	msgs, err := ch.Consume(
		s.config.RabbitMq.EventBus.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-msgs:
			if !ok {
				return nil
			}

			var event protoEventV1.Event
			err = proto.Unmarshal(msg.Body, &event)
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "could not unmarshal event bus message"))
				continue
			}

			domainEvent, err := protoEventConverterV1.ConvertEventToDomain(&event)
			if err != nil {
				s.loggerService.LogError(ctx, syserr.Wrap(err, "could not convert event to domain"))
				continue
			}

			listeners, ok := s.eventListeners[domainEvent.Type]
			if ok && len(listeners) > 0 {
				for _, listener := range listeners {
					listener(domainEvent)
				}
			}
		}
	}
}
