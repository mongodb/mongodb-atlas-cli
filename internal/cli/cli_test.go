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

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAliases(t *testing.T) {
	type args struct {
		use   string
		extra []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single word",
			args: args{use: "words"},
			want: []string{"word"},
		},
		{
			name: "camel case",
			args: args{use: "camelCases"},
			want: []string{"camelcases", "camel-cases", "camelCase", "camelcase", "camel-case"},
		},
		{
			name: "camel case with extra",
			args: args{use: "camelCases", extra: []string{"extra"}},
			want: []string{"camelcases", "camel-cases", "camelCase", "camelcase", "camel-case", "extra"},
		},
	}
	for _, tt := range tests {
		want := tt.want
		args := tt.args
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GenerateAliases(args.use, args.extra...)
			assert.Equal(t, want, got)
		})
	}
}
