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

package snapshots

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/spf13/cobra"
)

const (
	CannotUseNotFlexWithFlexApisErrorCode = "CANNOT_USE_NON_FLEX_CLUSTER_IN_FLEX_API"
	FeatureUnsupported                    = "FEATURE_UNSUPPORTED"
	ClusterNotFoundErrorCode              = "CLUSTER_NOT_FOUND"
)

func Builder() *cobra.Command {
	const use = "snapshots"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage cloud backup snapshots for your project.",
	}

	cmd.AddCommand(
		ListBuilder(),
		CreateBuilder(),
		DescribeBuilder(),
		WatchBuilder(),
		DeleteBuilder(),
		DownloadBuilder(),
	)

	return cmd
}
