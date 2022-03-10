// Copyright 2022 MongoDB Inc
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
// +build unit

package latestrelease

import (
	"testing"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
)

func TestFile(t *testing.T) {
	t.Run("save latest version mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		err := s.SaveLatestVersion("1.2.3")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		path, _ := config.Path(stateFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		v, err := s.LoadLatestVersion()
		if err != nil || v != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version mongocli empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		_, err := s.LoadLatestVersion()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
	t.Run("save latest version atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		err := s.SaveLatestVersion("1.2.3")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		path, _ := config.Path(stateFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		v, err := s.LoadLatestVersion()
		if err != nil || v != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version atlascli empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		_, err := s.LoadLatestVersion()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
}
