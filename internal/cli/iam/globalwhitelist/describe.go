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
package globalwhitelist

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	id    string
	store store.GlobalAPIKeyWhitelistDescriber
}

func (opts *DescribeOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

const describeTemplate = `ID	CIDR BLOCK	CREATED AT
{{.ID}}	{{.CidrBlock}}	{{.Created}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.GlobalAPIKeyWhitelist(opts.id)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), describeTemplate, r)
}

// mongocli iam globalWhitelist(s) describe <ID>
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show"},
		Args:    cobra.ExactArgs(1),
		Short:   describeEntry,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	return cmd
}
