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
	"path/filepath"
	"slices"
	"strings"
	"text/template"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/speakeasy-api/openapi-overlay/pkg/overlay"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

//go:embed commands.go.tmpl
var commandsTemplateContent string

//go:embed metadata.go.tmpl
var metadataTemplateContent string

type OutputType string

const (
	Commands OutputType = "commands"
	Metadata OutputType = "metadata"
)

// Returns all possible values of OutputType.
func AllOutputTypes() []OutputType {
	return []OutputType{Commands, Metadata}
}

func main() {
	var (
		specPath      string
		overlayPath   string
		outputTypeStr string
	)

	var rootCmd = &cobra.Command{
		Use:   "api-generator",
		Short: "CLI which generates api command definitions from a OpenAPI spec",
		RunE: func(command *cobra.Command, _ []string) error {
			outputType := OutputType(outputTypeStr)
			if !slices.Contains(AllOutputTypes(), outputType) {
				return fmt.Errorf("'%s' is not a valid output type", outputType)
			}

			return run(command.Context(), specPath, overlayPath, outputType, command.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVar(&specPath, "spec", "", "Path to spec file")
	rootCmd.Flags().StringVar(&overlayPath, "overlay", "", "Path to overlay folder")
	_ = rootCmd.MarkFlagRequired("spec")
	_ = rootCmd.MarkFlagFilename("spec")

	rootCmd.Flags().StringVar(&outputTypeStr, "output-type", "", "Set output type [overlay/commands/metadata]")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, specPath, overlayPath string, outputType OutputType, w io.Writer) error {
	specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer specFile.Close()

	files, err := os.ReadDir(overlayPath)
	if err != nil {
		return err
	}

	overlayFiles := make([]io.Reader, 0, len(files))

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yaml") && !strings.HasSuffix(file.Name(), ".yml") {
			continue
		}

		fileName := filepath.Join(overlayPath, file.Name())

		overlayFile, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}

		defer overlayFile.Close() //nolint // required

		overlayFiles = append(overlayFiles, overlayFile)
	}

	switch outputType {
	case Commands:
		return convertSpecToAPICommands(ctx, specFile, overlayFiles, w)
	case Metadata:
		return convertSpecToMetadata(ctx, specFile, overlayFiles, w)
	default:
		return fmt.Errorf("'%s' is not a valid outputType", outputType)
	}
}

func applyOverlays(r io.Reader, overlayFiles []io.Reader) (io.Reader, error) {
	if len(overlayFiles) == 0 {
		return r, nil
	}

	var spec yaml.Node
	err := yaml.NewDecoder(r).Decode(&spec)
	if err != nil {
		return nil, err
	}

	for _, overlayFile := range overlayFiles {
		var o overlay.Overlay
		dec := yaml.NewDecoder(overlayFile)

		if err := dec.Decode(&o); err != nil {
			return nil, err
		}

		if err := o.Validate(); err != nil {
			return nil, err
		}

		if err = o.ApplyTo(&spec); err != nil {
			return nil, err
		}
	}

	buf, err := yaml.Marshal(&spec)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(buf), nil
}

func convertSpecToAPICommands(ctx context.Context, r io.Reader, overlayFiles []io.Reader, w io.Writer) error {
	return convertSpec(ctx, r, overlayFiles, w, specToCommands, commandsTemplateContent)
}

func convertSpecToMetadata(ctx context.Context, r io.Reader, overlayFiles []io.Reader, w io.Writer) error {
	return convertSpec(ctx, r, overlayFiles, w, specToMetadata, metadataTemplateContent)
}

func convertSpec[T any](ctx context.Context, r io.Reader, overlayFiles []io.Reader, w io.Writer, mapper func(spec *openapi3.T) (T, error), templateContent string) error {
	overlaySpec, err := applyOverlays(r, overlayFiles)
	if err != nil {
		return fmt.Errorf("failed to apply overlays, error: %w", err)
	}

	spec, err := loadSpec(overlaySpec)
	if err != nil {
		return fmt.Errorf("failed to load spec, error: %w", err)
	}

	if templateContent != metadataTemplateContent {
		if err := spec.Validate(ctx, openapi3.DisableSchemaPatternValidation(), openapi3.DisableExamplesValidation()); err != nil {
			return fmt.Errorf("spec validation failed, error: %w", err)
		}
	}

	commands, err := mapper(spec)
	if err != nil {
		return fmt.Errorf("failed convert spec to api commands: %w", err)
	}

	return writeCommands(w, templateContent, commands)
}

func loadSpec(r io.Reader) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	return loader.LoadFromIoReader(r)
}

func writeCommands[T any](w io.Writer, templateContent string, data T) error {
	tmpl, err := template.New("output").Funcs(template.FuncMap{
		"currentYear": func() int {
			return time.Now().UTC().Year()
		},
		"replace": func(o, n, s string) string {
			return strings.ReplaceAll(s, o, n)
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
