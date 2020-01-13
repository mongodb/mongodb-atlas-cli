// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package usage

const (
	ProjectID     = "The project ID to use. Overrides config/env settings."
	OrgID         = "The organization ID to use. Overrides config/env settings."
	Profile       = "Use a specific profile from your configuration file."
	Members       = "Number of replica set members."
	InstanceSize  = "Tier for each data-bearing server in the cluster."
	DiskSize      = "Capacity, in gigabytes, of the host’s root volume."
	Backup        = "If true, the cluster uses Atlas Continuous Backups for backing up cluster data."
	MDBVersion    = "Version of the cluster to deploy."
	Page          = "Page number."
	Limit         = "Items per page."
	Username      = "Username for authenticating to MongoDB."
	Password      = "The user’s password."
	Roles         = "User’s roles and the databases/collections on which the roles apply."
	Comment       = "Optional comment associated with the whitelist entry."
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
