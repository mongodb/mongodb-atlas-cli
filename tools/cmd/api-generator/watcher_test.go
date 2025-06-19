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
	"fmt"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/require"
)

func TestNewWatcherGetProperties(t *testing.T) {
	tests := []struct {
		input          map[string]any
		expectedOutput *api.WatcherGetProperties
		expectedErr    error
	}{
		{
			input:          nil,
			expectedOutput: nil,
			expectedErr:    ErrWatcherGetPropertiesExtIsNil,
		},
		{
			input:          map[string]any{},
			expectedOutput: nil,
			expectedErr:    ErrWatcherGetPropertiesInvalidOperationID,
		},
		{
			input: map[string]any{
				"operation-id": "",
			},
			expectedOutput: nil,
			expectedErr:    ErrWatcherGetPropertiesInvalidOperationID,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Params:      map[string]string{},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
				"version":      123,
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Params:      map[string]string{},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
				"version":      "1991-05-17",
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Version:     api.StableVersion{Date: api.VersionDate{Year: 1991, Month: 5, Day: 17}},
				Params:      map[string]string{},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
				"version":      "1991-05-17",
				"params": map[string]any{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Version:     api.StableVersion{Date: api.VersionDate{Year: 1991, Month: 5, Day: 17}},
				Params: map[string]string{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
				"version":      "1991-05-17.upcoming",
				"params": map[string]any{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Version:     api.UpcomingVersion{Date: api.VersionDate{Year: 1991, Month: 5, Day: 17}},
				Params: map[string]string{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"operation-id": "getCluster",
				"version":      "preview",
				"params": map[string]any{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedOutput: &api.WatcherGetProperties{
				OperationID: "getCluster",
				Version:     api.PreviewVersion{},
				Params: map[string]string{
					"groupId":     "input:groupId",
					"clusterName": "body:$.name",
				},
			},
			expectedErr: nil,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("TestNewWatcherGetProperties[%v]", idx), func(t *testing.T) {
			output, err := newWatcherGetProperties(tt.input)

			if tt.expectedErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expectedOutput, output)
			} else {
				require.Nil(t, output)
				require.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestNewWatcherExpectProperties(t *testing.T) {
	tests := []struct {
		input          map[string]any
		expectedOutput *api.WatcherExpectProperties
		expectedErr    error
	}{
		{
			input: map[string]any{
				"http-code": 200,
				"match": map[string]any{
					"values": []any{
						"IDLE",
					},
				},
			},
			expectedOutput: nil,
			expectedErr:    ErrWatcherMatchPropertiesPathIsMissing,
		},
		{
			input: map[string]any{
				"http-code": 200,
				"match": map[string]any{
					"path": "$.stateName",
				},
			},
			expectedOutput: nil,
			expectedErr:    ErrWatcherMatchPropertiesValuesAreMissing,
		},
		{
			input: map[string]any{
				"http-code": 200,
				"match": map[string]any{
					"path": "$.stateName",
					"values": []any{
						"IDLE",
					},
				},
			},
			expectedOutput: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.stateName",
					Values: []string{"IDLE"},
				},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"http-code": 200,
				"match": map[string]any{
					"path": "$.stateName",
					"values": []any{
						"IDLE",
						"IDLE2",
					},
				},
			},
			expectedOutput: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.stateName",
					Values: []string{"IDLE", "IDLE2"},
				},
			},
			expectedErr: nil,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("TestNewWatcherExpectProperties[%v]", idx), func(t *testing.T) {
			output, err := newWatcherExpectProperties(tt.input)

			if tt.expectedErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expectedOutput, output)
			} else {
				require.Nil(t, output)
				require.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
