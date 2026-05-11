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

package cli

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
)

// AnnotationKeyPermission is the Cobra annotation key used to declare a command's pledge tier.
const AnnotationKeyPermission = "atlas.permission"

// SetPermission annotates a Cobra command with the required pledge permission tier.
// Call this in each command's Builder() immediately before returning the command.
func SetPermission(cmd *cobra.Command, tier api.PermissionTier) {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}
	cmd.Annotations[AnnotationKeyPermission] = string(tier)
}
