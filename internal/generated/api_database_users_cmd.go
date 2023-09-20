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

type createDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *createDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.CreateDatabaseUserApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DatabaseUsersApi.CreateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "createDatabaseUser",
		Short: "Create One Database User in One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVar(&opts.awsIAMType, "awsIAMType", "&quot;NONE&quot;", `Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role.`)

	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "&quot;admin&quot;", `Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.`)

	cmd.Flags().StringVar(&opts.deleteAfterDate, "deleteAfterDate", "", `Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.`)

	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies the project.`)

	cmd.Flags().ArraySliceVar(&opts.labels, "labels", nil, `List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console.`)

	cmd.Flags().StringVar(&opts.ldapAuthType, "ldapAuthType", "&quot;NONE&quot;", `Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host.`)

	cmd.Flags().ArraySliceVar(&opts.links, "links", nil, `List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.`)

	cmd.Flags().StringVar(&opts.password, "password", "", `Alphanumeric string that authenticates this database user against the database specified in &#x60;databaseName&#x60;. To authenticate with SCRAM-SHA, you must specify this parameter. This parameter doesn&#39;t appear in this response.`)

	cmd.Flags().ArraySliceVar(&opts.roles, "roles", nil, `List that provides the pairings of one role with one applicable database.`)

	cmd.Flags().ArraySliceVar(&opts.scopes, "scopes", nil, `List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.`)

	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
`)

	cmd.Flags().StringVar(&opts.x509Type, "x509Type", "&quot;NONE&quot;", `X.509 method that MongoDB Cloud uses to authenticate the database user.

- For application-managed X.509, specify &#x60;MANAGED&#x60;.
- For self-managed X.509, specify &#x60;CUSTOMER&#x60;.

Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deleteDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string
}

func (opts *deleteDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.DeleteDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,
	}
	resp, _, err := opts.client.DatabaseUsersApi.DeleteDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deleteDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deleteDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "deleteDatabaseUser",
		Short: "Remove One Database User from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

type getDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string
}

func (opts *getDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.GetDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,
	}
	resp, _, err := opts.client.DatabaseUsersApi.GetDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "getDatabaseUser",
		Short: "Return One Database User from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

type listDatabaseUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listDatabaseUsersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listDatabaseUsersOpts) Run(ctx context.Context) error {
	params := &admin.ListDatabaseUsersApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.DatabaseUsersApi.ListDatabaseUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listDatabaseUsersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listDatabaseUsersOpts{}
	cmd := &cobra.Command{
		Use:   "listDatabaseUsers",
		Short: "Return All Database Users from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type updateDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string
}

func (opts *updateDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,
	}
	resp, _, err := opts.client.DatabaseUsersApi.UpdateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "updateDatabaseUser",
		Short: "Update One Database User in One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
`)

	cmd.Flags().StringVar(&opts.awsIAMType, "awsIAMType", "&quot;NONE&quot;", `Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role.`)

	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "&quot;admin&quot;", `Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.`)

	cmd.Flags().StringVar(&opts.deleteAfterDate, "deleteAfterDate", "", `Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.`)

	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies the project.`)

	cmd.Flags().ArraySliceVar(&opts.labels, "labels", nil, `List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console.`)

	cmd.Flags().StringVar(&opts.ldapAuthType, "ldapAuthType", "&quot;NONE&quot;", `Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host.`)

	cmd.Flags().ArraySliceVar(&opts.links, "links", nil, `List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.`)

	cmd.Flags().StringVar(&opts.password, "password", "", `Alphanumeric string that authenticates this database user against the database specified in &#x60;databaseName&#x60;. To authenticate with SCRAM-SHA, you must specify this parameter. This parameter doesn&#39;t appear in this response.`)

	cmd.Flags().ArraySliceVar(&opts.roles, "roles", nil, `List that provides the pairings of one role with one applicable database.`)

	cmd.Flags().ArraySliceVar(&opts.scopes, "scopes", nil, `List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.`)

	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |
`)

	cmd.Flags().StringVar(&opts.x509Type, "x509Type", "&quot;NONE&quot;", `X.509 method that MongoDB Cloud uses to authenticate the database user.

- For application-managed X.509, specify &#x60;MANAGED&#x60;.
- For self-managed X.509, specify &#x60;CUSTOMER&#x60;.

Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

func databaseUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "databaseUsers",
		Short: `Returns, adds, edits, and removes database users.`,
	}
	cmd.AddCommand(
		createDatabaseUserBuilder(),
		deleteDatabaseUserBuilder(),
		getDatabaseUserBuilder(),
		listDatabaseUsersBuilder(),
		updateDatabaseUserBuilder(),
	)
	return cmd
}
