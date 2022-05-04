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
	"github.com/mongodb/mongocli/internal/cli/separate_flow"
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
	Login LoginOpts
	flow  separate_flow.Flow
}

func (opts *RegisterOpts) RegisterAndAuthenticate(ctx context.Context) error {
	// TODO:CLOUDP-121210 - Replace with new request and remove URI override.
	code, _, err := opts.Login.flow.RequestCode(ctx)
	if err != nil {
		return err
	}

	if opts.Login.IsGov {
		code.VerificationURI = govAccountURI
	} else {
		code.VerificationURI = accountURI
	}

	opts.Login.printAuthInstructions(code)

	if !opts.Login.NoBrowser {
		if errBrowser := browser.OpenURL(code.VerificationURI); errBrowser != nil {
			_, _ = fmt.Fprintf(os.Stderr, "There was an issue opening your browser\n")
		}
	}

	accessToken, _, err := opts.Login.flow.PollToken(ctx, code)
	var target *atlas.ErrorResponse
	if errors.As(err, &target) && target.ErrorCode == authExpiredError {
		return errTimedOut
	}
	if err != nil {
		return err
	}
	opts.Login.AccessToken = accessToken.AccessToken
	opts.Login.RefreshToken = accessToken.RefreshToken
	return nil
}

func (opts *RegisterOpts) Run() error {
	return opts.flow.Flow(opts)
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
			opts.flow = separate_flow.Flow{cmd.Context()}
			if hasUserProgrammaticKeys() {
				return fmt.Errorf(`you have already set the programmatic keys for this profile. 

Run '%s auth register --profile <profileName>' to use your username and password with a new profile`, config.BinName())
			}

			opts.Login.OutWriter = cmd.OutOrStdout()
			opts.Login.Config = config.Default()
			if config.OpsManagerURL() != "" {
				opts.Login.OpsManagerURL = config.OpsManagerURL()
			}
			return opts.Login.initFlow()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.Login.IsGov, "gov", false, "Register to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.Login.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.Login.SkipConfig, "skipConfig", false, "Skip profile configuration.")

	return cmd
}
