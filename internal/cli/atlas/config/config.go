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

package config

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage your profile to store access settings for your MongoDB deployment.",
		Long: `Configure settings in a user profile.
All settings are optional. You can specify settings individually by running: 
$ atlas config set --help 

You can also use environment variables (MONGODB_ATLAS_*) when running the tool.
To find out more, see the documentation: https://dochub.mongodb.org/core/atlas-cli-env-variables.`,
		Annotations: map[string]string{
			"toc": "true",
		},
		Args: require.NoArgs,
	}

	cmd.AddCommand(
		InitBuilder(),
		config.SetBuilder(),
		config.ListBuilder(),
		config.DescribeBuilder(),
		config.RenameBuilder(),
		config.DeleteBuilder(),
		config.EditBuilder(),
	)

	return cmd
}
