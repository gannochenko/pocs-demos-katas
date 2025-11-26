package domain

const (
	GenerateReportGithubWorkflowName = "GenerateReportGithubWorkflow"
)

type GenerateReportGithubWorkflowInput struct {
	Repository string `json:"repository"`
}

type GenerateReportGithubWorkflowOutput struct {
}