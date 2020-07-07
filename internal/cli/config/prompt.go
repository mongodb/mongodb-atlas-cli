package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
)

const (
	omBaseURLHelp = "FQDN and port number of the Ops Manager Application."
	projectHelp   = "ID of an existing project that your API keys have access to. If you don't enter an ID, you must use --projectId for every command that requires it."
	orgHelp       = "ID of an existing organization that your API keys have access to. If you don't enter an ID, you must use --orgId for every command that requires it."
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
		Message: "Default Org ID:",
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
		Message: "Default Project ID:",
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
