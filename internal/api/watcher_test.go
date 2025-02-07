package api

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
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
		expect            *WatcherExpectProperties
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
			expect: &WatcherExpectProperties{
				HTTPCode: 404,
			},
			commandRequest: CommandRequest{
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
				Match: &WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE"},
				},
			},
			commandRequest: CommandRequest{
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
				Match: &WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE", "DONE"},
				},
			},
			commandRequest: CommandRequest{
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
			},
			commandRequest: CommandRequest{
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
			},
			commandRequest: CommandRequest{
				Command: Command{
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
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
				Match: &WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE"},
				},
			},
			commandRequest: CommandRequest{
				Command: Command{
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
			expect: &WatcherExpectProperties{
				HTTPCode: 200,
				Match: &WatcherMatchProperties{
					Path:   "$.status",
					Values: []string{"IDLE", "DONE"},
				},
			},
			commandRequest: CommandRequest{
				Command: Command{
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
