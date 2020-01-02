package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestIAMProjectsCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectCreator(ctrl)

	defer ctrl.Finish()

	expected := mocks.Project1

	mockStore.
		EXPECT().
		CreateProject(gomock.Eq("ProjectBar"), gomock.Eq("5a0a1e7e0f2912c554080adc")).Return(expected, nil).
		Times(1)

	createOpts := &IAMProjectsCreateOpts{
		store: mockStore,
		name:  "ProjectBar",
		orgID: "5a0a1e7e0f2912c554080adc",
	}
	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
