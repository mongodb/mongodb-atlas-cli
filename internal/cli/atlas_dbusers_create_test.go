package cli

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasDBUserCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseUserCreator(ctrl)

	defer ctrl.Finish()

	expected := mocks.DatabaseUserMock()

	createOpts := &AtlasDBUsersCreateOpts{
		projectID: "5a0a1e7e0f2912c554080adc",
		username:  "ProjectBar",
		password:  "US",
		roles:     []string{"admin@admin"},
		store:     mockStore,
	}

	mockStore.
		EXPECT().
		CreateDatabaseUser(createOpts.newDatabaseUser()).Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
