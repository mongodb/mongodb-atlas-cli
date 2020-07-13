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
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	atlasProfile   = "atlas"
	newProfileName = "newProfileName"
)

var testConfigDir = "/test_dir"
var fs afero.Fs
var filename = fmt.Sprintf("%v/%v.toml", testConfigDir, ToolName)

func TestMain(m *testing.M) {
	fs = afero.NewMemMapFs()

	viper.SetFs(fs)
	SetConfigPath(testConfigDir)

	code := m.Run()
	os.Exit(code)
}

func createProfile(profileContents string) (*profile, error) {
	err := afero.WriteFile(fs, filename, []byte(profileContents), 0600)
	if err != nil {
		return nil, err
	}

	return &profile{
		name:      "default",
		configDir: testConfigDir,
		fs:        fs,
	}, nil
}

func createProfileWithJustDefaultUser() (*profile, error) {
	const contents = `
[default]
  base_url = "http://base_url.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
`
	return createProfile(contents)
}

func createProfileWithOneNonDefaultUser() (*profile, error) {
	const contents = `
[atlas]
  base_url = "http://base_url.com/"
  ops_manager_skip_verify = "false"
  org_id = "5cac6a2179358edabd12b572"
`
	return createProfile(contents)
}

func createProfileWithOneDefaultUserOneNonDefault() (*profile, error) {
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

func createProfileWithFullDescription() (*profile, error) {
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
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	availableProfiles := profile.List()
	assert.Equal(t, 2, len(availableProfiles), "expected to find 2 profiles")
}

func TestProfileDescribeFullProfile(t *testing.T) {
	profile, err := createProfileWithFullDescription()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	desc := profile.Get()
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
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	desc := profile.Get()
	assert.Equal(t, DefaultProfile, Name())
	assert.Equal(t, 4, len(desc))
}

func TestProfileDescribeWithNonDefaultProfile(t *testing.T) {
	profile, err := createProfileWithOneNonDefaultUser()
	if err != nil {
		t.Error(err)
	}

	profile.SetName(atlasProfile)

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	desc := profile.Get()
	assert.Equal(t, atlasProfile, profile.Name(), "expected atlas profile to be described")
	assert.Equal(t, 3, len(desc))

	assert.Equal(t, "5cac6a2179358edabd12b572", profile.OrgID(), "project id should match")
	assert.Equal(t, "", profile.ProjectID(), "project id should not be set")
}

func TestProfileDelete(t *testing.T) {
	profile, err := createProfileWithJustDefaultUser()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	if err := profile.Delete(); err != nil {
		t.Error(err)
	}

	availableProfiles := List()
	assert.Equal(t, 0, len(availableProfiles), "0 profiles should exist after deleting")
}

func TestProfileDeleteNonDefault(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	profile.SetName(atlasProfile)

	if err := profile.Delete(); err != nil {
		t.Error(err)
	}

	availableProfiles := List()
	assert.Equal(t, 1, len(availableProfiles), "1 profiles should exist after deleting")
	assert.Equal(t, DefaultProfile, availableProfiles[0], "the default profile should remain")
}

func TestProfileExists(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, profile.Exists(DefaultProfile), "default profile should exist")
	assert.Equal(t, true, profile.Exists(atlasProfile), "atlas profile should exist")
	assert.Equal(t, false, profile.Exists("not_a_profile"), "not_a_profile profile should not exist")
}

func TestProfileRename(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	profile.SetName(DefaultProfile)

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	defaultDescription := profile.Get()

	if err := profile.Rename(newProfileName); err != nil {
		t.Error(err)
	}

	profile.SetName(newProfileName)
	descriptionAfterRename := profile.Get()

	// after renaming, one profile should exist
	assert.Equal(t, false, profile.Exists(DefaultProfile), "default profile should not exist after rename")
	assert.Equal(t, true, profile.Exists(newProfileName), "new profile should exist after rename")
	assert.Equal(t, defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfileRenameOverwriteExisting(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	profile.SetName(DefaultProfile)

	defaultDescription := profile.Get()

	if err := profile.Rename(atlasProfile); err != nil {
		t.Error(err)
	}

	profile.SetName(atlasProfile)

	descriptionAfterRename := profile.Get()

	// after renaming, one profile should exist
	assert.Equal(t, false, profile.Exists(DefaultProfile), "default profile should not exist after rename")
	assert.Equal(t, true, profile.Exists(atlasProfile), "atlas profile should exist after rename")
	assert.Equal(t, defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfileSet(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := profile.Load(false); err != nil {
		t.Error(err)
	}

	profile.SetName(DefaultProfile)

	profile.Set(projectID, "newProjectId")

	assert.Equal(t, "newProjectId", ProjectID(), "project id should be set to new value")
}
