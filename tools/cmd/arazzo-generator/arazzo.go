// Copyright 2026 MongoDB Inc
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

package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Document is the subset of the Arazzo Description the generator consumes.
type Document struct {
	Arazzo             string              `yaml:"arazzo"`
	Info               Info                `yaml:"info"`
	SourceDescriptions []SourceDescription `yaml:"sourceDescriptions"`
	Workflows          []Workflow          `yaml:"workflows"`
}

type Info struct {
	Title       string `yaml:"title"`
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

type SourceDescription struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Type string `yaml:"type"`
}

type Workflow struct {
	WorkflowID  string       `yaml:"workflowId"`
	Summary     string       `yaml:"summary"`
	Description string       `yaml:"description"`
	Inputs      InputsSchema `yaml:"inputs"`
	Steps       []Step       `yaml:"steps"`
}

type InputsSchema struct {
	Type       string                   `yaml:"type"`
	Properties map[string]InputProperty `yaml:"properties"`
	Required   []string                 `yaml:"required"`
}

type InputProperty struct {
	Type        string   `yaml:"type"`
	Description string   `yaml:"description"`
	Enum        []string `yaml:"enum"`
	Pattern     string   `yaml:"pattern"`
	Default     any      `yaml:"default"`
	XBinding    string   `yaml:"x-binding"`
	XDerived    *Derived `yaml:"x-derived"`
}

// Derived declares how the translator resolves a derived input or output.
// Template substitutes embedded runtime expressions; From with Values selects
// a value for the given service and leaves the value unset when the service
// has no mapping; Transform post-processes the resolved template value.
type Derived struct {
	Template  string            `yaml:"template"`
	From      string            `yaml:"from"`
	Values    map[string]string `yaml:"values"`
	Transform string            `yaml:"transform"`
}

type Step struct {
	StepID          string             `yaml:"stepId"`
	Description     string             `yaml:"description"`
	OperationID     string             `yaml:"operationId"`
	Parameters      []StepParameter    `yaml:"parameters"`
	RequestBody     *RequestBody       `yaml:"requestBody"`
	SuccessCriteria []Criterion        `yaml:"successCriteria"`
	OnFailure       []FailureAction    `yaml:"onFailure"`
	Outputs         map[string]string  `yaml:"outputs"`
	XDerivedOutputs map[string]Derived `yaml:"x-derived-outputs"`
	XCleanup        string             `yaml:"x-cleanup"`
}

type StepParameter struct {
	Name  string `yaml:"name"`
	In    string `yaml:"in"`
	Value any    `yaml:"value"`
}

type RequestBody struct {
	ContentType string `yaml:"contentType"`
	Payload     any    `yaml:"payload"`
}

type Criterion struct {
	Condition string `yaml:"condition"`
}

type FailureAction struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	RetryAfter int    `yaml:"retryAfter"`
	RetryLimit int    `yaml:"retryLimit"`
}

func (s Step) hasRetryOnFailure() bool {
	for _, action := range s.OnFailure {
		if action.Type == retryAction {
			return true
		}
	}

	return false
}

func loadDocument(path string) (*Document, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var doc Document
	if err := yaml.Unmarshal(content, &doc); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	if len(doc.Workflows) == 0 {
		return nil, fmt.Errorf("%s: no workflows defined", path)
	}

	return &doc, nil
}
