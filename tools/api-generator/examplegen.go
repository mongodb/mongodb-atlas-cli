// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/template"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

// Returns a map of operationID:*Examples.
func extractExamples(operation *openapi3.Operation, examplesMap map[string]*api.Examples) error {
	paramMap := extractParameterExamples(operation.Parameters)

	requestBodyExamples, err := extractRequestBodyExamples(operation.RequestBody)
	if err != nil {
		return err
	}

	_, exists := examplesMap[operation.OperationID]
	if exists {
		return fmt.Errorf("expect no example to already be stored for OperationID: %s", operation.OperationID)
	}

	examplesMap[operation.OperationID] = &api.Examples{
		ParameterExample:    paramMap,
		RequestBodyExamples: requestBodyExamples,
	}

	return nil
}

// For each parameter in an operation, the parameter name and example is extracted.
// A map of parameterName:example is returned.
func extractParameterExamples(parameters openapi3.Parameters) map[string]string {
	result := make(map[string]string)

	for _, parameterRef := range parameters {
		parameterExample, ok := parameterRef.Value.Schema.Value.Example.(string)
		if !ok {
			continue
		}

		result[parameterRef.Value.Name] = parameterExample
	}

	return result
}

// For each verion of an operation, the version and examples are extracted.
// A map of version:[]examples is returned.
func extractRequestBodyExamples(requestBody *openapi3.RequestBodyRef) (map[string][]api.RequestBodyExample, error) {
	examples := make([]api.RequestBodyExample, 0)
	results := make(map[string][]api.RequestBodyExample, 0)
	if requestBody == nil {
		return nil, nil
	}

	for versionedContentType, mediaType := range requestBody.Value.Content {
		version, _, err := extractVersionAndContentType(versionedContentType)
		if err != nil {
			return nil, fmt.Errorf("unsupported version %q error: %w", versionedContentType, err)
		}

		for name, exampleRef := range mediaType.Examples {
			if exampleRef != nil {
				result := api.RequestBodyExample{
					Name:        name,
					Description: exampleRef.Value.Description,
					Value:       toJSONString(exampleRef.Value.Value.(map[string]any)),
				}

				examples = append(examples, result)
			}
		}

		if len(examples) != 0 {
			// Ensure list of examples are sorted in same order
			sort.Slice(examples, func(i, j int) bool {
				return examples[i].Name < examples[j].Name
			})
			results[version] = examples
		}
	}

	return results, nil
}

func toJSONString(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		_, _ = log.Warningln("Unable to convert to JSON string")
		return ""
	}
	return string(jsonData)
}

func exportExamples(examplesMap map[string]*api.Examples) error {
	file, err := os.Create("./internal/api/examples.go")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("examples.go.tmpl").Funcs(template.FuncMap{
		"currentYear": func() int {
			return time.Now().UTC().Year()
		},
	}).Parse(examplesTemplateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template and write to file
	if err := tmpl.Execute(file, examplesMap); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
