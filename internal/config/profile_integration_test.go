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

// +build integration

package config

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const (
	atlasProfile   = "atlas"
	newProfileName = "newProfileName"
)

func createProfile(profileContents string) error {
	p.fs = afero.NewMemMapFs()

	testConfigDir, err := configHome()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v/%v.toml", testConfigDir, ToolName)

	err = afero.WriteFile(p.fs, filename, []byte(profileContents), 0600)
	if err != nil {
		return err
	}

	return nil
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

func TestProfile_List(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	availableProfiles := List()
	assert.Len(t, availableProfiles, 2, "expected to find 2 profiles")
}

func TestProfile_Get_FullProfile(t *testing.T) {
	err := createProfileWithFullDescription()
	if err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := Get()

	a := assert.New(t)
	a.Equal("default", Name())
	a.Len(desc, 9)
	a.Equal("cloud", Service())
	a.Equal("some_public_key", PublicAPIKey())
	a.Equal("some_private_key", PrivateAPIKey())
	a.Equal("http://om_url.com/", OpsManagerURL())
	a.Equal("/path/to/certificate", OpsManagerCACertificate())
	a.Equal("false", OpsManagerSkipVerify())
}

func TestProfile_Get_Default(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := Get()

	a := assert.New(t)
	a.Equal(DefaultProfile, Name())
	a.Len(desc, 4)
}

func TestProfile_Get_NonDefault(t *testing.T) {
	err := createProfileWithOneNonDefaultUser()
	if err != nil {
		t.Error(err)
	}

	SetName(atlasProfile)

	if err := Load(); err != nil {
		t.Error(err)
	}

	desc := Get()

	a := assert.New(t)
	a.Equal(atlasProfile, Name(), "expected atlas profile to be described")
	a.Len(desc, 3)
	a.Equal("5cac6a2179358edabd12b572", OrgID(), "project id should match")
	a.Empty(ProjectID(), "project id should not be set")
}

func TestProfile_Delete_NonDefault(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
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
	a := assert.New(t)
	a.Len(availableProfiles, 1, "1 profiles should exist after deleting")
	a.Equal(DefaultProfile, availableProfiles[0], "the default profile should remain")
}

func TestProfile_Exists(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	a := assert.New(t)
	a.True(Exists(DefaultProfile), "default profile should exist")
	a.True(Exists(atlasProfile), "atlas profile should exist")
	a.False(Exists("not_a_profile"), "not_a_profile profile should not exist")
}

func TestProfile_Rename(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	if err := Load(); err != nil {
		t.Error(err)
	}

	defaultDescription := Get()

	if err := Rename(newProfileName); err != nil {
		t.Error(err)
	}

	SetName(newProfileName)
	descriptionAfterRename := Get()

	// after renaming, one profile should exist
	a := assert.New(t)
	a.False(Exists(DefaultProfile), "default profile should not exist after rename")
	a.True(Exists(newProfileName), "new profile should exist after rename")
	a.Equal(defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfile_Rename_OverwriteExisting(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	defaultDescription := Get()

	if err := Rename(atlasProfile); err != nil {
		t.Error(err)
	}

	SetName(atlasProfile)

	descriptionAfterRename := Get()

	// after renaming, one profile should exist
	a := assert.New(t)
	a.False(Exists(DefaultProfile), "default profile should not exist after rename")
	a.True(Exists(atlasProfile), "atlas profile should exist after rename")
	a.Equal(defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfile_Set(t *testing.T) {
	err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	SetName(DefaultProfile)

	Set(projectID, "newProjectId")

	assert.Equal(t, "newProjectId", ProjectID(), "project id should be set to new value")
}
