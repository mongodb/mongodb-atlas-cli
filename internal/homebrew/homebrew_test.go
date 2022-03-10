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
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"testing"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
)

func TestFile(t *testing.T) {
	t.Run("save brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		err := s.SaveBrewPath("a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath mcli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		path, _ := config.Path(brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath()
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath mcli is empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "mongocli")

		_, _, err := s.LoadBrewPath()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
	t.Run("save brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		err := s.SaveBrewPath("a/b/c", "d/e/f")
		if err != nil {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath atlascli", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		path, _ := config.Path(brewFileSubPath)
		_ = afero.WriteFile(appFS, path, []byte(""), 0600)

		p1, p2, err := s.LoadBrewPath()
		if err != nil || p1 != "" || p2 != "" {
			t.Errorf("LoadLatestVersion() unexpected error: %v", err)
		}
	})
	t.Run("load brewpath atlascli is empty", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		s := NewLoaderSaver(appFS, "atlascli")

		_, _, err := s.LoadBrewPath()
		if err == nil {
			t.Errorf("LoadLatestVersion() expected error: file not found")
		}
	})
}

func TestOutputOpts_testIsHomebrew(t *testing.T) {
	tests := []struct {
		tool string
		isHb bool
	}{
		{"atlascli", false},
		{"mongocli", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_ishomebrew_%v", tt.tool, tt.isHb), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := mocks.NewMockPathStore(ctrl)
			defer ctrl.Finish()

			mockStore.EXPECT().LoadBrewPath().Return("", "", nil)
			mockStore.EXPECT().SaveBrewPath(gomock.Any(), gomock.Any()).Return(nil)

			result := IsHomebrew(mockStore)
			if result != tt.isHb {
				t.Errorf("got = %v, want %v", result, tt.isHb)
			}
		})
	}
}
