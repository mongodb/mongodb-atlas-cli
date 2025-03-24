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

package snapshots

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
	store       store.SnapshotsLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	TYPE	STATUS	CREATED AT	EXPIRES AT{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.SnapshotType}}	{{.Status}}	{{.CreatedAt}}	{{.ExpiresAt}}{{end}}
`
var listTemplateFlex = `ID	STATUS	MONGODB VERSION	START TIME	FINISH TIME	EXPIRATION{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Status}}	{{.MongoDBVersion}}	{{.StartTime}}	{{.FinishTime}}	{{.Expiration}}{{end}}
`

func (opts *ListOpts) Run() error {
	r, err := opts.store.FlexClusterSnapshots(opts.newListFlexBackupsAPIParams())
	if err == nil {
		opts.Template = listTemplateFlex
		return opts.Print(r)
	}

	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != cannotUseNotFlexWithFlexApisErrorCode && apiError.ErrorCode != featureUnsupported {
		return err
	}

	listOpts := opts.NewAtlasListOptions()
	snapshotsList, err := opts.store.Snapshots(opts.ConfigProjectID(), opts.clusterName, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(snapshotsList)
}

func (opts *ListOpts) newListFlexBackupsAPIParams() *atlasv2.ListFlexBackupsApiParams {
	includeCount := !opts.OmitCount
	return &atlasv2.ListFlexBackupsApiParams{
		GroupId:      opts.ConfigProjectID(),
		Name:         opts.clusterName,
		IncludeCount: &includeCount,
		ItemsPerPage: &opts.ItemsPerPage,
		PageNum:      &opts.PageNum,
	}
}

// ListBuilder builds a cobra.Command that can run as:
// atlas backups snapshots list <clusterName> [--projectId projectId] [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list <clusterName>",
		Short:   "Return all cloud backup snapshots for your project and cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster that contains the snapshots you want to retrieve.",
			"output":          listTemplate,
		},
		Example: `  # Return a JSON-formatted list of snapshots for the cluster named myDemo 
  atlas backups snapshots list myDemo --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
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
