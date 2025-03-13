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

package api

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/require"
)

func TestAnyToString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		// Test nil
		{
			name:     "nil value",
			input:    nil,
			expected: "",
		},

		// Test basic types
		{
			name:     "string value",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "integer value",
			input:    42,
			expected: "42",
		},
		{
			name:     "negative integer",
			input:    -123,
			expected: "-123",
		},
		{
			name:     "float value",
			input:    3.14,
			expected: "3.14",
		},
		{
			name:     "negative float",
			input:    -0.5,
			expected: "-0.5",
		},
		{
			name:     "boolean true",
			input:    true,
			expected: "true",
		},
		{
			name:     "boolean false",
			input:    false,
			expected: "false",
		},

		// Test slices
		{
			name:     "empty slice",
			input:    []string{},
			expected: "",
		},
		{
			name:     "string slice with one element",
			input:    []string{"first"},
			expected: "first",
		},
		{
			name:     "string slice with multiple elements",
			input:    []string{"first", "second", "third"},
			expected: "first",
		},
		{
			name:     "integer slice",
			input:    []int{1, 2, 3},
			expected: "1",
		},
		{
			name:     "float slice",
			input:    []float64{1.1, 2.2, 3.3},
			expected: "1.1",
		},
		{
			name:     "boolean slice",
			input:    []bool{true, false},
			expected: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := anyToString(tt.input)
			if result != tt.expected {
				t.Errorf("anyToString(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestApplyJsonPath(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		jsonPath    string
		expected    string
		expectError bool
	}{
		{
			name:        "basic object with string",
			json:        `{"name": "Cluster0", "status": "IDLE"}`,
			jsonPath:    "$.name",
			expected:    "Cluster0",
			expectError: false,
		},
		{
			name:        "basic object with different field",
			json:        `{"name": "Cluster0", "status": "IDLE"}`,
			jsonPath:    "$.status",
			expected:    "IDLE",
			expectError: false,
		},
		{
			name:        "object with integer",
			json:        `{"id": 123, "name": "Test"}`,
			jsonPath:    "$.id",
			expected:    "123",
			expectError: false,
		},
		{
			name:        "object with boolean",
			json:        `{"active": true, "name": "Test"}`,
			jsonPath:    "$.active",
			expected:    "true",
			expectError: false,
		},
		{
			name:        "array first element",
			json:        `{"tags": ["dev", "prod", "staging"]}`,
			jsonPath:    "$.tags[0]",
			expected:    "dev",
			expectError: false,
		},
		{
			name:        "nested object",
			json:        `{"cluster": {"name": "Cluster0", "status": "IDLE"}}`,
			jsonPath:    "$.cluster.name",
			expected:    "Cluster0",
			expectError: false,
		},
		{
			name:        "invalid JSON",
			json:        `{"name": "Cluster0", "status": "IDLE"`, // Missing closing brace
			jsonPath:    "$.name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid JSONPath",
			json:        `{"name": "Cluster0", "status": "IDLE"}`,
			jsonPath:    "$.[invalid",
			expected:    "",
			expectError: true,
		},
		{
			name:        "non-existent path",
			json:        `{"name": "Cluster0", "status": "IDLE"}`,
			jsonPath:    "$.nonexistent",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty JSON object",
			json:        `{}`,
			jsonPath:    "$.name",
			expected:    "",
			expectError: true,
		},
		{
			name:        "array of objects first element",
			json:        `{"clusters": [{"name": "Cluster0"}, {"name": "Cluster1"}]}`,
			jsonPath:    "$.clusters[0].name",
			expected:    "Cluster0",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ReadCloser from the JSON string
			reader := io.NopCloser(strings.NewReader(tt.json))

			result, err := applyJSONPath(reader, tt.jsonPath)

			// Check error expectations
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
				return
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// If we're not expecting an error, check the result
			if !tt.expectError {
				if result != tt.expected {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

// TestApplyJsonPathReaderError tests the case where the reader fails.
func TestApplyJsonPathReaderError(t *testing.T) {
	// Create a mock reader that always fails
	failingReader := &mockFailingReader{
		err: io.ErrClosedPipe,
	}

	_, err := applyJSONPath(failingReader, "$.name")
	if err == nil {
		t.Error("expected error from failing reader, got nil")
	}
}

// mockFailingReader implements io.ReadCloser and always returns an error.
type mockFailingReader struct {
	err error
}

func (m *mockFailingReader) Read(_ []byte) (n int, err error) {
	return 0, m.err
}

func (*mockFailingReader) Close() error {
	return nil
}

// MockCommandExecutor implements CommandExecutor interface for testing.
type MockCommandExecutor struct {
	response *CommandResponse
	err      error
}

func (m *MockCommandExecutor) ExecuteCommand(_ context.Context, _ CommandRequest) (*CommandResponse, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

func TestWatchInner(t *testing.T) {
	tests := []struct {
		name              string
		executor          *MockCommandExecutor
		expect            *api.WatcherExpectProperties
		commandRequest    CommandRequest
		expectedCompleted bool
		expectedErr       error
	}{
		{
			name: "Delete cluster - waiting for 404",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: false,
					HTTPCode:  404,
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 404,
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: true,
			expectedErr:       nil,
		},
		{
			name: "Create cluster - wait for IDLE status",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: true,
					HTTPCode:  200,
					Output:    io.NopCloser(strings.NewReader(`{"status": "IDLE"}`)),
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE"},
				},
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: true,
			expectedErr:       nil,
		},
		{
			name: "Create cluster - wait for IDLE or DONE status",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: true,
					HTTPCode:  200,
					Output:    io.NopCloser(strings.NewReader(`{"status": "DONE"}`)),
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE", "DONE"},
				},
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: true,
			expectedErr:       nil,
		},
		{
			name: "Executor returns error",
			executor: &MockCommandExecutor{
				err: errors.New("execution failed"),
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: false,
			expectedErr:       ErrWatcherFailedToExecuteWatchRequest,
		},
		{
			name: "HTTP code mismatch",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: false,
					HTTPCode:  400,
					Output:    io.NopCloser(strings.NewReader(`{}`)),
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: false,
			expectedErr:       nil,
		},
		{
			name: "Default HTTP code (200)",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: true,
					HTTPCode:  200,
					Output:    io.NopCloser(strings.NewReader(`{}`)),
				},
			},
			expect: nil,
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: true,
			expectedErr:       nil,
		},
		{
			name: "Invalid JSON in response",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: true,
					HTTPCode:  200,
					Output:    io.NopCloser(strings.NewReader(`invalid json`)),
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE"},
				},
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: false,
			expectedErr:       ErrWatcherFailedToApplyJSONPathToWatcherResponse,
		},
		{
			name: "Status not in expected values",
			executor: &MockCommandExecutor{
				response: &CommandResponse{
					IsSuccess: true,
					HTTPCode:  200,
					Output:    io.NopCloser(strings.NewReader(`{"status": "CREATING"}`)),
				},
			},
			expect: &api.WatcherExpectProperties{
				HTTPCode: 200,
				Match: &api.WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE", "DONE"},
				},
			},
			commandRequest: CommandRequest{
				Command: api.Command{
					OperationID: "GetCluster",
				},
			},
			expectedCompleted: false,
			expectedErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			result, err := watchInner(ctx, tt.executor, tt.expect, tt.commandRequest)

			// Check error
			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error containing %v, got nil", tt.expectedErr)
					return
				}
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("expected error containing %v, got %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Check result
			if result != tt.expectedCompleted {
				t.Errorf("got result %v, want %v", result, tt.expectedCompleted)
			}
		})
	}
}

func TestBuildRequestParameters(t *testing.T) {
	// Test setup
	validJSON := []byte(`{
		"data": {
			"id": "abc123",
			"nested": {
				"value": "def456"
			},
			"array": ["item1", "item2"]
		}
	}`)

	invalidJSON := []byte(`{invalid json`)

	tests := []struct {
		name          string
		requestParams map[string][]string
		responseBody  []byte
		watcherParams map[string]string
		expected      map[string][]string
		expectedError error
	}{
		{
			name: "Success - All parameter types",
			requestParams: map[string][]string{
				"ProjectID": {"abcdef"},
			},
			responseBody: validJSON,
			watcherParams: map[string]string{
				"fromBody":   "body:$.data.id",
				"fromInput":  "input:ProjectID",
				"fromConst":  "const:fixed-value",
				"fromNested": "body:$.data.nested.value",
			},
			expected: map[string][]string{
				"fromBody":   {"abc123"},
				"fromInput":  {"abcdef"},
				"fromConst":  {"fixed-value"},
				"fromNested": {"def456"},
			},
			expectedError: nil,
		},
		{
			name:          "Error - Invalid parameter format (missing colon)",
			requestParams: map[string][]string{},
			responseBody:  validJSON,
			watcherParams: map[string]string{
				"invalid": "bodyno-colon",
			},
			expected:      nil,
			expectedError: ErrWatcherFailedToBuildRequestParametersInvalidParameter,
		},
		{
			name:          "Error - Invalid parameter function",
			requestParams: map[string][]string{},
			responseBody:  validJSON,
			watcherParams: map[string]string{
				"invalid": "unknown:value",
			},
			expected:      nil,
			expectedError: ErrWatcherFailedToBuildRequestInvalidParameterOperation,
		},
		{
			name:          "Error - Input parameter not found",
			requestParams: map[string][]string{},
			responseBody:  validJSON,
			watcherParams: map[string]string{
				"missing": "input:NonExistentParam",
			},
			expected:      nil,
			expectedError: ErrWatcherFailedToBuildRequestInputParameterNotFound,
		},
		{
			name:          "Error - Invalid JSONPath",
			requestParams: map[string][]string{},
			responseBody:  validJSON,
			watcherParams: map[string]string{
				"invalid": "body:$.invalid.path",
			},
			expected:      nil,
			expectedError: ErrWatcherFailedToBuildRequestFailedToApplyJSONPath,
		},
		{
			name: "Success - Empty response body with no body parameters",
			requestParams: map[string][]string{
				"ProjectID": {"abc123"},
			},
			responseBody: []byte{},
			watcherParams: map[string]string{
				"fromInput": "input:ProjectID",
				"fromConst": "const:fixed-value",
			},
			expected: map[string][]string{
				"fromInput": {"abc123"},
				"fromConst": {"fixed-value"},
			},
			expectedError: nil,
		},
		{
			name: "Success - Invalid JSON with no body parameters",
			requestParams: map[string][]string{
				"ProjectID": {"abc123"},
			},
			responseBody: invalidJSON,
			watcherParams: map[string]string{
				"fromInput": "input:ProjectID",
				"fromConst": "const:fixed-value",
			},
			expected: map[string][]string{
				"fromInput": {"abc123"},
				"fromConst": {"fixed-value"},
			},
			expectedError: nil,
		},
		{
			name:          "Error - Invalid JSON with body parameter",
			requestParams: map[string][]string{},
			responseBody:  invalidJSON,
			watcherParams: map[string]string{
				"fromBody": "body:$.data.id",
			},
			expected:      nil,
			expectedError: ErrWatcherFailedToBuildRequestFailedToApplyJSONPath,
		},
		{
			name: "Success - Multiple input values",
			requestParams: map[string][]string{
				"MultiValue": {"value1", "value2", "value3"},
			},
			responseBody: validJSON,
			watcherParams: map[string]string{
				"fromInput": "input:MultiValue",
			},
			expected: map[string][]string{
				"fromInput": {"value1", "value2", "value3"},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := buildRequestParameters(tt.requestParams, tt.responseBody, tt.watcherParams)

			// Check error
			if tt.expectedError != nil {
				require.Error(t, actualErr)
				require.ErrorIs(t, actualErr, tt.expectedError)
			} else {
				require.NoError(t, actualErr)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}
