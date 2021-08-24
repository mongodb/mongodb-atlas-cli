package events

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func Test_orgListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationEventLister(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.EventResponse{}
	listOpts := &orgListOpts{
		store: mockStore,
	}
	listOpts.OrgID = "1"

	mockStore.
		EXPECT().OrganizationEvents(listOpts.OrgID, listOpts.newEventListOptions()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestOrgListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		OrgListBuilder(),
		0,
		[]string{
			flag.Limit,
			flag.Page,
			flag.Output,
			flag.OrgID,
			flag.Type,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}

func TestOrgsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		OrgsBuilder(),
		1,
		[]string{},
	)
}
