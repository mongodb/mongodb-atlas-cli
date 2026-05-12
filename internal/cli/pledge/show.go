// Copyright 2026 MongoDB Inc
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

package pledge

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/spf13/cobra"
)

// ShowBuilder returns the cobra command for `atlas pledge show`.
func ShowBuilder() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show the active pledge for the current session.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			key, err := pledge.ResolveSessionKey()
			if err != nil {
				return fmt.Errorf("pledge is not supported on this platform: %w", err)
			}
			pf, err := pledge.Load(key)
			if errors.Is(err, pledge.ErrNoPledge) {
				fmt.Fprintln(cmd.OutOrStdout(), "No pledge set for this session.")
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Profile:    %s\n", pf.Profile)
			fmt.Fprintf(cmd.OutOrStdout(), "Max tier:   %s\n", pf.MaxTier)
			fmt.Fprintf(cmd.OutOrStdout(), "Narrowed:   %s\n", pf.NarrowedAt.Format("2006-01-02T15:04:05Z07:00"))
			if len(pf.AllowedOps) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Allowed ops: %v\n", pf.AllowedOps)
			}
			return nil
		},
	}
}
