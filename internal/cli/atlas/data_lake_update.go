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

const AWS = "AWS"

type DataLakeUpdateOpts struct {
	cli.GlobalOpts
	store      store.DataLakeUpdater
	Name       string
	Region     string
	Role       string
	TestBucket string
}

func (opts *DataLakeUpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DataLakeUpdateOpts) Run() error {
	updateRequest := mongodbatlas.DataLakeUpdateRequest{}

	if opts.Region != "" {
		updateRequest.DataProcessRegion = &mongodbatlas.DataProcessRegion{
			CloudProvider: AWS,
			Region:        opts.Region,
		}
	}

	if opts.Role != "" || opts.TestBucket != "" {
		updateRequest.CloudProviderConfig = &mongodbatlas.CloudProviderConfig{
			AWSConfig: mongodbatlas.AwsCloudProviderConfig{
				IAMAssumedRoleARN: opts.Role,
				TestS3Bucket:      opts.TestBucket,
			},
		}
	}

	result, err := opts.store.UpdateDataLake(opts.ProjectID, opts.Name, &updateRequest)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas datalake(s) update name --projectId projectId
func DataLakeUpdateBuilder() *cobra.Command {
	opts := &DataLakeUpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <name>",
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

	cmd.Flags().StringVar(&opts.Region, flag.Region, "", usage.DataLakeRegion)
	cmd.Flags().StringVar(&opts.Role, flag.Role, "", usage.DataLakeRole)
	cmd.Flags().StringVar(&opts.TestBucket, flag.TestBucket, "", usage.DataLakeTestBucket)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
