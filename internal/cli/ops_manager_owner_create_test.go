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

	"go.mongodb.org/ops-manager/opsmngr"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestManagerOwnerCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOwnerCreator(ctrl)

	defer ctrl.Finish()

	email := "test@test.com"
	password := "Passw0rd!"
	firstName := "Testy"
	lastName := "McTestyson"
	expected := &opsmngr.CreateUserResponse{}

	t.Run("no whitelist", func(t *testing.T) {
		createOpts := &opsManagerOwnerCreateOpts{
			email:     email,
			password:  password,
			firstName: firstName,
			lastName:  lastName,
			store:     mockStore,
		}

		mockStore.
			EXPECT().
			CreateOwner(createOpts.newOwner(), createOpts.whitelistIps).Return(expected, nil).
			Times(1)

		err := createOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("with whitelist", func(t *testing.T) {
		createOpts := &opsManagerOwnerCreateOpts{
			email:        email,
			password:     password,
			firstName:    firstName,
			lastName:     lastName,
			store:        mockStore,
			whitelistIps: []string{"192.168.0.1"},
		}

		mockStore.
			EXPECT().
			CreateOwner(createOpts.newOwner(), createOpts.whitelistIps).Return(expected, nil).
			Times(1)

		err := createOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
