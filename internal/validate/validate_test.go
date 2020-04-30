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
