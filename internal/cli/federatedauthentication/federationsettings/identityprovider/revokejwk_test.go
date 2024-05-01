// Copyright 2024 MongoDB Inc
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

package identityprovider

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func TestRevoke_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIdentityProviderJwkRevoker(ctrl)

	revokeOpts := &RevokeOpts{
		store:                mockStore,
		FederationSettingsID: "federationSettingsID",
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "id",
			Confirm: true,
		},
	}

	mockStore.
		EXPECT().
		RevokeJwksFromIdentityProvider(revokeOpts.FederationSettingsID, revokeOpts.Entry).
		Return(nil).
		Times(1)

	if err := revokeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestRevokeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RevokeBuilder(),
		0,
		[]string{flag.Output, flag.FederationSettingsID},
	)
}
