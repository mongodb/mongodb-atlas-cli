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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestTelemetry_Track(t *testing.T) {
	config.ToolName = config.AtlasCLI

	a := assert.New(t)
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.ToolName+"*")
	a.NoError(err)

	tracker := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}

	cmd := cobra.Command{
		Use: "test-command",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	err = tracker.track(TrackOptions{
		Cmd: &cmd,
	})
	a.NoError(err)
	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tracker.fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.True(info.Size() > minExpectedSize)
}

func TestTelemetry_TrackWithError(t *testing.T) {
	config.ToolName = config.AtlasCLI

	a := assert.New(t)
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.ToolName+"*")
	a.NoError(err)

	tracker := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}

	cmd := cobra.Command{
		Use: "test-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("test")
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())

	err = tracker.track(TrackOptions{
		Cmd: &cmd,
		Err: errCmd,
	})
	a.NoError(err)

	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tracker.fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.True(info.Size() > minExpectedSize)
}

func TestTelemetry_Save(t *testing.T) {
	config.ToolName = config.AtlasCLI

	a := assert.New(t)
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.ToolName+"*")
	a.NoError(err)

	tracker := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}

	var properties = map[string]interface{}{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.ToolName,
		Properties: properties,
	}
	a.NoError(tracker.save(event))
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tracker.fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.True(info.Size() > minExpectedSize)
}

func TestTelemetry_Save_MaxCacheFileSize(t *testing.T) {
	config.ToolName = config.AtlasCLI

	a := assert.New(t)
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.ToolName+"*")
	a.NoError(err)

	tracker := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	var properties = map[string]interface{}{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.ToolName,
		Properties: properties,
	}

	// First save will work as the cache file will be new
	a.NoError(tracker.save(event))
	// Second save should fail as the file will be larger than 10 bytes
	a.Error(tracker.save(event))
}

func TestTelemetry_OpenCacheFile(t *testing.T) {
	config.ToolName = config.AtlasCLI

	a := assert.New(t)
	cacheDir, err := os.MkdirTemp(os.TempDir(), config.ToolName+"*")
	a.NoError(err)

	tracker := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	_, err = tracker.openCacheFile()
	a.NoError(err)
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tracker.fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file is empty
	var expectedSize int64 // The nil value is zero
	a.Equal(info.Size(), expectedSize)
}
