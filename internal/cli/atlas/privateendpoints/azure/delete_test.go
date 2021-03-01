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

// +build unit

package azure

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPrivateEndpointDeleter(ctrl)
	defer ctrl.Finish()

	deleteOpts := &DeleteOpts{
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "to_delete",
			Confirm: true,
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DeletePrivateEndpoint(deleteOpts.ProjectID, provider, deleteOpts.Entry).
		Return(nil).
		Times(1)

	err := deleteOpts.Run()
	assert.NoError(t, err)
}

func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DeleteBuilder(),
		0,
		[]string{flag.ProjectID, flag.Force},
	)
}
