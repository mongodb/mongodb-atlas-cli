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

package logs

import (
	"github.com/spf13/cobra"
)

func JobsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "jobs",
		Aliases: []string{"job"},
		Short:   LogCollection,
	}

	cmd.AddCommand(JobsCollectOptsBuilder())
	cmd.AddCommand(JobsListOptsBuilder())
	cmd.AddCommand(JobsDownloadOptsBuilder())
	cmd.AddCommand(JobsDeleteOptsBuilder())

	return cmd
}
