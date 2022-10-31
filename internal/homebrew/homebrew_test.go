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

package homebrew

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestChecker_IsHomebrew(t *testing.T) {
	tests := []struct {
		paths      *homebrew
		isHomebrew bool
	}{
		{
			paths: &homebrew{
				CheckedAt:      time.Now(),
				ExecutablePath: "/workplace/mongocli/bin/mongocli",
				FormulaPath:    "/opt/homebrew/Cellar/mongocli/1.22.0",
			},
			isHomebrew: false,
		},
		{
			paths: &homebrew{
				CheckedAt:      time.Now(),
				ExecutablePath: "",
				FormulaPath:    "/opt/homebrew/Cellar/mongocli/1.22.0",
			},
			isHomebrew: false,
		},
		{
			paths: &homebrew{
				CheckedAt:      time.Now(),
				ExecutablePath: "/workplace/mongocli/bin/mongocli",
				FormulaPath:    "",
			},
			isHomebrew: false,
		},
		{
			paths: &homebrew{
				CheckedAt:      time.Now(),
				ExecutablePath: "/workplace/mongocli/bin/mongocli",
				FormulaPath:    ".",
			},
			isHomebrew: false,
		},
		{
			paths: &homebrew{
				CheckedAt:      time.Now(),
				ExecutablePath: "/opt/homebrew/Cellar/mongocli/1.22.0/bin",
				FormulaPath:    "/opt/homebrew/Cellar/mongocli/1.22.0",
			},
			isHomebrew: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("path:%v/formula:%v", tt.paths.ExecutablePath, tt.paths.FormulaPath), func(t *testing.T) {
			appFS := afero.NewMemMapFs()
			c, err := NewChecker(appFS)
			if err != nil {
				t.Errorf("NewChecker() unexpected error: %v", err)
			}

			err = c.save(tt.paths)
			if err != nil {
				t.Errorf("save() unexpected error: %v", err)
			}

			result := c.IsHomebrew()
			if result != tt.isHomebrew {
				t.Errorf("got = %v, want %v", result, tt.isHomebrew)
			}
		})
	}
}
