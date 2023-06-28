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

package peering

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	id    string
	store store.PeeringConnectionDescriber
}

var watchTemplate = "\nNetwork peering changes completed.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const (
	waitingForUser    = "WAITING_FOR_USER"
	failed            = "FAILED"
	available         = "AVAILABLE"
	pendingAcceptance = "PENDING_ACCEPTANCE"
)

func (opts *WatchOpts) watcher() (bool, error) {
	result, err := opts.store.PeeringConnection(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return false, err
	}

	switch v := result.(type) {
	case *atlasv2.AwsNetworkPeeringConnectionSettings:
		return watcherAWS(v), nil
	case *atlasv2.AzureNetworkPeeringConnectionSettings:
		return watcherAzure(v), nil
	case *atlasv2.GCPNetworkPeeringConnectionSettings:
		return watcherGCP(v), nil
	}
	return false, nil
}

func watcherGCP(peer *atlasv2.GCPNetworkPeeringConnectionSettings) bool {
	return *peer.Status == waitingForUser || *peer.Status == failed || *peer.Status == available
}

func watcherAzure(peer *atlasv2.AzureNetworkPeeringConnectionSettings) bool {
	return *peer.Status == waitingForUser || *peer.Status == failed || *peer.Status == available
}

func watcherAWS(peer *atlasv2.AwsNetworkPeeringConnectionSettings) bool {
	return *peer.StatusName == pendingAcceptance || *peer.StatusName == failed || *peer.StatusName == available
}

func (opts *WatchOpts) Run() error {
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// mongocli atlas networking peering watch <ID> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <peerId>",
		Short: "Watch the specified peering connection in your project until it becomes available.",
		Long: `This command checks the peering connection's status periodically until it becomes available. 
Once it reaches the expected state, the command prints "Network peering changes completed."
If you run the command in the terminal, it blocks the terminal session until the resource is available.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: fmt.Sprintf(`  Watch for the network peering connection with the ID 5f621dc701240c5b7c3a888e to become available in the project with the ID 5e2211c17a3e5a48f5497de3:
  %s networking peering watch 5f621dc701240c5b7c3a888e --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"peerIdDesc": "Unique ID of the network peering connection that you want to watch.",
			"output":     watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	return cmd
}
