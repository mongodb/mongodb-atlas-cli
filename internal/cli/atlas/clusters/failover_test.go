// Copyright 2023 MongoDB Inc
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

package clusters

import (
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/mocks"
	"github.com/andreangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
)

func TestFailoverBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		FailoverBuilder(),
		0,
		[]string{flag.Force, flag.ProjectID},
	)
}

func TestFailoverOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterTester(ctrl)

	deleteOpts := &FailoverOpts{
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
			Entry:   "test",
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		TestClusterFailover(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(nil).
		Times(1)

	if err := deleteOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
