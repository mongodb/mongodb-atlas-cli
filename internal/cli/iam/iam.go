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

package iam

import (
	"github.com/mongodb/mongocli/internal/cli/iam/globalaccesslist"
	"github.com/mongodb/mongocli/internal/cli/iam/globalapikeys"
	"github.com/mongodb/mongocli/internal/cli/iam/organizations"
	"github.com/mongodb/mongocli/internal/cli/iam/projects"
	"github.com/mongodb/mongocli/internal/cli/iam/teams"
	"github.com/mongodb/mongocli/internal/cli/iam/users"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
)

const (
	iam  = "Organization and projects operations."
	long = "Identity and Access Management."
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use: "iam",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return validate.Credentials()
		},
		Short: iam,
		Long:  long,
	}
	cmd.AddCommand(
		projects.Builder(),
		organizations.Builder(),
		globalapikeys.Builder(),
		globalaccesslist.Builder(),
		users.Builder(),
		teams.Builder(),
	)

	return cmd
}
