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
//go:build !windows && !darwin
// +build !windows,!darwin

package telemetry

import (
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/homebrew"
	"github.com/spf13/afero"
)

func readPackageSource() *string {
	if c, err := homebrew.NewChecker(afero.NewOsFs()); err == nil && c.IsHomebrew() {
		s := "homebrew"
		return &s
	}

	if b, err := os.ReadFile("/etc/atlascli/package_source"); err == nil {
		s := string(b)
		return &s
	}

	return nil
}
