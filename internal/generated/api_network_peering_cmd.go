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

type createPeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *createPeeringConnectionOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createPeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.CreatePeeringConnectionApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createPeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createPeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "createPeeringConnection",
		Short: "Create One New Network Peering Connection",
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

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type createPeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *createPeeringContainerOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createPeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.CreatePeeringContainerApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createPeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createPeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "createPeeringContainer",
		Short: "Create One New Network Peering Container",
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

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deletePeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
	peerId  string
}

func (opts *deletePeeringConnectionOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.DeletePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deletePeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deletePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "deletePeeringConnection",
		Short: "Remove One Existing Network Peering Connection",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type deletePeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	containerId string
}

func (opts *deletePeeringContainerOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.DeletePeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deletePeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deletePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "deletePeeringContainer",
		Short: "Remove One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type disablePeeringOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *disablePeeringOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *disablePeeringOpts) Run(ctx context.Context) error {
	params := &admin.DisablePeeringApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DisablePeeringWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func disablePeeringBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := disablePeeringOpts{}
	cmd := &cobra.Command{
		Use:   "disablePeering",
		Short: "Disable Connect via Peering Only Mode for One Project",
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

	cmd.Flags().BoolVar(&opts.enabled, "enabled", false, `Flag that indicates whether someone enabled **Connect via Peering Only** mode for the specified project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getPeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
	peerId  string
}

func (opts *getPeeringConnectionOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.GetPeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getPeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getPeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "getPeeringConnection",
		Short: "Return One Network Peering Connection in One Project",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type getPeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	containerId string
}

func (opts *getPeeringContainerOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.GetPeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getPeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getPeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "getPeeringContainer",
		Short: "Return One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type listPeeringConnectionsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
	providerName string
}

func (opts *listPeeringConnectionsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPeeringConnectionsOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringConnectionsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		ProviderName: &opts.providerName,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringConnectionsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listPeeringConnectionsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listPeeringConnectionsOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringConnections",
		Short: "Return All Network Peering Connections in One Project",
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
	cmd.Flags().StringVar(&opts.providerName, "providerName", "&quot;AWS&quot;", `Cloud service provider to use for this VPC peering connection.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type listPeeringContainerByCloudProviderOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	providerName string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listPeeringContainerByCloudProviderOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPeeringContainerByCloudProviderOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringContainerByCloudProviderApiParams{
		GroupId:      opts.groupId,
		ProviderName: &opts.providerName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainerByCloudProviderWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listPeeringContainerByCloudProviderBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listPeeringContainerByCloudProviderOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringContainerByCloudProvider",
		Short: "Return All Network Peering Containers in One Project for One Cloud Provider",
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
	cmd.Flags().StringVar(&opts.providerName, "providerName", "&quot;AWS&quot;", `Cloud service provider that serves the desired network peering containers.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("providerName")
	return cmd
}

type listPeeringContainersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listPeeringContainersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPeeringContainersOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringContainersApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listPeeringContainersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listPeeringContainersOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringContainers",
		Short: "Return All Network Peering Containers in One Project",
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

type updatePeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
	peerId  string
}

func (opts *updatePeeringConnectionOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updatePeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updatePeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updatePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "updatePeeringConnection",
		Short: "Update One New Network Peering Connection",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type updatePeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	containerId string
}

func (opts *updatePeeringContainerOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updatePeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updatePeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updatePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "updatePeeringContainer",
		Short: "Update One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type verifyConnectViaPeeringOnlyModeForOneProjectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *verifyConnectViaPeeringOnlyModeForOneProjectOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *verifyConnectViaPeeringOnlyModeForOneProjectOpts) Run(ctx context.Context) error {
	params := &admin.VerifyConnectViaPeeringOnlyModeForOneProjectApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProjectWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func verifyConnectViaPeeringOnlyModeForOneProjectBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := verifyConnectViaPeeringOnlyModeForOneProjectOpts{}
	cmd := &cobra.Command{
		Use:   "verifyConnectViaPeeringOnlyModeForOneProject",
		Short: "Verify Connect via Peering Only Mode for One Project",
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

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func networkPeeringBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use: "networkPeering",
		Short: `Returns, adds, edits, and removes network peering containers and peering connections.
When you deploy an M10+ dedicated cluster, Atlas creates a VPC for the selected provider and region or regions if no existing VPC or VPC peering connection exists for that provider and region. Atlas assigns the VPC a Classless Inter-Domain Routing (CIDR) block.`,
	}
	cmd.AddCommand(
		createPeeringConnectionBuilder(),
		createPeeringContainerBuilder(),
		deletePeeringConnectionBuilder(),
		deletePeeringContainerBuilder(),
		disablePeeringBuilder(),
		getPeeringConnectionBuilder(),
		getPeeringContainerBuilder(),
		listPeeringConnectionsBuilder(),
		listPeeringContainerByCloudProviderBuilder(),
		listPeeringContainersBuilder(),
		updatePeeringConnectionBuilder(),
		updatePeeringContainerBuilder(),
		verifyConnectViaPeeringOnlyModeForOneProjectBuilder(),
	)
	return cmd
}
