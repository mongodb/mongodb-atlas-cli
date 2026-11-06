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

func TestE2EManifestFilesAreValid(t *testing.T) {
	manifests, err := walkE2EManifests(func(path string) error {
		manifest, err := loade2eManifest(path)
		if err != nil {
			return err
		}
		if manifest == nil {
			return fmt.Errorf("%s: empty %s", path, e2eManifestName)
		}
		return nil
	})
	require.NoError(t, err)
	assert.Positive(t, manifests)
}

func TestEveryE2EPackageHasManifest(t *testing.T) {
	packages, err := loade2eUnits()
	require.NoError(t, err)

	manifests, err := walkE2EManifests(func(string) error { return nil })
	require.NoError(t, err)

	assert.Equal(t, len(packages), manifests, "every e2e unit must have exactly one "+e2eManifestName+" and there must be no orphan "+e2eManifestName+" files")
}

func walkE2EManifests(visit func(path string) error) (int, error) {
	repoRoot, err := e2eRepoRoot()
	if err != nil {
		return 0, err
	}

	found := 0
	err = filepath.WalkDir(filepath.Join(repoRoot, "test", "e2e"), func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || entry.Name() != e2eManifestName {
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
