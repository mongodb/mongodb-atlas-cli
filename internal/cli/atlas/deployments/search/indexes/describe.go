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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var describeTemplate = `ID	NAME	DATABASE	COLLECTION
{{.IndexID}}	{{.Name}}	{{.Database}}	{{.CollectionName}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	options.DeploymentOpts
	indexID       string
	mongodbClient mongodbclient.MongoDBClient
	store         store.SearchIndexDescriber
}

func (opts *DescribeOpts) Run(ctx context.Context) error {
	if err := opts.validateAndPrompt(ctx); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() {
		return opts.RunAtlas()
	}

	return opts.RunLocal(ctx)
}

func (opts *DescribeOpts) RunAtlas() error {
	r, err := opts.store.SearchIndex(opts.ConfigProjectID(), opts.DeploymentName, opts.indexID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) RunLocal(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	if err = opts.mongodbClient.Connect(connectionString, connectWaitSeconds); err != nil {
		return err
	}
	defer opts.mongodbClient.Disconnect()

	r, err := opts.mongodbClient.SearchIndex(opts.indexID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) initMongoDBClient(ctx context.Context) func() error {
	return func() error {
		opts.mongodbClient = mongodbclient.NewClientWithContext(ctx)
		return nil
	}
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) validateAndPrompt(ctx context.Context) error {
	if err := opts.ValidateAndPromptDeploymentType(); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() && opts.DeploymentName == "" {
		return ErrNoDeploymentName
	}

	if opts.DeploymentName == "" {
		if err := opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	}

	if opts.indexID == "" {
		if err := promptRequiredName("Search Index ID", &opts.indexID); err != nil {
			return err
		}
	}

	return nil
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			opts.PodmanClient = podman.NewClient(log.IsDebugLevel(), w)
			return opts.PreRunE(
				opts.InitOutput(w, describeTemplate),
				opts.InitStore(opts.PodmanClient),
				opts.initStore(cmd.Context()),
				opts.initMongoDBClient(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.indexID = args[0]
			}
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
