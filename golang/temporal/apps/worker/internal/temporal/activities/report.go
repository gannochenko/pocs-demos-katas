package activities

import (
	"context"
	"fmt"
	"time"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"github.com/google/go-github/v62/github"
	"github.com/samber/lo"
)

type ReportActivityGroup struct {
	githubClient interfaces.GitHubClient
}

func NewReportActivityGroup(githubClient interfaces.GitHubClient) interfaces.TemporalActivityGroup {
	return &ReportActivityGroup{
		githubClient: githubClient,
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
	fmt.Println(input.Text)
	return domain.MakeHumanReadableSummaryActivityOutput{}, nil
}

func (a *ReportActivityGroup) PostToSlackActivity(ctx context.Context, input domain.PostToSlackActivityInput) (domain.PostToSlackActivityOutput, error) {
	return domain.PostToSlackActivityOutput{}, nil
}

func (a *ReportActivityGroup) GetActivities() map[string]any {
	schema := make(map[string]any)

	schema[domain.GetRepositoryStatsActivityName] = a.GetRepositoryStatsActivity
	schema[domain.MakeHumanReadableSummaryActivityName] = a.MakeHumanReadableSummaryActivity
	schema[domain.PostToSlackActivityName] = a.PostToSlackActivity

	return schema
}
