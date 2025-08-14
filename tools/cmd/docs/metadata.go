// Copyright 2025 MongoDB Inc
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
	`acceptVpcPeeringConnection`: {
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
	`acknowledgeAlert`: {
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
	`addAllTeamsToProject`: {
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
	`addOrganizationRole`: {
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
	`addProjectApiKey`: {
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
	`addProjectRole`: {
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
	`addProjectServiceAccount`: {
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
	`addProjectUser`: {
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
	`addTeamUser`: {
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
	`addUserToProject`: {
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
	`addUserToTeam`: {
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
	`authorizeCloudProviderAccessRole`: {
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
	`autoScalingConfiguration`: {
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
	`cancelBackupRestoreJob`: {
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
	`createAlertConfiguration`: {
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
	`createApiKey`: {
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
	`createApiKeyAccessList`: {
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
	`createAtlasSearchDeployment`: {
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
	`createAtlasSearchIndex`: {
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
	`createAtlasSearchIndexDeprecated`: {
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
	`createBackupExportJob`: {
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
	`createBackupRestoreJob`: {
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
	`createCloudProviderAccessRole`: {
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
	`createCluster`: {
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
	`createCostExplorerQueryProcess`: {
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
	`createCustomDatabaseRole`: {
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
	`createCustomZoneMapping`: {
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
	`createDataFederationPrivateEndpoint`: {
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
	`createDatabaseUser`: {
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
	`createDatabaseUserCertificate`: {
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
	`createEncryptionAtRestPrivateEndpoint`: {
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
	`createExportBucket`: {
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
  "iamRoleId": "668c5f0ed436263134491592"
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
  "iamRoleId": "668c5f0ed436263134491592"
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
	`createFederatedDatabase`: {
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
	`createFlexBackupRestoreJob`: {
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
	`createFlexCluster`: {
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
	`createIdentityProvider`: {
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
	`createLegacyBackupRestoreJob`: {
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
	`createLinkToken`: {
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
	`createManagedNamespace`: {
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
	`createOneDataFederationQueryLimit`: {
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
	`createOnlineArchive`: {
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
	`createOrganization`: {
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
	`createOrganizationInvitation`: {
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
	`createOrganizationUser`: {
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
	`createPeeringConnection`: {
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
	`createPeeringContainer`: {
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
	`createPipeline`: {
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
	`createPrivateEndpoint`: {
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
	`createPrivateEndpointService`: {
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
	`createPrivateLinkConnection`: {
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
	`createProject`: {
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
	`createProjectApiKey`: {
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
	`createProjectInvitation`: {
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
	`createProjectIpAccessList`: {
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
	`createProjectServiceAccount`: {
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
	`createProjectServiceAccountAccessList`: {
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
	`createProjectServiceAccountSecret`: {
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
	`createPushBasedLogConfiguration`: {
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
	`createPushMigration`: {
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
	`createRoleMapping`: {
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
	`createRollingIndex`: {
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
	`createServerlessBackupRestoreJob`: {
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
	`createServerlessInstance`: {
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
	`createServerlessPrivateEndpoint`: {
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
	`createServiceAccount`: {
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
	`createServiceAccountAccessList`: {
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
	`createServiceAccountSecret`: {
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
	`createSharedClusterBackupRestoreJob`: {
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
	`createStreamConnection`: {
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
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: nil,
	},
	`createStreamInstance`: {
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
	`createStreamInstanceWithSampleConnections`: {
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
	`createStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: nil,
	},
	`createTeam`: {
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
	`createThirdPartyIntegration`: {
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
	`cutoverMigration`: {
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
	`deauthorizeCloudProviderAccessRole`: {
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
	`deferMaintenanceWindow`: {
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
	`deleteAlertConfiguration`: {
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
	`deleteAllBackupSchedules`: {
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
	`deleteAllCustomZoneMappings`: {
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
	`deleteApiKey`: {
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
	`deleteApiKeyAccessListEntry`: {
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
	`deleteAtlasSearchDeployment`: {
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
	`deleteAtlasSearchIndex`: {
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
	`deleteAtlasSearchIndexByName`: {
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
	`deleteAtlasSearchIndexDeprecated`: {
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
	`deleteCluster`: {
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
	`deleteCustomDatabaseRole`: {
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
	`deleteDataFederationPrivateEndpoint`: {
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
	`deleteDatabaseUser`: {
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
	`deleteExportBucket`: {
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
	`deleteFederatedDatabase`: {
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
	`deleteFederationApp`: {
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
	`deleteFlexCluster`: {
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
	`deleteIdentityProvider`: {
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
	`deleteLdapConfiguration`: {
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
	`deleteLegacySnapshot`: {
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
	`deleteLinkToken`: {
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
	`deleteManagedNamespace`: {
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
	`deleteOneDataFederationInstanceQueryLimit`: {
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
	`deleteOnlineArchive`: {
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
	`deleteOrganization`: {
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
	`deleteOrganizationInvitation`: {
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
	`deletePeeringConnection`: {
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
	`deletePeeringContainer`: {
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
	`deletePipeline`: {
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
	`deletePipelineRunDataset`: {
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
	`deletePrivateEndpoint`: {
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
	`deletePrivateEndpointService`: {
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
	`deletePrivateLinkConnection`: {
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
	`deleteProject`: {
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
	`deleteProjectInvitation`: {
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
	`deleteProjectIpAccessList`: {
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
	`deleteProjectLimit`: {
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
| atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
| atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
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
	`deleteProjectServiceAccount`: {
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
	`deleteProjectServiceAccountAccessListEntry`: {
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
	`deleteProjectServiceAccountSecret`: {
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
	`deletePushBasedLogConfiguration`: {
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
	`deleteReplicaSetBackup`: {
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
	`deleteRoleMapping`: {
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
	`deleteServerlessInstance`: {
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
	`deleteServerlessPrivateEndpoint`: {
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
	`deleteServiceAccount`: {
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
	`deleteServiceAccountAccessListEntry`: {
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
	`deleteServiceAccountSecret`: {
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
	`deleteShardedClusterBackup`: {
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
	`deleteStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Human-readable label that identifies the stream connection.`,
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
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`deleteStreamInstance`: {
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
				Usage: `Human-readable label that identifies the stream instance to delete.`,
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
	`deleteStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`deleteTeam`: {
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
	`deleteThirdPartyIntegration`: {
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
	`deleteVpcPeeringConnection`: {
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
	`disableCustomerManagedX509`: {
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
	`disableDataProtectionSettings`: {
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
	`disablePeering`: {
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
	`disableSlowOperationThresholding`: {
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
	`downloadFederatedDatabaseQueryLogs`: {
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
	`downloadFlexBackup`: {
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
	`downloadInvoiceCsv`: {
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
	`downloadOnlineArchiveQueryLogs`: {
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
	`downloadSharedClusterBackup`: {
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
	`downloadStreamTenantAuditLogs`: {
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
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`enableSlowOperationThresholding`: {
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
	`endOutageSimulation`: {
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
	`getAccountDetails`: {
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
	`getActiveVpcPeeringConnections`: {
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
	`getAlert`: {
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
	`getAlertConfiguration`: {
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
	`getApiKey`: {
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
	`getApiKeyAccessList`: {
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
	`getAtlasProcess`: {
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
	`getAtlasSearchDeployment`: {
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
	`getAtlasSearchIndex`: {
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
	`getAtlasSearchIndexByName`: {
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
	`getAtlasSearchIndexDeprecated`: {
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
	`getAuditingConfiguration`: {
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
	`getAwsCustomDns`: {
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
	`getBackupExportJob`: {
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
	`getBackupRestoreJob`: {
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
	`getBackupSchedule`: {
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
	`getCloudProviderAccessRole`: {
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
	`getCluster`: {
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
	`getClusterAdvancedConfiguration`: {
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
	`getClusterStatus`: {
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
	`getCollStatsLatencyNamespaceClusterMeasurements`: {
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
	`getCollStatsLatencyNamespaceHostMeasurements`: {
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
	`getCollStatsLatencyNamespaceMetrics`: {
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
	`getCollStatsLatencyNamespacesForCluster`: {
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
	`getCollStatsLatencyNamespacesForHost`: {
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
	`getConnectedOrgConfig`: {
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
	`getCostExplorerQueryProcess`: {
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
	`getCustomDatabaseRole`: {
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
	`getDataFederationPrivateEndpoint`: {
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
	`getDataProtectionSettings`: {
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
	`getDatabase`: {
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
	`getDatabaseMeasurements`: {
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
	`getDatabaseUser`: {
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
	`getDiskMeasurements`: {
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
	`getEncryptionAtRest`: {
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
	`getEncryptionAtRestPrivateEndpoint`: {
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
	`getEncryptionAtRestPrivateEndpointsForCloudProvider`: {
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
	`getExportBucket`: {
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
	`getFederatedDatabase`: {
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
	`getFederationSettings`: {
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
	`getFlexBackup`: {
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
	`getFlexBackupRestoreJob`: {
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
	`getFlexCluster`: {
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
	`getGroupClusterQueryShapeInsightSummaries`: {
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
	`getHostLogs`: {
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
	`getHostMeasurements`: {
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
	`getIdentityProvider`: {
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
	`getIdentityProviderMetadata`: {
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
	`getIndexMetrics`: {
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
	`getInvoice`: {
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
	`getLdapConfiguration`: {
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
	`getLdapConfigurationStatus`: {
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
	`getLegacyBackupCheckpoint`: {
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
	`getLegacyBackupRestoreJob`: {
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
	`getLegacySnapshot`: {
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
	`getLegacySnapshotSchedule`: {
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
	`getMaintenanceWindow`: {
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
	`getManagedNamespace`: {
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
	`getManagedSlowMs`: {
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
	`getMeasurements`: {
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
	`getOnlineArchive`: {
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
	`getOrganization`: {
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
	`getOrganizationEvent`: {
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
	`getOrganizationInvitation`: {
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
	`getOrganizationSettings`: {
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
	`getOrganizationUser`: {
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
	`getOutageSimulation`: {
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
	`getPeeringConnection`: {
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
	`getPeeringContainer`: {
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
	`getPinnedNamespaces`: {
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
	`getPipeline`: {
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
	`getPipelineRun`: {
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
	`getPrivateEndpoint`: {
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
	`getPrivateEndpointService`: {
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
	`getPrivateLinkConnection`: {
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
	`getProject`: {
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
	`getProjectByName`: {
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
	`getProjectEvent`: {
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
	`getProjectInvitation`: {
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
	`getProjectIpAccessListStatus`: {
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
	`getProjectIpList`: {
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
	`getProjectLimit`: {
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
| atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
| atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
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
	`getProjectLtsVersions`: {
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
	`getProjectServiceAccount`: {
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
	`getProjectSettings`: {
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
	`getProjectTeam`: {
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
	`getProjectUser`: {
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
	`getPushBasedLogConfiguration`: {
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
	`getPushMigration`: {
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
	`getRegionalizedPrivateEndpointSetting`: {
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
	`getReplicaSetBackup`: {
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
	`getResourcesNonCompliant`: {
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
	`getRoleMapping`: {
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
	`getSampleDatasetLoadStatus`: {
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
	`getServerlessAutoIndexing`: {
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
	`getServerlessBackup`: {
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
	`getServerlessBackupRestoreJob`: {
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
	`getServerlessInstance`: {
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
	`getServerlessPrivateEndpoint`: {
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
	`getServiceAccount`: {
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
	`getShardedClusterBackup`: {
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
	`getSharedClusterBackup`: {
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
	`getSharedClusterBackupRestoreJob`: {
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
	`getStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Human-readable label that identifies the stream connection to return.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance to return.`,
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
	`getStreamInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeConnections`: {
				Usage: `Flag to indicate whether connections information should be included in the stream instance.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance to return.`,
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
	`getStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`getTeamById`: {
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
	`getTeamByName`: {
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
	`getThirdPartyIntegration`: {
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
	`getUserByUsername`: {
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
	`getValidationStatus`: {
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
	`getVpcPeeringConnections`: {
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
	`grantMongoDbEmployeeAccess`: {
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
	`listAccessLogsByClusterName`: {
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
	`listAccessLogsByHostname`: {
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
	`listAlertConfigurationMatchersFieldNames`: {
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
	`listAlertConfigurations`: {
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
	`listAlertConfigurationsByAlertId`: {
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
	`listAlerts`: {
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
	`listAlertsByAlertConfigurationId`: {
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
	`listApiKeyAccessListsEntries`: {
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
	`listApiKeys`: {
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
	`listAtlasProcesses`: {
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
	`listAtlasSearchIndexes`: {
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
	`listAtlasSearchIndexesCluster`: {
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
	`listAtlasSearchIndexesDeprecated`: {
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
	`listBackupExportJobs`: {
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
	`listBackupRestoreJobs`: {
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
	`listCloudProviderAccessRoles`: {
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
	`listCloudProviderRegions`: {
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
	`listClusterSuggestedIndexes`: {
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
	`listClusters`: {
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
	`listClustersForAllProjects`: {
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
	`listConnectedOrgConfigs`: {
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
	`listCustomDatabaseRoles`: {
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
	`listDataFederationPrivateEndpoints`: {
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
	`listDatabaseUserCertificates`: {
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
	`listDatabaseUsers`: {
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
	`listDatabases`: {
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
	`listDiskMeasurements`: {
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
	`listDiskPartitions`: {
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
	`listDropIndexes`: {
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
	`listExportBuckets`: {
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
	`listFederatedDatabases`: {
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
	`listFlexBackupRestoreJobs`: {
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
	`listFlexBackups`: {
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
	`listFlexClusters`: {
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
	`listIdentityProviders`: {
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
	`listIndexMetrics`: {
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
	`listInvoices`: {
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
	`listLegacyBackupCheckpoints`: {
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
	`listLegacyBackupRestoreJobs`: {
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
	`listLegacySnapshots`: {
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
	`listMetricTypes`: {
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
	`listOnlineArchives`: {
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
	`listOrganizationEvents`: {
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
	`listOrganizationInvitations`: {
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
	`listOrganizationProjects`: {
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
	`listOrganizationTeams`: {
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
	`listOrganizationUsers`: {
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
	`listOrganizations`: {
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
	`listPeeringConnections`: {
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
	`listPeeringContainerByCloudProvider`: {
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
	`listPeeringContainers`: {
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
	`listPendingInvoices`: {
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
	`listPipelineRuns`: {
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
	`listPipelineSchedules`: {
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
	`listPipelineSnapshots`: {
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
	`listPipelines`: {
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
	`listPrivateEndpointServices`: {
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
	`listPrivateLinkConnections`: {
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
	`listProjectApiKeys`: {
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
	`listProjectEvents`: {
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
	`listProjectInvitations`: {
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
	`listProjectIpAccessLists`: {
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
	`listProjectLimits`: {
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
	`listProjectServiceAccountAccessList`: {
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
	`listProjectServiceAccounts`: {
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
	`listProjectTeams`: {
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
	`listProjectUsers`: {
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
	`listProjects`: {
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
	`listReplicaSetBackups`: {
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
	`listRoleMappings`: {
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
	`listSchemaAdvice`: {
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
	`listServerlessBackupRestoreJobs`: {
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
	`listServerlessBackups`: {
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
	`listServerlessInstances`: {
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
	`listServerlessPrivateEndpoints`: {
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
	`listServiceAccountAccessList`: {
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
	`listServiceAccountProjects`: {
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
	`listServiceAccounts`: {
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
	`listShardedClusterBackups`: {
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
	`listSharedClusterBackupRestoreJobs`: {
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
	`listSharedClusterBackups`: {
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
	`listSlowQueries`: {
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
	`listSlowQueryNamespaces`: {
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
	`listSourceProjects`: {
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
	`listStreamConnections`: {
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
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`listStreamInstances`: {
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
	`listStreamProcessors`: {
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
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`listSuggestedIndexes`: {
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
	`listTeamUsers`: {
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
	`listThirdPartyIntegrations`: {
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
	`loadSampleDataset`: {
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
	`migrateProjectToAnotherOrg`: {
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
	`modifyStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: nil,
	},
	`pausePipeline`: {
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
	`pinFeatureCompatibilityVersion`: {
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
	`pinNamespacesPatch`: {
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
	`pinNamespacesPut`: {
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
	`queryLineItemsFromSingleInvoice`: {
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
	`rejectVpcPeeringConnection`: {
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
	`removeConnectedOrgConfig`: {
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
	`removeOrganizationRole`: {
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
	`removeOrganizationUser`: {
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
	`removeProjectApiKey`: {
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
	`removeProjectRole`: {
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
	`removeProjectTeam`: {
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
	`removeProjectUser`: {
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
	`removeTeamUser`: {
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
	`removeUserFromTeam`: {
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
	`renameTeam`: {
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
	`requestEncryptionAtRestPrivateEndpointDeletion`: {
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
	`resetMaintenanceWindow`: {
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
	`resumePipeline`: {
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
	`returnAllControlPlaneIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		Examples: nil,
	},
	`returnAllIpAddresses`: {
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
	`returnFederatedDatabaseQueryLimit`: {
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
	`returnFederatedDatabaseQueryLimits`: {
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
	`revokeJwksFromIdentityProvider`: {
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
	`revokeMongoDbEmployeeAccess`: {
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
	`saveLdapConfiguration`: {
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
	`setProjectLimit`: {
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
| atlas.project.deployment.clusters | Limit on the number of clusters in this project if the org is not sales-sold (If org is sales-sold, the limit is 100) | 25 | 90 |
| atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 |
| atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 |
| atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 |
| atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 |
| atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 |
| dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A |
| dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A |
| dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A |
| dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A |
| atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Service Connections per Region Group | 50 | 100|
| atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
| atlas.project.deployment.salesSoldM0s | Limit on the number of M0 clusters in this project if the org is sales-sold | 100 | 100 |
`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the prettyprint format.`,
			},
		},
		Examples: nil,
	},
	`setServerlessAutoIndexing`: {
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
	`startOutageSimulation`: {
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
	`startStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`startStreamProcessorWith`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: nil,
	},
	`stopStreamProcessor`: {
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
				Usage: `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
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
	`takeSnapshot`: {
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
	`testFailover`: {
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
	`toggleAlertConfiguration`: {
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
	`toggleAwsCustomDns`: {
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
	`toggleMaintenanceAutoDefer`: {
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
	`toggleRegionalizedPrivateEndpointSetting`: {
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
	`triggerSnapshotIngestion`: {
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
	`unpinFeatureCompatibilityVersion`: {
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
	`unpinNamespaces`: {
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
	`updateAlertConfiguration`: {
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
	`updateApiKey`: {
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
	`updateApiKeyRoles`: {
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
	`updateAtlasSearchDeployment`: {
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
	`updateAtlasSearchIndex`: {
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
	`updateAtlasSearchIndexByName`: {
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
	`updateAtlasSearchIndexDeprecated`: {
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
	`updateAuditingConfiguration`: {
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
	`updateBackupSchedule`: {
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
	`updateCluster`: {
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
	`updateClusterAdvancedConfiguration`: {
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
	`updateConnectedOrgConfig`: {
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
	`updateCustomDatabaseRole`: {
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
	`updateDataProtectionSettings`: {
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
	`updateDatabaseUser`: {
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
	`updateEncryptionAtRest`: {
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
	`updateFederatedDatabase`: {
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
	`updateFlexCluster`: {
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
	`updateIdentityProvider`: {
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
	`updateLegacySnapshotRetention`: {
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
	`updateLegacySnapshotSchedule`: {
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
	`updateMaintenanceWindow`: {
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
	`updateOnlineArchive`: {
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
	`updateOrganization`: {
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
	`updateOrganizationInvitation`: {
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
	`updateOrganizationInvitationById`: {
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
	`updateOrganizationRoles`: {
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
	`updateOrganizationSettings`: {
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
	`updateOrganizationUser`: {
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
	`updatePeeringConnection`: {
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
	`updatePeeringContainer`: {
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
	`updatePipeline`: {
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
	`updateProject`: {
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
	`updateProjectInvitation`: {
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
	`updateProjectInvitationById`: {
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
	`updateProjectRoles`: {
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
	`updateProjectServiceAccount`: {
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
	`updateProjectSettings`: {
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
	`updatePushBasedLogConfiguration`: {
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
	`updateRoleMapping`: {
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
	`updateServerlessInstance`: {
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
	`updateServerlessPrivateEndpoint`: {
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
	`updateServiceAccount`: {
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
	`updateSnapshotRetention`: {
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
	`updateStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Usage: `Human-readable label that identifies the stream connection.`,
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
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: nil,
	},
	`updateStreamInstance`: {
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
				Usage: `Human-readable label that identifies the stream instance to update.`,
			},
		},
		Examples: nil,
	},
	`updateTeamRoles`: {
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
	`updateThirdPartyIntegration`: {
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
	`upgradeFlexCluster`: {
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
	`upgradeSharedCluster`: {
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
	`upgradeSharedClusterToServerless`: {
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
	`validateAtlasResourcePolicy`: {
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
	`validateMigration`: {
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
	`verifyConnectViaPeeringOnlyModeForOneProject`: {
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
	`verifyLdapConfiguration`: {
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
