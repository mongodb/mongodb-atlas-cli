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

package accesslists

import (
	"testing"

	"go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockOrganizationAPIKeyAccessListCreator(ctrl)

	createOpts := &CreateOpts{
		store:  mockStore,
		apyKey: "1",
		ips:    []string{"77.54.32.11"},
	}

	r, err := createOpts.newAccessListAPIKeysReq()

	params := &admin.CreateApiKeyAccessListApiParams{
		OrgId:                 createOpts.OrgID,
		ApiUserId:             createOpts.apyKey,
		UserAccessListRequest: r,
	}
	if err != nil {
		t.Fatalf("newAccessListAPIKeysReq() unexpected error: %v", err)
	}

	mockStore.
		EXPECT().
		CreateOrganizationAPIKeyAccessList(params).
		Return(&admin.PaginatedApiUserAccessListResponse{}, nil).
		Times(1)

	if err = createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
