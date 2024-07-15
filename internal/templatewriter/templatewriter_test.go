// Copyright 2024 MongoDB Inc
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

package templatewriter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type printTest struct {
	name     string
	template string
	data     any
	expected string
	wantErr  require.ErrorAssertionFunc
}

func Test_Print(t *testing.T) {
	var buf bytes.Buffer

	tests := []printTest{
		{
			name:     "primitive data",
			template: "name: {{.}}",
			data:     "Jane",
			expected: "name: Jane",
			wantErr:  require.NoError,
		},
		{
			name:     "nil data",
			template: "name: {{.}}",
			data:     nil,
			expected: "name: <no value>",
			wantErr:  require.NoError,
		},
		{
			name:     "pointer of non empty slice",
			template: "items: {{range .Items}}{{.}} {{end}}",
			data:     struct{ Items *[]string }{Items: &[]string{"AWS", "GCP", "Azure"}},
			expected: "items: AWS GCP Azure ",
			wantErr:  require.NoError,
		},
		{
			name:     "nil pointer of slice",
			template: "items: {{range .Items}}{{.}} {{end}}",
			data:     struct{ Items *[]string }{Items: nil},
			expected: "",
			wantErr:  require.Error, // expected to fail, as Items is nil
		},
		{
			name:     "nil pointer of slice",
			template: "items: {{range valueOrEmptySlice .Items}}{{.}} {{end}}",
			data:     struct{ Items *[]string }{Items: nil},
			expected: "items: ",
			wantErr:  require.NoError,
		},
		{
			name:     "non empty slice",
			template: "items: {{range valueOrEmptySlice .Items}}{{.}} {{end}}",
			data:     struct{ Items []string }{Items: []string{"AWS", "GCP", "Azure"}},
			expected: "items: AWS GCP Azure ",
			wantErr:  require.NoError,
		},
		{
			name:     "empty slice",
			template: "items: {{range valueOrEmptySlice .Items}}{{.}} {{end}}",
			data:     struct{ Items []string }{Items: []string{}},
			expected: "items: ",
			wantErr:  require.NoError,
		},
	}

	for _, conf := range tests {
		t.Run(conf.name, func(t *testing.T) {
			conf.wantErr(t, Print(&buf, conf.template, conf.data))
			assert.Equal(t, conf.expected, buf.String())
		})
		buf.Reset()
	}
}
