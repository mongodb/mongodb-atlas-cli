// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=token_mock_test.go -package=auth . TokenConfig

type TokenConfig interface {
	AccessToken() string
	Name() string
}

type tokenOpts struct {
	cli.OutputOpts
	config TokenConfig
}

func (opts *tokenOpts) Run() error {
	accessToken := opts.config.AccessToken()
	if accessToken == "" {
		return fmt.Errorf("no access token found for profile %s", opts.config.Name())
	}
	return opts.Print(accessToken)
}

func TokenBuilder() *cobra.Command {
	opts := &tokenOpts{}
	cmd := &cobra.Command{
		Use:    "token",
		Hidden: true,
		Short:  "Return the token for the current profile.",
		Example: `  # Return the token for the current profile:
  atlas auth token

  # Return the token for the current profile and save it to a file:
  atlas auth token > token.txt

  # Return the token for a specific profile:
  atlas auth token --profile <profile_name>
  `,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			// If the profile is set in the context, use it instead of the default profile
			profile, ok := config.ProfileFromContext(cmd.Context())
			if ok {
				opts.config = profile
			} else {
				opts.config = config.Default()
			}
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddOutputOptFlags(cmd)

	return cmd
}
