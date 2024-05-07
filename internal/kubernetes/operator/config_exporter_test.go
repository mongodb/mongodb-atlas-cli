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

//go:build unit

package operator

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20231115013/admin"
)

const projectID = "TestProjectID"

func Test_fetchDataFederationNames(t *testing.T) {
	ctl := gomock.NewController(t)
	atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)

	t.Run("Can fetch Data Federation Instance names", func(t *testing.T) {
		dataFederations := []admin.DataLakeTenant{
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance0"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance1"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
			{
				DataProcessRegion: &admin.DataLakeDataProcessRegion{
					CloudProvider: "TestProvider",
					Region:        "TestRegion",
				},
				Name:  pointer.Get("DataFederationInstance2"),
				State: pointer.Get("TestState"),
				Storage: &admin.DataLakeStorage{
					Databases: nil,
					Stores:    nil,
				},
			},
		}

		atlasOperatorGenericStore.EXPECT().DataFederationList(projectID).Return(dataFederations, nil)
		expected := []string{"DataFederationInstance0", "DataFederationInstance1", "DataFederationInstance2"}
		ce := NewConfigExporter(
			atlasOperatorGenericStore,
			nil,           // credsProvider (not used)
			projectID, "", // orgID (not used)
		)
		got, err := ce.fetchDataFederationNames()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if diff := deep.Equal(got, expected); diff != nil {
			t.Error(diff)
		}
	})
}

func TestProjectWithWrongOrgID(t *testing.T) {
	ctl := gomock.NewController(t)
	atlasOperatorGenericStore := mocks.NewMockOperatorGenericStore(ctl)

	t.Run("should fail flagging when the org id does not match the owner of the project", func(t *testing.T) {
		project := &admin.Group{
			Id:    pointer.Get("project-id"),
			Name:  "test-project",
			OrgId: "right-org-id",
		}

		atlasOperatorGenericStore.EXPECT().Project(projectID).Return(project, nil)
		ce := NewConfigExporter(
			atlasOperatorGenericStore,
			nil, // credsProvider (not used)
			projectID, "wrong-org-id",
		)
		_, got := ce.Run()
		expected := errors.New("the project test-project (project-id) is not part of the " +
			"organization \"wrong-org-id\", please confirm the arguments provided " +
			"to the command or you are using the correct profile")
		if diff := deep.Equal(got, expected); diff != nil {
			t.Error(diff)
		}
	})
}
