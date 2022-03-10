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

package logs

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DecryptOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	inFileName                    string
	outFileName                   string
	kmipServerCAFileName          string
	kmipClientCertificateFileName string
	localKeyFileName              string
}

func (opts *DecryptOpts) initStore() func() error {
	// Keeping for now, not sure if needed
	return func() error {
		return nil
	}
}

func (opts *DecryptOpts) initFiles() func() error {
	// Validate the provided files can be open
	return func() error {
		return nil
	}
}

func (opts *DecryptOpts) Run() error {
	// Run the command. For now printing the file path.
	return opts.Print(opts.inFileName)
}

// mongocli om logs decrypt --localKey <localKeyFile> --kmipServerCAFile <caFile> –-kmipClientCertificateFile <certFile> --file <encryptedLogFile> --out <outputLogFile>.
func DecryptBuilder() *cobra.Command {
	opts := &DecryptOpts{}
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts a log file with the provided local key file or KMIP files.",
		Example: `
  $ mongocli ops-manager decrypt --localKey filePath --file logPath --out resultPath`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initFiles(), opts.initStore())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.inFileName, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.outFileName, flag.Out, "", "", usage.OutputLogFile)

	cmd.Flags().StringVarP(&opts.localKeyFileName, flag.LocalKeyFile, "", "", usage.LocalKeyFile)
	cmd.Flags().StringVarP(&opts.kmipServerCAFileName, flag.KMIPServerCAFile, "", "", usage.KMIPServerCAFile)
	cmd.Flags().StringVarP(&opts.kmipClientCertificateFileName, flag.KMIPClientCertificateFile, "", "", usage.KMIPClientCertificateFile)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.File)

	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagFilename(flag.Out)
	_ = cmd.MarkFlagFilename(flag.LocalKeyFile)
	_ = cmd.MarkFlagFilename(flag.KMIPServerCAFile)
	_ = cmd.MarkFlagFilename(flag.KMIPClientCertificateFile)

	return cmd
}
