package eventBus

import (
	"context"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	"backend/interfaces"
	"backend/internal/domain"
	v1 "backend/internal/proto/entity/image/event/v1"
	"backend/internal/util/syserr"
)

type Service struct {
	configService interfaces.ConfigService
	connection    *amqp091.Connection
	channel       *amqp091.Channel
	queueName     string
}

func NewEventBusService(configService interfaces.ConfigService) *Service {
	return &Service{
		configService: configService,
	}
}

func (s *Service) Start(ctx context.Context) error {
	config, err := s.configService.GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not extract config")
	}

	conn, err := amqp091.Dial(fmt.Sprintf("amqp://guest:guest@%s:%d/", config.RabbitMq.Host, config.RabbitMq.Port))
	if err != nil {
		return syserr.Wrap(err, "could not open connection")
	}

	s.connection = conn

	s.channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	err = s.consumeMessages(ctx, s.channel, config.RabbitMq.QueueName)
	if err != nil {
		return syserr.Wrap(err, "could not start consuming messages")
	}

	s.queueName = config.RabbitMq.QueueName

	return nil
}

func (s *Service) Stop() error {
	err := s.connection.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close connection")
	}

	err = s.channel.Close()
	if err != nil {
		return syserr.Wrap(err, "could not close channel")
	}

	return nil
}

func (s *Service) AddEventListener(eventType domain.EventBusEventType, cb func(payload []byte)) error {
	return nil
}

func (s *Service) TriggerEvent(event *domain.EventBusEvent) error {
	headers := map[string]any{}
	headers["eventType"] = event.Type

	protoEvent := v1.ConvertEventFromProto(event)

	msgBytes, err := proto.Marshal(protoEvent)
	if err != nil {
		return syserr.Wrap(err, "could not marshal an event payload")
	}

	err = s.channel.Publish(
		"",
		s.queueName,
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

func (s *Service) consumeMessages(ctx context.Context, ch *amqp091.Channel, queueName string) error {
	msgs, err := ch.Consume(
		queueName,
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

			// unmarshall body
			log.Printf("Received a message: %s", msg.Body)
		}
	}
}
