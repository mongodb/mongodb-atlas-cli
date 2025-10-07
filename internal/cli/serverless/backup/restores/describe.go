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

package restores

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=restores . ServerlessRestoreJobsDescriber

type ServerlessRestoreJobsDescriber interface {
	ServerlessRestoreJob(string, string, string) (*atlasv2.ServerlessBackupRestoreJob, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	id          string
	clusterName string
	store       ServerlessRestoreJobsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var restoreDescribeTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT	URLs
{{.Id}}	{{.SnapshotId}}	{{.TargetClusterName}}	{{.DeliveryType}}	{{.ExpiresAt}}	{{range $index, $element := .DeliveryUrl}}{{if $index}}; {{end}}{{$element}}{{end}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.ServerlessRestoreJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas serverless backup(s) restore(s) job(s) describe.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a cloud backup restore job.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.NoArgs,
		Example: `  # Return the details for the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the serverless isntance named Cluster0:
  atlas serverless backup restore describe --restoreJobId 507f1f77bcf86cd799439011 --clusterName Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), restoreDescribeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "please use the 'atlas backup restores describe' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	cmd.Flags().StringVar(&opts.id, flag.RestoreJobID, "", usage.RestoreJobID)
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.RestoreJobID)

	return cmd
}
