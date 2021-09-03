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
// +build unit

package accesslists

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationAPIKeyAccessListWhitelistCreator(ctrl)
	defer ctrl.Finish()

	prevServ := config.Service()
	config.SetService(config.OpsManagerService)
	defer func() {
		config.SetService(prevServ)
	}()

	createOpts := &CreateOpts{
		store:  mockStore,
		apyKey: "1",
		ips:    []string{"77.54.32.11"},
	}

	r, err := createOpts.newAccessListAPIKeysReq()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	t.Run("OM 5.0", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "5.0.0.100"}, nil).
			Times(1)

		expected := &atlas.AccessListAPIKeys{}

		mockStore.
			EXPECT().
			CreateOrganizationAPIKeyAccessList(createOpts.OrgID, createOpts.apyKey, r).
			Return(expected, nil).
			Times(1)

		if err = createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("OM 4.4", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "4.4.0.100"}, nil).
			Times(1)

		req := fromAccessListAPIKeysReqToWhitelistAPIKeysReq(r)
		expected := &atlas.WhitelistAPIKeys{}

		mockStore.
			EXPECT().
			CreateOrganizationAPIKeyWhitelist(createOpts.OrgID, createOpts.apyKey, req).
			Return(expected, nil).
			Times(1)

		if err = createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.OrgID, flag.Output, flag.APIKey, flag.CIDR, flag.IP},
	)
}
