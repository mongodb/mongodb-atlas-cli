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
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const (
	atlasAPIHelp  = "Please provide your API keys. To create new keys, see the documentation: https://docs.atlas.mongodb.com/configure-api-access/"
	omAPIHelp     = "Please provide your API keys. To create new keys, see the documentation: https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
	omBaseURLHelp = "FQDN and port number of the Ops Manager Application."
	projectHelp   = "ID of an existing project that your API keys have access to. If you don't enter an ID, you must use --projectId for every command that requires it."
	orgHelp       = "ID of an existing organization that your API keys have access to. If you don't enter an ID, you must use --orgId for every command that requires it."
)

type configOpts struct {
	Service       string
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
	ProjectID     string
	OrgID         string
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
	if opts.OrgID != "" {
		config.SetOrgID(opts.OrgID)
	}

	return config.Save()
}

func (opts *configOpts) Run() error {
	fmt.Printf(`You are configuring a profile for %s.

All values are optional and you can use environment variables (MCLI_*) instead.

Enter [?] on any option to get help.

`, config.ToolName)

	q := opts.accessQuestions()
	if err := survey.Ask(q, opts); err != nil {
		return err
	}

	q = opts.defaultQuestions()
	if err := survey.Ask(q, opts); err != nil {
		return err
	}

	if err := opts.Save(); err != nil {
		return err
	}

	fmt.Printf("\nYour profile is now configured.\n")
	if config.Name() != config.DefaultProfile {
		fmt.Printf("To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}
	fmt.Printf("You can use [%s config set] to change these settings at a later time.\n", config.ToolName)
	return nil
}

func (opts *configOpts) apiKeyHelp() string {
	if opts.IsOpsManager() {
		return omAPIHelp
	}
	return atlasAPIHelp
}

func (opts *configOpts) accessQuestions() []*survey.Question {
	helpLink := opts.apiKeyHelp()

	q := []*survey.Question{
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
	}
	if opts.IsOpsManager() {
		opsManagerQuestions := []*survey.Question{
			{
				Name: "opsManagerURL",
				Prompt: &survey.Input{
					Message: "URL to Access Ops Manager:",
					Default: config.OpsManagerURL(),
					Help:    omBaseURLHelp,
				},
				Validate: validate.URL,
			},
		}
		q = append(opsManagerQuestions, q...)
	}
	return q
}

func (opts *configOpts) defaultQuestions() []*survey.Question {
	q := []*survey.Question{
		{
			Name: "projectId",
			Prompt: &survey.Input{
				Message: "Default Project ID:",
				Help:    projectHelp,
				Default: config.ProjectID(),
			},
			Validate: validate.OptionalObjectID,
		},
		{
			Name: "orgId",
			Prompt: &survey.Input{
				Message: "Default Org ID:",
				Help:    orgHelp,
				Default: config.OrgID(),
			},
			Validate: validate.OptionalObjectID,
		},
	}
	return q
}

func Builder() *cobra.Command {
	opts := &configOpts{}
	cmd := &cobra.Command{
		Use:   "config",
		Short: description.ConfigDescription,
		Long:  description.ConfigLongDescription,
		Example: `
  To configure the tool to work with Atlas
  $ mongocli config
  
  To configure the tool to work with Cloud Manager
  $ mongocli config --service cloud-manager

  To configure the tool to work with Ops Manager
  $ mongocli config --service ops-manager
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.Service, flag.Service, config.CloudService, usage.Service)
	cmd.AddCommand(SetBuilder())
	cmd.AddCommand(ListBuilder())
	cmd.AddCommand(DescribeBuilder())
	cmd.AddCommand(DeleteBuilder())

	return cmd
}
