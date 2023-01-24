// Copyright 2022 MongoDB Inc
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

package processes

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
)

const describeTemplate = `ID	REPLICA SET NAME	SHARD NAME	VERSION
{{.ID}}	{{.ReplicaSetName}}	{{.ShardName}}	{{.Version}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	host  string
	port  int
	store store.ProcessDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Process(opts.ConfigProjectID(), opts.host, opts.port)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// DescribeBuilder atlas process(es) describe <hostname:port> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <hostname:port>",
		Short: "Return the details for the specified MongoDB process for your project.",
		Example: fmt.Sprintf(`  # Return the JSON-formatted details for the MongoDB process with hostname and port atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017
  %s process describe atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 --output json`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the Atlas MongoDB process.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = cli.GetHostnameAndPort(args[0])
			if err != nil {
				return err
			}
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
