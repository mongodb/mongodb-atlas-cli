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
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const describeTemplate = `ID	REPLICA SET NAME	SHARD NAME	VERSION
{{.Id}}	{{.ReplicaSetName}}	{{.ShardName}}	{{.Version}}
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
	processID := opts.host + ":" + strconv.Itoa(opts.port)
	listParams := newProcessParams(opts.ConfigProjectID(), processID)
	r, err := opts.store.Process(listParams)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func newProcessParams(projectID string, processID string) *atlasv2.GetAtlasProcessApiParams {
	return &atlasv2.GetAtlasProcessApiParams{
		GroupId:   projectID,
		ProcessId: processID,
	}
}

// DescribeBuilder atlas process(es) describe <hostname:port> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <hostname:port>",
		Short: "Return the details for the specified MongoDB process for your project.",
		Example: `  # Return the JSON-formatted details for the MongoDB process with hostname and port atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017
  atlas process describe atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 --output json`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the Atlas MongoDB process.",
			"output":            describeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
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
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
