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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=pause_mock_test.go -package=clusters . ClusterPauser

type ClusterPauser interface {
	PauseCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
}

type PauseOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name  string
	store ClusterPauser
}

func (opts *PauseOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var pauseTmpl = "Pausing cluster '{{.Name}}'.\n"

func (opts *PauseOpts) Run() error {
	r, err := opts.store.PauseCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return commonerrors.Check(err)
	}

	return opts.Print(r)
}

// atlas cluster(s) pause <clusterName> [--projectId projectId].
func PauseBuilder() *cobra.Command {
	opts := &PauseOpts{}
	cmd := &cobra.Command{
		Use:   "pause <clusterName>",
		Short: "Pause the specified running MongoDB cluster.",
		Long:  fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Cluster Manager"), "Atlas supports this command only for M10+ clusters."),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to pause.",
			"output":          pauseTmpl,
		},
		Example: `  # Pause the cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas clusters pause myCluster --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), pauseTmpl),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
