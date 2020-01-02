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

	listOpts := &AtlasClustersDescribeOpts{
		projectID: "5a0a1e7e0f2912c554080adc",
		name:      "test",
		store:     mockStore,
	}

	mockStore.
		EXPECT().
		Cluster(listOpts.projectID, listOpts.name).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
