// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fixture

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
)

func NewMockLocalDeploymentOpts(ctrl *gomock.Controller, deploymentName string) MockDeploymentOpts {
	mockContainerEngine := mocks.NewMockEngine(ctrl)
	mockDeploymentTelemetry := mocks.NewMockDeploymentTelemetry(ctrl)
	mockOpts := MockDeploymentOpts{
		ctrl:                    ctrl,
		MockContainerEngine:     mockContainerEngine,
		MockDeploymentTelemetry: mockDeploymentTelemetry,
		Opts: &options.DeploymentOpts{
			ContainerEngine:     mockContainerEngine,
			DeploymentName:      deploymentName,
			DeploymentType:      "local",
			DeploymentTelemetry: mockDeploymentTelemetry,
		},
	}
	return mockOpts
}

func (m *MockDeploymentOpts) LocalMockFlowWithMockContainer(ctx context.Context, mockContainer []container.Container) {
	m.MockContainerEngine.
		EXPECT().
		Ready().
		Return(nil).
		Times(1)
	m.MockContainerEngine.
		EXPECT().
		ContainerList(ctx, options.ContainerFilter).
		Return(mockContainer, nil).
		Times(1)

	m.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)
}

func (m *MockDeploymentOpts) LocalMockFlow(ctx context.Context) {
	m.LocalMockFlowWithMockContainer(ctx, m.MockContainerWithState("running"))
}

func (m *MockDeploymentOpts) MockContainerWithState(state string) []container.Container {
	return []container.Container{
		{
			Names:  []string{m.Opts.DeploymentName},
			State:  state,
			Labels: map[string]string{"version": "7.0.9"},
			ID:     m.Opts.DeploymentName,
		},
	}
}
