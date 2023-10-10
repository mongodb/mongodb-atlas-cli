package fixture

import (
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
)

type MockDeploymentOpts struct {
	ctrl                      *gomock.Controller
	MockCredentialsGetter     *mocks.MockCredentialsGetter
	MockAtlasClusterListStore *mocks.MockClusterLister
	MockPodman                *mocks.MockClient
	Opts                      *options.DeploymentOpts
}
