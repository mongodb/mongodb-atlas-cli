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
	"compress/gzip"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/spf13/afero"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func TestLogsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LogsBuilder(),
		0,
		[]string{
			flag.DeploymentName,
			flag.Output,
		},
	)
}

func TestLogs_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	buf := new(bytes.Buffer)
	expectedLocalDeployment := "localDeployment"
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)
	mockPodman := deploymentTest.MockPodman

	downloadOpts := &DownloadOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.LocalMockFlow(ctx)

	mockPodman.
		EXPECT().
		ContainerLogs(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
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
	ctx := context.Background()
	buf := new(bytes.Buffer)
	atlasDeployment := "localDeployment1"
	mockStore := mocks.NewMockLogsDownloader(ctrl)
	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, atlasDeployment)

	downloadOpts := &DownloadOpts{
		OutputOpts:     cli.OutputOpts{OutWriter: buf},
		GlobalOpts:     cli.GlobalOpts{ProjectID: "ProjectID"},
		DownloaderOpts: cli.DownloaderOpts{},
		DeploymentOpts: *deploymentTest.Opts,
		downloadStore:  mockStore,
		host:           "test",
		name:           "mongodb.gz",
	}

	downloadOpts.Fs = afero.NewMemMapFs()
	downloadOpts.Out = downloadOpts.name
	deploymentTest.CommonAtlasMocks(downloadOpts.ProjectID)

	mockStore.
		EXPECT().
		DownloadLog(gomock.Any(), downloadOpts.newHostLogsParams()).
		Times(1).
		DoAndReturn(func(_ io.Writer, _ *admin.GetHostLogsApiParams) error {
			f, err := downloadOpts.Fs.OpenFile(downloadOpts.Out, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				return err
			}

			zw := gzip.NewWriter(f)
			defer zw.Close()
			zw.Name = downloadOpts.Out

			if _, err = zw.Write([]byte("log\nlog\n")); err != nil {
				return err
			}

			return nil
		})

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
