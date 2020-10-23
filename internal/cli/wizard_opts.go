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

package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type WizardOpts struct {
	Wizard bool
}

type Flag struct {
	Name     string
	Usage    string
	Password bool
	Options  []string
}

// RunWizard allows to set flags with interactive prompts
func (opts WizardOpts) RunWizard(required, optional []*Flag) (map[string]string, error) {
	answers, err := opts.askRequiredFlags(required)
	if err != nil {
		return nil, err
	}

	answerOptional, err := opts.askOptionalFlags(optional)
	if err != nil {
		return nil, err
	}

	// Merging two maps
	for k, v := range answerOptional {
		answers[k] = v
	}

	return answers, nil
}

// askRequiredFlags allows the user to set required flags by using interactive prompts and stores the answers
func (opts WizardOpts) askRequiredFlags(flags []*Flag) (map[string]string, error) {
	m := make(map[string]string, len(flags))
	for _, flag := range flags {
		answer, err := opts.newAnswer(opts.newQuestion(flag))
		if err != nil {
			return nil, err
		}
		m[flag.Name] = answer
	}

	return m, nil
}

// askOptionalFlags allows the user to set optional flags by using interactive prompts and stores the answers
//
// Workflow:
//
// Step 1: The user selects one optional flag from a list of flags or select "done" to exit the loop
//
// Step 2: The user inserts the value for the selected flag
//
// Repeat Step 1 and Step 2 until the user selects "done"

func (opts WizardOpts) askOptionalFlags(flags []*Flag) (map[string]string, error) {
	m := make(map[string]*Flag, len(flags))
	answers := make(map[string]string, len(flags))

	optionalFlags := []string{"done"}
	for _, flag := range flags {
		m[flag.Name] = flag
		optionalFlags = append(optionalFlags, flag.Name)
	}

	done := false
	for !done {
		flag, err := opts.newAnswer(opts.newOptionalFlagQuestion(optionalFlags))
		if err != nil {
			return nil, err
		}

		if flag == "done" {
			done = true
		} else {
			answer, err := opts.newAnswer(opts.newQuestion(m[flag]))
			if err != nil {
				return nil, err
			}
			answers[flag] = answer
		}
	}

	return answers, nil
}

// newOptionalFlagQuestion generates the interactive prompt for selecting an optional flag from a list of flags
func (opts WizardOpts) newOptionalFlagQuestion(optionalFlags []string) []*survey.Question {
	return []*survey.Question{
		{
			Name:     "Flag",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: `Select one optional flag or select "done" to run the command:`,
				Options: optionalFlags,
			},
		},
	}
}

// newQuestion generates the interactive prompt for a flag
func (opts WizardOpts) newQuestion(flag *Flag) []*survey.Question {
	question := &survey.Question{Name: "Flag", Validate: survey.Required}
	switch {
	case flag.Password:
		question.Prompt = &survey.Password{
			Message: "Insert " + flag.Name,
			Help:    flag.Usage,
		}
	case flag.Options != nil:
		question.Prompt = &survey.Select{
			Message: "Select one option for " + flag.Name + ":",
			Help:    flag.Usage,
			Options: flag.Options,
		}
	default:
		question.Prompt = &survey.Input{
			Message: "Insert " + flag.Name,
			Help:    flag.Usage,
		}
	}

	return []*survey.Question{question}
}

func (opts WizardOpts) newAnswer(question []*survey.Question) (string, error) {
	answer := struct {
		Flag string
	}{}

	if err := survey.Ask(question, &answer); err != nil {
		return "", err
	}

	return answer.Flag, nil
}

func (opts WizardOpts) GetAnswer(answers map[string]string, key string) (string, error) {
	if answer, ok := answers[key]; ok {
		return answer, nil
	}

	return "", fmt.Errorf("the flag %s is required", key)
}
