package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestAtlasEventsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsStore(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Events()

	t.Run("for a project", func(t *testing.T) {
		listOpts := &atlasEventsListOpts{
			store: mockStore,
		}
		listOpts.orgID = "1"

		mockStore.
			EXPECT().OrganizationEvents(listOpts.orgID, listOpts.newEventListOptions()).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
	t.Run("for an org", func(t *testing.T) {
		listOpts := &atlasEventsListOpts{
			store: mockStore,
		}

		listOpts.projectID = "1"
		mockStore.
			EXPECT().ProjectEvents(listOpts.projectID, &listOpts.newEventListOptions().ListOptions).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
