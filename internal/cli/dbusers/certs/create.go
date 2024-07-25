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

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store             store.DBUserCertificateCreator
	username          string
	monthsUntilExpiry int
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateDBUserCertificate(opts.ConfigProjectID(), opts.username, opts.monthsUntilExpiry)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

var createTemplate = "{{.}}\n"

// atlas dbuser(s) certs create --username <username> [--monthsUntilExpiration number] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Atlas-managed X.509 certificate for the specified database user.",
		Long:  `The user you specify must authenticate using X.509 certificates. You can't use this command to create certificates if you are managing your own Certificate Authority (CA) in self-managed X.509 mode.`,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
		Example: `  # Create an Atlas-managed X.509 certificate that expires in 5 months for a MongoDB user named dbuser for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers certs create --username dbuser --monthsUntilExpiration 5 --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	const defaultExpiration = 3

	cmd.Flags().IntVar(&opts.monthsUntilExpiry, flag.MonthsUntilExpiration, defaultExpiration, usage.MonthsUntilExpiration)
	cmd.Flags().StringVar(&opts.username, flag.Username, "", usage.DatabaseUser)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Username)

	return cmd
}
