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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

var listTemplate = `ID	NAME	DATABASE	COLLECTION	STATUS	TYPE{{range valueOrEmptySlice .}}
{{.IndexID}}	{{.Name}}	{{.Database}}	{{.CollectionName}}	{{.Status}}	{{if .Type}}{{.Type}}{{else}}` + search.DefaultType + `{{end}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=indexes . Lister

type Lister interface {
	SearchIndexesDeprecated(string, string, string, string) ([]atlasv2.ClusterSearchIndex, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	options.DeploymentOpts
	search.IndexOpts
	mongodbClient mongodbclient.MongoDBClient
	store         Lister
}

func (opts *ListOpts) Run(ctx context.Context) error {
	if _, err := opts.SelectDeployments(ctx, opts.ConfigProjectID(), options.IdleState, options.UpdatingState); err != nil {
		return err
	}

	if err := opts.validateAndPrompt(); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() {
		return opts.RunAtlas()
	}

	return opts.RunLocal(ctx)
}

func (opts *ListOpts) RunAtlas() error {
	r, err := opts.store.SearchIndexesDeprecated(opts.ConfigProjectID(), opts.DeploymentName, opts.DBName, opts.Collection)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *ListOpts) RunLocal(ctx context.Context) error {
	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	if err = opts.mongodbClient.Connect(ctx, connectionString, connectWaitSeconds); err != nil {
		return err
	}
	defer func() {
		_ = opts.mongodbClient.Disconnect(ctx)
	}()

	r, err := opts.mongodbClient.Database(opts.DBName).Collection(opts.Collection).SearchIndexes(ctx)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) initMongoDBClient() error {
	opts.mongodbClient = mongodbclient.NewClient()
	return nil
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) validateAndPrompt() error {
	if opts.DBName == "" {
		if err := promptRequiredName("Database", &opts.DBName); err != nil {
			return err
		}
	}

	if opts.Collection == "" {
		return promptRequiredName("Collection", &opts.Collection)
	}

	return nil
}

func (opts *ListOpts) PostRun() {
	opts.DeploymentTelemetry.AppendDeploymentType()
	opts.DeploymentTelemetry.AppendDeploymentUUID()
}

func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all Atlas Search indexes for a deployment.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		GroupID: "all",
		Annotations: map[string]string{
			"output": listTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			w := cmd.OutOrStdout()
			return opts.PreRunE(
				opts.InitOutput(w, listTemplate),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.initMongoDBClient,
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DBName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.Collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
