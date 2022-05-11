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
	"errors"
	"reflect"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
)

func readAnswer(response interface{}, name string) (interface{}, error) { // based on https://pkg.go.dev/github.com/AlecAivazis/survey/v2/core#WriteAnswer
	v, ok := response.(map[string]interface{})
	if ok {
		return v[name], nil
	}

	target := reflect.ValueOf(response)
	if target.Kind() != reflect.Ptr {
		return nil, errors.New("invalid response: expected a map[string]interface{} or a pointer to struct")
	}
	elem := target.Elem()
	if elem.Kind() != reflect.Struct {
		return nil, errors.New("invalid response: expected a map[string]interface{} or a pointer to struct")
	}

	return elem.FieldByName(name).Interface(), nil
}

func TrackAsk(qs []*survey.Question, response interface{}, opts ...survey.AskOpt) error {
	err := survey.Ask(qs, response, opts...)
	for _, q := range qs {
		answer, _ := readAnswer(response, q.Name)
		trackSurvey(q.Prompt, answer, err)
	}
	return err
}

func TrackAskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	err := survey.AskOne(p, response, opts...)
	trackSurvey(p, response, err)
	return err
}

func trackSurvey(p survey.Prompt, response interface{}, e error) {
	if !config.TelemetryEnabled() {
		return
	}

	t, err := newTracker()
	if err != nil {
		logError(err)
		return
	}

	options := []eventOpt{}

	if e != nil {
		options = append(options, withError(e))
	}

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
