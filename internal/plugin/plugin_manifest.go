// Copyright 2024 MongoDB Inc
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

package plugin

import (
	"fmt"
	"regexp"
)

type PluginManifest struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Binary      string `yaml:"binary,omitempty"`
	Version     string `yaml:"version,omitempty"`
	Commands    map[string]struct {
		Description string `yaml:"description,omitempty"`
	} `yaml:"commands,omitempty"`
}

func (p *PluginManifest) IsValid() (bool, []error) {
	var errors []error
	errorMessage := `value "%s" is not defined`

	if p.Name == "" {
		errors = append(errors, fmt.Errorf(errorMessage, "name"))
	}
	if p.Description == "" {
		errors = append(errors, fmt.Errorf(errorMessage, "description"))
	}
	if p.Binary == "" {
		errors = append(errors, fmt.Errorf(errorMessage, "binary"))
	}
	if p.Version == "" {
		errors = append(errors, fmt.Errorf(errorMessage, "version"))
	} else if valid, _ := regexp.MatchString(`^\d+\.\d+\.\d+$`, p.Version); !valid {
		errors = append(errors, fmt.Errorf(`value in field "version" is not a valid semantic version`))
	}
	if p.Commands == nil {
		errors = append(errors, fmt.Errorf(errorMessage, "commands"))
	} else {
		for command, value := range p.Commands {
			if value.Description == "" {
				errors = append(errors, fmt.Errorf(`value "description" in command "%s" is not defined`, command))
			}
		}
	}

	if len(errors) > 0 {
		return false, errors
	}
	return true, nil
}
