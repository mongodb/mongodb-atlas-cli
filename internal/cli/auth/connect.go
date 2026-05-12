// Copyright 2026 MongoDB Inc
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
	"net/http"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/auth"
)

// ConnectConfig defines the profile operations needed by the connect command.
type ConnectConfig interface {
	SetSaver
	SetTokenExpiry(string) // TODO: remove when Token() reads expiry directly instead of via tokenClaims
	AuthServerMetadata() map[string]any
	SetAuthServerMetadata(map[string]any)
	Service() string
	ClientID() string
	AuthServerURL() string
}

type ConnectOpts struct {
	config         ConnectConfig
	OutWriter      io.Writer
	NoBrowser      bool
	Discover bool
}

// discoverOrLoadMetadata returns cached AS metadata if still valid and from
// the expected issuer, otherwise fetches fresh metadata via RFC 8414 discovery.
func (opts *ConnectOpts) discoverOrLoadMetadata(ctx context.Context, authCfg *auth.Config) (map[string]any, error) {
	if opts.Discover {
		opts.config.SetAuthServerMetadata(nil)
		if err := opts.config.Save(); err != nil {
			return nil, fmt.Errorf("failed to clear metadata cache: %w", err)
		}
	}

	if cached := opts.config.AuthServerMetadata(); cached != nil {
		if !metadataExpired(cached) {
			if metadata, ok := cached["metadata"].(map[string]any); ok {
				// Verify the cached metadata came from the AS we're configured to use
				if issuerStr, ok := metadata["issuer"].(string); ok && issuerStr == authCfg.AuthServerURL.String() {
					return metadata, nil
				}
			}
		}
	}

	cached, err := authCfg.DiscoverAuthServer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover authorization server: %w", err)
	}

	opts.config.SetAuthServerMetadata(cached)
	if err := opts.config.Save(); err != nil {
		return nil, fmt.Errorf("failed to save discovery metadata: %w", err)
	}

	if metadata, ok := cached["metadata"].(map[string]any); ok {
		return metadata, nil
	}

	return nil, errors.New("discovery response missing metadata")
}

func (opts *ConnectOpts) Run(ctx context.Context) error {
	client := &http.Client{Transport: transport.Default()}
	authCfg, err := transport.FlowForAuthIssuer(opts.config, client, version.Version)
	if err != nil {
		return fmt.Errorf("failed to configure auth issuer: %w", err)
	}

	metadata, err := opts.discoverOrLoadMetadata(ctx, authCfg)
	if err != nil {
		return err
	}

	authorizationEndpoint, ok := metadata["authorization_endpoint"].(string)
	if !ok || authorizationEndpoint == "" {
		return errors.New("authorization_endpoint not found in server metadata")
	}
	tokenEndpoint, ok := metadata["token_endpoint"].(string)
	if !ok || tokenEndpoint == "" {
		return errors.New("token_endpoint not found in server metadata")
	}

	// Generate PKCE and state
	pkce, err := auth.GeneratePKCE()
	if err != nil {
		return err
	}
	state, err := auth.GenerateState()
	if err != nil {
		return err
	}

	var code string
	var redirectURI string

	if opts.NoBrowser {
		// Manual flow: user authorizes in their own browser and pastes the redirect URL back
		redirectURI = auth.NoBrowserRedirectURI()
		authURL, err := authCfg.AuthorizationURL(authorizationEndpoint, redirectURI, state, pkce)
		if err != nil {
			return err
		}

		_, _ = fmt.Fprintf(opts.OutWriter, "\nTo authenticate, visit the following URL in your browser:\n\n%s\n", authURL)
		_, _ = fmt.Fprintln(opts.OutWriter, "\nAfter you approve, your browser will redirect to a page that fails to load.")
		_, _ = fmt.Fprintln(opts.OutWriter, "This is expected. Copy the full URL from your browser's address bar and paste it here.")
		_, _ = fmt.Fprint(opts.OutWriter, "\nPaste the URL: ")

		code, err = auth.ParseCodeFromRedirectURL(state)
		if err != nil {
			return fmt.Errorf("authorization failed: %w", err)
		}
	} else {
		// Browser flow: start a local callback server and open the browser
		callbackServer, err := auth.StartCallbackServer(state)
		if err != nil {
			return fmt.Errorf("failed to start callback server: %w", err)
		}
		defer callbackServer.Close()

		redirectURI = callbackServer.RedirectURI()
		authURL, err := authCfg.AuthorizationURL(authorizationEndpoint, redirectURI, state, pkce)
		if err != nil {
			return err
		}

		if errBrowser := browser.OpenURL(authURL); errBrowser != nil {
			_, _ = fmt.Fprintf(opts.OutWriter, "\nThere was an issue opening your browser. To authenticate, visit:\n%s\n", authURL)
		} else if log.IsDebugLevel() {
			_, _ = fmt.Fprintf(opts.OutWriter, "\nAuthorization URL: %s\n", authURL)
		}

		_, _ = fmt.Fprintln(opts.OutWriter, "\nWaiting for authorization...")
		code, err = callbackServer.WaitForCallback(ctx)
		if err != nil {
			return fmt.Errorf("authorization failed: %w", err)
		}
	}

	// Exchange the authorization code for tokens
	token, err := authCfg.ExchangeCode(ctx, tokenEndpoint, code, redirectURI, pkce.CodeVerifier)
	if err != nil {
		return err
	}

	// Store the tokens, expiry, and auth type
	opts.config.SetAccessToken(token.AccessToken)
	opts.config.SetRefreshToken(token.RefreshToken)
	if !token.Expiry.IsZero() {
		opts.config.SetTokenExpiry(token.Expiry.Format(time.RFC3339))
	}
	opts.config.SetAuthType(config.UserDelegation)
	if err := opts.config.Save(); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	_, _ = fmt.Fprintln(opts.OutWriter, "Successfully connected to MongoDB Atlas.")
	return nil
}

func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}

	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect to MongoDB Atlas using the dedicated authorization server.",
		Long:  `This command authenticates with MongoDB Atlas via the dedicated OAuth Authorization Server at authorize.mongodb.com.`,
		Example: `  # Connect to your MongoDB Atlas account:
  atlas auth connect
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.config = config.Default()
			opts.OutWriter = cmd.OutOrStdout()
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't automatically open a browser session.")
	cmd.Flags().BoolVar(&opts.Discover, "discover", false, "Force re-discovery of authorization server metadata.")

	return cmd
}

// metadataExpired checks whether the cached metadata has passed its expiry.
func metadataExpired(cached map[string]any) bool {
	expiry, ok := cached["expiry"].(string)
	if !ok {
		return true
	}
	t, err := time.Parse(time.RFC3339, expiry)
	if err != nil {
		return true
	}
	return time.Now().After(t)
}
