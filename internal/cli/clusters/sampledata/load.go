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

package sampledata

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=load_mock_test.go -package=sampledata . Adder

type Adder interface {
	AddSampleData(string, string) (*atlasv2.SampleDatasetStatus, error)
}

type LoadOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name  string
	store Adder
}

func (opts *LoadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var addTmpl = "Sample Data Job {{.Id}} created.\n"

func (opts *LoadOpts) Run() error {
	r, err := opts.store.AddSampleData(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas cluster sampleData load <clusterName> --projectId projectId -o json.
func LoadBuilder() *cobra.Command {
	opts := &LoadOpts{}
	cmd := &cobra.Command{
		Use:   "load <clusterName>",
		Short: "Load sample data into the specified cluster for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster for which you want to load sample data.",
			"output":          addTmpl,
		},
		Example: `  # Load sample data into the cluster named myCluster:
  atlas clusters sampleData load myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), addTmpl),
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
