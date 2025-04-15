// Copyright 2020 MongoDB Inc
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

package clusters

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

const idle = "IDLE"

type WatchOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	cli.RefresherOpts
	name          string
	isFlexCluster bool
	store         store.ClusterDescriber
}

var watchTemplate = "\nCluster available.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) flexClusterWatcher(ctx context.Context) func() (any, bool, error) {
	return func() (any, bool, error) {
		result, err := opts.store.FlexCluster(opts.ConfigProjectID(), opts.name)
		if err != nil {
			apiError, ok := atlasv2.AsError(err)
			if !ok {
				return nil, false, err
			}

			if apiError.Error == http.StatusUnauthorized {
				// Refresh the access token
				// Note: this only updates the config, so we have to re-initialize the store
				if err := opts.RefreshAccessToken(ctx); err != nil {
					return nil, false, err
				}

				// Re-initialize store, refreshAccessToken only refreshes the config
				return nil, false, opts.initStore(ctx)()
			}
		}

		if err != nil {
			return nil, false, err
		}

		return nil, result.GetStateName() == idle, nil
	}
}

func (opts *WatchOpts) watcher(ctx context.Context) func() (any, bool, error) {
	return func() (any, bool, error) {
		result, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
		if err != nil {
			var atlasClustersPinnedErr *atlasClustersPinned.GenericOpenAPIError

			if errors.As(err, &atlasClustersPinnedErr) {
				if *atlasClustersPinnedErr.Model().Error == http.StatusUnauthorized {
					// Refresh the access token
					// Note: this only updates the config, so we have to re-initialize the store
					if err := opts.RefreshAccessToken(ctx); err != nil {
						return nil, false, err
					}

					// Re-initialize store, refreshAccessToken only refreshes the config
					return nil, false, opts.initStore(ctx)()
				}
			}
		}
		if err != nil {
			return nil, false, err
		}
		return nil, result.GetStateName() == idle, nil
	}
}

func (opts *WatchOpts) Run(ctx context.Context) error {
	if opts.isFlexCluster {
		if _, err := opts.Watch(opts.flexClusterWatcher(ctx)); err != nil {
			return err
		}
		return opts.Print(nil)
	}

	if _, err := opts.Watch(opts.watcher(ctx)); err != nil {
		return err
	}

	return opts.Print(nil)
}

// newIsFlexCluster sets the opts.isFlexCluster that indicates if the cluster to create is
// a FlexCluster.
func (opts *WatchOpts) newIsFlexCluster() error {
	_, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		var atlasClustersPinnedErr *atlasClustersPinned.GenericOpenAPIError
		if errors.As(err, &atlasClustersPinnedErr) {
			if *atlasClustersPinnedErr.Model().ErrorCode == cannotUseFlexWithClusterApisErrorCode {
				opts.isFlexCluster = true
				return nil
			}
		}

		return err
	}

	opts.isFlexCluster = false
	return nil
}

// WatchBuilder builds a cobra.Command that can run as:
// atlas cluster(s) watch <clusterName> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <clusterName>",
		Short: "Watch the specified cluster in your project until it becomes available.",
		Long: `This command checks the cluster's status periodically until it reaches an IDLE state. 
Once the cluster reaches the expected state, the command prints "Cluster available."
If you run the command in the terminal, it blocks the terminal session until the resource state changes to IDLE.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Watch for the cluster named myCluster to become available for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas clusters watch myCluster --projectId 5e2211c17a3e5a48f5497de3`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to watch.",
			"output":          watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
				opts.InitFlow(config.Default()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			if err := opts.newIsFlexCluster(); err != nil {
				return err
			}
			return opts.Run(cmd.Context())
		},
	}

	opts.AddProjectOptsFlags(cmd)

	return cmd
}
