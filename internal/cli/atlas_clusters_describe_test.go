package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasClustersDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterDescriber(ctrl)

	defer ctrl.Finish()

	expected := mocks.ClusterMock()

	describeOpts := &atlasClustersDescribeOpts{
		globalOpts: newGlobalOpts(),
		name:       "test",
		store:      mockStore,
	}

	mockStore.
		EXPECT().
		Cluster(describeOpts.projectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	err := describeOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
