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
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/validate"
)

func newOMURLInput() survey.Prompt {
	return &survey.Input{
		Message: "URL to Access Ops Manager:",
		Help:    "FQDN and port number of the Ops Manager Application.",
		Default: config.OpsManagerURL(),
	}
}

func newOrgIDInput() survey.Prompt {
	return &survey.Input{
		Message: "Default Org ID:",
		Help:    "ID of an existing organization that your API keys have access to. If you don't enter an ID, you must use --orgId for every command that requires it.",
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
		Message: "Default Project ID:",
		Help:    "ID of an existing project that your API keys have access to. If you don't enter an ID, you must use --projectId for every command that requires it.",
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
	helpLink := "Please provide your API keys. To create new keys, see the documentation: https://docs.atlas.mongodb.com/configure-api-access/"
	if isOM {
		helpLink = "Please provide your API keys. To create new keys, see the documentation: https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
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

const (
	plaintextFormat = "plaintext"
	jsonFormat      = "json"
)

func defaultQuestions(isAtlas bool) []*survey.Question {
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
	if isAtlas {
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

func tenantQuestions() []*survey.Question {
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
