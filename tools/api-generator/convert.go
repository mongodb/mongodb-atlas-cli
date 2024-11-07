// Copyright 2024 MongoDB Inc
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
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
)

var (
	versionRegex = regexp.MustCompile(`^application/vnd\.atlas\.(?P<version>\d{4}-\d{2}-\d{2})\+(?P<contentType>[\w]+)$`)
)

func specToCommands(spec *openapi3.T) (api.GroupedAndSortedCommands, error) {
	groups := make(map[string]*api.Group, 0)

	for path, item := range spec.Paths.Map() {
		for verb, operation := range item.Operations() {
			command, err := operationToCommand(path, verb, operation)
			if err != nil {
				return nil, fmt.Errorf("failed to convert operation to command: %w", err)
			}

			if len(operation.Tags) != 1 {
				return nil, fmt.Errorf("expect every operation to have exactly 1 tag, got: %v", len(operation.Tags))
			}

			tag := operation.Tags[0]
			if _, ok := groups[tag]; !ok {
				group, err := groupForTag(spec, tag)
				if err != nil {
					return nil, err
				}

				groups[tag] = group
			}

			groups[tag].Commands = append(groups[tag].Commands, *command)
		}
	}

	// Sort commands inside of groups
	sortedGroups := make([]api.Group, 0, len(groups))
	for _, group := range groups {
		sort.Slice(group.Commands, func(i, j int) bool {
			return group.Commands[i].OperationID < group.Commands[j].OperationID
		})

		sortedGroups = append(sortedGroups, *group)
	}

	// Sort groups
	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Name < sortedGroups[j].Name
	})

	return sortedGroups, nil
}

func operationToCommand(path, verb string, operation *openapi3.Operation) (*api.Command, error) {
	httpVerb, err := api.ToHTTPVerb(verb)
	if err != nil {
		return nil, err
	}

	parameters, err := extractParameters(operation.Parameters)
	if err != nil {
		return nil, err
	}

	versions, err := buildVersions(operation)
	if err != nil {
		return nil, err
	}

	description, err := Clean(operation.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to clean description: %w", err)
	}

	command := api.Command{
		OperationID: operation.OperationID,
		Description: description,
		RequestParameters: api.RequestParameters{
			URL:             path,
			QueryParameters: parameters.query,
			URLParameters:   parameters.url,
			Verb:            httpVerb,
		},
		Versions: versions,
	}

	return &command, nil
}

// Struct to hold both types of parameters.
type parameterSet struct {
	query []api.Parameter
	url   []api.Parameter
}

// Extract and categorize parameters.
func extractParameters(parameters openapi3.Parameters) (parameterSet, error) {
	queryParameters := make([]api.Parameter, 0)
	urlParameters := make([]api.Parameter, 0)

	for _, parameterRef := range parameters {
		parameter := parameterRef.Value

		description, err := Clean(parameter.Description)
		if err != nil {
			return parameterSet{}, fmt.Errorf("failed to clean description: %w", err)
		}

		apiParameter := api.Parameter{
			Name:        parameter.Name,
			Description: description,
			Required:    parameter.Required,
		}

		switch parameter.In {
		case "query":
			queryParameters = append(queryParameters, apiParameter)
		case "path":
			urlParameters = append(urlParameters, apiParameter)
		default:
			return parameterSet{}, fmt.Errorf("invalid parameter 'in' location: %s", parameter.In)
		}
	}

	return parameterSet{
		query: queryParameters,
		url:   urlParameters,
	}, nil
}

// Build versions from responses and request body.
func buildVersions(operation *openapi3.Operation) ([]api.Version, error) {
	versionsMap := make(map[string]*api.Version)

	if err := processResponses(operation.Responses, versionsMap); err != nil {
		return nil, err
	}

	if err := processRequestBody(operation.RequestBody, versionsMap); err != nil {
		return nil, err
	}

	return sortVersions(versionsMap), nil
}

// Process response content types.
func processResponses(responses *openapi3.Responses, versionsMap map[string]*api.Version) error {
	for statusString, responses := range responses.Map() {
		statusCode, err := strconv.Atoi(statusString)
		if err != nil {
			return fmt.Errorf("http status code '%s' is not numeric: %w", statusString, err)
		}

		if statusCode < 200 || statusCode >= 300 {
			continue
		}

		for versionedContentType := range responses.Value.Content {
			if err := addContentTypeToVersion(versionedContentType, versionsMap, false); err != nil {
				return err
			}
		}
	}
	return nil
}

// Process request body content types.
func processRequestBody(requestBody *openapi3.RequestBodyRef, versionsMap map[string]*api.Version) error {
	if requestBody == nil {
		return nil
	}

	for versionedContentType := range requestBody.Value.Content {
		if err := addContentTypeToVersion(versionedContentType, versionsMap, true); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to add content type to version map.
func addContentTypeToVersion(versionedContentType string, versionsMap map[string]*api.Version, isRequest bool) error {
	version, contentType, err := extractVersionAndContentType(versionedContentType)
	if err != nil {
		return fmt.Errorf("unsupported version '%s' error: %w", versionedContentType, err)
	}

	if _, ok := versionsMap[version]; !ok {
		versionsMap[version] = &api.Version{
			Version:              version,
			RequestContentTypes:  []string{},
			ResponseContentTypes: []string{},
		}
	}

	if isRequest {
		versionsMap[version].RequestContentTypes = append(versionsMap[version].RequestContentTypes, contentType)
	} else {
		versionsMap[version].ResponseContentTypes = append(versionsMap[version].ResponseContentTypes, contentType)
	}

	return nil
}

// Sort versions and their content types.
func sortVersions(versionsMap map[string]*api.Version) []api.Version {
	versions := make([]api.Version, 0)

	for _, version := range versionsMap {
		sort.Slice(version.RequestContentTypes, func(i, j int) bool {
			return version.RequestContentTypes[i] < version.RequestContentTypes[j]
		})

		sort.Slice(version.ResponseContentTypes, func(i, j int) bool {
			return version.ResponseContentTypes[i] < version.ResponseContentTypes[j]
		})

		versions = append(versions, *version)
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Version < versions[j].Version
	})

	return versions
}

func groupForTag(spec *openapi3.T, tag string) (*api.Group, error) {
	description := ""

	if specTag := spec.Tags.Get(tag); specTag != nil {
		cleanDescription, err := Clean(specTag.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to clean description: %w", err)
		}
		description = cleanDescription
	}

	return &api.Group{
		Name:        tag,
		Description: description,
		Commands:    []api.Command{},
	}, nil
}

func extractVersionAndContentType(input string) (version string, contentType string, err error) {
	matches := versionRegex.FindStringSubmatch(input)
	if matches == nil {
		return "", "", errors.New("invalid format")
	}

	// Get the named group indices
	versionIndex := versionRegex.SubexpIndex("version")
	contentTypeIndex := versionRegex.SubexpIndex("contentType")

	return matches[versionIndex], matches[contentTypeIndex], nil
}
