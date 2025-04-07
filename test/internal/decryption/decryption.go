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

package decryption

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func GenerateFileName(dir, suffix string) string {
	return path.Join(dir, "test-"+suffix)
}

func DumpToTemp(files embed.FS, srcFile, destFile string) error {
	content, err := files.ReadFile(srcFile)
	if err != nil {
		return err
	}

	return os.WriteFile(destFile, content, fs.ModePerm)
}

func parseJSON(contents []byte) ([]map[string]any, error) {
	var res []map[string]any

	s := bufio.NewScanner(bytes.NewReader(contents))
	for s.Scan() {
		var item map[string]any
		err := json.Unmarshal(s.Bytes(), &item)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return res, nil
}

func LogsAreEqual(t *testing.T, expected, got []byte) {
	t.Helper()
	expectedLines, err := parseJSON(expected)
	require.NoError(t, err)

	gotLines, err := parseJSON(got)
	require.NoError(t, err)

	if diff := deep.Equal(expectedLines, gotLines); diff != nil {
		t.Logf("=== expected audit lines ===\n%s\n=== end expected lines ===", expected)
		t.Logf("=== actual audit lines ===\n%s\n=== end actual lines ===", string(got))
		t.Error(diff)
	}
}
