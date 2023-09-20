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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
}

func (opts *createUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createUserOpts) Run(ctx context.Context) error {
	params := &admin.CreateUserApiParams{}
	resp, _, err := opts.client.MongoDBCloudUsersApi.CreateUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createUserOpts{}
	cmd := &cobra.Command{
		Use:   "createUser",
		Short: "Create One MongoDB Cloud User",
		Annotations: map[string]string{
			"output": template,
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

	cmd.Flags().StringVar(&opts.country, "country", "", `Two alphabet characters that identifies MongoDB Cloud user&#39;s geographic location. This parameter uses the ISO 3166-1a2 code format.`)

	cmd.Flags().StringVar(&opts.createdAt, "createdAt", "", `Date and time when the current account is created. This value is in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.emailAddress, "emailAddress", "", `Email address that belongs to the MongoDB Cloud user.`)

	cmd.Flags().StringVar(&opts.firstName, "firstName", "", `First or given name that belongs to the MongoDB Cloud user.`)

	cmd.Flags().StringVar(&opts.id, "id", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.`)

	cmd.Flags().StringVar(&opts.lastAuth, "lastAuth", "", `Date and time when the current account last authenticated. This value is in the ISO 8601 timestamp format in UTC.`)

	cmd.Flags().StringVar(&opts.lastName, "lastName", "", `Last name, family name, or surname that belongs to the MongoDB Cloud user.`)

	cmd.Flags().ArraySliceVar(&opts.links, "links", nil, `List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.`)

	cmd.Flags().StringVar(&opts.mobileNumber, "mobileNumber", "", `Mobile phone number that belongs to the MongoDB Cloud user.`)

	cmd.Flags().StringVar(&opts.password, "password", "", `Password applied with the username to log in to MongoDB Cloud. MongoDB Cloud does not return this parameter except in response to creating a new MongoDB Cloud user. Only the MongoDB Cloud user can update their password after it has been set from the MongoDB Cloud console.`)

	cmd.Flags().ArraySliceVar(&opts.roles, "roles", nil, `List of objects that display the MongoDB Cloud user&#39;s roles and the corresponding organization or project to which that role applies. A role can apply to one organization or one project but not both.`)

	cmd.Flags().SetSliceVar(&opts.teamIds, "teamIds", nil, `List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs.`)

	cmd.Flags().StringVar(&opts.username, "username", "", `Email address that represents the username of the MongoDB Cloud user.`)

	return cmd
}

type getUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	userId string
}

func (opts *getUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getUserOpts) Run(ctx context.Context) error {
	params := &admin.GetUserApiParams{
		UserId: opts.userId,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getUserOpts{}
	cmd := &cobra.Command{
		Use:   "getUser",
		Short: "Return One MongoDB Cloud User using Its ID",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies this user.`)

	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type getUserByUsernameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client   *admin.APIClient
	userName string
}

func (opts *getUserByUsernameOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getUserByUsernameOpts) Run(ctx context.Context) error {
	params := &admin.GetUserByUsernameApiParams{
		UserName: opts.userName,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserByUsernameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getUserByUsernameBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getUserByUsernameOpts{}
	cmd := &cobra.Command{
		Use:   "getUserByUsername",
		Short: "Return One MongoDB Cloud User using Their Username",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.userName, "userName", "", `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`)

	_ = cmd.MarkFlagRequired("userName")
	return cmd
}

func mongoDBCloudUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mongoDBCloudUsers",
		Short: `Returns, adds, and edits MongoDB Cloud users.`,
	}
	cmd.AddCommand(
		createUserBuilder(),
		getUserBuilder(),
		getUserByUsernameBuilder(),
	)
	return cmd
}
