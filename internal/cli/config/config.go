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
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	nameIDFormat = "%s (%s)"
)

type ProjectOrgsLister interface {
	Projects(*atlas.ListOptions) (interface{}, error)
	Organizations(*atlas.OrganizationsListOptions) (*atlas.Organizations, error)
}

type configOpts struct {
	Service        string
	PublicAPIKey   string
	PrivateAPIKey  string
	OpsManagerURL  string
	ProjectID      string
	OrgID          string
	MongoShellPath string
	Output         string
	store          ProjectOrgsLister
}

func (opts *configOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.PublicAuthenticatedPreset(config.Default()))
	return err
}

func (opts *configOpts) IsCloud() bool {
	return opts.Service == config.CloudService
}

func (opts *configOpts) IsOpsManager() bool {
	return opts.Service == config.OpsManagerService
}

func (opts *configOpts) SetUpAccess() {
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
}

func (opts *configOpts) SetUpProject() {
	if opts.ProjectID != "" {
		config.SetProjectID(opts.ProjectID)
	}
}

func (opts *configOpts) SetUpOrg() {
	if opts.OrgID != "" {
		config.SetOrgID(opts.OrgID)
	}
}

func (opts *configOpts) setUpMongoSHPath() {
	if opts.MongoShellPath != "" {
		config.SetMongoShellPath(opts.MongoShellPath)
	}
}

func (opts *configOpts) setUpOutput() {
	if opts.Output != plaintextFormat {
		config.SetOutput(opts.Output)
	}
}

func (opts *configOpts) Run() error {
	fmt.Printf(`You are configuring a profile for %s.

All values are optional and you can use environment variables (MCLI_*) instead.

Enter [?] on any option to get help.

`, config.ToolName)

	q := accessQuestions(opts.IsOpsManager())
	if err := survey.Ask(q, opts); err != nil {
		return err
	}
	opts.SetUpAccess()

	if err := opts.initStore(); err != nil {
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

	q2 := defaultQuestions(opts.IsCloud())
	if err := survey.Ask(q2, opts); err != nil {
		return err
	}
	opts.setUpOutput()
	opts.setUpMongoSHPath()

	if err := config.Save(); err != nil {
		return err
	}

	fmt.Printf("\nYour profile is now configured.\n")
	if config.Name() != config.DefaultProfile {
		fmt.Printf("To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}
	fmt.Printf("You can use [%s config set] to change these settings at a later time.\n", config.ToolName)
	return nil
}

// askProject will try to construct a select based on fetched projects.
// If it fails or there are no projects to show we fallback to ask for project by ID.
func (opts *configOpts) askProject() error {
	pMap, pSlice, err := opts.projects()
	var projectID string
	if err != nil || len(pSlice) == 0 {
		prompt := newProjectIDInput()
		return survey.AskOne(prompt, &opts.ProjectID, survey.WithValidator(validate.OptionalObjectID))
	}

	prompt := newProjectSelect(pSlice)
	if err := survey.AskOne(prompt, &projectID); err != nil {
		return err
	}
	opts.ProjectID = pMap[projectID]
	return nil
}

// projects fetches projects and returns then as a slice of the format `nameIDFormat`,
// and a map such as `map[nameIDFormat]=ID`.
// This is necessary as we can only prompt using `nameIDFormat`
// and we want them to get the ID mapping to store on the config.
func (opts *configOpts) projects() (pMap map[string]string, pSlice []string, err error) {
	projects, err := opts.store.Projects(nil)
	if err != nil {
		fmt.Printf("there was a problem fetching projects: %s\n", err)
		return nil, nil, err
	}
	if opts.IsCloud() {
		pMap, pSlice = atlasProjects(projects.(*atlas.Projects).Results)
	} else {
		pMap, pSlice = omProjects(projects.(*opsmngr.Projects).Results)
	}
	return pMap, pSlice, nil
}

// askOrg will try to construct a select based on fetched organizations.
// If it fails or there are no organizations to show we fallback to ask for org by ID.
func (opts *configOpts) askOrg() error {
	oMap, oSlice, err := opts.orgs()
	var orgID string
	if err != nil || len(oSlice) == 0 {
		prompt := newOrgIDInput()
		return survey.AskOne(prompt, &opts.OrgID, survey.WithValidator(validate.OptionalObjectID))
	}

	prompt := newOrgSelect(oSlice)
	if err := survey.AskOne(prompt, &orgID); err != nil {
		return err
	}
	opts.OrgID = oMap[orgID]
	return nil
}

// orgs fetches organizations and returns then as a slice of the format `nameIDFormat`,
// and a map such as `map[nameIDFormat]=ID`.
// This is necessary as we can only prompt using `nameIDFormat`
// and we want them to get the ID mapping to store on the config.
func (opts *configOpts) orgs() (oMap map[string]string, oSlice []string, err error) {
	orgs, err := opts.store.Organizations(nil)
	if err != nil {
		fmt.Printf("there was a problem fetching orgs: %s\n", err)
		return nil, nil, err
	}
	oMap = make(map[string]string, len(orgs.Results))
	oSlice = make([]string, len(orgs.Results))
	for i, o := range orgs.Results {
		d := fmt.Sprintf(nameIDFormat, o.Name, o.ID)
		oMap[d] = o.ID
		oSlice[i] = d
	}
	return oMap, oSlice, nil
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
  
  To configure the tool to work with Cloud Manager
  $ mongocli config --service cloud-manager

  To configure the tool to work with Ops Manager
  $ mongocli config --service ops-manager
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
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

// atlasProjects transform []*atlas.Project to a map[string]string and []string.
func atlasProjects(projects []*atlas.Project) (pMap map[string]string, pSlice []string) {
	pMap = make(map[string]string, len(projects))
	pSlice = make([]string, len(projects))
	for i, p := range projects {
		d := fmt.Sprintf(nameIDFormat, p.Name, p.ID)
		pMap[d] = p.ID
		pSlice[i] = d
	}
	return pMap, pSlice
}

// omProjects transform []*opsmngr.Project to a map[string]string and []string.
func omProjects(projects []*opsmngr.Project) (pMap map[string]string, pSlice []string) {
	pMap = make(map[string]string, len(projects))
	pSlice = make([]string, len(projects))
	for i, p := range projects {
		d := fmt.Sprintf(nameIDFormat, p.Name, p.ID)
		pMap[d] = p.ID
		pSlice[i] = d
	}
	return pMap, pSlice
}
