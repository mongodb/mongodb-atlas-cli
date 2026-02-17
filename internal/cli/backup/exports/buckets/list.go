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

package buckets

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=buckets . ExportBucketsLister

type ExportBucketsLister interface {
	ExportBuckets(string, *store.ListOptions) (*atlasv2.PaginatedBackupSnapshotExportBuckets, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	store ExportBucketsLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	BUCKET NAME	CLOUD PROVIDER	IAM ROLE ID{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.BucketName}}	{{.CloudProvider}}	{{.IamRoleId}}{{end}}
`

func (opts *ListOpts) Run() error {
	listOpts := opts.NewAtlasListOptions()
	r, err := opts.store.ExportBuckets(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas backup(s) export(s) bucket(s) list [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List cloud backup restore buckets for your project and cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Example: `  # Return all continuous backup export buckets for your project:
  atlas backup exports buckets list`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
