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

package telemetry

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
)

func TrackAskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	err := survey.AskOne(p, response, opts...)
	trackSurvey(p, response)
	return err
}

func trackSurvey(p survey.Prompt, response interface{}) {
	if !config.TelemetryEnabled() {
		return
	}

	t, err := newTracker()
	if err != nil {
		logError(err)
		return
	}

	options := []eventOpt{}

	switch v := p.(type) {
	case *survey.Confirm:
		r := *response.(*bool)

		options = append(options, withPrompt(v.Message, "confirm"), withDefault(r == v.Default))
	case *survey.Input:
		r := *response.(*string)

		options = append(options, withPrompt(v.Message, "input"), withDefault(r == v.Default), withEmpty(r == ""))
	case *survey.Password:
		r := *response.(*string)

		options = append(options, withPrompt(v.Message, "input"), withEmpty(r == ""))
	case *survey.Select:
		r := *response.(*string)

		options = append(options, withPrompt(v.Message, "select"), withDefault(r == v.Default), withEmpty(r == ""))
	}

	event := newEvent(options...)

	if err = t.save(event); err != nil {
		logError(err)
	}
}
