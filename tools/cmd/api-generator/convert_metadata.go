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
	"errors"
	"fmt"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"
)

func specToMetadata(spec *openapi3.T) (metadatatypes.Metadata, error) {
	metadataMap := make(metadatatypes.Metadata, 0)

	for _, item := range spec.Paths.Map() {
		for httpMethod, operation := range item.Operations() {
			metadata, err := extractMetadata(httpMethod, operation)
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
func extractMetadata(httpMethod string, operation *openapi3.Operation) (*metadatatypes.OperationMetadata, error) {
	if operation == nil {
		return nil, nil
	}

	paramMetadata := extractParameterMetadata(operation.Parameters)

	requestBodyExamples, err := extractRequestBodyExamples(operation.RequestBody)
	if err != nil {
		return nil, err
	}

	paramExamples := extractParameterExamples(operation.Parameters)

	examples, err := buildExamples(requestBodyExamples, paramExamples, httpMethod, operation)
	if err != nil {
		return nil, err
	}

	return &metadatatypes.OperationMetadata{
		Parameters: paramMetadata,
		Examples:   examples,
	}, nil
}

func extractParameterMetadata(parameters openapi3.Parameters) map[string]metadatatypes.ParameterMetadata {
	result := make(map[string]metadatatypes.ParameterMetadata)

	for _, parameterRef := range parameters {
		result[parameterRef.Value.Name] = metadatatypes.ParameterMetadata{
			Usage: parameterRef.Value.Description,
		}
	}

	return result
}

type extractedExamples struct {
	Example  any
	Examples openapi3.Examples
}

func extractParameterExamples(parameters openapi3.Parameters) map[string]extractedExamples {
	result := make(map[string]extractedExamples)

	for _, parameterRef := range parameters {
		defaultExample := parameterRef.Value.Example
		if defaultExample == nil {
			defaultExample = parameterRef.Value.Schema.Value.Example
		}
		result[parameterRef.Value.Name] = extractedExamples{
			Example:  defaultExample,
			Examples: parameterRef.Value.Examples,
		}
	}

	return result
}

func extractRequestBodyExamples(requestBody *openapi3.RequestBodyRef) (map[string]extractedExamples, error) {
	if requestBody == nil || requestBody.Value == nil {
		return nil, nil
	}

	results := make(map[string]extractedExamples, 0)

	for versionedContentType, mediaType := range requestBody.Value.Content {
		version, _, err := extractVersionAndContentType(versionedContentType)
		if err != nil {
			return nil, fmt.Errorf("unsupported version %q error: %w", versionedContentType, err)
		}

		if shouldIgnoreVersion(version) {
			continue
		}

		results[version] = extractedExamples{
			Example:  mediaType.Example,
			Examples: mediaType.Examples,
		}
	}

	return results, nil
}

func extractDefaultVersion(operation *openapi3.Operation) (string, error) {
	defaultVersion := ""

	var defaultResponse *openapi3.ResponseRef
	for code := 200; code < 300; code++ {
		if response := operation.Responses.Status(code); response != nil {
			defaultResponse = response
			break
		}
	}
	if defaultResponse == nil {
		return "", errors.New("default version not found")
	}

	for mime := range defaultResponse.Value.Content {
		version, _, err := extractVersionAndContentType(mime)
		if err != nil {
			return "", fmt.Errorf("unsupported version %q error: %w", mime, err)
		}
		defaultVersion = version
	}

	return defaultVersion, nil
}

func extractAllKeys(parameterExamples map[string]extractedExamples) map[string]bool {
	allKeys := map[string]bool{}
	for _, examples := range parameterExamples {
		if examples.Example != nil {
			allKeys["-"] = true
		}
		for key := range examples.Examples {
			allKeys[key] = true
		}
	}
	return allKeys
}

func extractRequiredFlagNames(operation *openapi3.Operation) map[string]bool {
	requiredFlags := map[string]bool{}
	for _, parameterRef := range operation.Parameters {
		if parameterRef.Value.Required {
			requiredFlags[parameterRef.Value.Name] = true
		}
	}
	return requiredFlags
}

func extractFlagValue(key string, flagName string, examples extractedExamples, required bool) string {
	if key != "-" {
		if examples.Examples[key] != nil {
			return toValueString(examples.Examples[key].Value)
		}
	}
	if examples.Example != nil {
		return toValueString(examples.Example)
	}
	if required {
		return fmt.Sprintf("[%s]", flagName)
	}

	return ""
}

func buildExamplesNoResponseBody(parameterExamples map[string]extractedExamples, httpMethod string, operation *openapi3.Operation) (map[string][]metadatatypes.Example, error) {
	examples := map[string][]metadatatypes.Example{}

	requiredFlagNames := extractRequiredFlagNames(operation)

	if httpMethod == "POST" || httpMethod == "PUT" {
		return nil, nil // don't bother we need a request body
	}

	defaultVersion, err := extractDefaultVersion(operation)
	if err != nil {
		return nil, err
	}

	allKeys := extractAllKeys(parameterExamples)
	for key := range allKeys {
		example := metadatatypes.Example{
			Source:      key,
			Name:        "",
			Description: "",
			Value:       "",
			Flags:       map[string]string{},
		}
		for flagName, flagExample := range parameterExamples {
			if flagExample.Examples[key] != nil {
				if example.Name == "" && flagExample.Examples[key].Value.Summary != "" {
					example.Name = flagExample.Examples[key].Value.Summary
				}
				if example.Description == "" && flagExample.Examples[key].Value.Description != "" {
					example.Description = flagExample.Examples[key].Value.Description
				}
			}
			value := extractFlagValue(key, flagName, flagExample, requiredFlagNames[flagName])
			if value != "" {
				example.Flags[flagName] = value
			}
		}

		if examples[defaultVersion] == nil {
			examples[defaultVersion] = []metadatatypes.Example{}
		}
		examples[defaultVersion] = append(examples[defaultVersion], example)
	}

	if len(examples) == 0 {
		return nil, nil
	}

	return examples, nil
}

func buildExamplesWithResponseBody(requestBodyExamples, parameterExamples map[string]extractedExamples, operation *openapi3.Operation) (map[string][]metadatatypes.Example, error) {
	examples := map[string][]metadatatypes.Example{}

	requiredFlagNames := extractRequiredFlagNames(operation)

	for version, requestBodyExamples := range requestBodyExamples {
		if requestBodyExamples.Example != nil {
			example := metadatatypes.Example{
				Source: "-",
				Value:  toValueString(requestBodyExamples.Example),
				Flags:  map[string]string{},
			}
			for flagName, flagExample := range parameterExamples {
				value := extractFlagValue("-", flagName, flagExample, requiredFlagNames[flagName])
				if value != "" {
					example.Flags[flagName] = value
				}
			}
			if examples[version] == nil {
				examples[version] = []metadatatypes.Example{}
			}
			examples[version] = append(examples[version], example)
		}
		for exampleName, requestBodyExample := range requestBodyExamples.Examples {
			example := metadatatypes.Example{
				Source:      exampleName,
				Name:        requestBodyExample.Value.Summary,
				Description: requestBodyExample.Value.Description,
				Value:       toValueString(requestBodyExample.Value.Value),
				Flags:       map[string]string{},
			}
			for flagName, flagExample := range parameterExamples {
				value := extractFlagValue(exampleName, flagName, flagExample, requiredFlagNames[flagName])
				if value != "" {
					example.Flags[flagName] = value
				}
			}
			if examples[version] == nil {
				examples[version] = []metadatatypes.Example{}
			}
			examples[version] = append(examples[version], example)
		}
	}

	if len(examples) == 0 {
		return nil, nil
	}

	return examples, nil
}

func buildExamples(requestBodyExamples, parameterExamples map[string]extractedExamples, httpMethod string, operation *openapi3.Operation) (map[string][]metadatatypes.Example, error) {
	if len(requestBodyExamples) == 0 {
		return buildExamplesNoResponseBody(parameterExamples, httpMethod, operation)
	}

	return buildExamplesWithResponseBody(requestBodyExamples, parameterExamples, operation)
}

func toValueString(data any) string {
	if data == nil {
		return ""
	}

	switch value := data.(type) {
	case string:
		return value
	case bool:
		return strconv.FormatBool(value)
	case float64:
		return fmt.Sprintf("%v", value)
	case map[string]any:
		if len(value) == 0 {
			return ""
		}

		jsonData, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			_, _ = log.Warningln("unable to convert to JSON string")
			return ""
		}

		return string(jsonData)
	case []any:
		if len(value) == 0 {
			return ""
		}

		jsonData, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			_, _ = log.Warningln("unable to convert to JSON string")
			return ""
		}

		return string(jsonData)
	default:
		_, _ = log.Warningln("unable to find type")
		return fmt.Sprintf("%v", value)
	}
}
