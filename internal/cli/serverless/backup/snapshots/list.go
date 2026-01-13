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

package snapshots

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=snapshots . ServerlessSnapshotsLister

type ServerlessSnapshotsLister interface {
	ServerlessSnapshots(string, string, *store.ListOptions) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	clusterName string
	store       ServerlessSnapshotsLister
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

func (opts *ListOpts) Run() error {
	listOpts := opts.NewAtlasListOptions()
	r, err := opts.store.ServerlessSnapshots(opts.ConfigProjectID(), opts.clusterName, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas serverless backups snapshots list <clusterName> [--projectId projectId] [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list <clusterName>",
		Short:   "Return all cloud backup snapshots for the specified serverless instance in your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Label that identifies the Atlas serverless instance that contains the snapshots you want to return.",
		},
		Example: `  # Return a JSON-formatted list of snapshots for the instance named myDemo 
  atlas serverless backups snapshots list myDemo --output json`,
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
		Deprecated: "please use the 'atlas backup snapshots list' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
