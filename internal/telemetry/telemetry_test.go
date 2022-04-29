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
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/spf13/cobra"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

/*
const cacheDir = "/path/to/mock/dir"

func TestTelemetry_Save(t *testing.T) {
	config.ToolName = config.AtlasCLI
	fs = afero.NewMemMapFs()
	now := time.Now()
	var properties = map[string]interface{}{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  now.Format(time.RFC3339Nano),
		Source:     config.ToolName,
		Name:       config.ToolName + "-event",
		Properties: properties,
	}
	a := assert.New(t)
	a.NoError(save(event, cacheDir))
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.True(info.Size() > minExpectedSize)
}

func TestTelemetry_Save_MaxCacheFileSize(t *testing.T) {
	config.ToolName = config.AtlasCLI
	fs = afero.NewMemMapFs()
	now := time.Now()
	var properties = map[string]interface{}{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  now.Format(time.RFC3339Nano),
		Source:     config.ToolName,
		Name:       config.ToolName + "-event",
		Properties: properties,
	}
	maxCacheFileSize = 10 // 10 bytes
	a := assert.New(t)
	// First save will work as the cache file will be new
	a.NoError(save(event, cacheDir))
	// Second save should fail as the file will be larger than 10 bytes
	a.Error(save(event, cacheDir))
}

func TestTelemetry_OpenCacheFile(t *testing.T) {
	config.ToolName = config.AtlasCLI
	fs = afero.NewMemMapFs()
	a := assert.New(t)
	_, err := openCacheFile(cacheDir)
	a.NoError(err)
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := fs.Stat(filename)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file is empty
	var expectedSize int64 // The nil value is zero
	a.Equal(info.Size(), expectedSize)
}
*/

func TestTelemetry_TrackCommand(t *testing.T) {
	config.ToolName = config.AtlasCLI
	config.SetTelemetryEnabled(true)
	fs = afero.NewMemMapFs()
	cmd := cobra.Command{
		Use: "test-command",
	}
	TrackCommand(&cmd)
	a := assert.New(t)
	// Verify that the file exists
	cacheDir, err := os.UserCacheDir()
	a.NoError(err)
	cacheDir = path.Join(cacheDir, config.ToolName)
	filename := path.Join(cacheDir, cacheFilename)
	// TODO: Temporary print to debug unit test failure on Windows
	fmt.Printf("--- filename: %s\n", filename)
	info, statError := fs.Stat(filename)
	fmt.Printf("--- statErr: %+v\n", statError)
	a.NoError(statError)
	// Verify the file name
	a.Equal(info.Name(), cacheFilename)
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.True(info.Size() > minExpectedSize)
}
