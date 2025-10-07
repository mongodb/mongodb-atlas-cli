// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snapshots

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

const describeTemplate = `ID	SNAPSHOT TYPE	EXPIRES AT
{{.Id}}	{{.SnapshotType}}	{{.ExpiresAt}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=snapshots . ServerlessSnapshotsDescriber

type ServerlessSnapshotsDescriber interface {
	ServerlessSnapshot(string, string, string) (*atlasv2.ServerlessBackupSnapshot, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store       ServerlessSnapshotsDescriber
	snapshot    string
	clusterName string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.ServerlessSnapshot(opts.ConfigProjectID(), opts.clusterName, opts.snapshot)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas backup snapshots describe --snapshotId snapshotId --clusterName clusterName --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:         "describe",
		Short:       "Return the details for the specified snapshot for your project.",
		Long:        fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:        require.NoArgs,
		Annotations: map[string]string{"output": describeTemplate},
		Example: `  # Return the details for the backup snapshot with the ID 5f4007f327a3bd7b6f4103c5 for the instance named myDemo:
  atlas serverless backups snapshots describe --snapshotId 5f4007f327a3bd7b6f4103c5 --clusterName myDemo`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "please use the 'atlas backup snapshots describe' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.snapshot, flag.SnapshotID, "", usage.SnapshotID)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.SnapshotID)
	return cmd
}
