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

package atlas

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

type DataLakeCreateOpts struct {
	cli.GlobalOpts
	store store.DataLakeCreator
	Name  string
}

func (opts *DataLakeCreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DataLakeCreateOpts) Run() error {
	createRequest := mongodbatlas.DataLakeCreateRequest{
		Name: opts.Name,
	}

	result, err := opts.store.CreateDataLake(opts.ProjectID, &createRequest)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas datalake(s) create name --projectId projectId
func DataLakeCreateBuilder() *cobra.Command {
	opts := &DataLakeCreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: description.CreateDataLake,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
