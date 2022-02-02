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
	"fmt"
	"io"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/oauth"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type logoutOpts struct {
	*cli.DeleteOpts
	OutWriter io.Writer
	config    ConfigDeleter
	flow      Revoker
}

//go:generate mockgen -destination=../../mocks/mock_logout.go -package=mocks github.com/mongodb/mongocli/internal/cli/auth Revoker,ConfigDeleter

type ConfigDeleter interface {
	Delete() error
}

type Revoker interface {
	Revoke(context.Context, string, string) (*atlas.Response, error)
}

func (opts *logoutOpts) initFlow() error {
	var err error
	opts.flow, err = oauth.FlowWithConfig(config.Default())
	return err
}

func (opts *logoutOpts) Run(ctx context.Context) error {
	_, err := opts.flow.Revoke(ctx, config.RefreshToken(), "refresh_token")
	if err != nil {
		return err
	}

	if err := opts.Delete(opts.config.Delete); err != nil {
		return err
	}

	_, _ = fmt.Fprintln(opts.OutWriter, "Successfully logged out.")
	return nil
}

func LogoutBuilder() *cobra.Command {
	opts := &logoutOpts{
		DeleteOpts: cli.NewDeleteOpts("Successfully logged out\n", "Project not deleted"),
	}

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out the CLI.",
		Example: `  To log out from the CLI:
  $ mongocli auth logout
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			opts.config = config.Default()
			return opts.initFlow()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PromptWithMessage("Are you sure you want to log out"); err != nil {
				return err
			}
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	return cmd
}
