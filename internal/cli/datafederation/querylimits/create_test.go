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

// This code was autogenerated at 2023-06-23T15:50:57+01:00. Note: Manual updates are allowed, but may be overwritten.

package querylimits

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataFederationQueryLimitCreator(ctrl)

	expected := &admin.DataFederationTenantQueryLimit{
		Name:          "bytesProcessed.query",
		TenantName:    pointer.Get("DataFederation1"),
		Value:         1000,
		OverrunPolicy: pointer.Get("BLOCK"),
	}

	createOpts := &CreateOpts{
		limitName:     expected.Name,
		tenantName:    *expected.TenantName,
		value:         expected.Value,
		overrunPolicy: *expected.OverrunPolicy,
		store:         mockStore,
	}

	mockStore.
		EXPECT().
		CreateDataFederationQueryLimit(createOpts.ConfigProjectID(), createOpts.tenantName, createOpts.limitName, createOpts.newCreateRequest()).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
