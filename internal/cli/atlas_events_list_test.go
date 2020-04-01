package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestAtlasOrganizationEventsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsStore(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Events()

	listOpts := &atlasEventsListOpts{
		globalOpts: newGlobalOpts(),
		source:     "organization",
		store:      mockStore,
	}

	mockStore.
		EXPECT().OrganizationEvents(listOpts.projectID, listOpts.newEventListOptions()).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestAtlasProjectEventsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsStore(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Events()

	listOpts := &atlasEventsListOpts{
		globalOpts: newGlobalOpts(),
		source:     "project",
		store:      mockStore,
	}

	mockStore.
		EXPECT().ProjectEvents(listOpts.projectID, &listOpts.newEventListOptions().ListOptions).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
