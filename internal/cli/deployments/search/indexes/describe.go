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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

var describeTemplate = `ID	NAME	DATABASE	COLLECTION	STATUS	TYPE
{{.IndexID}}	{{.Name}}	{{.Database}}	{{.CollectionName}}	{{.Status}}	{{if .Type}}{{.Type}}{{else}}` + search.DefaultType + `{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=indexes . Describer

type Describer interface {
	SearchIndexDeprecated(string, string, string) (*atlasv2.ClusterSearchIndex, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	options.DeploymentOpts
	indexID       string
	mongodbClient mongodbclient.MongoDBClient
	store         Describer
}

func (opts *DescribeOpts) Run(ctx context.Context) error {
	_, err := opts.SelectDeployments(ctx, opts.ConfigProjectID(), options.IdleState)
	if err != nil {
		return err
	}

	if opts.indexID == "" {
		if err := promptRequiredName("Search Index ID", &opts.indexID); err != nil {
			return err
		}
	}
	if opts.IsAtlasDeploymentType() {
		return opts.RunAtlas()
	}

	return opts.RunLocal(ctx)
}

func (opts *DescribeOpts) RunAtlas() error {
	r, err := opts.store.SearchIndexDeprecated(opts.ConfigProjectID(), opts.DeploymentName, opts.indexID)
	if err != nil {
		return err
	}

	telemetry.AppendOption(telemetry.WithSearchIndexType(r.GetType()))

	return opts.Print(r)
}

func (opts *DescribeOpts) RunLocal(ctx context.Context) error {
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

	r, err := opts.mongodbClient.SearchIndex(ctx, opts.indexID)
	if err != nil {
		return err
	}

	if r.Type != nil {
		telemetry.AppendOption(telemetry.WithSearchIndexType(*r.Type))
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) initMongoDBClient() error {
	opts.mongodbClient = mongodbclient.NewClient()
	return nil
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) PostRun() {
	opts.DeploymentTelemetry.AppendDeploymentType()
	opts.DeploymentTelemetry.AppendDeploymentUUID()
}

func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe [indexId]",
		Short:   "Describe a search index for the specified deployment.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"indexIdDesc": "ID of the index.",
			"output":      describeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			w := cmd.OutOrStdout()
			return opts.PreRunE(
				opts.InitOutput(w, describeTemplate),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.initMongoDBClient,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.indexID = args[0]
			}
			return opts.Run(cmd.Context())
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
