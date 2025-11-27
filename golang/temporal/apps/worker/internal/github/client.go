package github

import (
	"context"
	"errors"
	"strings"
	"time"
	"worker/internal/domain"
	"worker/internal/interfaces"

	libCtx "lib/ctx"

	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	delegateClient *github.Client
	config *domain.Config
}

func NewClient(config *domain.Config) interfaces.GitHubClient {
	return &GitHubClient{
		delegateClient: nil,
		config: config,
	}
}

func (c *GitHubClient) Connect(ctx context.Context) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.config.GitHub.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	c.delegateClient = github.NewClient(tc)
}

func (c *GitHubClient) FetchCommits(ctx context.Context, repository string) ([]*github.RepositoryCommit, error) {
	if c.delegateClient == nil {
		return nil, errors.New("client not connected")
	}

	since := libCtx.GetTime(ctx).Add(-24 * time.Hour)
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    1,
		},
		Since: since,
	}

	parts := strings.Split(repository, "/")
	if len(parts) != 2 {
		return nil, errors.New("invalid repository format, should be <owner>/<repo>")
	}

	var allCommits []*github.RepositoryCommit

	for {
		commits, resp, err := c.delegateClient.Repositories.ListCommits(ctx, parts[0], parts[1], opts)
		if err != nil {
			return nil, err
		}

		allCommits = append(allCommits, commits...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage

		// Small delay to avoid hitting rate limits
		time.Sleep(200 * time.Millisecond)
	}

	return allCommits, nil
}
