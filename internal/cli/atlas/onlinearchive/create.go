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
	"fmt"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	clusterName  string
	dbName       string
	collection   string
	dateField    string
	archiveAfter float64
	partitions   []string
	store        store.OnlineArchiveCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	archive, err := opts.newOnlineArchive()
	if err != nil {
		return err
	}
	result, err := opts.store.CreateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *CreateOpts) newOnlineArchive() (*atlas.OnlineArchive, error) {
	partitions, err := opts.partitionFields()
	if err != nil {
		return nil, err
	}
	a := &atlas.OnlineArchive{
		CollName: opts.collection,
		Criteria: &atlas.OnlineArchiveCriteria{
			DateField:       opts.dateField,
			ExpireAfterDays: opts.archiveAfter,
		},
		DBName:          opts.dbName,
		PartitionFields: partitions,
	}
	return a, nil
}
func (opts *CreateOpts) partitionFields() ([]*atlas.PartitionFields, error) {
	fields := make([]*atlas.PartitionFields, len(opts.partitions))
	for i, p := range opts.partitions {
		f := strings.Split(p, ":")
		if len(f) != 2 {
			return nil, fmt.Errorf("invalid partition, got: %s", p)
		}
		order := float64(i)
		fields[i] = &atlas.PartitionFields{
			FieldName: f[0],
			FieldType: f[1],
			Order:     &order,
		}
	}
	return fields, nil
}

// mongocli atlas cluster(s) onlineArchive(s) create [--clusterName clusterName] [--db dbName][--collection collection][--projectId projectId]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateOnlineArchive,
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.dbName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.dateField, flag.DateField, "", usage.DateField)
	cmd.Flags().Float64Var(&opts.archiveAfter, flag.ArchiveAfter, 0, usage.ArchiveAfter)
	cmd.Flags().StringSliceVar(&opts.partitions, flag.Partition, nil, usage.PartitionFields)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.Database)
	_ = cmd.MarkFlagRequired(flag.Collection)
	_ = cmd.MarkFlagRequired(flag.DateField)
	_ = cmd.MarkFlagRequired(flag.ArchiveAfter)

	return cmd
}
