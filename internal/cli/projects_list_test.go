package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestProjectsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectLister(ctrl)

	defer ctrl.Finish()

	expected := mocks.ProjectsMock()

	mockStore.
		EXPECT().
		GetAllProjects().
		Return(expected, nil).
		Times(1)

	listOpts := &ListProjectOpts{store: mockStore}
	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
