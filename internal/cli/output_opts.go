// Copyright 2020 MongoDB Inc
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

package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PaesslerAG/jsonpath"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/jsonpathwriter"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/jsonwriter"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/templatewriter"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/spf13/cobra"
)

const (
	jsonFormat      = "json"
	plaintextFormat = "plaintext"
	jsonPath        = "json-path"
	goTemplate      = "go-template"
	goTemplateFile  = "go-template-file"
)

var templateFormats = []string{goTemplate, goTemplateFile, jsonPath}

type OutputOpts struct {
	Template  string
	OutWriter io.Writer
	Output    string
}

// InitOutput allow to init the OutputOpts in a functional way.
func (opts *OutputOpts) InitOutput(w io.Writer, t string) func() error {
	return func() error {
		opts.Template = t
		opts.OutWriter = w
		return nil
	}
}

func (*OutputOpts) AutoCompleteOutputFlag() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "json-path", "go-template", "go-template-file"}, cobra.ShellCompDirectiveDefault
	}
}

// ConfigOutput returns the output format.
// If the format is empty, it caches it after querying config.
func (opts *OutputOpts) ConfigOutput() string {
	if opts.Output != "" {
		return opts.Output
	}
	opts.Output = config.Output()
	return opts.Output
}

func (opts *OutputOpts) IsJSONOutput() bool {
	return opts.ConfigOutput() == jsonFormat
}

func (opts *OutputOpts) IsGoTemplate() bool {
	return opts.ConfigOutput() == goTemplate || opts.ConfigOutput() == goTemplateFile
}

func (opts *OutputOpts) IsJSONPathOutput() bool {
	return opts.ConfigOutput() == jsonPath
}

func (opts *OutputOpts) IsPlainOutput() bool {
	return !(opts.IsJSONOutput() || opts.IsGoTemplate() || opts.IsJSONPathOutput())
}

// ConfigWriter returns the io.Writer.
// If the writer is nil, it defaults to os.Stdout and caches it.
func (opts *OutputOpts) ConfigWriter() io.Writer {
	if opts.OutWriter != nil {
		return opts.OutWriter
	}
	opts.OutWriter = os.Stdout
	return opts.OutWriter
}

// IsTerminal returns true is the current file descriptor is TTY kind of terminal.
func (opts *OutputOpts) IsTerminal() bool {
	return terminal.IsTerminal(opts.OutWriter)
}

// IsCygwinTerminal returns true is the current file descriptor is cygwin.
func (opts *OutputOpts) IsCygwinTerminal() bool {
	return terminal.IsCygwinTerminal(opts.OutWriter)
}

// Print will evaluate the defined format and try to parse it accordingly outputting to the set writer.
func (opts *OutputOpts) Print(o any) error {
	if opts.ConfigOutput() == jsonFormat {
		return jsonwriter.Print(opts.ConfigWriter(), o)
	}

	outputType, val := opts.outputTypeAndValue()
	if outputType == jsonPath {
		return jsonpathwriter.Print(opts.ConfigWriter(), val, o)
	}

	t, err := template(outputType, val)
	if err != nil {
		return err
	}

	if t != "" {
		return templatewriter.Print(opts.ConfigWriter(), t, o)
	}
	_, err = fmt.Fprintln(opts.ConfigWriter(), o)
	return err
}

func (opts *OutputOpts) PrintForCompactResultsResponse(o any) error {
	if opts.ConfigOutput() == jsonFormat {
		compactResponse, err := mapReduceResults(o)
		if err == nil {
			return jsonwriter.Print(opts.ConfigWriter(), compactResponse)
		}
	}

	outputType, val := opts.outputTypeAndValue()
	if outputType == jsonPath {
		compactResponse, err := mapReduceResults(o)
		if err == nil {
			return jsonpathwriter.Print(opts.ConfigWriter(), val, compactResponse)
		}
	}

	t, err := template(outputType, val)
	if err != nil {
		return err
	}

	if t != "" {
		return templatewriter.Print(opts.ConfigWriter(), t, o)
	}
	_, err = fmt.Fprintln(opts.ConfigWriter(), o)
	return err
}

func mapReduceResults(o any) (any, error) {
	jsonString, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var val any
	if e := json.Unmarshal(jsonString, &val); e != nil {
		return nil, e
	}

	return jsonpath.Get("results", val)
}

// outputTypeAndValue returns the output type and the associated value
// Current available output types are:
// "go-template=Template string",
// "go-template-file=path/to/template",
// and "json-path=path".
func (opts *OutputOpts) outputTypeAndValue() (outputType, v string) {
	v = opts.Template
	for _, format := range templateFormats {
		format += "="
		if strings.HasPrefix(opts.ConfigOutput(), format) {
			v = opts.ConfigOutput()[len(format):]
			outputType = format[:len(format)-1]
			break
		}
	}

	return
}

var errTemplate = errors.New("error loading template")

// template returns the correct template from the output type.
func template(outputType, value string) (string, error) {
	if outputType == goTemplateFile {
		data, err := os.ReadFile(value)
		if err != nil {
			return "", fmt.Errorf("%w: %s, %w", errTemplate, value, err)
		}

		value = string(data)
	}
	return value, nil
}
