// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package nodes

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const testJSON = `{"Specs":[{"instanceSize": "S20_HIGHCPU_NVME", "nodeCount": 2}, {"instanceSize": "S110_LOWCPU_NVME", "nodeCount": 42}]}`
const testInvalidJSON = `(╯°□°)╯︵ ┻━┻`
const fileName = "spec.json"

var testJSONParsed = atlasv2.ApiSearchDeploymentRequest{
	Specs: []atlasv2.ApiSearchDeploymentSpec{
		{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2},
		{InstanceSize: "S110_LOWCPU_NVME", NodeCount: 42},
	},
}

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSearchNodesCreator(ctrl)

	t.Run("valid file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		_ = afero.WriteFile(appFS, fileName, []byte(testJSON), 0600)

		opts := &CreateOpts{
			store: mockStore,
		}
		opts.filename = fileName
		opts.fs = appFS

		expected := &atlasv2.ApiSearchDeploymentResponse{
			GroupId: pointer.Get("32b6e34b3d91647abb20e111"),
			Id:      pointer.Get("32b6e34b3d91647abb20e222"),
			Specs: &[]atlasv2.ApiSearchDeploymentSpec{
				{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2},
				{InstanceSize: "S110_LOWCPU_NVME", NodeCount: 42},
			},
			StateName: pointer.Get("UPDATING"),
		}

		mockStore.
			EXPECT().
			CreateSearchNodes(opts.ProjectID, opts.clusterName, &testJSONParsed).Return(expected, nil).
			Times(1)

		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("invalid file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		_ = afero.WriteFile(appFS, fileName, []byte(testInvalidJSON), 0600)

		opts := &CreateOpts{
			store: mockStore,
		}
		opts.filename = fileName
		opts.fs = appFS

		err := opts.Run()
		if err == nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		const expectedError = "failed to parse JSON file due to"
		if !strings.Contains(err.Error(), expectedError) {
			t.Fatalf("newSearchIndex() unexpected error: %v expected: %s", err, expectedError)
		}
	})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.File,
			flag.ProjectID,
			flag.Output,
		},
	)
}
