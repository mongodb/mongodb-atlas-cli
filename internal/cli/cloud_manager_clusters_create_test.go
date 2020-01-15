// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"testing"

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
		Return(nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
