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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/cobra"
)

type RegisterOpts struct {
	LoginOpts
}

func (opts *RegisterOpts) RegisterRun(ctx context.Context) error {
	if err := opts.oauthFlow(ctx); err != nil {
		return err
	}

	opts.SetUpOAuthAccess()
	s, err := opts.config.AccessTokenSubject()
	if err != nil {
		return err
	}

	if err := opts.CheckProfile(ctx); err != nil {
		return err
	}

	if err := opts.config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, "Successfully logged in as %s.\n", s)

	return opts.setUpProfile(ctx)
}

func RegisterBuilder() *cobra.Command {
	opts := &RegisterOpts{}

	cmd := &cobra.Command{
		Use:    "register",
		Short:  "Register with MongoDB Atlas.",
		Hidden: false,
		Example: fmt.Sprintf(`  # To start the interactive setup:
  %s auth register
`, config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return prerun.ExecuteE(
				opts.LoginPreRun,
				opts.InitFlow(config.Default()),
				validate.NoAPIKeys,
				validate.NoAccessToken)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _ = fmt.Fprintf(opts.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")

			return opts.RegisterRun(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.IsGov, "gov", false, "Register with Atlas for Government.")
	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")

	return cmd
}
