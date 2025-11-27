package interfaces

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type GitHubClient interface {
	Connect(ctx context.Context)
	FetchCommits(ctx context.Context, repository string) ([]*github.RepositoryCommit, error)
}

type OpenAIClient interface {
	Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error)
}

type SlackClient interface {
	SendMessage(ctx context.Context, channelID, message string) error
}