package restores

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestWatchBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.ProjectID,
			flag.Output,
		},
	)
}

func TestWatchOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsDescriber(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.CloudProviderSnapshotRestoreJob{
		Failed: pointy.Bool(true),
	}

	describeOpts := &WatchOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		id:          "1",
	}

	mockStore.
		EXPECT().
		RestoreJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
