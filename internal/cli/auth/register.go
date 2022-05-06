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
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/mock_register.go -package=mocks github.com/mongodb/mongocli/internal/cli/auth RegisterFlow

const (
	WithProfileMsg                      = `Run "atlas auth register --profile <profile_name>" to create a new Atlas account on a new Atlas CLI profile.`
	accountURI                          = "https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect"
	govAccountURI                       = "https://account.mongodbgov.com/account/register?fromURI=https://account.mongodbgov.com/account/connect"
	RegisterWithProfileForNewAccountMsg = `Run "atlas auth register --profile <profile_name>" to create a new account on a new profile.`
)

type registerSurvey struct {
	confirm func(message string, defaultResponse bool) (response bool, err error)
}

type registerOpts struct {
	cli.DefaultSetterOpts
	login          *LoginOpts
	registerSurvey *registerSurvey
}

var defaultRegisterSurvey = registerSurvey{
	confirm: func(message string, defaultResponse bool) (response bool, err error) {
		p := &survey.Confirm{
			Message: message,
			Default: defaultResponse,
		}
		err = survey.AskOne(p, &response)
		return response, err
	},
}

func NewRegisterFlow(l *LoginOpts) RegisterFlow {
	return &registerOpts{
		registerSurvey: &defaultRegisterSurvey,
		login:          l,
	}
}

type RegisterFlow interface {
	Run(ctx context.Context) error
	PreRun(outWriter io.Writer) error
}

func (opts *registerOpts) registerAndAuthenticate(ctx context.Context) error {
	// TODO:CLOUDP-121210 - Replace with new request and remove URI override.
	for {
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
		if retry, errRetry := opts.shouldRetryRegister(err); errRetry != nil {
			return errRetry
		} else if retry {
			continue
		}

		if err != nil {
			return err
		}

		opts.login.AccessToken = accessToken.AccessToken
		opts.login.RefreshToken = accessToken.RefreshToken
		return nil
	}
}

func (opts *registerOpts) shouldRetryRegister(err error) (retry bool, errSurvey error) {
	var target *atlas.ErrorResponse
	tokenExpired := err == auth.ErrTimeout || (errors.As(err, &target) && target.ErrorCode == authExpiredError)
	if !tokenExpired {
		return false, nil
	}

	return opts.registerSurvey.confirm("Your one-time verification code is expired. Would you like to generate a new one?", true)
}

func (opts *registerOpts) setUpProfile(ctx context.Context) error {
	opts.login.SetOAuthUpAccess()
	s, err := opts.login.config.AccessTokenSubject()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, "Successfully logged in as %s.\n", s)
	if opts.login.SkipConfig {
		return opts.login.config.Save()
	}
	if err := opts.InitStore(ctx); err != nil {
		return err
	}

	if err := opts.AskOrg(); err != nil {
		return err
	}
	opts.SetUpOrg()
	if err := opts.AskProject(); err != nil {
		return err
	}
	opts.SetUpProject()

	opts.SetUpMongoSHPath()
	opts.SetUpTelemetryEnabled()
	if err := opts.login.config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprint(opts.OutWriter, "\nYour profile is now configured.\n")
	if config.Name() != config.DefaultProfile {
		_, _ = fmt.Fprintf(opts.OutWriter, "To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}

	_, _ = fmt.Fprintf(opts.OutWriter, "You can use [%s config set] to change these settings at a later time.\n", config.BinName())

	return nil
}

func (opts *registerOpts) Run(ctx context.Context) error {
	_, _ = fmt.Fprintf(opts.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")

	if err := opts.registerAndAuthenticate(ctx); err != nil {
		return err
	}

	return opts.setUpProfile(ctx)
}

func (opts *registerOpts) PreRun(outWriter io.Writer) error {
	opts.OutWriter = outWriter
	opts.login.OutWriter = outWriter
	opts.login.config = config.Default()
	if config.OpsManagerURL() != "" {
		opts.login.OpsManagerURL = config.OpsManagerURL()
	}
	return opts.login.initFlow()
}

func (opts *registerOpts) registerPreRun() error {
	if hasUserProgrammaticKeys() {
		msg := fmt.Sprintf(AlreadyAuthenticatedMsg, config.PublicAPIKey())
		return fmt.Errorf(`%s

%s`, msg, WithProfileMsg)
	}

	if account, err := AccountWithAccessToken(); err == nil {
		msg := fmt.Sprintf(AlreadyAuthenticatedEmailMsg, account)
		return fmt.Errorf(`%s

%s
`, msg, RegisterWithProfileForNewAccountMsg)
	}
	return nil
}

func RegisterBuilder() *cobra.Command {
	opts := &registerOpts{
		registerSurvey: &defaultRegisterSurvey,
		login:          &LoginOpts{},
	}
	cmd := &cobra.Command{
		Use:    "register",
		Short:  "Register with MongoDB Atlas.",
		Hidden: true,
		Example: fmt.Sprintf(`  To start the interactive setup:
  $ %s auth register
`, config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.registerPreRun(); err != nil {
				return err
			}
			return opts.PreRun(cmd.OutOrStdout())
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
