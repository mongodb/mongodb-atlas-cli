package main

import (
	"fmt"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
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
				Version:     "1991-05-17",
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
				Version:     "1991-05-17",
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
				"http-code": "200",
				"match": map[string]any{
					"values": "IDLE",
				},
			},
			expectedOutput: nil,
			expectedErr:    ErrWatcherMatchPropertiesPathIsMissing,
		},
		{
			input: map[string]any{
				"http-code": "200",
				"match": map[string]any{
					"path": "$.stateName",
				},
			},
			expectedOutput: nil,
			expectedErr:    ErrWatcherMatchPropertiesValuesAreMissing,
		},
		{
			input: map[string]any{
				"http-code": "200",
				"match": map[string]any{
					"path":   "$.stateName",
					"values": "IDLE",
				},
			},
			expectedOutput: &api.WatcherExpectProperties{
				HTTPCode: "200",
				Match: &api.WatcherMatchProperties{
					Path:   "$.stateName",
					Values: []string{"IDLE"},
				},
			},
			expectedErr: nil,
		},
		{
			input: map[string]any{
				"http-code": "200",
				"match": map[string]any{
					"path":   "$.stateName",
					"values": "IDLE,IDLE2",
				},
			},
			expectedOutput: &api.WatcherExpectProperties{
				HTTPCode: "200",
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
