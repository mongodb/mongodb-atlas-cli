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
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/auth"
)

//go:generate mockgen -destination=../../mocks/mock_login.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/cli/auth LoginConfig,LoginFlow

type LoginConfig interface {
	config.SetSaver
	AccessTokenSubject() (string, error)
}

var (
	ErrProjectIDNotFound = errors.New("you don't have access to this or it doesn't exist")
	ErrOrgIDNotFound     = errors.New("you don't have access to this organization ID or it doesn't exist")
)

type LoginOpts struct {
	cli.DefaultSetterOpts
	cli.RefresherOpts
	AccessToken    string
	RefreshToken   string
	IsGov          bool
	isCloudManager bool
	NoBrowser      bool
	SkipConfig     bool
	config         LoginConfig
}

type LoginFlow interface {
	LoginRun(ctx context.Context) error
	LoginPreRun() error
}

// SyncWithOAuthAccessProfile returns a function that is synchronizing the oauth settings
// from a login config profile with the provided command opts.
func (opts *LoginOpts) SyncWithOAuthAccessProfile(c LoginConfig) func() error {
	return func() error {
		opts.config = c

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
		if config.ClientID() != "" {
			opts.config.Set(config.ClientIDField, config.ClientID())
		}

		// sync OpsManagerURL from command opts (higher priority)
		// and OpsManagerURL from default profile
		if opts.OpsManagerURL != "" {
			opts.config.Set(config.OpsManagerURLField, opts.OpsManagerURL)
		}
		if config.OpsManagerURL() != "" {
			opts.OpsManagerURL = config.OpsManagerURL()
		}

		return nil
	}
}

func (opts *LoginOpts) LoginRun(ctx context.Context) error {
	if err := opts.oauthFlow(ctx); err != nil {
		return err
	}
	// oauth config might have changed,
	// re-sync config profile with login opts
	if err := opts.SyncWithOAuthAccessProfile(opts.config)(); err != nil {
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

	if opts.SkipConfig {
		return nil
	}

	if err := opts.setUpProfile(ctx); err != nil {
		return err
	}

	if config.Name() != config.DefaultProfile {
		_, _ = fmt.Fprintf(opts.OutWriter, "To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}

	return nil
}

func (opts *LoginOpts) checkProfile(ctx context.Context) error {
	if err := opts.InitStore(ctx); err != nil {
		return err
	}
	if config.OrgID() != "" && !opts.OrgExists(config.OrgID()) {
		opts.config.Set("org_id", "")
	}

	if config.ProjectID() != "" && !opts.ProjectExists(config.ProjectID()) {
		opts.config.Set("project_id", "")
	}
	return nil
}

func (opts *LoginOpts) setUpProfile(ctx context.Context) error {
	if err := opts.InitStore(ctx); err != nil {
		return err
	}
	// Initialize the text to be displayed if users are asked to select orgs or projects
	opts.OnMultipleOrgsOrProjects = func() {
		if !opts.AskedOrgsOrProjects {
			_, _ = fmt.Fprintln(opts.OutWriter, `Now set your default organization and project.

You have multiple organizations or projects, select one to proceed.`)
		}
	}

	if config.OrgID() == "" || !opts.OrgExists(config.OrgID()) {
		if err := opts.AskOrg(); err != nil {
			return err
		}
	}

	opts.SetUpOrg()

	if config.ProjectID() == "" || !opts.ProjectExists(config.ProjectID()) {
		if err := opts.AskProject(); err != nil {
			return err
		}
	}
	opts.SetUpProject()

	// Only make references to profile if user was asked about org or projects
	if opts.AskedOrgsOrProjects && opts.ProjectID != "" && opts.OrgID != "" {
		if !opts.ProjectExists(config.ProjectID()) {
			return ErrProjectIDNotFound
		}

		if !opts.OrgExists(config.OrgID()) {
			return ErrOrgIDNotFound
		}

		_, _ = fmt.Fprintf(opts.OutWriter, `
Your profile is now configured.
You can use [%s config set] to change these settings at a later time.
`, config.BinName())
	}

	return opts.config.Save()
}

func (opts *LoginOpts) printAuthInstructions(code *auth.DeviceCode) {
	codeDuration := time.Duration(code.ExpiresIn) * time.Second
	_, _ = fmt.Fprintf(opts.OutWriter, `
To verify your account, copy your one-time verification code:
`)

	userCode := fmt.Sprintf("%s-%s", code.UserCode[0:len(code.UserCode)/2], code.UserCode[len(code.UserCode)/2:])
	_, _ = fmt.Fprintln(opts.OutWriter, userCode)

	_, _ = fmt.Fprintf(opts.OutWriter, `
Paste the code in the browser when prompted to activate your Atlas CLI. Your code will expire after %.0f minutes.

To continue, go to `,
		codeDuration.Minutes(),
	)
	_, _ = fmt.Fprintln(opts.OutWriter, code.VerificationURI)
}

func (opts *LoginOpts) handleBrowser(uri string) {
	if opts.NoBrowser {
		return
	}

	if errBrowser := browser.OpenURL(uri); errBrowser != nil {
		_, _ = log.Warningln("There was an issue opening your browser")
	}
}

func (opts *LoginOpts) oauthFlow(ctx context.Context) error {
	askedToOpenBrowser := false
	for {
		code, _, err := opts.RequestCode(ctx)
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

func shouldRetryAuthenticate(err error, p survey.Prompt) (retry bool, errSurvey error) {
	if err == nil || !auth.IsTimeoutErr(err) {
		return false, nil
	}
	err = telemetry.TrackAskOne(p, &retry)
	return retry, err
}

func newRegenerationPrompt() survey.Prompt {
	return &survey.Confirm{
		Message: "Your one-time verification code is expired. Would you like to generate a new one?",
		Default: true,
	}
}

func (opts *LoginOpts) LoginPreRun(ctx context.Context) func() error {
	return func() error {
		// ignore expired tokens since logging in
		if err := opts.RefreshAccessToken(ctx); err != nil {
			// clean up any expired or invalid tokens
			opts.config.Set(config.AccessTokenField, "")

			if !errors.Is(err, cli.ErrInvalidRefreshToken) {
				return err
			}
		}

		return nil
	}
}

func tool() string {
	if config.ToolName == config.MongoCLI {
		return "Atlas or Cloud Manager"
	}
	return "Atlas"
}

func LoginBuilder() *cobra.Command {
	opts := &LoginOpts{}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with MongoDB Atlas.",
		Example: fmt.Sprintf(`  # To start the interactive login for your MongoDB %s account:
  %s auth login
`, tool(), config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			defaultProfile := config.Default()
			return prerun.ExecuteE(
				opts.SyncWithOAuthAccessProfile(defaultProfile),
				opts.InitFlow(defaultProfile),
				opts.LoginPreRun(cmd.Context()),
				validate.NoAPIKeys,
				validate.NoAccessToken,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.LoginRun(cmd.Context())
		},
		Args: require.NoArgs,
	}

	if config.ToolName == config.MongoCLI {
		cmd.Flags().BoolVar(&opts.isCloudManager, "cm", false, "Log in to Cloud Manager.")
	}

	cmd.Flags().BoolVar(&opts.IsGov, "gov", false, "Log in to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.SkipConfig, "skipConfig", false, "Skip profile configuration.")
	_ = cmd.Flags().MarkDeprecated("skipConfig", "if profile is configured, the login flow will skip the config step by default.")
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
