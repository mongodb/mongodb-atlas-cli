// Copyright 2023 MongoDB Inc
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

package indexes

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	deletedMessage      = "Index '%s' deleted"
	deleteMessageFailed = "Index not deleted"
)

type DeleteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	*cli.DeleteOpts
	options.DeploymentOpts
	search.IndexOpts
	mongodbClient mongodbclient.MongoDBClient
	store         store.SearchIndexDeleter
}

func (opts *DeleteOpts) Run(ctx context.Context) error {
	if _, err := opts.SelectDeployments(ctx, opts.ConfigProjectID()); err != nil {
		return err
	}

	if opts.Entry == "" {
		if err := promptRequiredName("Search Index ID", &opts.Entry); err != nil {
			return err
		}
	}

	if err := opts.Prompt(); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() {
		return opts.RunAtlas()
	}

	return opts.RunLocal(ctx)
}

func (opts *DeleteOpts) RunAtlas() error {
	return opts.Delete(opts.store.DeleteSearchIndex, opts.ConfigProjectID(), opts.DeploymentName)
}

func (opts *DeleteOpts) RunLocal(ctx context.Context) error {
	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	if err = opts.mongodbClient.Connect(connectionString, connectWaitSeconds); err != nil {
		return err
	}
	defer opts.mongodbClient.Disconnect()

	return opts.Delete(opts.mongodbClient.DeleteSearchIndex, opts.Entry)
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) initMongoDBClient(ctx context.Context) func() error {
	return func() error {
		opts.mongodbClient = mongodbclient.NewClientWithContext(ctx)
		return nil
	}
}

func (opts *DeleteOpts) PostRun() {
	opts.DeploymentTelemetry.AppendDeploymentType()
}

func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts(deletedMessage, deleteMessageFailed),
	}
	cmd := &cobra.Command{
		Use:     "delete <indexId>",
		Aliases: []string{"rm"},
		Short:   "Delete the specified search index from the specified deployment.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"indexIdDesc": "ID of the index.",
			"output":      opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.initMongoDBClient(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.Entry = args[0]
			}

			return opts.Run(cmd.Context())
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUserPassword, flag.Password, "", usage.Password)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
