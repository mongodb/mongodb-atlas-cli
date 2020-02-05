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

package usage

const (
	ProjectID     = "Project ID to use. Overrides configuration file or environment variable settings."
	OrgID         = "Organization ID to use. Overrides configuration file or environment variable settings."
	Profile       = "Profile to use from your configuration file."
	Members       = "Number of members in the replica set."
	InstanceSize  = "Tier for each data-bearing server in the cluster."
	DiskSizeGB    = "Capacity, in gigabytes, of the host’s root volume."
	Backup        = "If true, uses Atlas Continuous Backups to back up cluster data."
	MDBVersion    = "MongoDB version of the cluster to deploy."
	Page          = "Page number."
	Limit         = "Number of items per page."
	Username      = "Username for authenticating to MongoDB."
	Password      = "User’s password."
	Roles         = "User's roles and the databases or collections on which the roles apply."
	Comment       = "Optional description of the whitelist entry."
	Force         = "Don't ask for confirmation."
	Email         = "User’s email address."
	FirstName     = "User’s first name."
	LastName      = "User’s last name."
	WhitelistIps  = "IP addresses to add to the new user’s whitelist."
	WhitelistType = `Type of whitelist entry.
Valid values: cidrBlock|ipAddress`
	Service = `Type of MongoDB service.
Valid values: cloud|cloud-manager|ops-manager`
	Provider = `Name of your cloud service provider.
Valid values: AWS|AZURE|GCP.`
	Region = `Physical location of your MongoDB cluster.
For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws
For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure
For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp`
)
