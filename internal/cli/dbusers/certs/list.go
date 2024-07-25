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

package certs

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	store    store.DBUserCertificateLister
	username string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.DBUserCertificates(opts.ConfigProjectID(), opts.username, nil)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

var listTemplate = `ID SUBJECT CREATED AT{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Subject}}	{{.CreatedAt}}{{end}}
`

// atlas dbuser(s) certs list|ls <username> [--projectId projectId].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list <username>",
		Aliases: []string{"ls"},
		Short:   "Return all Atlas-managed, unexpired X.509 certificates for the specified database user.",
		Long: `You can't use this command to return certificates if you are managing your own Certificate Authority (CA) in self-managed X.509 mode.
		
The user you specify must authenticate using X.509 certificates.`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"usernameDesc": "Username of the database user for whom you want to list Atlas-managed certificates.",
		},
		Example: `  # Return a JSON-formatted list of all Atlas-managed X.509 certificates for a MongoDB user named dbuser for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers certs list dbuser --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
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
