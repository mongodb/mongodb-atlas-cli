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

	"github.com/google/go-github/v61/github"
)

//go:generate mockgen -destination=../mocks/mock_release_version.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version ReleaseVersionDescriber

const (
	maxWaitTime = 500 * time.Millisecond
)

type Criteria func(tag string) bool

type ReleaseVersionDescriber interface {
	LatestWithCriteria(n int, matchCriteria Criteria) (*github.RepositoryRelease, error)
}

func NewReleaseVersionDescriber() ReleaseVersionDescriber {
	return &releaseVersionFetcher{ctx: context.Background()}
}

type releaseVersionFetcher struct {
	ctx context.Context
}

const (
	owner   = "mongodb"
	project = "mongodb-atlas-cli"
)

// LatestWithCriteria retrieves the first release version that matches the criteria. We assume that ListReleases returns releases sorted by created_at value.
func (s *releaseVersionFetcher) LatestWithCriteria(n int, matchCriteria Criteria) (*github.RepositoryRelease, error) {
	startTime := time.Now()
	client := github.NewClient(nil)

	opt := &github.ListOptions{PerPage: n}

	if matchCriteria == nil {
		matchCriteria = func(_ string) bool {
			return true
		}
	}

	for {
		releases, resp, err := client.Repositories.ListReleases(s.ctx, owner, project, opt)
		if err != nil {
			return nil, err
		}
		// Returns as soon as criteria is matched
		for _, release := range releases {
			if matchCriteria(release.GetTagName()) && !release.GetPrerelease() {
				return release, nil
			}
		}
		// Reached the end of the pages or timed out
		if resp.NextPage == 0 || time.Now().After(startTime.Add(maxWaitTime)) {
			return nil, nil
		}
		opt.Page = resp.NextPage
	}
}
