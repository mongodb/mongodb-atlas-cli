// Copyright 2021 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

const describeTemplate = `ID	SNAPSHOT TYPE	TYPE	DESCRIPTION	EXPIRES AT
{{.Id}}	{{.SnapshotType}}	{{.Type}}	{{.Description}}	{{.ExpiresAt}}
`

const describeTemplateFlex = `ID	STATUS	MONGODB VERSION	START TIME	FINISH TIME	EXPIRATION
{{.Id}}	{{.Status}}	{{.MongoDBVersion}}	{{.StartTime}}	{{.FinishTime}}	{{.Expiration}}
`

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store       store.SnapshotsDescriber
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
	r, err := opts.store.FlexClusterSnapshot(opts.ConfigProjectID(), opts.clusterName, opts.snapshot)
	if err == nil {
		opts.Template = describeTemplateFlex
		return opts.Print(r)
	}

	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != cannotUseNotFlexWithFlexApisErrorCode && apiError.ErrorCode != featureUnsupported {
		return err
	}

	snapshots, err := opts.store.Snapshot(opts.ConfigProjectID(), opts.clusterName, opts.snapshot)
	if err != nil {
		return err
	}

	return opts.Print(snapshots)
}

// DescribeBuilder builds a cobra.Command that can run as:
// atlas backup snapshots describe snapshotId  --clusterName clusterName --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <snapshotId>",
		Aliases: []string{"get"},
		Short:   "Return the details for the specified snapshot for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Example: `  # Return the details for the backup snapshot with the ID 5f4007f327a3bd7b6f4103c5 for the cluster named myDemo:
  atlas backups snapshots describe 5f4007f327a3bd7b6f4103c5 --clusterName myDemo`,
		Annotations: map[string]string{
			"snapshotIdDesc": "Unique identifier of the snapshot you want to retrieve.",
			"output":         describeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.snapshot = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	return cmd
}
