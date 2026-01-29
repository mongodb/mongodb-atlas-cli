// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated using `make gen-api-commands`. DO NOT EDIT.
// Don't make any manual changes to this file.

package main

import "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"

var metadata = metadatatypes.Metadata{
	`acceptGroupStreamVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `The VPC Peering Connection id.`,
			},
		},
		Examples: nil,
	},
	`acknowledgeGroupAlert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`addGroupAccessUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`addGroupApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`addGroupTeams`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`addGroupUserRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		Examples: nil,
	},
	`addGroupUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`addOrgTeamUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team to add the MongoDB Cloud user to.`,
			},
		},
		Examples: nil,
	},
	`addOrgTeamUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.`,
			},
		},
		Examples: nil,
	},
	`addOrgUserRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		Examples: nil,
	},
	`authorizeGroupCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		Examples: nil,
	},
	`autoGroupClusterScalingConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`cancelGroupClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`restoreJobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to remove.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`restoreJobId`: `[restoreJobId]`,
				},
			},
			},
		},
	},
	`createFederationSettingConnectedOrgConfigRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: nil,
	},
	`createFederationSettingIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
		},
		Examples: nil,
	},
	`createGroup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`projectOwnerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user to whom to grant the Project Owner role on the specified project. If you set this parameter, it overrides the default value of the oldest Organization Owner.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `create_project`,
				Name:        `Create a basic project with environment tag`,
				Description: `Creates a new project named "MongoTube" with an environment tag set to "e2e"`,
				Value: `{
  "name": "MongoTube",
  "orgId": "67b715468c10250b968dcb84",
  "tags": [
    {
      "key": "environment",
      "value": "e2e"
    }
  ]
}`,
			},
			},
		},
	},
	`createGroupAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_add`,
				Name:        `Add Entries to Project IP Access List`,
				Description: `Adds multiple access list entries to the specified project`,
				Value: `[
  {
    "cidrBlock": "192.168.1.0/24",
    "comment": "Internal network range"
  },
  {
    "cidrBlock": "10.0.0.0/16",
    "comment": "VPC network range"
  }
]`,
				Flags: map[string]string{
					`groupId`: `[your-project-id]`,
				},
			},
			},
		},
	},
	`createGroupAiModelApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupAlertConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupBackupExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `AWS`,

				Description: `AWS`,
				Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592",
  "requirePrivateNetworking": false
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
			`2024-05-30`: {{
				Source: `AWS`,

				Description: `AWS`,
				Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592",
  "requirePrivateNetworking": false
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `Azure`,

				Description: `Azure`,
				Value: `{
  "bucketName": "examplecontainer",
  "cloudProvider": "AZURE",
  "roleId": "668c5f0ed436263134491592",
  "serviceUrl": "https://examplestorageaccount.blob.core.windows.net/examplecontainer"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `GCP`,

				Description: `GCP`,
				Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "GCP",
  "roleId": "668c5f0ed436263134491592"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`createGroupBackupPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider for the private endpoint to create.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupCloudProviderAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`Use-Effective-Instance-Fields`: {
				Usage: `Controls how hardware specification fields are returned in the response after cluster creation. When set to true, returns the original client-specified values and provides separate effective fields showing current operational values. When false (default), hardware specification fields show current operational values directly. Primarily used for autoscaling compatibility.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `Cluster`,

				Description: `Cluster`,
				Value: `{
  "clusterType": "SHARDED",
  "name": "myCluster",
  "replicationSpecs": [
    {
      "regionConfigs": [
        {
          "analyticsAutoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": false
            },
            "diskGB": {
              "enabled": true
            }
          },
          "analyticsSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 0
          },
          "autoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": false
            },
            "diskGB": {
              "enabled": true
            }
          },
          "electableSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M50",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M50",
            "nodeCount": 0
          },
          "regionName": "US_EAST_1"
        }
      ],
      "zoneName": "Zone 1"
    },
    {
      "regionConfigs": [
        {
          "analyticsAutoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": false
            },
            "diskGB": {
              "enabled": true
            }
          },
          "analyticsSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M30",
            "nodeCount": 0
          },
          "autoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": false
            },
            "diskGB": {
              "enabled": true
            }
          },
          "electableSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 0
          },
          "regionName": "US_EAST_1"
        }
      ],
      "zoneName": "Zone 1"
    }
  ]
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
			`2024-10-23`: {{
				Source: `Cluster`,

				Description: `Cluster`,
				Value: `{
  "clusterType": "SHARDED",
  "name": "myCluster",
  "replicationSpecs": [
    {
      "regionConfigs": [
        {
          "analyticsAutoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": true,
              "maxInstanceSize": "M40",
              "minInstanceSize": "M30",
              "scaleDownEnabled": true
            },
            "diskGB": {
              "enabled": true
            }
          },
          "analyticsSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 0
          },
          "autoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": true,
              "maxInstanceSize": "M60",
              "minInstanceSize": "M30",
              "scaleDownEnabled": true
            },
            "diskGB": {
              "enabled": true
            }
          },
          "electableSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M60",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M60",
            "nodeCount": 0
          },
          "regionName": "US_EAST_1"
        }
      ],
      "zoneName": "Zone 1"
    },
    {
      "regionConfigs": [
        {
          "analyticsAutoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": true,
              "maxInstanceSize": "M40",
              "minInstanceSize": "M30",
              "scaleDownEnabled": true
            },
            "diskGB": {
              "enabled": true
            }
          },
          "analyticsSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M30",
            "nodeCount": 0
          },
          "autoScaling": {
            "autoIndexing": {
              "enabled": false
            },
            "compute": {
              "enabled": true,
              "maxInstanceSize": "M60",
              "minInstanceSize": "M30",
              "scaleDownEnabled": true
            },
            "diskGB": {
              "enabled": true
            }
          },
          "electableSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M40",
            "nodeCount": 0
          },
          "regionName": "US_EAST_1"
        }
      ],
      "zoneName": "Zone 1"
    }
  ]
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source:      `create_cluster`,
				Name:        `Create a basic MongoDB Atlas cluster`,
				Description: `Creates a new M10 replica set cluster in AWS US East region running MongoDB 6.0`,
				Value: `{
  "clusterType": "REPLICASET",
  "name": "MyCluster",
  "replicationSpecs": [
    {
      "regionConfigs": [
        {
          "electableSpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M10",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "regionName": "US_EAST_1"
        }
      ]
    }
  ]
}`,
				Flags: map[string]string{
					`groupId`: `[your-project-id]`,
				},
			},
			},
		},
	},
	`createGroupClusterBackupExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterBackupTenantRestore`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterFtsIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection on which to create an Atlas Search index.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterGlobalWriteCustomZoneMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterGlobalWriteManagedNamespace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterIndexRollingIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster on which MongoDB Cloud creates an index.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `2dspere Index`,

				Description: `2dspere Index`,
				Value: `{
  "collation": {
    "alternate": "non-ignorable",
    "backwards": false,
    "caseFirst": "lower",
    "caseLevel": false,
    "locale": "af",
    "maxVariable": "punct",
    "normalization": false,
    "numericOrdering": false,
    "strength": 3
  },
  "collection": "accounts",
  "db": "sample_airbnb",
  "keys": [
    {
      "property_type": "1"
    }
  ],
  "options": {
    "name": "PartialIndexTest",
    "partialFilterExpression": {
      "limit": {
        "$gt": 900
      }
    }
  }
}`,
				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `Partial Index`,

				Description: `Partial Index`,
				Value: `{
  "collation": {
    "alternate": "non-ignorable",
    "backwards": false,
    "caseFirst": "lower",
    "caseLevel": false,
    "locale": "af",
    "maxVariable": "punct",
    "normalization": false,
    "numericOrdering": false,
    "strength": 3
  },
  "collection": "accounts",
  "db": "sample_airbnb",
  "keys": [
    {
      "property_type": "1"
    }
  ],
  "options": {
    "name": "PartialIndexTest",
    "partialFilterExpression": {
      "limit": {
        "$gt": 900
      }
    }
  }
}`,
				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `Sparse Index`,

				Description: `Sparse Index`,
				Value: `{
  "collation": {
    "alternate": "non-ignorable",
    "backwards": false,
    "caseFirst": "lower",
    "caseLevel": false,
    "locale": "af",
    "maxVariable": "punct",
    "normalization": false,
    "numericOrdering": false,
    "strength": 3
  },
  "collection": "accounts",
  "db": "sample_airbnb",
  "keys": [
    {
      "test_field": "1"
    }
  ],
  "options": {
    "name": "SparseIndexTest",
    "sparse": true
  }
}`,
				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`createGroupClusterOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Label that identifies the cluster to create Search Nodes for.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupClusterSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection on which to create an Atlas Search index.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupCustomDbRoleRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupDataFederation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`skipRoleValidation`: {
				Usage: `Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.`,
			},
		},
		Examples: nil,
	},
	`createGroupDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `AWS IAM Authentication`,

				Description: `AWS IAM Authentication`,
				Value: `{
  "awsIAMType": "USER",
  "databaseName": "$external",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "arn:aws:iam::358363220050:user/mongodb-aws-iam-auth-test-user"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `LDAP Authentication`,

				Description: `LDAP Authentication`,
				Value: `{
  "databaseName": "admin",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "ldapAuthType": "GROUP",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "CN=marketing,OU=groups,DC=example,DC=com"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `OIDC Workforce Federated Authentication`,

				Description: `OIDC Workforce Federated Authentication`,
				Value: `{
  "databaseName": "admin",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "oidcAuthType": "IDP_GROUP",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "5dd7496c7a3e5a648454341c/sales"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `OIDC Workload Federated Authentication`,

				Description: `OIDC Workload Federated Authentication`,
				Value: `{
  "databaseName": "$external",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "oidcAuthType": "USER",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "5dd7496c7a3e5a648454341c/sales"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `SCRAM-SHA Authentication`,

				Description: `SCRAM-SHA Authentication`,
				Value: `{
  "databaseName": "admin",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "password": "changeme123",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "david"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			}, {
				Source: `X509 Authentication`,

				Description: `X509 Authentication`,
				Value: `{
  "databaseName": "$external",
  "groupId": "32b6e34b3d91647abb20e7b8",
  "roles": [
    {
      "databaseName": "sales",
      "roleName": "readWrite"
    },
    {
      "databaseName": "marketing",
      "roleName": "read"
    }
  ],
  "scopes": [
    {
      "name": "myCluster",
      "type": "CLUSTER"
    }
  ],
  "username": "CN=david@example.com,OU=users,DC=example,DC=com",
  "x509Type": "CUSTOMER"
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`createGroupDatabaseUserCert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the MongoDB database user account for whom to create a certificate.`,
			},
		},
		Examples: nil,
	},
	`createGroupEncryptionAtRestPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider for the private endpoint to create.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupFlexClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster whose snapshot you want to restore.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`integrationType`: {
				Usage: `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupLiveMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupLogIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupPeer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`createGroupPrivateEndpointEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupPrivateEndpointEndpointServiceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint.`,
			},
			`endpointServiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupPrivateEndpointServerlessInstanceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Usage: `Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.`,
			},
		},
		Examples: nil,
	},
	`createGroupPrivateNetworkSettingEndpointId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupPushBasedLogExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupServerlessBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the serverless instance whose snapshot you want to restore.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupStandbyLink`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether to wrap the response in an envelope.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project containing the clusters.`,
			},
		},
		Examples: nil,
	},
	`createGroupStandbyLinkFailover`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether to wrap the response in an envelope.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`standbyLinkId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the standby link.`,
			},
		},
		Examples: nil,
	},
	`createGroupStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: nil,
	},
	`createGroupStreamPrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: nil,
	},
	`createGroupStreamWorkspace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrg`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgApiKeyAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgBillingCostExplorerUsageProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: nil,
	},
	`createOrgInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgLiveMigrationLinkToken`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgSandboxConfig`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createOrgUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`createUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`cutoverGroupLiveMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`liveMigrationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the migration.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`liveMigrationId`: `6296fb4c7c7aa997cf94e9a8`,
				},
			},
			},
		},
	},
	`deauthorizeGroupCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider of the role to deauthorize.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`roleId`:        `[roleId]`,
				},
			},
			},
		},
	},
	`deferGroupMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteFederationSetting`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
				},
			},
			},
		},
	},
	`deleteFederationSettingConnectedOrgConfigRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`id`:                   `32b6e34b3d91647abb20e7b8`,
					`orgId`:                `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteFederationSettingIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the identity provider to connect.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`identityProviderId`:   `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `delete_project`,
				Name:        `Delete a project`,
				Description: `Deletes an existing project`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Usage: `Access list entry that you want to remove from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (` + "`" + `/` + "`" + `) with its URL-encoded value (` + "`" + `%2F` + "`" + `). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:

- how your application established the connection,
- how MongoDB Cloud or the driver using the address behaves, and
- which protocol (like TCP or UDP) the connection uses.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_delete`,
				Name:        `Remove One Entry from One Project IP Access List`,
				Description: `Removes one access list entry from the specified project's IP access list`,

				Flags: map[string]string{
					`entryValue`: `10.0.0.0/16`,
					`groupId`:    `[your-project-id]`,
				},
			},
			},
		},
	},
	`deleteGroupAiModelApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiKeyId`: {
				Usage: `The id of the API key to be deleted.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiKeyId`: `[apiKeyId]`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupAlertConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupBackupExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportBucketId`: {
				Usage: `Unique 24-hexadecimal character string that identifies the Export Bucket.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`exportBucketId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupBackupPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider of the private endpoint to delete.`,
			},
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint to delete.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`retainBackups`: {
				Usage: `Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source:      `delete_cluster`,
				Name:        `Delete a cluster`,
				Description: `Deletes the specified cluster from the project`,

				Flags: map[string]string{
					`clusterName`: `[your-cluster-name]`,
					`groupId`:     `[your-project-id]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterBackupSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupClusterBackupSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterBackupSnapshotShardedCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterFtsIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterGlobalWriteCustomZoneMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupClusterGlobalWriteManagedNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`collection`: {
				Usage: `Human-readable label that identifies the collection associated with the managed namespace.`,
			},
			`db`: {
				Usage: `Human-readable label that identifies the database that contains the collection.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupClusterOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the online archive to delete.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`archiveId`:   `[archiveId]`,
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupClusterSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Label that identifies the cluster to delete.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupClusterSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`collectionName`: {
				Usage: `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Usage: `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Usage: `Name of the Atlas Search index to delete.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`indexName`:      `[indexName]`,
				},
			},
			},
		},
	},
	`deleteGroupClusterSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`deleteGroupContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`containerId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupCustomDbRoleRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleName`: {
				Usage: `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`roleName`: `[roleName]`,
				},
			},
			},
		},
	},
	`deleteGroupDataFederation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to remove.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`deleteGroupDataFederationLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this data federation instance limit.

| Limit Name | Description | Default |
| --- | --- | --- |
| bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A |
| bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A |
| bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A |
| bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`limitName`:  `[limitName]`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`deleteGroupDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Usage: `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsIAMType | ROLE | <abbr title="Amazon Resource Name">ARN</abbr> |
| AWS IAM | awsIAMType | USER | <abbr title="Amazon Resource Name">ARN</abbr> |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP group name |
| OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP user name |
| SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`databaseName`: `[databaseName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`username`:     `SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName`,
				},
			},
			},
		},
	},
	`deleteGroupFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`deleteGroupIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`integrationType`: {
				Usage: `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`integrationType`: `[integrationType]`,
				},
			},
			},
		},
	},
	`deleteGroupInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`invitationId`: `[invitationId]`,
				},
			},
			},
		},
	},
	`deleteGroupLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this project limit.

| Limit Name | Description | Default | API Override Limit |
| --- | --- | --- | --- |
| atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`limitName`: `[limitName]`,
				},
			},
			},
		},
	},
	`deleteGroupLogIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `Unique identifier of the log integration configuration.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`id`:      `[id]`,
				},
			},
			},
		},
	},
	`deleteGroupPeer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`peerId`:  `[peerId]`,
				},
			},
			},
		},
	},
	`deleteGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`deleteGroupPipelineRun`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pipelineRunId`: {
				Usage: `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pipelineName`:  `[pipelineName]`,
					`pipelineRunId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupPrivateEndpointEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint service.`,
			},
			`endpointServiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupPrivateEndpointEndpointServiceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint.`,
			},
			`endpointId`: {
				Usage: `Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`,
			},
			`endpointServiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointId`:        `[endpointId]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupPrivateEndpointServerlessInstanceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Usage: `Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`:   `[endpointId]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`instanceName`: `[instanceName]`,
				},
			},
			},
		},
	},
	`deleteGroupPrivateNetworkSettingEndpointId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Usage: `Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`: `[endpointId]`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupPushBasedLogExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`deleteGroupServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupServiceAccountAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`ipAddress`: {
				Usage: `One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:  `mdb_sa_id_1234567890abcdef12345678`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`ipAddress`: `192.0.2.0%2F24`,
				},
			},
			},
		},
	},
	`deleteGroupServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`secretId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the secret.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`secretId`: `[secretId]`,
				},
			},
			},
		},
	},
	`deleteGroupStandbyLink`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether to wrap the response in an envelope.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`standbyLinkId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the standby link.`,
			},
		},
		Examples: nil,
	},
	`deleteGroupStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Label that identifies the stream connection.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`connectionName`: `[connectionName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`tenantName`:     `[tenantName]`,
				},
			},
			},
		},
	},
	`deleteGroupStreamPrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionId`: {
				Usage: `Unique ID that identifies the Private Link connection.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`connectionId`: `[connectionId]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`processorName`: `[processorName]`,
					`tenantName`:    `[tenantName]`,
				},
			},
			},
		},
	},
	`deleteGroupStreamVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `The VPC Peering Connection id.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`id`:      `[id]`,
				},
			},
			},
		},
	},
	`deleteGroupStreamWorkspace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace to delete.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`deleteGroupUserSecurityLdapUserToDnMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteOrg`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgApiKeyAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Usage: `One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`invitationId`: `[invitationId]`,
					`orgId`:        `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgLiveMigrationLinkTokens`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:            `4888442a3354817a7320eb61`,
					`resourcePolicyId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`deleteOrgSandboxConfig`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`sandboxConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the sandbox configuration.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:           `4888442a3354817a7320eb61`,
					`sandboxConfigId`: `507f1f77bcf86cd799439011`,
				},
			},
			},
		},
	},
	`deleteOrgServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`orgId`:    `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgServiceAccountAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Usage: `One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:  `mdb_sa_id_1234567890abcdef12345678`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`deleteOrgServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`secretId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the secret.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`secretId`: `[secretId]`,
				},
			},
			},
		},
	},
	`deleteOrgTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team that you want to delete.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:  `4888442a3354817a7320eb61`,
					`teamId`: `[teamId]`,
				},
			},
			},
		},
	},
	`disableGroupBackupCompliancePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`disableGroupManagedSlowMs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`disableGroupPrivateIpModePeering`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`disableGroupUserSecurityCustomerX509`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`downloadGroupClusterBackupTenant`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`downloadGroupClusterLog`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Usage: `Specifies the date and time for the ending point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours after the start date. If the start date is also unspecified, the value will default to the time of the request.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`hostName`: {
				Usage: `Human-readable label that identifies the host that stores the log files that you want to download.`,
			},
			`logName`: {
				Usage: `Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project.`,
			},
			`startDate`: {
				Usage: `Specifies the date and time for the starting point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours prior to the end date. If the end date is also unspecified, the value will default to 24 hours prior to the time of the request.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source:      `get_host_logs`,
				Name:        `Download MongoDB logs for a specific cluster host`,
				Description: `Returns a compressed (.gz) MongoDB log file containing log messages for the specified host`,

				Flags: map[string]string{
					`endDate`:   `1609545600`,
					`groupId`:   `[your-project-id]`,
					`hostName`:  `[your-host-name]`,
					`logName`:   `mongodb`,
					`startDate`: `1609459200`,
				},
			},
			},
		},
	},
	`downloadGroupClusterOnlineArchiveQueryLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveOnly`: {
				Usage: `Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.`,
			},
			`endDate`: {
				Usage: `Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Usage: `Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`endDate`:     `1.636481348e+09`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`startDate`:   `1.636481348e+09`,
				},
			},
			},
		},
	},
	`downloadGroupDataFederationQueryLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Usage: `Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Usage: `Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance for which you want to download query logs.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endDate`:    `1.636481348e+09`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`startDate`:  `1.636466948e+09`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`downloadGroupFlexClusterBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`downloadGroupStreamAuditLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Usage: `Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Usage: `Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endDate`:    `1.636481348e+09`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`startDate`:  `1.636466948e+09`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`enableGroupManagedSlowMs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`endGroupClusterOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that is undergoing outage simulation.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`generateGroupSandboxClusterDescription`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`sandboxConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the sandbox configuration.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`sandboxConfigId`: `507f1f77bcf86cd799439011`,
				},
			},
			},
		},
	},
	`getFederationSettingConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`orgId`:                `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getFederationSettingConnectedOrgConfigRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`id`:                   `32b6e34b3d91647abb20e7b8`,
					`orgId`:                `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getFederationSettingIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Usage: `Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`identityProviderId`:   `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getFederationSettingIdentityProviderMetadata`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Usage: `Legacy 20-hexadecimal digit string that identifies the identity provider. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`identityProviderId`:   `c2777a9eca931f29fc2f`,
				},
			},
			},
		},
	},
	`getGroup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `get_project`,
				Name:        `Get a project`,
				Description: `Get a project using a project id`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Usage: `Access list entry that you want to return from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (` + "`" + `/` + "`" + `) with its URL-encoded value (` + "`" + `%2F` + "`" + `).`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_get`,
				Name:        `Return One Project IP Access List Entry`,
				Description: `Returns one access list entry from the specified project's IP access list: 10.0.0.0/16`,

				Flags: map[string]string{
					`entryValue`: `10.0.0.0/16`,
					`groupId`:    `[your-project-id]`,
				},
			},
			},
		},
	},
	`getGroupAccessListStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Usage: `Network address or cloud provider security construct that identifies which project access list entry to be verified.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_get_status`,
				Name:        `Return Status of One Project IP Access List Entry`,
				Description: `Returns the status of 10.0.0.0/16`,

				Flags: map[string]string{
					`entryValue`: `10.0.0.0/16`,
					`groupId`:    `[your-project-id]`,
				},
			},
			},
		},
	},
	`getGroupActivityFeed`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`eventType`: {
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`maxDate`: {
				Usage: `End date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Usage: `Start date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAiModelApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiKeyId`: {
				Usage: `The id of the API key to be retrieved.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiKeyId`: `[apiKeyId]`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAiModelRateLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`modelGroupName`: {
				Usage: `The name of the model group to be retrieved.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`modelGroupName`: `[modelGroupName]`,
				},
			},
			},
		},
	},
	`getGroupAlert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertId`: `[alertId]`,
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAlertAlertConfigs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertId`: `[alertId]`,
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAlertConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAlertConfigAlerts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAuditLog`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupAwsCustomDns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupBackupCompliancePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-10-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupBackupExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportBucketId`: {
				Usage: `Unique 24-hexadecimal character string that identifies the Export Bucket.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`exportBucketId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupBackupPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider of the private endpoint.`,
			},
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupName`: {
				Usage: `Human-readable label that identifies this project.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`getGroupCloudProviderAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`roleId`:  `[roleId]`,
				},
			},
			},
		},
	},
	`getGroupCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`Use-Effective-Instance-Fields`: {
				Usage: `Controls how hardware specification fields are returned in the response. When set to true, returns the original client-specified values and provides separate effective fields showing current operational values. When false (default), hardware specification fields show current operational values directly. Primarily used for autoscaling compatibility.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source:      `get_cluster`,
				Name:        `Get details of a specific cluster`,
				Description: `Returns the details for the specified cluster in the project`,

				Flags: map[string]string{
					`clusterName`: `[your-cluster-name]`,
					`groupId`:     `[your-project-id]`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupCheckpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`checkpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the checkpoint.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`checkpointId`: `[checkpointId]`,
					`clusterName`:  `[clusterName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportId`: {
				Usage: `Unique 24-hexadecimal character string that identifies the Export Job.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`exportId`:    `32b6e34b3d91647abb20e7b8`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the restore jobs you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`restoreJobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`restoreJobId`: `[restoreJobId]`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupSnapshotShardedCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupTenantRestore`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`restoreId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`restoreId`:   `[restoreId]`,
				},
			},
			},
		},
	},
	`getGroupClusterBackupTenantSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`getGroupClusterCollStatNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`clusterView`: {
				Usage: `Human-readable label that identifies the cluster topology to retrieve metrics for.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`clusterView`: `[clusterView]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`period`:      `PT10H`,
				},
			},
			},
		},
	},
	`getGroupClusterFtsIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
				},
			},
			},
		},
	},
	`getGroupClusterGlobalWrites`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the online archive to return.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`archiveId`:   `[archiveId]`,
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that is undergoing outage simulation.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterProcessArgs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterQueryShape`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`queryShapeHash`: {
				Usage: `A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`queryShapeHash`: `[queryShapeHash]`,
				},
			},
			},
		},
	},
	`getGroupClusterQueryShapeInsightDetails`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processIds`: {
				Usage: `ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each processId.`,
			},
			`queryShapeHash`: {
				Usage: `A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs.`,
			},
			`since`: {
				Usage: `Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`until`: {
				Usage: `Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you specify the **until** parameter, you must specify the **since** parameter.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`queryShapeHash`: `[queryShapeHash]`,
				},
			},
			},
		},
	},
	`getGroupClusterRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`jobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`jobId`:       `[jobId]`,
				},
			},
			},
		},
	},
	`getGroupClusterSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Label that identifies the cluster to return the Search Nodes for.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
				},
			},
			},
		},
	},
	`getGroupClusterSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Usage: `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Usage: `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Usage: `Name of the Atlas Search index to return.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`indexName`:      `[indexName]`,
				},
			},
			},
		},
	},
	`getGroupClusterSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `[snapshotId]`,
				},
			},
			},
		},
	},
	`getGroupClusterSnapshotSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupClusterStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`containerId`: `32b6e34b3d91647abb20e7b8`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupCustomDbRoleRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleName`: {
				Usage: `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`roleName`: `[roleName]`,
				},
			},
			},
		},
	},
	`getGroupDataFederation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the Federated Database to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupDataFederationLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this data federation instance limit.

| Limit Name | Description | Default |
| --- | --- | --- |
| bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A |
| bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A |
| bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A |
| bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`limitName`:  `[limitName]`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Usage: `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsIAMType | ROLE | <abbr title="Amazon Resource Name">ARN</abbr> |
| AWS IAM | awsIAMType | USER | <abbr title="Amazon Resource Name">ARN</abbr> |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP group name |
| OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP user name |
| SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`databaseName`: `[databaseName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`username`:     `SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName`,
				},
			},
			},
		},
	},
	`getGroupDbAccessHistoryCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`authResult`: {
				Usage: `Flag that indicates whether the response returns the successful authentication attempts only.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`end`: {
				Usage: `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`ipAddress`: {
				Usage: `One Internet Protocol address that attempted to authenticate with the database.`,
			},
			`nLogs`: {
				Usage: `Maximum number of lines from the log to return.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupDbAccessHistoryProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`authResult`: {
				Usage: `Flag that indicates whether the response returns the successful authentication attempts only.`,
			},
			`end`: {
				Usage: `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`hostname`: {
				Usage: `Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.`,
			},
			`ipAddress`: {
				Usage: `One Internet Protocol address that attempted to authenticate with the database.`,
			},
			`nLogs`: {
				Usage: `Maximum number of lines from the log to return.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`hostname`: `[hostname]`,
				},
			},
			},
		},
	},
	`getGroupEncryptionAtRest`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupEncryptionAtRestPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider of the private endpoint.`,
			},
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupEvent`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the event that you want to return.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`eventId`: `[eventId]`,
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`getGroupFlexClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`restoreJobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`name`:         `[name]`,
					`restoreJobId`: `[restoreJobId]`,
				},
			},
			},
		},
	},
	`getGroupFlexClusterBackupSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`name`:       `[name]`,
					`snapshotId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupHostFtsMetricIndexMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Usage: `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Usage: `Human-readable label that identifies the database.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Usage: `Human-readable label that identifies the index.`,
			},
			`metrics`: {
				Usage: `List that contains the measurements that MongoDB Atlas reports for the associated data series.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`collectionName`: `mycoll`,
					`databaseName`:   `mydb`,
					`granularity`:    `PT1M`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`indexName`:      `myindex`,
					`metrics`:        `[metrics]`,
					`period`:         `PT10H`,
					`processId`:      `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`getGroupIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`integrationType`: {
				Usage: `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`integrationType`: `[integrationType]`,
				},
			},
			},
		},
	},
	`getGroupInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`invitationId`: `[invitationId]`,
				},
			},
			},
		},
	},
	`getGroupIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this project limit.

| Limit Name | Description | Default | API Override Limit |
| --- | --- | --- | --- |
| atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`limitName`: `[limitName]`,
				},
			},
			},
		},
	},
	`getGroupLiveMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`liveMigrationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the migration.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`liveMigrationId`: `6296fb4c7c7aa997cf94e9a8`,
				},
			},
			},
		},
	},
	`getGroupLiveMigrationValidateStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`validationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the validation job.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`validationId`: `507f1f77bcf86cd799439011`,
				},
			},
			},
		},
	},
	`getGroupLogIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `Unique identifier of the log integration configuration.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`id`:      `[id]`,
				},
			},
			},
		},
	},
	`getGroupMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupManagedSlowMs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupMongoDbVersions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Filter results to only one cloud provider.`,
			},
			`defaultStatus`: {
				Usage: `Filter results to only the default values per tier. This value must be DEFAULT.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceSize`: {
				Usage: `Filter results to only one instance size.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`instanceSize`: `M10`,
					`itemsPerPage`: `100`,
				},
			},
			},
		},
	},
	`getGroupPeer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`peerId`:  `[peerId]`,
				},
			},
			},
		},
	},
	`getGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`getGroupPipelineAvailableSchedules`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`getGroupPipelineAvailableSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`completedAfter`: {
				Usage: `Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`getGroupPipelineRun`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pipelineRunId`: {
				Usage: `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pipelineName`:  `[pipelineName]`,
					`pipelineRunId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupPrivateEndpointEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint service.`,
			},
			`endpointServiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupPrivateEndpointEndpointServiceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint.`,
			},
			`endpointId`: {
				Usage: `Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`,
			},
			`endpointServiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointId`:        `[endpointId]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupPrivateEndpointRegionalMode`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupPrivateEndpointServerlessInstanceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the tenant endpoint.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Usage: `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`:   `[endpointId]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`instanceName`: `[instanceName]`,
				},
			},
			},
		},
	},
	`getGroupPrivateNetworkSettingEndpointId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Usage: `Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`: `[endpointId]`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessCollStatNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`period`:    `PT10H`,
					`processId`: `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Usage: `Human-readable label that identifies the database that the specified MongoDB process serves.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`databaseName`: `[databaseName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`processId`:    `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessDatabaseMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Usage: `Human-readable label that identifies the database that the specified MongoDB process serves.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Usage: `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`databaseName`: `[databaseName]`,
					`granularity`:  `PT1M`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`period`:       `PT10H`,
					`processId`:    `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessDisk`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`partitionName`: {
				Usage: `Human-readable label of the disk or partition to which the measurements apply.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`partitionName`: `[partitionName]`,
					`processId`:     `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessDiskMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Usage: `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`partitionName`: {
				Usage: `Human-readable label of the disk or partition to which the measurements apply.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`granularity`:   `PT1M`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`partitionName`: `[partitionName]`,
					`period`:        `PT10H`,
					`processId`:     `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupProcessMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Usage: `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`granularity`: `PT1M`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`period`:      `PT10H`,
					`processId`:   `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getGroupPushBasedLogExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupSampleDatasetLoad`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`sampleDatasetId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the loaded sample dataset.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`sampleDatasetId`: `[sampleDatasetId]`,
				},
			},
			},
		},
	},
	`getGroupServerlessBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`restoreJobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`restoreJobId`: `[restoreJobId]`,
				},
			},
			},
		},
	},
	`getGroupServerlessBackupSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`snapshotId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`getGroupServerlessPerformanceAdvisorAutoIndexing`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupStandbyLink`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether to wrap the response in an envelope.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`standbyLinkId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the standby link.`,
			},
		},
		Examples: nil,
	},
	`getGroupStandbyLinkFailover`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether to wrap the response in an envelope.`,
			},
			`failoverId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the failover operation.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`standbyLinkId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the standby link.`,
			},
		},
		Examples: nil,
	},
	`getGroupStreamAccountDetails`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `One of "aws", "azure" or "gcp".`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`regionName`: {
				Usage: `The cloud provider specific region name, i.e. "US_EAST_1" for cloud provider "aws".`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`regionName`:    `[regionName]`,
				},
			},
			},
		},
	},
	`getGroupStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Label that identifies the stream connection to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`connectionName`: `[connectionName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`tenantName`:     `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupStreamPrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionId`: {
				Usage: `Unique ID that identifies the Private Link connection.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`connectionId`: `[connectionId]`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`processorName`: `[processorName]`,
					`tenantName`:    `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupStreamProcessors`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupStreamWorkspace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeConnections`: {
				Usage: `Flag to indicate whether connections information should be included in the stream workspace.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`getGroupTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team for which you want to get.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`teamId`:  `[teamId]`,
				},
			},
			},
		},
	},
	`getGroupUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`userId`:  `[userId]`,
				},
			},
			},
		},
	},
	`getGroupUserSecurity`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getGroupUserSecurityLdapVerify`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`requestId`: {
				Usage: `Unique string that identifies the request to verify an Lightweight Directory Access Protocol (LDAP) configuration.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`requestId`: `[requestId]`,
				},
			},
			},
		},
	},
	`getOrg`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgActivityFeed`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`eventType`: {
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`maxDate`: {
				Usage: `End date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Usage: `Start date and time for events to include in the activity feed link. ISO 8601 timestamp format in UTC.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgAiModelApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiKeyId`: {
				Usage: `The id of the API key to be retrieved.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiKeyId`: `[apiKeyId]`,
					`orgId`:    `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgAiModelRateLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`modelGroupName`: {
				Usage: `The name of the model group to be retrieved.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`modelGroupName`: `[modelGroupName]`,
					`orgId`:          `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgApiKeyAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Usage: `One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgBillingCostExplorerUsage`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`token`: {
				Usage: `Unique 64 digit string that identifies the Cost Explorer query.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
					`token`: `4ABBE973862346D40F3AE859D4BE96E0F895764EB14EAB039E7B82F9D638C05C`,
				},
			},
			},
		},
	},
	`getOrgEvent`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the event that you want to return.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`eventId`: `[eventId]`,
					`orgId`:   `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgFederationSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgGroups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`name`: {
				Usage: `Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgInvite`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`invitationId`: `[invitationId]`,
					`orgId`:        `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgInvoice`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`invoiceId`: `[invoiceId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgInvoiceCsv`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`invoiceId`: `[invoiceId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgNonCompliantResources`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:            `4888442a3354817a7320eb61`,
					`resourcePolicyId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`getOrgSandboxConfig`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`sandboxConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the sandbox configuration.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:           `4888442a3354817a7320eb61`,
					`sandboxConfigId`: `507f1f77bcf86cd799439011`,
				},
			},
			},
		},
	},
	`getOrgServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`orgId`:    `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgServiceAccountGroups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`orgId`:    `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`getOrgTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team whose information you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:  `4888442a3354817a7320eb61`,
					`teamId`: `[teamId]`,
				},
			},
			},
		},
	},
	`getOrgTeamByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamName`: {
				Usage: `Name of the team whose information you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:    `4888442a3354817a7320eb61`,
					`teamName`: `[teamName]`,
				},
			},
			},
		},
	},
	`getOrgUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:  `4888442a3354817a7320eb61`,
					`userId`: `[userId]`,
				},
			},
			},
		},
	},
	`getSku`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`skuId`: {
				Usage: `Unique identifier of the SKU to retrieve.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`skuId`: `ATLAS_AWS_INSTANCE_M10`,
				},
			},
			},
		},
	},
	`getSystemStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`getUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this user.`,
			},
		},
		Examples: nil,
	},
	`getUserByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userName`: {
				Usage: `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`,
			},
		},
		Examples: nil,
	},
	`grantGroupClusterMongoDbEmployeeAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`inviteGroupServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listAlertConfigMatcherFieldNames`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listClusterDetails`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listControlPlaneIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		Examples: nil,
	},
	`listEventTypes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listFederationSettingConnectedOrgConfigRoleMappings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`orgId`:                `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listFederationSettingConnectedOrgConfigs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
				},
			},
			},
		},
	},
	`listFederationSettingIdentityProviders`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`idpType`: {
				Usage: `The types of the target identity providers.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`protocol`: {
				Usage: `The protocols of the target identity providers.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
				},
			},
			},
		},
	},
	`listGroupAccessListEntries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_list`,
				Name:        `Return project IP access list`,
				Description: `Returns all access list entries from the specified project's IP access list.`,

				Flags: map[string]string{
					`groupId`: `[your-project-id]`,
				},
			},
			},
		},
	},
	`listGroupAiModelApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupAiModelRateLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupAlertConfigs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupAlerts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`status`: {
				Usage: `Status of the alerts to return. Omit this parameter to return all alerts in all statuses. TRACKING indicates the alert condition exists but has not persisted for the minimum notification delay. OPEN indicates the alert condition currently exists. CLOSED indicates the alert condition has been resolved.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupBackupExportBuckets`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupBackupPrivateEndpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider for the private endpoints to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupCloudProviderAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupCheckpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupExports`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the restore jobs you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupSnapshotShardedClusters`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupTenantRestores`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterBackupTenantSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterCollStatMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to retrieve metrics for.`,
			},
			`clusterView`: {
				Usage: `Human-readable label that identifies the cluster topology to retrieve metrics for.`,
			},
			`collectionName`: {
				Usage: `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Usage: `Human-readable label that identifies the database.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Usage: `List that contains the metrics that you want to retrieve for the associated data series. If you don't set this parameter, this resource returns data series for all Coll Stats Latency metrics.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`clusterView`:    `[clusterView]`,
					`collectionName`: `mycoll`,
					`databaseName`:   `mydb`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`period`:         `PT10H`,
				},
			},
			},
		},
	},
	`listGroupClusterCollStatPinnedNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to retrieve pinned namespaces for.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterFtsIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Usage: `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Usage: `Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterOnlineArchives`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterPerformanceAdvisorDropIndexSuggestions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterPerformanceAdvisorSchemaAdvice`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterPerformanceAdvisorSuggestedIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`namespaces`: {
				Usage: `Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`processIds`: {
				Usage: `ProcessIds from which to retrieve suggested indexes. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each processId.`,
			},
			`since`: {
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`until`: {
				Usage: `Date and time up until which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you specify the **until** parameter, you must specify the **since** parameter.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterProviderRegions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`providers`: {
				Usage: `Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters.`,
			},
			`tier`: {
				Usage: `Cluster tier for which to retrieve the regions.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterQueryShapeInsightSummaries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`commands`: {
				Usage: `Retrieve query shape statistics matching specified MongoDB commands. To include multiple commands, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each command. The currently supported parameters are find, distinct, and aggregate. Omit this parameter to return results for all supported commands.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`nSummaries`: {
				Usage: `Maximum number of query statistic summaries to return.`,
			},
			`namespaces`: {
				Usage: `Namespaces from which to retrieve query shape statistics. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processIds`: {
				Usage: `ProcessIds from which to retrieve query shape statistics. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each processId.`,
			},
			`queryShapeHashes`: {
				Usage: `A list of SHA256 hashes of desired query shapes, output by MongoDB commands like $queryStats and $explain or slow query logs. To include multiple series, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each series. Omit this parameter to return results for all available series.`,
			},
			`series`: {
				Usage: `Query shape statistics data series to retrieve. A series represents a specific metric about query execution. To include multiple series, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each series. Omit this parameter to return results for all available series.`,
			},
			`since`: {
				Usage: `Date and time from which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`until`: {
				Usage: `Date and time up until which to retrieve query shape statistics. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you specify the **until** parameter, you must specify the **since** parameter.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterQueryShapes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`status`: {
				Usage: `The status of query shapes to retrieve. Only REJECTED status is supported. If omitted, defaults to REJECTED.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`batchId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can't be part of a batch.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Usage: `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Usage: `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterSearchIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusterSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`completed`: {
				Usage: `Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupClusters`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`Use-Effective-Instance-Fields`: {
				Usage: `Controls how hardware specification fields are returned in the response. When set to true, returns the original client-specified values and provides separate effective fields showing current operational values. When false (default), hardware specification fields show current operational values directly. Primarily used for autoscaling compatibility.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeDeletedWithRetainedBackups`: {
				Usage: `Flag that indicates whether to return Clusters with retain backups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source:      `list_clusters`,
				Name:        `List all clusters in a project`,
				Description: `Returns all clusters in the specified project`,

				Flags: map[string]string{
					`groupId`: `[your-project-id]`,
				},
			},
			},
		},
	},
	`listGroupCollStatMetrics`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupContainerAll`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupContainers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`providerName`: {
				Usage: `Cloud service provider that serves the desired network peering containers.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`providerName`: `[providerName]`,
				},
			},
			},
		},
	},
	`listGroupCustomDbRoleRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupDataFederation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`type`: {
				Usage: `Type of Federated Database Instances to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupDataFederationLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance for which you want to retrieve query limits.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`listGroupDatabaseUserCerts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the MongoDB database user account whose certificates you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`username`: `[username]`,
				},
			},
			},
		},
	},
	`listGroupDatabaseUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupEncryptionAtRestPrivateEndpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider for the private endpoints to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupEvents`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterNames`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventType`: {
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`excludedEventType`: {
				Usage: `Category of event that you would like to exclude from query results, such as CLUSTER_CREATED

**IMPORTANT**: Event type names change frequently. Verify that you specify the event type correctly by checking the complete list of event types.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`maxDate`: {
				Usage: `Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Usage: `Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupFlexClusterBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`listGroupFlexClusterBackupSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`listGroupFlexClusters`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupHostFtsMetricIndexMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Usage: `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Usage: `Human-readable label that identifies the database.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Usage: `List that contains the measurements that MongoDB Atlas reports for the associated data series.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`collectionName`: `mycoll`,
					`databaseName`:   `mydb`,
					`granularity`:    `PT1M`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`metrics`:        `[metrics]`,
					`period`:         `PT10H`,
					`processId`:      `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`listGroupHostFtsMetricMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Usage: `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Usage: `List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don't set this parameter, this resource returns all hardware and status metrics for the associated data series.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`granularity`: `PT1M`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`metrics`:     `[metrics]`,
					`period`:      `PT10H`,
					`processId`:   `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`listGroupHostFtsMetrics`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`listGroupIntegrations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupInvites`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Email address of the user account invited to this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupLogIntegrations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`integrationType`: {
				Usage: `Optional filter by integration type (e.g., 'S3_LOG_EXPORT').`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-03-12`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupPeers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`providerName`: {
				Usage: `Cloud service provider to use for this VPC peering connection.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupPipelineRuns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`createdBefore`: {
				Usage: `If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`listGroupPipelines`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupPrivateEndpointEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Cloud service provider that manages this private endpoint service.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupPrivateEndpointServerlessInstanceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Usage: `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`instanceName`: `[instanceName]`,
				},
			},
			},
		},
	},
	`listGroupPrivateNetworkSettingEndpointIds`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupProcessCollStatMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Usage: `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Usage: `Human-readable label that identifies the database.`,
			},
			`end`: {
				Usage: `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Usage: `List that contains the metrics that you want to retrieve for the associated data series. If you don't set this parameter, this resource returns data series for all Coll Stats Latency metrics.`,
			},
			`period`: {
				Usage: `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Usage: `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`collectionName`: `mycoll`,
					`databaseName`:   `mydb`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`period`:         `PT10H`,
					`processId`:      `my.host.name.com:27017`,
				},
			},
			},
		},
	},
	`listGroupProcessDatabases`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`listGroupProcessDisks`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`listGroupProcessPerformanceAdvisorNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Usage: `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `[processId]`,
				},
			},
			},
		},
	},
	`listGroupProcessPerformanceAdvisorSlowQueryLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Usage: `Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeMetrics`: {
				Usage: `Whether or not to include metrics extracted from the slow query log as separate fields.`,
			},
			`includeOpType`: {
				Usage: `Whether or not to include the operation type (read/write/command) extracted from the slow query log as a separate field.`,
			},
			`includeReplicaState`: {
				Usage: `Whether or not to include the replica state of the host when the slow query log was generated as a separate field.`,
			},
			`nLogs`: {
				Usage: `Maximum number of lines from the log to return.`,
			},
			`namespaces`: {
				Usage: `Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Usage: `Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `[processId]`,
				},
			},
			},
		},
	},
	`listGroupProcessPerformanceAdvisorSuggestedIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Usage: `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`nExamples`: {
				Usage: `Maximum number of example queries that benefit from the suggested index.`,
			},
			`nIndexes`: {
				Usage: `Number that indicates the maximum indexes to suggest.`,
			},
			`namespaces`: {
				Usage: `Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processId`: {
				Usage: `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`processId`: `[processId]`,
				},
			},
			},
		},
	},
	`listGroupProcesses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `list_atlas_processes`,
				Name:        `List all processes for a project`,
				Description: `Returns details of all processes for the specified project`,

				Flags: map[string]string{
					`groupId`: `[your-project-id]`,
				},
			},
			},
		},
	},
	`listGroupServerlessBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupServerlessBackupSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupServerlessInstances`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupServiceAccounts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupStandbyLinkFailovers`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`standbyLinkId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the standby link.`,
			},
		},
		Examples: nil,
	},
	`listGroupStandbyLinks`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the project.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listGroupStreamActiveVpcPeeringConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupStreamConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`tenantName`: `[tenantName]`,
				},
			},
			},
		},
	},
	`listGroupStreamPrivateLinkConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupStreamVpcPeeringConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`requesterAccountId`: {
				Usage: `The Account ID of the VPC Peering connection/s.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:            `32b6e34b3d91647abb20e7b8`,
					`requesterAccountId`: `[requesterAccountId]`,
				},
			},
			},
		},
	},
	`listGroupStreamWorkspaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupTeams`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`listGroupUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`flattenTeams`: {
				Usage: `Flag that indicates whether the returned list should include users who belong to a team with a role in this project. You might not have assigned the individual users a role in this project. If ` + "`" + `"flattenTeams" : false` + "`" + `, this resource returns only users with a role in the project.  If ` + "`" + `"flattenTeams" : true` + "`" + `, this resource returns both users with roles in the project and users who belong to teams with roles in the project.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeOrgUsers`: {
				Usage: `Flag that indicates whether the returned list should include users with implicit access to the project, the Organization Owner or Organization Read Only role. You might not have assigned the individual users a role in this project. If ` + "`" + `"includeOrgUsers": false` + "`" + `, this resource returns only users with a role in the project. If ` + "`" + `"includeOrgUsers": true` + "`" + `, this resource returns both users with roles in the project and users who have implicit access to the project through their organization role.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgMembershipStatus`: {
				Usage: `Flag that indicates whether to filter the returned list by users organization membership status. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Email address to filter users by. Not supported in deprecated versions.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:             `32b6e34b3d91647abb20e7b8`,
					`orgMembershipStatus`: `ACTIVE`,
				},
			},
			},
		},
	},
	`listGroups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `list_projects`,
				Name:        `Get a list of all projects`,
				Description: `Get a list of all projects inside of the organisation`,
			},
			},
		},
	},
	`listOrgAiModelApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgAiModelRateLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgApiKeyAccessListEntries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgEvents`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventType`: {
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`maxDate`: {
				Usage: `Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Usage: `Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgInvites`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgInvoicePending`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgInvoices`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`fromDate`: {
				Usage: `Retrieve the invoices the startDates of which are greater than or equal to the fromDate. If omit, the invoices return will go back to earliest startDate.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orderBy`: {
				Usage: `Field used to order the returned invoices by. Use in combination of sortBy parameter to control the order of the result.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`sortBy`: {
				Usage: `Field used to sort the returned invoices by. Use in combination with orderBy parameter to control the order of the result.`,
			},
			`statusNames`: {
				Usage: `Statuses of the invoice to be retrieved. Omit to return invoices of all statuses.`,
			},
			`toDate`: {
				Usage: `Retrieve the invoices the endDates of which are smaller than or equal to the toDate. If omit, the invoices return will go further to latest endDate.`,
			},
			`viewLinkedInvoices`: {
				Usage: `Flag that indicates whether to return linked invoices in the linkedInvoices field.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`fromDate`: `2023-01-01`,
					`orderBy`:  `desc`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`toDate`:   `2023-01-01`,
				},
			},
			},
		},
	},
	`listOrgLiveMigrationAvailableProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgResourcePolicies`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgSandboxConfig`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`orgId`:    `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgServiceAccounts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgTeamUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`orgMembershipStatus`: {
				Usage: `Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string to filter users by. Not supported in deprecated versions.`,
			},
			`username`: {
				Usage: `Email address to filter users by. Not supported in deprecated versions.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:               `4888442a3354817a7320eb61`,
					`orgMembershipStatus`: `ACTIVE`,
					`teamId`:              `[teamId]`,
				},
			},
			},
		},
	},
	`listOrgTeams`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`: `4888442a3354817a7320eb61`,
				},
			},
			},
		},
	},
	`listOrgUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`orgMembershipStatus`: {
				Usage: `Organization membership status to filter users by. If you exclude this parameter, this resource returns both pending and active users. Not supported in deprecated versions.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Email address to filter users by. Not supported in deprecated versions.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:               `4888442a3354817a7320eb61`,
					`orgMembershipStatus`: `ACTIVE`,
				},
			},
			},
		},
	},
	`listOrgs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`name`: {
				Usage: `Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`listSkus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`migrateGroup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`pauseGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`pinGroupClusterCollStatPinnedNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`pinGroupClusterFeatureCompatibilityVersion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`rejectGroupStreamVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `The VPC Peering Connection id.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`id`:      `[id]`,
				},
			},
			},
		},
	},
	`removeFederationSettingConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`orgId`:                `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`removeGroupApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`removeGroupTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`teamId`:  `[teamId]`,
				},
			},
			},
		},
	},
	`removeGroupUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Project](#tag/MongoDB-Cloud-Users/operation/listProjectUsers) resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`userId`:  `[userId]`,
				},
			},
			},
		},
	},
	`removeGroupUserRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		Examples: nil,
	},
	`removeOrgTeamUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team to remove the MongoDB user from.`,
			},
		},
		Examples: nil,
	},
	`removeOrgTeamUserFromTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:  `4888442a3354817a7320eb61`,
					`teamId`: `[teamId]`,
					`userId`: `[userId]`,
				},
			},
			},
		},
	},
	`removeOrgUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the [Return All MongoDB Cloud Users in One Organization](#tag/MongoDB-Cloud-Users/operation/listOrganizationUsers) resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2025-02-19`: {{
				Source: `-`,

				Flags: map[string]string{
					`orgId`:  `4888442a3354817a7320eb61`,
					`userId`: `[userId]`,
				},
			},
			},
		},
	},
	`removeOrgUserRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		Examples: nil,
	},
	`renameOrgTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team that you want to rename.`,
			},
		},
		Examples: nil,
	},
	`requestGroupEncryptionAtRestPrivateEndpointDeletion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Usage: `Human-readable label that identifies the cloud provider of the private endpoint to delete.`,
			},
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the private endpoint to delete.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`requestGroupSampleDatasetLoad`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the cluster into which you load the sample dataset.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
					`name`:    `[name]`,
				},
			},
			},
		},
	},
	`resetGroupAiModelRateLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`preview`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`resetGroupMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`restartGroupClusterPrimaries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`resumeGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
				},
			},
			},
		},
	},
	`revokeFederationSettingIdentityProviderJwks`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the identity provider to connect.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`identityProviderId`:   `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`revokeGroupClusterMongoDbEmployeeAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`searchOrgInvoiceLineItems`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
		},
		Examples: nil,
	},
	`setGroupDataFederationLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this data federation instance limit.

| Limit Name | Description | Default |
| --- | --- | --- |
| bytesProcessed.query | Limit on the number of bytes processed during a single data federation query | N/A |
| bytesProcessed.daily | Limit on the number of bytes processed for the data federation instance for the current day | N/A |
| bytesProcessed.weekly | Limit on the number of bytes processed for the data federation instance for the current week | N/A |
| bytesProcessed.monthly | Limit on the number of bytes processed for the data federation instance for the current month | N/A |
`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		Examples: nil,
	},
	`setGroupLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Usage: `Human-readable label that identifies this project limit.

| Limit Name | Description | Default | API Override Limit |
| --- | --- | --- | --- |
| atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 100 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 100 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`setGroupServerlessPerformanceAdvisorAutoIndexing`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`enable`: {
				Usage: `Value that we want to set for the Serverless Auto Indexing toggle.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`enable`:      `[enable]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`startGroupClusterOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to undergo an outage simulation.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`startGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`processorName`: `[processorName]`,
					`tenantName`:    `[tenantName]`,
				},
			},
			},
		},
	},
	`startGroupStreamProcessorWith`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: nil,
	},
	`stopGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`processorName`: `[processorName]`,
					`tenantName`:    `[tenantName]`,
				},
			},
			},
		},
	},
	`takeGroupClusterBackupSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`tenantGroupFlexClusterUpgrade`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`toggleGroupAlertConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`toggleGroupAwsCustomDns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`toggleGroupMaintenanceWindowAutoDefer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`toggleGroupPrivateEndpointRegionalMode`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`triggerGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`unpinGroupClusterCollStatUnpinNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to unpin namespaces from.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`unpinGroupClusterFeatureCompatibilityVersion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`updateFederationSettingConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.`,
			},
		},
		Examples: nil,
	},
	`updateFederationSettingConnectedOrgConfigRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		Examples: nil,
	},
	`updateFederationSettingIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Usage: `Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id.`,
			},
		},
		Examples: nil,
	},
	`updateGroup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `update_project`,
				Name:        `Update project name and tags`,
				Description: `Update the value of the existing project to "MongoTube - Production" and change the tags to an environment tag set to "production"`,
				Value: `{
  "name": "MongoTube - Production",
  "tags": [
    {
      "key": "environment",
      "value": "production"
    }
  ]
}`,
				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`updateGroupAiModelApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiKeyId`: {
				Usage: `The id of the API key to be updated.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupAiModelRateLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`modelGroupName`: {
				Usage: `The name of the model group to be updated.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupAlertConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupApiKeyRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupAuditLog`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupBackupCompliancePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`overwriteBackupPolicies`: {
				Usage: `Flag that indicates whether to overwrite non complying backup policies with the new data protection settings or not.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupBackupExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportBucketId`: {
				Usage: `Unique 24-hexadecimal character string that identifies the snapshot export bucket.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`Use-Effective-Instance-Fields`: {
				Usage: `Controls how hardware specification fields are returned in the response after cluster updates. When set to true, returns the original client-specified values and provides separate effective fields showing current operational values. When false (default), hardware specification fields show current operational values directly. Note: When using this header with autoscaling enabled, MongoDB ignores replicationSpecs changes during updates. To intentionally override the replicationSpecs, disable this header.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-10-23`: {{
				Source:      `update_cluster`,
				Name:        `Update a cluster's configuration`,
				Description: `Updates the cluster to increase disk size to 20GB and increase node count to 5`,
				Value: `{
  "replicationSpecs": [
    {
      "regionConfigs": [
        {
          "electableSpecs": {
            "diskSizeGB": 20,
            "instanceSize": "M10",
            "nodeCount": 5
          },
          "priority": 7,
          "providerName": "AWS",
          "regionName": "US_EAST_1"
        }
      ]
    }
  ]
}`,
				Flags: map[string]string{
					`clusterName`: `[your-cluster-name]`,
					`groupId`:     `[your-project-id]`,
				},
			},
			},
		},
	},
	`updateGroupClusterBackupSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterBackupSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterCollStatPinnedNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterFtsIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection whose Atlas Search index to update.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the online archive to update.`,
			},
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterProcessArgs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterQueryShape`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`queryShapeHash`: {
				Usage: `A SHA256 hash of a query shape, output by MongoDB commands like $queryStats and $explain or slow query logs.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Label that identifies the cluster to update the Search Nodes for.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection whose Atlas Search index you want to update.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Name of the cluster that contains the collection whose Atlas Search index you want to update.`,
			},
			`collectionName`: {
				Usage: `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Usage: `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Usage: `Name of the Atlas Search index to update.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`snapshotId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		Examples: nil,
	},
	`updateGroupClusterSnapshotSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Usage: `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupCustomDbRoleRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`roleName`: {
				Usage: `Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.`,
			},
		},
		Examples: nil,
	},
	`updateGroupDataFederation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`skipRoleValidation`: {
				Usage: `Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to update.`,
			},
		},
		Examples: nil,
	},
	`updateGroupDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Usage: `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsIAMType | ROLE | <abbr title="Amazon Resource Name">ARN</abbr> |
| AWS IAM | awsIAMType | USER | <abbr title="Amazon Resource Name">ARN</abbr> |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC Workforce | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP group name |
| OIDC Workload | oidcAuthType | USER | Atlas OIDC IdP ID (found in federation settings), followed by a '/', followed by the IdP user name |
| SCRAM-SHA | awsIAMType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`,
			},
		},
		Examples: nil,
	},
	`updateGroupEncryptionAtRest`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Usage: `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`integrationType`: {
				Usage: `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupInviteById`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
		},
		Examples: nil,
	},
	`updateGroupInvites`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupLogIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Usage: `Unique identifier of the log integration configuration.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		Examples: nil,
	},
	`updateGroupPeer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Usage: `Human-readable label that identifies the Data Lake Pipeline.`,
			},
		},
		Examples: nil,
	},
	`updateGroupPrivateEndpointServerlessInstanceEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Usage: `Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.`,
			},
		},
		Examples: nil,
	},
	`updateGroupPushBasedLogExport`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Usage: `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateGroupStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Label that identifies the stream connection.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: nil,
	},
	`updateGroupStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`processorName`: {
				Usage: `Label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace.`,
			},
		},
		Examples: nil,
	},
	`updateGroupStreamWorkspace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`tenantName`: {
				Usage: `Label that identifies the stream workspace to update.`,
			},
		},
		Examples: nil,
	},
	`updateGroupTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.`,
			},
		},
		Examples: nil,
	},
	`updateGroupUserRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the user to modify.`,
			},
		},
		Examples: nil,
	},
	`updateGroupUserSecurity`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrg`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgInviteById`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgInvites`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		Examples: nil,
	},
	`updateOrgSandboxConfig`: {
		OnlyPrivatePreview: true,
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`sandboxConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the sandbox configuration.`,
			},
		},
		Examples: nil,
	},
	`updateOrgServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Usage: `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`updateOrgUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		Examples: nil,
	},
	`updateOrgUserRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the user to modify.`,
			},
		},
		Examples: nil,
	},
	`upgradeGroupClusterTenantUpgrade`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`upgradeGroupClusterTenantUpgradeToServerless`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`validateGroupLiveMigrations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`validateOrgResourcePolicies`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`verifyGroupPrivateIpMode`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`groupId`: `32b6e34b3d91647abb20e7b8`,
				},
			},
			},
		},
	},
	`verifyGroupUserSecurityLdap`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`withGroupStreamSampleConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
}
