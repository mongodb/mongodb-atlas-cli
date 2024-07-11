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
	"errors"

	"github.com/AlecAivazis/survey/v2"
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
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	connectWaitSeconds = 10
	createTemplate     = "Search index created with ID: {{.IndexID}}\n"
	notFoundState      = "NOT_FOUND"
)

var ErrSearchIndexDuplicated = errors.New("search index is duplicated")

type CreateOpts struct {
	cli.WatchOpts
	cli.GlobalOpts
	cli.OutputOpts
	options.DeploymentOpts
	search.IndexOpts
	mongodbClient    mongodbclient.MongoDBClient
	connectionString string
	index            *admin.ClusterSearchIndex
	store            store.SearchIndexCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) RunLocal(ctx context.Context) error {
	var err error

	if err = opts.validateAndPrompt(); err != nil {
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

	opts.index, err = opts.DeprecatedNewSearchIndex()
	if err != nil {
		return err
	}

	db := opts.mongodbClient.Database(opts.index.Database)
	if idx, _ := db.SearchIndexByName(ctx, opts.index.Name, opts.index.CollectionName); idx != nil {
		return ErrSearchIndexDuplicated
	}

	opts.index, err = db.CreateSearchIndex(ctx, opts.index.CollectionName, opts.index)
	return err
}

func (opts *CreateOpts) RunAtlas() error {
	if err := opts.validateAndPrompt(); err != nil {
		return err
	}

	index, err := opts.NewSearchIndex()
	if err != nil {
		return err
	}

	opts.index, err = opts.store.CreateSearchIndexes(opts.ConfigProjectID(), opts.DeploymentName, index)
	return err
}

func (opts *CreateOpts) Run(ctx context.Context) error {
	_, err := opts.SelectDeployments(ctx, opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if opts.IsLocalDeploymentType() {
		return opts.RunLocal(ctx)
	}

	return opts.RunAtlas()
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

func (opts *CreateOpts) watchLocal(ctx context.Context) (any, bool, error) {
	state, err := opts.status(ctx)
	if err != nil {
		return nil, false, err
	}
	if state == "READY" {
		opts.index.Status = &state
		return opts.index, true, nil
	}
	return nil, false, nil
}

func (opts *CreateOpts) watchAtlas(_ context.Context) (any, bool, error) {
	index, err := opts.store.SearchIndex(opts.ConfigProjectID(), opts.DeploymentName, *opts.index.IndexID)
	if err != nil {
		return nil, false, err
	}
	if index.GetStatus() == "STEADY" {
		return index, true, nil
	}
	return nil, false, nil
}

func (opts *CreateOpts) PostRun(ctx context.Context) error {
	opts.AppendDeploymentType()
	if !opts.EnableWatch {
		return opts.Print(opts.index)
	}

	watch := opts.watchLocal
	if opts.IsAtlasDeploymentType() {
		watch = opts.watchAtlas
	}

	watchResult, err := opts.Watch(func() (any, bool, error) {
		return watch(ctx)
	})

	if err != nil {
		return err
	}
	opts.index = watchResult.(*admin.ClusterSearchIndex)

	if err := opts.Print(opts.index); err != nil {
		return err
	}
	return opts.PostRunMessages()
}

func (opts *CreateOpts) validateAndPrompt() error {
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
		Use:     "create [indexName]",
		Short:   "Create a search index for the specified deployment.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"indexNameDesc": "Name of the index.",
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			w := cmd.OutOrStdout()
			opts.WatchOpts.OutWriter = w

			if opts.Filename != "" && (opts.DBName != "" || opts.Collection != "") {
				return errors.New("the '-file' flag cannot be used in conjunction with the 'db' and 'collection' flags, please choose either 'file' or 'db' and '-collection', but not both")
			}

			return opts.PreRunE(
				opts.InitOutput(w, createTemplate),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.initStore(cmd.Context()),
				opts.initMongoDBClient(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.Name = args[0]
			}
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PostRun(cmd.Context())
		},
	}

	// Atlas and Local
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)
	cmd.Flags().StringVar(&opts.DBName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.Collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVarP(&opts.Filename, flag.File, flag.FileShort, "", usage.SearchFilename)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	// Local only
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().UintVar(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUserPassword, flag.Password, "", usage.Password)

	// Atlas only
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.MarkFlagsMutuallyExclusive(flag.Database, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.Collection, flag.File)
	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}
