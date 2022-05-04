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
	"errors"
	"fmt"
	"os"

	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const accountURI = "https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect"
const govAccountURI = "https://account.mongodbgov.com/account/register?fromURI=https://account.mongodbgov.com/account/connect"

type RegisterOpts struct {
	login LoginOpts
}

func (opts *RegisterOpts) registerAndAuthenticate(ctx context.Context) error {
	// TODO:CLOUDP-121210 - Replace with new request and remove URI override.
	code, _, err := opts.login.flow.RequestCode(ctx)
	if err != nil {
		return err
	}

	if opts.login.IsGov {
		code.VerificationURI = govAccountURI
	} else {
		code.VerificationURI = accountURI
	}

	opts.login.printAuthInstructions(code)

	if !opts.login.NoBrowser {
		if errBrowser := browser.OpenURL(code.VerificationURI); errBrowser != nil {
			_, _ = fmt.Fprintf(os.Stderr, "There was an issue opening your browser\n")
		}
	}

	accessToken, _, err := opts.login.flow.PollToken(ctx, code)
	var target *atlas.ErrorResponse
	if errors.As(err, &target) && target.ErrorCode == authExpiredError {
		return errTimedOut
	}
	if err != nil {
		return err
	}
	opts.login.AccessToken = accessToken.AccessToken
	opts.login.RefreshToken = accessToken.RefreshToken
	return nil
}

func (opts *RegisterOpts) Run(ctx context.Context) error {
	_, _ = fmt.Fprintf(opts.login.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")

	if err := opts.registerAndAuthenticate(ctx); err != nil {
		return err
	}

	opts.login.SetOAuthUpAccess()
	s, err := opts.login.config.AccessTokenSubject()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.login.OutWriter, "Successfully logged in as %s.\n", s)
	if opts.login.SkipConfig {
		return opts.login.config.Save()
	}

	return nil
}

func RegisterBuilder() *cobra.Command {
	opts := &RegisterOpts{}
	cmd := &cobra.Command{
		Use:    "register",
		Short:  "Register with MongoDB Atlas.",
		Hidden: true,
		Example: fmt.Sprintf(`  To start the interactive setup:
  $ %s auth register
`, config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if hasUserProgrammaticKeys() {
				return fmt.Errorf(`you have already set the programmatic keys for this profile. 

Run '%s auth register --profile <profileName>' to use your username and password with a new profile`, config.BinName())
			}

			opts.login.OutWriter = cmd.OutOrStdout()
			opts.login.config = config.Default()
			if config.OpsManagerURL() != "" {
				opts.login.OpsManagerURL = config.OpsManagerURL()
			}
			return opts.login.initFlow()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.login.IsGov, "gov", false, "Register to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.login.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.login.SkipConfig, "skipConfig", false, "Skip profile configuration.")

	return cmd
}
