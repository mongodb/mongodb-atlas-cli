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

package version

import (
	"context"
	"time"

	"github.com/google/go-github/v38/github"
)

//go:generate mockgen -destination=../mocks/mock_release_version.go -package=mocks github.com/mongodb/mongocli/internal/version ReleaseVersionDescriber

const (
	maxPageSize = 100
	maxWaitTime = 1 * time.Second
)

type ReleaseVersionDescriber interface {
	AllVersions() ([]*github.RepositoryRelease, error)
}

func NewReleaseVersionDescriber() ReleaseVersionDescriber {
	return &releaseVersionFetcher{ctx: context.Background()}
}

type releaseVersionFetcher struct {
	ctx context.Context
}

// NVersions retrieves the first n versions returned by Github list releases endpoint.
func (s *releaseVersionFetcher) NVersions(n int) ([]*github.RepositoryRelease, error) {
	var releases []*github.RepositoryRelease
	var page = 1
	var pageSize = maxPageSize

	if n <= maxPageSize {
		pageSize = n
	}

	startTime := time.Now()
	client := github.NewClient(nil)

	for {
		slice, _, err := client.Repositories.ListReleases(s.ctx, owner, mongoCLI, &github.ListOptions{PerPage: pageSize, Page: page})
		releases = append(releases, slice...)

		if err != nil {
			return releases, err
		}

		if len(slice) < pageSize || startTime.Add(maxWaitTime).After(time.Now()) {
			return releases, nil
		}
		page++
	}
}

// NVersions retrieves all versions returned by Github list releases endpoint.
func (s *releaseVersionFetcher) AllVersions() ([]*github.RepositoryRelease, error) {
	return s.NVersions(maxPageSize)
}
