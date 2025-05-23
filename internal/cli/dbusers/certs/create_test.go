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

package certs

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateBuilder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDBUserCertificateCreator(ctrl)

	username := "to_create"
	monthsUntilExpiry := 12

	createOpts := &CreateOpts{
		store:             mockStore,
		username:          username,
		monthsUntilExpiry: monthsUntilExpiry,
	}

	mockStore.
		EXPECT().
		CreateDBUserCertificate(createOpts.ProjectID, username, monthsUntilExpiry).
		Return("", nil).
		Times(1)

	require.NoError(t, createOpts.Run())
}

func TestTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, createTemplate, "")
}
