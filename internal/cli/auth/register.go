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
	"net/url"
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

var errUnexpectedRegistrationURL = errors.New("unexpected registration URL")

type RegisterOpts struct {
	LoginOpts
}

func (opts *RegisterOpts) RegisterRun(ctx context.Context) error {
	if err := opts.checkNoOtherCredentials(config.UserAccount); err != nil {
		return err
	}
	opts.config.SetAuthType(config.UserAccount)

	c, _, err := opts.RegistrationConfig(ctx)
	if err != nil {
		return err
	}
	if c == nil {
		return errUnexpectedRegistrationURL
	}
	registerURL, err := accountRegisterURL(c.RegistrationURL)
	if err != nil {
		return err
	}
	if err = opts.authCodeFlow(ctx, func(authURL string) string {
		u := *registerURL
		u.RawQuery = url.Values{"n": {authURL}}.Encode()
		return u.String()
	}); err != nil {
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

// accountRegisterURL keeps only the account host, since /account/register/cli is the device-code sign-up and ignores n.
func accountRegisterURL(registrationURL string) (*url.URL, error) {
	u, err := url.Parse(registrationURL)
	if err != nil || u.Scheme != "https" ||
		(!strings.HasSuffix(u.Hostname(), ".mongodb.com") && !strings.HasSuffix(u.Hostname(), ".mongodbgov.com")) {
		return nil, fmt.Errorf("%w: %q", errUnexpectedRegistrationURL, registrationURL)
	}
	return &url.URL{Scheme: u.Scheme, Host: u.Host, Path: "/account/register"}, nil
}

// Asker must be set, promptOutput calls it and a nil TrackAsker panics.
func newRegisterOpts() *RegisterOpts {
	return &RegisterOpts{
		LoginOpts: LoginOpts{
			Asker: &telemetry.Ask{},
		},
	}
}

func RegisterBuilder() *cobra.Command {
	opts := newRegisterOpts()

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
				opts.InitLegacyDeviceCodeFlow(defaultProfile),
				opts.LoginPreRun(cmd.Context()),
				validate.NoAPIKeys,
				validate.NoAccessToken)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, _ = fmt.Fprintf(opts.OutWriter, "Create and verify your MongoDB Atlas account in your browser. Atlas CLI continues automatically after you sign up.\n")

			return opts.RegisterRun(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")

	return cmd
}
