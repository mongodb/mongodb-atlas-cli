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

type getAuditingConfigurationOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *getAuditingConfigurationOpts) preRun() (err error) {
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

	return nil
}

func (opts *getAuditingConfigurationOpts) run(ctx context.Context, w io.Writer) error {

	params := &admin.GetAuditingConfigurationApiParams{
		GroupId: opts.groupId,
	}

	resp, _, err := opts.client.AuditingApi.GetAuditingConfigurationWithParams(ctx, params).Execute()
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

func getAuditingConfigurationBuilder() *cobra.Command {
	opts := getAuditingConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "getAuditingConfiguration",
		Short: "Return the Auditing Configuration for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	return cmd
}

type updateAuditingConfigurationOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *updateAuditingConfigurationOpts) preRun() (err error) {
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

	return nil
}

func (opts *updateAuditingConfigurationOpts) readData() (*admin.AuditLog, error) {
	var out *admin.AuditLog

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

func (opts *updateAuditingConfigurationOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}

	params := &admin.UpdateAuditingConfigurationApiParams{
		GroupId: opts.groupId,

		AuditLog: data,
	}

	resp, _, err := opts.client.AuditingApi.UpdateAuditingConfigurationWithParams(ctx, params).Execute()
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

func updateAuditingConfigurationBuilder() *cobra.Command {
	opts := updateAuditingConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateAuditingConfiguration",
		Short: "Update Auditing Configuration for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	return cmd
}

func auditingBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auditing",
		Short: `Returns and edits database auditing settings for MongoDB Cloud projects.`,
	}
	cmd.AddCommand(
		getAuditingConfigurationBuilder(),
		updateAuditingConfigurationBuilder(),
	)
	return cmd
}
