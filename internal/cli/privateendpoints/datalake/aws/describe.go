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

package aws

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var describeTemplate = `ID	ENDPOINT PROVIDER	TYPE	COMMENT
{{.EndpointId}}	{{.Provider}}	{{.Type}}	{{.Comment}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	privateEndpointID string
	store             store.DataLakePrivateEndpointDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DataLakePrivateEndpoint(opts.ConfigProjectID(), opts.privateEndpointID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas privateEndpoint(s)|privateendpoint(s) dataLakes aws describe|get <privateEndpointId> [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <privateEndpointId>",
		Aliases: []string{"get"},
		Short:   "Return a specific Data Lake private endpoint for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{
			"privateEndpointIdDesc": "Unique 22-character alphanumeric string that identifies the private endpoint.",
			"output":                describeTemplate,
		},
		Example: `  # This example uses the profile named "myprofile" for accessing Atlas.
  atlas privateEndpoint dataLake aws describe vpce-abcdefg0123456789 -P myprofile`,
		Args:       require.ExactArgs(1),
		Deprecated: "Please use 'atlas datafederation privateendpoints describe'",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.privateEndpointID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
