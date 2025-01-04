package client

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

const (
	GitHubTokenEnv = "GIT_TOKEN"
)

type GitHubClient struct {
	Client *github.Client
	C      context.Context
}

// Initializes a new GitHub client using an OAuth Token
func InitClient() (*GitHubClient, error) {
	token := os.Getenv(GitHubTokenEnv)
	if token == "" {
		return nil, fmt.Errorf("environment variable not set: %v", GitHubTokenEnv)
	}

	c := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(c, ts)

	return &GitHubClient{
		Client: github.NewClient(tc),
		C:      c,
	}, nil
}
