package cli

import (
	"testing"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/spf13/afero"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestCloudManagerClustersCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationStore(ctrl)

	defer ctrl.Finish()

	expected := mocks.AutomationMock()
	appFS := afero.NewMemMapFs()
	// create test file
	fileYML := `
---
name: "cluster_2"
version: 4.2.2
feature_compatibility_version: 4.2
processes:
  - hostname: host0
    db_path: /data/cluster_2/rs1
    log_path: /data/cluster_2/rs1/mongodb.log
    priority: 1
    votes: 1
    port: 29010
  - hostname: host1
    db_path: /data/cluster_2/rs2
    log_path: /data/cluster_2/rs2/mongodb.log
    priority: 1
    votes: 1
    port: 29020
  - hostname: host2
    db_path: /data/cluster_2/rs3
    log_path: /data/cluster_2/rs3/mongodb.log
    priority: 1
    votes: 1
    port: 29030`
	fileName := "test.yml"
	_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

	createOpts := &cmClustersCreateOpts{
		globalOpts: newGlobalOpts(),
		store:      mockStore,
		fs:         appFS,
		filename:   fileName,
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(createOpts.projectID).
		Return(expected, nil).
		Times(1)

	mockStore.
		EXPECT().
		UpdateAutomationConfig(createOpts.projectID, expected).
		Return(new(cloudmanager.AutomationConfig), nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
