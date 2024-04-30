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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/decryption"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
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

var decryptTemplate = "Decrypt of %s to %s completed.\n"

func (opts *DecryptOpts) newDecryption() *decryption.Decryption {
	return decryption.NewDecryption(
		decryption.WithAWSOpts(opts.awsOpts.awsAccessKey, opts.awsOpts.awsSecretAccessKey, opts.awsOpts.awsSessionToken),
		decryption.WithGCPOpts(opts.gcpOpts.gcpServiceAccountKey),
		decryption.WithAzureOpts(opts.azureOpts.azureTenantID, opts.azureOpts.azureClientID, opts.azureOpts.azureSecret),
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
		fmt.Printf(decryptTemplate, opts.inFileName, opts.Out)
	}

	return nil
}

// atlas logs decrypt --localKey <localKeyFile> --kmipServerCAFile <caFile> –-kmipClientCertificateFile <certFile> --file <encryptedLogFile> --out <outputLogFile>.
func DecryptBuilder() *cobra.Command {
	opts := &DecryptOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:    "decrypt",
		Hidden: true,
		Short:  "Decrypts an audit log file with the provided AWS, GCP or Azure key management services.",
		Annotations: map[string]string{
			"output": decryptTemplate,
		},
		Example: `  # Decrypt using AWS credentials:
  atlas logs decrypt --file /path/to/logFile.json --awsAccessKey <accessKey> --awsSecretAccessKey <secretKey> --awsSessionToken <sessionToken>
  # Decrypt using GCP credentials:
  atlas logs decrypt --file /path/to/logFile.json --gcpServiceAccountKey <serviceAccountKey>
  # Decrypt using Azure credentials:
  atlas logs decrypt --file /path/to/logFile.json --azureClientId <clientId> --azureTenantId <tenantId> --azureSecret <secret>
`,
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return opts.PreRunE(opts.initDefaultOut)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
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
