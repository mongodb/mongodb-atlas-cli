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

const digestAuth = "DigestAuth"

const rootTag = "Root"

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

func TestOperationRequiresAuthentication(t *testing.T) {
	rootSecurity := openapi3.SecurityRequirements{
		{digestAuth: []string{}},
	}
	explicitNoSecurity := openapi3.SecurityRequirements{}
	explicitSecurity := openapi3.SecurityRequirements{
		{"ServiceAccounts": []string{}},
	}

	tests := []struct {
		name      string
		operation *openapi3.Operation
		want      bool
	}{
		{
			name:      "inherits root security",
			operation: &openapi3.Operation{},
			want:      true,
		},
		{
			name: "operation explicitly disables security",
			operation: &openapi3.Operation{
				Security: &explicitNoSecurity,
			},
			want: false,
		},
		{
			name: "operation explicitly requires security",
			operation: &openapi3.Operation{
				Security: &explicitSecurity,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := operationRequiresAuthentication(tt.operation, rootSecurity)
			if got != tt.want {
				t.Fatalf("operationRequiresAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecToCommandsSetsUnauthenticatedFromOperationSecurity(t *testing.T) {
	noSecurity := openapi3.SecurityRequirements{}
	spec := &openapi3.T{
		Security: openapi3.SecurityRequirements{
			{digestAuth: []string{}},
		},
		Tags: openapi3.Tags{
			{Name: rootTag},
		},
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/atlas/v2/groups", &openapi3.PathItem{
				Get: testOperation("listGroups", nil),
			}),
			openapi3.WithPath("/api/atlas/v2/unauth/ephemeralClusters:create", &openapi3.PathItem{
				Post: testOperation("createEphemeralCluster", &noSecurity),
			}),
		),
	}

	groups, err := specToCommands(time.Now(), spec)

	if err != nil {
		t.Fatalf("specToCommands() error = %v", err)
	}
	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}

	commands := groups[0].Commands
	if len(commands) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(commands))
	}

	for _, command := range commands {
		switch command.OperationID {
		case "createEphemeralCluster":
			if !command.Unauthenticated {
				t.Fatal("expected createEphemeralCluster to be unauthenticated")
			}
		case "listGroups":
			if command.Unauthenticated {
				t.Fatal("expected listGroups to require authentication")
			}
		default:
			t.Fatalf("unexpected command %q", command.OperationID)
		}
	}
}

func testOperation(operationID string, security *openapi3.SecurityRequirements) *openapi3.Operation {
	return &openapi3.Operation{
		OperationID:  operationID,
		Tags:         []string{rootTag},
		Description:  operationID,
		Security:     security,
		Responses:    testResponses(),
		Parameters:   openapi3.Parameters{},
		RequestBody:  nil,
		ExternalDocs: nil,
	}
}

func testResponses() *openapi3.Responses {
	return openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{
			"application/vnd.atlas.2025-01-01+json": &openapi3.MediaType{},
		},
	}}))
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
