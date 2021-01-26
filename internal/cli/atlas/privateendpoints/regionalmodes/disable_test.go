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

package regionalmodes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestDisableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRegionalizedPrivateEndpointSettingUpdater(ctrl)
	defer ctrl.Finish()

	expected := &atlas.RegionalizedPrivateEndpointSetting{
		Enabled: true,
	}

	opts := &DisableOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		UpdateRegionalizedPrivateEndpointSetting(opts.ProjectID, false).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDisableBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DisableBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
