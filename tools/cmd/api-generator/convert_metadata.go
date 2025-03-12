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
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"
)

func specToMetadata(spec *openapi3.T) (map[string]*metadatatypes.Metadata, error) {
	metadataMap := make(map[string]*metadatatypes.Metadata, 0)

	for _, item := range spec.Paths.Map() {
		for _, operation := range item.Operations() {
			metadata, err := extractMetadata(operation)
			if err != nil {
				return nil, fmt.Errorf("failed to extract example: %w", err)
			}
			if metadata != nil {
				metadataMap[operation.OperationID] = metadata
			}
		}
	}
	return metadataMap, nil
}

// Returns a map of operationID:*Metadata.
func extractMetadata(operation *openapi3.Operation) (*metadatatypes.Metadata, error) {
	if operation == nil {
		return nil, nil
	}

	requestBodyExamples, err := extractRequestBodyExamples(operation.RequestBody)
	if err != nil {
		return nil, err
	}

	paramMap := extractParameterMetadata(operation.Parameters)

	return &metadatatypes.Metadata{
		Parameters:          paramMap,
		RequestBodyExamples: requestBodyExamples,
	}, nil
}

// For each parameter in an operation, the parameter name and example is extracted.
// A map of parameterName:example is returned.
func extractParameterMetadata(parameters openapi3.Parameters) map[string]metadatatypes.ParameterMetadata {
	result := make(map[string]metadatatypes.ParameterMetadata)

	for _, parameterRef := range parameters {
		example := ""
		if parameterExample, ok := parameterRef.Value.Schema.Value.Example.(string); ok {
			example = parameterExample
		}
		result[parameterRef.Value.Name] = metadatatypes.ParameterMetadata{
			Example: example,
			Usage:   parameterRef.Value.Description,
		}
	}

	return result
}

// For each verion of an operation, the version and examples are extracted.
// A map of version:[]examples is returned.
func extractRequestBodyExamples(requestBody *openapi3.RequestBodyRef) (map[string][]metadatatypes.RequestBodyExample, error) {
	if requestBody == nil || requestBody.Value == nil {
		return nil, nil
	}

	results := make(map[string][]metadatatypes.RequestBodyExample, 0)

	for versionedContentType, mediaType := range requestBody.Value.Content {
		examples := make([]metadatatypes.RequestBodyExample, 0)
		version, _, err := extractVersionAndContentType(versionedContentType)
		if err != nil {
			return nil, fmt.Errorf("unsupported version %q error: %w", versionedContentType, err)
		}

		if shouldIgnoreVersion(version) {
			continue
		}

		for name, exampleRef := range mediaType.Examples {
			if exampleRef == nil || exampleRef.Value == nil {
				continue
			}

			exampleName := name
			if exampleRef.Value.Summary != "" {
				exampleName = exampleRef.Value.Summary
			}

			var value any
			if exampleRef.Value != nil && exampleRef.Value.Value != nil {
				value = exampleRef.Value.Value
			}

			result := metadatatypes.RequestBodyExample{
				Name:        exampleName,
				Description: exampleRef.Value.Description,
				Value:       toJSONString(value),
			}

			examples = append(examples, result)
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
	if data == nil {
		return ""
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		_, _ = log.Warningln("Unable to convert to JSON string")
		return ""
	}
	return string(jsonData)
}
