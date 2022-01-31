// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit
// +build unit

package accesslists

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestWhitelistCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	var expected *mongodbatlas.ProjectIPAccessLists

	createOpts := &CreateOpts{
		entry:     "37.228.254.100",
		entryType: ipAddress,
		store:     mockStore,
	}

	mockStore.
		EXPECT().
		CreateProjectIPAccessList(createOpts.newProjectIPAccessList()).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Type, flag.Comment, flag.DeleteAfter},
	)
}

func TestValidateCurrentIPFlagNoFlagNoIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: ipAddress,
		store:     mockStore,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{}); errF() == nil {
		t.Fatalf("Error expected for empty args and no current ip flag.")
	}
}

func TestValidateCurrentIPFlagWithFlagWithIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: ipAddress,
		store:     mockStore,
		currentIP: true,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{"37.228.254.100"}); errF() == nil {
		t.Fatalf("Error expected for args and current ip flag in the same command.")
	}
}

func TestValidateCurrentIPFlagWithFlagNoIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: ipAddress,
		store:     mockStore,
		currentIP: true,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{}); errF() != nil {
		t.Fatalf("Error not expected for no args and current ip flag in the same command.")
	}
}

func TestValidateCurrentIPFlagNoFlagWithIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: ipAddress,
		store:     mockStore,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{"37.228.254.100"}); errF() != nil {
		t.Fatalf("Error not expected for args and no current ip flag in the same command.")
	}
}

func TestValidateCurrentIPFlagWithFlagNoIPWithCIDR(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: cidrBlock,
		store:     mockStore,
		currentIP: true,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{}); errF() == nil {
		t.Fatalf("Error expected for CIDR with no args current ip flag in the same command.")
	}
}

func TestValidateCurrentIPFlagWithFlagNoIPWithAWSSecurityGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: awsSecurityGroup,
		store:     mockStore,
		currentIP: true,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{}); errF() == nil {
		t.Fatalf("Error expected for awsSecurityGroup with no args current ip flag in the same command.")
	}
}

func TestValidateCurrentIPFlagWithFlagWithIPWithAWSSecurityGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectIPAccessListCreator(ctrl)
	defer ctrl.Finish()

	createOpts := &CreateOpts{
		entryType: awsSecurityGroup,
		store:     mockStore,
		currentIP: true,
	}

	if errF := createOpts.validateCurrentIPFlag(CreateBuilder(), []string{"37.228.254.100"}); errF() == nil {
		t.Fatalf("Error expected for awsSecurityGroup with no args current ip flag in the same command.")
	}
}
