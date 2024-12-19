package processor

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
)

func TestDeleteOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessorDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		store: mockStore,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
			Entry:   "test",
		},
		streamsInstance: "ExampleInstance",
	}

	mockStore.
		EXPECT().
		DeleteStreamProcessor(deleteOpts.ProjectID, deleteOpts.streamsInstance, deleteOpts.Entry).
		Return(nil).
		Times(1)

	if err := deleteOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
