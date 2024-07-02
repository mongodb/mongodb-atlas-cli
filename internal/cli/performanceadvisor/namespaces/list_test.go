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

//go:build unit

package namespaces

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestNamespacesList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPerformanceAdvisorNamespacesLister(ctrl)

	expected := atlasv2.Namespaces{
		Namespaces: &[]atlasv2.NamespaceObj{
			{
				Namespace: pointer.Get("test"),
			},
		},
	}

	listOpts := &ListOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		PerformanceAdvisorNamespaces(listOpts.newNamespaceOptions(listOpts.ProjectID, listOpts.ProcessName)).
		Return(&expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, listTemplate, expected)
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Duration,
			flag.Since,
			flag.ProcessName,
		},
	)
}

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		1,
		[]string{},
	)
}
