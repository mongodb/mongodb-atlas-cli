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

package create

import (
	"github.com/spf13/cobra"
)

const longDesc = `The network peering create command checks if a VPC exists in the region you specify for your Atlas project. If one exists, this command creates the peering connection between that VPC and your VPC. If an Atlas VPC doesn't exist, this command creates one and creates a connection between it and your VPC.
		
To learn more about network peering connections, see https://www.mongodb.com/docs/atlas/security-vpc-peering/.

`

func Builder() *cobra.Command {
	const use = "create"
	cmd := &cobra.Command{
		Use:   use,
		Short: "Create a connection with AWS, Azure and Google Cloud.",
	}
	cmd.AddCommand(
		AzureBuilder(),
		AwsBuilder(),
		GCPBuilder(),
	)

	return cmd
}
