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

package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

var (
	ErrFailedToBuildHTTPRequest = errors.New("failed to build http request")
	ErrUnsupportedContentType   = errors.New("unsupported content type")
	ErrFailedToBuildPath        = errors.New("failed to build path")
	ErrFailedToBuildQuery       = errors.New("failed to build query")
	ErrInvalidBaseURL           = errors.New("invalid base url")
	ErrVersionNotFound          = errors.New("version not found")
)

func ConvertToHTTPRequest(baseURL string, request CommandRequest) (*http.Request, error) {
	// Find the version details
	version, err := selectVersion(request.Version, request.Command.Versions)
	if err != nil {
		return nil, errors.Join(ErrFailedToBuildHTTPRequest, err)
	}

	// Create base url object
	requestURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Join(ErrInvalidBaseURL, err)
	}

	// Convert the URL template into a specific URL
	// Example: /api/atlas/v2/groups/{groupId}/clusters -> /api/atlas/v2/groups/a12345b/clusters
	path, err := buildPath(request.Command.RequestParameters.URL, request.Command.RequestParameters.URLParameters, request.Parameters)
	if err != nil {
		return nil, errors.Join(ErrFailedToBuildPath, err)
	}

	// Add path to base url
	requestURL = requestURL.JoinPath(path)

	// Build the query parameters
	query, err := buildQueryParameters(request.Command.RequestParameters.QueryParameters, request.Parameters)
	if err != nil {
		return nil, errors.Join(ErrFailedToBuildQuery, err)
	}
	requestURL.RawQuery = query

	// Get the content type
	contentType := contentType(version)

	// Only Set the HTTP body when we have a content type
	// GET/DELETE/certain POSTs don't have a content type => no body
	requestBody := request.Content
	if contentType == nil {
		requestBody = http.NoBody
	}

	// Build the actual http request
	//nolint: noctx // we're not passing the context here on purpose, the caller adds the context
	httpRequest, err := http.NewRequest(request.Command.RequestParameters.Verb, requestURL.String(), requestBody)
	if err != nil {
		return nil, errors.Join(ErrVersionNotFound, err)
	}

	// Set accept header
	accept, err := acceptHeader(version, request.ContentType)
	if err != nil {
		return nil, errors.Join(ErrUnsupportedContentType, err)
	}
	httpRequest.Header.Add("Accept", accept)

	// Set content type header, if needed
	if contentType != nil {
		httpRequest.Header.Add("Content-Type", *contentType)
	}

	return httpRequest, nil
}

// https://swagger.io/docs/specification/v3_0/serialization/#path-parameters
// Our spec only contains the default:
// > The default serialization method is style: simple and explode: false.
// > Given the path /users/{id}, the path parameter id is serialized as follows:
// > style: simple + exploded: false, example template /users/{id}, single value: /users/5, array: /users/3,4,5.
func buildPath(path string, commandURLParameters []shared_api.Parameter, parameterValues map[string][]string) (string, error) {
	for _, commandURLParameter := range commandURLParameters {
		values, exist := parameterValues[commandURLParameter.Name]
		if !exist || len(values) == 0 {
			return "", fmt.Errorf("url parameter '%s' is not set", commandURLParameter.Name)
		}

		// Path encode all values
		for i := range values {
			values[i] = url.PathEscape(values[i])
		}

		value := strings.Join(values, ",")
		path = strings.ReplaceAll(path, "{"+commandURLParameter.Name+"}", value)
	}

	return path, nil
}

// https://swagger.io/docs/specification/v3_0/serialization/#query-parameters
// Our spec only contains the default:
// style: form + exploded: true, example template /users{?id*}, single value: /users?id=5, array: /users?id=3&id=4&id=5
func buildQueryParameters(commandQueryParameters []shared_api.Parameter, parameterValues map[string][]string) (string, error) {
	query := new(url.URL).Query()

	for _, commandQueryParameter := range commandQueryParameters {
		values, exist := parameterValues[commandQueryParameter.Name]

		if !exist || len(values) == 0 {
			if commandQueryParameter.Required {
				return "", fmt.Errorf("query parameter '%s' is missing", commandQueryParameter.Name)
			}

			continue
		}

		for _, value := range values {
			query.Add(commandQueryParameter.Name, value)
		}
	}

	return query.Encode(), nil
}

// select a version from a list of versions, throws an error if no match is found.
func selectVersion(selectedVersion shared_api.Version, versions []shared_api.CommandVersion) (*shared_api.CommandVersion, error) {
	for _, version := range versions {
		if version.Version.Equal(selectedVersion) {
			return &version, nil
		}
	}

	return nil, fmt.Errorf("version '%s' not found", selectedVersion)
}

// generate the accept header using the given format string
// try to find the content type in the list of response content types, if not found set the type to json.
func acceptHeader(version *shared_api.CommandVersion, requestedContentType string) (string, error) {
	contentType := ""
	supportedTypes := make([]string, 0, len(version.ResponseContentTypes))

	for _, responseContentType := range version.ResponseContentTypes {
		if responseContentType == requestedContentType {
			contentType = requestedContentType
		}

		supportedTypes = append(supportedTypes, responseContentType)
	}

	if contentType == "" {
		return "", fmt.Errorf("expected one of the following values: [%s], but got '%s' instead", strings.Join(supportedTypes, ","), requestedContentType)
	}

	return fmt.Sprintf("application/vnd.atlas.%s+%s", version.Version, contentType), nil
}

func contentType(version *shared_api.CommandVersion) *string {
	if version.RequestContentType != "" {
		contentType := fmt.Sprintf("application/vnd.atlas.%s+%s", version.Version, version.RequestContentType)
		return &contentType
	}

	return nil
}

type DefaultCommandConverter struct {
	configProvider ConfigProvider
}

func (c *DefaultCommandConverter) ConvertToHTTPRequest(request CommandRequest) (*http.Request, error) {
	// Get the base URL
	baseURL, err := c.configProvider.GetBaseURL()
	if err != nil {
		return nil, errors.Join(ErrFailedToGetBaseURL, err)
	}

	return ConvertToHTTPRequest(baseURL, request)
}

func NewDefaultCommandConverter(configProvider ConfigProvider) (*DefaultCommandConverter, error) {
	if configProvider == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("configProvider is nil"))
	}

	return &DefaultCommandConverter{
		configProvider: configProvider,
	}, nil
}
