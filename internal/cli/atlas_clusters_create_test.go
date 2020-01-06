package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasClustersCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterCreator(ctrl)

	defer ctrl.Finish()

	expected := mocks.ClusterMock()

	createOpts := &atlasClustersCreateOpts{
		globalOpts:   newGlobalOpts(),
		name:         "ProjectBar",
		region:       "US",
		instanceSize: atlasM2,
		nodes:        3,
		diskSize:     10,
		backup:       false,
		mdbVersion:   currentMDBVersion,
		store:        mockStore,
	}

	mockStore.
		EXPECT().
		CreateCluster(createOpts.newCluster()).Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
