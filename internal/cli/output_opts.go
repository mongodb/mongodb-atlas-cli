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
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/jsonwriter"
	"github.com/mongodb/mongocli/internal/templatewriter"
)

const (
	jsonFormat     = "json"
	goTemplate     = "go-template"
	goTemplateFile = "go-template-file"
)

var templateFormats = []string{goTemplate, goTemplateFile}

type OutputOpts struct {
	Template  string
	OutWriter io.Writer
	Output    string
}

// InitOutput allow to init the OutputOpts in a functional way
func (opts *OutputOpts) InitOutput(w io.Writer, t string) func() error {
	return func() error {
		opts.Template = t
		opts.OutWriter = w
		return nil
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

// ConfigWriter returns the io.Writer.
// If the writer is nil, it defaults to os.Stdout and caches it.
func (opts *OutputOpts) ConfigWriter() io.Writer {
	if opts.OutWriter != nil {
		return opts.OutWriter
	}
	opts.OutWriter = os.Stdout
	return opts.OutWriter
}

// IsTerminal returns true is the current file descriptor is TTY kind of terminal
func (opts *OutputOpts) IsTerminal() bool {
	if f, isFile := opts.OutWriter.(*os.File); isFile {
		return isatty.IsTerminal(f.Fd()) || opts.IsCygwinTerminal()
	}

	return false
}

// IsCygwinTerminal returns true is the current file descriptor is cygwin
func (opts *OutputOpts) IsCygwinTerminal() bool {
	if f, isFile := opts.OutWriter.(*os.File); isFile {
		return isatty.IsCygwinTerminal(f.Fd())
	}

	return false
}

// Print will evaluate the defined format and try to parse it accordingly outputting to the set writer
func (opts *OutputOpts) Print(v interface{}) error {
	if opts.ConfigOutput() == jsonFormat {
		return jsonwriter.Print(opts.ConfigWriter(), v)
	}
	t, err := opts.parseTemplate()
	if err != nil {
		return err
	}
	if t != "" {
		return templatewriter.Print(opts.ConfigWriter(), t, v)
	}
	_, err = fmt.Fprintln(opts.ConfigWriter(), v)
	return err
}

// parseTemplate will try to find if the given format is a user given template, either by string or file and use it.
// Current available user templates are  "go-template=Template string" and "go-template-file=path/to/template"
func (opts *OutputOpts) parseTemplate() (string, error) {
	value := opts.Template
	templateFormat := ""
	for _, format := range templateFormats {
		format += "="
		if strings.HasPrefix(opts.ConfigOutput(), format) {
			value = opts.ConfigOutput()[len(format):]
			templateFormat = format[:len(format)-1]
			break
		}
	}
	if templateFormat == goTemplateFile {
		data, err := ioutil.ReadFile(value)
		if err != nil {
			return "", fmt.Errorf("error loading template: %s, %v", value, err)
		}

		value = string(data)
	}
	return value, nil
}
