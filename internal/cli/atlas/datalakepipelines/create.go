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

// This code was autogenerated at 2023-04-12T16:00:41+01:00. Note: Manual updates are allowed, but may be overwritten.

package datalakepipelines

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store        store.PipelinesCreator
	pipelineName string

	sinkType             string
	sinkMetadataProvider string
	sinkMetadataRegion   string
	sinkPartitionField   []string
	sourceType           string
	sourceClusterName    string
	sourceCollectionName string
	sourceDatabaseName   string
	sourcePolicyItemID   string
	transform            []string
	filename             string
	fs                   afero.Fs
}

var ErrSourceTypeInvalid = errors.New("--sourceType invalid")
var ErrSinkTransformInvalid = errors.New("--transform format invalid")

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = `Pipeline {{.Name}} created.`

func (opts *CreateOpts) Run() error {
	createRequest, err := opts.newCreateRequest()
	if err != nil {
		return err
	}

	r, err := opts.store.CreatePipeline(opts.ConfigProjectID(), *createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) validate() error {
	if !strings.EqualFold(opts.sourceType, periodicCPS) && !strings.EqualFold(opts.sourceType, onDemandCPS) {
		return fmt.Errorf("%w: expected either '%s' or '%s' got '%s'", ErrSourceTypeInvalid, periodicCPS, onDemandCPS, opts.sourceType)
	}

	for _, entry := range opts.transform {
		entries := strings.Split(entry, ":")
		if len(entries) != 2 || len(entries[0]) == 0 || len(entries[1]) == 0 {
			return fmt.Errorf("%w: expected format is 'type:field1,fieldN' got '%s'", ErrSinkTransformInvalid, entry)
		}
	}

	return nil
}

func (opts *CreateOpts) newCreateRequest() (*atlasv2.IngestionPipeline, error) {
	if opts.filename != "" {
		pipeline := &atlasv2.IngestionPipeline{}
		if err := file.Load(opts.fs, opts.filename, pipeline); err != nil {
			return nil, err
		}
		return pipeline, nil
	}

	pipeline := &atlasv2.IngestionPipeline{
		Name: &opts.pipelineName,
		Sink: &atlasv2.IngestionSink{
			DLSIngestionSink: &atlasv2.DLSIngestionSink{
				Type:             &opts.sinkType,
				MetadataProvider: &opts.sinkMetadataProvider,
				MetadataRegion:   &opts.sinkMetadataRegion,
			},
		},
		Source: &atlasv2.IngestionSource{},
	}

	for i, fieldName := range opts.sinkPartitionField {
		pipeline.Sink.DLSIngestionSink.PartitionFields = append(pipeline.Sink.DLSIngestionSink.PartitionFields, *atlasv2.NewPartitionField(fieldName, int32(i)))
	}

	for _, entry := range opts.transform {
		entries := strings.Split(entry, ":")
		transformType := entries[0]
		transformFieldNames := strings.Split(entries[1], ",")
		for i := range transformFieldNames {
			pipeline.Transformations = append(pipeline.Transformations, atlasv2.FieldTransformation{Field: &transformFieldNames[i], Type: &transformType})
		}
	}

	if strings.EqualFold(opts.sourceType, periodicCPS) {
		pipeline.Source.PeriodicCpsSnapshotSource = &atlasv2.PeriodicCpsSnapshotSource{
			Type:           &opts.sourceType,
			ClusterName:    &opts.sourceClusterName,
			CollectionName: &opts.sourceCollectionName,
			DatabaseName:   &opts.sourceDatabaseName,
			PolicyItemId:   &opts.sourcePolicyItemID,
		}
	} else if strings.EqualFold(opts.sourceType, onDemandCPS) {
		pipeline.Source.OnDemandCpsSnapshotSource = &atlasv2.OnDemandCpsSnapshotSource{
			Type:           &opts.sourceType,
			ClusterName:    &opts.sourceClusterName,
			CollectionName: &opts.sourceCollectionName,
			DatabaseName:   &opts.sourceDatabaseName,
		}
	}

	return pipeline, nil
}

// atlas dataLakePipelines create <pipelineName> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{fs: afero.NewOsFs()}
	cmd := &cobra.Command{
		Use:   "create <pipelineName>",
		Short: "Creates a new Data Lake Pipeline.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"pipelineNameDesc": "Label that identifies the pipeline",
		},
		Example: `# create data lake pipeline:
  atlas dataLakePipelines create Pipeline1 --sinkType CPS --sinkMetadataProvider AWS --sinkMetadataRegion us-east-1 --sinkPartitionField name:0,summary:1 --sourceType PERIODIC_CPS --sourceClusterName Cluster1 --sourceDatabaseName sample_airbnb --sourceCollectionName listingsAndReviews --sourcePolicyItemId 507f1f77bcf86cd799439011 --transform EXCLUDE:space,EXCLUDE:notes`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.pipelineName = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.sinkType, flag.SinkType, "", usage.SinkType)
	cmd.Flags().StringVar(&opts.sinkMetadataProvider, flag.SinkMetadataProvider, "", usage.SinkMetadataProvider)
	cmd.Flags().StringVar(&opts.sinkMetadataRegion, flag.SinkMetadataRegion, "", usage.SinkMetadataRegion)
	cmd.Flags().StringSliceVar(&opts.sinkPartitionField, flag.SinkPartitionField, nil, usage.SinkPartitionField)
	cmd.Flags().StringVar(&opts.sourceType, flag.SourceType, "", usage.SourceType)
	cmd.Flags().StringVar(&opts.sourceClusterName, flag.SourceClusterName, "", usage.SourceClusterName)
	cmd.Flags().StringVar(&opts.sourceCollectionName, flag.SourceCollectionName, "", usage.SourceCollectionName)
	cmd.Flags().StringVar(&opts.sourceDatabaseName, flag.SourceDatabaseName, "", usage.SourceDatabaseName)
	cmd.Flags().StringVar(&opts.sourcePolicyItemID, flag.SourcePolicyItemID, "", usage.SourcePolicyItemID)
	cmd.Flags().StringSliceVar(&opts.transform, flag.Transform, nil, usage.Transform)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.PipelineFilename)

	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.RegisterFlagCompletionFunc(flag.SourceType, autoCompleteSourceType)

	cmd.MarkFlagsMutuallyExclusive(flag.SinkType, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SinkMetadataProvider, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SinkMetadataRegion, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SinkPartitionField, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SourceType, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SourceClusterName, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SourceCollectionName, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SourceDatabaseName, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.SourcePolicyItemID, flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.Transform, flag.File)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
