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

type addAllTeamsToProjectOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *addAllTeamsToProjectOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *addAllTeamsToProjectOpts) readData() (*[]admin.TeamRole, error) {
	var out *[]admin.TeamRole

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

func (opts *addAllTeamsToProjectOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.AddAllTeamsToProjectApiParams{
		GroupId: opts.groupId,

		TeamRole: data,
	}

	resp, _, err := opts.client.TeamsApi.AddAllTeamsToProjectWithParams(ctx, params).Execute()
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

func addAllTeamsToProjectBuilder() *cobra.Command {
	opts := addAllTeamsToProjectOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "addAllTeamsToProject",
		Short: "Add One or More Teams to One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type addTeamUserOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string

	filename string
	fs       afero.Fs
}

func (opts *addTeamUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *addTeamUserOpts) readData() (*[]admin.AddUserToTeam, error) {
	var out *[]admin.AddUserToTeam

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

func (opts *addTeamUserOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.AddTeamUserApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,

		AddUserToTeam: data,
	}

	resp, _, err := opts.client.TeamsApi.AddTeamUserWithParams(ctx, params).Execute()
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

func addTeamUserBuilder() *cobra.Command {
	opts := addTeamUserOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "addTeamUser",
		Short: "Assign MongoDB Cloud Users from One Organization to One Team",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type createTeamOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *createTeamOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createTeamOpts) readData() (*admin.Team, error) {
	var out *admin.Team

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

func (opts *createTeamOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.CreateTeamApiParams{
		OrgId: opts.orgId,

		Team: data,
	}

	resp, _, err := opts.client.TeamsApi.CreateTeamWithParams(ctx, params).Execute()
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

func createTeamBuilder() *cobra.Command {
	opts := createTeamOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createTeam",
		Short: "Create One Team in One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteTeamOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
}

func (opts *deleteTeamOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteTeamOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.DeleteTeamApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
	}

	resp, _, err := opts.client.TeamsApi.DeleteTeamWithParams(ctx, params).Execute()
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

func deleteTeamBuilder() *cobra.Command {
	opts := deleteTeamOpts{}
	cmd := &cobra.Command{
		Use:   "deleteTeam",
		Short: "Remove One Team from One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to delete.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type getTeamByIdOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
}

func (opts *getTeamByIdOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getTeamByIdOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.GetTeamByIdApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
	}

	resp, _, err := opts.client.TeamsApi.GetTeamByIdWithParams(ctx, params).Execute()
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

func getTeamByIdBuilder() *cobra.Command {
	opts := getTeamByIdOpts{}
	cmd := &cobra.Command{
		Use:   "getTeamById",
		Short: "Return One Team using its ID",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team whose information you want to return.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type getTeamByNameOpts struct {
	client   *admin.APIClient
	orgId    string
	teamName string
}

func (opts *getTeamByNameOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getTeamByNameOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.GetTeamByNameApiParams{
		OrgId:    opts.orgId,
		TeamName: opts.teamName,
	}

	resp, _, err := opts.client.TeamsApi.GetTeamByNameWithParams(ctx, params).Execute()
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

func getTeamByNameBuilder() *cobra.Command {
	opts := getTeamByNameOpts{}
	cmd := &cobra.Command{
		Use:   "getTeamByName",
		Short: "Return One Team using its Name",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamName, "teamName", "", `Name of the team whose information you want to return.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamName")
	return cmd
}

type listOrganizationTeamsOpts struct {
	client       *admin.APIClient
	orgId        string
	itemsPerPage int
	includeCount bool
	pageNum      int
}

func (opts *listOrganizationTeamsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listOrganizationTeamsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.ListOrganizationTeamsApiParams{
		OrgId:        opts.orgId,
		ItemsPerPage: &opts.itemsPerPage,
		IncludeCount: &opts.includeCount,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.TeamsApi.ListOrganizationTeamsWithParams(ctx, params).Execute()
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

func listOrganizationTeamsBuilder() *cobra.Command {
	opts := listOrganizationTeamsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationTeams",
		Short: "Return All Teams in One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listProjectTeamsOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listProjectTeamsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listProjectTeamsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListProjectTeamsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.TeamsApi.ListProjectTeamsWithParams(ctx, params).Execute()
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

func listProjectTeamsBuilder() *cobra.Command {
	opts := listProjectTeamsOpts{}
	cmd := &cobra.Command{
		Use:   "listProjectTeams",
		Short: "Return All Teams in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type listTeamUsersOpts struct {
	client       *admin.APIClient
	orgId        string
	teamId       string
	itemsPerPage int
	pageNum      int
}

func (opts *listTeamUsersOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listTeamUsersOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.ListTeamUsersApiParams{
		OrgId:        opts.orgId,
		TeamId:       opts.teamId,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.TeamsApi.ListTeamUsersWithParams(ctx, params).Execute()
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

func listTeamUsersBuilder() *cobra.Command {
	opts := listTeamUsersOpts{}
	cmd := &cobra.Command{
		Use:   "listTeamUsers",
		Short: "Return All MongoDB Cloud Users Assigned to One Team",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type removeProjectTeamOpts struct {
	client  *admin.APIClient
	groupId string
	teamId  string
}

func (opts *removeProjectTeamOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *removeProjectTeamOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.RemoveProjectTeamApiParams{
		GroupId: opts.groupId,
		TeamId:  opts.teamId,
	}

	_, err := opts.client.TeamsApi.RemoveProjectTeamWithParams(ctx, params).Execute()
	return err
}

func removeProjectTeamBuilder() *cobra.Command {
	opts := removeProjectTeamOpts{}
	cmd := &cobra.Command{
		Use:   "removeProjectTeam",
		Short: "Remove One Team from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type removeTeamUserOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
	userId string
}

func (opts *removeTeamUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *removeTeamUserOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.RemoveTeamUserApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
		UserId: opts.userId,
	}

	_, err := opts.client.TeamsApi.RemoveTeamUserWithParams(ctx, params).Execute()
	return err
}

func removeTeamUserBuilder() *cobra.Command {
	opts := removeTeamUserOpts{}
	cmd := &cobra.Command{
		Use:   "removeTeamUser",
		Short: "Remove One MongoDB Cloud User from One Team",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.`)
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type renameTeamOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string

	filename string
	fs       afero.Fs
}

func (opts *renameTeamOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *renameTeamOpts) readData() (*admin.Team, error) {
	var out *admin.Team

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

func (opts *renameTeamOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.RenameTeamApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,

		Team: data,
	}

	resp, _, err := opts.client.TeamsApi.RenameTeamWithParams(ctx, params).Execute()
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

func renameTeamBuilder() *cobra.Command {
	opts := renameTeamOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "renameTeam",
		Short: "Rename One Team",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to rename.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type updateTeamRolesOpts struct {
	client  *admin.APIClient
	groupId string
	teamId  string

	filename string
	fs       afero.Fs
}

func (opts *updateTeamRolesOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateTeamRolesOpts) readData() (*admin.TeamRole, error) {
	var out *admin.TeamRole

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

func (opts *updateTeamRolesOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdateTeamRolesApiParams{
		GroupId: opts.groupId,
		TeamId:  opts.teamId,

		TeamRole: data,
	}

	resp, _, err := opts.client.TeamsApi.UpdateTeamRolesWithParams(ctx, params).Execute()
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

func updateTeamRolesBuilder() *cobra.Command {
	opts := updateTeamRolesOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateTeamRoles",
		Short: "Update Team Roles in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

func teamsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "teams",
		Short: `Returns, adds, edits, or removes teams.`,
	}
	cmd.AddCommand(
		addAllTeamsToProjectBuilder(),
		addTeamUserBuilder(),
		createTeamBuilder(),
		deleteTeamBuilder(),
		getTeamByIdBuilder(),
		getTeamByNameBuilder(),
		listOrganizationTeamsBuilder(),
		listProjectTeamsBuilder(),
		listTeamUsersBuilder(),
		removeProjectTeamBuilder(),
		removeTeamUserBuilder(),
		renameTeamBuilder(),
		updateTeamRolesBuilder(),
	)
	return cmd
}
