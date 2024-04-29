// Copyright 2024 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type createPushBasedLogConfigurationOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPushBasedLogConfigurationOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
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

	return nil
}

func (opts *createPushBasedLogConfigurationOpts) readData(r io.Reader) (*admin.CreatePushBasedLogExportProjectRequest, error) {
	var out *admin.CreatePushBasedLogExportProjectRequest

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

func (opts *createPushBasedLogConfigurationOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreatePushBasedLogConfigurationApiParams{
		GroupId: opts.groupId,

		CreatePushBasedLogExportProjectRequest: data,
	}

	var err error
	_, err = opts.client.PushBasedLogExportApi.CreatePushBasedLogConfigurationWithParams(ctx, params).Execute()
	return err
}

func (opts *createPushBasedLogConfigurationOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func createPushBasedLogConfigurationBuilder() *cobra.Command {
	opts := createPushBasedLogConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPushBasedLogConfiguration",
		Short: "Enable the push-based log export feature for a project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	return cmd
}

type deletePushBasedLogConfigurationOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *deletePushBasedLogConfigurationOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
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

	return nil
}

func (opts *deletePushBasedLogConfigurationOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DeletePushBasedLogConfigurationApiParams{
		GroupId: opts.groupId,
	}

	var err error
	_, err = opts.client.PushBasedLogExportApi.DeletePushBasedLogConfigurationWithParams(ctx, params).Execute()
	return err
}

func (opts *deletePushBasedLogConfigurationOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func deletePushBasedLogConfigurationBuilder() *cobra.Command {
	opts := deletePushBasedLogConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "deletePushBasedLogConfiguration",
		Short: "Disable the push-based log export feature for a project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	return cmd
}

type getPushBasedLogConfigurationOpts struct {
	client  *admin.APIClient
	groupId string
	format  string
	tmpl    *template.Template
	resp    *admin.PushBasedLogExportProject
}

func (opts *getPushBasedLogConfigurationOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getPushBasedLogConfigurationOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetPushBasedLogConfigurationApiParams{
		GroupId: opts.groupId,
	}

	var err error
	opts.resp, _, err = opts.client.PushBasedLogExportApi.GetPushBasedLogConfigurationWithParams(ctx, params).Execute()
	return err
}

func (opts *getPushBasedLogConfigurationOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func getPushBasedLogConfigurationBuilder() *cobra.Command {
	opts := getPushBasedLogConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "getPushBasedLogConfiguration",
		Short: "Get the push-based log export configuration for a project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type updatePushBasedLogConfigurationOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *updatePushBasedLogConfigurationOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
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

	return nil
}

func (opts *updatePushBasedLogConfigurationOpts) readData(r io.Reader) (*admin.PushBasedLogExportProject, error) {
	var out *admin.PushBasedLogExportProject

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

func (opts *updatePushBasedLogConfigurationOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.UpdatePushBasedLogConfigurationApiParams{
		GroupId: opts.groupId,

		PushBasedLogExportProject: data,
	}

	var err error
	_, err = opts.client.PushBasedLogExportApi.UpdatePushBasedLogConfigurationWithParams(ctx, params).Execute()
	return err
}

func (opts *updatePushBasedLogConfigurationOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func updatePushBasedLogConfigurationBuilder() *cobra.Command {
	opts := updatePushBasedLogConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updatePushBasedLogConfiguration",
		Short: "Update the push-based log export feature for a project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	return cmd
}

func pushBasedLogExportBuilder() *cobra.Command {
	const use = "pushBasedLogExport"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `You can continually push logs from mongod, mongos, and audit logs to an AWS S3 bucket. Atlas exports logs every 5 minutes.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		createPushBasedLogConfigurationBuilder(),
		deletePushBasedLogConfigurationBuilder(),
		getPushBasedLogConfigurationBuilder(),
		updatePushBasedLogConfigurationBuilder(),
	)
	return cmd
}
