// Copyright 2020 MongoDB Inc
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
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

var testConfigDir = ""

const atlasProfile = "atlas"
const newProfileName = "newProfileName"

func TestMain(m *testing.M) {
	var err error
	testConfigDir, err = ioutil.TempDir("", "")
	if err != nil {
		log.Fatalf("failed to set up for test %v", err)
	}

	SetConfigPath(testConfigDir)
	fmt.Printf("set config path to %v\n", testConfigDir)

	code := m.Run()

	if err := os.RemoveAll(testConfigDir); err != nil {
		log.Fatalf("failed to clean up after test: %v", err)
	}

	os.Exit(code)
}

func createProfile(profileContents string) error {
	return ioutil.WriteFile(fmt.Sprintf("%v/%v.toml", testConfigDir, ToolName), []byte(profileContents), 0600)
}

func createProfileWithJustDefaultUser() error {
	const contents = `
[default]
  base_url = "http://base_url.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
`
	return createProfile(contents)
}

func createProfileWithOneNonDefaultUser() error {
	const contents = `
[atlas]
  base_url = "http://base_url.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
`
	return createProfile(contents)
}

func createProfileWithOneDefaultUserOneNonDefault() error {
	const contents = `
[default]
  base_url = "http://base_url.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
  profile_id = "5cac6a2179358edabd12b572"

[atlas]
  base_url = "http://atlas.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
`
	return createProfile(contents)
}

func createProfileWithFullDescription() error {
	const contents = `
[default]
  base_url = "http://base_url.com/"
  ops_manager_url = "http://om_url.com/"
  ops_manager_ca_certificate = "/path/to/certificate"
  ops_manager_skip_verify = "false"
  public_api_key = "some_public_key"
  private_api_key = "some_private_key"
  org_id = "5cac6a2179358edabd12b572"
  profile_id = "5cac6a2179358edabd12b572"
  service = "cloud"
`
	return createProfile(contents)
}

func TestProfileList(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	availableProfiles := List()
	assert.Equal(t, 2, len(availableProfiles), "expected to find 2 profiles")
}

func TestProfileDescribeFullProfile(t *testing.T) {
	if err := createProfileWithFullDescription(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := GetConfigDescription()
	assert.Equal(t, "default", Name())
	assert.Equal(t, 9, len(desc))
	assert.Equal(t, "cloud", Service())
	assert.Equal(t, "some_public_key", PublicAPIKey())
	assert.Equal(t, "some_private_key", PrivateAPIKey())
	assert.Equal(t, "http://om_url.com/", OpsManagerURL())
	assert.Equal(t, "/path/to/certificate", OpsManagerCACertificate())
	assert.Equal(t, "false", OpsManagerSkipVerify())
}

func TestProfileDescribeWithOneDefaultProfile(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := GetConfigDescription()
	assert.Equal(t, DefaultProfile, Name())
	assert.Equal(t, 4, len(desc))
}

func TestProfileDescribeWithNonDefaultProfile(t *testing.T) {
	if err := createProfileWithOneNonDefaultUser(); err != nil {
		t.Error(err)
	}

	SetName(atlasProfile)

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := GetConfigDescription()
	assert.Equal(t, atlasProfile, Name(), "expected atlas profile to be described")
	assert.Equal(t, 3, len(desc))

	assert.Equal(t, "5cac6a2179358edabd12b572", OrgID(), "project id should match")
	assert.Equal(t, "", ProjectID(), "project id should not be set")
}

func TestProfileDelete(t *testing.T) {
	if err := createProfileWithJustDefaultUser(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	if err := Delete(); err != nil {
		t.Error(err)
	}

	// TODO :: Fix this
	viper.Reset()

	availableProfiles := List()
	assert.Equal(t, 0, len(availableProfiles), "0 profiles should exist after deleting")
}

func TestProfileDeleteNonDefault(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	SetName(atlasProfile)

	if err := Delete(); err != nil {
		t.Error(err)
	}

	availableProfiles := List()
	assert.Equal(t, 1, len(availableProfiles), "1 profiles should exist after deleting")
	assert.Equal(t, DefaultProfile, availableProfiles[0], "the default profile should remain")
}

func TestProfileExists(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, Exists(DefaultProfile), "default profile should exist")
	assert.Equal(t, true, Exists(atlasProfile), "atlas profile should exist")
	assert.Equal(t, false, Exists("not_a_profile"), "not_a_profile profile should not exist")
}

func TestProfileRename(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	if err := Load(); err != nil {
		t.Error(err)
	}

	defaultDescription := GetConfigDescription()

	if err := Rename(newProfileName); err != nil {
		t.Error(err)
	}

	SetName(newProfileName)
	descriptionAfterRename := GetConfigDescription()

	// after renaming, one profile should exist
	assert.Equal(t, false, Exists(DefaultProfile), "default profile should not exist after rename")
	assert.Equal(t, true, Exists(newProfileName), "new profile should exist after rename")
	assert.Equal(t, defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfileRenameOverwriteExisting(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	defaultDescription := GetConfigDescription()

	if err := Rename(atlasProfile); err != nil {
		t.Error(err)
	}

	SetName(atlasProfile)

	descriptionAfterRename := GetConfigDescription()

	// after renaming, one profile should exist
	assert.Equal(t, false, Exists(DefaultProfile), "default profile should not exist after rename")
	assert.Equal(t, true, Exists(atlasProfile), "atlas profile should exist after rename")
	assert.Equal(t, defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfileSet(t *testing.T) {
	if err := createProfileWithOneDefaultUserOneNonDefault(); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	Set(projectID, "newProjectId")

	assert.Equal(t, "newProjectId", ProjectID(), "project id should be set to new value")
}
