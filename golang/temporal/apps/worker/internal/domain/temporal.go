package domain

import "fmt"

const (
	GenerateReportGithubWorkflowName = "GenerateReportGithubWorkflow"

	GetRepositoryStatsActivityName = "GetRepositoryStatsActivity"
	MakeHumanReadableSummaryActivityName = "MakeHumanReadableSummaryActivity"
	PostToSlackActivityName = "PostToSlackActivity"
)

type GenerateReportGithubWorkflowInput struct {
	Repository string `json:"repository"`
}

type GenerateReportGithubWorkflowOutput struct {
}

type GetRepositoryStatsActivityInput struct {
	Repository string `json:"repository"`
}

type Commit struct {
	SHA string `json:"sha"`
	Message string `json:"message"`
	AuthorName string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Date string `json:"date"`
}

type GetRepositoryStatsActivityOutput struct {
	Commits []Commit `json:"commits"`
}

func (o *GetRepositoryStatsActivityOutput) ToText() string {
	result := ""
	for _, commit := range o.Commits {
	result += fmt.Sprintf("Commit SHA: %s\n", commit.SHA)
		result += fmt.Sprintf("Message: %s\n", commit.Message)
		result += fmt.Sprintf("Author Name: %s\n", commit.AuthorName)
		result += fmt.Sprintf("Author Email: %s\n", commit.AuthorEmail)
		result += fmt.Sprintf("Date: %s\n", commit.Date)
		result += "=========\n"
	}
	return result
}

type MakeHumanReadableSummaryActivityInput struct {
	Text string `json:"text"`
}

type MakeHumanReadableSummaryActivityOutput struct {
	Summary string `json:"summary"`
}

type PostToSlackActivityInput struct {
	Repository string `json:"repository"`
	Summary string `json:"summary"`
}

type PostToSlackActivityOutput struct {
}
