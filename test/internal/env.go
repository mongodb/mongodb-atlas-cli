// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
)

const (
	cloudgov = "cloudgov"
)

func ProfileName() string {
	profileName := os.Getenv("E2E_PROFILE_NAME")
	if profileName != "" {
		return profileName
	}

	return "__e2e"
}

func ProfileProjectID() string {
	profile, err := ProfileData()
	if err != nil {
		return ""
	}
	return profile["project_id"]
}

func FlexInstanceName() (string, error) {
	instanceName, ok := os.LookupEnv("E2E_FLEX_INSTANCE_NAME")
	if !ok || instanceName == "" {
		return "", errors.New("environment variable is missing: E2E_FLEX_INSTANCE_NAME")
	}

	return instanceName, nil
}

func CloudRoleID() (string, error) {
	roleID, ok := os.LookupEnv("E2E_CLOUD_ROLE_ID")
	if !ok || roleID == "" {
		return "", errors.New("environment variable is missing: E2E_CLOUD_ROLE_ID")
	}

	return roleID, nil
}

func TestBucketName() (string, error) {
	bucketName, ok := os.LookupEnv("E2E_TEST_BUCKET")
	if !ok || bucketName == "" {
		return "", errors.New("environment variable is missing: E2E_TEST_BUCKET")
	}

	return bucketName, nil
}

func GCPCredentials() (string, error) {
	credentials, ok := os.LookupEnv("GCP_CREDENTIALS")
	if !ok || credentials == "" {
		return "", errors.New("environment variable is missing: GCP_CREDENTIALS")
	}

	return credentials, nil
}

func repoPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	parts := strings.Split(wd, fmt.Sprintf("%ctest%c", os.PathSeparator, os.PathSeparator))

	return parts[0], nil
}

func AtlasCLIBin() (string, error) {
	repo, err := repoPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(repo, "bin", "atlas"), nil
}

func ProfileData() (map[string]string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		cliPath,
		"config",
		"describe",
		ProfileName(),
		"-o=json",
	)

	cmd.Stderr = os.Stderr

	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var profile map[string]string
	if err := json.Unmarshal(buf, &profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func IsGov() bool {
	profile, err := ProfileData()
	if err != nil {
		return false
	}

	return profile["service"] == cloudgov
}

// LocalDevMDBVersionPrefix is the MongoDB version reported by the major-8 local deployment image.
const LocalDevMDBVersionPrefix = "8.3."

// LocalDevLatestMDBVersionPrefix is the MongoDB version reported by the untagged (latest) LocalDevImage.
const LocalDevLatestMDBVersionPrefix = "9.0."

func LocalDevImage() string {
	image, ok := os.LookupEnv("LOCALDEV_IMAGE")
	if !ok || image == "" {
		image = options.LocalDevImage
	}

	return image
}
