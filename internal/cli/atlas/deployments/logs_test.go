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
	"context"
	"strings"
	"testing"

	"github.com/containers/podman/v4/libpod/define"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/spf13/afero"
)

func TestLogsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LogsBuilder(),
		0,
		[]string{},
	)
}

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()
	expectedLocalDeployment := "localDeployment1"
	buf := new(bytes.Buffer)

	downloadOpts := &DownloadOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		fs: afero.NewMemMapFs(),
	}

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]*define.InspectContainerData{
			{
				Name: options.MongodHostnamePrefix + "-" + expectedLocalDeployment,
				Config: &define.InspectContainerConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				HostConfig: &define.InspectContainerHostConfig{
					PortBindings: map[string][]define.InspectHostPort{
						"27017/tcp": {
							{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
					},
				},
				Mounts: []define.InspectMount{
					{
						Name: downloadOpts.DeploymentOpts.LocalMongodDataVolume(),
					},
				},
			},
		}, nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerLogs(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]string{"log1", "log2"}, nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	expectedOutput := "localDeployment1.log"
	if !strings.Contains(buf.String(), expectedOutput) {
		t.Fatalf("Run() expected output: %s, got: %s", expectedOutput, buf.String())
	}
}
