// Copyright 2026 MongoDB Inc
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

package pledge

import (
	"os"
	"path/filepath"
	"runtime"
)

// StateDir returns the base directory for pledge state files.
// It honours XDG_STATE_HOME on Linux, uses ~/Library/Application Support on macOS,
// and falls back to ~/.local/state on other unix systems.
func StateDir() (string, error) {
	if dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR"); dir != "" {
		// Allow tests to override.
		return dir, nil
	}

	var base string
	switch runtime.GOOS {
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, "Library", "Application Support")
	default:
		if xdg := os.Getenv("XDG_STATE_HOME"); xdg != "" {
			base = xdg
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			base = filepath.Join(home, ".local", "state")
		}
	}

	return filepath.Join(base, "atlascli", "pledge"), nil
}

// ensureDir creates dir with mode 0700 if it does not exist.
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o700)
}

// atomicWrite writes data to path via a tmp file + rename, ensuring the final
// file never contains a partial write.
func atomicWrite(path string, data []byte, perm os.FileMode) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, perm); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
