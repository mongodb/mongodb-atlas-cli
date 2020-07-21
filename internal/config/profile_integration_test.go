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

func createProfile(profileContents string) (*profile, error) {
	fs := afero.NewMemMapFs()

	testConfigDir := "/test"
	filename := fmt.Sprintf("%v/%v.toml", testConfigDir, ToolName)

	p := profile{
		name:      DefaultProfile,
		configDir: testConfigDir,
		fs:        fs,
	}

	err := afero.WriteFile(fs, filename, []byte(profileContents), 0600)
	return &p, err
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

func TestProfile_Get_FullProfile(t *testing.T) {
	profile, err := createProfileWithFullDescription()
	if err != nil {
		t.Fatal(err)
	}

	profile.SetName(DefaultProfile)

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	desc := profile.Get()

	a := assert.New(t)
	a.Equal("default", profile.Name())
	a.Len(desc, 9)
	a.Equal("cloud", profile.Service())
	a.Equal("some_public_key", profile.PublicAPIKey())
	a.Equal("some_private_key", profile.PrivateAPIKey())
	a.Equal("http://om_url.com/", profile.OpsManagerURL())
	a.Equal("/path/to/certificate", profile.OpsManagerCACertificate())
	a.Equal("false", profile.OpsManagerSkipVerify())
}

func TestProfile_Get_Default(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Fatal(err)
	}

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	desc := profile.Get()

	a := assert.New(t)
	a.Equal(DefaultProfile, profile.Name())
	a.Len(desc, 4)
}

func TestProfile_Get_NonDefault(t *testing.T) {
	profile, err := createProfileWithOneNonDefaultUser()
	if err != nil {
		t.Fatal(err)
	}

	profile.SetName(atlasProfile)

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	desc := profile.Get()

	a := assert.New(t)
	a.Equal(atlasProfile, profile.Name(), "expected atlas profile to be described")
	a.Len(desc, 3)
	a.Equal("5cac6a2179358edabd12b572", profile.OrgID(), "project id should match")
	a.Empty(profile.ProjectID(), "project id should not be set")
}

func TestProfile_Delete_NonDefault(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Fatal(err)
	}

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	profile.SetName(atlasProfile)

	if err := profile.Delete(); err != nil {
		t.Fatal(err)
	}

	desc := profile.Get()
	a := assert.New(t)
	a.Len(desc, 0, "profile should have no properties")
}

func TestProfile_Rename(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Fatal(err)
	}

	profile.SetName(DefaultProfile)

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	defaultDescription := profile.Get()

	if err := profile.Rename(newProfileName); err != nil {
		t.Fatal(err)
	}

	profile.SetName(newProfileName)
	descriptionAfterRename := profile.Get()

	// after renaming, one profile should exist
	a := assert.New(t)
	a.False(Exists(DefaultProfile), "default profile should not exist after rename")
	a.True(Exists(newProfileName), "new profile should exist after rename")
	a.Equal(defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfile_Rename_OverwriteExisting(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Fatal(err)
	}

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	profile.SetName(DefaultProfile)

	defaultDescription := profile.Get()

	if err := profile.Rename(atlasProfile); err != nil {
		t.Fatal(err)
	}

	profile.SetName(atlasProfile)

	descriptionAfterRename := profile.Get()

	// after renaming, one profile should exist
	a := assert.New(t)
	a.Equal(defaultDescription, descriptionAfterRename, "descriptions should be equal after renaming")
}

func TestProfile_Set(t *testing.T) {
	profile, err := createProfileWithOneDefaultUserOneNonDefault()
	if err != nil {
		t.Fatal(err)
	}

	if err := profile.Load(false); err != nil {
		t.Fatal(err)
	}

	profile.SetName(DefaultProfile)

	profile.Set(projectID, "newProjectId")

	assert.Equal(t, "newProjectId", profile.ProjectID(), "project id should be set to new value")
}
