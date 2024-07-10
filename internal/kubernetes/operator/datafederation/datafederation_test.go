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

///go:build unit

package datafederation

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	projectName        = "testProject-1"
	dataFederationName = "testDataFederation-1"
	targetNamespace    = "test-namespace-1"
	resourceVersion    = "x.y.z"
	projectID          = "test-project-id"
)

func Test_BuildAtlasDataFederation(t *testing.T) {
	ctl := gomock.NewController(t)
	dataFederationStore := mocks.NewMockDataFederationStore(ctl)
	dictionary := resources.AtlasNameToKubernetesName()

	t.Run("Can import Data Federations", func(t *testing.T) {
		dataFederation := &atlasv2.DataLakeTenant{
			CloudProviderConfig: &atlasv2.DataLakeCloudProviderConfig{
				Aws: atlasv2.DataLakeAWSCloudProviderConfig{
					RoleId:       "TestRoleID",
					TestS3Bucket: "TestBucket",
				},
			},
			DataProcessRegion: &atlasv2.DataLakeDataProcessRegion{
				CloudProvider: "TestProvider",
				Region:        "TestRegion",
			},
			Hostnames: &[]string{"TestHostname"},
			Name:      pointer.Get(dataFederationName),
			State:     pointer.Get("TestState"),
			Storage: &atlasv2.DataLakeStorage{
				Databases: &[]atlasv2.DataLakeDatabaseInstance{
					{
						Collections: &[]atlasv2.DataLakeDatabaseCollection{
							{
								DataSources: &[]atlasv2.DataLakeDatabaseDataSourceSettings{
									{
										AllowInsecure:       pointer.Get(true),
										Collection:          pointer.Get("TestCollection"),
										CollectionRegex:     pointer.Get("TestCollectionRegex"),
										Database:            pointer.Get("TestDatabase"),
										DatabaseRegex:       pointer.Get("TestDatabaseRegex"),
										DefaultFormat:       pointer.Get("TestFormat"),
										Path:                pointer.Get("TestPath"),
										ProvenanceFieldName: pointer.Get("TestFieldName"),
										StoreName:           pointer.Get("TestStoreName"),
										Urls:                &[]string{"TestUrl"},
									},
								},
								Name: pointer.Get("TestName"),
							},
						},
						MaxWildcardCollections: pointer.Get(10),
						Name:                   pointer.Get("TestName"),
						Views: &[]atlasv2.DataLakeApiBase{
							{
								Name:     pointer.Get("TestName"),
								Pipeline: pointer.Get("TestPipeline"),
								Source:   pointer.Get("TestSource"),
							},
						},
					},
				},
				Stores: &[]atlasv2.DataLakeStoreSettings{
					{
						Name:                     pointer.Get("TestName"),
						Provider:                 "TestProvider",
						AdditionalStorageClasses: &[]string{"TestClasses"},
						Bucket:                   pointer.Get("TestBucket"),
						Delimiter:                pointer.Get("TestDelimiter"),
						IncludeTags:              pointer.Get(true),
						Prefix:                   pointer.Get("TestPrefix"),
						Public:                   pointer.Get(true),
						Region:                   pointer.Get("TestRegion"),
					},
				},
			},
		}

		dataFederationStore.EXPECT().DataFederation(projectID, dataFederationName).Return(dataFederation, nil)

		expected := &akov2.AtlasDataFederation{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasDataFederation",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.GetName()), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: akov2.DataFederationSpec{
				Project: akov2common.ResourceRefNamespaced{
					Name:      projectName,
					Namespace: targetNamespace,
				},
				Name: dataFederationName,
				CloudProviderConfig: &akov2.CloudProviderConfig{
					AWS: &akov2.AWSProviderConfig{
						RoleID:       dataFederation.CloudProviderConfig.Aws.RoleId,
						TestS3Bucket: dataFederation.CloudProviderConfig.Aws.TestS3Bucket,
					},
				},
				DataProcessRegion: &akov2.DataProcessRegion{
					CloudProvider: dataFederation.DataProcessRegion.CloudProvider,
					Region:        dataFederation.DataProcessRegion.Region,
				},
				Storage: &akov2.Storage{
					Databases: []akov2.Database{
						{
							Collections: []akov2.Collection{
								{
									DataSources: []akov2.DataSource{
										{
											AllowInsecure:       true,
											Collection:          *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].Collection,
											CollectionRegex:     *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].CollectionRegex,
											Database:            *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].Database,
											DatabaseRegex:       *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].DatabaseRegex,
											DefaultFormat:       *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].DefaultFormat,
											Path:                *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].Path,
											ProvenanceFieldName: *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].ProvenanceFieldName,
											StoreName:           *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].StoreName,
											Urls:                []string{dataFederation.Storage.GetDatabases()[0].GetCollections()[0].GetDataSources()[0].GetUrls()[0]},
										},
									},
									Name: *dataFederation.Storage.GetDatabases()[0].GetCollections()[0].Name,
								},
							},
							MaxWildcardCollections: *dataFederation.Storage.GetDatabases()[0].MaxWildcardCollections,
							Name:                   *dataFederation.Storage.GetDatabases()[0].Name,
							Views: []akov2.View{
								{
									Name:     *dataFederation.Storage.GetDatabases()[0].GetViews()[0].Name,
									Pipeline: *dataFederation.Storage.GetDatabases()[0].GetViews()[0].Pipeline,
									Source:   *dataFederation.Storage.GetDatabases()[0].GetViews()[0].Source,
								},
							},
						},
					},
					Stores: []akov2.Store{
						{
							Name:                     *dataFederation.Storage.GetStores()[0].Name,
							Provider:                 dataFederation.Storage.GetStores()[0].Provider,
							AdditionalStorageClasses: []string{dataFederation.Storage.GetStores()[0].GetAdditionalStorageClasses()[0]},
							Bucket:                   *dataFederation.Storage.GetStores()[0].Bucket,
							Delimiter:                *dataFederation.Storage.GetStores()[0].Delimiter,
							IncludeTags:              *dataFederation.Storage.GetStores()[0].IncludeTags,
							Prefix:                   *dataFederation.Storage.GetStores()[0].Prefix,
							Public:                   *dataFederation.Storage.GetStores()[0].Public,
							Region:                   *dataFederation.Storage.GetStores()[0].Region,
						},
					},
				},
			},
			Status: akov2status.DataFederationStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		got, err := BuildAtlasDataFederation(dataFederationStore, dataFederationName, projectID, projectName, resourceVersion, targetNamespace, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if diff := deep.Equal(got, expected); diff != nil {
			t.Error(diff)
		}
	})
}
