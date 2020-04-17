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

package cli

import (
	"testing"

	"github.com/go-test/deep"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestOpsManagerClustersIndexesCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationPatcher(ctrl)

	defer ctrl.Finish()

	expected := fixtures.AutomationConfig()

	createOpts := &opsManagerClustersIndexesCreateOpts{
		globalOpts:      globalOpts{},
		name:            "index",
		db:              "db",
		collection:      "db",
		rsName:          "rsName",
		locale:          "locale",
		caseFirst:       "test",
		alternate:       "alternate",
		maxVariable:     "max",
		strength:        0,
		caseLevel:       false,
		numericOrdering: false,
		normalization:   false,
		backwards:       false,
		keys:            []string{"field:key", "field:key", "field:key"},
		options:         []string{"field:key", "field:key", "field:key"},
		store:           mockStore,
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(createOpts.projectID).
		Return(expected, nil).
		Times(1)

	mockStore.
		EXPECT().
		UpdateAutomationConfig(createOpts.projectID, expected).
		Return(nil).
		Times(1)

	err := createOpts.Run()

	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestOpsManagerClustersIndexesCreate_newIndexMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationPatcher(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Indexes()

	createOpts := &opsManagerClustersIndexesCreateOpts{
		globalOpts:      globalOpts{},
		name:            "index",
		db:              "dbname",
		collection:      "collection",
		rsName:          "name",
		locale:          "loc",
		caseFirst:       "test",
		alternate:       "test",
		maxVariable:     "test",
		strength:        1,
		caseLevel:       true,
		background:      true,
		unique:          true,
		sparse:          true,
		numericOrdering: true,
		backwards:       true,
		keys:            []string{"field:key"},
		options:         []string{"field:key"},
		store:           mockStore,
	}

	indexMap, err := createOpts.newIndexMap()

	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if diff := deep.Equal(indexMap, expected); diff[0] != "slice[0].map[keys]: []interface {} != []map[string]interface {}" {
		t.Error(diff)
	}

}
