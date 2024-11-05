// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/autogeneration/L1"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

//go:embed commands.go.tmpl
var templateContent string

func main() {
	var (
		specPath   string
		outputPath string
	)

	var rootCmd = &cobra.Command{
		Use:   "L1-generator",
		Short: "CLI which generates L1 command definitions from a OpenAPI spec",
		PreRunE: func(_ *cobra.Command, _ []string) error {
			if specPath == "" {
				return errors.New("--spec is a required flag")
			}

			if outputPath == "" {
				return errors.New("--output is a required flag")
			}

			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return convertSpecToL1Commands(afero.NewOsFs(), specPath, outputPath)
		},
	}

	rootCmd.PersistentFlags().StringVar(&specPath, "spec", "", "Path to spec file")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output", "", "Path to output file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertSpecToL1Commands(fs afero.Fs, specPath, outputPath string) error {
	spec, err := loadSpec(fs, specPath)
	if err != nil {
		return fmt.Errorf("failed to load spec: '%s', error: %w", specPath, err)
	}

	commands, err := specToCommands(spec)
	if err != nil {
		return fmt.Errorf("failed convert spec to L1 commands: %w", err)
	}

	return writeCommands(fs, outputPath, commands)
}

func loadSpec(fs afero.Fs, specPath string) (*openapi3.T, error) {
	file, err := fs.Open(specPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	loader := openapi3.NewLoader()
	return loader.LoadFromIoReader(file)
}

func writeCommands(fs afero.Fs, outputPath string, data L1.GroupedAndSortedCommands) error {
	tmpl, err := template.New("commands.go.tmpl").Parse(templateContent)
	if err != nil {
		return err
	}

	file, err := fs.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to open output file %s, error: %w", outputPath, err)
	}

	defer func() {
		_ = file.Close()
	}()

	return tmpl.Execute(file, data)
}
