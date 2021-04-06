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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type LoadSampleDataOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.SampleDataAdder
}

func (opts *LoadSampleDataOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var addTmpl = "Loading sample dataset into {{.ClusterName}}.\n"

func (opts *LoadSampleDataOpts) Run() error {
	r, err := opts.store.AddSampleData(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas cluster loadSampleData <clusterName> --projectId projectId -o json
func AddBuilder() *cobra.Command {
	opts := &LoadSampleDataOpts{}
	cmd := &cobra.Command{
		Use:   "loadSampleData <clusterName>",
		Short: "Load sample data into a MongoDB cluster in Atlas.",
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), addTmpl),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagFilename(flag.File)

	return cmd
}
