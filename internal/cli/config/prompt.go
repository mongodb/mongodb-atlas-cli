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
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
)

const (
	omBaseURLHelp = "FQDN and port number of the Ops Manager Application."
	projectHelp   = "id of an existing project that your API keys have access to. If you don't enter an id, you must use --projectId for every command that requires it."
	orgHelp       = "id of an existing organization that your API keys have access to. If you don't enter an id, you must use --orgId for every command that requires it."
	atlasAPIHelp  = "Please provide your API keys. To create new keys, see the documentation: https://docs.atlas.mongodb.com/configure-api-access/"
	omAPIHelp     = "Please provide your API keys. To create new keys, see the documentation: https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
)

func newOMURLInput() survey.Prompt {
	return &survey.Input{
		Message: "URL to Access Ops Manager:",
		Help:    omBaseURLHelp,
		Default: config.OpsManagerURL(),
	}
}

func newOrgIDInput() survey.Prompt {
	return &survey.Input{
		Message: "Default Org id:",
		Help:    orgHelp,
		Default: config.OrgID(),
	}
}

func newOrgSelect(options []string) survey.Prompt {
	return &survey.Select{
		Message: "Choose a default organization:",
		Options: options,
	}
}

func newProjectIDInput() survey.Prompt {
	return &survey.Input{
		Message: "Default Project id:",
		Help:    projectHelp,
		Default: config.ProjectID(),
	}
}

func newProjectSelect(options []string) survey.Prompt {
	return &survey.Select{
		Message: "Choose a default project:",
		Options: options,
	}
}

func accessQuestions(isOM bool) []*survey.Question {
	helpLink := atlasAPIHelp
	if isOM {
		helpLink = omAPIHelp
	}

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
	if isOM {
		omQuestions := []*survey.Question{
			{
				Name:     "opsManagerURL",
				Prompt:   newOMURLInput(),
				Validate: validate.OptionalURL,
			},
		}
		q = append(omQuestions, q...)
	}
	return q
}

func defaultQuestions() []*survey.Question {
	q := []*survey.Question{
		{
			Name:     "projectId",
			Prompt:   newProjectIDInput(),
			Validate: validate.OptionalObjectID,
		},
		{
			Name:     "orgId",
			Prompt:   newOrgIDInput(),
			Validate: validate.OptionalObjectID,
		},
	}
	return q
}
