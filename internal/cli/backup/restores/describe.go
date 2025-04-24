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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

type DescribeOpts struct {
	cli.ProjectOpts
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
{{.Id}}	{{.SnapshotId}}	{{.TargetClusterName}}	{{.DeliveryType}}	{{.ExpiresAt}}	{{range $index, $element := valueOrEmptySlice .DeliveryUrl}}{{if $index}}; {{end}}{{$element}}{{end}}
`

var restoreDescribeFlexClusterTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT	URLs
{{.Id}}	{{.SnapshotId}}	{{.TargetDeploymentItemName}}	{{.DeliveryType}}	{{.ExpirationDate}}	{{.SnapshotUrl}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.RestoreFlexClusterJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err == nil {
		opts.Template = restoreDescribeFlexClusterTemplate
		return opts.Print(r)
	}

	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != cannotUseNotFlexWithFlexApisErrorCode && apiError.ErrorCode != featureUnsupported {
		return err
	}

	restoreJob, err := opts.store.RestoreJob(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return err
	}

	return opts.Print(restoreJob)
}

// DescribeBuilder builds a cobra.Command that can run as:
// atlas backup(s) restore(s) job(s) describe <ID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <restoreJobId>",
		Aliases: []string{"get"},
		Short:   "Describe a cloud backup restore job.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"restoreJobIdDesc": "ID of the restore job.",
		},
		Example: `  # Return the details for the continuous backup restore job with the ID 507f1f77bcf86cd799439011 for the cluster named Cluster0:
  atlas backup restore describe 507f1f77bcf86cd799439011 --clusterName Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), restoreDescribeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
