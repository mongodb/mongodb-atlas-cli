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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_MongoCLIConfigHome(t *testing.T) {
	t.Run("with env set", func(t *testing.T) {
		expHome, err := os.UserConfigDir()
		expected := fmt.Sprintf("%s/mongocli", expHome)
		if err != nil {
			t.Fatalf("os.UserConfigDir() unexpected error: %v", err)
		}

		home, err := MongoCLIConfigHome()
		if err != nil {
			t.Fatalf("MongoCLIConfigHome() unexpected error: %v", err)
		}
		if home != expected {
			t.Errorf("MongoCLIConfigHome() = %s; want '%s'", home, expected)
		}
	})
}

func TestConfig_OldMongoCLIConfigHome(t *testing.T) {
	t.Run("old home with XDG_CONFIG_HOME", func(t *testing.T) {
		const xdgHome = "my_config"
		t.Setenv("XDG_CONFIG_HOME", xdgHome)
		home, err := OldMongoCLIConfigHome()
		if err != nil {
			t.Fatalf("OldMongoCLIConfigHome() unexpected error: %v", err)
		}
		if home != xdgHome {
			t.Errorf("MongoCLIConfigHome() = %s; want '%s'", home, xdgHome)
		}
	})
	t.Run("old home without XDG_CONFIG_HOME", func(t *testing.T) {
		t.Setenv("XDG_CONFIG_HOME", "")
		home, err := OldMongoCLIConfigHome()
		if err != nil {
			t.Fatalf("OldMongoCLIConfigHome() unexpected error: %v", err)
		}
		osHome, _ := os.UserHomeDir()
		if home != osHome+"/.config" {
			t.Errorf("OldMongoCLIConfigHome() = %s; want '%s/.config'", home, osHome)
		}
	})
}

func TestConfig_AtlasCLIConfigHome(t *testing.T) {
	t.Run("with env set", func(t *testing.T) {
		expHome, err := os.UserConfigDir()
		expected := fmt.Sprintf("%s/atlascli", expHome)
		if err != nil {
			t.Fatalf("os.UserConfigDir() unexpected error: %v", err)
		}

		home, err := AtlasCLIConfigHome()
		if err != nil {
			t.Fatalf("AtlasCLIConfigHome() unexpected error: %v", err)
		}
		if home != expected {
			t.Errorf("AtlasCLIConfigHome() = %s; want '%s'", home, expected)
		}
	})
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
		if got := IsTrue(tt.input); got != tt.want {
			t.Errorf("IsTrue() get: %v, want %v", got, tt.want)
		}
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
		fields := tt.fields
		expectedHostName := tt.expectedHostName
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(AtlasActionHostNameEnv, fields.atlasActionEnv)
			t.Setenv(GitHubActionsHostNameEnv, fields.ghActionsEnv)
			t.Setenv(ContainerizedHostNameEnv, fields.containerizedEnv)
			actualHostName := getConfigHostnameFromEnvs()

			assert.Equal(t, expectedHostName, actualHostName)
		})
	}
}

func TestConfig_SetName(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		require.NoError(t, SetName("default"))
		require.NoError(t, SetName("default-123"))
		require.NoError(t, SetName("default-test"))
	})

	t.Run("invalid", func(t *testing.T) {
		require.Error(t, SetName("d.efault"))
		require.Error(t, SetName("default.123"))
		require.Error(t, SetName("default.test"))
	})
}

func TestConfig_Rename(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		require.NoError(t, Rename("default"))
		require.NoError(t, Rename("default-123"))
		require.NoError(t, Rename("default-test"))
	})

	t.Run("invalid", func(t *testing.T) {
		require.Error(t, Rename("d.efault"))
		require.Error(t, Rename("default.123"))
		require.Error(t, Rename("default.test"))
	})
}
