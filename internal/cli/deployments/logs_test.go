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

package deployments

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/spf13/afero"
	"go.uber.org/mock/gomock"
)

func TestLogs_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	buf := new(bytes.Buffer)
	want := "localDeployment"
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, want)

	downloadOpts := &DownloadOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.LocalMockFlow(ctx)

	deploymentTest.MockContainerEngine.
		EXPECT().
		ContainerLogs(ctx, want).
		Return([]string{"log1", "log2"}, nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	expectedLogs := "log1\nlog2\n"

	if !strings.Contains(buf.String(), expectedLogs) {
		t.Fatalf("Run() expected output: %s, got: %s", expectedLogs, buf.String())
	}
}

func TestLogs_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	atlasDeployment := "localDeployment1"
	mockStore := NewMockLogsDownloader(ctrl)
	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, atlasDeployment)

	downloadOpts := &DownloadOpts{
		ProjectOpts: cli.ProjectOpts{ProjectID: "ProjectID"},
		DownloaderOpts: cli.DownloaderOpts{
			Out: "out",
		},
		DeploymentOpts: *deploymentTest.Opts,
		downloadStore:  mockStore,
		Host:           "test",
		Name:           "mongodb.gz",
	}

	downloadOpts.Fs = afero.NewMemMapFs()
	deploymentTest.CommonAtlasMocks(downloadOpts.ProjectID)

	b, _ := base64.RawStdEncoding.DecodeString("H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQQAAP//hRFKDQsAAAA") // "hello world" gzipped
	mockStore.
		EXPECT().
		DownloadLog(downloadOpts.newHostLogsParams()).
		Return(io.NopCloser(bytes.NewReader(b)), nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDownloadOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &DownloadOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	deploymentTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentUUID().
		Times(1)

	opts.PostRun()
}
