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

package organizations

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const describeTemplate = `ID	NAME
{{.ID}}	{{.Name}}
`

type DescribeOpts struct {
	cli.OutputOpts
	id    string
	store store.OrganizationDescriber
}

func (opts *DescribeOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Organization(opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam organizations(s) describe <ID>
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	opts.Template = describeTemplate
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Short:   "Describe an organizations.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
