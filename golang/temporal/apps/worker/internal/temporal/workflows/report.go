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
	var repoStats domain.GetRepositoryStatsActivityOutput
	err := workflow.ExecuteActivity(w.withSafeguard(ctx, 10*time.Second, 5), domain.GetRepositoryStatsActivityName, domain.GetRepositoryStatsActivityInput{
		Repository: input.Repository,
	}).Get(ctx, &repoStats)
	if err != nil {
		return errors.Wrap(err, "failed to execute activity")
	}

	if len(repoStats.Commits) > 0 {
		var summary domain.MakeHumanReadableSummaryActivityOutput
		err = workflow.ExecuteActivity(w.withSafeguard(ctx, 3*time.Minute, 3), domain.MakeHumanReadableSummaryActivityName, domain.MakeHumanReadableSummaryActivityInput{
			Text: repoStats.ToText(),
		}).Get(ctx, &summary)
		if err != nil {
			return errors.Wrap(err, "failed to execute activity")
		}

		var postToSlackOutput domain.PostToSlackActivityOutput
		err = workflow.ExecuteActivity(w.withSafeguard(ctx, 10*time.Second, 5), domain.PostToSlackActivityName, domain.PostToSlackActivityInput{
			Repository: input.Repository,
			Summary: summary.Summary,
		}).Get(ctx, &postToSlackOutput)
		if err != nil {
			return errors.Wrap(err, "failed to execute activity")
		}
	}

	err = workflow.Sleep(ctx, w.getDurationToNext9AM(workflow.Now(ctx)))
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

func (w *ReportWorkflowGroup) getDurationToNext9AM(now time.Time) time.Duration {
	next9AM := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())

	if now.Hour() >= 9 {
		next9AM = next9AM.Add(24 * time.Hour)
	}

	return next9AM.Sub(now)
}

func (w *ReportWorkflowGroup) withSafeguard(ctx workflow.Context, timeout time.Duration, attempts int32) workflow.Context {
	return workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: timeout,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: attempts,
		},
	})
}