// Copyright 2022 MongoDB Inc
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

package telemetry

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(nil).
		Times(1)
	require.NoError(t, tr.trackCommand(TrackOptions{}))
}

func TestTrackCommandWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsSender(ctrl)

	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	cmd := &cobra.Command{
		Use: "test-command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return errors.New("test command error")
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())
	require.Error(t, errCmd)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(nil).
		Times(1)

	err = tr.trackCommand(TrackOptions{
		Err: errCmd,
	})
	require.NoError(t, err)
}

func TestTrackCommandWithSendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventsSender(ctrl)

	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())

	require.NoError(t, errCmd)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(errors.New("test send error")).
		Times(1)

	err = tr.trackCommand(TrackOptions{
		Err: errCmd,
	})
	require.NoError(t, err)

	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	assert.Equal(t, cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	assert.Greater(t, info.Size(), minExpectedSize)
}

func TestSave(t *testing.T) {
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}

	var properties = map[string]any{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.AtlasCLI,
		Properties: properties,
	}
	require.NoError(t, tr.save(event))
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.Greater(info.Size(), minExpectedSize)
}

func TestSaveOverMaxCacheFileSize(t *testing.T) {
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	var properties = map[string]any{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.AtlasCLI,
		Properties: properties,
	}

	// First save will work as the cache file will be new
	require.NoError(t, tr.save(event))
	// Second save should fail as the file will be larger than 10 bytes
	require.Error(t, tr.save(event))
}

func TestOpenCacheFile(t *testing.T) {
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	_, err = tr.openCacheFile()
	require.NoError(t, err)
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file is empty
	var expectedSize int64 // The nil value is zero
	a.Equal(expectedSize, info.Size())
}

func TestTrackSurvey(t *testing.T) {
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.AtlasCLI+"*")
	require.NoError(t, err)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())
	require.NoError(t, errCmd)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		cmd:              cmd,
	}

	response := true
	err = tr.trackSurvey(
		&survey.Confirm{Message: "test"},
		&response,
		nil,
	)
	require.NoError(t, err)
	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.Greater(info.Size(), minExpectedSize)
}
