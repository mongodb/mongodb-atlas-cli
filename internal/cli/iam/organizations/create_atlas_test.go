package organizations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCreateAtlasBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateAtlasBuilder(),
		0,
		[]string{
			flag.OwnerID,
			flag.APIKeyDescription,
			flag.APIKeyRole,
			flag.Output,
		},
	)
}

func TestCreateAtlasOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasOrganizationCreator(ctrl)

	expected := &mongodbatlas.CreateOrganizationRequest{
		APIKey:     nil,
		Name:       "Org 0",
		OrgOwnerID: nil,
	}
	resp := &mongodbatlas.CreateOrganizationResponse{}
	mockStore.
		EXPECT().
		CreateAtlasOrganization(expected).Return(resp, nil).
		Times(1)

	createOpts := &CreateAtlasOpts{
		store: mockStore,
		name:  "Org 0",
	}
	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
