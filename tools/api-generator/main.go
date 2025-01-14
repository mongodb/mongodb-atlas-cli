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
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
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
		Use:   "api-generator",
		Short: "CLI which generates api command definitions from a OpenAPI spec",
		PreRunE: func(_ *cobra.Command, _ []string) error {
			if specPath == "" {
				return errors.New("--spec is a required flag")
			}

			if outputPath == "" {
				return errors.New("--output is a required flag")
			}

			return nil
		},
		RunE: func(command *cobra.Command, _ []string) error {
			return convertSpecToAPICommands(command.Context(), afero.NewOsFs(), specPath, outputPath)
		},
	}

	rootCmd.PersistentFlags().StringVar(&specPath, "spec", "", "Path to spec file")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output", "", "Path to output file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertSpecToAPICommands(ctx context.Context, fs afero.Fs, specPath, outputPath string) error {
	spec, err := loadSpec(fs, specPath)
	if err != nil {
		return fmt.Errorf("failed to load spec: '%s', error: %w", specPath, err)
	}

	if err := spec.Validate(ctx, openapi3.DisableSchemaPatternValidation(), openapi3.DisableExamplesValidation()); err != nil {
		return fmt.Errorf("spec validation failed, error: %w", err)
	}

	commands, err := specToCommands(spec)
	if err != nil {
		return fmt.Errorf("failed convert spec to api commands: %w", err)
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

func writeCommands(fs afero.Fs, outputPath string, data api.GroupedAndSortedCommands) error {
	tmpl, err := template.New("commands.go.tmpl").Parse(templateContent)
	if err != nil {
		return err
	}

	// Write template output to a buffer first
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// Format the generated code
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format generated code: %w", err)
	}

	// Write the formatted code to file
	file, err := fs.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to open output file %s, error: %w", outputPath, err)
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = file.Write(formatted)
	return err
}
