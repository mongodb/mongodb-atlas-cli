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

package link

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
	"go.uber.org/mock/gomock"
)

func TestLinkTokenCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTokenCreator(ctrl)

	expected := &atlasv2.TargetOrg{}
	createOpts := &CreateOpts{
		OrgOpts:      cli.OrgOpts{OrgID: "1"},
		accessListIP: []string{"1.2.3.4", "5.6.7.8"},
		store:        mockStore,
	}

	createRequest := createOpts.newTokenCreateRequest()

	mockStore.
		EXPECT().CreateLinkToken(createOpts.OrgID, createRequest).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
