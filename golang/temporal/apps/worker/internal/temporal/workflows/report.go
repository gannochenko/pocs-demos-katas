package workflows

import (
	"time"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type ReportWorkflowGroup struct {
}

func NewReportWorkflowGroup() interfaces.TemporalWorkflowGroup {
	return &ReportWorkflowGroup{}
}

func (w *ReportWorkflowGroup) GenerateReportGithubWorkflow(ctx workflow.Context, input domain.GenerateReportGithubWorkflowInput) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	})

	var repoStats domain.GetRepositoryStatsActivityOutput
	err := workflow.ExecuteActivity(ctx, domain.GetRepositoryStatsActivityName, domain.GetRepositoryStatsActivityInput{
		Repository: input.Repository,
	}).Get(ctx, &repoStats)
	if err != nil {
		return errors.Wrap(err, "failed to execute activity")
	}

	var summary domain.MakeHumanReadableSummaryActivityOutput
	err = workflow.ExecuteActivity(ctx, domain.MakeHumanReadableSummaryActivityName, domain.MakeHumanReadableSummaryActivityInput{
		Text: repoStats.ToText(),
	}).Get(ctx, &summary)
	if err != nil {
		return errors.Wrap(err, "failed to execute activity")
	}

	err = workflow.Sleep(ctx, 5*time.Minute)
	if err != nil {
		return errors.Wrap(err, "failed to sleep")
	}

	return workflow.NewContinueAsNewError(ctx, w.GenerateReportGithubWorkflow, input)
}

func (w *ReportWorkflowGroup) GetWorkflows() map[string]any {
	schema := make(map[string]any)

	schema[domain.GenerateReportGithubWorkflowName] = w.GenerateReportGithubWorkflow

	return schema
}
