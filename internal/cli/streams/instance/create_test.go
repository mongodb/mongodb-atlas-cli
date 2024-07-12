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

package instance

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	testProjectID = "create-project-id"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsCreator(ctrl)

	t.Run("stream instances create", func(t *testing.T) {
		buf := new(bytes.Buffer)
		opts := &CreateOpts{
			store:    mockStore,
			name:     "ExampleStream",
			provider: "AWS",
			region:   "VIRGINIA_USA",
			tier:     "SP30", // Test non default case
		}
		opts.ProjectID = testProjectID

		expected := &atlasv2.StreamsTenant{
			Name:              &opts.name,
			GroupId:           &opts.ProjectID,
			DataProcessRegion: &atlasv2.StreamsDataProcessRegion{CloudProvider: "AWS", Region: "VIRGINIA_USA"},
			StreamConfig: &atlasv2.StreamConfig{
				Tier: &opts.tier,
			},
		}

		mockStore.
			EXPECT().
			CreateStream(opts.ProjectID, expected).
			Return(expected, nil).
			Times(1)

		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())
		test.VerifyOutputTemplate(t, createTemplate, expected)
	})

	t.Run("stream instances create --tier", func(t *testing.T) {
		buf := new(bytes.Buffer)
		opts := &CreateOpts{
			store:    mockStore,
			name:     "ExampleStream",
			provider: "AWS",
			region:   "VIRGINIA_USA",
			tier:     "SP10", // Test non default case
		}
		opts.ProjectID = testProjectID

		expected := &atlasv2.StreamsTenant{
			Name:              &opts.name,
			GroupId:           &opts.ProjectID,
			DataProcessRegion: &atlasv2.StreamsDataProcessRegion{CloudProvider: "AWS", Region: "VIRGINIA_USA"},
			StreamConfig: &atlasv2.StreamConfig{
				Tier: &opts.tier,
			},
		}

		mockStore.
			EXPECT().
			CreateStream(opts.ProjectID, expected).
			Return(expected, nil).
			Times(1)

		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())
		test.VerifyOutputTemplate(t, createTemplate, expected)
	})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.Provider, flag.Region, flag.Tier},
	)
}
