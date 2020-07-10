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
	"fmt"

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
	name       string
	region     string
	role       string
	testBucket string
}

func (opts *DataLakeUpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DataLakeUpdateOpts) newUpdateRequest() *mongodbatlas.DataLakeUpdateRequest {
	updateRequest := &mongodbatlas.DataLakeUpdateRequest{}

	if opts.region != "" {
		updateRequest.DataProcessRegion = &mongodbatlas.DataProcessRegion{
			CloudProvider: AWS,
			Region:        opts.region,
		}
	}

	if opts.role != "" || opts.testBucket != "" {
		updateRequest.CloudProviderConfig = &mongodbatlas.CloudProviderConfig{
			AWSConfig: mongodbatlas.AwsCloudProviderConfig{
				IAMAssumedRoleARN: opts.role,
				TestS3Bucket:      opts.testBucket,
			},
		}
	}

	return updateRequest
}

func (opts *DataLakeUpdateOpts) Run() error {
	updateRequest := opts.newUpdateRequest()

	result, err := opts.store.UpdateDataLake(opts.ProjectID, opts.name, updateRequest)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas datalake(s) update name --projectId projectId [--role role] [--testBucket bucket] [--region region]
func DataLakeUpdateBuilder() *cobra.Command {
	opts := &DataLakeUpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <name>",
		Short: description.UpdateDataLake,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			if opts.region == "" && opts.role == "" && opts.testBucket == "" {
				return fmt.Errorf("must provide something to update")
			}

			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.DataLakeRegion)
	cmd.Flags().StringVar(&opts.role, flag.Role, "", usage.DataLakeRole)
	cmd.Flags().StringVar(&opts.testBucket, flag.TestBucket, "", usage.DataLakeTestBucket)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
