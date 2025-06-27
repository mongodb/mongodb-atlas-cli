package plugin

import (
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/google/go-github/v61/github"
)

func NewAuthenticatedGithubClient() *github.Client {
	// create a new github client
	ghClient := github.NewClient(nil)

	// try to get the gh token from either the gh cli config or the environment variable (GH_TOKEN/GITHUB_TOKEN)
	if token, _ := auth.TokenForHost("github.com"); token != "" {
		ghClient = ghClient.WithAuthToken(token)
	}

	return ghClient
}
