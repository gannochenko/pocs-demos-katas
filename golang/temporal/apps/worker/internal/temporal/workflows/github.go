package workflows

import (
	"worker/internal/domain"
	"worker/internal/interfaces"

	"go.temporal.io/sdk/workflow"
)

type GithubWorkflowGroup struct {
}

func NewGithubWorkflow() interfaces.TemporalWorkflowGroup {
	return &GithubWorkflowGroup{}
}

func (w *GithubWorkflowGroup) GenerateReportGithubWorkflow(ctx workflow.Context, input domain.GenerateReportGithubWorkflowInput) (domain.GenerateReportGithubWorkflowOutput, error) {
	return domain.GenerateReportGithubWorkflowOutput{}, nil
}

func (w *GithubWorkflowGroup) GetWorkflows() map[string]any {
	schema := make(map[string]any)

	schema[domain.GenerateReportGithubWorkflowName] = w.GenerateReportGithubWorkflow

	return schema
}
