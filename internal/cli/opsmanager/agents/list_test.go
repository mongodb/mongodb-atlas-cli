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

package agents

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestAgentsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAgentLister(ctrl)

	expected := &opsmngr.Agents{}

	listOpts := ListOpts{
		store:     mockStore,
		agentType: "MONITORING",
	}
	listOpts.ProjectID = "1"

	mockStore.
		EXPECT().
		Agents(listOpts.ProjectID, listOpts.agentType, listOpts.NewListOptions()).
		Return(expected, nil).
		Times(1)

	config.SetService(config.OpsManagerService)

	require.NoError(t, listOpts.Run())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Page, flag.Limit},
	)
}
