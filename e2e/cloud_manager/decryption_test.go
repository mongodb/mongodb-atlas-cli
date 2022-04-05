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

package cloud_manager_test

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"reflect"
)

func generateFileName(dir string, i int, suffix string) string {
	return path.Join(dir, fmt.Sprintf("test%v-%v", i, suffix))
}

func dumpToTemp(files embed.FS, srcDir string, i int, suffix, destDir string) (string, error) {
	inputFile := generateFileName(srcDir, i, suffix)
	outputFile := generateFileName(destDir, i, suffix)

	err := dumpToTempFile(files, inputFile, outputFile)
	if err != nil {
		return "", err
	}

	return outputFile, nil
}

func dumpToTempFile(files embed.FS, srcFile, destFile string) error {
	content, err := files.ReadFile(srcFile)
	if err != nil {
		return err
	}

	return os.WriteFile(destFile, content, fs.ModePerm)
}

func parseJSON(contents []byte) ([]map[string]interface{}, error) {
	res := []map[string]interface{}{}

	s := bufio.NewScanner(bytes.NewReader(contents))
	for s.Scan() {
		var item map[string]interface{}
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

func logsAreEqual(expected, got []byte) (bool, error) {
	expectedLines, err := parseJSON(expected)
	if err != nil {
		return false, err
	}

	gotLines, err := parseJSON(got)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(expectedLines, gotLines), nil
}
