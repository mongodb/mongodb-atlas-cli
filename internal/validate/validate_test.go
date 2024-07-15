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

package validate

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestURL(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantErr bool
	}{
		{
			name:    "Valid URL",
			val:     "http://test.com/",
			wantErr: false,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: true,
		},
		{
			name:    "missing trailing slash",
			val:     "http://test.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := URL(val); (err != nil) != wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestOptionalObjectID(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantErr bool
	}{
		{
			name:    "Empty value",
			val:     "",
			wantErr: false,
		},
		{
			name:    "nil value",
			val:     nil,
			wantErr: false,
		},
		{
			name:    "Valid ObjectID",
			val:     "5e9f088b4797476aa0a5d56a",
			wantErr: false,
		},
		{
			name:    "Short ObjectID",
			val:     "5e9f088b4797476aa0a5d56",
			wantErr: true,
		},
		{
			name:    "Invalid ObjectID",
			val:     "5e9f088b4797476aa0a5d56z",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := OptionalObjectID(val); (err != nil) != wantErr {
				t.Errorf("OptionalObjectID() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestObjectID(t *testing.T) {
	tests := []struct {
		name    string
		val     string
		wantErr bool
	}{
		{
			name:    "Empty value",
			val:     "",
			wantErr: true,
		},
		{
			name:    "Valid ObjectID",
			val:     "5e9f088b4797476aa0a5d56a",
			wantErr: false,
		},
		{
			name:    "Short ObjectID",
			val:     "5e9f088b4797476aa0a5d56",
			wantErr: true,
		},
		{
			name:    "Invalid ObjectID",
			val:     "5e9f088b4797476aa0a5d56z",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := ObjectID(val); (err != nil) != wantErr {
				t.Errorf("OptionalObjectID() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestCredentials(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		if err := Credentials(); err == nil {
			t.Fatal("Credentials() expected an error\n")
		}
	})
	t.Run("with api key credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("PUBLIC_API_KEY", "test")
		t.Setenv("PRIVATE_API_KEY", "test")
		if err := Credentials(); err != nil {
			t.Fatalf("Credentials() unexpected error %v\n", err)
		}
	})
	t.Run("with auth token credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("ACCESS_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
		t.Setenv("REFRESH_TOKEN", "test")
		if err := Credentials(); err != nil {
			t.Fatalf("Credentials() unexpected error %v\n", err)
		}
	})
}

func TestNoAPIKeyss(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		if err := NoAPIKeys(); err != nil {
			t.Fatalf("NoAPIKeys() unexpected error %v\n", err)
		}
	})
	t.Run("with api key credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("PUBLIC_API_KEY", "test")
		t.Setenv("PRIVATE_API_KEY", "test")
		if err := NoAPIKeys(); err == nil {
			t.Fatalf("NoAPIKeys() expected error\n")
		}
	})
	t.Run("with auth token credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("ACCESS_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
		t.Setenv("REFRESH_TOKEN", "test")
		if err := NoAPIKeys(); err != nil {
			t.Fatalf("NoAPIKeys() unexpected error %v\n", err)
		}
	})
}

func TestNoAccessToken(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		if err := NoAccessToken(); err != nil {
			t.Fatalf("NoAccessToken() unexpected error %v\n", err)
		}
	})
	t.Run("with api key credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("PUBLIC_API_KEY", "test")
		t.Setenv("PRIVATE_API_KEY", "test")
		if err := NoAccessToken(); err != nil {
			t.Fatalf("NoAccessToken() unexpected error %v\n", err)
		}
	})
	t.Run("with auth token credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		t.Setenv("ACCESS_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
		t.Setenv("REFRESH_TOKEN", "test")
		if err := NoAccessToken(); err == nil {
			t.Fatalf("NoAccessToken() expected error\n")
		}
	})
}

func TestFlagInSlice(t *testing.T) {
	t.Parallel()
	type args struct {
		value       string
		flag        string
		validValues []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "value is present",
			args: args{
				value:       "test",
				flag:        "flag",
				validValues: []string{"test", "not-test"},
			},
			wantErr: false,
		},
		{
			name: "value is present",
			args: args{
				value:       "test",
				flag:        "flag",
				validValues: []string{"not-test"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		value := tt.args.value
		flag := tt.args.flag
		validValues := tt.args.validValues
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := FlagInSlice(value, flag, validValues); (err != nil) != wantErr {
				t.Errorf("FlagInSlice() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestOptionalURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		val     any
		wantErr bool
	}{
		{
			name:    "valid",
			val:     "https://test.com/",
			wantErr: false,
		},
		{
			name:    "empty",
			val:     "",
			wantErr: false,
		},
		{
			name:    "nil",
			val:     nil,
			wantErr: false,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := OptionalURL(val); (err != nil) != wantErr {
				t.Errorf("OptionalURL() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestPath(t *testing.T) {
	f, err := os.CreateTemp("", "TestPath")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
		require.NoError(t, os.Remove(f.Name()))
	})
	tests := []struct {
		name    string
		val     any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			val:     f.Name(),
			wantErr: require.NoError,
		},
		{
			name:    "empty",
			val:     "",
			wantErr: require.Error,
		},
		{
			name:    "nil",
			val:     nil,
			wantErr: require.Error,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantErr(t, Path(val))
		})
	}
}

func TestOptionalPath(t *testing.T) {
	f, err := os.CreateTemp("", "TestOptionalPath")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
		require.NoError(t, os.Remove(f.Name()))
	})
	tests := []struct {
		name    string
		val     any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			val:     f.Name(),
			wantErr: require.NoError,
		},
		{
			name:    "empty",
			val:     "",
			wantErr: require.NoError,
		},
		{
			name:    "nil",
			val:     nil,
			wantErr: require.NoError,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantErr(t, OptionalPath(val))
		})
	}
}

func TestClusterName(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid (single word)",
			val:     "Cluster0",
			wantErr: require.NoError,
		},
		{
			name:    "valid (dashed)",
			val:     "Cluster-0",
			wantErr: require.NoError,
		},
		{
			name:    "invalid (space)",
			val:     "Cluster 0",
			wantErr: require.Error,
		},
		{
			name:    "invalid (underscore)",
			val:     "Cluster_0",
			wantErr: require.Error,
		},
		{
			name:    "invalid (spacial char)",
			val:     "Cluster-ñ",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantErr(t, ClusterName(val))
		})
	}
}

func TestDBUsername(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid (single word)",
			val:     "admin",
			wantErr: require.NoError,
		},
		{
			name:    "valid (dashed)",
			val:     "admin-test",
			wantErr: require.NoError,
		},
		{
			name:    "valid (underscore)",
			val:     "admin_test",
			wantErr: require.NoError,
		},
		{
			name:    "invalid (space)",
			val:     "admin test",
			wantErr: require.Error,
		},
		{
			name:    "invalid (spacial char)",
			val:     "admin-ñ",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantErr(t, DBUsername(val))
		})
	}
}

func TestWeakPassword(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "valid password",
			val:     "password!@3!",
			wantErr: require.NoError,
		},
		{
			name:    "weak password",
			val:     "password",
			wantErr: require.Error,
		},
		{
			name:    "weak password",
			val:     "password1",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantErr(t, WeakPassword(val))
		})
	}
}
