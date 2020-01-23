// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/usage"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type configOpts struct {
	*globalOpts
	Service       string
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
}

func (opts *configOpts) IsCloud() bool {
	return opts.Service == config.CloudService
}

func (opts *configOpts) IsOpsManager() bool {
	return opts.Service == config.OpsManagerService
}

func (opts *configOpts) IsCloudManager() bool {
	return opts.Service == config.CloudManagerService
}

func (opts *configOpts) Save() error {
	opts.SetService(opts.Service)
	if opts.PublicAPIKey != "" {
		opts.SetPublicAPIKey(opts.PublicAPIKey)
	}
	if opts.PrivateAPIKey != "" {
		opts.SetPrivateAPIKey(opts.PrivateAPIKey)
	}
	if opts.IsOpsManager() && opts.OpsManagerURL != "" {
		opts.SetOpsManagerURL(opts.OpsManagerURL)
	}

	return opts.Config.Save()
}

func (opts *configOpts) Run() error {
	helpLink := "https://docs.atlas.mongodb.com/configure-api-access/"

	if opts.IsOpsManager() {
		helpLink = "https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
	}

	var defaultQuestions = []*survey.Question{
		{
			Name: "publicAPIKey",
			Prompt: &survey.Input{
				Message: "Public API Key:",
				Help:    helpLink,
				Default: opts.Config.PublicAPIKey(),
			},
		},
		{
			Name: "privateAPIKey",
			Prompt: &survey.Password{
				Message: "Private API Key:",
				Help:    helpLink,
			},
		},
	}

	if opts.IsOpsManager() {
		var opsManagerQuestions = []*survey.Question{
			{
				Name: "opsManagerURL",
				Prompt: &survey.Input{
					Message: "Ops Manager Base URL:",
					Default: opts.Config.OpsManagerURL(),
					Help:    "Ops Manager host URL",
				},
				Validate: validURL,
			},
		}
		defaultQuestions = append(opsManagerQuestions, defaultQuestions...)
	}

	err := survey.Ask(defaultQuestions, opts)
	if err != nil {
		return err
	}

	return opts.Save()
}

func ConfigBuilder() *cobra.Command {
	opts := &configOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the tool.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.globalOpts.loadConfig()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.Service, flags.Service, config.CloudService, usage.Service)
	cmd.Flags().StringVarP(&opts.profile, flags.Profile, flags.ProfileShort, config.DefaultProfile, usage.ProfileConfig)
	cmd.AddCommand(ConfigSetBuilder())

	return cmd
}
