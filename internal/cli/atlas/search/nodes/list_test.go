package nodes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSearchNodesLister(ctrl)

	listOpts := &ListOpts{
		store: mockStore,
	}

	expected := &atlasv2.ApiSearchDeploymentResponse{}

	mockStore.
		EXPECT().
		SearchNodes(listOpts.ProjectID, listOpts.clusterName).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, listTemplate, expected)
}
