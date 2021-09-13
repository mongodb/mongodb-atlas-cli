// Copyright 2021 MongoDB Inc
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

package store

import (
	"context"

	"github.com/google/go-github/v38/github"
)

//go:generate mockgen -destination=../mocks/mock_version.go -package=mocks github.com/mongodb/mongocli/internal/store VersionDescriber

type VersionDescriber interface {
	LatestVersion() (string, error)
}

// LatestVersion encapsulates the logic to manage different cloud providers.
func (s *Store) LatestVersion() (string, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "mongodb", "mongocli")
	if err != nil {
		return "", err
	}
	return release.GetTagName(), nil
}
