// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
