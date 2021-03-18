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

package quickstart

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/usage"
)

const (
	mongoshHelp         = "MongoDB CLI will use the MongoDB shell version provided to allow you to access your deployments."
	mongoshNotFoundHelp = "MongoDB CLI will store the path in your profile, type ‘mongocli config’ to change it."
)

func newAccessListQuestion(publicIP, message string) *survey.Question {
	return &survey.Question{
		Name: "ipAddress",
		Prompt: &survey.Input{
			Message: message,
			Help:    usage.NetworkAccessListIPEntry,
			Default: publicIP,
		},
	}
}

func newRegionQuestions(region, provider string) *survey.Question {
	if region != "" {
		return nil
	}
	return &survey.Question{
		Name: "region",
		Prompt: &survey.Select{
			Message: "Cloud Provider Region:",
			Help:    usage.Region,
			Options: DefaultRegions[strings.ToUpper(provider)],
		},
	}
}
func newDBUsernameQuestion(dbUser, message string, validation func(val interface{}) error) *survey.Question {
	q := &survey.Question{
		Validate: validation,
		Name:     "dbUsername",
		Prompt: &survey.Input{
			Message: message,
			Help:    usage.DBUsername,
			Default: dbUser,
		},
	}
	return q
}

func newDBUserPasswordQuestion(password, message string) *survey.Question {
	return &survey.Question{
		Name: "DBUserPassword",
		Prompt: &survey.Input{
			Message: message,
			Help:    usage.Password,
			Default: password,
		},
	}
}

func newClusterNameQuestion(clusterName, message string) *survey.Question {
	return &survey.Question{
		Name: "clusterName",
		Prompt: &survey.Input{
			Message: message,
			Help:    usage.ClusterName,
			Default: clusterName,
		},
	}
}

func newClusterProviderQuestion() *survey.Question {
	return &survey.Question{
		Name: "provider",
		Prompt: &survey.Select{
			Message: "Cloud Provider:",
			Help:    usage.Provider,
			Options: []string{"AWS", "GCP", "AZURE"},
		},
	}
}

func newMongoShellQuestionAccessDeployment(clusterName string) *survey.Confirm {
	return &survey.Confirm{
		Message: fmt.Sprintf("Do you want to access %s with MongoDB Shell?", clusterName),
		Help:    mongoshHelp,
	}
}

func newMongoShellQuestionProvidePath() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want to provide the path to your MongoDB shell binary?",
		Help:    mongoshNotFoundHelp,
	}
}

func newMongoShellQuestion() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you have a MongoDB shell version installed on your machine?",
	}
}

func newMongoShellPathInput(defaultValue string, validation func(val interface{}) error) *survey.Question {
	return &survey.Question{
		Validate: validation,
		Name:     "mongoShellPath",
		Prompt: &survey.Input{
			Message: "Default MongoDB Shell Path:",
			Help:    mongoshHelp,
			Default: defaultValue,
		},
	}
}

func newMongoShellQuestionOpenBrowser() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want to download MongoDB Shell [This will open www.mongodb.com on your browser]?",
	}
}
