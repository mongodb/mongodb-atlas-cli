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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/search"
)

// StableIds is a struct that holds a list of OpenAPI operation IDs.
// We rely on SDK convention to name methods based on operations IDs.
type StableIds struct {
	StableIds []string `json:"stableIds"`
}

// Generating list of used GO SDK operation IDs.
func main() {
	argsNumber := 3
	// Check if the folder path is provided as a command-line argument
	if len(os.Args) != argsNumber {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <folder_path> <output_file>")
		os.Exit(1)
	}

	folderPath := os.Args[1]
	outputFile := os.Args[2]

	if folderPath == "" {
		fmt.Fprintln(os.Stderr, "Please provide a folder path")
		os.Exit(1)
	}
	if outputFile == "" {
		fmt.Fprintln(os.Stderr, "Please provide an output file")
		os.Exit(1)
	}

	// Define the regular expression pattern
	pattern := `s\.clientv2\.[\w\r\n\s]+\.([\w\r\n\s]+)\(`

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error compiling regular expression:", err)
		os.Exit(1)
	}
	// Perform regexp search on all Go files
	stableIds, err := walkFiles(folderPath, regex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error walking through directory:", err)
		os.Exit(1)
	}
	sort.Strings(stableIds.StableIds)

	err = writeStringsToJSONFile(*stableIds, outputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error saving operations file:", err)
		os.Exit(1)
	}
}

func walkFiles(folderPath string, regex *regexp.Regexp) (*StableIds, error) {
	stableIds := StableIds{
		StableIds: []string{},
	}
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading file:", err)
				return nil
			}
			matches := regex.FindAllStringSubmatch(string(content), -1)
			for _, match := range matches {
				if len(match) > 1 {
					value := match[1]
					value = strings.TrimSuffix(value, "WithParams")
					value = strings.TrimSpace(value)
					value = strings.ToLower(value[:1]) + value[1:]
					if !search.StringInSlice(stableIds.StableIds, value) {
						stableIds.StableIds = append(stableIds.StableIds, value)
					}
				}
			}
		}
		return nil
	})
	return &stableIds, err
}

func writeStringsToJSONFile(values StableIds, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(values)
	if err != nil {
		return err
	}

	return nil
}
