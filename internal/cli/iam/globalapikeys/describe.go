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

package globalapikeys

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.OutputOpts
	id    string
	store store.GlobalAPIKeyDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const describeTemplate = `ID	DESCRIPTION	PUBLIC KEY	PRIVATE KEY
{{.ID}}	{{.Desc}}	{{.PublicKey}}	{{.PrivateKey}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.GlobalAPIKey(opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam globalApiKey(s) describe <ID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	opts.Template = describeTemplate
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Short:   "Return the details for the specified global API key for your Ops Manager instance.",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the global API key.",
		},
		Example: `  # Return the JSON-formatted details for the global API key with the ID 5f5bad7a57aef32b04ed0210:
  mongocli iam globalApiKeys describe 5f5bad7a57aef32b04ed0210 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
