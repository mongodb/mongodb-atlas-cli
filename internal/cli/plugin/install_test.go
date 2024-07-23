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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_InstallBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		InstallBuilder(nil, nil),
		0,
		[]string{},
	)
}

func Test_validateForExistingPlugins(t *testing.T) {
	github1 := &plugin.Github{Name: "repoName1", Owner: "repoOwner1"}
	github2 := &plugin.Github{Name: "repoName2", Owner: "repoOwner2"}

	plugins := []*plugin.Plugin{
		{
			Name:        "plugin1",
			Description: "plugin1 description",
			Version:     "1.4.5",
			Github:      github1,
			Commands: []*plugin.Command{
				{Name: "command1"},
				{Name: "command 2"},
			},
		},
		{
			Name:        "plugin2",
			Description: "plugin2 description",
			Version:     "1.2.3",
			Github:      github2,
			Commands: []*plugin.Command{
				{Name: "command3"},
				{Name: "command4"},
			},
		},
	}

	opts := InstallOpts{plugins: plugins}

	opts.repositoryName = github1.Name
	opts.repositoryOwner = github1.Owner
	err := opts.validateForExistingPlugins()
	require.Error(t, err)

	opts.repositoryName = github2.Name
	opts.repositoryOwner = "differentOwner"
	err = opts.validateForExistingPlugins()
	assert.NoError(t, err)
}
