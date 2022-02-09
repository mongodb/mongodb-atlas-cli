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
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/prompt"
	"github.com/mongodb/mongocli/internal/toolname"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

type configOpts struct {
	cli.DefaultSetterOpts
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
}

func (opts *configOpts) SetUpAccess() {
	config.SetService(opts.Service)
	if opts.PublicAPIKey != "" {
		config.SetPublicAPIKey(opts.PublicAPIKey)
	}
	if opts.PrivateAPIKey != "" {
		config.SetPrivateAPIKey(opts.PrivateAPIKey)
	}
	if opts.OpsManagerURL != "" {
		config.SetOpsManagerURL(opts.OpsManagerURL)
	}
}

func (opts *configOpts) Run(ctx context.Context) error {
	fmt.Printf(`You are configuring a profile for %s.

All values are optional and you can use environment variables (MCLI_*) instead.

Enter [?] on any option to get help.

`, toolname.ToolName)

	q := accessQuestions(opts.IsOpsManager())
	if err := survey.Ask(q, opts); err != nil {
		return err
	}
	opts.SetUpAccess()

	if err := opts.InitStore(ctx); err != nil {
		return err
	}

	if config.IsAccessSet() {
		if err := opts.askOrg(); err != nil {
			return err
		}
		if err := opts.askProject(); err != nil {
			return err
		}
	} else {
		q = tenantQuestions()
		if err := survey.Ask(q, opts); err != nil {
			return err
		}
	}
	opts.SetUpProject()
	opts.SetUpOrg()

	if err := survey.Ask(opts.DefaultQuestions(), opts); err != nil {
		return err
	}
	opts.SetUpOutput()
	opts.SetUpMongoSHPath()

	if err := config.Save(); err != nil {
		return err
	}

	fmt.Printf("\nYour profile is now configured.\n")
	if config.Name() != config.DefaultProfile {
		fmt.Printf("To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}
	fmt.Printf("You can use [%s config set] to change these settings at a later time.\n", toolname.ToolName)
	return nil
}

// askProject will try to construct a select based on fetched projects.
// If it fails or there are no projects to show we fallback to ask for project by ID.
func (opts *configOpts) askProject() error {
	pMap, pSlice, err := opts.Projects()
	var projectID string
	if err != nil || len(pSlice) == 0 {
		p := newProjectIDInput()
		return survey.AskOne(p, &opts.ProjectID, survey.WithValidator(validate.OptionalObjectID))
	}

	p := prompt.NewProjectSelect(pSlice)
	if err := survey.AskOne(p, &projectID); err != nil {
		return err
	}
	opts.ProjectID = pMap[projectID]
	return nil
}

// askOrg will try to construct a select based on fetched organizations.
// If it fails or there are no organizations to show we fallback to ask for org by ID.
func (opts *configOpts) askOrg() error {
	oMap, oSlice, err := opts.Orgs()
	var orgID string
	if err != nil || len(oSlice) == 0 {
		p := newOrgIDInput()
		return survey.AskOne(p, &opts.OrgID, survey.WithValidator(validate.OptionalObjectID))
	}

	p := prompt.NewOrgSelect(oSlice)
	if err := survey.AskOne(p, &orgID); err != nil {
		return err
	}
	opts.OrgID = oMap[orgID]
	return nil
}

func Builder() *cobra.Command {
	opts := &configOpts{}
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure a profile to store access settings for your MongoDB deployment.",
		Long: `Configure settings in a user profile.
All settings are optional. You can specify settings individually by running: 
$ mongocli config set --help 

You can also use environment variables (MCLI_*) when running the tool.
To find out more, see the documentation: https://docs.mongodb.com/mongocli/stable/configure/environment-variables/.`,
		Example: `
  To configure the tool to work with Atlas
  $ mongocli config

  To configure the tool to work with Atlas for Government
  $ mongocli config --service cloudgov
  
  To configure the tool to work with Cloud Manager
  $ mongocli config --service cloud-manager

  To configure the tool to work with Ops Manager
  $ mongocli config --service ops-manager
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		Annotations: map[string]string{
			"toc": "true",
		},
		Args: require.NoArgs,
	}
	cmd.Flags().StringVar(&opts.Service, flag.Service, config.CloudService, usage.Service)
	cmd.AddCommand(
		SetBuilder(),
		ListBuilder(),
		DescribeBuilder(),
		RenameBuilder(),
		DeleteBuilder(),
	)

	return cmd
}
