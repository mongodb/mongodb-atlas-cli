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

package aws

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
)

const (
	provider            = "AWS"
	privateEndpointType = "DATA_LAKE"
)

func Builder() *cobra.Command {
	const use = "aws"
	cmd := &cobra.Command{
		Use:   use,
		Short: "Manage AWS private endpoints.",
	}
	cmd.AddCommand(
		CreateBuilder(),
		ListBuilder(),
		DescribeBuilder(),
		DeleteBuilder(),
	)

	if config.ToolName == config.AtlasCLI {
		cmd.Deprecated = "Please use 'atlas datafederation privateendpoints'"
	}

	return cmd
}
