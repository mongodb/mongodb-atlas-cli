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

package onlinearchive

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName     string
	dbName          string
	collection      string
	dateField       string
	dateFormat      string
	archiveAfter    int
	partitions      []string
	expireAfterDays int
	store           store.OnlineArchiveCreator
	filename        string
	fs              afero.Fs
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Online archive '{{.Id}}' created.\n"

func (opts *CreateOpts) Run() error {
	archive, err := opts.newOnlineArchive()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newOnlineArchive() (*atlasv2.BackupOnlineArchiveCreate, error) {
	if opts.filename != "" {
		var archive *atlasv2.BackupOnlineArchiveCreate
		if err := file.Load(opts.fs, opts.filename, &archive); err != nil {
			return nil, err
		}
		return archive, nil
	}

	archive := &atlasv2.BackupOnlineArchiveCreate{
		CollName: opts.collection,
		Criteria: atlasv2.Criteria{
			DateField:       &opts.dateField,
			DateFormat:      &opts.dateFormat,
			ExpireAfterDays: pointer.Get(opts.archiveAfter),
		},
		DbName:          opts.dbName,
		PartitionFields: pointer.Get(opts.partitionFields()),
	}

	if opts.expireAfterDays > 0 {
		archive.DataExpirationRule = &atlasv2.DataExpirationRule{
			ExpireAfterDays: &opts.expireAfterDays,
		}
	}

	return archive, nil
}

func (opts *CreateOpts) partitionFields() []atlasv2.PartitionField {
	fields := make([]atlasv2.PartitionField, len(opts.partitions))
	for i, p := range opts.partitions {
		fields[i] = atlasv2.PartitionField{
			FieldName: p,
			Order:     i,
		}
	}
	return fields
}

// atlas cluster(s) onlineArchive(s) create [--clusterName clusterName] [--db dbName][--collection collection][--partition fieldName:fieldType][--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an online archive for a collection in the specified cluster.",
		Long: `You can create an online archive for an M10 or larger cluster.
		
To learn more about online archives, see https://www.mongodb.com/docs/atlas/online-archive/manage-online-archive/.

` + fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args: require.NoArgs,
		Example: `  # Create an online archive for the sample_mflix.movies collection in a cluster named myTestCluster when the current date is greater than the value of released date plus 2 days:
  atlas clusters onlineArchive create --clusterName myTestCluster --db sample_mflix --collection movies --dateField released --archiveAfter 2 --output json
  
  # Create an online archive for the sample_mflix.movies collection in a cluster named myTestCluster using a profile named egAtlasProfile when the current date is greater than the value of the released date plus 2 days. Data is partitioned based on the title field, year field, and released field from the documents in the collection:
  atlas clusters onlineArchive create --clusterName myTestCluster --db sample_mflix --collection movies --dateField released --archiveAfter 2 --partition title,year --output json -P egAtlasProfile `,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.dbName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.dateField, flag.DateField, "", usage.DateField)
	cmd.Flags().StringVar(&opts.dateFormat, flag.DateFormat, "ISODATE", usage.DateFormat)
	cmd.Flags().IntVar(&opts.archiveAfter, flag.ArchiveAfter, 0, usage.ArchiveAfter)
	cmd.Flags().IntVar(&opts.expireAfterDays, flag.ExpireAfterDays, 0, usage.ExpireAfterDays)
	cmd.Flags().StringSliceVar(&opts.partitions, flag.Partition, nil, usage.PartitionFields)
	cmd.Flags().StringVar(&opts.filename, flag.File, "", usage.OnlineArchiveFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	cmd.MarkFlagsRequiredTogether(flag.Database, flag.Collection, flag.DateField, flag.ArchiveAfter)

	cmd.MarkFlagsOneRequired(flag.Database, flag.Collection, flag.DateField, flag.ArchiveAfter, flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Database)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Collection)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DateField)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.DateFormat)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ArchiveAfter)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ExpireAfterDays)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Partition)

	return cmd
}
