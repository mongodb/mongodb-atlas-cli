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
	"testing"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

const headerParam = "headerParam"

func TestExtractVersionAndContentType(t *testing.T) {
	tests := []struct {
		input           string
		wantVersion     api.Version
		wantContentType string
	}{
		{
			"application/vnd.atlas.2025-01-01+json",
			api.NewStableVersion(2025, 1, 1),
			"json",
		},
		{
			"application/vnd.atlas.2024-08-05+json",
			api.NewStableVersion(2024, 8, 5),
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01+csv",
			api.NewStableVersion(2023, 1, 1),
			"csv",
		},
		{
			"application/vnd.atlas.preview+json",
			api.NewPreviewVersion(),
			"json",
		},
		{
			"application/vnd.atlas.preview+csv",
			api.NewPreviewVersion(),
			"csv",
		},
		{
			"application/vnd.atlas.2024-08-05.upcoming+json",
			api.NewUpcomingVersion(2024, 8, 5),
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01.upcoming+csv",
			api.NewUpcomingVersion(2023, 1, 1),
			"csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotVersion, gotContentType, err := extractVersionAndContentType(tt.input)
			if err != nil {
				t.Fatalf("Error = %v", err)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("Expected: %s. Got: %s", tt.wantVersion, gotVersion)
			}
			if gotContentType != tt.wantContentType {
				t.Errorf("Expected: %s Got: %s,", tt.wantContentType, gotContentType)
			}
		})
	}
}

func TestExtractParameters_HeaderParametersSkipped(t *testing.T) {
	// Create test parameters with different 'in' locations
	parameters := openapi3.Parameters{
		{
			Value: &openapi3.Parameter{
				Name:     "queryParam",
				In:       "query",
				Required: false,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{"string"},
					},
				},
			},
		},
		{
			Value: &openapi3.Parameter{
				Name:     "pathParam",
				In:       "path",
				Required: true,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{"string"},
					},
				},
			},
		},
		{
			Value: &openapi3.Parameter{
				Name:     headerParam,
				In:       "header",
				Required: false,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{"string"},
					},
				},
			},
		},
		{
			Value: &openapi3.Parameter{
				Name:     "anotherQueryParam",
				In:       "query",
				Required: false,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{"string"},
					},
				},
			},
		},
	}

	result, err := extractParameters(parameters)
	if err != nil {
		t.Fatalf("extractParameters failed: %v", err)
	}

	// Verify query parameters are included
	if len(result.query) != 2 {
		t.Errorf("Expected 2 query parameters, got %d", len(result.query))
	}

	queryParamNames := make(map[string]bool)
	for _, param := range result.query {
		queryParamNames[param.Name] = true
	}
	if !queryParamNames["queryParam"] {
		t.Error("Expected 'queryParam' to be in query parameters")
	}
	if !queryParamNames["anotherQueryParam"] {
		t.Error("Expected 'anotherQueryParam' to be in query parameters")
	}

	// Verify path parameters are included
	if len(result.url) != 1 {
		t.Errorf("Expected 1 path parameter, got %d", len(result.url))
	}
	if result.url[0].Name != "pathParam" {
		t.Errorf("Expected path parameter name 'pathParam', got '%s'", result.url[0].Name)
	}

	// Verify header parameter is NOT included (skipped)
	for _, param := range result.query {
		if param.Name == headerParam {
			t.Error("Header parameter 'headerParam' should not be in query parameters")
		}
	}
	for _, param := range result.url {
		if param.Name == headerParam {
			t.Error("Header parameter 'headerParam' should not be in URL parameters")
		}
	}
}

func TestBuildVersions_DeprecatedOperation(t *testing.T) {
	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/vnd.atlas.2023-01-01+json": &openapi3.MediaType{
					Extensions: map[string]any{},
				},
				"application/vnd.atlas.2024-01-01+json": &openapi3.MediaType{
					Extensions: map[string]any{},
				},
			},
		},
	})

	operation := &openapi3.Operation{
		Deprecated: true,
		Responses:  responses,
	}

	versions, err := buildVersions(now, operation)
	if err != nil {
		t.Fatalf("buildVersions() error = %v", err)
	}

	if len(versions) != 2 {
		t.Fatalf("Expected 2 versions, got %d", len(versions))
	}

	for _, version := range versions {
		if !version.Deprecated {
			t.Errorf("Expected version %s to be deprecated, but it's not", version.Version)
		}
	}
}

func TestAddContentTypeToVersion_DeprecatedWithSunset(t *testing.T) {
	versionsMap := make(map[string]*api.CommandVersion)
	sunsetDate := time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)

	extensions := map[string]any{
		"x-sunset": "2026-01-15",
	}

	err := addContentTypeToVersion("application/vnd.atlas.2023-01-01+json", versionsMap, extensions, false)
	if err != nil {
		t.Fatalf("addContentTypeToVersion() error = %v", err)
	}

	versionString := api.NewStableVersion(2023, 1, 1).String()
	version, ok := versionsMap[versionString]
	if !ok {
		t.Fatalf("Expected version %s to be in versionsMap", versionString)
	}

	if version.Sunset == nil {
		t.Error("Expected sunset date to be set")
	} else if !version.Sunset.Equal(sunsetDate) {
		t.Errorf("Expected sunset date %v, got %v", sunsetDate, version.Sunset)
	}

	if !version.Deprecated {
		t.Error("Expected version to be deprecated when it has a sunset date")
	}
}

func TestAddContentTypeToVersion_NotDeprecated(t *testing.T) {
	versionsMap := make(map[string]*api.CommandVersion)

	extensions := map[string]any{}

	err := addContentTypeToVersion("application/vnd.atlas.2023-01-01+json", versionsMap, extensions, false)
	if err != nil {
		t.Fatalf("addContentTypeToVersion() error = %v", err)
	}

	versionString := api.NewStableVersion(2023, 1, 1).String()
	version, ok := versionsMap[versionString]
	if !ok {
		t.Fatalf("Expected version %s to be in versionsMap", versionString)
	}

	if version.Deprecated {
		t.Error("Expected version not to be deprecated when no deprecation indicators are present")
	}
}
