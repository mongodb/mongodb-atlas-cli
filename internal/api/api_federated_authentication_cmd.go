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

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createRoleMappingOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	orgId                string

	filename string
	fs       afero.Fs
}

func (opts *createRoleMappingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createRoleMappingOpts) readData() (*admin.AuthFederationRoleMapping, error) {
	var out *admin.AuthFederationRoleMapping

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *createRoleMappingOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.CreateRoleMappingApiParams{
		FederationSettingsId: opts.federationSettingsId,
		OrgId:                opts.orgId,

		AuthFederationRoleMapping: data,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.CreateRoleMappingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func createRoleMappingBuilder() *cobra.Command {
	opts := createRoleMappingOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createRoleMapping",
		Short: "Add One Role Mapping to One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteFederationAppOpts struct {
	client               *admin.APIClient
	federationSettingsId string
}

func (opts *deleteFederationAppOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteFederationAppOpts) run(ctx context.Context, _ io.Writer) error {

	params := &admin.DeleteFederationAppApiParams{
		FederationSettingsId: opts.federationSettingsId,
	}

	_, err := opts.client.FederatedAuthenticationApi.DeleteFederationAppWithParams(ctx, params).Execute()
	return err
}

func deleteFederationAppBuilder() *cobra.Command {
	opts := deleteFederationAppOpts{}
	cmd := &cobra.Command{
		Use:   "deleteFederationApp",
		Short: "Delete the federation settings instance.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	return cmd
}

type deleteRoleMappingOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	id                   string
	orgId                string
}

func (opts *deleteRoleMappingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteRoleMappingOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.DeleteRoleMappingApiParams{
		FederationSettingsId: opts.federationSettingsId,
		Id:                   opts.id,
		OrgId:                opts.orgId,
	}

	_, err := opts.client.FederatedAuthenticationApi.DeleteRoleMappingWithParams(ctx, params).Execute()
	return err
}

func deleteRoleMappingBuilder() *cobra.Command {
	opts := deleteRoleMappingOpts{}
	cmd := &cobra.Command{
		Use:   "deleteRoleMapping",
		Short: "Remove One Role Mapping from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.id, "id", "", `Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getConnectedOrgConfigOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	orgId                string
}

func (opts *getConnectedOrgConfigOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getConnectedOrgConfigOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsId,
		OrgId:                opts.orgId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.GetConnectedOrgConfigWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getConnectedOrgConfigBuilder() *cobra.Command {
	opts := getConnectedOrgConfigOpts{}
	cmd := &cobra.Command{
		Use:   "getConnectedOrgConfig",
		Short: "Return One Org Config Connected to One Federation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getFederationSettingsOpts struct {
	client *admin.APIClient
	orgId  string
}

func (opts *getFederationSettingsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getFederationSettingsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.GetFederationSettingsApiParams{
		OrgId: opts.orgId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.GetFederationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getFederationSettingsBuilder() *cobra.Command {
	opts := getFederationSettingsOpts{}
	cmd := &cobra.Command{
		Use:   "getFederationSettings",
		Short: "Return Federation Settings for One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getIdentityProviderOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	identityProviderId   string
}

func (opts *getIdentityProviderOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getIdentityProviderOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetIdentityProviderApiParams{
		FederationSettingsId: opts.federationSettingsId,
		IdentityProviderId:   opts.identityProviderId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.GetIdentityProviderWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getIdentityProviderBuilder() *cobra.Command {
	opts := getIdentityProviderOpts{}
	cmd := &cobra.Command{
		Use:   "getIdentityProvider",
		Short: "Return one identity provider from the specified federation.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.identityProviderId, "identityProviderId", "", `Unique 20-hexadecimal digit string that identifies the identity provider.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("identityProviderId")
	return cmd
}

type getIdentityProviderMetadataOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	identityProviderId   string
}

func (opts *getIdentityProviderMetadataOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getIdentityProviderMetadataOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetIdentityProviderMetadataApiParams{
		FederationSettingsId: opts.federationSettingsId,
		IdentityProviderId:   opts.identityProviderId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.GetIdentityProviderMetadataWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getIdentityProviderMetadataBuilder() *cobra.Command {
	opts := getIdentityProviderMetadataOpts{}
	cmd := &cobra.Command{
		Use:   "getIdentityProviderMetadata",
		Short: "Return the metadata of one identity provider in the specified federation.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.identityProviderId, "identityProviderId", "", `Unique 20-hexadecimal digit string that identifies the identity provider.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("identityProviderId")
	return cmd
}

type getRoleMappingOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	id                   string
	orgId                string
}

func (opts *getRoleMappingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getRoleMappingOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.GetRoleMappingApiParams{
		FederationSettingsId: opts.federationSettingsId,
		Id:                   opts.id,
		OrgId:                opts.orgId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.GetRoleMappingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getRoleMappingBuilder() *cobra.Command {
	opts := getRoleMappingOpts{}
	cmd := &cobra.Command{
		Use:   "getRoleMapping",
		Short: "Return One Role Mapping from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.id, "id", "", `Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listConnectedOrgConfigsOpts struct {
	client               *admin.APIClient
	federationSettingsId string
}

func (opts *listConnectedOrgConfigsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listConnectedOrgConfigsOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListConnectedOrgConfigsApiParams{
		FederationSettingsId: opts.federationSettingsId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.ListConnectedOrgConfigsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listConnectedOrgConfigsBuilder() *cobra.Command {
	opts := listConnectedOrgConfigsOpts{}
	cmd := &cobra.Command{
		Use:   "listConnectedOrgConfigs",
		Short: "Return All Connected Org Configs from the Federation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	return cmd
}

type listIdentityProvidersOpts struct {
	client               *admin.APIClient
	federationSettingsId string
}

func (opts *listIdentityProvidersOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listIdentityProvidersOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListIdentityProvidersApiParams{
		FederationSettingsId: opts.federationSettingsId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.ListIdentityProvidersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listIdentityProvidersBuilder() *cobra.Command {
	opts := listIdentityProvidersOpts{}
	cmd := &cobra.Command{
		Use:   "listIdentityProviders",
		Short: "Return all identity providers from the specified federation.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	return cmd
}

type listRoleMappingsOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	orgId                string
}

func (opts *listRoleMappingsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listRoleMappingsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.ListRoleMappingsApiParams{
		FederationSettingsId: opts.federationSettingsId,
		OrgId:                opts.orgId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.ListRoleMappingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listRoleMappingsBuilder() *cobra.Command {
	opts := listRoleMappingsOpts{}
	cmd := &cobra.Command{
		Use:   "listRoleMappings",
		Short: "Return All Role Mappings from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type removeConnectedOrgConfigOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	orgId                string
}

func (opts *removeConnectedOrgConfigOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *removeConnectedOrgConfigOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.RemoveConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsId,
		OrgId:                opts.orgId,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.RemoveConnectedOrgConfigWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func removeConnectedOrgConfigBuilder() *cobra.Command {
	opts := removeConnectedOrgConfigOpts{}
	cmd := &cobra.Command{
		Use:   "removeConnectedOrgConfig",
		Short: "Remove One Org Config Connected to One Federation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.`)

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateConnectedOrgConfigOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	orgId                string

	filename string
	fs       afero.Fs
}

func (opts *updateConnectedOrgConfigOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateConnectedOrgConfigOpts) readData() (*admin.ConnectedOrgConfig, error) {
	var out *admin.ConnectedOrgConfig

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *updateConnectedOrgConfigOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.UpdateConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsId,
		OrgId:                opts.orgId,

		ConnectedOrgConfig: data,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.UpdateConnectedOrgConfigWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func updateConnectedOrgConfigBuilder() *cobra.Command {
	opts := updateConnectedOrgConfigOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateConnectedOrgConfig",
		Short: "Update One Org Config Connected to One Federation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateIdentityProviderOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	identityProviderId   string

	filename string
	fs       afero.Fs
}

func (opts *updateIdentityProviderOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateIdentityProviderOpts) readData() (*admin.SamlIdentityProviderUpdate, error) {
	var out *admin.SamlIdentityProviderUpdate

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *updateIdentityProviderOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.UpdateIdentityProviderApiParams{
		FederationSettingsId: opts.federationSettingsId,
		IdentityProviderId:   opts.identityProviderId,

		SamlIdentityProviderUpdate: data,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.UpdateIdentityProviderWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func updateIdentityProviderBuilder() *cobra.Command {
	opts := updateIdentityProviderOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateIdentityProvider",
		Short: "Update the identity provider.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.identityProviderId, "identityProviderId", "", `Unique 20-hexadecimal digit string that identifies the identity provider.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("identityProviderId")
	return cmd
}

type updateRoleMappingOpts struct {
	client               *admin.APIClient
	federationSettingsId string
	id                   string
	orgId                string

	filename string
	fs       afero.Fs
}

func (opts *updateRoleMappingOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateRoleMappingOpts) readData() (*admin.AuthFederationRoleMapping, error) {
	var out *admin.AuthFederationRoleMapping

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *updateRoleMappingOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.UpdateRoleMappingApiParams{
		FederationSettingsId: opts.federationSettingsId,
		Id:                   opts.id,
		OrgId:                opts.orgId,

		AuthFederationRoleMapping: data,
	}

	resp, _, err := opts.client.FederatedAuthenticationApi.UpdateRoleMappingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func updateRoleMappingBuilder() *cobra.Command {
	opts := updateRoleMappingOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateRoleMapping",
		Short: "Update One Role Mapping in One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.federationSettingsId, "federationSettingsId", "", `Unique 24-hexadecimal digit string that identifies your federation.`)
	cmd.Flags().StringVar(&opts.id, "id", "", `Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.`)
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("federationSettingsId")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

func federatedAuthenticationBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "federatedAuthentication",
		Short: `Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.`,
	}
	cmd.AddCommand(
		createRoleMappingBuilder(),
		deleteFederationAppBuilder(),
		deleteRoleMappingBuilder(),
		getConnectedOrgConfigBuilder(),
		getFederationSettingsBuilder(),
		getIdentityProviderBuilder(),
		getIdentityProviderMetadataBuilder(),
		getRoleMappingBuilder(),
		listConnectedOrgConfigsBuilder(),
		listIdentityProvidersBuilder(),
		listRoleMappingsBuilder(),
		removeConnectedOrgConfigBuilder(),
		updateConnectedOrgConfigBuilder(),
		updateIdentityProviderBuilder(),
		updateRoleMappingBuilder(),
	)
	return cmd
}
