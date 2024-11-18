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
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildQueryParameters(t *testing.T) {
	tests := []struct {
		name                   string
		commandQueryParameters []Parameter
		parameterValues        map[string][]string
		shouldFail             bool
		expectedValue          string
	}{
		{
			name:                   "empty parameters",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues:        map[string][]string{},
			shouldFail:             false,
			expectedValue:          "",
		},
		{
			name:                   "envelope set",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues: map[string][]string{
				"envelope": {"true"},
			},
			shouldFail:    false,
			expectedValue: "envelope=true",
		},
		{
			name:                   "single metric",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues: map[string][]string{
				"metrics": {"metric1"},
			},
			shouldFail:    false,
			expectedValue: "metrics=metric1",
		},
		{
			name:                   "multiple metrics",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues: map[string][]string{
				"metrics": {"metric1", "metric2"},
			},
			shouldFail:    false,
			expectedValue: "metrics=metric1&metrics=metric2",
		},
		{
			name:                   "multiple query parameters set",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues: map[string][]string{
				"envelope": {"true"},
				"metrics":  {"metric1", "metric2"},
			},
			shouldFail:    false,
			expectedValue: "envelope=true&metrics=metric1&metrics=metric2",
		},
		{
			name:                   "query encoding test",
			commandQueryParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.QueryParameters,
			parameterValues: map[string][]string{
				"period": {"when stars fade and dawn breaks - a moment both past and eternal"},
			},
			shouldFail:    false,
			expectedValue: "period=when+stars+fade+and+dawn+breaks+-+a+moment+both+past+and+eternal",
		},

		{
			name:                   "missing query parameters",
			commandQueryParameters: []Parameter{{Name: "required-test-query-parameter", Required: true, Type: ParameterType{IsArray: false, Type: TypeBool}}},
			parameterValues:        map[string][]string{},
			shouldFail:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryString, err := buildQueryParameters(tt.commandQueryParameters, tt.parameterValues)

			if tt.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedValue, queryString)
			}
		})
	}
}

func TestBuildPath(t *testing.T) {
	tests := []struct {
		name                 string
		pathTemplate         string
		commandURLParameters []Parameter
		parameterValues      map[string][]string
		shouldFail           bool
		expectedValue        string
	}{
		{
			name:                 "everything missing",
			pathTemplate:         getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URL,
			commandURLParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URLParameters,
			parameterValues:      map[string][]string{},
			shouldFail:           true,
			expectedValue:        "",
		},
		{
			name:                 "everything set",
			pathTemplate:         getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URL,
			commandURLParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URLParameters,
			parameterValues: map[string][]string{
				"groupId":        {"abcdef1234"},
				"clusterName":    {"cluster-01"},
				"clusterView":    {"view-42"},
				"databaseName":   {"metrics"},
				"collectionName": {"pageviews"},
			},
			shouldFail:    false,
			expectedValue: "/api/atlas/v2/groups/abcdef1234/clusters/cluster-01/view-42/metrics/pageviews/collStats/measurements",
		},
		{
			name:                 "groupId is missing",
			pathTemplate:         getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URL,
			commandURLParameters: getCollStatsLatencyNamespaceClusterMeasurementsCommand.RequestParameters.URLParameters,
			parameterValues: map[string][]string{
				"clusterName":    {"cluster-01"},
				"clusterView":    {"view-42"},
				"databaseName":   {"metrics"},
				"collectionName": {"pageviews"},
			},
			shouldFail:    true,
			expectedValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := buildPath(tt.pathTemplate, tt.commandURLParameters, tt.parameterValues)

			if tt.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedValue, path)
			}
		})
	}
}

func TestConvertToHttpRequest(t *testing.T) {
	tests := []struct {
		name                          string
		baseURL                       string
		request                       CommandRequest
		shouldFail                    bool
		expectedURL                   string
		expectedHTTPVerb              string
		expectedHTTPAcceptHeader      string
		expectedHTTPContentTypeHeader string
		expectedHTTPBody              string
	}{
		{
			name:    "valid get request (getCollStatsLatencyNamespaceClusterMeasurementsCommand)",
			baseURL: "https://base_url",
			request: CommandRequest{
				Command: getCollStatsLatencyNamespaceClusterMeasurementsCommand,
				Content: strings.NewReader(""),
				Format:  "json",
				Parameters: map[string][]string{
					"groupId":        {"abcdef1234"},
					"clusterName":    {"cluster-01"},
					"clusterView":    {"view-42"},
					"databaseName":   {"metrics"},
					"collectionName": {"pageviews"},
					"envelope":       {"true"},
					"metrics":        {"metric1", "metric2"},
				},
				Version: "2023-11-15",
			},
			shouldFail:                    false,
			expectedURL:                   "https://base_url/api/atlas/v2/groups/abcdef1234/clusters/cluster-01/view-42/metrics/pageviews/collStats/measurements?envelope=true&metrics=metric1&metrics=metric2",
			expectedHTTPVerb:              http.MethodGet,
			expectedHTTPAcceptHeader:      "application/vnd.atlas.2023-11-15+json",
			expectedHTTPContentTypeHeader: "",
			expectedHTTPBody:              "",
		},
		{
			name:    "valid post request (createClusterCommand)",
			baseURL: "http://another_base",
			request: CommandRequest{
				Command: createClusterCommand,
				Content: strings.NewReader(`{"very_pretty_json":true}`),
				Format:  "json",
				Parameters: map[string][]string{
					"groupId": {"0ff00ff00ff0"},
					"pretty":  {"true"},
				},
				Version: "2024-08-05",
			},
			shouldFail:                    false,
			expectedURL:                   "http://another_base/api/atlas/v2/groups/0ff00ff00ff0/clusters?pretty=true",
			expectedHTTPVerb:              http.MethodPost,
			expectedHTTPAcceptHeader:      "application/vnd.atlas.2024-08-05+json",
			expectedHTTPContentTypeHeader: "application/vnd.atlas.2024-08-05+json",
			expectedHTTPBody:              `{"very_pretty_json":true}`,
		},
		{
			name:    "valid post request, custom format (createClusterCommand), should work after CLOUDP-280747 is implemented",
			baseURL: "http://another_base",
			request: CommandRequest{
				Command: createClusterCommand,
				Content: strings.NewReader(`{"very_pretty_json":true}`),
				Format:  `{{ .Id }}`,
				Parameters: map[string][]string{
					"groupId": {"0ff00ff00ff0"},
					"pretty":  {"true"},
				},
				Version: "2024-08-05",
			},
			shouldFail:                    true, // TODO: should fail until CLOUDP-280747 is implemented
			expectedURL:                   "http://another_base/api/atlas/v2/groups/0ff00ff00ff0/clusters?pretty=true",
			expectedHTTPVerb:              http.MethodPost,
			expectedHTTPAcceptHeader:      "application/vnd.atlas.2024-08-05+json",
			expectedHTTPContentTypeHeader: "application/vnd.atlas.2024-08-05+json",
			expectedHTTPBody:              ``,
		},
		{
			name:    "invalid post request, missing groupId (createClusterCommand)",
			baseURL: "http://another_base",
			request: CommandRequest{
				Command: createClusterCommand,
				Content: strings.NewReader(`{"very_pretty_json":true}`),
				Format:  "json",
				Parameters: map[string][]string{
					"pretty": {"true"},
				},
				Version: "2024-08-05",
			},
			shouldFail: true,
		},
		{
			name:    "invalid post request, invalid version (createClusterCommand)",
			baseURL: "http://another_base",
			request: CommandRequest{
				Command: createClusterCommand,
				Content: strings.NewReader(`{"very_pretty_json":true}`),
				Format:  "json",
				Parameters: map[string][]string{
					"groupId": {"0ff00ff00ff0"},
					"pretty":  {"true"},
				},
				Version: "1991-05-17",
			},
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpRequest, err := ConvertToHTTPRequest(tt.baseURL, tt.request)

			if tt.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, tt.expectedURL, httpRequest.URL.String())
				require.Equal(t, tt.expectedHTTPVerb, httpRequest.Method)
				require.Equal(t, tt.expectedHTTPAcceptHeader, httpRequest.Header.Get("Accept"))
				require.Equal(t, tt.expectedHTTPContentTypeHeader, httpRequest.Header.Get("Content-Type"))

				bytes, err := io.ReadAll(httpRequest.Body)
				require.NoError(t, err)
				body := string(bytes)
				require.Equal(t, tt.expectedHTTPBody, body)
			}
		})
	}
}

// Please keep fixtures below this command
// Trying to keep this file readable.
var getCollStatsLatencyNamespaceClusterMeasurementsCommand = Command{
	OperationID: `getCollStatsLatencyNamespaceClusterMeasurements`,
	Description: `Get a list of the Coll Stats Latency cluster-level measurements for the given namespace.`,
	RequestParameters: RequestParameters{
		URL: `/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements`,
		QueryParameters: []Parameter{
			{
				Name:        `envelope`,
				Description: `Flag that indicates whether Application wraps the response in an envelope JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `bool`,
				},
			},
			{
				Name:        `metrics`,
				Description: `List that contains the metrics that you want to retrieve for the associated data series. If you don&#39;t set this parameter, this resource returns data series for all Coll Stats Latency metrics.`,
				Required:    false,
				Type: ParameterType{
					IsArray: true,
					Type:    `string`,
				},
			},
			{
				Name:        `start`,
				Description: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set period.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `end`,
				Description: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set period.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `period`,
				Description: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set start and end.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
		},
		URLParameters: []Parameter{
			{
				Name: `groupId`,
				Description: `Unique 24-hexadecimal digit string that identifies your project. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.


NOTE: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
				Required: true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `clusterName`,
				Description: `Human-readable label that identifies the cluster to retrieve metrics for.`,
				Required:    true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `clusterView`,
				Description: `Human-readable label that identifies the cluster topology to retrieve metrics for.`,
				Required:    true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `databaseName`,
				Description: `Human-readable label that identifies the database.`,
				Required:    true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
			{
				Name:        `collectionName`,
				Description: `Human-readable label that identifies the collection.`,
				Required:    true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
		},
		Verb: http.MethodGet,
	},
	Versions: []Version{
		{
			Version:            `2023-11-15`,
			RequestContentType: ``,
			ResponseContentTypes: []string{
				`json`,
			},
		},
	},
}

var createClusterCommand = Command{
	OperationID: `createCluster`,
	Description: `Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters.`,
	RequestParameters: RequestParameters{
		URL: `/api/atlas/v2/groups/{groupId}/clusters`,
		QueryParameters: []Parameter{
			{
				Name:        `envelope`,
				Description: `Flag that indicates whether Application wraps the response in an envelope JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `bool`,
				},
			},
			{
				Name:        `pretty`,
				Description: `Flag that indicates whether the response body should be in the prettyprint format.`,
				Required:    false,
				Type: ParameterType{
					IsArray: false,
					Type:    `bool`,
				},
			},
		},
		URLParameters: []Parameter{
			{
				Name: `groupId`,
				Description: `Unique 24-hexadecimal digit string that identifies your project. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.


NOTE: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
				Required: true,
				Type: ParameterType{
					IsArray: false,
					Type:    `string`,
				},
			},
		},
		Verb: http.MethodPost,
	},
	Versions: []Version{
		{
			Version:            `2023-01-01`,
			RequestContentType: `json`,
			ResponseContentTypes: []string{
				`json`,
			},
		},
		{
			Version:            `2023-02-01`,
			RequestContentType: `json`,
			ResponseContentTypes: []string{
				`json`,
			},
		},
		{
			Version:            `2024-08-05`,
			RequestContentType: `json`,
			ResponseContentTypes: []string{
				`json`,
			},
		},
		{
			Version:            `2024-10-23`,
			RequestContentType: `json`,
			ResponseContentTypes: []string{
				`json`,
			},
		},
	},
}
