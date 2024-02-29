// Copyright 2021 MongoDB Inc
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

package atlas

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	Use               = "atlas"
	deprecatedMessage = "There's a new, dedicated Atlas CLI available for Atlas users. Install the Atlas CLI to enjoy the same capabilities and keep getting new features: https://dochub.mongodb.org/core/migrate-to-atlas-cli."
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:        Use,
		Short:      "MongoDB Atlas operations.",
		Deprecated: deprecatedMessage,
		RunE: func(_ *cobra.Command, _ []string) error {
			return errors.New("deprecated")
		},
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	return cmd
}
