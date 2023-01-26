// Copyright 2022 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id          string
	clusterName string
	store       store.RestoreJobsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var restoreDescribeTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT	URLs
{{.ID}}	{{.SnapshotID}}	{{.TargetClusterName}}	{{.DeliveryType}}	{{.ExpiresAt}}	{{range $index, $element := .DeliveryURL}}{{if $index}}; {{end}}{{$element}}{{end}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.RestoreJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas backup(s) restore(s) job(s) describe <ID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <restoreJobId>",
		Short: "Describe a cloud backup restore job.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"restoreJobIdDesc": "ID of the restore job.",
		},
		Example: fmt.Sprintf(`  # Return the details for the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the cluster named Cluster0:
  %s backup restore describe 507f1f77bcf86cd799439011 --clusterName Cluster0`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), restoreDescribeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
