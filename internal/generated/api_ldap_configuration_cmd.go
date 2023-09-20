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

type deleteLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *deleteLDAPConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteLDAPConfigurationOpts) Run(ctx context.Context, _ io.Writer) error {
	params := &admin.DeleteLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.client.LDAPConfigurationApi.DeleteLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return nil
}

func deleteLDAPConfigurationBuilder() *cobra.Command {
	opts := deleteLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteLDAPConfiguration",
		Short: "Remove the Current LDAP User to DN Mapping",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *getLDAPConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getLDAPConfigurationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.GetLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getLDAPConfigurationBuilder() *cobra.Command {
	opts := getLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "getLDAPConfiguration",
		Short: "Return the Current LDAP or X.509 Configuration",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getLDAPConfigurationStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client    *admin.APIClient
	groupId   string
	requestId string
}

func (opts *getLDAPConfigurationStatusOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getLDAPConfigurationStatusOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetLDAPConfigurationStatusApiParams{
		GroupId:   opts.groupId,
		RequestId: opts.requestId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.GetLDAPConfigurationStatusWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getLDAPConfigurationStatusBuilder() *cobra.Command {
	opts := getLDAPConfigurationStatusOpts{}
	cmd := &cobra.Command{
		Use:   "getLDAPConfigurationStatus",
		Short: "Return the Status of One Verify LDAP Configuration Request",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.requestId, "requestId", "", `Unique string that identifies the request to verify an &lt;abbr title&#x3D;&quot;Lightweight Directory Access Protocol&quot;&gt;LDAP&lt;/abbr&gt; configuration.`)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("requestId")
	return cmd
}

type saveLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *saveLDAPConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *saveLDAPConfigurationOpts) readData() (*admin.UserSecurity, error) {
	var out *admin.UserSecurity

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

func (opts *saveLDAPConfigurationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.SaveLDAPConfigurationApiParams{
		GroupId: opts.groupId,

		UserSecurity: data,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.SaveLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func saveLDAPConfigurationBuilder() *cobra.Command {
	opts := saveLDAPConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "saveLDAPConfiguration",
		Short: "Edit the LDAP or X.509 Configuration",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type verifyLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *verifyLDAPConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *verifyLDAPConfigurationOpts) readData() (*admin.NDSLDAPVerifyConnectivityJobRequestParams, error) {
	var out *admin.NDSLDAPVerifyConnectivityJobRequestParams

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

func (opts *verifyLDAPConfigurationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.VerifyLDAPConfigurationApiParams{
		GroupId: opts.groupId,

		NDSLDAPVerifyConnectivityJobRequestParams: data,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.VerifyLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func verifyLDAPConfigurationBuilder() *cobra.Command {
	opts := verifyLDAPConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "verifyLDAPConfiguration",
		Short: "Verify the LDAP Configuration in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func lDAPConfigurationBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lDAPConfiguration",
		Short: `Returns, edits, verifies, and removes LDAP configurations. An LDAP configuration defines settings for MongoDB Cloud to connect to your LDAP server over TLS for user authentication and authorization. Your LDAP server must be visible to the internet or connected to your MongoDB Cloud cluster with VPC Peering. Also, your LDAP server must use TLS. You must have the MongoDB Cloud admin user privilege to use these endpoints. Also, to configure user authentication and authorization with LDAPS, your cluster must run MongoDB 3.6 or higher. Groups for which you have configured LDAPS can&#39;t create a cluster using a version of MongoDB 3.6 or lower.`,
	}
	cmd.AddCommand(
		deleteLDAPConfigurationBuilder(),
		getLDAPConfigurationBuilder(),
		getLDAPConfigurationStatusBuilder(),
		saveLDAPConfigurationBuilder(),
		verifyLDAPConfigurationBuilder(),
	)
	return cmd
}
