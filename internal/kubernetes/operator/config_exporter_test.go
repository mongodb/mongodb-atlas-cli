// Copyright 2022 MongoDB Inc
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
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const projectID = "TestProjectID"

func Test_fetchDataFederationNames(t *testing.T) {
	ctl := gomock.NewController(t)
	atlasOperatorGenericStore := mocks.NewMockAtlasOperatorGenericStore(ctl)

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
			nil, /* credsProvider (not used) */
			projectID, "" /* orgID (not used) */)
		got, err := ce.fetchDataFederationNames()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(expected, got) {
			expJs, _ := json.MarshalIndent(expected, "", " ")
			gotJs, _ := json.MarshalIndent(got, "", " ")
			fmt.Printf("E:%s\r\n; G:%s\r\n", expJs, gotJs)
			t.Fatalf("Data federation names mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}
