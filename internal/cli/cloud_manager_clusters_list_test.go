package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestCloudManagerClustersList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationGetter(ctrl)

	defer ctrl.Finish()

	expected := mocks.AutomationMock()

	listOpts := &cmClustersListOpts{
		globalOpts: newGlobalOpts(),
		store:      mockStore,
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(listOpts.projectID).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
