// Copyright 2021 MongoDB Inc
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

package livemigrations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/livemigrations/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113002/admin"
)

func TestLiveMigrationCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockLiveMigrationCreator(ctrl)

	expected := atlasv2.LiveMigrationResponse{}

	createOpts := &CreateOpts{
		LiveMigrationsOpts: options.LiveMigrationsOpts{
			ProjectOpts:                 cli.ProjectOpts{ProjectID: "1"},
			SourceProjectID:             "2",
			SourceClusterName:           "testSrc",
			SourceManagedAuthentication: true,
			DestinationClusterName:      "testDest",
			MigrationHosts:              []string{"mig1"},
		},
		store: mockStore,
	}

	createRequest := createOpts.NewCreateRequest()

	mockStore.
		EXPECT().LiveMigrationCreate(createOpts.ProjectID, createRequest).Return(&expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
