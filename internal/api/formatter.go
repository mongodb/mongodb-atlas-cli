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

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"text/template"
)

var (
	ErrFormatterFailedToParseFormatString = errors.New("failed to parse format string")
)

type Formatter struct {
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(format string, in io.ReadCloser) (io.ReadCloser, error) {
	if isGoTemplate(format) {
		return f.formatJSON(format, in)
	}

	return in, nil
}

func (*Formatter) formatJSON(format string, readerCloser io.ReadCloser) (io.ReadCloser, error) {
	// Make sure the readerCloser reader gets closed
	defer readerCloser.Close()

	// Attempt to parse the format string
	tmpl, err := template.New("formatter-template").Parse(format)
	if err != nil {
		return nil, errors.Join(ErrFormatterFailedToParseFormatString, err)
	}

	// Decode the Reader as a json
	var data any
	if err := json.NewDecoder(readerCloser).Decode(&data); err != nil {
		return nil, err
	}

	// buffer contains the result of the execute template
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return nil, err
	}

	// Transform the buffer into a io.ReadCloser
	reader := bytes.NewReader(buffer.Bytes())
	readCloser := io.NopCloser(reader)

	return readCloser, nil
}

func (*Formatter) ContentType(format string) (string, error) {
	if isGoTemplate(format) {
		return "json", nil
	}

	return format, nil
}

// checks if the format string is a go template or not
// this is a basic check checking if the format string has an opening {{ and closing template }} tag.
func isGoTemplate(format string) bool {
	openTagIdx := strings.Index(format, "{{")
	closeTagIdx := strings.Index(format, "}}")

	return openTagIdx != -1 && closeTagIdx != -1 && openTagIdx < closeTagIdx
}
