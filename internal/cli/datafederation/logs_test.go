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

// This code was autogenerated at 2023-06-21T13:32:20+01:00. Note: Manual updates are allowed, but may be overwritten.

package datafederation

import (
	"io"
	"os"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestLogOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataFederationLogDownloader(ctrl)

	const contents = "expected"

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

	const outFilename = "logs.gz"

	fs := afero.NewMemMapFs()

	opts := &LogOpts{
		store: mockStore,
		DownloaderOpts: cli.DownloaderOpts{
			Out: outFilename,
			Fs:  fs,
		},
	}

	mockStore.
		EXPECT().
		DataFederationLogs(opts.ProjectID, opts.id, int64(0), int64(0)).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	of, _ := fs.Open(outFilename)
	defer of.Close()
	b, _ := io.ReadAll(of)
	require.Equal(t, contents, string(b))
}

func TestLogBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LogBuilder(),
		0,
		[]string{flag.ProjectID, flag.Out, flag.Force},
	)
}
