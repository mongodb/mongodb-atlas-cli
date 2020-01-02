package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasClustersList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterLister(ctrl)

	defer ctrl.Finish()

	expected := mocks.ClustersMock()

	listOpts := &AtlasClustersListOpts{
		projectID: "5a0a1e7e0f2912c554080adc",
		store:     mockStore,
	}

	mockStore.
		EXPECT().
		ProjectClusters(listOpts.projectID, listOpts.newListOptions()).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
