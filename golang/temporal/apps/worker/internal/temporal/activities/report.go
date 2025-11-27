package activities

import (
	"context"
	"time"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"github.com/google/go-github/v62/github"
	"github.com/samber/lo"
)

type ReportActivityGroup struct {
	config *domain.Config
	githubClient interfaces.GitHubClient
	openaiClient interfaces.OpenAIClient
	slackClient interfaces.SlackClient
}

func NewReportActivityGroup(config *domain.Config, githubClient interfaces.GitHubClient, openaiClient interfaces.OpenAIClient, slackClient interfaces.SlackClient) interfaces.TemporalActivityGroup {
	return &ReportActivityGroup{
		config: config,
		githubClient: githubClient,
		openaiClient: openaiClient,
		slackClient: slackClient,
	}
}

func (a *ReportActivityGroup) GetRepositoryStatsActivity(ctx context.Context, input domain.GetRepositoryStatsActivityInput) (domain.GetRepositoryStatsActivityOutput, error) {
	commits, err := a.githubClient.FetchCommits(ctx, input.Repository)
	if err != nil {
		return domain.GetRepositoryStatsActivityOutput{}, err
	}

	commitItems := make([]domain.Commit, len(commits))

	for i, commit := range commits {
		comm := commit.Commit

		author := lo.FromPtrOr(comm.Author, github.CommitAuthor{})

		commitItems[i] = domain.Commit{
			SHA: lo.FromPtrOr(comm.SHA, ""),
			Message: lo.FromPtrOr(comm.Message, ""),
			AuthorName: lo.FromPtrOr(author.Name, ""),
			AuthorEmail: lo.FromPtrOr(author.Email, ""),
			Date: lo.FromPtrOr(author.Date, github.Timestamp{}).Format(time.RFC3339),
		}
	}

	return domain.GetRepositoryStatsActivityOutput{
		Commits: commitItems,
	}, nil
}

func (a *ReportActivityGroup) MakeHumanReadableSummaryActivity(ctx context.Context, input domain.MakeHumanReadableSummaryActivityInput) (domain.MakeHumanReadableSummaryActivityOutput, error) {
	summary, err := a.openaiClient.Complete(
		ctx,
		"You are a helpful assistant that summarizes text. The text is a list of commits from a repository. Return a summary of the commits in a human readable format and in markdown. Dont use ordered lists for formatting. Use emoji to make lists. Make it sound exciting! Don't mention every commit, just the most important ones. Conduct the generic yet precise summary. Menion the champion contributor(s) of this round.",
		input.Text,
	)
	if err != nil {
		return domain.MakeHumanReadableSummaryActivityOutput{}, err
	}

	return domain.MakeHumanReadableSummaryActivityOutput{
		Summary: summary,
	}, nil
}

func (a *ReportActivityGroup) PostToSlackActivity(ctx context.Context, input domain.PostToSlackActivityInput) (domain.PostToSlackActivityOutput, error) {
	err := a.slackClient.SendMessage(ctx, a.config.Slack.Channel, input.Summary)
	if err != nil {
		return domain.PostToSlackActivityOutput{}, err
	}

	return domain.PostToSlackActivityOutput{}, nil
}

func (a *ReportActivityGroup) GetActivities() map[string]any {
	schema := make(map[string]any)

	schema[domain.GetRepositoryStatsActivityName] = a.GetRepositoryStatsActivity
	schema[domain.MakeHumanReadableSummaryActivityName] = a.MakeHumanReadableSummaryActivity
	schema[domain.PostToSlackActivityName] = a.PostToSlackActivity

	return schema
}
