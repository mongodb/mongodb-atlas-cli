// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package oauth

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_patchConfigHostname(t *testing.T) {
	type fields struct {
		containerizedEnv string
		hostnameEnv      string
	}
	tests := []struct {
		name             string
		fields           fields
		expectedHostName string
	}{
		{
			name: "sets native hostname when no hostname env var is set",
			fields: fields{
				containerizedEnv: "",
				hostnameEnv:      "",
			},
			expectedHostName: config.NativeHostName,
		},
		{
			name: "sets container hostname when legacy containerized env var is set to true",
			fields: fields{
				containerizedEnv: "true",
				hostnameEnv:      "",
			},
			expectedHostName: config.DockerContainerHostName,
		},
		{
			name: "sets action hostname when valid action hostname env var is set",
			fields: fields{
				containerizedEnv: "",
				hostnameEnv:      config.GitHubActionsHostName,
			},
			expectedHostName: config.GitHubActionsHostName,
		},
		{
			name: "does not set hostname when invalid hostname env var is set",
			fields: fields{
				containerizedEnv: "",
				hostnameEnv:      "nonsense",
			},
			expectedHostName: config.NativeHostName,
		},
	}
	for _, tt := range tests {
		fields := tt.fields
		expectedHostName := tt.expectedHostName
		t.Run(tt.name, func(t *testing.T) {
			config.HostName = config.NativeHostName
			t.Setenv("MONGODB_ATLAS_IS_CONTAINERIZED", fields.containerizedEnv)
			t.Setenv("MONGODB_ATLAS_HOSTNAME", fields.hostnameEnv)
			patchConfigHostnameFromEnv()

			assert.Equal(t, expectedHostName, config.HostName)
		})
	}
}
