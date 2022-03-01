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

	"github.com/spf13/afero"
)

func TestFile(t *testing.T) {
	t.Run("save latest version mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		err := s.SaveLatestVersion("mongocli", "1.2.3")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		path, _ := filePath("mongocli", stateFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		v, err := s.LoadLatestVersion("mongocli")
		if err != nil || v != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("save latest version atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		err := s.SaveLatestVersion("atlascli", "1.2.3")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load latest version atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		path, _ := filePath("atlascli", stateFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		v, err := s.LoadLatestVersion("atlascli")
		if err != nil || v != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("save brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		err := s.SaveBrewPath("mongocli", "a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		path, _ := filePath("mongocli", brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath("mongocli")
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("save brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		err := s.SaveBrewPath("atlascli", "a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewStore(appFS)

		path, _ := filePath("atlascli", brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath("atlascli")
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
}
