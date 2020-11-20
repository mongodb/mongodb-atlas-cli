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
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.CloudManagerClustersDescriber
}

func (opts *DescribeOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var describeTemplate = `ID	NAME	TYPE
{{.ID}}	{{.ClusterName}}	{{.TypeName}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.cluster()
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) cluster() (interface{}, error) {
	if opts.ConfigOutput() == "" {
		return opts.store.OpsManagerCluster(opts.ConfigProjectID(), opts.name)
	}
	c, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return nil, err
	}
	r := convert.FromAutomationConfig(c)
	for _, rs := range r {
		if rs.Name == opts.name {
			return rs, nil
		}
	}
	return nil, fmt.Errorf("replica set %s not found", opts.name)
}

// mongocli cloud-manager cluster(s) describe <name> --projectId projectId
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <name>",
		Short: DescribeCluster,
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
