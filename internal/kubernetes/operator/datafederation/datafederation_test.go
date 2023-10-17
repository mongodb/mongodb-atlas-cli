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
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		dataFederation := &admin.DataLakeTenant{
			CloudProviderConfig: &admin.DataLakeCloudProviderConfig{
				Aws: admin.DataLakeAWSCloudProviderConfig{
					RoleId:       "TestRoleID",
					TestS3Bucket: "TestBucket",
				},
			},
			DataProcessRegion: &admin.DataLakeDataProcessRegion{
				CloudProvider: "TestProvider",
				Region:        "TestRegion",
			},
			Hostnames: []string{"TestHostname"},
			Name:      pointer.Get(dataFederationName),
			State:     pointer.Get("TestState"),
			Storage: &admin.DataLakeStorage{
				Databases: []admin.DataLakeDatabaseInstance{
					{
						Collections: []admin.DataLakeDatabaseCollection{
							{
								DataSources: []admin.DataLakeDatabaseDataSourceSettings{
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
										Urls:                []string{"TestUrl"},
									},
								},
								Name: pointer.Get("TestName"),
							},
						},
						MaxWildcardCollections: pointer.Get(10),
						Name:                   pointer.Get("TestName"),
						Views: []admin.DataLakeApiBase{
							{
								Name:     pointer.Get("TestName"),
								Pipeline: pointer.Get("TestPipeline"),
								Source:   pointer.Get("TestSource"),
							},
						},
					},
				},
				Stores: []admin.DataLakeStoreSettings{
					{
						Name:                     pointer.Get("TestName"),
						Provider:                 "TestProvider",
						AdditionalStorageClasses: []string{"TestClasses"},
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

		expected := &atlasV1.AtlasDataFederation{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDataFederation",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.GetName()), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: atlasV1.DataFederationSpec{
				Project: common.ResourceRefNamespaced{
					Name:      projectName,
					Namespace: targetNamespace,
				},
				Name: dataFederationName,
				CloudProviderConfig: &atlasV1.CloudProviderConfig{
					AWS: &atlasV1.AWSProviderConfig{
						RoleID:       dataFederation.CloudProviderConfig.Aws.RoleId,
						TestS3Bucket: dataFederation.CloudProviderConfig.Aws.TestS3Bucket,
					},
				},
				DataProcessRegion: &atlasV1.DataProcessRegion{
					CloudProvider: dataFederation.DataProcessRegion.CloudProvider,
					Region:        dataFederation.DataProcessRegion.Region,
				},
				Storage: &atlasV1.Storage{
					Databases: []atlasV1.Database{
						{
							Collections: []atlasV1.Collection{
								{
									DataSources: []atlasV1.DataSource{
										{
											AllowInsecure:       true,
											Collection:          *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Collection,
											CollectionRegex:     *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].CollectionRegex,
											Database:            *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Database,
											DatabaseRegex:       *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].DatabaseRegex,
											DefaultFormat:       *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].DefaultFormat,
											Path:                *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Path,
											ProvenanceFieldName: *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].ProvenanceFieldName,
											StoreName:           *dataFederation.Storage.Databases[0].Collections[0].DataSources[0].StoreName,
											Urls:                []string{dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Urls[0]},
										},
									},
									Name: *dataFederation.Storage.Databases[0].Collections[0].Name,
								},
							},
							MaxWildcardCollections: *dataFederation.Storage.Databases[0].MaxWildcardCollections,
							Name:                   *dataFederation.Storage.Databases[0].Name,
							Views: []atlasV1.View{
								{
									Name:     *dataFederation.Storage.Databases[0].Views[0].Name,
									Pipeline: *dataFederation.Storage.Databases[0].Views[0].Pipeline,
									Source:   *dataFederation.Storage.Databases[0].Views[0].Source,
								},
							},
						},
					},
					Stores: []atlasV1.Store{
						{
							Name:                     *dataFederation.Storage.Stores[0].Name,
							Provider:                 dataFederation.Storage.Stores[0].Provider,
							AdditionalStorageClasses: []string{dataFederation.Storage.Stores[0].AdditionalStorageClasses[0]},
							Bucket:                   *dataFederation.Storage.Stores[0].Bucket,
							Delimiter:                *dataFederation.Storage.Stores[0].Delimiter,
							IncludeTags:              *dataFederation.Storage.Stores[0].IncludeTags,
							Prefix:                   *dataFederation.Storage.Stores[0].Prefix,
							Public:                   *dataFederation.Storage.Stores[0].Public,
							Region:                   *dataFederation.Storage.Stores[0].Region,
						},
					},
				},
			},
			Status: status.DataFederationStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
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
