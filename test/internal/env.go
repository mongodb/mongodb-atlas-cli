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
	cloudgov                   = "cloudgov"
	snapshotCloudRoleID        = "c0123456789abcdef012345c"
	snapshotTestBucket         = "test-bucket"
	snapshotFlexInstanceName   = "test-flex"
	snapshotIdentityProviderID = "d0123456789abcdef012345d"
	snapshotOrgID              = "a0123456789abcdef012345a"
	snapshotProjectID          = "b0123456789abcdef012345b"
	snapshotOpsManagerURL      = "http://localhost:8080/"
)

type TestMode string

const (
	TestModeLive   TestMode = "live"   // run tests against a live Atlas instance (this do not replay or record snapshots)
	TestModeRecord TestMode = "record" // record snapshots
	TestModeReplay TestMode = "replay" // replay snapshots
)

func TestRunMode() (TestMode, error) {
	mode := os.Getenv("TEST_MODE")
	if mode == "" || strings.EqualFold(mode, "live") {
		return TestModeLive, nil
	}

	if strings.EqualFold(mode, "record") {
		return TestModeRecord, nil
	}

	if strings.EqualFold(mode, "replay") {
		return TestModeReplay, nil
	}

	return TestModeLive, fmt.Errorf("invalid value for environment variable TEST_MODE: %s, expected 'live', 'record' or 'replay'", mode)
}

func ProfileName() string {
	profileName := os.Getenv("E2E_PROFILE_NAME")
	if profileName != "" {
		return profileName
	}

	mode, err := TestRunMode()
	if err != nil || mode != TestModeReplay {
		return "__e2e"
	}

	return "__e2e_snapshot"
}

func ProfileProjectID() string {
	profile, err := ProfileData()
	if err != nil {
		return ""
	}
	return profile["project_id"]
}

func SkipCleanup() bool {
	mode, err := TestRunMode()
	if err != nil {
		return false
	}

	if mode == TestModeLive {
		return false
	}

	return true
}

func IdentityProviderID() (string, error) {
	mode, err := TestRunMode()
	if err == nil && mode == TestModeReplay {
		return snapshotIdentityProviderID, nil
	}

	idpID, ok := os.LookupEnv("IDENTITY_PROVIDER_ID")
	if !ok || idpID == "" {
		return "", errors.New("environment variable is missing: IDENTITY_PROVIDER_ID")
	}

	return idpID, nil
}

func FlexInstanceName() (string, error) {
	mode, err := TestRunMode()
	if err == nil && mode == TestModeReplay {
		return snapshotFlexInstanceName, nil
	}

	instanceName, ok := os.LookupEnv("E2E_FLEX_INSTANCE_NAME")
	if !ok || instanceName == "" {
		return "", errors.New("environment variable is missing: E2E_FLEX_INSTANCE_NAME")
	}

	return instanceName, nil
}

func CloudRoleID() (string, error) {
	mode, err := TestRunMode()
	if err == nil && mode == TestModeReplay {
		return snapshotCloudRoleID, nil
	}

	roleID, ok := os.LookupEnv("E2E_CLOUD_ROLE_ID")
	if !ok || roleID == "" {
		return "", errors.New("environment variable is missing: E2E_CLOUD_ROLE_ID")
	}

	return roleID, nil
}

func TestBucketName() (string, error) {
	mode, err := TestRunMode()
	if err == nil && mode == TestModeReplay {
		return snapshotTestBucket, nil
	}

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

	parts := strings.Split(wd, "/test/")

	return parts[0], nil
}

func AtlasCLIBin() (string, error) {
	repo, err := repoPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(repo, "bin", "atlas"), nil
}

func snapshotBasePath() (string, error) {
	repo, err := repoPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(repo, "test", "e2e", "testdata", ".snapshots"), nil
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

func LocalDevImage() string {
	image, ok := os.LookupEnv("LOCALDEV_IMAGE")
	if !ok || image == "" {
		image = options.LocalDevImage
	}

	return image
}
