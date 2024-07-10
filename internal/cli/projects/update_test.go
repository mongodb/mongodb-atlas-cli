// Copyright 2024 MongoDB Inc
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

//go:build unit

package projects

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectUpdater(ctrl)

	const projectJSON = `{ "name": "testProject", "tags": [ { "key": "env", "value": "dev" }, { "key": "app", "value": "cli" } ] }`
	const filename = "myProject.json"

	appFS := afero.NewMemMapFs()
	_ = afero.WriteFile(appFS, filename, []byte(projectJSON), 0600)

	expected := &atlasv2.Group{
		Name: "testProject",
		Tags: &[]atlasv2.ResourceTag{
			{Key: "env", Value: "dev"},
			{Key: "app", Value: "cli"},
		},
	}

	updateOpts := &UpdateOpts{
		fs:        appFS,
		store:     mockStore,
		filename:  filename,
		projectID: "5a0a1e7e0f2912c554080add",
	}
	params, err := updateOpts.newUpdateProjectParams()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	mockStore.
		EXPECT().
		UpdateProject(params).Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, updateTemplate, atlasv2.Group{})
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.File},
	)
}
