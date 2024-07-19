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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store           store.SnapshotsCreator
	clusterName     string
	desc            string
	retentionInDays int
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Snapshot '{{.Id}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newCloudProviderSnapshot()

	r, err := opts.store.CreateSnapshot(opts.ConfigProjectID(), opts.clusterName, createRequest)
	if err != nil {
		return commonerrors.Check(err)
	}
	return opts.Print(r)
}

func (opts *CreateOpts) newCloudProviderSnapshot() *atlasv2.DiskBackupOnDemandSnapshotRequest {
	createRequest := &atlasv2.DiskBackupOnDemandSnapshotRequest{
		RetentionInDays: &opts.retentionInDays,
		Description:     &opts.desc,
	}
	return createRequest
}

// atlas backup snapshots create|take clusterName --desc description --retention days [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <clusterName>",
		Aliases: []string{"take"},
		Short:   "Create a backup snapshot for your project and cluster.",
		Long: `You can create on-demand backup snapshots for Atlas cluster tiers M10 and larger.

` + fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Owner"), "Atlas supports this command only for M10+ clusters."),
		Args: require.ExactArgs(1),
		Example: `  # Create a backup snapshot for the cluster named myDemo that Atlas retains for 30 days:
  atlas backups snapshots create myDemo --desc "test" --retention 30`,
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster whose snapshot you want to restore.",
			"output":          createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.clusterName = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.desc, flag.Description, "", usage.SnapshotDescription)
	cmd.Flags().IntVar(&opts.retentionInDays, flag.Retention, 1, usage.RetentionInDays)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Description)

	return cmd
}
