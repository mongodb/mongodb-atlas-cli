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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationAPIKeyAccessListWhitelistDeleter(ctrl)
	defer ctrl.Finish()

	prevServ := config.Service()
	config.SetService(config.OpsManagerService)
	defer func() {
		config.SetService(prevServ)
	}()

	deleteOpts := &DeleteOpts{
		store: mockStore,
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "5a0a1e7e0f2912c554080adc",
			Confirm: true,
		},
		apiKey: "1",
		GlobalOpts: cli.GlobalOpts{
			OrgID: "2",
		},
	}

	t.Run("OM 5.0", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "5.0.0.100.20210101T0000Z"}, nil).
			Times(1)

		mockStore.
			EXPECT().
			DeleteOrganizationAPIKeyAccessList(deleteOpts.OrgID, deleteOpts.apiKey, gomock.Eq("5a0a1e7e0f2912c554080adc")).
			Return(nil).
			Times(1)

		if err := deleteOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("OM 5.0-rc0", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "5.0.0-rc0.100.20210101T0000Z"}, nil).
			Times(1)

		mockStore.
			EXPECT().
			DeleteOrganizationAPIKeyAccessList(deleteOpts.OrgID, deleteOpts.apiKey, gomock.Eq("5a0a1e7e0f2912c554080adc")).
			Return(nil).
			Times(1)

		if err := deleteOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("OM 5.0", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "5.0.0.100.20210101T0000Z"}, nil).
			Times(1)

		mockStore.
			EXPECT().
			DeleteOrganizationAPIKeyAccessList(deleteOpts.OrgID, deleteOpts.apiKey, gomock.Eq("5a0a1e7e0f2912c554080adc")).
			Return(nil).
			Times(1)

		if err := deleteOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("OM 4.4", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetServiceVersion().
			Return(&atlas.ServiceVersion{GitHash: "some git hash", Version: "4.4.0.100.20210101T0000Z"}, nil).
			Times(1)

		mockStore.
			EXPECT().
			DeleteOrganizationAPIKeyWhitelist(deleteOpts.OrgID, deleteOpts.apiKey, gomock.Eq("5a0a1e7e0f2912c554080adc")).
			Return(nil).
			Times(1)

		if err := deleteOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DeleteBuilder(),
		0,
		[]string{flag.OrgID, flag.Force, flag.APIKey},
	)
}
