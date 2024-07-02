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

package datafederation

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DeletingState = "DELETING"
	DeletedState  = "DELETED"
)

func BuildAtlasDataFederation(dataFederationStore store.DataFederationStore, dataFederationName, projectID, projectName, operatorVersion, targetNamespace string, dictionary map[string]string) (*akov2.AtlasDataFederation, error) {
	dataFederation, err := dataFederationStore.DataFederation(projectID, dataFederationName)
	if err != nil {
		return nil, err
	}
	if !isDataFederationExportable(dataFederation) {
		return nil, nil
	}
	atlasDataFederation := &akov2.AtlasDataFederation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasDataFederation",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: getDataFederationSpec(dataFederation, targetNamespace, projectName),
		Status: akov2status.DataFederationStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}
	return atlasDataFederation, nil
}

func isDataFederationExportable(dataFederation *atlasv2.DataLakeTenant) bool {
	state := dataFederation.GetState()
	return state != DeletingState && state != DeletedState
}

func getDataFederationSpec(dataFederationSpec *atlasv2.DataLakeTenant, targetNamespace, projectName string) akov2.DataFederationSpec {
	return akov2.DataFederationSpec{
		Project:             akov2common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		Name:                dataFederationSpec.GetName(),
		CloudProviderConfig: getCloudProviderConfig(dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   getDataProcessRegion(dataFederationSpec.DataProcessRegion),
		Storage:             getStorage(dataFederationSpec.Storage),
	}
}

func getCloudProviderConfig(cloudProviderConfig *atlasv2.DataLakeCloudProviderConfig) *akov2.CloudProviderConfig {
	if cloudProviderConfig == nil {
		return &akov2.CloudProviderConfig{}
	}
	return &akov2.CloudProviderConfig{
		AWS: &akov2.AWSProviderConfig{
			RoleID:       cloudProviderConfig.Aws.RoleId,
			TestS3Bucket: cloudProviderConfig.Aws.TestS3Bucket,
		},
	}
}

func getDataProcessRegion(dataProcessRegion *atlasv2.DataLakeDataProcessRegion) *akov2.DataProcessRegion {
	if dataProcessRegion == nil {
		return &akov2.DataProcessRegion{}
	}
	return &akov2.DataProcessRegion{
		CloudProvider: dataProcessRegion.CloudProvider,
		Region:        dataProcessRegion.Region,
	}
}

func getStorage(storage *atlasv2.DataLakeStorage) *akov2.Storage {
	if storage == nil {
		return &akov2.Storage{}
	}
	return &akov2.Storage{
		Databases: getDatabases(storage.GetDatabases()),
		Stores:    getStores(storage.GetStores()),
	}
}

func getDatabases(database []atlasv2.DataLakeDatabaseInstance) []akov2.Database {
	if database == nil {
		return []akov2.Database{}
	}
	result := make([]akov2.Database, 0, len(database))

	for _, obj := range database {
		result = append(result, akov2.Database{
			Collections:            getCollection(obj.GetCollections()),
			MaxWildcardCollections: obj.GetMaxWildcardCollections(),
			Name:                   obj.GetName(),
			Views:                  getViews(obj.GetViews()),
		})
	}
	return result
}

func getCollection(collections []atlasv2.DataLakeDatabaseCollection) []akov2.Collection {
	if collections == nil {
		return []akov2.Collection{}
	}
	result := make([]akov2.Collection, 0, len(collections))

	for _, obj := range collections {
		result = append(result, akov2.Collection{
			DataSources: getDataSources(obj.GetDataSources()),
			Name:        obj.GetName(),
		})
	}
	return result
}

func getDataSources(dataSources []atlasv2.DataLakeDatabaseDataSourceSettings) []akov2.DataSource {
	if dataSources == nil {
		return []akov2.DataSource{}
	}
	result := make([]akov2.DataSource, 0, len(dataSources))

	for _, obj := range dataSources {
		result = append(result, akov2.DataSource{
			AllowInsecure:       obj.GetAllowInsecure(),
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

func getViews(views []atlasv2.DataLakeApiBase) []akov2.View {
	if views == nil {
		return []akov2.View{}
	}
	result := make([]akov2.View, 0, len(views))

	for _, obj := range views {
		result = append(result, akov2.View{
			Name:     obj.GetName(),
			Pipeline: obj.GetPipeline(),
			Source:   obj.GetSource(),
		})
	}
	return result
}

func getStores(stores []atlasv2.DataLakeStoreSettings) []akov2.Store {
	if stores == nil {
		return []akov2.Store{}
	}
	result := make([]akov2.Store, 0, len(stores))

	for _, obj := range stores {
		result = append(result, akov2.Store{
			Name:                     obj.GetName(),
			Provider:                 obj.Provider,
			AdditionalStorageClasses: obj.GetAdditionalStorageClasses(),
			Bucket:                   obj.GetBucket(),
			Delimiter:                obj.GetDelimiter(),
			IncludeTags:              obj.GetIncludeTags(),
			Prefix:                   obj.GetPrefix(),
			Public:                   obj.GetPublic(),
			Region:                   obj.GetRegion(),
		})
	}
	return result
}
