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

package homebrew

import (
	"testing"

	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
)

func TestFile(t *testing.T) {
	t.Run("save brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "mongocli")

		err := s.SaveBrewPath("a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "mongocli")

		path, _ := file.Path("mongocli", brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath()
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath mcli is empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "mongocli")

		_, _, err := s.LoadBrewPath()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
	t.Run("save brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "atlascli")

		err := s.SaveBrewPath("a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "mongocli")

		path, _ := file.Path("atlascli", brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath()
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath atlascli is empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewPathStore(appFS, "mongocli")

		_, _, err := s.LoadBrewPath()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
}
