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
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createCostExplorerQueryProcessOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *createCostExplorerQueryProcessOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *createCostExplorerQueryProcessOpts) readData() (*admin.CostExplorerFilterRequestBody, error) {
	var out *admin.CostExplorerFilterRequestBody

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

func (opts *createCostExplorerQueryProcessOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.CreateCostExplorerQueryProcessApiParams{
		OrgId: opts.orgId,

		CostExplorerFilterRequestBody: data,
	}

	resp, _, err := opts.client.InvoicesApi.CreateCostExplorerQueryProcessWithParams(ctx, params).Execute()
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

func createCostExplorerQueryProcessBuilder() *cobra.Command {
	opts := createCostExplorerQueryProcessOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createCostExplorerQueryProcess",
		Short: "Create Cost Explorer query process",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	return cmd
}

type createCostExplorerQueryProcess1Opts struct {
	client *admin.APIClient
	orgId  string
	token  string
}

func (opts *createCostExplorerQueryProcess1Opts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *createCostExplorerQueryProcess1Opts) run(ctx context.Context, w io.Writer) error {

	params := &admin.CreateCostExplorerQueryProcess1ApiParams{
		OrgId: opts.orgId,
		Token: opts.token,
	}

	resp, _, err := opts.client.InvoicesApi.CreateCostExplorerQueryProcess1WithParams(ctx, params).Execute()
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

func createCostExplorerQueryProcess1Builder() *cobra.Command {
	opts := createCostExplorerQueryProcess1Opts{}
	cmd := &cobra.Command{
		Use:   "createCostExplorerQueryProcess1",
		Short: "Return results from a given Cost Explorer query, or notify that the results are not ready yet.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.token, "token", "", `Unique 64 digit string that identifies the Cost Explorer query.`)

	_ = cmd.MarkFlagRequired("token")
	return cmd
}

type downloadInvoiceCSVOpts struct {
	client    *admin.APIClient
	orgId     string
	invoiceId string
}

func (opts *downloadInvoiceCSVOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *downloadInvoiceCSVOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.DownloadInvoiceCSVApiParams{
		OrgId:     opts.orgId,
		InvoiceId: opts.invoiceId,
	}

	resp, _, err := opts.client.InvoicesApi.DownloadInvoiceCSVWithParams(ctx, params).Execute()
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

func downloadInvoiceCSVBuilder() *cobra.Command {
	opts := downloadInvoiceCSVOpts{}
	cmd := &cobra.Command{
		Use:   "downloadInvoiceCSV",
		Short: "Return One Organization Invoice as CSV",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.invoiceId, "invoiceId", "", `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`)

	_ = cmd.MarkFlagRequired("invoiceId")
	return cmd
}

type getInvoiceOpts struct {
	client    *admin.APIClient
	orgId     string
	invoiceId string
}

func (opts *getInvoiceOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *getInvoiceOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetInvoiceApiParams{
		OrgId:     opts.orgId,
		InvoiceId: opts.invoiceId,
	}

	resp, _, err := opts.client.InvoicesApi.GetInvoiceWithParams(ctx, params).Execute()
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

func getInvoiceBuilder() *cobra.Command {
	opts := getInvoiceOpts{}
	cmd := &cobra.Command{
		Use:   "getInvoice",
		Short: "Return One Organization Invoice",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.invoiceId, "invoiceId", "", `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`)

	_ = cmd.MarkFlagRequired("invoiceId")
	return cmd
}

type listInvoicesOpts struct {
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listInvoicesOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *listInvoicesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListInvoicesApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.InvoicesApi.ListInvoicesWithParams(ctx, params).Execute()
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

func listInvoicesBuilder() *cobra.Command {
	opts := listInvoicesOpts{}
	cmd := &cobra.Command{
		Use:   "listInvoices",
		Short: "Return All Invoices for One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	return cmd
}

type listPendingInvoicesOpts struct {
	client *admin.APIClient
	orgId  string
}

func (opts *listPendingInvoicesOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *listPendingInvoicesOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.ListPendingInvoicesApiParams{
		OrgId: opts.orgId,
	}

	resp, _, err := opts.client.InvoicesApi.ListPendingInvoicesWithParams(ctx, params).Execute()
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

func listPendingInvoicesBuilder() *cobra.Command {
	opts := listPendingInvoicesOpts{}
	cmd := &cobra.Command{
		Use:   "listPendingInvoices",
		Short: "Return All Pending Invoices for One Organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	return cmd
}

func invoicesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoices",
		Short: `Returns invoices.`,
	}
	cmd.AddCommand(
		createCostExplorerQueryProcessBuilder(),
		createCostExplorerQueryProcess1Builder(),
		downloadInvoiceCSVBuilder(),
		getInvoiceBuilder(),
		listInvoicesBuilder(),
		listPendingInvoicesBuilder(),
	)
	return cmd
}
