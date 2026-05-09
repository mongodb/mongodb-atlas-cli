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
	"net/http"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/cobra"
)

// ConnectConfig defines the profile operations needed by the connect command.
type ConnectConfig interface {
	SetSaver
	AuthServerMetadata() map[string]any
	SetAuthServerMetadata(map[string]any)
	Service() string
}

type ConnectOpts struct {
	config ConnectConfig
}

// discoverOrLoadMetadata returns cached AS metadata if still valid,
// otherwise fetches fresh metadata via RFC 8414 discovery and caches it.
func (opts *ConnectOpts) discoverOrLoadMetadata(ctx context.Context) (map[string]any, error) {
	if cached := opts.config.AuthServerMetadata(); cached != nil {
		if !metadataExpired(cached) {
			if metadata, ok := cached["metadata"].(map[string]any); ok {
				return metadata, nil
			}
		}
	}

	client := &http.Client{Transport: transport.Default()}
	authCfg, err := transport.FlowForAuthIssuer(opts.config, client, version.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to configure auth issuer: %w", err)
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
	metadata, err := opts.discoverOrLoadMetadata(ctx)
	if err != nil {
		return err
	}

	_ = metadata // will be used by the authorization code flow

	opts.config.SetAuthType(config.UserDelegation)
	if err := opts.config.Save(); err != nil {
		return err
	}

	return errors.New("connect is not yet implemented")
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
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

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
