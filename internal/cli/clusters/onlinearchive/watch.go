// Copyright 2021 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	id          string
	clusterName string
	store       store.OnlineArchiveDescriber
}

var watchTemplate = "\nOnline archive available.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (any, bool, error) {
	result, err := opts.store.OnlineArchive(opts.ConfigProjectID(), opts.clusterName, opts.id)
	if err != nil {
		return nil, false, err
	}
	return nil, result.GetState() != "PENDING" && result.GetState() != "PAUSING", nil
}

func (opts *WatchOpts) Run() error {
	if _, err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// WatchBuilder
//
//	atlas cluster(s) onlineArchive watch <archiveId> --clusterName=<name>	[--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <archiveId>",
		Short: "Watch for an archive to be available.",
		Long: `This command checks the archive's status periodically until it reaches a state different from PENDING or PAUSING. 
Once the archive reaches the expected status, the command prints "Online archive available."
If you run the command in the terminal, it blocks the terminal session until the resource status changes to the expected status.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  atlas cluster onlineArchive watch archiveIdSample --clusterName clusterNameSample`,
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"archiveIdDesc": "Unique identifier of the online archive to watch.",
			"output":        watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
