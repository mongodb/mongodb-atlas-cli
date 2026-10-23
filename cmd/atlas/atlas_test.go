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

package main

import "testing"

func TestIsAgentSkillsInvocation(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{name: "skills list", args: []string{"skills", "list"}, want: true},
		{name: "skills install after profile", args: []string{"--profile", "non-cloud", "skills", "install"}, want: true},
		{name: "skills alias", args: []string{"skill", "list"}, want: true},
		{name: "combined flags before non-skills command", args: []string{"-DP", "skills", "clusters", "list"}, want: false},
		{name: "combined flags before skills command", args: []string{"-DP", "non-cloud", "skills", "list"}, want: true},
		{name: "profile shorthand with attached value", args: []string{"-PD", "skills", "list"}, want: true},
		{name: "other command", args: []string{"clusters", "list"}, want: false},
		{name: "profile value that resembles command", args: []string{"--profile", "skills", "clusters", "list"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAgentSkillsInvocation(tt.args); got != tt.want {
				t.Errorf("isAgentSkillsInvocation(%v) = %t, want %t", tt.args, got, tt.want)
			}
		})
	}
}

func TestBuild(t *testing.T) {
	t.Setenv("DO_NOT_TRACK", "1")
	main()
}
