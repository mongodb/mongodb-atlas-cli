// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodes

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	clusterName string
	filename    string
	fs          afero.Fs
	store       store.SearchNodesCreator
}

const (
	stateIdle = "IDLE"
)

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Search nodes are being created.\n"
var createWatchTemplate = "Created new search nodes.\n"

func (opts *CreateOpts) Run() error {
	spec, err := loadAPISearchDeploymentSpec(opts.fs, opts.filename)
	if err != nil {
		return err
	}

	r, err := opts.store.CreateSearchNodes(opts.ConfigProjectID(), opts.clusterName, spec)
	if err != nil {
		return err
	}

	if opts.EnableWatch {
		watchResult, err := opts.Watch(opts.watcher)
		if err != nil {
			return err
		}
		opts.Template = createWatchTemplate
		r = watchResult.(*admin.ApiSearchDeploymentResponse)
	}

	return opts.Print(r)
}

func (opts *CreateOpts) watcher() (any, bool, error) {
	res, err := opts.store.SearchNodes(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return nil, false, err
	}
	return res, *res.StateName == stateIdle, nil
}

// atlas clusters search nodes create [--clusterName clusterName] [--file filePath].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	opts.fs = afero.NewOsFs()

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a search node for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Organization Owner or Project Owner"),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: `  # Create a search node for the cluster named myCluster using a JSON node spec configuration file named spec.json:
  atlas clusters search nodes create --clusterName myCluster --file spec.json --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	// Command specific flags
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	_ = cmd.MarkFlagRequired(flag.ClusterName)

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.SearchNodesFilename)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagRequired(flag.File)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().UintVar(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)

	// Global flags
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
