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

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	namePattern        = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
	connectWaitSeconds = 10
	createTemplate     = "Search index created\n"
	notFoundState      = "NOT_FOUND"
)

type CreateOpts struct {
	cli.WatchOpts
	cli.GlobalOpts
	cli.OutputOpts
	options.DeploymentOpts
	search.IndexOpts
	mongodbClient    mongodbclient.MongoDBClient
	connectionString string
	index            *admin.ClusterSearchIndex
}

func (opts *CreateOpts) Run(ctx context.Context) error {
	var err error
	if err := opts.PodmanClient.Ready(ctx); err != nil {
		return err
	}

	if err := opts.validateAndPrompt(ctx); err != nil {
		return err
	}

	opts.connectionString, err = opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	if err = opts.mongodbClient.Connect(opts.connectionString, connectWaitSeconds); err != nil {
		return err
	}
	defer opts.mongodbClient.Disconnect()

	opts.index, err = opts.NewSearchIndex()
	if err != nil {
		return err
	}

	return opts.mongodbClient.Database(opts.index.Database).CreateSearchIndex(ctx, opts.index.CollectionName, opts.index)
}

func (opts *CreateOpts) initMongoDBClient(ctx context.Context) func() error {
	return func() error {
		opts.mongodbClient = mongodbclient.NewClientWithContext(ctx)
		return nil
	}
}

func (opts *CreateOpts) status(ctx context.Context) (string, error) {
	if err := opts.mongodbClient.Connect(opts.connectionString, connectWaitSeconds); err != nil {
		return "", err
	}
	defer opts.mongodbClient.Disconnect()

	db := opts.mongodbClient.Database(opts.index.Database)
	col := db.Collection(opts.index.CollectionName)
	cursor, err := col.Aggregate(ctx, mongo.Pipeline{
		{
			{Key: "$listSearchIndexes", Value: bson.D{}},
		},
	})
	if err != nil {
		return "", err
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return "", err
	}
	if len(results) == 0 {
		return notFoundState, nil
	}
	status, ok := results[0]["status"].(string)
	if !ok {
		return notFoundState, nil
	}
	return status, nil
}

func (opts *CreateOpts) watch(ctx context.Context) (bool, error) {
	state, err := opts.status(ctx)
	if err != nil {
		return false, err
	}
	if state == "STEADY" {
		return true, nil
	}
	return false, nil
}

func (opts *CreateOpts) PostRun(ctx context.Context) error {
	if !opts.EnableWatch {
		return opts.Print(nil)
	}

	if err := opts.Watch(func() (bool, error) {
		return opts.watch(ctx)
	}); err != nil {
		return err
	}

	return opts.Print(nil)
}

func (opts *CreateOpts) validateAndPrompt(ctx context.Context) error {
	if opts.DeploymentName != "" {
		if err := opts.DeploymentOpts.CheckIfDeploymentExists(ctx); err != nil {
			return err
		}
	} else {
		if err := opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	}

	if opts.Filename != "" {
		return nil
	}

	if opts.Name == "" {
		if err := promptRequiredName("Search Index Name", &opts.Name); err != nil {
			return err
		}
	}

	if opts.DBName == "" {
		if err := promptRequiredName("Database", &opts.DBName); err != nil {
			return err
		}
	}

	if opts.Collection == "" {
		if err := promptRequiredName("Collection", &opts.Collection); err != nil {
			return err
		}
	}

	return nil
}

func promptRequiredName(message string, response *string) error {
	return telemetry.TrackAskOne(
		&survey.Input{Message: message},
		response,
		survey.WithValidator(survey.Required),
	)
}

func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		IndexOpts: search.IndexOpts{
			Analyzer: search.DefaultAnalyzer,
			Dynamic:  true,
			Fs:       afero.NewOsFs(),
		},
	}

	cmd := &cobra.Command{
		Use:   "create [indexName]",
		Short: "Create a search index for the specified deployment.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"indexNameDesc": "Name of the index.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			opts.PodmanClient = podman.NewClient(log.IsDebugLevel(), w)
			return opts.PreRunE(
				opts.InitOutput(w, createTemplate),
				opts.InitStore(opts.PodmanClient),
				opts.initMongoDBClient(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.Name = args[0]
			}
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PostRun(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().StringVar(&opts.DBName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.Collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVarP(&opts.Filename, flag.File, flag.FileShort, "", usage.SearchFilename)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().UintVar(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Database)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Collection)

	return cmd
}
