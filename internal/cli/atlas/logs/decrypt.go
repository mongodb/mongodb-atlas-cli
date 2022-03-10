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
	inFileName  string
	outFileName string
	awsOpts     DecryptAWSOpts
	gcpOpts     DecryptGCPOpts
	azureOpts   DecryptAzureOpts
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

func (opts *DecryptOpts) initFiles() func() error {
	// Validate the provided files can be open
	return func() error {
		return nil
	}
}

func (opts *DecryptOpts) initFlags() func() error {
	// Validate all needed flags are set or try to retrieve right credentials.
	return func() error {
		return nil
	}
}

func (opts *DecryptOpts) Run() error {
	// Run the command. For now printing the file path.
	return opts.Print(opts.inFileName)
}

// atlas logs decrypt --file <encryptedLogFile> --out <outputLogFile>.
func DecryptBuilder() *cobra.Command {
	opts := &DecryptOpts{}
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypts a log file with the provided local key file or KMIP files.",
		Example: `
  $ atlascli decrypt --file logPath --out resultPath --awsAccessKey accessKey --awsSecretKey secretKey --awsSessionToken sessionToken`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initFiles(), opts.initFlags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.inFileName, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.outFileName, flag.Out, "", "", usage.OutputLogFile)

	cmd.Flags().StringVarP(&opts.awsOpts.awsAccessKey, flag.AWSAccessKey, "", "", usage.DecryptAWSAccessKey)
	cmd.Flags().StringVarP(&opts.awsOpts.awsSecretAccessKey, flag.AWSSecretKey, "", "", usage.DecryptAWSSecretKey)
	cmd.Flags().StringVarP(&opts.awsOpts.awsSessionToken, flag.AWSSessionToken, "", "", usage.AWSSessionToken)

	cmd.Flags().StringVarP(&opts.azureOpts.azureClientID, flag.AzureClientID, "", "", usage.AzureClientID)
	cmd.Flags().StringVarP(&opts.azureOpts.azureTenantID, flag.AzureTenantID, "", "", usage.AzureTenantID)
	cmd.Flags().StringVarP(&opts.azureOpts.azureSecret, flag.AzureSecret, "", "", usage.AzureSecret)

	cmd.Flags().StringVarP(&opts.gcpOpts.gcpServiceAccountKey, flag.GCPServiceAccountKey, "", "", usage.GCPServiceAccountKey)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.File)

	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}
