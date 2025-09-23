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

package plugin

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func TestList_Run(t *testing.T) {
	plugins := getTestPlugins(t)

	listOpts := &ListOps{
		Opts: Opts{
			plugins: plugins,
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: new(bytes.Buffer),
		},
	}

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestList_Template(t *testing.T) {
	test.VerifyOutputTemplate(t, listTemplate, getTestPlugins(t).GetValidAndInvalidPlugins())
}

func TestList_TemplateWithAliases(t *testing.T) {
	plugins := getTestPlugins(t)

	listOpts := &ListOps{
		Opts: Opts{
			plugins: plugins,
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: new(bytes.Buffer),
		},
	}

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	output := listOpts.OutWriter.(*bytes.Buffer).String()

	// Check that aliases are included in the output
	if !strings.Contains(output, "[aliases: cmd1, c1]") {
		t.Errorf("Expected to find aliases for command1, but output was: %s", output)
	}

	if !strings.Contains(output, "[aliases: tf]") {
		t.Errorf("Expected to find aliases for command3, but output was: %s", output)
	}
}
