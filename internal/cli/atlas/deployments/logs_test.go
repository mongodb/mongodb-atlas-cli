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

package deployments

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
)

func TestLogsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LogsBuilder(),
		0,
		[]string{},
	)
}

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	downloadOpts := &DownloadOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient: mockPodman,
		},
	}

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
