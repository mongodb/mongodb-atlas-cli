// Copyright 2021 MongoDB Inc
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

package connectionstring

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name   string
	store  store.ClusterDescriber
	csType string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplateStandard = `STANDARD CONNECTION STRING
{{.StandardSrv}}
`

var describeTemplatePrivate = `PRIVATE CONNECTION STRING
{{.PrivateSrv}}
`
var describeTemplateShardOptimized = `SHARD OPTIMIZED CONNECTION STRING
{{range .PrivateEndpoint }}{{.SrvShardOptimizedConnectionString}}{{end}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.AtlasCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}
	return opts.Print(r.ConnectionStrings)
}

// atlas cluster(s) connectionString describe <clusterName> --type standard|private --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe <clusterName>",
		Aliases: []string{"get"},
		Short:   "Return the SRV connection string for the cluster you specify.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster for which you want to retrieve connection strings.",
		},
		Example: `  # Return the JSON-formatted connection strings for the cluster named myCluster:
  atlas clusters connectionStrings describe myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplateStandard),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]

			switch strings.ToLower(opts.csType) {
			case "private":
				opts.Template = describeTemplatePrivate
			case "privateendpoint", "privateendpoints":
				opts.Template = describeTemplateShardOptimized
			default:
				opts.Template = describeTemplateStandard
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.csType, flag.TypeFlag, "standard", usage.ConnectionStringType)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
