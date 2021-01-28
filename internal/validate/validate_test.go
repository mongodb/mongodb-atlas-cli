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

// +build unit

package validate

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestURL(t *testing.T) {
	t.Run("Valid URL", func(t *testing.T) {
		err := URL("http://test.com/")
		if err != nil {
			t.Fatalf("URL() unexpected error %v\n", err)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		err := URL(1)
		if err == nil {
			t.Fatalf("URL() unexpected error %v\n", err)
		}
	})
	t.Run("missing trailing slash", func(t *testing.T) {
		err := URL("http://test.com")
		if err == nil {
			t.Fatalf("URL() unexpected error %v\n", err)
		}
	})
}

func TestOptionalObjectID(t *testing.T) {
	t.Run("Empty value", func(t *testing.T) {
		err := OptionalObjectID("")
		if err != nil {
			t.Fatalf("OptionalObjectID() unexpected error %v\n", err)
		}
	})
	t.Run("nil value", func(t *testing.T) {
		err := OptionalObjectID(nil)
		if err != nil {
			t.Fatalf("OptionalObjectID() unexpected error %v\n", err)
		}
	})
	t.Run("Valid ObjectID", func(t *testing.T) {
		err := OptionalObjectID("5e9f088b4797476aa0a5d56a")
		if err != nil {
			t.Fatalf("OptionalObjectID() unexpected error %v\n", err)
		}
	})
	t.Run("Short ObjectID", func(t *testing.T) {
		err := OptionalObjectID("5e9f088b4797476aa0a5d56")
		if err == nil {
			t.Fatal("OptionalObjectID() expected an error\n")
		}
	})
	t.Run("Invalid ObjectID", func(t *testing.T) {
		err := OptionalObjectID("5e9f088b4797476aa0a5d56z")
		if err == nil {
			t.Fatal("OptionalObjectID() expected an error\n")
		}
	})
}

func TestObjectID(t *testing.T) {
	t.Run("Valid ObjectID", func(t *testing.T) {
		err := ObjectID("5e9f088b4797476aa0a5d56a")
		if err != nil {
			t.Fatalf("ObjectID() unexpected error %v\n", err)
		}
	})
	t.Run("Short ObjectID", func(t *testing.T) {
		err := ObjectID("5e9f088b4797476aa0a5d56")
		if err == nil {
			t.Fatal("ObjectID() expected an error\n")
		}
	})
	t.Run("Invalid ObjectID", func(t *testing.T) {
		err := ObjectID("5e9f088b4797476aa0a5d56z")
		if err == nil {
			t.Fatal("ObjectID() expected an error\n")
		}
	})
}

func TestCredentials(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		err := Credentials()
		if err == nil {
			t.Fatal("Credentials() expected an error\n")
		}
	})
	t.Run("with credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		_ = os.Setenv("PUBLIC_API_KEY", "test")
		_ = os.Setenv("PRIVATE_API_KEY", "test")

		err := Credentials()
		if err != nil {
			t.Fatalf("Credentials() unexpected error %v\n", err)
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
		val     interface{}
		wantErr bool
	}{
		{
			name:    "valid",
			val:     "http://test.com/",
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
