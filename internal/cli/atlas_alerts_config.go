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

package cli

import (
	"github.com/spf13/cobra"
)

func AtlasAlertsConfigBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "configs",
		Aliases: []string{"config"},
		Short:   "Manage alerts configuration for your project.",
		Long:    "The configs command provides access to your alerts configurations. You can create, edit, and delete alert configurations.",
	}

	cmd.AddCommand(AtlasAlertsConfigCreateBuilder())
	cmd.AddCommand(AtlasAlertConfigListBuilder())
	cmd.AddCommand(AtlasAlertConfigDeleteBuilder())
	cmd.AddCommand(AtlasAlertConfigsFieldsBuilder())
	cmd.AddCommand(AtlasAlertsConfigUpdateBuilder())

	return cmd
}
