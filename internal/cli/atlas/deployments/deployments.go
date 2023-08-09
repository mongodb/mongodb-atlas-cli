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

package local

import (
	"errors"

	searchindexes "github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/searchindexes"
	"github.com/spf13/cobra"
)

var ErrInstanceNotFound = errors.New("instance not found")

const (
	localUser     = "mongoUser"
	localPassword = "hunter1"
	localURI      = "mongodb://localhost:37017"
)

var (
	localData = map[string]string{"ConnectionString": localURI, "User": localUser, "Password": localPassword}
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deployments",
		Short: "Manage Atlas and local instances.",
	}

	cmd.AddCommand(
		StartBuilder(),
		StopBuilder(),
		ClearBuilder(),
		DescribeBuilder(),
		ConnectBuilder(),
		SampleDataBuilder(),
		searchindexes.Builder(),
	)

	return cmd
}
