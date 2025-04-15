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

package restores

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	clusterName string
	store       store.RestoreJobsLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var restoreListTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.SnapshotId}}	{{.TargetClusterName}}	{{.DeliveryType}}	{{.ExpiresAt}}{{end}}
`
var restoreListFlexClusterTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.SnapshotId}}	{{.TargetDeploymentItemName}}	{{.DeliveryType}}	{{.ExpirationDate}}{{end}}
`

func (opts *ListOpts) Run() error {
	r, err := opts.store.RestoreFlexClusterJobs(opts.newListFlexBackupRestoreJobsAPIParams())
	if err == nil {
		opts.Template = restoreListFlexClusterTemplate
		return opts.Print(r)
	}

	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != cannotUseNotFlexWithFlexApisErrorCode && apiError.ErrorCode != featureUnsupported {
		return err
	}

	listOpts := opts.NewListOptions()
	restoreJobs, err := opts.store.RestoreJobs(opts.ConfigProjectID(), opts.clusterName, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(restoreJobs)
}

func (opts *ListOpts) newListFlexBackupRestoreJobsAPIParams() *atlasv2.ListFlexBackupRestoreJobsApiParams {
	includeCount := !opts.OmitCount
	return &atlasv2.ListFlexBackupRestoreJobsApiParams{
		GroupId:      opts.ConfigProjectID(),
		Name:         opts.clusterName,
		IncludeCount: &includeCount,
		ItemsPerPage: &opts.ItemsPerPage,
		PageNum:      &opts.PageNum,
	}
}

// ListBuilder builds a cobra.Command that can run as:
// atlas backup(s) restore(s) job(s) list <clusterName> [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list <clusterName>",
		Aliases: []string{"ls"},
		Short:   "Return all cloud backup restore jobs for your project and cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster for which you want to retrieve restore jobs.",
			"output":          restoreListTemplate,
		},
		Example: `  # Return all continuous backup restore jobs for the cluster Cluster0:
  atlas backup restore list Cluster0`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), restoreListTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.clusterName = args[0]

			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
