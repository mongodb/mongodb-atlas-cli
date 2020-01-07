package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestCloudManagerClustersDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationGetter(ctrl)

	defer ctrl.Finish()

	expected := mocks.AutomationMock()

	descOpts := &cmClustersDescribeOpts{
		globalOpts: newGlobalOpts(),
		store:      mockStore,
		name:       "myReplicaSet",
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(descOpts.projectID).
		Return(expected, nil).
		Times(1)

	err := descOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
