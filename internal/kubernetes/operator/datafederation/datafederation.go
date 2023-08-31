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

package datafederation

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20230201004/admin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DeletingState = "DELETING"
	DeletedState  = "DELETED"
)

func BuildAtlasDataFederation(dataFederationStore store.DataFederationStore, dataFederationName, projectID, projectName, operatorVersion, targetNamespace string, dictionary map[string]string) (*atlasV1.AtlasDataFederation, error) {
	dataFederation, err := dataFederationStore.DataFederation(projectID, dataFederationName)
	if err != nil {
		return nil, err
	}
	if !isDataFederationExportable(dataFederation) {
		return nil, nil
	}
	atlasDataFederation := &atlasV1.AtlasDataFederation{
		TypeMeta: v1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasDataFederation",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: GetDataFederationSpec(dataFederation, targetNamespace, projectName),
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
	return atlasDataFederation, nil
}

func isDataFederationExportable(dataFederation *admin.DataLakeTenant) bool {
	state := dataFederation.GetState()
	if state == DeletingState || state == DeletedState {
		return false
	}
	return true
}

func GetDataFederationSpec(dataFederationSpec *admin.DataLakeTenant, targetNamespace, projectName string) atlasV1.DataFederationSpec {
	return atlasV1.DataFederationSpec{
		Project:             common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		Name:                dataFederationSpec.GetName(),
		CloudProviderConfig: GetCloudProviderConfig(dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   GetDataProcessRegion(dataFederationSpec.DataProcessRegion),
		Storage:             GetStorage(dataFederationSpec.Storage),
	}
}

func GetCloudProviderConfig(cloudProviderConfig *admin.DataLakeCloudProviderConfig) *atlasV1.CloudProviderConfig {
	if cloudProviderConfig == nil {
		return &atlasV1.CloudProviderConfig{}
	}
	return &atlasV1.CloudProviderConfig{
		AWS: &atlasV1.AWSProviderConfig{
			RoleID:       cloudProviderConfig.Aws.RoleId,
			TestS3Bucket: cloudProviderConfig.Aws.TestS3Bucket,
		},
	}
}

func GetDataProcessRegion(dataProcessRegion *admin.DataLakeDataProcessRegion) *atlasV1.DataProcessRegion {
	if dataProcessRegion == nil {
		return &atlasV1.DataProcessRegion{}
	}
	return &atlasV1.DataProcessRegion{
		CloudProvider: dataProcessRegion.CloudProvider,
		Region:        dataProcessRegion.Region,
	}
}

func GetStorage(storage *admin.DataLakeStorage) *atlasV1.Storage {
	if storage == nil {
		return &atlasV1.Storage{}
	}
	return &atlasV1.Storage{
		Databases: GetDatabases(storage.Databases),
		Stores:    GetStores(storage.Stores),
	}
}
func GetDatabases(database []admin.DataLakeDatabaseInstance) []atlasV1.Database {
	if database == nil {
		return []atlasV1.Database{}
	}
	var result []atlasV1.Database
	for _, obj := range database {
		result = append(result, atlasV1.Database{
			Collections:            GetCollection(obj.GetCollections()),
			MaxWildcardCollections: obj.GetMaxWildcardCollections(),
			Name:                   obj.GetName(),
			Views:                  GetViews(obj.GetViews()),
		})
	}
	return result
}

func GetCollection(collections []admin.DataLakeDatabaseCollection) []atlasV1.Collection {
	if collections == nil {
		return []atlasV1.Collection{}
	}
	var result []atlasV1.Collection
	for _, obj := range collections {
		result = append(result, atlasV1.Collection{
			DataSources: GetDataSources(obj.GetDataSources()),
			Name:        obj.GetName(),
		})
	}
	return result
}

func GetDataSources(dataSources []admin.DataLakeDatabaseDataSourceSettings) []atlasV1.DataSource {
	if dataSources == nil {
		return []atlasV1.DataSource{}
	}
	var result []atlasV1.DataSource
	for _, obj := range dataSources {
		result = append(result, atlasV1.DataSource{
			AllowInsecure:       aws.ToBool(obj.AllowInsecure),
			Collection:          obj.GetCollection(),
			CollectionRegex:     obj.GetCollectionRegex(),
			Database:            obj.GetDatabase(),
			DatabaseRegex:       obj.GetDatabaseRegex(),
			DefaultFormat:       obj.GetDefaultFormat(),
			Path:                obj.GetPath(),
			ProvenanceFieldName: obj.GetProvenanceFieldName(),
			StoreName:           obj.GetStoreName(),
			Urls:                obj.GetUrls(),
		})
	}
	return result
}

func GetViews(views []admin.DataLakeApiBase) []atlasV1.View {
	if views == nil {
		return []atlasV1.View{}
	}
	var result []atlasV1.View
	for _, obj := range views {
		result = append(result, atlasV1.View{
			Name:     obj.GetName(),
			Pipeline: obj.GetPipeline(),
			Source:   obj.GetSource(),
		})
	}
	return result
}

func GetStores(stores []admin.DataLakeStoreSettings) []atlasV1.Store {
	if stores == nil {
		return []atlasV1.Store{}
	}
	var result []atlasV1.Store
	for _, obj := range stores {
		result = append(result, atlasV1.Store{
			Name:                     obj.GetName(),
			Provider:                 obj.Provider,
			AdditionalStorageClasses: obj.GetAdditionalStorageClasses(),
			Bucket:                   obj.GetBucket(),
			Delimiter:                obj.GetDelimiter(),
			IncludeTags:              aws.ToBool(obj.IncludeTags),
			Prefix:                   obj.GetPrefix(),
			Public:                   aws.ToBool(obj.Public),
			Region:                   obj.GetRegion(),
		})
	}
	return result
}
