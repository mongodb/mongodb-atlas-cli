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

package decryption

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/decryption"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type KeyProviderListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	file string
	fs   afero.Fs
}

var listTmpl = `{{ range . }}{{ .Provider }}:{{if .Filename}} Filename = {{ .Filename }}{{end}}{{if .UID}} Unique Key ID = "{{ .UID }}"{{end}}{{if .KMIPServerName}} KMIP Server Name = "{{.KMIPServerName}}"{{end}}{{if .KMIPPort}} KMIP Port = "{{ .KMIPPort }}"{{end}}{{if .KeyWrapMethod}} Key Wrap Method = "{{ .KeyWrapMethod }}"{{end}}
{{ end }}`

func (opts *KeyProviderListOpts) Run() error {
	f, err := opts.fs.Open(opts.file)
	if err != nil {
		return err
	}
	defer f.Close()

	logs, err := decryption.ListKeyProviders(f)
	if err != nil {
		return err
	}

	return opts.Print(logs)
}

// mongocli om logs listKeyProvider --file <encryptedLogFile>.
func KeyProvidersListBuilder() *cobra.Command {
	opts := &KeyProviderListOpts{fs: afero.NewOsFs()}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists all key provider configurations found in the encrypted audit log file.",
		Example: `
  $ mongocli ops-manager listKeyProvider --file log.gz`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), listTmpl),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.file, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)
	return cmd
}
