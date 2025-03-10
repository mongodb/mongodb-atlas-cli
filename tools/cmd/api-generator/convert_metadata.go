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
	if len(requestBodyExamples) == 0 {
		return nil, nil
	}

	paramMap := extractParameterExamples(operation.Parameters)

	return &metadatatypes.Metadata{
		ParameterExample:    paramMap,
		RequestBodyExamples: requestBodyExamples,
	}, nil
}

// For each parameter in an operation, the parameter name and example is extracted.
// A map of parameterName:example is returned.
func extractParameterExamples(parameters openapi3.Parameters) map[string]string {
	result := make(map[string]string)

	for _, parameterRef := range parameters {
		parameterExample, ok := parameterRef.Value.Schema.Value.Example.(string)
		if ok {
			result[parameterRef.Value.Name] = parameterExample
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

		for name, exampleRef := range mediaType.Examples {
			if exampleRef == nil || exampleRef.Value == nil {
				continue
			}
			result := metadatatypes.RequestBodyExample{
				Name:        name,
				Description: exampleRef.Value.Description,
				Value:       toJSONString(exampleRef.Value.Value.(map[string]any)),
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
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		_, _ = log.Warningln("Unable to convert to JSON string")
		return ""
	}
	return string(jsonData)
}
