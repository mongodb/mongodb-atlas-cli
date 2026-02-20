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

package vscode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVscode_BuildDeeplink(t *testing.T) {
	testCases := []struct {
		name           string
		uri            string
		deploymentName string
		deploymentType string
		telemetry      bool
		expected       string
	}{
		{
			name:           "atlas deployment with telemetry enabled",
			uri:            "mongodb+srv://testDeployment.abcdef.mongodb-qa.net",
			deploymentName: "testDeployment",
			deploymentType: "atlas",
			telemetry:      true,
			expected:       "vscode://mongodb.mongodb-vscode/connectWithURI?connectionString=mongodb%252Bsrv%253A%252F%252FtestDeployment.abcdef.mongodb-qa.net&name=testDeployment+%2528Atlas%2529&reuseExisting=true&utm_source=AtlasCLI",
		},
		{
			name:           "local deployment with telemetry disabled",
			uri:            "mongodb://localhost:11111/?directConnection=true",
			deploymentName: "testDeployment",
			deploymentType: "local",
			telemetry:      false,
			expected:       "vscode://mongodb.mongodb-vscode/connectWithURI?connectionString=mongodb%253A%252F%252Flocalhost%253A11111%252F%253FdirectConnection%253Dtrue&name=testDeployment+%2528Local%2529&reuseExisting=true",
		},
		{ //nolint:gosec // G101: false positive, this is a test URI with a literal password to verify special character encoding
			name:           "special characters present in password",
			uri:            "mongodb://user:123%%%@localhost:11111/?directConnection=true",
			deploymentName: "testDeployment",
			deploymentType: "local",
			telemetry:      false,
			expected:       "vscode://mongodb.mongodb-vscode/connectWithURI?connectionString=mongodb%253A%252F%252Fuser%253A123%2525%2525%2525%2540localhost%253A11111%252F%253FdirectConnection%253Dtrue&name=testDeployment+%2528Local%2529&reuseExisting=true",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			output := buildDeeplink(testCase.uri, testCase.deploymentName, testCase.deploymentType, testCase.telemetry)
			assert.Equal(t, testCase.expected, output)
		})
	}
}
