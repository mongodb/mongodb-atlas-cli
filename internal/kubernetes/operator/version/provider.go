// Copyright 2023 MongoDB Inc
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
	"fmt"
	"io"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
)

const (
	operatorRepositoryOrg = "mongodb"

	operatorRepository        = "mongodb-atlas-kubernetes"
	maxMajorVersionsSupported = 3
)

type AtlasOperatorVersionProvider interface {
	GetLatest() (string, error)
	IsSupported(version string) (bool, error)
	DownloadResource(ctx context.Context, version, path string) (io.ReadCloser, error)
}

type OperatorVersion struct {
	ghClient *github.Client
}

func (v *OperatorVersion) GetLatest() (string, error) {
	latest, _, err := v.ghClient.Repositories.GetLatestRelease(context.Background(), operatorRepositoryOrg, operatorRepository)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve latest version: %w", err)
	}

	return strings.Trim(*latest.Name, "v"), nil
}

func (v *OperatorVersion) IsSupported(version string) (bool, error) {
	latest, err := v.GetLatest()
	if err != nil {
		return false, err
	}

	latestSemVer, err := semver.NewVersion(latest)
	if err != nil {
		return false, fmt.Errorf("latest operator version %s is invalid", latest)
	}

	requestedOperatorVersionSem, err := semver.NewVersion(version)
	if err != nil {
		return false, fmt.Errorf("requested operator version %s is invalid", version)
	}

	if requestedOperatorVersionSem.Major() != latestSemVer.Major() {
		return false, nil
	}

	if requestedOperatorVersionSem.GreaterThan(latestSemVer) {
		return false, nil
	}

	if requestedOperatorVersionSem.LessThan(latestSemVer) {
		return latestSemVer.Minor()-requestedOperatorVersionSem.Minor() < maxMajorVersionsSupported, nil
	}

	return true, nil
}

func (v *OperatorVersion) DownloadResource(ctx context.Context, version, path string) (io.ReadCloser, error) {
	data, _, err := v.ghClient.Repositories.DownloadContents(
		ctx,
		operatorRepositoryOrg,
		operatorRepository,
		path,
		&github.RepositoryContentGetOptions{
			Ref: "v" + version,
		})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewOperatorVersion(ghClient *github.Client) *OperatorVersion {
	return &OperatorVersion{
		ghClient: ghClient,
	}
}
