// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package globalaccesslists

import (
	"context"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.OutputOpts
	id    string
	store store.GlobalAPIKeyWhitelistDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const describeTemplate = `ID	CIDR BLOCK	CREATED AT
{{.ID}}	{{.CidrBlock}}	{{.Created}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.GlobalAPIKeyWhitelist(opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam globalAccessList(s) describe <ID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	opts.Template = describeTemplate
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified access list entry for your global API key.",
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the access list entry.",
		},
		Example: `  # Return the JSON-formatted details for the access list entry with the ID 5f5bad7a57aef32b04ed0210 from the access list for the global API key:
  mongocli iam globalAccessLists describe 5f5bad7a57aef32b04ed0210 --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
