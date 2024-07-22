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

package compass

import (
	"fmt"
	"os/exec"
	"path"

	"golang.org/x/sys/windows/registry"
)

const compassBin = "MongoDBCompass.exe"

func binPath() string {
	// Find the path to the MongoDB Compass binary using the registry
	directory, err := compassDirectory()
	if directory == "" || err != nil {
		// If the directory is not found in the registry, search in the PATH environment variable
		p, err := exec.LookPath(compassBin)
		if err != nil {
			return ""
		}

		return p
	}

	return path.Join(directory, compassBin)
}

// Finds the path to the MongoDB Compass directory.
// The path can be found in the registry:
// Location: HKEY_LOCAL_MACHINE\SOFTWARE\MongoDB\MongoDB Compass
// Key: Directory
func compassDirectory() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\MongoDB\MongoDB Compass`, registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("Error opening registry key: %v", err)
	}
	defer key.Close()

	directory, _, err := key.GetStringValue("Directory")
	if err != nil {
		return "", fmt.Errorf("Error reading Directory value: %v", err)
	}

	return directory, nil
}
