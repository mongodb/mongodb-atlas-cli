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
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_checkForDuplicatePlugins(t *testing.T) {
	github1 := &plugin.Github{Name: "repoName1", Owner: "repoOwner1"}
	github2 := &plugin.Github{Name: "repoName2", Owner: "repoOwner2"}

	version1, err := semver.NewVersion("1.4.5")
	require.NoError(t, err)
	version2, err := semver.NewVersion("1.2.3")
	require.NoError(t, err)

	plugins := []*plugin.Plugin{
		{
			Name:        "plugin1",
			Description: "plugin1 description",
			Version:     version1,
			Github:      github1,
			Commands: []*plugin.Command{
				{Name: "command1"},
				{Name: "command 2"},
			},
		},
		{
			Name:        "plugin2",
			Description: "plugin2 description",
			Version:     version2,
			Github:      github2,
			Commands: []*plugin.Command{
				{Name: "command3"},
				{Name: "command4"},
			},
		},
	}

	opts := InstallOpts{
		Opts: Opts{
			plugins: &plugin.ValidatedPlugins{
				ValidPlugins: plugins,
			},
		},
	}

	opts.githubAsset = &GithubAsset{name: github1.Name, owner: github1.Owner}
	err = opts.checkForDuplicatePlugins()
	require.Error(t, err)

	opts.githubAsset.name = github2.Name
	opts.githubAsset.owner = "differentOwner"
	err = opts.checkForDuplicatePlugins()
	assert.NoError(t, err)
}
