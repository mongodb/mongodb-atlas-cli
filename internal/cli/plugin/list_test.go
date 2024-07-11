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
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func TestList_Run(t *testing.T) {
	plugins := []*plugin.Plugin{
		{
			Name:        "plugin1",
			Description: "plugin1 description",
			Version:     "1.4.5",
			Commands: []*plugin.Command{
				{Name: "command1"},
				{Name: "command 2"},
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
		},
	}

	listOpts := &ListOps{
		plugins: plugins,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: new(bytes.Buffer),
		},
	}

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	test.VerifyOutputTemplate(t, listTemplate, plugins)
}
