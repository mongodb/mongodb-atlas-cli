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
	"strconv"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/randgen"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
)

const (
	mongoShellHelp         = "MongoDB CLI will use the MongoDB shell version provided to allow you to access your deployments."
	mongoShellNotFoundHelp = "MongoDB CLI will store the path in your profile, type ‘mongocli config’ to change it."
	SampleDataHelp         = "Add a sample dataset into your new Atlas cluster."
	passwordLength         = 12
)

func newClusterNameQuestion(clusterName, message string) *survey.Question {
	return &survey.Question{
		Name: "clusterName",
		Prompt: &survey.Input{
			Message: fmt.Sprintf("Cluster Name%s:", message),
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
			Message: "Cloud Provider:",
			Help:    usage.Provider,
			Options: []string{"AWS", "GCP", "AZURE"},
		},
		Validate: survey.Required,
	}
}

func newRegionQuestions(defaultRegions []string) *survey.Question {
	return &survey.Question{
		Name: "region",
		Prompt: &survey.Select{
			Message: "Cloud Provider Region:",
			Help:    usage.Region,
			Options: defaultRegions,
		},
		Validate: survey.Required,
	}
}

func newAccessListQuestion(publicIP, message string) *survey.Question {
	return &survey.Question{
		Name: "ipAddress",
		Prompt: &survey.Input{
			Message: fmt.Sprintf("Access List Entry%s:", message),
			Help:    usage.NetworkAccessListIPEntry,
			Default: publicIP,
		},
		Validate: survey.Required,
	}
}

func newDBUsernameQuestion(dbUser, message string, validation survey.Validator) *survey.Question {
	q := &survey.Question{
		Name: "dbUsername",
		Prompt: &survey.Input{
			Message: fmt.Sprintf("Database User Username%s:", message),
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
			Message: fmt.Sprintf("Database User Password%s:", message),
			Help:    usage.Password,
			Default: password,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.WeakPassword),
	}
}

func newSampleDataQuestion(clusterName string) *survey.Confirm {
	return &survey.Confirm{
		Message: fmt.Sprintf("Do you want to load sample data into %s?", clusterName),
		Help:    SampleDataHelp,
	}
}

func newMongoShellQuestionAccessDeployment(clusterName string) *survey.Confirm {
	return &survey.Confirm{
		Message: fmt.Sprintf("Do you want to access %s with MongoDB Shell?", clusterName),
		Help:    mongoShellHelp,
	}
}

func newMongoShellPathQuestion() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want to provide the path to your MongoDB shell binary?",
		Help:    mongoShellNotFoundHelp,
	}
}

func newIsMongoShellInstalledQuestion() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you have a MongoDB shell version installed on your machine?",
	}
}

func newMongoShellPathInput(defaultValue string) *survey.Question {
	return &survey.Question{
		Name: "mongoShellPath",
		Prompt: &survey.Input{
			Message: "Default MongoDB Shell Path:",
			Help:    mongoShellHelp,
			Default: defaultValue,
		},
		Validate: validate.Path,
	}
}

func newMongoShellQuestionOpenBrowser() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want to download MongoDB Shell [This will open www.mongodb.com on your browser]?",
	}
}

func newAtlasAccountQuestionOpenBrowser() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want to create an Atlas account [This will open www.mongodb.com on your browser]?",
	}
}

func newProfileDocQuestionOpenBrowser() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want more information to set up your profile [This will open www.mongodb.com on your browser]?",
	}
}

func providerQuestion(provider string) *survey.Question {
	if provider == "" {
		return newClusterProviderQuestion()
	}

	return nil
}

func clusterNameQuestion(clusterName string) *survey.Question {
	if clusterName != "" {
		return nil
	}

	message := ""
	newClusterName := autogeneratedName()
	if newClusterName != "" {
		message = fmt.Sprintf(" [Press Enter to use the auto-generated name '%s']", newClusterName)
	}

	return newClusterNameQuestion(newClusterName, message)
}

func dbUserPasswordQuestion(password string) (string, *survey.Question) {
	if password != "" {
		return "", nil
	}

	message := ""
	pwd, err := randgen.GenerateRandomBase64String(passwordLength)
	if err == nil {
		message = fmt.Sprintf(" [Press Enter to use an auto-generated password '%s']", pwd)
	}

	return pwd, newDBUserPasswordQuestion(pwd, message)
}

func dbUsernameQuestion(dbUser string, validation survey.Validator) *survey.Question {
	if dbUser != "" {
		return nil
	}

	message := ""
	newDBUser := autogeneratedName()
	if newDBUser != "" {
		message = fmt.Sprintf(" [Press Enter to use '%s']", newDBUser)
	}

	return newDBUsernameQuestion(newDBUser, message, validation)
}

func accessListQuestion(ipAddresses []string) *survey.Question {
	if len(ipAddresses) > 0 {
		return nil
	}

	message := ""
	publicIP := store.IPAddress()
	if publicIP != "" {
		message = fmt.Sprintf(" [Press Enter to use your public IP address '%s']", publicIP)
	}
	return newAccessListQuestion(publicIP, message)
}

func accessDeploymentQuestion(skip bool, clusterName string) *survey.Confirm {
	if skip {
		return nil
	}
	return newMongoShellQuestionAccessDeployment(clusterName)
}

func autogeneratedName() string {
	return "Quickstart-" + strconv.FormatInt(time.Now().Unix(), 10)
}
