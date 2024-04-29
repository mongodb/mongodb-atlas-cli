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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasauth "go.mongodb.org/atlas/auth"
)

type RegisterOpts struct {
	LoginOpts
}

func (opts *RegisterOpts) RegisterRun(ctx context.Context) error {
	c, _, err := opts.RegistrationConfig(ctx)
	if err != nil {
		return err
	}
	if err = opts.registerFlow(ctx, c); err != nil {
		return err
	}

	// oauth config might have changed,
	// re-sync config profile with login opts
	if err = opts.SyncWithOAuthAccessProfile(opts.config)(); err != nil {
		return err
	}

	s, err := opts.config.AccessTokenSubject()
	if err != nil {
		return err
	}

	if err := opts.checkProfile(ctx); err != nil {
		return err
	}

	if err := opts.config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, "Successfully logged in as %s.\n", s)

	return opts.setUpProfile(ctx)
}

func (opts *LoginOpts) registerFlow(ctx context.Context, conf *atlasauth.RegistrationConfig) error {
	askedToOpenBrowser := false
	for {
		code, _, err := opts.RequestCode(ctx)
		code.VerificationURI = conf.RegistrationURL
		if err != nil {
			return err
		}

		opts.printAuthInstructions(code)
		if !askedToOpenBrowser {
			opts.handleBrowser(code.VerificationURI)
			askedToOpenBrowser = true
		}

		accessToken, _, err := opts.PollToken(ctx, code)
		if retry, errRetry := shouldRetryAuthenticate(err, newRegenerationPrompt()); errRetry != nil {
			return errRetry
		} else if retry {
			continue
		}
		if err != nil {
			return err
		}

		opts.AccessToken = accessToken.AccessToken
		opts.RefreshToken = accessToken.RefreshToken
		return nil
	}
}

func RegisterBuilder() *cobra.Command {
	opts := &RegisterOpts{}

	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register with MongoDB Atlas.",
		Example: `  # To start the interactive setup:
  atlas auth register
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			defaultProfile := config.Default()
			return prerun.ExecuteE(
				opts.SyncWithOAuthAccessProfile(defaultProfile),
				opts.InitFlow(defaultProfile),
				opts.LoginPreRun(cmd.Context()),
				validate.NoAPIKeys,
				validate.NoAccessToken)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, _ = fmt.Fprintf(opts.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")

			return opts.RegisterRun(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")

	return cmd
}
