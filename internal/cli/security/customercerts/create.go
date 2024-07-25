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

package customercerts

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const createTemplate = "Certificate successfully created.\n"

type SaveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store   store.X509CertificateConfSaver
	casPath string
	fs      afero.Fs
}

func (opts *SaveOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *SaveOpts) Run() error {
	fileBytes, err := afero.ReadFile(opts.fs, opts.casPath)
	if err != nil {
		return err
	}

	caFileContents := string(fileBytes)

	r, err := opts.store.SaveX509Configuration(opts.ConfigProjectID(), caFileContents)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas security customercerts create --projectId projectId --casFile /path/to/certificates.pem.
func CreateBuilder() *cobra.Command {
	opts := &SaveOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Saves a customer-managed X.509 configuration for your project.",
		Long: `Saving a customer-managed X.509 configuration triggers a rolling restart.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
		Example: `  # Save the file named ca.pem stored in the files directory to the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas security customerCerts create --casFile files/ca.pem --projectId 5e2211c17a3e5a48f5497de3 --output json`,
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

	cmd.Flags().StringVar(&opts.casPath, flag.CASFilePath, "", usage.CASFilePath)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.CASFilePath)

	return cmd
}
