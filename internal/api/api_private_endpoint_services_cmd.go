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

type createPrivateEndpointOpts struct {
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *createPrivateEndpointOpts) preRun() (err error) {
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

func (opts *createPrivateEndpointOpts) readData(r io.Reader) (*admin.CreateEndpointRequest, error) {
	var out *admin.CreateEndpointRequest

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

func (opts *createPrivateEndpointOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreatePrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,

		CreateEndpointRequest: data,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.CreatePrivateEndpointWithParams(ctx, params).Execute()
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

func createPrivateEndpointBuilder() *cobra.Command {
	opts := createPrivateEndpointOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPrivateEndpoint",
		Short: "Create One Private Endpoint for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type createPrivateEndpointServiceOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *createPrivateEndpointServiceOpts) preRun() (err error) {
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

func (opts *createPrivateEndpointServiceOpts) readData(r io.Reader) (*admin.CloudProviderEndpointServiceRequest, error) {
	var out *admin.CloudProviderEndpointServiceRequest

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

func (opts *createPrivateEndpointServiceOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreatePrivateEndpointServiceApiParams{
		GroupId: opts.groupId,

		CloudProviderEndpointServiceRequest: data,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.CreatePrivateEndpointServiceWithParams(ctx, params).Execute()
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

func createPrivateEndpointServiceBuilder() *cobra.Command {
	opts := createPrivateEndpointServiceOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPrivateEndpointService",
		Short: "Create One Private Endpoint Service for One Provider",
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

type deletePrivateEndpointOpts struct {
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
	format            string
	tmpl              *template.Template
}

func (opts *deletePrivateEndpointOpts) preRun() (err error) {
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

func (opts *deletePrivateEndpointOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.DeletePrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointId:        opts.endpointId,
		EndpointServiceId: opts.endpointServiceId,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.DeletePrivateEndpointWithParams(ctx, params).Execute()
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

func deletePrivateEndpointBuilder() *cobra.Command {
	opts := deletePrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "deletePrivateEndpoint",
		Short: "Remove One Private Endpoint for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointId")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type deletePrivateEndpointServiceOpts struct {
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string
	format            string
	tmpl              *template.Template
}

func (opts *deletePrivateEndpointServiceOpts) preRun() (err error) {
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

func (opts *deletePrivateEndpointServiceOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.DeletePrivateEndpointServiceApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.DeletePrivateEndpointServiceWithParams(ctx, params).Execute()
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

func deletePrivateEndpointServiceBuilder() *cobra.Command {
	opts := deletePrivateEndpointServiceOpts{}
	cmd := &cobra.Command{
		Use:   "deletePrivateEndpointService",
		Short: "Remove One Private Endpoint Service for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getPrivateEndpointOpts struct {
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
	format            string
	tmpl              *template.Template
}

func (opts *getPrivateEndpointOpts) preRun() (err error) {
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

func (opts *getPrivateEndpointOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetPrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointId:        opts.endpointId,
		EndpointServiceId: opts.endpointServiceId,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.GetPrivateEndpointWithParams(ctx, params).Execute()
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

func getPrivateEndpointBuilder() *cobra.Command {
	opts := getPrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "getPrivateEndpoint",
		Short: "Return One Private Endpoint for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointId")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getPrivateEndpointServiceOpts struct {
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string
	format            string
	tmpl              *template.Template
}

func (opts *getPrivateEndpointServiceOpts) preRun() (err error) {
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

func (opts *getPrivateEndpointServiceOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetPrivateEndpointServiceApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.GetPrivateEndpointServiceWithParams(ctx, params).Execute()
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

func getPrivateEndpointServiceBuilder() *cobra.Command {
	opts := getPrivateEndpointServiceOpts{}
	cmd := &cobra.Command{
		Use:   "getPrivateEndpointService",
		Short: "Return One Private Endpoint Service for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getRegionalizedPrivateEndpointSettingOpts struct {
	client  *admin.APIClient
	groupId string
	format  string
	tmpl    *template.Template
}

func (opts *getRegionalizedPrivateEndpointSettingOpts) preRun() (err error) {
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

func (opts *getRegionalizedPrivateEndpointSettingOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.GetRegionalizedPrivateEndpointSettingApiParams{
		GroupId: opts.groupId,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSettingWithParams(ctx, params).Execute()
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

func getRegionalizedPrivateEndpointSettingBuilder() *cobra.Command {
	opts := getRegionalizedPrivateEndpointSettingOpts{}
	cmd := &cobra.Command{
		Use:   "getRegionalizedPrivateEndpointSetting",
		Short: "Return Regionalized Private Endpoint Status",
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

type listPrivateEndpointServicesOpts struct {
	client        *admin.APIClient
	groupId       string
	cloudProvider string
	format        string
	tmpl          *template.Template
}

func (opts *listPrivateEndpointServicesOpts) preRun() (err error) {
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

func (opts *listPrivateEndpointServicesOpts) run(ctx context.Context, _ io.Reader, w io.Writer) error {

	params := &admin.ListPrivateEndpointServicesApiParams{
		GroupId:       opts.groupId,
		CloudProvider: opts.cloudProvider,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.ListPrivateEndpointServicesWithParams(ctx, params).Execute()
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

func listPrivateEndpointServicesBuilder() *cobra.Command {
	opts := listPrivateEndpointServicesOpts{}
	cmd := &cobra.Command{
		Use:   "listPrivateEndpointServices",
		Short: "Return All Private Endpoint Services for One Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)

	_ = cmd.MarkFlagRequired("cloudProvider")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type toggleRegionalizedPrivateEndpointSettingOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) preRun() (err error) {
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

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) readData(r io.Reader) (*admin.ProjectSettingItem, error) {
	var out *admin.ProjectSettingItem

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

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.ToggleRegionalizedPrivateEndpointSettingApiParams{
		GroupId: opts.groupId,

		ProjectSettingItem: data,
	}

	resp, _, err := opts.client.PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSettingWithParams(ctx, params).Execute()
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

func toggleRegionalizedPrivateEndpointSettingBuilder() *cobra.Command {
	opts := toggleRegionalizedPrivateEndpointSettingOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "toggleRegionalizedPrivateEndpointSetting",
		Short: "Toggle Regionalized Private Endpoint Status",
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

func privateEndpointServicesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "privateEndpointServices",
		Short: `Returns, adds, edits, and removes private endpoint services.`,
	}
	cmd.AddCommand(
		createPrivateEndpointBuilder(),
		createPrivateEndpointServiceBuilder(),
		deletePrivateEndpointBuilder(),
		deletePrivateEndpointServiceBuilder(),
		getPrivateEndpointBuilder(),
		getPrivateEndpointServiceBuilder(),
		getRegionalizedPrivateEndpointSettingBuilder(),
		listPrivateEndpointServicesBuilder(),
		toggleRegionalizedPrivateEndpointSettingBuilder(),
	)
	return cmd
}
