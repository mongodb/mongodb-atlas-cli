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
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestAddParameters(t *testing.T) {
	testCmd := &cobra.Command{}
	err := addParameters(testCmd, []api.Parameter{
		{Name: "test1", Description: "description1", Required: false, Type: api.ParameterType{IsArray: false, Type: api.TypeString}},
		{Name: "test2", Description: "description2", Required: true, Type: api.ParameterType{IsArray: true, Type: api.TypeBool}},
	})

	require.NoError(t, err)

	test1 := testCmd.Flag("test1")
	require.NotNil(t, test1)
	require.Equal(t, "test1", test1.Name)
	require.Equal(t, "description1", test1.Usage)
	require.Equal(t, "string", test1.Value.Type())
	_, required1 := test1.Annotations[cobra.BashCompOneRequiredFlag]
	require.False(t, required1)

	test2 := testCmd.Flag("test2")
	require.NotNil(t, test2)
	require.Equal(t, "test2", test2.Name)
	require.Equal(t, "description2", test2.Usage)
	require.Equal(t, "boolSlice", test2.Value.Type())
	_, required2 := test2.Annotations[cobra.BashCompOneRequiredFlag]
	require.True(t, required2)
}

func TestAddParametersDuplicates(t *testing.T) {
	testCmd := &cobra.Command{}
	err := addParameters(testCmd, []api.Parameter{
		{Name: "test1", Description: "description1", Required: false, Type: api.ParameterType{IsArray: false, Type: api.TypeString}},
		{Name: "test1", Description: "description2", Required: true, Type: api.ParameterType{IsArray: true, Type: api.TypeBool}},
	})
	require.Error(t, err)
}

type fixedValueProvider struct{}

func (*fixedValueProvider) ValueForFlag(flagName string) (*string, error) {
	if flagName == "b" {
		return nil, nil
	}

	return pointer.Get("default-for-" + flagName), nil
}

func TestSetUnTouchedFlags(t *testing.T) {
	// Create cobra command with flags
	testCmd := &cobra.Command{}
	testCmd.Flags().String("a", "a", "will remain unchanged by test")
	testCmd.Flags().String("b", "b", "will remain unchanged by test, it also won't be changed by the fixedValueProvider")
	testCmd.Flags().String("c", "c", "will be changed by test")
	testCmd.Flags().String("d", "d", "will remain unchanged by test")
	testCmd.Flags().String("e", "e", "will be changed by test")

	// Parse flags, this will also set `Changed`
	// Adding test here in case cobra behavior changes
	require.NoError(t, testCmd.ParseFlags([]string{"--c", "user-changed-c", "--e", "user-changed-e"}))

	// This is the state before PreRunE
	require.Equal(t, "a", testCmd.Flag("a").Value.String())
	require.Equal(t, "b", testCmd.Flag("b").Value.String())
	require.Equal(t, "user-changed-c", testCmd.Flag("c").Value.String())
	require.Equal(t, "d", testCmd.Flag("d").Value.String())
	require.Equal(t, "user-changed-e", testCmd.Flag("e").Value.String())

	require.False(t, testCmd.Flag("a").Changed)
	require.False(t, testCmd.Flag("b").Changed)
	require.True(t, testCmd.Flag("c").Changed)
	require.False(t, testCmd.Flag("d").Changed)
	require.True(t, testCmd.Flag("e").Changed)

	// The actual method we're testing
	require.NoError(t, setUnTouchedFlags(&fixedValueProvider{}, testCmd))

	// Verify that untouched values have been updated with their respective new values
	require.Equal(t, "default-for-a", testCmd.Flag("a").Value.String())
	require.Equal(t, "b", testCmd.Flag("b").Value.String())
	require.Equal(t, "user-changed-c", testCmd.Flag("c").Value.String())
	require.Equal(t, "default-for-d", testCmd.Flag("d").Value.String())
	require.Equal(t, "user-changed-e", testCmd.Flag("e").Value.String())
	require.True(t, testCmd.Flag("a").Changed)
	require.False(t, testCmd.Flag("b").Changed)
	require.True(t, testCmd.Flag("c").Changed)
	require.True(t, testCmd.Flag("d").Changed)
	require.True(t, testCmd.Flag("e").Changed)
}

func TestPrintDeprecatedWarning(t *testing.T) {
	tests := []struct {
		name        string
		apiCommand  api.Command
		version     string
		shouldPrint bool
		expectedMsg string
	}{
		{
			name: "deprecated version without sunset",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
						Sunset:     nil,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: false,
						Sunset:     nil,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: true,
			expectedMsg: "warning: version '2023-01-01' is deprecated. Consider upgrading to a newer version: 2024-01-01.\n",
		},
		{
			name: "deprecated version with sunset (should not print separate warning)",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "non-deprecated version",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: false,
						Sunset:     nil,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "version not found",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: true,
						Sunset:     nil,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "invalid version string",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
						Sunset:     nil,
					},
				},
			},
			version:     "invalid-version",
			shouldPrint: false,
		},
		{
			name: "version not deprecated and no sunset date",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "all versions deprecated should not print warning",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{Version: api.NewStableVersion(2023, 1, 1), Deprecated: true, Sunset: nil},
					{Version: api.NewStableVersion(2024, 1, 1), Deprecated: true, Sunset: nil},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stderr output
			oldStderr := os.Stderr
			r, w, err := os.Pipe()
			require.NoError(t, err)
			os.Stderr = w

			// Read output in a goroutine to avoid blocking
			outputChan := make(chan string, 1)
			go func() {
				buf := make([]byte, 1024)
				n, _ := r.Read(buf)
				outputChan <- string(buf[:n])
			}()

			printDeprecatedVersionWarning(tt.apiCommand, &tt.version)

			w.Close()
			os.Stderr = oldStderr

			// Get captured output
			output := <-outputChan

			if tt.shouldPrint {
				require.Contains(t, output, tt.expectedMsg, "Expected deprecation warning to be printed")
			} else {
				require.Empty(t, output, "Expected no output for non-deprecated or invalid cases")
			}
		})
	}
}

func TestPrintDeprecatedWarningWithSunset(t *testing.T) {
	tests := []struct {
		name        string
		apiCommand  api.Command
		version     string
		shouldPrint bool
		expectedMsg string
	}{
		{
			name: "version with future sunset date",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Sunset:     nil,
						Deprecated: false,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: true,
			expectedMsg: "warning: version '2023-01-01' is deprecated for this command and will be sunset on 2026-01-15. Consider upgrading to a newer version if available.",
		},
		{
			name: "version with past sunset date",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Sunset:     pointer.Get(time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Sunset:     nil,
						Deprecated: false,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: true,
			expectedMsg: "error: version '2023-01-01' is deprecated for this command and has already been sunset since 2020-01-15. Consider upgrading to a newer version if available.",
		},
		{
			name: "version without sunset date",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
						Sunset:  nil,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "version not found",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
						Deprecated: true,
					},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "invalid version string",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
						Deprecated: true,
					},
				},
			},
			version:     "invalid-version",
			shouldPrint: false,
		},
		{
			name: "all versions with sunset date should not print warning",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{Version: api.NewStableVersion(2023, 1, 1), Deprecated: true, Sunset: nil},
					{Version: api.NewStableVersion(2024, 1, 1), Deprecated: true, Sunset: nil},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
		{
			name: "all versions with sunset date and deprecated should not print warning",
			apiCommand: api.Command{
				Versions: []api.CommandVersion{
					{Version: api.NewStableVersion(2023, 1, 1), Deprecated: true, Sunset: pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))},
					{Version: api.NewStableVersion(2024, 1, 1), Deprecated: true, Sunset: nil},
				},
			},
			version:     "2023-01-01",
			shouldPrint: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stderr output
			oldStderr := os.Stderr
			r, w, err := os.Pipe()
			require.NoError(t, err)
			os.Stderr = w

			// Read output in a goroutine to avoid blocking
			outputChan := make(chan string, 1)
			var output string
			go func() {
				var buf bytes.Buffer
				_, err = io.Copy(&buf, r)
				require.NoError(t, err) //nolint:testifylint // this is a test
				output = buf.String()
				outputChan <- output
			}()

			printDeprecatedVersionWarning(tt.apiCommand, &tt.version)

			w.Close()
			os.Stderr = oldStderr

			// Wait for goroutine to finish
			<-outputChan

			if tt.shouldPrint {
				require.Contains(t, output, tt.expectedMsg, "Expected sunset warning to be printed")
			} else {
				require.Empty(t, output, "Expected no output for versions without sunset or invalid cases")
			}
		})
	}
}

func TestConvertAPIToCobraCommand(t *testing.T) {
	tests := []struct {
		name          string
		command       api.Command
		expectError   bool
		expectedError error
		validate      func(t *testing.T, cmd *cobra.Command)
	}{
		{
			name: "basic command creation",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				Aliases:     []string{"getInstance"},
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				require.Equal(t, "getServerlessInstance", cmd.Use)
				require.Equal(t, []string{"getInstance"}, cmd.Aliases)
				require.Equal(t, "getServerlessInstance", cmd.Annotations["operationId"])
				require.NotNil(t, cmd.PreRunE)
				require.NotNil(t, cmd.RunE)
				// Check that flags are added
				require.NotNil(t, cmd.Flag(flag.Version))
				require.Nil(t, cmd.Flag(flag.File))
				require.NotNil(t, cmd.Flag(flag.Output))
			},
		},
		{
			name: "command with ShortOperationID",
			command: api.Command{
				OperationID:      "getGroupServerlessInstance",
				ShortOperationID: "getServerlessInstance",
				Description:      "Returns one serverless instance from the specified project.",
				Aliases:          []string{"getInstance"},
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				require.Equal(t, "getServerlessInstance", cmd.Use)
				// Original operation ID should be in aliases
				require.Contains(t, cmd.Aliases, "getGroupServerlessInstance")
				require.Contains(t, cmd.Aliases, "getInstance")
				require.Equal(t, "getGroupServerlessInstance", cmd.Annotations["operationId"])
			},
		},
		{
			name: "command with deprecated version",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				require.NotNil(t, cmd.PreRunE)
				// Command should be marked as deprecated since all versions are deprecated
				require.NotEmpty(t, cmd.Deprecated)
				require.Contains(t, cmd.Deprecated, "all of the available endpoint versions have been deprecated")
				require.Contains(t, cmd.Deprecated, "2023-01-01")
				require.Contains(t, cmd.Deprecated, "2026-01-15")
			},
		},
		{
			name: "command with all versions deprecated (no sunset)",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				// Command should be marked as deprecated
				require.NotEmpty(t, cmd.Deprecated)
				require.Contains(t, cmd.Deprecated, "all of the available endpoint versions have been deprecated")
			},
		},
		{
			name: "command with mixed versions (not all deprecated)",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
						Sunset:     pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: false,
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				// Command should NOT be marked as deprecated since not all versions are deprecated
				require.Empty(t, cmd.Deprecated)
			},
		},
		{
			name: "command with multiple versions",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
					},
					{
						Version: api.NewStableVersion(2024, 1, 1),
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				require.NotNil(t, cmd.Flag(flag.Version))
				// Version flag should have choices for both versions
				versionFlag := cmd.Flag(flag.Version)
				require.NotNil(t, versionFlag)
			},
		},
		{
			name: "command with no versions",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{},
			},
			expectError:   true,
			expectedError: ErrAPICommandsHasNoVersions,
		},
		{
			name: "command with request body (needs file flag)",
			command: api.Command{
				OperationID: "createServerlessInstance",
				Description: "Creates one serverless instance in the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless",
					Verb: "http.MethodPost",
				},
				Versions: []api.CommandVersion{
					{
						Version:            api.NewStableVersion(2023, 1, 1),
						RequestContentType: "json",
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				// File flag should be present when request body is needed
				require.NotNil(t, cmd.Flag(flag.File))
			},
		},
		{
			name: "command with watcher",
			command: api.Command{
				OperationID: "getServerlessInstance",
				Description: "Returns one serverless instance from the specified project.",
				RequestParameters: api.RequestParameters{
					URL:  "/api/atlas/v2/groups/{groupId}/serverless/{name}",
					Verb: "http.MethodGet",
				},
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
					},
				},
				Watcher: &api.WatcherProperties{
					Get: api.WatcherGetProperties{
						OperationID: "getServerlessInstance",
						Version:     api.NewStableVersion(2023, 1, 1),
						Params:      map[string]string{"groupId": "groupId", "name": "name"},
					},
				},
			},
			expectError: false,
			validate: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()
				// Watch flags should be present when watcher is configured
				require.NotNil(t, cmd.Flag(flag.EnableWatch))
				require.NotNil(t, cmd.Flag(flag.WatchTimeout))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := convertAPIToCobraCommand(tt.command)

			if tt.expectError {
				require.Error(t, err)
				if tt.expectedError != nil {
					require.Equal(t, tt.expectedError, err)
				}
				require.Nil(t, cmd)
			} else {
				require.NoError(t, err)
				require.NotNil(t, cmd)
				if tt.validate != nil {
					tt.validate(t, cmd)
				}
			}
		})
	}
}

func TestAllVersionsDeprecated(t *testing.T) {
	tests := []struct {
		name     string
		command  api.Command
		expected bool
	}{
		{
			name: "all versions have sunset dates",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
						Sunset:  pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
					{
						Version: api.NewStableVersion(2024, 1, 1),
						Sunset:  pointer.Get(time.Date(2027, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			expected: true,
		},
		{
			name: "all versions are deprecated",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: true,
					},
				},
			},
			expected: true,
		},
		{
			name: "mix of sunset and deprecated",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
						Sunset:  pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: true,
					},
				},
			},
			expected: true,
		},
		{
			name: "one version not deprecated",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: false,
						Sunset:     nil,
					},
				},
			},
			expected: false,
		},
		{
			name: "no versions deprecated",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: false,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: false,
					},
				},
			},
			expected: false,
		},
		{
			name: "empty versions",
			command: api.Command{
				Versions: []api.CommandVersion{},
			},
			expected: true, // Empty is considered all deprecated (edge case)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := allVersionsDeprecated(tt.command)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestAddDeprecationMessageIfNeeded(t *testing.T) {
	tests := []struct {
		name            string
		command         api.Command
		shouldDeprecate bool
		expectedParts   []string // Parts that should be contained in the message
	}{
		{
			name: "single version with sunset",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
						Sunset:  pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			shouldDeprecate: true,
			expectedParts: []string{
				"all of the available endpoint versions have been deprecated",
				"2023-01-01",
				"2026-01-15",
				"sunset date",
			},
		},
		{
			name: "multiple versions with sunset",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version: api.NewStableVersion(2023, 1, 1),
						Sunset:  pointer.Get(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
					{
						Version: api.NewStableVersion(2024, 1, 1),
						Sunset:  pointer.Get(time.Date(2027, 1, 15, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			shouldDeprecate: true,
			expectedParts: []string{
				"all of the available endpoint versions have been deprecated",
				"2023-01-01",
				"2026-01-15",
				"sunset date",
			},
		},
		{
			name: "all versions deprecated without sunset",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: true,
					},
				},
			},
			shouldDeprecate: true,
			expectedParts: []string{
				"all of the available endpoint versions have been deprecated",
			},
		},
		{
			name: "not all versions deprecated",
			command: api.Command{
				Versions: []api.CommandVersion{
					{
						Version:    api.NewStableVersion(2023, 1, 1),
						Deprecated: true,
					},
					{
						Version:    api.NewStableVersion(2024, 1, 1),
						Deprecated: false,
					},
				},
			},
			shouldDeprecate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			addDeprecationMessageIfNeeded(cmd, tt.command)

			if tt.shouldDeprecate {
				require.NotEmpty(t, cmd.Deprecated)
				require.Contains(t, cmd.Deprecated, "all of the available endpoint versions have been deprecated")
				for _, part := range tt.expectedParts {
					require.Contains(t, cmd.Deprecated, part, "Expected deprecation message to contain: %s", part)
				}
			} else {
				require.Empty(t, cmd.Deprecated)
			}
		})
	}
}
