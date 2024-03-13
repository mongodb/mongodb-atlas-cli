// Copyright 2024 MongoDB Inc
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

package instance

import (
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestDownloadOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsDownloader(ctrl)

	const contents = "expected"
	const fileName = "auditLogs.gz"

	file, err := os.CreateTemp("", "")
	if err != nil {
		require.NoError(t, err)
	}
	filename := file.Name()
	defer os.Remove(filename)
	_, _ = file.WriteString(contents)
	_ = file.Close()

	expected, _ := os.Open(filename)
	defer expected.Close()

	fs := afero.NewMemMapFs()
	downloadOpts := &DownloadOpts{
		store: mockStore,
		DownloaderOpts: cli.DownloaderOpts{
			Out: fileName,
			Fs:  fs,
		},
	}

	downloadOpts.ProjectID = "download-project-id"
	downloadOpts.decompress = true
	downloadOpts.fileName = fileName

	endDate := int64(0)
	startDate := int64(0)

	downloadParams := new(store.DownloadStreamTenantAuditLogsApiParams)
	downloadParams.EndDate = &endDate
	downloadParams.StartDate = &startDate
	downloadParams.TenantName = "streams-tenant"

	mockStore.
		EXPECT().
		DownloadAuditLog(downloadParams).
		Return(expected, nil).
		Times(1)

	if err := downloadOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	of, _ := fs.Open(fileName)
	defer of.Close()
	b, _ := io.ReadAll(of)
	require.Equal(t, contents, string(b))
}

func TestDownloadBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DownloadBuilder(),
		0,
		[]string{flag.Out, flag.Start, flag.End, flag.Force, flag.Decompress, flag.ProjectID},
	)
}
