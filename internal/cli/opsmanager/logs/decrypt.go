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
	"fmt"
	"io"
	"os"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/decryption"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type DecryptOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.DownloaderOpts
	inFileName                    string
	kmipServerCAFileName          string
	kmipClientCertificateFileName string
	localKeyFileName              string
}

// shouldPrintResultsToStdout returns true when the results should be printed to Stdout (--out|-o flag is not set).
func (opts *DecryptOpts) shouldPrintResultsToStdout() bool {
	return opts.Out == ""
}

func (opts *DecryptOpts) Run() error {
	var outWriter io.WriteCloser = os.Stdout
	if !opts.shouldPrintResultsToStdout() {
		var err error
		outWriter, err = opts.NewWriteCloser()
		if err != nil {
			return err
		}
		defer outWriter.Close()
	}

	inReader, err := opts.Fs.Open(opts.inFileName)
	if err != nil {
		return err
	}
	defer inReader.Close()

	keyProviderOpts := &decryption.KeyProviderOpts{
		Local: decryption.KeyProviderLocalOpts{
			KeyFileName: opts.localKeyFileName,
		},
		KMIP: decryption.KeyProviderKMIPOpts{
			ServerCAFileName:          opts.kmipServerCAFileName,
			ClientCertificateFileName: opts.kmipClientCertificateFileName,
		},
	}

	if err := decryption.Decrypt(inReader, outWriter, keyProviderOpts); err != nil && !opts.shouldPrintResultsToStdout() {
		_ = opts.OnError(outWriter)
		return err
	}

	if !opts.shouldPrintResultsToStdout() {
		fmt.Printf("Decrypt of %s to %s completed.\n", opts.inFileName, opts.Out)
	}

	return nil
}

// mongocli om logs decrypt --localKey <localKeyFile> --kmipServerCAFile <caFile> –-kmipClientCertificateFile <certFile> --file <encryptedLogFile> --out <outputLogFile>.
func DecryptBuilder() *cobra.Command {
	opts := &DecryptOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts an audit log file with the provided local key file or with a server that supports KMIP.",
		Example: `
	For audit logs in BSON format:
  $ mongocli ops-manager logs decrypt --localKeyFile /path/to/keyFile --file /path/to/logFile.bson --out /path/to/file.json
	For audit logs in JSON format:
  $ mongocli ops-manager logs decrypt --localKeyFile /path/to/keyFile --file /path/to/logFile.json --out /path/to/file.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.inFileName, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.Out, flag.Out, flag.OutputShort, "", usage.OutputLogFile)

	cmd.Flags().StringVarP(&opts.localKeyFileName, flag.LocalKeyFile, "", "", usage.LocalKeyFile)
	cmd.Flags().StringVarP(&opts.kmipServerCAFileName, flag.KMIPServerCAFile, "", "", usage.KMIPServerCAFile)
	cmd.Flags().StringVarP(&opts.kmipClientCertificateFileName, flag.KMIPClientCertificateFile, "", "", usage.KMIPClientCertificateFile)

	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagFilename(flag.Out)
	_ = cmd.MarkFlagFilename(flag.LocalKeyFile)
	_ = cmd.MarkFlagFilename(flag.KMIPServerCAFile)
	_ = cmd.MarkFlagFilename(flag.KMIPClientCertificateFile)

	return cmd
}
