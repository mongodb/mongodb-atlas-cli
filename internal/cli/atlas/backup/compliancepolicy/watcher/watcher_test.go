// Copyright 2023 MongoDB Inc
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

package watcher

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	fakeProjectID = "stringstringstringstring"
)

func TestCompliancePolicyWatcherFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyDescriber(ctrl)
	projectID := fakeProjectID

	actual := &atlasv2.DataProtectionSettings{
		CopyProtectionEnabled: new(bool),
	}
	expected := &atlasv2.DataProtectionSettings{}
	expected.SetCopyProtectionEnabled(true)
	expected.SetState(active)

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(projectID).
		Return(expected, nil).
		Times(1)

	watcher := CompliancePolicyWatcherFactory(projectID, mockStore, actual)
	isStateActive, err := watcher()
	if err != nil {
		t.Fatalf("watcher function unexpected error: %v", err)
	}
	assert.True(t, isStateActive)
	assert.Equal(t, expected, actual)
}

func TestCompliancePolicyWatcherFactory_Fail_State(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyDescriber(ctrl)
	projectID := fakeProjectID

	expected := &atlasv2.DataProtectionSettings{}
	expected.SetState("")

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(projectID).
		Return(expected, nil).
		Times(1)

	watcher := CompliancePolicyWatcherFactory(projectID, mockStore, atlasv2.NewDataProtectionSettings())
	_, err := watcher()
	assert.ErrorIs(t, err, errInvalidStateField)
}

func TestCompliancePolicyWatcherFactory_Fail_APICall(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyDescriber(ctrl)
	projectID := fakeProjectID

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(projectID).
		Return(nil, errors.New("network error")).
		Times(1)

	watcher := CompliancePolicyWatcherFactory(projectID, mockStore, atlasv2.NewDataProtectionSettings())
	_, err := watcher()
	assert.Error(t, err)
}
