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

// +build unit

package atlas

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/spf13/afero"
)

func TestLogsDownloadOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockLogsDownloader(ctrl)

	appFS := afero.NewMemMapFs()

	opts := &LogsDownloadOpts{
		name:  "mongo.gz",
		fs:    appFS,
		store: mockStore,
	}

	f, err := opts.newWriteCloser()
	if err != nil {
		t.Fatalf("newWriteCloser() unexpected error: %v", err)
	}

	mockStore.
		EXPECT().
		DownloadLog(opts.ProjectID, opts.host, opts.name, f, opts.newDateRangeOpts()).
		Return(nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
