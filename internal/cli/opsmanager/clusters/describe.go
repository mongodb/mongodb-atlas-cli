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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.CloudManagerClustersDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	NAME	TYPE	REPLICASET NAME
{{.ID}}	{{.ClusterName}}	{{.TypeName}}	{{.ReplicaSetName}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.cluster()
	if err != nil {
		return err
	}

	return opts.Print(r)
}
func (opts *DescribeOpts) validateArg() error {
	if opts.ConfigOutput() != "" {
		return nil
	}
	if err := validate.ObjectID(opts.name); err != nil {
		return fmt.Errorf("please provide a valid cluster ID or provide an output format to use names, %w", err)
	}
	return nil
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

// mongocli cloud-manager cluster(s) describe <name> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <id|name>",
		Short: "Describe a cluster.",
		Long: `When describing cluster with no output format please provide the cluster ID.
When using an output format please provide the cluster name.`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"id|nameDesc": "Name or ID of the cluster.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.validateArg,
				opts.initStore(cmd.Context()),
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
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
