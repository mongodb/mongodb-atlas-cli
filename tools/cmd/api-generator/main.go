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
	"reflect"
	"slices"
	"strings"
	"text/template"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
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

			return run(command.Context(), specPath, outputType, command.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVar(&specPath, "spec", "", "Path to spec file")
	_ = rootCmd.MarkFlagFilename("spec")

	rootCmd.Flags().StringVar(&outputTypeStr, "output-type", "", "Set output type [commands/metadata]")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, specPath string, outputType OutputType, w io.Writer) error {
	var spec io.Reader = os.Stdin
	if specPath != "" {
		specFile, err := os.OpenFile(specPath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		spec = specFile
		defer specFile.Close()
	}

	now := time.Now()

	switch outputType {
	case Commands:
		return convertSpecToAPICommands(ctx, now, spec, w)
	case Metadata:
		return convertSpecToMetadata(ctx, now, spec, w)
	default:
		return fmt.Errorf("'%s' is not a valid outputType", outputType)
	}
}

func convertSpecToAPICommands(ctx context.Context, now time.Time, r io.Reader, w io.Writer) error {
	return convertSpec(ctx, now, r, w, specToCommands, commandsTemplateContent)
}

func convertSpecToMetadata(ctx context.Context, now time.Time, r io.Reader, w io.Writer) error {
	return convertSpec(ctx, now, r, w, func(_ time.Time, spec *openapi3.T) (metadatatypes.Metadata, error) {
		return specToMetadata(spec)
	}, metadataTemplateContent)
}

func convertSpec[T any](ctx context.Context, now time.Time, r io.Reader, w io.Writer, mapper func(now time.Time, spec *openapi3.T) (T, error), templateContent string) error {
	spec, err := loadSpec(r)
	if err != nil {
		return fmt.Errorf("failed to load spec, error: %w", err)
	}

	if templateContent != metadataTemplateContent {
		if err := spec.Validate(ctx, openapi3.DisableSchemaPatternValidation(), openapi3.DisableExamplesValidation()); err != nil {
			return fmt.Errorf("spec validation failed, error: %w", err)
		}
	}

	commands, err := mapper(now, spec)
	if err != nil {
		return fmt.Errorf("failed convert spec to api commands: %w", err)
	}

	return writeCommands(w, templateContent, commands)
}

func loadSpec(r io.Reader) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	return loader.LoadFromIoReader(r)
}

func sortedKeys(v any) []string {
	if reflect.TypeOf(v).Kind() != reflect.Map {
		return nil
	}

	if reflect.TypeOf(v).Key().Kind() != reflect.String {
		return nil
	}

	keys := make([]string, 0, reflect.ValueOf(v).Len())
	for _, key := range reflect.ValueOf(v).MapKeys() {
		keys = append(keys, key.String())
	}

	slices.Sort(keys)

	return keys
}

func writeCommands[T any](w io.Writer, templateContent string, data T) error {
	tmpl, err := template.New("output").Funcs(template.FuncMap{
		"currentYear": func() int {
			return time.Now().UTC().Year()
		},
		"replace": func(o, n, s string) string {
			return strings.ReplaceAll(s, o, n)
		},
		"sortedKeys": sortedKeys,
		"createVersion": func(version api.Version) string {
			switch v := version.(type) {
			case api.PreviewVersion:
				return "shared_api.NewPreviewVersion()"
			case api.UpcomingVersion:
				return fmt.Sprintf("shared_api.NewUpcomingVersion(%d, %d, %d)", v.Date.Year, v.Date.Month, v.Date.Day)
			case api.StableVersion:
				return fmt.Sprintf("shared_api.NewStableVersion(%d, %d, %d)", v.Date.Year, v.Date.Month, v.Date.Day)
			}

			panic("unreachable")
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
