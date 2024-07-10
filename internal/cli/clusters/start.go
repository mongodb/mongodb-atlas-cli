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

package clusters

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
)

type StartOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.ClusterStarter
}

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var startTmpl = "Starting cluster '{{.Name}}'.\n"

func (opts *StartOpts) Run() error {
	r, err := opts.store.StartCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

// atlas cluster(s) start <clusterName> [--projectId projectId].
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start <clusterName>",
		Short: "Start the specified paused MongoDB cluster.",
		Long:  fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Cluster Manager"), "Atlas supports this command only for M10+ clusters."),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to start.",
			"output":          startTmpl,
		},
		Example: `  # Start a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas clusters start myCluster --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), startTmpl),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
