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
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type authorizeCloudProviderAccessRoleOpts struct {
	client  *admin.APIClient
	groupId string
	roleId  string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *authorizeCloudProviderAccessRoleOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *authorizeCloudProviderAccessRoleOpts) readData(r io.Reader) (*admin.CloudProviderAccessRole, error) {
	var out *admin.CloudProviderAccessRole

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(r)
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

func (opts *authorizeCloudProviderAccessRoleOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.AuthorizeCloudProviderAccessRoleApiParams{
		GroupId: opts.groupId,
		RoleId:  opts.roleId,

		CloudProviderAccessRole: data,
	}

	resp, _, err := opts.client.CloudProviderAccessApi.AuthorizeCloudProviderAccessRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func authorizeCloudProviderAccessRoleBuilder() *cobra.Command {
	opts := authorizeCloudProviderAccessRoleOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "authorizeCloudProviderAccessRole",
		Short: "Authorize One Cloud Provider Access Role",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.roleId, "roleId", "", `Unique 24-hexadecimal digit string that identifies the role.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("roleId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type createCloudProviderAccessRoleOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *createCloudProviderAccessRoleOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *createCloudProviderAccessRoleOpts) readData(r io.Reader) (*admin.CloudProviderAccessRole, error) {
	var out *admin.CloudProviderAccessRole

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(r)
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

func (opts *createCloudProviderAccessRoleOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateCloudProviderAccessRoleApiParams{
		GroupId: opts.groupId,

		CloudProviderAccessRole: data,
	}

	resp, _, err := opts.client.CloudProviderAccessApi.CreateCloudProviderAccessRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func createCloudProviderAccessRoleBuilder() *cobra.Command {
	opts := createCloudProviderAccessRoleOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createCloudProviderAccessRole",
		Short: "Create One Cloud Provider Access Role",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type deauthorizeCloudProviderAccessRoleOpts struct {
	client        *admin.APIClient
	groupId       string
	cloudProvider string
	roleId        string
}

func (opts *deauthorizeCloudProviderAccessRoleOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	return err
}

func (opts *deauthorizeCloudProviderAccessRoleOpts) run(ctx context.Context, _ io.Reader, _ io.Writer) error {

	params := &admin.DeauthorizeCloudProviderAccessRoleApiParams{
		GroupId:       opts.groupId,
		CloudProvider: opts.cloudProvider,
		RoleId:        opts.roleId,
	}

	_, err := opts.client.CloudProviderAccessApi.DeauthorizeCloudProviderAccessRoleWithParams(ctx, params).Execute()
	return err
}

func deauthorizeCloudProviderAccessRoleBuilder() *cobra.Command {
	opts := deauthorizeCloudProviderAccessRoleOpts{}
	cmd := &cobra.Command{
		Use:   "deauthorizeCloudProviderAccessRole",
		Short: "Deauthorize One Cloud Provider Access Role",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "", `Human-readable label that identifies the cloud provider of the role to deauthorize.`)
	cmd.Flags().StringVar(&opts.roleId, "roleId", "", `Unique 24-hexadecimal digit string that identifies the role.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("roleId")
	return cmd
}

type getCloudProviderAccessRoleOpts struct {
	client  *admin.APIClient
	groupId string
	roleId  string
	format  string
	tmpl    *template.Template
}

func (opts *getCloudProviderAccessRoleOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *getCloudProviderAccessRoleOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetCloudProviderAccessRoleApiParams{
		GroupId: opts.groupId,
		RoleId:  opts.roleId,
	}

	resp, _, err := opts.client.CloudProviderAccessApi.GetCloudProviderAccessRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func getCloudProviderAccessRoleBuilder() *cobra.Command {
	opts := getCloudProviderAccessRoleOpts{}
	cmd := &cobra.Command{
		Use:   "getCloudProviderAccessRole",
		Short: "Return specified Cloud Provider Access Role",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.roleId, "roleId", "", `Unique 24-hexadecimal digit string that identifies the role.`)

	_ = cmd.MarkFlagRequired("roleId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listCloudProviderAccessRolesOpts struct {
	client  *admin.APIClient
	groupId string
	format  string
	tmpl    *template.Template
}

func (opts *listCloudProviderAccessRolesOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *listCloudProviderAccessRolesOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.ListCloudProviderAccessRolesApiParams{
		GroupId: opts.groupId,
	}

	resp, _, err := opts.client.CloudProviderAccessApi.ListCloudProviderAccessRolesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func listCloudProviderAccessRolesBuilder() *cobra.Command {
	opts := listCloudProviderAccessRolesOpts{}
	cmd := &cobra.Command{
		Use:   "listCloudProviderAccessRoles",
		Short: "Return All Cloud Provider Access Roles",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func cloudProviderAccessBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloudProviderAccess",
		Short: `Returns, adds, authorizes, and removes AWS IAM roles in Atlas.`,
	}
	cmd.AddCommand(
		authorizeCloudProviderAccessRoleBuilder(),
		createCloudProviderAccessRoleBuilder(),
		deauthorizeCloudProviderAccessRoleBuilder(),
		getCloudProviderAccessRoleBuilder(),
		listCloudProviderAccessRolesBuilder(),
	)
	return cmd
}
