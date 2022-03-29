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
	inFileName string
	awsOpts    DecryptAWSOpts
	gcpOpts    DecryptGCPOpts
	azureOpts  DecryptAzureOpts
}

type DecryptAWSOpts struct {
	awsAccessKey       string
	awsSecretAccessKey string
	awsSessionToken    string
}

type DecryptGCPOpts struct {
	gcpServiceAccountKey string
}

type DecryptAzureOpts struct {
	azureClientID string
	azureTenantID string
	azureSecret   string
}

// stdOutMode returns true when the results should be printed to Stdout (--out|-o flag is not set).
func (opts *DecryptOpts) stdOutMode() bool {
	return opts.Out == ""
}

func (opts *DecryptOpts) Run() error {
	var outWriter io.WriteCloser = os.Stdout
	if !opts.stdOutMode() {
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

	keyProviderOpts := decryption.KeyProviderOpts{
		AWS: decryption.KeyProviderAWSOpts{
			AccessKey:       opts.awsOpts.awsAccessKey,
			SecretAccessKey: opts.awsOpts.awsSecretAccessKey,
			SessionToken:    opts.awsOpts.awsSessionToken,
		},
		GCP: decryption.KeyProviderGCPOpts{
			ServiceAccountKey: opts.gcpOpts.gcpServiceAccountKey,
		},
		Azure: decryption.KeyProviderAzureOpts{
			ClientID: opts.azureOpts.azureClientID,
			Secret:   opts.azureOpts.azureSecret,
			TenantID: opts.azureOpts.azureTenantID,
		},
	}

	if err := decryption.Decrypt(inReader, outWriter, keyProviderOpts); err != nil && !opts.stdOutMode() {
		_ = opts.OnError(outWriter)
		return err
	}

	if !opts.stdOutMode() {
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
		Short: "Decrypts an audit log file with the provided AWS, GCP or Azure key management services.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.inFileName, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.Out, flag.Out, flag.OutputShort, "", usage.OutputLogFile)

	cmd.Flags().StringVarP(&opts.awsOpts.awsAccessKey, flag.AWSAccessKey, "", "", usage.DecryptAWSAccessKey)
	cmd.Flags().StringVarP(&opts.awsOpts.awsSecretAccessKey, flag.AWSSecretKey, "", "", usage.DecryptAWSSecretKey)
	cmd.Flags().StringVarP(&opts.awsOpts.awsSessionToken, flag.AWSSessionToken, "", "", usage.AWSSessionToken)

	cmd.Flags().StringVarP(&opts.azureOpts.azureClientID, flag.AzureClientID, "", "", usage.AzureClientID)
	cmd.Flags().StringVarP(&opts.azureOpts.azureTenantID, flag.AzureTenantID, "", "", usage.AzureTenantID)
	cmd.Flags().StringVarP(&opts.azureOpts.azureSecret, flag.AzureSecret, "", "", usage.AzureSecret)

	cmd.Flags().StringVarP(&opts.gcpOpts.gcpServiceAccountKey, flag.GCPServiceAccountKey, "", "", usage.GCPServiceAccountKey)

	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}
