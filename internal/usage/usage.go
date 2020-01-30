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
	ProjectID     = "The project ID to use. Overrides config/env settings."
	OrgID         = "The organization ID to use. Overrides config/env settings."
	Profile       = "Use a specific profile from your configuration file."
	Members       = "Number of replica set members."
	InstanceSize  = "Tier for each data-bearing server in the cluster."
	DiskSizeGB    = "Capacity, in gigabytes, of the host’s root volume."
	Backup        = "If true, the cluster uses Atlas Continuous Backups for backing up cluster data."
	MDBVersion    = "Version of the cluster to deploy."
	Page          = "Page number."
	Limit         = "Items per page."
	Username      = "Username for authenticating to MongoDB."
	Password      = "The user’s password."
	Roles         = "User’s roles and the databases/collections on which the roles apply."
	Comment       = "Optional comment associated with the whitelist entry."
	Force         = "Don't ask for confirmation."
	Email         = "The user’s email"
	FirstName     = "The user’s first name"
	LastName      = "The user’s last name"
	WhitelistIps  = "API key allowed IPs"
	WhitelistType = `Type of entry.
On of: cidrBlock|ipAddress`
	Service = `Cloud service type.
On of: cloud|cloud-manager|ops-manager`
	Provider = `Provider name.
One of: AWS|AZURE|GCP.`
	Region = `Physical location of your MongoDB cluster.
For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws
For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure
For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp`
)
