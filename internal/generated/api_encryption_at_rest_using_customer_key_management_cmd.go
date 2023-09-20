// Copyright 2023 MongoDB Inc
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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type getEncryptionAtRestOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *getEncryptionAtRestOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getEncryptionAtRestOpts) Run(ctx context.Context) error {
	params := &admin.GetEncryptionAtRestApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRestWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getEncryptionAtRestBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getEncryptionAtRestOpts{}
	cmd := &cobra.Command{
		Use: "getEncryptionAtRest",
		Short: "Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)


	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}
type updateEncryptionAtRestOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	
}

func (opts *updateEncryptionAtRestOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateEncryptionAtRestOpts) Run(ctx context.Context) error {
	params := &admin.UpdateEncryptionAtRestApiParams{
		GroupId: opts.groupId,
		
	}
	resp, _, err := opts.client.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRestWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateEncryptionAtRestBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateEncryptionAtRestOpts{}
	cmd := &cobra.Command{
		Use: "updateEncryptionAtRest",
		Short: "Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	

	cmd.Flags().AWSKMSVar(&opts.awsKms, "awsKms", , ``)

	cmd.Flags().AzureKeyVaultVar(&opts.azureKeyVault, "azureKeyVault", , ``)

	cmd.Flags().GoogleCloudKMSVar(&opts.googleCloudKms, "googleCloudKms", , ``)


	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func encryptionAtRestUsingCustomerKeyManagementBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "encryptionAtRestUsingCustomerKeyManagement",
		Short:   `Returns and edits the Encryption at Rest using Customer Key Management configuration. MongoDB Cloud encrypts all storage whether or not you use your own key management.`,
	}
	cmd.AddCommand(
		getEncryptionAtRestBuilder(),
		updateEncryptionAtRestBuilder(),
	)
	return cmd
}

