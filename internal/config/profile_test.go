// Copyright 2020 MongoDB Inc
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

package config

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCLIConfigHome(t *testing.T) {
	expHome, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("os.UserConfigDir() unexpected error: %v", err)
	}
	home, err := CLIConfigHome()
	if err != nil {
		t.Fatalf("AtlasCLIConfigHome() unexpected error: %v", err)
	}
	expected := path.Join(expHome, "atlascli")
	if home != expected {
		t.Errorf("AtlasCLIConfigHome() = %s; want '%s'", home, expected)
	}
}

func TestConfig_IsTrue(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "true",
			want:  true,
		},
		{
			input: "True",
			want:  true,
		},
		{
			input: "t",
			want:  true,
		},
		{
			input: "T",
			want:  true,
		},
		{
			input: "TRUE",
			want:  true,
		},
		{
			input: "y",
			want:  true,
		},
		{
			input: "Y",
			want:  true,
		},
		{
			input: "yes",
			want:  true,
		},
		{
			input: "Yes",
			want:  true,
		},
		{
			input: "YES",
			want:  true,
		},
		{
			input: "1",
			want:  true,
		},
		{
			input: "false",
			want:  false,
		},
		{
			input: "f",
			want:  false,
		},
		{
			input: "unknown",
			want:  false,
		},
		{
			input: "0",
			want:  false,
		},
		{
			input: "",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			if got := IsTrue(tt.input); got != tt.want {
				t.Errorf("IsTrue() get: %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConfigHostname(t *testing.T) {
	type fields struct {
		containerizedEnv string
		atlasActionEnv   string
		ghActionsEnv     string
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
				atlasActionEnv:   "",
				ghActionsEnv:     "",
			},
			expectedHostName: NativeHostName,
		},
		{
			name: "sets container hostname when containerized env var is set",
			fields: fields{
				containerizedEnv: "true",
				atlasActionEnv:   "",
				ghActionsEnv:     "",
			},
			expectedHostName: "-|-|" + DockerContainerHostName,
		},
		{
			name: "sets atlas action hostname when containerized env var is set",
			fields: fields{
				containerizedEnv: "",
				atlasActionEnv:   "true",
				ghActionsEnv:     "",
			},
			expectedHostName: AtlasActionHostName + "|-|-",
		},
		{
			name: "sets github actions hostname when action env var is set",
			fields: fields{
				containerizedEnv: "",
				atlasActionEnv:   "",
				ghActionsEnv:     "true",
			},
			expectedHostName: "-|" + GitHubActionsHostName + "|-",
		},
		{
			name: "sets actions and containerized hostnames when both env vars are set",
			fields: fields{
				containerizedEnv: "true",
				atlasActionEnv:   "true",
				ghActionsEnv:     "true",
			},
			expectedHostName: AtlasActionHostName + "|" + GitHubActionsHostName + "|" + DockerContainerHostName,
		},
	}
	for _, tt := range tests {
		f := tt.fields
		expectedHostName := tt.expectedHostName
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(AtlasActionHostNameEnv, f.atlasActionEnv)
			t.Setenv(GitHubActionsHostNameEnv, f.ghActionsEnv)
			t.Setenv(ContainerizedHostNameEnv, f.containerizedEnv)
			actualHostName := getConfigHostnameFromEnvs()

			assert.Equal(t, expectedHostName, actualHostName)
		})
	}
}

func TestProfile_Rename(t *testing.T) {
	tests := []struct {
		name    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "default",
			wantErr: require.NoError,
		},
		{
			name:    "default-123",
			wantErr: require.NoError,
		},
		{
			name:    "default-test",
			wantErr: require.NoError,
		},
		{
			name:    "default.123",
			wantErr: require.Error,
		},
		{
			name:    "default.test",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Profile{
				name: tt.name,
				fs:   afero.NewMemMapFs(),
			}
			tt.wantErr(t, p.Rename(tt.name), fmt.Sprintf("Rename(%v)", tt.name))
		})
	}
}

func TestProfile_SetName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "default",
			wantErr: require.NoError,
		},
		{
			name:    "default-123",
			wantErr: require.NoError,
		},
		{
			name:    "default-test",
			wantErr: require.NoError,
		},
		{
			name:    "default.123",
			wantErr: require.Error,
		},
		{
			name:    "default.test",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Profile{
				name: tt.name,
				fs:   afero.NewMemMapFs(),
			}
			tt.wantErr(t, p.SetName(tt.name), fmt.Sprintf("SetName(%v)", tt.name))
		})
	}
}
