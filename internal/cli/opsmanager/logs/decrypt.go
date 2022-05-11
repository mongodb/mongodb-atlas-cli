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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/decryption"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
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
	kmipClientCertificatePassword string
	kmipUsername                  string
	kmipPassword                  string
	localKeyFileName              string
}

func (opts *DecryptOpts) newDecryption() *decryption.Decryption {
	return decryption.NewDecryption(
		decryption.WithLocalOpts(opts.localKeyFileName),
		decryption.WithKMIPOpts(&decryption.KeyProviderKMIPOpts{
			ServerCAFileName:          opts.kmipServerCAFileName,
			ClientCertificateFileName: opts.kmipClientCertificateFileName,
			ClientCertificatePassword: opts.kmipClientCertificatePassword,
			Username:                  opts.kmipUsername,
			Password:                  opts.kmipPassword,
		}),
	)
}

func (opts *DecryptOpts) initDefaultOut() error {
	if opts.Out == "" {
		opts.Out = cli.StdOutMode // sets to "-"
	}
	return nil
}

func (opts *DecryptOpts) Run() error {
	outWriter, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}
	defer outWriter.Close()

	inReader, err := opts.Fs.Open(opts.inFileName)
	if err != nil {
		return err
	}
	defer inReader.Close()

	d := opts.newDecryption()
	if err := d.Decrypt(inReader, outWriter); err != nil && !opts.ShouldDownloadToStdout() {
		_ = opts.OnError(outWriter)
		return err
	}

	if !opts.ShouldDownloadToStdout() {
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initDefaultOut)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.inFileName, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.Out, flag.Out, flag.OutputShort, "", usage.OutputLogFile)

	cmd.Flags().StringVarP(&opts.localKeyFileName, flag.LocalKeyFile, "", "", usage.LocalKeyFile)
	cmd.Flags().StringVarP(&opts.kmipServerCAFileName, flag.KMIPServerCAFile, "", "", usage.KMIPServerCAFile)
	cmd.Flags().StringVarP(&opts.kmipClientCertificateFileName, flag.KMIPClientCertificateFile, "", "", usage.KMIPClientCertificateFile)
	cmd.Flags().StringVar(&opts.kmipClientCertificatePassword, flag.KMIPClientCertificatePassword, "", usage.KMIPClientCertificatePassword)
	cmd.Flags().StringVarP(&opts.kmipUsername, flag.KMIPUsername, "", "", usage.KMIPUsername)
	cmd.Flags().StringVarP(&opts.kmipPassword, flag.KMIPPassword, "", "", usage.KMIPPassword)

	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagFilename(flag.Out)
	_ = cmd.MarkFlagFilename(flag.LocalKeyFile)
	_ = cmd.MarkFlagFilename(flag.KMIPServerCAFile)
	_ = cmd.MarkFlagFilename(flag.KMIPClientCertificateFile)

	return cmd
}
