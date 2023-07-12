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

//go:build unit

package test

import (
	"testing"
	"text/template"

	"github.com/jba/templatecheck"
)

// VerifyTemplate validates that the given template string is valid.
func VerifyOutputTemplate(t *testing.T, templateString string, expected interface{}) {
	t.Helper()
	parsedTemplate, err := template.New("output").Parse(templateString)
	if err != nil {
		t.Fatalf("Failed to validate table format template: %v", err)
	}
	err = templatecheck.CheckText(parsedTemplate, expected)
	if err != nil {
		t.Fatalf("Failed to validate table format template: %v", err)
	}
}
