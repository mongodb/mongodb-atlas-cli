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
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/oauth"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/mock_login.go -package=mocks github.com/mongodb/mongocli/internal/cli/auth Authenticator,LoginConfig

type Authenticator interface {
	RequestCode(context.Context) (*auth.DeviceCode, *atlas.Response, error)
	PollToken(context.Context, *auth.DeviceCode) (*auth.Token, *atlas.Response, error)
}

type LoginConfig interface {
	config.SetSaver
	AccessTokenSubject() (string, error)
}

const (
	AlreadyAuthenticatedMsg      = "you are already authenticated with an API key (Public key: %s)"
	AlreadyAuthenticatedEmailMsg = "you are already authenticated with an email (%s)"
	LoginMsg                     = `run "atlas auth login" to refresh your session and continue`
	LoginWithProfileMsg          = `run "atlas auth login --profile <profile_name>"  to authenticate using your Atlas username and password on a new profile`
	LogoutToLoginAccountMsg      = `run "atlas auth logout" first if you want to login with another Atlas account on the same Atlas CLI profile`
)

var errTimedOut = errors.New("authentication timed out")

type LoginOpts struct {
	cli.DefaultSetterOpts
	AccessToken    string
	RefreshToken   string
	IsGov          bool
	isCloudManager bool
	NoBrowser      bool
	SkipConfig     bool
	config         LoginConfig
	flow           Authenticator
}

type LoginFlow interface {
	Run(ctx context.Context) error
	PreRun() error
}

func (opts *LoginOpts) initFlow() error {
	var err error
	opts.flow, err = oauth.FlowWithConfig(config.Default())
	return err
}

func (opts *LoginOpts) SetOAuthUpAccess() {
	switch {
	case opts.IsGov:
		opts.Service = config.CloudGovService
	case opts.isCloudManager:
		opts.Service = config.CloudManagerService
	default:
		opts.Service = config.CloudService
	}
	opts.config.Set("service", opts.Service)

	if opts.AccessToken != "" {
		opts.config.Set(config.AccessTokenField, opts.AccessToken)
	}
	if opts.RefreshToken != "" {
		opts.config.Set(config.RefreshTokenField, opts.RefreshToken)
	}
	if opts.OpsManagerURL != "" {
		opts.config.Set(config.OpsManagerURLField, opts.OpsManagerURL)
	}
	if config.ClientID() != "" {
		opts.config.Set(config.ClientIDField, config.ClientID())
	}
}

func (opts *LoginOpts) Run(ctx context.Context) error {
	if err := opts.oauthFlow(ctx); err != nil {
		return err
	}
	opts.SetOAuthUpAccess()
	s, err := opts.config.AccessTokenSubject()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, "Successfully logged in as %s.\n", s)
	if opts.SkipConfig {
		return opts.config.Save()
	}
	if err := opts.InitStore(ctx); err != nil {
		return err
	}
	_, _ = fmt.Fprint(opts.OutWriter, "Press Enter to continue your profile configuration")
	_, _ = fmt.Scanln()
	if err := opts.AskOrg(); err != nil {
		return err
	}
	opts.SetUpOrg()
	if err := opts.AskProject(); err != nil {
		return err
	}
	opts.SetUpProject()

	if err := survey.Ask(opts.DefaultQuestions(), opts); err != nil {
		return err
	}
	opts.SetUpOutput()
	opts.SetUpMongoSHPath()
	opts.SetUpTelemetryEnabled()
	if err := opts.config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprint(opts.OutWriter, "\nYour profile is now configured.\n")
	if config.Name() != config.DefaultProfile {
		_, _ = fmt.Fprintf(opts.OutWriter, "To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}

	_, _ = fmt.Fprintf(opts.OutWriter, "You can use [%s config set] to change these settings at a later time.\n", config.BinName())

	return nil
}

func (opts *LoginOpts) printAuthInstructions(code *auth.DeviceCode) {
	codeDuration := time.Duration(code.ExpiresIn) * time.Second
	_, _ = fmt.Fprintf(opts.OutWriter, `
First, copy your one-time code: %s-%s

Next, sign in with your browser and enter the code.

Or go to %s

Your code will expire after %.0f minutes.
`,
		code.UserCode[0:len(code.UserCode)/2],
		code.UserCode[len(code.UserCode)/2:],
		code.VerificationURI,
		codeDuration.Minutes(),
	)
}

func (opts *LoginOpts) oauthFlow(ctx context.Context) error {
	code, _, err := opts.flow.RequestCode(ctx)
	if err != nil {
		return err
	}

	opts.printAuthInstructions(code)

	if !opts.NoBrowser {
		if errBrowser := browser.OpenURL(code.VerificationURI); errBrowser != nil {
			_, _ = fmt.Fprintf(os.Stderr, "There was an issue opening your browser\n")
		}
	}

	var accessToken *auth.Token
	if accessToken, _, err = opts.flow.PollToken(ctx, code); err != nil {
		if auth.IsTimeoutErr(err) {
			return errTimedOut
		}
		return err
	}

	opts.AccessToken = accessToken.AccessToken
	opts.RefreshToken = accessToken.RefreshToken
	return nil
}

func hasUserProgrammaticKeys() bool {
	return config.PublicAPIKey() != "" && config.PrivateAPIKey() != ""
}

func (opts *LoginOpts) loginPreRun(ctx context.Context) error {
	if hasUserProgrammaticKeys() {
		msg := fmt.Sprintf(AlreadyAuthenticatedMsg, config.PublicAPIKey())
		return fmt.Errorf(`%s

%s`, msg, LoginWithProfileMsg)
	}

	if account, err := AccountWithAccessToken(); err == nil {
		if err := cli.RefreshToken(ctx); err == nil && validate.Token() == nil {
			msg := fmt.Sprintf(AlreadyAuthenticatedEmailMsg, account)
			return fmt.Errorf(`%s

%s`, msg, LogoutToLoginAccountMsg)
		}
	}
	return nil
}

func (opts *LoginOpts) PreRun() error {
	opts.config = config.Default()
	if config.OpsManagerURL() != "" {
		opts.OpsManagerURL = config.OpsManagerURL()
	}
	return opts.initFlow()
}

func Tool() string {
	if config.ToolName == config.MongoCLI {
		return "Cloud Manager"
	}
	return "Atlas"
}

func LoginBuilder() *cobra.Command {
	opts := &LoginOpts{}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with MongoDB Atlas.",
		Example: fmt.Sprintf(`  To start the interactive login for your MongoDB %s account:
  $ %s auth login
`, Tool(), config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			if err := opts.loginPreRun(cmd.Context()); err != nil {
				return err
			}
			return opts.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	if config.ToolName == config.MongoCLI {
		cmd.Flags().BoolVar(&opts.isCloudManager, "cm", false, "Log in to Cloud Manager.")
	}

	cmd.Flags().BoolVar(&opts.IsGov, "gov", false, "Log in to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.SkipConfig, "skipConfig", false, "Skip profile configuration.")
	return cmd
}

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage the CLI's authentication state.",
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		LoginBuilder(),
		WhoAmIBuilder(),
		LogoutBuilder(),
	)

	if config.ToolName == config.AtlasCLI {
		cmd.AddCommand(RegisterBuilder())
	}

	return cmd
}
