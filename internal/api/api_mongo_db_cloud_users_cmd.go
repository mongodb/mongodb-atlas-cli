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
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type createUserOpts struct {
	client *admin.APIClient

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.CloudAppUser
}

func (opts *createUserOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *createUserOpts) readData(r io.Reader) (*admin.CloudAppUser, error) {
	var out *admin.CloudAppUser

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

func (opts *createUserOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateUserApiParams{

		CloudAppUser: data,
	}

	var err error
	opts.resp, _, err = opts.client.MongoDBCloudUsersApi.CreateUserWithParams(ctx, params).Execute()
	return err
}

func (opts *createUserOpts) postRun(_ context.Context, w io.Writer) error {

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

func createUserBuilder() *cobra.Command {
	opts := createUserOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createUser",
		Short: "Create One MongoDB Cloud User",
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

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getUserOpts struct {
	client *admin.APIClient
	userId string
	format string
	tmpl   *template.Template
	resp   *admin.CloudAppUser
}

func (opts *getUserOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getUserOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetUserApiParams{
		UserId: opts.userId,
	}

	var err error
	opts.resp, _, err = opts.client.MongoDBCloudUsersApi.GetUserWithParams(ctx, params).Execute()
	return err
}

func (opts *getUserOpts) postRun(_ context.Context, w io.Writer) error {

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

func getUserBuilder() *cobra.Command {
	opts := getUserOpts{}
	cmd := &cobra.Command{
		Use:   "getUser",
		Short: "Return One MongoDB Cloud User using Its ID",
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
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies this user.`)

	_ = cmd.MarkFlagRequired("userId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getUserByUsernameOpts struct {
	client   *admin.APIClient
	userName string
	format   string
	tmpl     *template.Template
	resp     *admin.CloudAppUser
}

func (opts *getUserByUsernameOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getUserByUsernameOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetUserByUsernameApiParams{
		UserName: opts.userName,
	}

	var err error
	opts.resp, _, err = opts.client.MongoDBCloudUsersApi.GetUserByUsernameWithParams(ctx, params).Execute()
	return err
}

func (opts *getUserByUsernameOpts) postRun(_ context.Context, w io.Writer) error {

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

func getUserByUsernameBuilder() *cobra.Command {
	opts := getUserByUsernameOpts{}
	cmd := &cobra.Command{
		Use:   "getUserByUsername",
		Short: "Return One MongoDB Cloud User using Their Username",
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
	cmd.Flags().StringVar(&opts.userName, "userName", "", `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`)

	_ = cmd.MarkFlagRequired("userName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
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
