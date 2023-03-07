package main

import (
	"context"
	"log"
	"os"

	googlePubSub "cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func main() {
	projectID := os.Getenv("PROJECT_ID")
	topicID := os.Getenv("TOPIC_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION_ID")
	credentialsFile := os.Getenv("CREDENTIALS_FILE")

	credentials, err := os.ReadFile(credentialsFile)

	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	pubSubClient, err := googlePubSub.NewClient(ctx, projectID, option.WithCredentialsJSON(credentials))
	if err != nil {
		panic(err)
	}

	topic := pubSubClient.Topic(topicID)

	attributes := map[string]string{
		"Hello": "There",
	}

	message := &googlePubSub.Message{
		Data:       []byte("Hello!"),
		Attributes: attributes,
	}

	publishResult := topic.Publish(ctx, message)
	messageID, err := publishResult.Get(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("Message was published with id %s", messageID)

	subscription := pubSubClient.Subscription(subscriptionID)

	err = subscription.Receive(ctx, func(receiveCtx context.Context, message *googlePubSub.Message) {
		log.Printf("Message received %s", string(message.Data))
		message.Ack()

		cancelCtx()
	})

	log.Printf("Done")

	if err != nil {
		panic(err)
	}
}
