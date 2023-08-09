package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	argsNumber := 3
	// Check if the folder path is provided as a command-line argument
	if len(os.Args) != argsNumber {
		fmt.Println("Usage: go run main.go <folder_path> <output_file>")
		os.Exit(1)
	}

	folderPath := os.Args[1]
	outputFile := os.Args[2]

	if folderPath == "" {
		fmt.Println("Please provide a folder path")
		os.Exit(1)
	}
	if outputFile == "" {
		fmt.Println("Please provide an output file")
		os.Exit(1)
	}

	// Define the regular expression pattern
	pattern := `s\.clientv2\.\w+\.(\w+)\(`

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		os.Exit(1)
	}

	stableIds := StableIds{
		StableIds: []string{},
	}
	// Perform regexp search on all Go files
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return nil
			}
			matches := regex.FindAllStringSubmatch(string(content), -1)
			for _, match := range matches {
				if len(match) > 1 {
					value := match[1]
					value = strings.TrimSuffix(value, "WithParams")
					if !slices.Contains(stableIds.StableIds, value) {
						stableIds.StableIds = append(stableIds.StableIds, value)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
		os.Exit(1)
	}
	sort.Strings(stableIds.StableIds)

	err = writeStringsToJSONFile(stableIds, outputFile)
	if err != nil {
		fmt.Println("Error saving operations file:", err)
		os.Exit(1)
	}
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

type StableIds struct {
	StableIds []string `json:"stableIds"`
}
