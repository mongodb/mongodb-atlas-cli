package events

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func Test_projectListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectEventLister(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.EventResponse{}
	listOpts := &projectListOpts{
		store: mockStore,
	}
	listOpts.ProjectID = "1"

	mockStore.
		EXPECT().ProjectEvents(listOpts.ProjectID, listOpts.newEventListOptions()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestProjectListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectListBuilder(),
		0,
		[]string{
			flag.Limit,
			flag.Page,
			flag.Output,
			flag.ProjectID,
			flag.Type,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}

func TestProjectsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectsBuilder(),
		1,
		[]string{},
	)
}
