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

// This code was autogenerated at 2023-04-27T17:56:12+01:00. Note: Manual updates are allowed, but may be overwritten.

package availablesnapshots

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var listTemplate = `ID	DESCRIPTION	STATUS{{range .Results}}
{{if .DiskBackupReplicaSet}}{{.DiskBackupReplicaSet.Id}}	{{.DiskBackupReplicaSet.Description}}	{{.DiskBackupReplicaSet.Status}}{{else}}{{.DiskBackupShardedClusterSnapshot.Id}}	{{.DiskBackupShardedClusterSnapshot.Description}}	{{.DiskBackupShardedClusterSnapshot.Status}}{{end}}{{end}}`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	store store.PipelineAvailableSnapshotsLister

	pipelineName   string
	completedAfter string
}

var ErrCompletedAfterInvalidFormat = errors.New("--completedAfter flag is in an invalid format")

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func convertTime(value string) *time.Time {
	var result *time.Time
	if completedAfter, err := time.Parse(time.RFC3339, value); err == nil {
		result = &completedAfter
	}
	return result
}

func (opts *ListOpts) validate() error {
	if _, err := time.Parse(time.RFC3339, opts.completedAfter); opts.completedAfter != "" && err != nil {
		return fmt.Errorf("%w: expected format 'YYYY-MM-DD' got '%s'", ErrCompletedAfterInvalidFormat, opts.completedAfter)
	}
	return nil
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.PipelineAvailableSnapshots(opts.ConfigProjectID(), opts.pipelineName, convertTime(opts.completedAfter), opts.NewListOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dataLakePipelines availableSnapshots list [--projectId projectId].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all available backup snapshots for the specified data lake pipeline.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `# list available backup schedules for data lake pipeline called 'Pipeline1':
  atlas dataLakePipelines availableSnapshots list Pipeline1
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.pipelineName, flag.Pipeline, "", usage.Pipeline)
	_ = cmd.MarkFlagRequired(flag.Pipeline)

	cmd.Flags().StringVar(&opts.completedAfter, flag.CompletedAfter, "", usage.CompletedAfter)

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
