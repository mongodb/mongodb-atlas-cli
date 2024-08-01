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

//go:build unit

package plugin

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func getTestPlugins() []*plugin.Plugin {
	return []*plugin.Plugin{
		{
			Name:        "plugin1",
			Description: "plugin1 description",
			Version:     "1.4.5",
			Commands: []*plugin.Command{
				{Name: "command1"},
				{Name: "command 2"},
			},
			Github: &plugin.Github{
				Owner: "owner1",
				Name:  "repo1",
			},
		},
		{
			Name:        "plugin2",
			Description: "plugin2 description",
			Version:     "1.2.3",
			Commands: []*plugin.Command{
				{Name: "command3"},
				{Name: "command4"},
			},
			Github: &plugin.Github{
				Owner: "owner2",
				Name:  "repo2",
			},
		},
		{
			Name:        "plugin3",
			Description: "plugin3 description",
			Version:     "1.2.3",
			Commands: []*plugin.Command{
				{Name: "command5"},
				{Name: "command6"},
			},
			Github: &plugin.Github{
				Owner: "owner3",
				Name:  "repo3",
			},
		},
	}
}

func Test_findPluginWithGithubValues(t *testing.T) {
	tests := []struct {
		name      string
		owner     string
		nameValue string
		want      string
		wantErr   bool
	}{
		{
			name:      "Plugin found",
			owner:     "owner3",
			nameValue: "repo3",
			want:      "plugin3",
			wantErr:   false,
		},
		{
			name:      "Plugin not found",
			owner:     "owner4",
			nameValue: "repo4",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Opts{
				plugins: getTestPlugins(),
			}
			got, err := opts.findPluginWithGithubValues(tt.owner, tt.nameValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginByGithubValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want {
				t.Errorf("findPluginByGithubValues() = %v, want %v", got.Name, tt.want)
			}
			if got == nil && tt.want != "" {
				t.Errorf("findPluginByGithubValues() = nil, want %v", tt.want)
			}
		})
	}
}

func Test_findPluginWithName(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
		wantErr  bool
	}{
		{
			testName: "Plugin found",
			name:     "plugin3",
			want:     "plugin3",
			wantErr:  false,
		},
		{
			testName: "Plugin not found",
			name:     "plugin4",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			opts := &Opts{
				plugins: getTestPlugins(),
			}
			got, err := opts.findPluginWithName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want {
				t.Errorf("findPluginByName() = %v, want %v", got.Name, tt.want)
			}
			if got == nil && tt.want != "" {
				t.Errorf("findPluginByName() = nil, want %v", tt.want)
			}
		})
	}
}

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(nil, nil),
		3,
		[]string{},
	)
}
