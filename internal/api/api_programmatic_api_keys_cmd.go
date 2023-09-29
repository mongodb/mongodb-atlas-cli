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

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type addProjectApiKeyOpts struct {
	client    *admin.APIClient
	groupId   string
	apiUserId string

	filename string
	fs       afero.Fs
}

func (opts *addProjectApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *addProjectApiKeyOpts) readData() (*[]admin.UserAccessRoleAssignment, error) {
	var out *[]admin.UserAccessRoleAssignment

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

func (opts *addProjectApiKeyOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.AddProjectApiKeyApiParams{
		GroupId:   opts.groupId,
		ApiUserId: opts.apiUserId,

		UserAccessRoleAssignment: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.AddProjectApiKeyWithParams(ctx, params).Execute()
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

func addProjectApiKeyBuilder() *cobra.Command {
	opts := addProjectApiKeyOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "addProjectApiKey",
		Short: "Assign One Organization API Key to One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type createApiKeyOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *createApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createApiKeyOpts) readData() (*admin.CreateAtlasOrganizationApiKey, error) {
	var out *admin.CreateAtlasOrganizationApiKey

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

func (opts *createApiKeyOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.CreateApiKeyApiParams{
		OrgId: opts.orgId,

		CreateAtlasOrganizationApiKey: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateApiKeyWithParams(ctx, params).Execute()
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

func createApiKeyBuilder() *cobra.Command {
	opts := createApiKeyOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createApiKey",
		Short: "Create One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type createApiKeyAccessListOpts struct {
	client    *admin.APIClient
	orgId     string
	apiUserId string

	includeCount bool
	itemsPerPage int
	pageNum      int
	filename     string
	fs           afero.Fs
}

func (opts *createApiKeyAccessListOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createApiKeyAccessListOpts) readData() (*[]admin.UserAccessList, error) {
	var out *[]admin.UserAccessList

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

func (opts *createApiKeyAccessListOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.CreateApiKeyAccessListApiParams{
		OrgId:     opts.orgId,
		ApiUserId: opts.apiUserId,

		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,

		UserAccessList: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateApiKeyAccessListWithParams(ctx, params).Execute()
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

func createApiKeyAccessListBuilder() *cobra.Command {
	opts := createApiKeyAccessListOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createApiKeyAccessList",
		Short: "Create Access List Entries for One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.`)

	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type createProjectApiKeyOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createProjectApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createProjectApiKeyOpts) readData() (*admin.CreateAtlasProjectApiKey, error) {
	var out *admin.CreateAtlasProjectApiKey

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

func (opts *createProjectApiKeyOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.CreateProjectApiKeyApiParams{
		GroupId: opts.groupId,

		CreateAtlasProjectApiKey: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateProjectApiKeyWithParams(ctx, params).Execute()
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

func createProjectApiKeyBuilder() *cobra.Command {
	opts := createProjectApiKeyOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createProjectApiKey",
		Short: "Create and Assign One Organization API Key to One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deleteApiKeyOpts struct {
	client    *admin.APIClient
	orgId     string
	apiUserId string
}

func (opts *deleteApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteApiKeyOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.DeleteApiKeyApiParams{
		OrgId:     opts.orgId,
		ApiUserId: opts.apiUserId,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.DeleteApiKeyWithParams(ctx, params).Execute()
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

func deleteApiKeyBuilder() *cobra.Command {
	opts := deleteApiKeyOpts{}
	cmd := &cobra.Command{
		Use:   "deleteApiKey",
		Short: "Remove One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type deleteApiKeyAccessListEntryOpts struct {
	client    *admin.APIClient
	orgId     string
	apiUserId string
	ipAddress string
}

func (opts *deleteApiKeyAccessListEntryOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteApiKeyAccessListEntryOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.DeleteApiKeyAccessListEntryApiParams{
		OrgId:     opts.orgId,
		ApiUserId: opts.apiUserId,
		IpAddress: opts.ipAddress,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.DeleteApiKeyAccessListEntryWithParams(ctx, params).Execute()
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

func deleteApiKeyAccessListEntryBuilder() *cobra.Command {
	opts := deleteApiKeyAccessListEntryOpts{}
	cmd := &cobra.Command{
		Use:   "deleteApiKeyAccessListEntry",
		Short: "Remove One Access List Entry for One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	_ = cmd.MarkFlagRequired("ipAddress")
	return cmd
}

type getApiKeyOpts struct {
	client    *admin.APIClient
	orgId     string
	apiUserId string
}

func (opts *getApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getApiKeyOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetApiKeyApiParams{
		OrgId:     opts.orgId,
		ApiUserId: opts.apiUserId,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.GetApiKeyWithParams(ctx, params).Execute()
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

func getApiKeyBuilder() *cobra.Command {
	opts := getApiKeyOpts{}
	cmd := &cobra.Command{
		Use:   "getApiKey",
		Short: "Return One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type getApiKeyAccessListOpts struct {
	client    *admin.APIClient
	orgId     string
	ipAddress string
	apiUserId string
}

func (opts *getApiKeyAccessListOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getApiKeyAccessListOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetApiKeyAccessListApiParams{
		OrgId:     opts.orgId,
		IpAddress: opts.ipAddress,
		ApiUserId: opts.apiUserId,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.GetApiKeyAccessListWithParams(ctx, params).Execute()
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

func getApiKeyAccessListBuilder() *cobra.Command {
	opts := getApiKeyAccessListOpts{}
	cmd := &cobra.Command{
		Use:   "getApiKeyAccessList",
		Short: "Return One Access List Entry for One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", "", `One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("ipAddress")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type listApiKeyAccessListsEntriesOpts struct {
	client       *admin.APIClient
	orgId        string
	apiUserId    string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listApiKeyAccessListsEntriesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listApiKeyAccessListsEntriesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListApiKeyAccessListsEntriesApiParams{
		OrgId:        opts.orgId,
		ApiUserId:    opts.apiUserId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListApiKeyAccessListsEntriesWithParams(ctx, params).Execute()
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

func listApiKeyAccessListsEntriesBuilder() *cobra.Command {
	opts := listApiKeyAccessListsEntriesOpts{}
	cmd := &cobra.Command{
		Use:   "listApiKeyAccessListsEntries",
		Short: "Return All Access List Entries for One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type listApiKeysOpts struct {
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listApiKeysOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listApiKeysOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListApiKeysApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListApiKeysWithParams(ctx, params).Execute()
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

func listApiKeysBuilder() *cobra.Command {
	opts := listApiKeysOpts{}
	cmd := &cobra.Command{
		Use:   "listApiKeys",
		Short: "Return All Organization API Keys",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listProjectApiKeysOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listProjectApiKeysOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listProjectApiKeysOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListProjectApiKeysApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListProjectApiKeysWithParams(ctx, params).Execute()
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

func listProjectApiKeysBuilder() *cobra.Command {
	opts := listProjectApiKeysOpts{}
	cmd := &cobra.Command{
		Use:   "listProjectApiKeys",
		Short: "Return All Organization API Keys Assigned to One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type removeProjectApiKeyOpts struct {
	client    *admin.APIClient
	groupId   string
	apiUserId string
}

func (opts *removeProjectApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *removeProjectApiKeyOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.RemoveProjectApiKeyApiParams{
		GroupId:   opts.groupId,
		ApiUserId: opts.apiUserId,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.RemoveProjectApiKeyWithParams(ctx, params).Execute()
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

func removeProjectApiKeyBuilder() *cobra.Command {
	opts := removeProjectApiKeyOpts{}
	cmd := &cobra.Command{
		Use:   "removeProjectApiKey",
		Short: "Unassign One Organization API Key from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type updateApiKeyOpts struct {
	client    *admin.APIClient
	orgId     string
	apiUserId string

	filename string
	fs       afero.Fs
}

func (opts *updateApiKeyOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateApiKeyOpts) readData() (*admin.UpdateAtlasOrganizationApiKey, error) {
	var out *admin.UpdateAtlasOrganizationApiKey

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

func (opts *updateApiKeyOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.UpdateApiKeyApiParams{
		OrgId:     opts.orgId,
		ApiUserId: opts.apiUserId,

		UpdateAtlasOrganizationApiKey: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.UpdateApiKeyWithParams(ctx, params).Execute()
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

func updateApiKeyBuilder() *cobra.Command {
	opts := updateApiKeyOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateApiKey",
		Short: "Update One Organization API Key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

type updateApiKeyRolesOpts struct {
	client    *admin.APIClient
	groupId   string
	apiUserId string

	pageNum      int
	itemsPerPage int
	includeCount bool
	filename     string
	fs           afero.Fs
}

func (opts *updateApiKeyRolesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateApiKeyRolesOpts) readData() (*admin.UpdateAtlasProjectApiKey, error) {
	var out *admin.UpdateAtlasProjectApiKey

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

func (opts *updateApiKeyRolesOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.UpdateApiKeyRolesApiParams{
		GroupId:   opts.groupId,
		ApiUserId: opts.apiUserId,

		PageNum:      &opts.pageNum,
		ItemsPerPage: &opts.itemsPerPage,
		IncludeCount: &opts.includeCount,

		UpdateAtlasProjectApiKey: data,
	}

	resp, _, err := opts.client.ProgrammaticAPIKeysApi.UpdateApiKeyRolesWithParams(ctx, params).Execute()
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

func updateApiKeyRolesBuilder() *cobra.Command {
	opts := updateApiKeyRolesOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateApiKeyRoles",
		Short: "Update Roles of One Organization API Key to One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", "", `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`)

	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")
	return cmd
}

func programmaticAPIKeysBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "programmaticAPIKeys",
		Short: `Returns, adds, edits, and removes access tokens to use the MongoDB Cloud API. MongoDB Cloud applies these keys to organizations. These resources can return, assign, or revoke use of these keys within a specified project.`,
	}
	cmd.AddCommand(
		addProjectApiKeyBuilder(),
		createApiKeyBuilder(),
		createApiKeyAccessListBuilder(),
		createProjectApiKeyBuilder(),
		deleteApiKeyBuilder(),
		deleteApiKeyAccessListEntryBuilder(),
		getApiKeyBuilder(),
		getApiKeyAccessListBuilder(),
		listApiKeyAccessListsEntriesBuilder(),
		listApiKeysBuilder(),
		listProjectApiKeysBuilder(),
		removeProjectApiKeyBuilder(),
		updateApiKeyBuilder(),
		updateApiKeyRolesBuilder(),
	)
	return cmd
}
