package fixture

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
)

func NewMockLocalDeploymentOpts(ctrl *gomock.Controller, deploymentName string) MockDeploymentOpts {
	mockPodman := mocks.NewMockClient(ctrl)
	return MockDeploymentOpts{
		ctrl:       ctrl,
		MockPodman: mockPodman,
		Opts: &options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: deploymentName,
			DeploymentType: "local",
		},
	}
}
func (m *MockDeploymentOpts) LocalMockFlow(ctx context.Context) {
	m.MockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	m.MockPodman.
		EXPECT().
		ListContainers(ctx, options.MongodHostnamePrefix).
		Return(m.MockContainer(), nil).
		Times(1)
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
