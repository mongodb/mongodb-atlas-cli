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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
)

func NewMockLocalDeploymentOpts(ctrl *gomock.Controller, deploymentName string) MockDeploymentOpts {
	mockPodman := mocks.NewMockClient(ctrl)
	mockDeploymentTelemetry := mocks.NewMockDeploymentTelemetry(ctrl)
	mockOpts := MockDeploymentOpts{
		ctrl:                    ctrl,
		MockPodman:              mockPodman,
		MockDeploymentTelemetry: mockDeploymentTelemetry,
		Opts: &options.DeploymentOpts{
			PodmanClient:        mockPodman,
			DeploymentName:      deploymentName,
			DeploymentType:      "local",
			DeploymentTelemetry: mockDeploymentTelemetry,
		},
	}
	return mockOpts
}

func (m *MockDeploymentOpts) LocalMockFlowWithMockContainer(ctx context.Context, mockContainer []*podman.Container) {
	m.MockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	m.MockPodman.
		EXPECT().
		ListContainers(ctx, options.MongodHostnamePrefix).
		Return(mockContainer, nil).
		Times(1)
}

func (m *MockDeploymentOpts) LocalMockFlow(ctx context.Context) {
	m.LocalMockFlowWithMockContainer(ctx, m.MockContainer())
}

func (m *MockDeploymentOpts) MockContainer() []*podman.Container {
	return m.MockContainerWithState("running")
}

func (m *MockDeploymentOpts) MockContainerWithState(state string) []*podman.Container {
	return []*podman.Container{
		{
			Names:  []string{m.Opts.DeploymentName},
			State:  state,
			Labels: map[string]string{"version": "6.0.9"},
			ID:     m.Opts.DeploymentName,
		},
	}
}
