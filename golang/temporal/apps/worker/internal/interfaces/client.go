package interfaces

import (
	"context"

	"github.com/google/go-github/v62/github"
)

type GitHubClient interface {
	Connect(ctx context.Context)
	FetchCommits(ctx context.Context, repository string) ([]*github.RepositoryCommit, error)
}