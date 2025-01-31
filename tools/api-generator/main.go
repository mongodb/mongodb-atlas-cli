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
	"fmt"
	"go/format"
	"io"
	"os"
	"text/template"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/spf13/cobra"
)

//go:embed commands.go.tmpl
var templateContent string

func main() {
	var specPath string

	var rootCmd = &cobra.Command{
		Use:   "api-generator",
		Short: "CLI which generates api command definitions from a OpenAPI spec",
		RunE: func(command *cobra.Command, _ []string) error {
			return run(command.Context(), specPath, command.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVar(&specPath, "spec", "", "Path to spec file")
	_ = rootCmd.MarkFlagRequired("spec")
	_ = rootCmd.MarkFlagFilename("spec")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, specPath string, w io.Writer) error {
	specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer specFile.Close()

	return convertSpecToAPICommands(ctx, specFile, w)
}

func convertSpecToAPICommands(ctx context.Context, r io.Reader, w io.Writer) error {
	spec, err := loadSpec(r)
	if err != nil {
		return fmt.Errorf("failed to load spec, error: %w", err)
	}

	if err := spec.Validate(ctx, openapi3.DisableSchemaPatternValidation(), openapi3.DisableExamplesValidation()); err != nil {
		return fmt.Errorf("spec validation failed, error: %w", err)
	}

	commands, err := specToCommands(spec)
	if err != nil {
		return fmt.Errorf("failed convert spec to api commands: %w", err)
	}

	return writeCommands(w, commands)
}

func loadSpec(r io.Reader) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	return loader.LoadFromIoReader(r)
}

func writeCommands(w io.Writer, data api.GroupedAndSortedCommands) error {
	tmpl, err := template.New("commands.go.tmpl").Funcs(template.FuncMap{
		"currentYear": func() int {
			return time.Now().UTC().Year()
		},
	}).Parse(templateContent)
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

	_, err = w.Write(formatted)
	return err
}
