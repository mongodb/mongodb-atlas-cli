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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id          string
	clusterName string
	store       store.ServerlessRestoreJobsDescriber
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
	}

	cmd.Flags().StringVar(&opts.id, flag.RestoreJobID, "", usage.RestoreJobID)
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.RestoreJobID)

	return cmd
}
