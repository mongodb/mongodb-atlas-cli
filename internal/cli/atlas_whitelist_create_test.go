package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasWhitelistCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPWhitelistCreator(ctrl)

	defer ctrl.Finish()

	expected := mocks.ProjectIPWhitelistMock()

	createOpts := &AtlasWhitelistCreateOpts{
		projectID: "5a0a1e7e0f2912c554080adc",
		entry:     "37.228.254.100",
		entryType: ipAddress,
		store:     mockStore,
	}

	mockStore.
		EXPECT().
		CreateProjectIPWhitelist(createOpts.newWhitelist()).Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
