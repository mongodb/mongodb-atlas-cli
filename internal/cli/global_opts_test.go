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

// +build unit

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalOpts_PreRunE(t *testing.T) {
	t.Run("empty project ID", func(t *testing.T) {
		o := &GlobalOpts{}
		if err := o.PreRunE(); err != errMissingProjectID {
			t.Errorf("Expected err: %#v, got: %#v\n", errMissingProjectID, err)
		}
	})
	t.Run("invalid project ID", func(t *testing.T) {
		o := &GlobalOpts{ProjectID: "1"}
		if err := o.PreRunE(); err == nil {
			t.Errorf("Expected an error\n")
		}
	})
	t.Run("valid project ID", func(t *testing.T) {
		o := &GlobalOpts{ProjectID: "5e98249d937cfc52efdc2a9f"}
		if err := o.PreRunE(); err != nil {
			t.Fatalf("PreRunE() unexpected error %v\n", err)
		}
	})
	t.Run("empty org ID", func(t *testing.T) {
		o := &GlobalOpts{}
		o.PreRunEOrg = true
		if err := o.PreRunE(); err != ErrMissingOrgID {
			t.Errorf("Expected err: %#v, got: %#v\n", ErrMissingOrgID, err)
		}
	})
	t.Run("invalid org ID", func(t *testing.T) {
		o := &GlobalOpts{OrgID: "1"}
		o.PreRunEOrg = true
		if err := o.PreRunE(); err == nil {
			t.Errorf("Expected an error\n")
		}
	})
	t.Run("valid org ID", func(t *testing.T) {
		o := &GlobalOpts{OrgID: "5e98249d937cfc52efdc2a9f"}
		o.PreRunEOrg = true
		if err := o.PreRunE(); err != nil {
			t.Fatalf("PreRunE() unexpected error %v\n", err)
		}
	})
}

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
			got := GenerateAliases(args.use, args.extra...)
			assert.Equal(t, got, want)
		})
	}
}
