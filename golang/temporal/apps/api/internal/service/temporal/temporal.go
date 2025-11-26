package temporal

import (
	"api/internal/domain"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

type Service struct {
	client client.Client
}

func NewService(client client.Client) *Service {
	return &Service{client: client}
}

func (s *Service) ExecuteWorkflow(ctx context.Context, workflowName string, workflowID string, workflowInput any, executeAt *time.Time) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	
	delay := time.Duration(0)
	if executeAt != nil {
		delay = time.Until(*executeAt)
	}

	options := client.StartWorkflowOptions{
		ID: workflowID,
		TaskQueue: domain.WorkflowsTaskQueue,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING,
		StartDelay: delay,
	}
	_, err := s.client.ExecuteWorkflow(ctx, options, workflowName, workflowInput)
	if err != nil {
		return errors.Wrap(err, "failed to execute workflow")
	}
	return nil
}

func (s *Service) CancelWorkflow(ctx context.Context, workflowID string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	err := s.client.CancelWorkflow(ctx, workflowID, "")
	if err != nil {
		return errors.Wrap(err, "failed to cancel workflow")
	}
	return nil
}
