package temporal

import "go.temporal.io/sdk/client"

type Service struct {
	client client.Client
}

func NewService(client client.Client) *Service {
	return &Service{client: client}
}
