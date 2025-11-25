package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(address string) (*grpc.ClientConn, error) {
	// Define retry policy with exponential backoff
	retryPolicy := `{
		"methodConfig": [{
			"name": [{"service": "storage.task.v1.TaskService"}],
			"retryPolicy": {
				"MaxAttempts": 4,
				"InitialBackoff": "0.1s",
				"MaxBackoff": "5s",
				"BackoffMultiplier": 2.0,
				"RetryableStatusCodes": [ "UNAVAILABLE", "DEADLINE_EXCEEDED", "RESOURCE_EXHAUSTED" ]
			}
		}]
	}`

	// Create connection with retry policy
	clientConn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // todo: add tls
		grpc.WithDefaultServiceConfig(retryPolicy),
	)
	if err != nil {
		return nil, err
	}

	return clientConn, nil
}
