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

package ldap

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type StatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	store store.LDAPConfigurationDescriber
}

func (opts *StatusOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var verifyStatusTemplate = `REQUEST ID	PROJECT ID	STATUS
{{.RequestID}}	{{.GroupID}}	{{.Status}}
`

func (opts *StatusOpts) Run() error {
	r, err := opts.store.GetStatusLDAPConfiguration(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

<<<<<<< HEAD
// mongocli atlas ldap verify status <ID> [--projectId projectId]
=======
// mongocli atlas ldap status <ID> [--projectId projectId]
>>>>>>> origin/master
func StatusBuilder() *cobra.Command {
	opts := &StatusOpts{}
	cmd := &cobra.Command{
		Use:   "status <ID>",
		Args:  cobra.ExactValidArgs(1),
		Short: verify,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), verifyStatusTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
