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
	"encoding/json"
	"errors"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20250219001/admin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	connectWaitSeconds = 10
	createTemplate     = "Search index created with ID: {{.IndexID}}\n"
	notFoundState      = "NOT_FOUND"
)

var (
	ErrSearchIndexDuplicated = errors.New("search index is duplicated")
	errInvalidIndex          = errors.New("invalid index")
)

type IndexID struct {
	ID         string
	Name       string
	Collection string
	Database   string
	Index      any
}

type CreateOpts struct {
	cli.WatchOpts
	cli.ProjectOpts
	cli.OutputOpts
	options.DeploymentOpts
	search.IndexOpts
	mongodbClient    mongodbclient.MongoDBClient
	connectionString string
	store            store.SearchIndexCreatorDescriber
	indexID          IndexID
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

	if err = opts.mongodbClient.Connect(ctx, opts.connectionString, connectWaitSeconds); err != nil {
		return err
	}
	defer func() {
		_ = opts.mongodbClient.Disconnect(ctx)
	}()

	idx, err := opts.CreateSearchIndex()
	if err != nil {
		return err
	}

	var definition any
	var idxType *string

	switch index := idx.(type) {
	case *admin.SearchIndexCreateRequest:
		idxType = index.Type

		opts.indexID.Database = index.Database
		opts.indexID.Collection = index.CollectionName
		opts.indexID.Name = index.Name

		definition = index.Definition
	case *admin.ClusterSearchIndex:
		_, _ = log.Warningln("you're using an old search index definition")
		idxType = index.Type

		opts.indexID.Database = index.Database
		opts.indexID.Collection = index.CollectionName
		opts.indexID.Name = index.Name

		definition, err = buildIndexDefinition(index)
		if err != nil {
			return err
		}
	default:
		return errInvalidIndex
	}

	if idxType == nil {
		defaultType := search.DefaultType
		idxType = &defaultType
	}

	telemetry.AppendOption(telemetry.WithSearchIndexType(*idxType))

	coll := opts.mongodbClient.Database(opts.indexID.Database).Collection(opts.indexID.Collection)
	if idx, _ := coll.SearchIndexByName(ctx, opts.indexID.Name); idx != nil {
		return ErrSearchIndexDuplicated
	}

	result, err := coll.CreateSearchIndex(ctx, opts.indexID.Name, *idxType, definition)
	if err != nil {
		return err
	}

	opts.indexID.Index = result

	if result.IndexID != nil {
		opts.indexID.ID = *result.IndexID
	}

	return nil
}

func buildIndexDefinition(idx *admin.ClusterSearchIndex) (any, error) {
	// To maintain formatting of the SDK, marshal object into JSON and then unmarshal into BSON
	jsonIndex, err := json.Marshal(idx)
	if err != nil {
		return nil, err
	}

	var index bson.D
	err = bson.UnmarshalExtJSON(jsonIndex, true, &index)
	if err != nil {
		return nil, err
	}

	// Empty these fields so that they are not included into the index definition for the MongoDB command
	index = removeFields(index, "id", "collectionName", "database", "name", "type", "status")
	return index, nil
}

func removeFields(doc bson.D, fields ...string) bson.D {
	cleanedDoc := bson.D{}

	for _, elem := range doc {
		if slices.Contains(fields, elem.Key) {
			continue
		}

		cleanedDoc = append(cleanedDoc, elem)
	}
	return cleanedDoc
}

func (opts *CreateOpts) RunAtlas() error {
	if err := opts.validateAndPrompt(); err != nil {
		return err
	}

	idx, err := opts.CreateSearchIndex()
	if err != nil {
		return err
	}

	switch index := idx.(type) {
	case *admin.SearchIndexCreateRequest:
		telemetry.AppendOption(telemetry.WithSearchIndexType(index.GetType()))
		r, err := opts.store.CreateSearchIndexes(opts.ConfigProjectID(), opts.DeploymentName, index)
		if err != nil {
			return err
		}

		if r.Database != nil {
			opts.indexID.Database = *r.Database
		}
		if r.CollectionName != nil {
			opts.indexID.Collection = *r.CollectionName
		}
		if r.Name != nil {
			opts.indexID.Name = *r.Name
		}
		if r.IndexID != nil {
			opts.indexID.ID = *r.IndexID
		}

		opts.indexID.Index = r

		return nil
	case *admin.ClusterSearchIndex:
		_, _ = log.Warningln("you're using an old search index definition")
		telemetry.AppendOption(telemetry.WithSearchIndexType(index.GetType()))
		r, err := opts.store.CreateSearchIndexesDeprecated(opts.ConfigProjectID(), opts.DeploymentName, index)
		if err != nil {
			return err
		}

		opts.indexID.Database = r.Database
		opts.indexID.Collection = r.CollectionName
		opts.indexID.Name = r.Name
		if r.IndexID != nil {
			opts.indexID.ID = *r.IndexID
		}

		opts.indexID.Index = r

		return nil
	default:
		return errInvalidIndex
	}
}

func (opts *CreateOpts) Run(ctx context.Context) error {
	_, err := opts.SelectDeployments(ctx, opts.ConfigProjectID(), options.IdleState)
	if err != nil {
		return err
	}

	if opts.IsLocalDeploymentType() {
		return opts.RunLocal(ctx)
	}

	return opts.RunAtlas()
}

func (opts *CreateOpts) initMongoDBClient() error {
	opts.mongodbClient = mongodbclient.NewClient()
	return nil
}

func (opts *CreateOpts) status(ctx context.Context) (*mongodbclient.SearchIndexDefinition, string, error) {
	if err := opts.mongodbClient.Connect(ctx, opts.connectionString, connectWaitSeconds); err != nil {
		return nil, notFoundState, err
	}
	defer func() {
		_ = opts.mongodbClient.Disconnect(ctx)
	}()

	coll := opts.mongodbClient.Database(opts.indexID.Database).Collection(opts.indexID.Collection)
	index, err := coll.SearchIndexByName(ctx, opts.indexID.Name)
	if err != nil {
		return index, notFoundState, nil
	}
	if index.Status == nil {
		return index, notFoundState, nil
	}
	return index, *index.Status, nil
}

func (opts *CreateOpts) watchLocal(ctx context.Context) (any, bool, error) {
	index, state, err := opts.status(ctx)
	if err != nil {
		return nil, false, err
	}
	if state == "READY" {
		index.Status = &state
		return index, true, nil
	}
	return nil, false, nil
}

func (opts *CreateOpts) watchAtlas(_ context.Context) (any, bool, error) {
	index, err := opts.store.SearchIndexDeprecated(opts.ConfigProjectID(), opts.DeploymentName, opts.indexID.ID)
	if err != nil {
		return nil, false, err
	}
	if index.GetStatus() == "STEADY" {
		return index, true, nil
	}
	return nil, false, nil
}

func (opts *CreateOpts) PostRun(ctx context.Context) error {
	opts.DeploymentTelemetry.AppendDeploymentType()
	opts.DeploymentTelemetry.AppendDeploymentUUID()

	if !opts.EnableWatch {
		return opts.Print(opts.indexID.Index)
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
	if err := opts.Print(watchResult); err != nil {
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
				opts.initMongoDBClient,
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
	opts.AddOutputOptFlags(cmd)

	// Local only
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DeploymentOpts.DBUserPassword, flag.Password, "", usage.Password)

	// Atlas only
	opts.AddProjectOptsFlags(cmd)

	cmd.MarkFlagsMutuallyExclusive(flag.Database, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.Collection, flag.File)
	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}
