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

package validation

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/livemigrations/options"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
	"go.uber.org/mock/gomock"
)

func TestLiveMigrationValidationCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLiveMigrationValidationsCreator(ctrl)

	expected := atlasv2.LiveImportValidation{}

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
		EXPECT().CreateValidation(createOpts.ProjectID, createRequest).Return(&expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
