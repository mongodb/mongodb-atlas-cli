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

package generate

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2ESidecarFilesAreValid(t *testing.T) {
	sidecars, err := walkE2ESidecars(func(path string) error {
		manifest, err := loade2eManifest(path)
		if err != nil {
			return err
		}
		if manifest == nil {
			return fmt.Errorf("%s: empty %s", path, e2eSidecarName)
		}
		return nil
	})
	require.NoError(t, err)
	assert.Positive(t, sidecars)
}

func TestEveryE2EPackageHasSidecar(t *testing.T) {
	packages, err := loade2ePackages()
	require.NoError(t, err)

	sidecars, err := walkE2ESidecars(func(string) error { return nil })
	require.NoError(t, err)

	assert.Equal(t, len(packages), sidecars, "every e2e package must have exactly one "+e2eSidecarName+" and there must be no orphan "+e2eSidecarName+" files")
}

func walkE2ESidecars(visit func(path string) error) (int, error) {
	repoRoot, err := e2eRepoRoot()
	if err != nil {
		return 0, err
	}

	found := 0
	err = filepath.WalkDir(filepath.Join(repoRoot, "test", "e2e"), func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || entry.Name() != e2eSidecarName {
			return nil
		}
		found++
		return visit(path)
	})
	if err != nil {
		return found, err
	}
	return found, nil
}
