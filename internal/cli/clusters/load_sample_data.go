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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=load_sample_data_mock_test.go -package=clusters . SampleDataAdder

type SampleDataAdder interface {
	AddSampleData(string, string) (*atlasv2.SampleDatasetStatus, error)
}

type LoadSampleDataOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name  string
	store SampleDataAdder
}

func (opts *LoadSampleDataOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var addTmpl = "Sample data load job {{.Id}} created.\n"

func (opts *LoadSampleDataOpts) Run() error {
	r, err := opts.store.AddSampleData(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas cluster loadSampleData <clusterName> --projectId projectId -o json.
func LoadSampleDataBuilder() *cobra.Command {
	opts := &LoadSampleDataOpts{}
	cmd := &cobra.Command{
		Use:        "loadSampleData <clusterName>",
		Short:      "Load sample data into the specified cluster for your project.",
		Long:       fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Deprecated: "use 'atlas clusters sampleData load' instead",
		Args:       require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Label that identifies the cluster that you want to load sample data into.",
			"output":          addTmpl,
		},
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
