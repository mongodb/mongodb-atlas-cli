// Copyright 2023 MongoDB Inc
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

package deployments

import (
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/search"
	"github.com/spf13/cobra"
)

var (
	errCompassNotInstalled = errors.New("did not find MongoDB Compass, install: https://dochub.mongodb.org/core/install-compass")
	errMongoshNotInstalled = errors.New("did not find mongosh, install: https://dochub.mongodb.org/core/install-mongosh")
)

func Builder() *cobra.Command {
	const use = "deployments"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage cloud and local deployments.",
	}

	cmd.AddGroup(&cobra.Group{ID: "all", Title: "Cloud and local deployments commands:"})
	cmd.AddGroup(&cobra.Group{ID: "local", Title: "Local deployments commands:"})

	cmd.AddCommand(
		SetupBuilder(),
		DeleteBuilder(),
		ListBuilder(),
		ConnectBuilder(),
		DiagnosticsBuilder(),
		StartBuilder(),
		PauseBuilder(),
		search.Builder(),
	)

	return cmd
}
