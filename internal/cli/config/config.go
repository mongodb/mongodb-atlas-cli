// Copyright 2020 MongoDB Inc
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

package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

type configOpts struct {
	cli.GlobalOpts
	Service       string
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
	ProjectID     string
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
	config.SetService(opts.Service)
	if opts.PublicAPIKey != "" {
		config.SetPublicAPIKey(opts.PublicAPIKey)
	}
	if opts.PrivateAPIKey != "" {
		config.SetPrivateAPIKey(opts.PrivateAPIKey)
	}
	if opts.IsOpsManager() && opts.OpsManagerURL != "" {
		config.SetOpsManagerURL(opts.OpsManagerURL)
	}
	if opts.ProjectID != "" {
		config.SetProjectID(opts.ProjectID)
	}

	return config.Save()
}

func (opts *configOpts) Run() error {
	helpLink := "https://docs.atlas.mongodb.com/configure-api-access/"

	if opts.IsOpsManager() {
		helpLink = "https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
	}

	defaultQuestions := []*survey.Question{
		{
			Name: "publicAPIKey",
			Prompt: &survey.Input{
				Message: "Public API Key:",
				Help:    helpLink,
				Default: config.PublicAPIKey(),
			},
		},
		{
			Name: "privateAPIKey",
			Prompt: &survey.Password{
				Message: "Private API Key:",
				Help:    helpLink,
			},
		},
		{
			Name: "projectId",
			Prompt: &survey.Input{
				Message: "Default Project ID [optional]:",
				Help:    "This is the ID of an existing project your API keys have access to, you can leave this blank and specify it on every command with --projectId",
				Default: config.ProjectID(),
			},
			Validate: validate.OptionalObjectID,
		},
	}

	if opts.IsOpsManager() {
		opsManagerQuestions := []*survey.Question{
			{
				Name: "opsManagerURL",
				Prompt: &survey.Input{
					Message: "URL to Access Ops Manager:",
					Default: config.OpsManagerURL(),
					Help:    "https://docs.opsmanager.mongodb.com/current/reference/config/ui-settings/#URL-to-Access-Ops-Manager",
				},
				Validate: validate.URL,
			},
		}
		defaultQuestions = append(opsManagerQuestions, defaultQuestions...)
	}

	if err := survey.Ask(defaultQuestions, opts); err != nil {
		return err
	}

	return opts.Save()
}

func Builder() *cobra.Command {
	opts := &configOpts{}
	cmd := &cobra.Command{
		Use:   "config",
		Short: description.ConfigDescription,
		Long:  description.ConfigLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.Service, flag.Service, config.CloudService, usage.Service)
	cmd.AddCommand(SetBuilder())

	return cmd
}
