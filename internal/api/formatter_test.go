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

//go:build unit

package api

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsGoTemplate(t *testing.T) {
	tests := []struct {
		format       string
		isGoTemplate bool
	}{
		{
			format:       "json",
			isGoTemplate: false,
		},
		{
			format:       "csv",
			isGoTemplate: false,
		},
		{
			format:       "gzip",
			isGoTemplate: false,
		},
		{
			format:       "{foo}",
			isGoTemplate: false,
		},
		{
			format:       "{{bar}",
			isGoTemplate: false,
		},
		{
			format:       "{baz}}",
			isGoTemplate: false,
		},
		{
			format:       "{qux}}",
			isGoTemplate: false,
		},
		{
			format:       "}}quux{{",
			isGoTemplate: false,
		},
		{
			format:       "{{quuux}}",
			isGoTemplate: true,
		},
		{
			format:       "with prefix {{quuuux}}",
			isGoTemplate: true,
		},
		{
			format:       "with prefix {{quuuuux}} and suffix",
			isGoTemplate: true,
		},
		{
			format:       "{{quuuuuux}} with suffix",
			isGoTemplate: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			isGoTemplate := isGoTemplate(tt.format)
			require.Equal(t, tt.isGoTemplate, isGoTemplate)
		})
	}
}

func TestFormatJSON(t *testing.T) {
	const listClustersForAllProjectsJSON = `{"links":[{"href":"https://cloud-dev.mongodb.com/api/atlas/v2/clusters?envelope=false&includeCount=false&pretty=false&pageNum=1&itemsPerPage=1","rel":"self"},{"href":"https://cloud-dev.mongodb.com/api/atlas/v2/clusters?envelope=false&includeCount=false&pretty=false&itemsPerPage=1&pageNum=2","rel":"next"}],"results":[{"clusters":[{"alertCount":0,"atlasClusterType":"SERVERLESS","authEnabled":true,"availability":"available","backupEnabled":true,"clusterId":"6421b576bfa6c84992728bf7","dataSizeBytes":374626829,"isServerlessInstanceAboveFlexOpsThreshold":false,"isServerlessInstanceGrandfatheredToFlex":false,"name":"apix-atlascli-do-not-delete-e2e","nodeCount":-1,"serverlessPrivateEndpointId":null,"sslEnabled":true,"type":"sharded cluster","versions":["8.0.3"]}],"groupId":"5efda6aea3f2ed2e7dd6ce05","groupName":"Atlas CLI E2E","orgId":"5efda682a3f2ed2e7dd6cde4","orgName":"Atlas CLI E2E","planType":"Atlas","tags":[]}],"totalCount":2}`

	tests := []struct {
		name       string
		json       string
		template   string
		output     string
		shouldFail bool
	}{
		{
			name:       "listClustersForAllProjectsJson, valid 1",
			json:       listClustersForAllProjectsJSON,
			template:   "{{ (index (index .results 0).clusters 0).clusterId }}",
			output:     "6421b576bfa6c84992728bf7",
			shouldFail: false,
		},
		{
			name:       "listClustersForAllProjectsJson, valid 2",
			json:       listClustersForAllProjectsJSON,
			template:   "groupId: {{ index (index .results 0).groupId }}",
			output:     "groupId: 5efda6aea3f2ed2e7dd6ce05",
			shouldFail: false,
		},
		{
			name:       "invalid json",
			json:       "{notJson",
			template:   "groupId: {{ index (index .results 0).groupId }}",
			shouldFail: true,
		},
		{
			name:       "invalid template",
			json:       listClustersForAllProjectsJSON,
			template:   "groupId: {{ ** }}",
			shouldFail: true,
		},
		{
			name:       "incompatible response",
			json:       listClustersForAllProjectsJSON,
			template:   "foo: {{ .foo }}",
			output:     "foo: <no value>",
			shouldFail: false,
		},
	}

	for _, tt := range tests {
		formatter := NewFormatter()

		t.Run(tt.name, func(t *testing.T) {
			jsonReader := io.NopCloser(strings.NewReader(tt.json))
			output, err := formatter.formatJSON(tt.template, jsonReader)

			if tt.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				outputBytes, err := io.ReadAll(output)
				require.NoError(t, err)
				outputStr := string(outputBytes)

				require.Equal(t, tt.output, outputStr)
			}
		})
	}
}
