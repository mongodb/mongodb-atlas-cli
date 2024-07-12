// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodes

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type DeleteOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	*cli.DeleteOpts
	store store.SearchNodesDeleter
}

const atlasFtsDeploymentDoesNotExist = "ATLAS_SEARCH_DEPLOYMENT_DOES_NOT_EXIST"

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var deleteTemplate = "Started deleting search nodes for cluster '{{.}}'.\n"
var deleteWatchTemplate = "Search nodes for cluster '{{.}}' are deleted.\n"

func (opts *DeleteOpts) Run() error {
	if err := opts.Prompt(); err != nil {
		return err
	}

	if !opts.Confirm {
		fmt.Println(opts.FailMessage())
		return nil
	}
	if err := opts.store.DeleteSearchNodes(opts.ConfigProjectID(), opts.Entry); err != nil {
		return err
	}

	if opts.EnableWatch {
		if _, err := opts.Watch(opts.watcher); err != nil {
			return err
		}
		opts.Template = deleteWatchTemplate
	}

	return opts.Print(opts.Entry)
}

func (opts *DeleteOpts) watcher() (any, bool, error) {
	_, err := opts.store.SearchNodes(opts.ConfigProjectID(), opts.Entry)
	// Fallback case in case the backend starts returning 404 instead of the 400 it's returning currently
	target, ok := atlasv2.AsError(err)
	if ok && (target.GetErrorCode() == atlasFtsDeploymentDoesNotExist || target.GetError() == 404) {
		return nil, true, nil
	}

	return nil, false, err
}

// atlas clusters search nodes delete [--clusterName clusterName].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts(deleteTemplate, "Search node not deleted."),
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a search node for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Organization Owner or Project Owner"),
		Args:  require.NoArgs,
		Example: `  # Delete a search node for the cluster named myCluster:
  atlas clusters search nodes delete --clusterName myCluster`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), deleteTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	// Command specific flags
	cmd.Flags().StringVar(&opts.Entry, flag.ClusterName, "", usage.ClusterName)
	_ = cmd.MarkFlagRequired(flag.ClusterName)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().UintVar(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	// Global flags
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
