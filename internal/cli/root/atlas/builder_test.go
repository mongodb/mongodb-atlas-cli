// Copyright 2022 MongoDB Inc
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

//go:build unit
// +build unit

package atlas

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	type args struct {
		argsWithoutProg []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			want: 34,
			args: args{
				argsWithoutProg: []string{""},
			},
		},
		{
			name: "config",
			want: 34,
			args: args{
				argsWithoutProg: []string{"config"},
			},
		},
		{
			name: "auth",
			want: 34,
			args: args{
				argsWithoutProg: []string{"auth"},
			},
		},
		{
			name: "clusters",
			want: 34,
			args: args{
				argsWithoutProg: []string{"clusters"},
			},
		},
		{
			name: "autocomplete",
			want: 34,
			args: args{
				argsWithoutProg: []string{"__complete"},
			},
		},
		{
			name: "completion",
			want: 34,
			args: args{
				argsWithoutProg: []string{"completion"},
			},
		},
		{
			name: "--version",
			want: 34,
			args: args{
				argsWithoutProg: []string{"completion"},
			},
		},
	}
	var profile string
	for _, tt := range tests {
		name := tt.name
		args := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			got := Builder(&profile, args.argsWithoutProg)
			if len(got.Commands()) != want {
				t.Fatalf("got=%d, want=%d", len(got.Commands()), want)
			}
		})
	}
}
