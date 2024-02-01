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

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

const describeTemplate = `ID	SNAPSHOT TYPE	TYPE	DESCRIPTION	EXPIRES AT
{{.Id}}	{{.SnapshotType}}	{{.Type}}	{{.Description}}	{{.ExpiresAt}}
`

type DescribeOpts struct {
	cli.GlobalOpts
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
	r, err := opts.store.Snapshot(opts.ConfigProjectID(), opts.clusterName, opts.snapshot)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas backup snapshots describe snapshotId  --clusterName clusterName --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <snapshotId>",
		Short: "Return the details for the specified snapshot for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:  require.ExactArgs(1),
		Example: fmt.Sprintf(`  # Return the details for the backup snapshot with the ID 5f4007f327a3bd7b6f4103c5 for the cluster named myDemo:
  %s backups snapshots describe 5f4007f327a3bd7b6f4103c5 --clusterName myDemo`, cli.ExampleAtlasEntryPoint()),
		Annotations: map[string]string{
			"snapshotIdDesc": "Unique identifier of the snapshot you want to retrieve.",
			"output":         describeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.snapshot = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	return cmd
}
