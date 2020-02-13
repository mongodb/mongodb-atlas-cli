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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mcli/internal/mocks"
)

func TestIAMOrganizationsDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationDeleter(ctrl)

	defer ctrl.Finish()

	mockStore.
		EXPECT().
		DeleteOrganization(gomock.Eq("5a0a1e7e0f2912c554080adc")).
		Return(nil).
		Times(1)

	deleteOpts := &iamOrganizationsDeleteOpts{
		store: mockStore,
		deleteOpts: &deleteOpts{
			entry:   "5a0a1e7e0f2912c554080adc",
			confirm: true,
		},
	}
	err := deleteOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
