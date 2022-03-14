// Copyright 2022 MongoDB Inc
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

package cli

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/prompt"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/validate"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_default_opts.go -package=mocks github.com/mongodb/mongocli/internal/cli ProjectOrgsLister

type ProjectOrgsLister interface {
	Projects(*atlas.ListOptions) (interface{}, error)
	Organizations(*atlas.OrganizationsListOptions) (*atlas.Organizations, error)
	GetOrgProjects(string, *atlas.ListOptions) (interface{}, error)
}

type DefaultSetterOpts struct {
	Service        string
	OpsManagerURL  string
	ProjectID      string
	OrgID          string
	MongoShellPath string
	Output         string
	Store          ProjectOrgsLister
	OutWriter      io.Writer
}

func (opts *DefaultSetterOpts) InitStore(ctx context.Context) error {
	var err error
	opts.Store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
	return err
}

func (opts *DefaultSetterOpts) IsCloud() bool {
	return opts.Service == config.CloudService || opts.Service == config.CloudGovService
}

func (opts *DefaultSetterOpts) IsOpsManager() bool {
	return opts.Service == config.OpsManagerService
}

const resultsLimit = 500

var (
	errTooManyResults = errors.New("too many results")
	errNoResults      = errors.New("no results")
)

// Projects fetches projects and returns then as a slice of the format `nameIDFormat`,
// and a map such as `map[nameIDFormat]=ID`.
// This is necessary as we can only prompt using `nameIDFormat`
// and we want them to get the ID mapping to store in the config.
func (opts *DefaultSetterOpts) projects() (pMap map[string]string, pSlice []string, err error) {
	var projects interface{}
	if opts.OrgID == "" {
		projects, err = opts.Store.Projects(nil)
	} else {
		projects, err = opts.Store.GetOrgProjects(opts.OrgID, &atlas.ListOptions{ItemsPerPage: resultsLimit})
	}
	if err != nil {
		return nil, nil, err
	}
	switch r := projects.(type) {
	case *atlas.Projects:
		if r.TotalCount == 0 {
			return nil, nil, errNoResults
		}
		if r.TotalCount > resultsLimit {
			return nil, nil, errTooManyResults
		}
		pMap, pSlice = atlasProjects(r.Results)
	case *opsmngr.Projects:
		if r.TotalCount == 0 {
			return nil, nil, errNoResults
		}
		if r.TotalCount > resultsLimit {
			return nil, nil, errTooManyResults
		}
		pMap, pSlice = omProjects(r.Results)
	}

	return pMap, pSlice, nil
}

// Orgs fetches organizations and returns then as a slice of the format `nameIDFormat`,
// and a map such as `map[nameIDFormat]=ID`.
// This is necessary as we can only prompt using `nameIDFormat`
// and we want them to get the ID mapping to store on the config.
func (opts *DefaultSetterOpts) orgs() (oMap map[string]string, oSlice []string, err error) {
	includeDeleted := false
	pagination := &atlas.OrganizationsListOptions{IncludeDeletedOrgs: &includeDeleted}
	pagination.ItemsPerPage = resultsLimit
	orgs, err := opts.Store.Organizations(pagination)
	if err != nil {
		return nil, nil, err
	}
	if orgs.TotalCount == 0 {
		return nil, nil, errNoResults
	}
	if orgs.TotalCount > resultsLimit {
		return nil, nil, errTooManyResults
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

// AskProject will try to construct a select based on fetched projects.
// If it fails or there are no projects to show we fallback to ask for project by ID.
func (opts *DefaultSetterOpts) AskProject() error {
	pMap, pSlice, err := opts.projects()
	if err != nil {
		var target *atlas.ErrorResponse
		switch {
		case errors.Is(err, errNoResults):
			_, _ = fmt.Fprintln(opts.OutWriter, "You don't seem to have access to any project")
		case errors.Is(err, errTooManyResults):
			_, _ = fmt.Fprintf(opts.OutWriter, "You have access to more than %d projects\n", resultsLimit)
		case errors.As(err, &target):
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your projects: %s\n", target.Detail)
		default:
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your projects: %s\n", err)
		}
		p := &survey.Confirm{
			Message: "Do you want to enter the Project ID manually?",
		}
		manually := true
		if err2 := survey.AskOne(p, &manually); err2 != nil {
			return err2
		}
		if manually {
			p := prompt.NewProjectIDInput()
			return survey.AskOne(p, &opts.ProjectID, survey.WithValidator(validate.OptionalObjectID))
		}
		_, _ = fmt.Fprint(opts.OutWriter, "Skipping default project setting\n")
		return nil
	}

	p := prompt.NewProjectSelect(pSlice)
	var projectID string
	if err := survey.AskOne(p, &projectID); err != nil {
		return err
	}
	opts.ProjectID = pMap[projectID]
	return nil
}

// AskOrg will try to construct a select based on fetched organizations.
// If it fails or there are no organizations to show we fallback to ask for org by ID.
func (opts *DefaultSetterOpts) AskOrg() error {
	oMap, oSlice, err := opts.orgs()
	if err != nil {
		var target *atlas.ErrorResponse
		switch {
		case errors.Is(err, errNoResults):
			_, _ = fmt.Fprintln(opts.OutWriter, "You don't seem to have access to any organization")
		case errors.Is(err, errTooManyResults):
			_, _ = fmt.Fprintf(opts.OutWriter, "You have access to more than %d organizations\n", resultsLimit)
		case errors.As(err, &target):
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your organizations: %s\n", target.Detail)
		default:
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your organizations: %s\n", err)
		}
		p := &survey.Confirm{
			Message: "Do you want to enter the Org ID manually?",
		}
		manually := true
		if err2 := survey.AskOne(p, &manually); err2 != nil {
			return err2
		}
		if manually {
			p := prompt.NewOrgIDInput()
			return survey.AskOne(p, &opts.OrgID, survey.WithValidator(validate.OptionalObjectID))
		}
		_, _ = fmt.Fprint(opts.OutWriter, "Skipping default organization setting\n")
		return nil
	}
	p := prompt.NewOrgSelect(oSlice)
	var orgID string
	if err := survey.AskOne(p, &orgID); err != nil {
		return err
	}
	opts.OrgID = oMap[orgID]
	return nil
}

func (opts *DefaultSetterOpts) SetUpProject() {
	if opts.ProjectID != "" {
		config.SetProjectID(opts.ProjectID)
	}
}

func (opts *DefaultSetterOpts) SetUpOrg() {
	if opts.OrgID != "" {
		config.SetOrgID(opts.OrgID)
	}
}

func (opts *DefaultSetterOpts) SetUpMongoSHPath() {
	if opts.MongoShellPath != "" {
		config.SetMongoShellPath(opts.MongoShellPath)
	}
}

func (opts *DefaultSetterOpts) SetUpOutput() {
	if opts.Output != plaintextFormat {
		config.SetOutput(opts.Output)
	}
}

const nameIDFormat = "%s (%s)"

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

func (opts *DefaultSetterOpts) DefaultQuestions() []*survey.Question {
	q := []*survey.Question{
		{
			Name: "output",
			Prompt: &survey.Select{
				Message: "Default Output Format:",
				Options: []string{plaintextFormat, jsonFormat},
				Default: config.Output(),
			},
		},
	}
	if opts.IsCloud() {
		defaultPath := config.MongoShellPath()
		if defaultPath == "" {
			defaultPath = mongosh.Path()
		}
		atlasQuestion := &survey.Question{
			Name: "mongoShellPath",
			Prompt: &survey.Input{
				Message: "Default MongoDB Shell Path:",
				Help:    "MongoDB CLI will use the MongoDB shell version provided to allow you to access your deployments.",
				Default: defaultPath,
			},
			Validate: validate.OptionalPath,
		}
		q = append(q, atlasQuestion)
	}
	return q
}
