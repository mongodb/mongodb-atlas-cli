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
// +build unit

package create

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestOpsGenieOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIntegrationCreator(ctrl)
	defer ctrl.Finish()

	opts := &OpsGenieOpts{
		store: mockStore,
	}

	expected := &mongodbatlas.ThirdPartyIntegrations{}
	mockStore.
		EXPECT().
		CreateIntegration(opts.ProjectID, opsGenieType, opts.newOpsGenieIntegration()).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestOpsGenieBuilder(t *testing.T) {
	cli.CmdValidator(
		t,
		OpsGenieBuilder(),
		0,
		[]string{flag.APIKey},
	)
}
