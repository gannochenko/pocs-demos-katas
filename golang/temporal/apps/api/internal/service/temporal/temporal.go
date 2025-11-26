package temporal

import (
	"context"

	"go.temporal.io/sdk/client"
)

type Service struct {
	client client.Client
}

func NewService(client client.Client) *Service {
	return &Service{client: client}
}

func (s *Service) StartWorkflow(ctx context.Context, workflowName string, workflowID string, workflowInput any) error {
	return nil
}

func (s *Service) StopWorkflow(ctx context.Context, workflowID string) error {
	return nil
}
