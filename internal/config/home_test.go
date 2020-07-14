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

package config

import (
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestConfig_configHome(t *testing.T) {
	t.Run("with XDG_CONFIG_HOME", func(t *testing.T) {
		xdgHome := "my_config"
		_ = os.Setenv("XDG_CONFIG_HOME", xdgHome)
		home, err := configHome()
		if err != nil {
			t.Fatalf("configHome() unexpected error: %v", err)
		}
		if home != xdgHome {
			t.Errorf("configHome() = %s; want '%s'", home, xdgHome)
		}
		_ = os.Unsetenv("XDG_CONFIG_HOME")
	})
	t.Run("without XDG_CONFIG_HOME", func(t *testing.T) {
		homedir.DisableCache = true
		_ = os.Setenv("HOME", ".")
		home, err := configHome()
		if err != nil {
			t.Fatalf("configHome() unexpected error: %v", err)
		}
		if home != "./.config" {
			t.Errorf("configHome() = %s; want './.config'", home)
		}
		homedir.DisableCache = false
	})
}
