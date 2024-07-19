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

package onlinearchive

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type PauseOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id          string
	clusterName string
	store       store.OnlineArchiveUpdater
}

func (opts *PauseOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var pauseTemplate = "Online archive '{{.Id}}' paused.\n"

func (opts *PauseOpts) Run() error {
	cluster := &atlasv2.BackupOnlineArchive{
		Id:    &opts.id,
		State: pointer.Get("PAUSING"),
	}
	r, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, cluster)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas cluster(s) onlineArchive(s) pause <archiveId> [--clusterName name][--projectId projectId].
func PauseBuilder() *cobra.Command {
	opts := &PauseOpts{}
	cmd := &cobra.Command{
		Use:   "pause <archiveId>",
		Short: "Pause the specfied online archive for your cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"archiveIdDesc": "Unique identifier of the online archive to pause.",
			"output":        pauseTemplate,
		},
		Example: `  # Pause the online archive with the ID 5f189832e26ec075e10c32d3 for the cluster named myCluster:
  atlas clusters onlineArchives pause 5f189832e26ec075e10c32d3 --clusterName myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), pauseTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
