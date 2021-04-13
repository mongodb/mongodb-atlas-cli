// Copyright 2021 MongoDB Inc
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
package snapshots

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const describeTemplate = `ID	SNAPSHOT TYPE	TYPE	DESCRIPTION	EXPIRES AT
{{.ID}}	{{.SnapshotType}}	{{.Type}}	{{.Description}}	{{.ExpiresAt}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store       store.SnapshotsDescriber
	snapshot    string
	clusterName string
}

func (opts *DescribeOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.PublicAuthenticatedPreset(config.Default()))
	return err
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Snapshot(opts.ConfigProjectID(), opts.clusterName, opts.snapshot)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas backup snapshots describe snapshotID  --clusterName clusterName --projectId projectId
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <ID>",
		Short: "Get a specific snapshot for your project.",
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
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

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	return cmd
}
