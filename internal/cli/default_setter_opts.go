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
	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_default_opts.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli ProjectOrgsLister

type ProjectOrgsLister interface {
	Project(id string) (*atlasv2.Group, error)
	Projects(*store.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
	Organization(id string) (*atlasv2.AtlasOrganization, error)
	Organizations(*atlasv2.ListOrganizationsApiParams) (*atlasv2.PaginatedOrganization, error)
	GetOrgProjects(string, *store.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
}

type DefaultSetterOpts struct {
	Service                  string
	OpsManagerURL            string
	ProjectID                string
	OrgID                    string
	TelemetryEnabled         bool
	Output                   string
	Store                    ProjectOrgsLister
	OutWriter                io.Writer
	AskedOrgsOrProjects      bool
	OnMultipleOrgsOrProjects func()
}

func (opts *DefaultSetterOpts) InitStore(ctx context.Context) error {
	if opts.Store != nil {
		return nil
	}

	var err error
	opts.Store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
	return err
}

const resultsLimit = 500

var (
	errTooManyResults = errors.New("too many results")
	errNoResults      = errors.New("no results")
)

func newSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], speed)
}

// Projects fetches projects and returns then as two slices of string.
// One for names and another for ids.
func (opts *DefaultSetterOpts) projects() (ids, names []string, err error) {
	spin := newSpinner()
	spin.Start()
	defer spin.Stop()

	var projects *atlasv2.PaginatedAtlasGroup
	if opts.OrgID == "" {
		projects, err = opts.Store.Projects(&store.ListOptions{ItemsPerPage: resultsLimit})
	} else {
		list := &store.ListOptions{ItemsPerPage: resultsLimit}
		projects, err = opts.Store.GetOrgProjects(opts.OrgID, list)
	}
	if err != nil {
		err = commonerrors.Check(err)
		var atlasErr *atlas.ErrorResponse
		if errors.As(err, &atlasErr) && atlasErr.HTTPCode == 404 {
			return nil, nil, errNoResults
		}
		return nil, nil, err
	}
	if projects.GetTotalCount() == 0 {
		return nil, nil, errNoResults
	}
	if projects.GetTotalCount() > resultsLimit {
		return nil, nil, errTooManyResults
	}
	ids, names = atlasProjects(projects.GetResults())

	return ids, names, nil
}

// Orgs fetches organizations, filtering by name.
func (opts *DefaultSetterOpts) orgs(filter string) (results []atlasv2.AtlasOrganization, err error) {
	spin := newSpinner()
	spin.Start()
	defer spin.Stop()
	pagination := &atlasv2.ListOrganizationsApiParams{Name: &filter, ItemsPerPage: pointer.Get(resultsLimit)}
	orgs, err := opts.Store.Organizations(pagination)
	if err != nil {
		var atlasErr *atlas.ErrorResponse
		if errors.As(err, &atlasErr) && atlasErr.HTTPCode == 404 {
			return nil, errNoResults
		}
		return nil, commonerrors.Check(err)
	}
	if orgs == nil {
		return nil, errNoResults
	}

	if orgs.GetTotalCount() == 0 {
		return nil, errNoResults
	}
	if orgs.GetTotalCount() > resultsLimit {
		return nil, errTooManyResults
	}
	return orgs.GetResults(), nil
}

// ProjectExists checks if the project exists and the current user has access to it.
func (opts *DefaultSetterOpts) ProjectExists(id string) bool {
	if p, err := opts.Store.Project(id); p == nil || err != nil {
		return false
	}
	return true
}

// AskProject will try to construct a select based on fetched projects.
// If it fails or there are no projects to show we fall back to ask for project by ID.
// If only one project, select it by default without prompting the user.
func (opts *DefaultSetterOpts) AskProject() error {
	ids, names, err := opts.projects()
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
		if err2 := telemetry.TrackAskOne(p, &manually); err2 != nil {
			return err2
		}
		opts.AskedOrgsOrProjects = true
		if manually {
			p := prompt.NewProjectIDInput()
			return telemetry.TrackAskOne(p, &opts.ProjectID, survey.WithValidator(validate.OptionalObjectID))
		}
		_, _ = fmt.Fprint(opts.OutWriter, "Skipping default project setting\n")
		return nil
	}

	if len(ids) == 1 {
		opts.ProjectID = ids[0]
	} else {
		opts.runOnMultipleOrgsOrProjects()
		p := prompt.NewProjectSelect(ids, names)
		var projectID string
		if err := telemetry.TrackAskOne(p, &projectID); err != nil {
			return err
		}
		opts.ProjectID = projectID
		opts.AskedOrgsOrProjects = true
	}

	return nil
}

// OrgExists checks if the org exists and the current user has access to it.
func (opts *DefaultSetterOpts) OrgExists(id string) bool {
	if o, err := opts.Store.Organization(id); o == nil || err != nil {
		return false
	}
	return true
}

// AskOrg will try to construct a select based on fetched organizations.
// If it fails or there are no organizations to show we fall back to ask for org by ID.
// If only one organization, select it by default without prompting the user.
func (opts *DefaultSetterOpts) AskOrg() error {
	return opts.askOrgWithFilter("")
}

func (opts *DefaultSetterOpts) askOrgWithFilter(filter string) error {
	orgs, err := opts.orgs(filter)
	if err != nil {
		applyFilter := false
		var target *atlas.ErrorResponse
		switch {
		case errors.Is(err, errNoResults):
			if filter == "" {
				_, _ = fmt.Fprintln(opts.OutWriter, "You don't seem to have access to any organization")
			} else {
				_, _ = fmt.Fprintln(opts.OutWriter, "No results match, please type the organization ID or the organization name to filter.")
				applyFilter = true
			}
		case errors.Is(err, errTooManyResults):
			_, _ = fmt.Fprintln(opts.OutWriter, "Now set your default organization and project.")
			_, _ = fmt.Fprintf(opts.OutWriter, "Since you have access to more than %d organizations, please type the organization ID or the organization name to filter.\n", resultsLimit)
			applyFilter = true
		case errors.As(err, &target):
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your organizations: %s\n", target.Detail)
		default:
			_, _ = fmt.Fprintf(opts.OutWriter, "There was an error fetching your organizations: %s\n", err)
		}

		if applyFilter {
			filterPrompt := &survey.Input{
				Message: "Organization filter:",
				Help:    "Enter the 24 digit ID or type from the beginning of the name to filter.",
			}
			filterErr := telemetry.TrackAskOne(filterPrompt, &filter)
			if filterErr != nil {
				return filterErr
			}
			if filter != "" {
				if validate.ObjectID(filter) == nil {
					opts.OrgID = filter
					_, _ = fmt.Fprintf(opts.OutWriter, "Chosen default organization: %v\n", opts.OrgID)
					return nil
				}
				return opts.askOrgWithFilter(filter)
			}
		}

		return opts.manualOrgID()
	}

	return opts.selectOrg(orgs)
}

func (opts *DefaultSetterOpts) manualOrgID() error {
	p := &survey.Confirm{
		Message: "Do you want to enter the Organization ID manually?",
	}
	manually := true
	if err := telemetry.TrackAskOne(p, &manually); err != nil {
		return err
	}
	opts.AskedOrgsOrProjects = true
	if manually {
		p := prompt.NewOrgIDInput()
		return telemetry.TrackAskOne(p, &opts.OrgID, survey.WithValidator(validate.OptionalObjectID))
	}
	_, _ = fmt.Fprint(opts.OutWriter, "Skipping default organization setting\n")
	return nil
}

func (opts *DefaultSetterOpts) selectOrg(orgs []atlasv2.AtlasOrganization) error {
	if len(orgs) == 1 {
		opts.OrgID = orgs[0].GetId()
		return nil
	}

	opts.runOnMultipleOrgsOrProjects()

	p := prompt.NewOrgSelect(orgs)
	if err := telemetry.TrackAskOne(p, &opts.OrgID); err != nil {
		return err
	}
	opts.AskedOrgsOrProjects = true

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

func (opts *DefaultSetterOpts) SetUpOutput() {
	config.SetOutput(opts.Output)
}

// atlasProjects transform []*atlas.Project to a []string of ids and another for names.
func atlasProjects(projects []atlasv2.Group) (ids, names []string) {
	names = make([]string, len(projects))
	ids = make([]string, len(projects))
	for i, p := range projects {
		ids[i] = p.GetId()
		names[i] = p.GetName()
	}
	return ids, names
}

func (*DefaultSetterOpts) DefaultQuestions() []*survey.Question {
	defaultOutput := config.Output()
	if defaultOutput == "" {
		defaultOutput = plaintextFormat
	}
	q := []*survey.Question{
		{
			Name: "output",
			Prompt: &survey.Select{
				Message: "Default Output Format:",
				Options: []string{plaintextFormat, jsonFormat},
				Default: defaultOutput,
			},
		},
	}
	return q
}

func (opts *DefaultSetterOpts) runOnMultipleOrgsOrProjects() {
	if opts.OnMultipleOrgsOrProjects != nil {
		opts.OnMultipleOrgsOrProjects()
	}
}
