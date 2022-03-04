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

package latestrelease

import "github.com/google/go-github/v42/github"

type TestCase struct {
	tool             string
	currentVersion   string
	expectNewVersion bool
	release          *github.RepositoryRelease
}

func TestCases() []TestCase {
	f := false
	atlasV := "atlascli/v2.0.0"
	mcliV := "mongocli/v2.0.0"
	mcliOldV := "v2.0.0"

	tests := []TestCase{
		{
			tool:             "atlascli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "atlascli/v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "v3.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "v3.0.0-123",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "mongocli/v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v3.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v3.0.0-123",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &mcliV, Prerelease: &f, Draft: &f},
		},
	}
	return tests
}
