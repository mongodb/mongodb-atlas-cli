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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/jsonwriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient

	filename string
	fs       afero.Fs
}

func (opts *createOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createOrganizationOpts) readData() (*admin.CreateOrganizationRequest, error) {
	var out *admin.CreateOrganizationRequest

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

func (opts *createOrganizationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateOrganizationApiParams{

		CreateOrganizationRequest: data,
	}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createOrganizationBuilder() *cobra.Command {
	opts := createOrganizationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createOrganization",
		Short: "Create One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

type createOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *createOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createOrganizationInvitationOpts) readData() (*admin.OrganizationInvitationRequest, error) {
	var out *admin.OrganizationInvitationRequest

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

func (opts *createOrganizationInvitationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateOrganizationInvitationApiParams{
		OrgId: opts.orgId,

		OrganizationInvitationRequest: data,
	}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createOrganizationInvitationBuilder() *cobra.Command {
	opts := createOrganizationInvitationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createOrganizationInvitation",
		Short: "Invite One MongoDB Cloud User to Join One Atlas Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *deleteOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteOrganizationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeleteOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deleteOrganizationBuilder() *cobra.Command {
	opts := deleteOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOrganization",
		Short: "Remove One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string
}

func (opts *deleteOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteOrganizationInvitationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeleteOrganizationInvitationApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deleteOrganizationInvitationBuilder() *cobra.Command {
	opts := deleteOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOrganizationInvitation",
		Short: "Cancel One Organization Invitation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type getOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *getOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getOrganizationBuilder() *cobra.Command {
	opts := getOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganization",
		Short: "Return One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string
}

func (opts *getOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationInvitationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetOrganizationInvitationApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getOrganizationInvitationBuilder() *cobra.Command {
	opts := getOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganizationInvitation",
		Short: "Return One Organization Invitation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type getOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *getOrganizationSettingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationSettingsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetOrganizationSettingsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getOrganizationSettingsBuilder() *cobra.Command {
	opts := getOrganizationSettingsOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganizationSettings",
		Short: "Return Settings for One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationInvitationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client   *admin.APIClient
	orgId    string
	username string
}

func (opts *listOrganizationInvitationsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationInvitationsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListOrganizationInvitationsApiParams{
		OrgId:    opts.orgId,
		Username: &opts.username,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationInvitationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listOrganizationInvitationsBuilder() *cobra.Command {
	opts := listOrganizationInvitationsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationInvitations",
		Short: "Return All Organization Invitations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationProjectsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
	name         string
}

func (opts *listOrganizationProjectsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationProjectsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListOrganizationProjectsApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Name:         &opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationProjectsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listOrganizationProjectsBuilder() *cobra.Command {
	opts := listOrganizationProjectsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationProjects",
		Short: "Return One or More Projects in One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.name, "name", "", `Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listOrganizationUsersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationUsersOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListOrganizationUsersApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listOrganizationUsersBuilder() *cobra.Command {
	opts := listOrganizationUsersOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationUsers",
		Short: "Return All MongoDB Cloud Users in One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	includeCount bool
	itemsPerPage int
	pageNum      int
	name         string
}

func (opts *listOrganizationsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListOrganizationsApiParams{
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Name:         &opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listOrganizationsBuilder() *cobra.Command {
	opts := listOrganizationsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizations",
		Short: "Return All Organizations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.name, "name", "", `Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

type removeOrganizationUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
	userId string
}

func (opts *removeOrganizationUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *removeOrganizationUserOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.RemoveOrganizationUserApiParams{
		OrgId:  opts.orgId,
		UserId: opts.userId,
	}
	resp, _, err := opts.client.OrganizationsApi.RemoveOrganizationUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func removeOrganizationUserBuilder() *cobra.Command {
	opts := removeOrganizationUserOpts{}
	cmd := &cobra.Command{
		Use:   "removeOrganizationUser",
		Short: "Remove One MongoDB Cloud User from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies the user to be deleted.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type renameOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *renameOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *renameOrganizationOpts) readData() (*admin.AtlasOrganization, error) {
	var out *admin.AtlasOrganization

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

func (opts *renameOrganizationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.RenameOrganizationApiParams{
		OrgId: opts.orgId,

		AtlasOrganization: data,
	}
	resp, _, err := opts.client.OrganizationsApi.RenameOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func renameOrganizationBuilder() *cobra.Command {
	opts := renameOrganizationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "renameOrganization",
		Short: "Rename One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *updateOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationInvitationOpts) readData() (*admin.OrganizationInvitationRequest, error) {
	var out *admin.OrganizationInvitationRequest

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

func (opts *updateOrganizationInvitationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateOrganizationInvitationApiParams{
		OrgId: opts.orgId,

		OrganizationInvitationRequest: data,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateOrganizationInvitationBuilder() *cobra.Command {
	opts := updateOrganizationInvitationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateOrganizationInvitation",
		Short: "Update One Organization Invitation",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateOrganizationInvitationByIdOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string

	filename string
	fs       afero.Fs
}

func (opts *updateOrganizationInvitationByIdOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationInvitationByIdOpts) readData() (*admin.OrganizationInvitationUpdateRequest, error) {
	var out *admin.OrganizationInvitationUpdateRequest

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

func (opts *updateOrganizationInvitationByIdOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateOrganizationInvitationByIdApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,

		OrganizationInvitationUpdateRequest: data,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationByIdWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateOrganizationInvitationByIdBuilder() *cobra.Command {
	opts := updateOrganizationInvitationByIdOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateOrganizationInvitationById",
		Short: "Update One Organization Invitation by Invitation ID",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type updateOrganizationRolesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
	userId string

	filename string
	fs       afero.Fs
}

func (opts *updateOrganizationRolesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationRolesOpts) readData() (*admin.UpdateOrgRolesForUser, error) {
	var out *admin.UpdateOrgRolesForUser

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

func (opts *updateOrganizationRolesOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateOrganizationRolesApiParams{
		OrgId:  opts.orgId,
		UserId: opts.userId,

		UpdateOrgRolesForUser: data,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationRolesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateOrganizationRolesBuilder() *cobra.Command {
	opts := updateOrganizationRolesOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateOrganizationRoles",
		Short: "Update Organization Roles for One MongoDB Cloud User",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies the user to modify.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type updateOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *updateOrganizationSettingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationSettingsOpts) readData() (*admin.OrganizationSettings, error) {
	var out *admin.OrganizationSettings

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

func (opts *updateOrganizationSettingsOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateOrganizationSettingsApiParams{
		OrgId: opts.orgId,

		OrganizationSettings: data,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateOrganizationSettingsBuilder() *cobra.Command {
	opts := updateOrganizationSettingsOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateOrganizationSettings",
		Short: "Update Settings for One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

func organizationsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organizations",
		Short: `Returns, adds, and edits organizational units in MongoDB Cloud.`,
	}
	cmd.AddCommand(
		createOrganizationBuilder(),
		createOrganizationInvitationBuilder(),
		deleteOrganizationBuilder(),
		deleteOrganizationInvitationBuilder(),
		getOrganizationBuilder(),
		getOrganizationInvitationBuilder(),
		getOrganizationSettingsBuilder(),
		listOrganizationInvitationsBuilder(),
		listOrganizationProjectsBuilder(),
		listOrganizationUsersBuilder(),
		listOrganizationsBuilder(),
		removeOrganizationUserBuilder(),
		renameOrganizationBuilder(),
		updateOrganizationInvitationBuilder(),
		updateOrganizationInvitationByIdBuilder(),
		updateOrganizationRolesBuilder(),
		updateOrganizationSettingsBuilder(),
	)
	return cmd
}
