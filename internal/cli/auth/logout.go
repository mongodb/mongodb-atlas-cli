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
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20250312008/auth/clientcredentials"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=logout_mock_test.go -package=auth . ConfigDeleter,Revoker

type ConfigDeleter interface {
	Delete() error
	Name() string
	SetAccessToken(string)
	SetRefreshToken(string)
	SetProjectID(string)
	SetOrgID(string)
	SetPublicAPIKey(string)
	SetPrivateAPIKey(string)
	SetClientID(string)
	SetClientSecret(string)
	AuthType() config.AuthMechanism
	PublicAPIKey() string
	ClientID() string
	ClientSecret() string
	AccessTokenSubject() (string, error)
	RefreshToken() string
	Save() error
}

type Revoker interface {
	RevokeToken(context.Context, string, string) (*atlas.Response, error)
}

type logoutOpts struct {
	*cli.DeleteOpts
	cli.DefaultSetterOpts
	OutWriter                 io.Writer
	config                    ConfigDeleter
	flow                      Revoker
	keepConfig                bool
	revokeServiceAccountToken func() error
}

func (opts *logoutOpts) initFlow(ctx context.Context) error {
	var err error
	client := http.DefaultClient
	client.Transport = transport.Default()
	opts.flow, err = transport.FlowWithConfig(config.Default(), client, version.Version)
	opts.revokeServiceAccountToken = func() error {
		return revokeServiceAccountToken(ctx, opts.config.ClientID(), opts.config.ClientSecret())
	}
	return err
}

func revokeServiceAccountToken(ctx context.Context, clientID, clientSecret string) error {
	cfg := clientcredentials.NewConfig(clientID, clientSecret)
	if config.OpsManagerURL() != "" {
		// TokenURL and RevokeURL points to "https://cloud.mongodb.com/api/oauth/<token/revoke>". Modify TokenURL and RevokeURL if OpsManagerURL does not point to cloud.mongodb.com
		baseURL := strings.TrimSuffix(config.OpsManagerURL(), "/")
		cfg.TokenURL = baseURL + clientcredentials.TokenAPIPath
		cfg.RevokeURL = baseURL + clientcredentials.RevokeAPIPath
	}
	token, err := cfg.Token(ctx)
	if err != nil {
		return err
	}
	return cfg.RevokeToken(ctx, token)
}

func (opts *logoutOpts) Run(ctx context.Context) error {
	if !opts.Confirm {
		return nil
	}

	switch opts.config.AuthType() {
	case config.UserAccount:
		_, err := opts.flow.RevokeToken(ctx, config.RefreshToken(), "refresh_token")
		if err != nil {
			_, _ = log.Warningf("Warning: unable to revoke user account token: %v, proceeding with logout\n", err)
		}
	case config.ServiceAccount:
		if err := opts.revokeServiceAccountToken(); err != nil {
			_, _ = log.Warningf("Warning: unable to revoke service account token: %v, proceeding with logout\n", err)
		}
	case config.APIKeys, config.NoAuth, "":
	}

	// Clean up all the config
	opts.config.SetPublicAPIKey("")
	opts.config.SetPrivateAPIKey("")
	opts.config.SetAccessToken("")
	opts.config.SetRefreshToken("")
	opts.config.SetClientID("")
	opts.config.SetClientSecret("")
	opts.config.SetProjectID("")
	opts.config.SetOrgID("")

	if !opts.keepConfig {
		return opts.config.Delete()
	}

	return opts.config.Save()
}

func LogoutBuilder() *cobra.Command {
	opts := &logoutOpts{
		DeleteOpts: cli.NewDeleteOpts("Successfully logged out of '%s'\n", " "),
	}

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out of the CLI.",
		Example: `  # To log out from the CLI:
  atlas auth logout
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			// If the profile is set in the context, use it instead of the default profile
			profile, ok := config.ProfileFromContext(cmd.Context())
			if ok {
				opts.config = profile
			} else {
				opts.config = config.Default()
			}

			// Only initialize OAuth flow if we have OAuth-based auth
			if opts.config.AuthType() == config.UserAccount || opts.config.AuthType() == config.ServiceAccount {
				return opts.initFlow(cmd.Context())
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			var message string

			entry := opts.config.Name()
			logoutMessage := "Are you sure you want to log out of profile %s"

			switch opts.config.AuthType() {
			case config.APIKeys:
				message = logoutMessage + " with public API key " + opts.config.PublicAPIKey() + "?"
			case config.ServiceAccount:
				message = logoutMessage + " with service account " + opts.config.ClientID() + "?"
			case config.UserAccount:
				subject, err := opts.config.AccessTokenSubject()
				if err != nil {
					return err
				}

				if opts.config.RefreshToken() == "" {
					return ErrUnauthenticated
				}

				message = logoutMessage + " with user account " + subject + "?"
			case config.NoAuth, "":
				message = logoutMessage + "?"
			}

			opts.Entry = entry
			if err := opts.PromptWithMessage(message); err != nil {
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
