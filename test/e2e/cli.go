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

package e2e

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func AtlasCLIBin() (string, error) {
	path := os.Getenv("ATLAS_E2E_BINARY")
	cliPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin path %q", err, path)
	}

	if _, err := os.Stat(cliPath); err != nil {
		return "", fmt.Errorf("%w: invalid bin %q", err, path)
	}
	return cliPath, nil
}

func RandInt(max int64) (*big.Int, error) {
	return rand.Int(rand.Reader, big.NewInt(max))
}
