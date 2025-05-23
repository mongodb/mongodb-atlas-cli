// Copyright 2022 MongoDB Inc
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
	"context"
	"io"
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/oauth"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=logout_mock_test.go -package=auth . ConfigDeleter,Revoker

type ConfigDeleter interface {
	Delete() error
	SetAccessToken(string)
	SetRefreshToken(string)
	SetProjectID(string)
	SetOrgID(string)
	Save() error
}

type Revoker interface {
	RevokeToken(context.Context, string, string) (*atlas.Response, error)
}

type logoutOpts struct {
	*cli.DeleteOpts
	OutWriter  io.Writer
	config     ConfigDeleter
	flow       Revoker
	keepConfig bool
}

func (opts *logoutOpts) initFlow() error {
	var err error
	client := http.DefaultClient
	client.Transport = transport.Default()
	opts.flow, err = oauth.FlowWithConfig(config.Default(), client)
	return err
}

func (opts *logoutOpts) Run(ctx context.Context) error {
	// revoking a refresh token revokes the access token
	if _, err := opts.flow.RevokeToken(ctx, config.RefreshToken(), "refresh_token"); err != nil {
		return err
	}

	if !opts.keepConfig {
		return opts.Delete(opts.config.Delete)
	}
	opts.config.SetAccessToken("")
	opts.config.SetRefreshToken("")
	opts.config.SetProjectID("")
	opts.config.SetOrgID("")
	return opts.config.Save()
}

func LogoutBuilder() *cobra.Command {
	opts := &logoutOpts{
		DeleteOpts: cli.NewDeleteOpts("Successfully logged out of account %s\n", " "),
	}

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out of the CLI.",
		Example: `  # To log out from the CLI:
  atlas auth logout
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			opts.config = config.Default()
			return opts.initFlow()
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if config.RefreshToken() == "" {
				return ErrUnauthenticated
			}
			s, err := config.AccessTokenSubject()
			if err != nil {
				return err
			}
			opts.Entry = s
			if err := opts.PromptWithMessage("Are you sure you want to log out of account %s?"); err != nil || !opts.Confirm {
				return err
			}
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().BoolVar(&opts.keepConfig, "keep", false, usage.Keep)

	_ = cmd.Flags().MarkHidden("keep")
	return cmd
}
