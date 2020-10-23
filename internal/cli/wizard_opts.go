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

func (opts WizardOpts) newQuestion(flag *Flag) []*survey.Question {
	question := &survey.Question{Name: "Flag", Validate: survey.Required}

	if flag.Options == nil && !flag.Password {
		question.Prompt = &survey.Input{
			Message: "Insert " + flag.Name,
			Help:    flag.Usage,
		}
	} else {
		if flag.Password {
			question.Prompt = &survey.Password{
				Message: "Insert " + flag.Name,
				Help:    flag.Usage,
			}
		}

		if flag.Options != nil {
			question.Prompt = &survey.Select{
				Message: "Select one option for " + flag.Name + ":",
				Help:    flag.Usage,
				Options: flag.Options,
			}
		}
	}

	return []*survey.Question{question}
}

func (opts WizardOpts) newAnswer(question []*survey.Question) (string, error) {
	answer := struct {
		Flag string
	}{}

	err := survey.Ask(question, &answer)
	if err != nil {
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
