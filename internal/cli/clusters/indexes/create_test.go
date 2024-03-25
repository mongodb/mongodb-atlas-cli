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

//go:build unit

package indexes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIndexCreator(ctrl)

	createOpts := &CreateOpts{
		name:        "ProjectBar",
		clusterName: "US",
		db:          "test",
		collection:  "test",
		keys:        []string{"name:1"},
		store:       mockStore,
	}

	index, _ := createOpts.newIndex()
	mockStore.
		EXPECT().
		CreateIndex(createOpts.ProjectID, createOpts.clusterName, index).
		Return(nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateWithFile_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIndexCreator(ctrl)
	appFS := afero.NewMemMapFs()
	fileJSON := `
{
	"collection": "collectionName",
	"db": "dbName",
	"options":{
		"sparse": true,
		"unique": true,
		"textIndexVersion": 1,
		"name": "myIndex",
		"min": 1,
		"max": 10,
		"language_override": "test",
		"hidden": true,
		"expireAfterSeconds": 2,
		"default_language": "test",
		"default_language": "test",
		"columnstoreProjection": {"key":1, "key2":2},
		"bucketSize": 2,
		"bits": 222,
		"background": false,
		"2dsphereIndexVersion": 2
	}
}`
	fileName := "atlas_cluster_index_create_test.json"
	_ = afero.WriteFile(appFS, fileName, []byte(fileJSON), 0600)
	createOpts := &CreateOpts{
		filename: fileName,
		store:    mockStore,
		fs:       appFS,
	}

	index, _ := createOpts.newIndex()
	mockStore.
		EXPECT().
		CreateIndex(createOpts.ProjectID, createOpts.clusterName, index).
		Return(nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.ClusterName, flag.Database, flag.Collection, flag.Key, flag.Sparse, flag.ProjectID, flag.File},
	)
}
