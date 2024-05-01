// Copyright 2021 MongoDB Inc
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

package setup

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
)

func newClusterNameQuestion(clusterName string) *survey.Question {
	return &survey.Question{
		Name: "clusterName",
		Prompt: &survey.Input{
			Message: "Cluster Name [This can't be changed later]",
			Help:    usage.ClusterName,
			Default: clusterName,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.ClusterName),
	}
}

func newClusterProviderQuestion() *survey.Question {
	return &survey.Question{
		Name: "provider",
		Prompt: &survey.Select{
			Message: "Cloud Provider",
			Help:    usage.Provider,
			Options: []string{"AWS", "GCP", "AZURE"},
		},
		Validate: survey.Required,
	}
}

func newAccessListQuestion(publicIP, message string) survey.Prompt {
	extraInfo := "  Set to 0.0.0.0/0 if you want to enable connection from anywhere; use comma (,) to separate multiple entries."
	return &survey.Input{
		Message: "IP Access List Entry" + message,
		Help:    usage.NetworkAccessListIPEntry + extraInfo,
		Default: publicIP,
	}
}

func newDBUsernameQuestion(dbUser string, validation survey.Validator) *survey.Question {
	q := &survey.Question{
		Name: "dbUsername",
		Prompt: &survey.Input{
			Message: "Database User Username",
			Help:    usage.DBUsername,
			Default: dbUser,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.DBUsername, validation),
	}
	return q
}

func newDBUserPasswordQuestion(password, message string) *survey.Question {
	return &survey.Question{
		Name: "DBUserPassword",
		Prompt: &survey.Input{
			Message: "Database User Password" + message,
			Help:    usage.Password,
			Default: password,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.WeakPassword),
	}
}

func newSampleDataQuestion() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you want to load a sample dataset?",
		Help:    "Load a sample dataset to help you test features in your cluster. To learn more, see: https://dochub.mongodb.org/core/available-sample-datasets-atlas.",
		Default: true,
	}
}

func newClusterCreateConfirm() survey.Prompt {
	return &survey.Confirm{
		Message: "Are you ready to create your Atlas cluster with the above settings?",
		Default: true,
	}
}

func newClusterDefaultConfirm(tier string) survey.Prompt {
	message := "Do you want to set up your Atlas database with default settings?"
	if tier != DefaultAtlasTier {
		message = "Are you ready to create your Atlas cluster with the above settings?"
	}

	return &survey.Confirm{
		Message: message,
		Default: true,
	}
}
