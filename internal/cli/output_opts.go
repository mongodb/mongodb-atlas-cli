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
	"github.com/mongodb/mongocli/internal/jsonwritter"
	"github.com/mongodb/mongocli/internal/templatewritter"
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

func (opts *OutputOpts) InitOutput(w io.Writer, t string) func() error {
	return func() error {
		opts.Template = t
		opts.OutWriter = w
		return nil
	}
}

func (opts *OutputOpts) ConfigOutput() string {
	if opts.Output != "" {
		return opts.Output
	}
	opts.Output = config.Output()
	return opts.Output
}

func (opts *OutputOpts) ConfigWriter() io.Writer {
	if opts.OutWriter != nil {
		return opts.OutWriter
	}
	opts.OutWriter = os.Stdout
	return opts.OutWriter
}

func (opts *OutputOpts) IsTerminal() bool {
	if f, isFile := opts.OutWriter.(*os.File); isFile {
		return isatty.IsTerminal(f.Fd()) || opts.IsCygwinTerminal()
	}

	return false
}

func (opts *OutputOpts) IsCygwinTerminal() bool {
	if f, isFile := opts.OutWriter.(*os.File); isFile {
		return isatty.IsCygwinTerminal(f.Fd())
	}

	return false
}

func (opts *OutputOpts) Print(v interface{}) error {
	if opts.ConfigOutput() == jsonFormat {
		return jsonwritter.Print(opts.ConfigWriter(), v)
	}
	t, err := opts.parseTemplate()
	if err != nil {
		return err
	}
	if t != "" {
		return templatewritter.Print(opts.ConfigWriter(), t, v)
	}
	_, err = fmt.Fprintln(opts.ConfigWriter(), v)
	return err
}

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
