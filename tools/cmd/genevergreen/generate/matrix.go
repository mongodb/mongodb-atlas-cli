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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var minorVersionRe = regexp.MustCompile(`^\d+\.\d+$`)

// ServerMatrix maps a MongoDB minor version to the distros Atlas CLI publishes for it.
type ServerMatrix map[string][]string

// LoadServerMatrix reads and validates a server-matrix.json file.
func LoadServerMatrix(path string) (ServerMatrix, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m, err := ParseServerMatrix(b)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return m, nil
}

// ParseServerMatrix decodes and validates server matrix JSON.
func ParseServerMatrix(b []byte) (ServerMatrix, error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	var m ServerMatrix
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return nil, errors.New("unexpected data after server matrix")
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return m, nil
}

// Validate checks every version is major.minor and every distro is published by Atlas CLI.
func (m ServerMatrix) Validate() error {
	if len(m) == 0 {
		return errors.New("server matrix is empty")
	}
	for sv, ds := range m {
		if !IsMinorVersion(sv) {
			return fmt.Errorf("invalid version %q, want major.minor", sv)
		}
		if len(ds) == 0 {
			return fmt.Errorf("version %s has no distros", sv)
		}
		seen := map[string]bool{}
		for _, d := range ds {
			if _, ok := distros[d]; !ok {
				return fmt.Errorf("version %s: unknown distro %q", sv, d)
			}
			if seen[d] {
				return fmt.Errorf("version %s: duplicate distro %q", sv, d)
			}
			seen[d] = true
		}
	}
	return nil
}

// Versions returns the matrix versions, oldest first.
func (m ServerMatrix) Versions() []string {
	return slices.SortedFunc(maps.Keys(m), CompareMinorVersions)
}

// Supports reports whether Atlas CLI publishes distro for server version sv.
func (m ServerMatrix) Supports(sv, distro string) bool {
	return slices.Contains(m[sv], distro)
}

// PublishDistros returns the distros Atlas CLI publishes to, sorted.
func PublishDistros() []string {
	return slices.Sorted(maps.Keys(distros))
}

// IsMinorVersion reports whether v has the form major.minor.
func IsMinorVersion(v string) bool {
	return minorVersionRe.MatchString(v)
}

// CompareMinorVersions orders major.minor versions numerically. Both must pass IsMinorVersion.
func CompareMinorVersions(a, b string) int {
	am, an := splitMinorVersion(a)
	bm, bn := splitMinorVersion(b)
	if am != bm {
		return am - bm
	}
	return an - bn
}

func splitMinorVersion(v string) (int, int) {
	major, minor, _ := strings.Cut(v, ".")
	ma, _ := strconv.Atoi(major)
	mi, _ := strconv.Atoi(minor)
	return ma, mi
}
