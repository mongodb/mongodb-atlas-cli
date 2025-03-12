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
				Usage: `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
	},
	`createAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
				},
			}, {
				Source: `Azure`,

				Description: `Azure`,
				Value: `{
  "cloudProvider": "AZURE",
  "roleId": "668c5f0ed436263134491592",
  "serviceUrl": "https://examplestorageaccount.blob.core.windows.net/examplecontainer"
}`,
				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Flags: map[string]string{
					`envelope`: `false`,
					`pretty`:   `false`,
				},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
		Examples: nil,
	},
	`deleteAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`envelope`:  `false`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`envelope`:  `false`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
				},
			},
			},
		},
	},
	`deleteAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:         `false`,
					`orgId`:            `4888442a3354817a7320eb61`,
					`pretty`:           `false`,
					`resourcePolicyId`: `32b6e34b3d91647abb20e7b8`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`indexName`:      `[indexName]`,
					`pretty`:         `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`retainBackups`: {
				Usage: `Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleName`: {
				Usage: `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`: `[endpointId]`,
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
					`envelope`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to remove.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`name`:     `[name]`,
					`pretty`:   `false`,
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
					`envelope`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`archiveId`:   `[archiveId]`,
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`invitationId`: `[invitationId]`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`peerId`:   `[peerId]`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`containerId`: `32b6e34b3d91647abb20e7b8`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pipelineName`:  `[pipelineName]`,
					`pipelineRunId`: `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointId`:        `[endpointId]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`envelope`:          `false`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
					`pretty`:            `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`envelope`:          `false`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
					`pretty`:            `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`connectionId`: `[connectionId]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `delete_project`,
				Name:        `Delete a project`,
				Description: `Deletes an existing project`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_delete`,
				Name:        `Remove One Entry from One Project IP Access List`,
				Description: `Removes one access list entry from the specified project's IP access list`,

				Flags: map[string]string{
					`entryValue`: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`limitName`: `[limitName]`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:  `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`ipAddress`: `192.0.2.0%2F24`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`name`:     `[name]`,
					`pretty`:   `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:  `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`:  `false`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`pretty`:         `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance to delete.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team that you want to delete.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
					`teamId`:   `[teamId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:        `false`,
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`integrationType`: `[integrationType]`,
					`pretty`:          `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`id`:       `[id]`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`invoiceId`: `[invoiceId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
	},
	`getAlert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertId`:  `[alertId]`,
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
				},
			},
			},
		},
	},
	`getAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`envelope`:  `false`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`envelope`:  `false`,
					`ipAddress`: `192.0.2.0%2F24`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
				},
			},
			},
		},
	},
	`getApiVersions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`env`: {
				Usage: `The environment to get the versions from. If not provided, it returnsthe versions for the given MongoDB URL. (E.g. prod for cloud.mongodb.com)`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Usage: `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`pretty`:    `false`,
					`processId`: `mongodb.example.com:27017`,
				},
			},
			},
		},
	},
	`getAtlasResourcePolicies`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
				},
			},
			},
		},
	},
	`getAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:         `false`,
					`orgId`:            `4888442a3354817a7320eb61`,
					`pretty`:           `false`,
					`resourcePolicyId`: `32b6e34b3d91647abb20e7b8`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`indexName`:      `[indexName]`,
					`pretty`:         `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`indexId`:     `[indexId]`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
					`roleId`:   `[roleId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:       `false`,
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
					`envelope`:       `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
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
					`envelope`:    `false`,
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
					`envelope`:  `false`,
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
					`envelope`:             `false`,
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
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`token`:    `4ABBE973862346D40F3AE859D4BE96E0F895764EB14EAB039E7B82F9D638C05C`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleName`: {
				Usage: `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`endpointId`: `[endpointId]`,
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-10-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`granularity`:  `PT1M`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`period`:       `PT10H`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:      `false`,
					`granularity`:   `PT1M`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`partitionName`: `[partitionName]`,
					`period`:        `PT10H`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`includeCount`:  `true`,
					`itemsPerPage`:  `100`,
					`pageNum`:       `1`,
					`pretty`:        `false`,
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
					`envelope`:       `false`,
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
					`envelope`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
					`envelope`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreJobId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`name`:         `[name]`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`name`:     `[name]`,
					`pretty`:   `false`,
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
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`hostName`: `[hostName]`,
					`logName`:  `[logName]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`granularity`: `PT1M`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`period`:      `PT10H`,
					`pretty`:      `false`,
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
					`envelope`:             `false`,
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
					`envelope`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`invoiceId`: `[invoiceId]`,
					`orgId`:     `4888442a3354817a7320eb61`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`requestId`: {
				Usage: `Unique string that identifies the request to verify an <abbr title="Lightweight Directory Access Protocol">LDAP</abbr> configuration.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`checkpointId`: `[checkpointId]`,
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`jobId`:       `[jobId]`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`archiveId`:   `[archiveId]`,
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
				},
			},
			},
		},
	},
	`getOpenApiInfo`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`pretty`: `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`eventId`:  `[eventId]`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
					`userId`:   `[userId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`peerId`:   `[peerId]`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`containerId`: `32b6e34b3d91647abb20e7b8`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pipelineName`:  `[pipelineName]`,
					`pipelineRunId`: `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointId`:        `[endpointId]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`envelope`:          `false`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
					`pretty`:            `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`:     `[cloudProvider]`,
					`endpointServiceId`: `[endpointServiceId]`,
					`envelope`:          `false`,
					`groupId`:           `32b6e34b3d91647abb20e7b8`,
					`pretty`:            `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `get_project`,
				Name:        `Get a project`,
				Description: `Get a project using a project id`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`groupName`: `[groupName]`,
					`pretty`:    `false`,
				},
			},
			},
		},
	},
	`getProjectEvent`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeRaw`: {
				Usage: `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`eventId`:  `[eventId]`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`invitationId`: `[invitationId]`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_get_status`,
				Name:        `Return Status of One Project IP Access List Entry`,
				Description: `Returns the status of 10.0.0.0/16`,

				Flags: map[string]string{
					`entryValue`: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_get`,
				Name:        `Return One Project IP Access List Entry`,
				Description: `Returns one access list entry from the specified project's IP access list: 10.0.0.0/16`,

				Flags: map[string]string{
					`entryValue`: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`limitName`: `[limitName]`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`instanceSize`: `M10`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
					`userId`:   `[userId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:        `false`,
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`liveMigrationId`: `6296fb4c7c7aa997cf94e9a8`,
					`pretty`:          `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
					`envelope`:             `false`,
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
					`envelope`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`name`:     `[name]`,
					`pretty`:   `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`: `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
					`envelope`:       `false`,
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
					`envelope`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`pretty`:   `false`,
				},
			},
			},
		},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team whose information you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
					`teamId`:   `[teamId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamName`: {
				Usage: `Name of the team whose information you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:        `false`,
					`groupId`:         `32b6e34b3d91647abb20e7b8`,
					`integrationType`: `[integrationType]`,
					`pretty`:          `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies this user.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`pretty`:   `false`,
					`userId`:   `[userId]`,
				},
			},
			},
		},
	},
	`getUserByUsername`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userName`: {
				Usage: `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`pretty`:   `false`,
					`userName`: `[userName]`,
				},
			},
			},
		},
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`requesterAccountId`: {
				Usage: `The Account ID of the VPC Peering connection/s.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:           `false`,
					`groupId`:            `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`:       `100`,
					`pageNum`:            `1`,
					`pretty`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`start`: {
				Usage: `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`hostname`: `[hostname]`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`pretty`:   `false`,
				},
			},
			},
		},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
	},
	`listAlertConfigurationsByAlertId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertId`:      `[alertId]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`status`: {
				Usage: `Status of the alerts to return. Omit to return all alerts in all statuses.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
	},
	`listAlertsByAlertConfigurationId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`alertConfigId`: `32b6e34b3d91647abb20e7b8`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`includeCount`:  `true`,
					`itemsPerPage`:  `100`,
					`pageNum`:       `1`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`:    `[apiUserId]`,
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`pretty`:         `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:    `[clusterName]`,
					`collectionName`: `[collectionName]`,
					`databaseName`:   `[databaseName]`,
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`pretty`:         `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
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
					`envelope`:             `false`,
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`itemsPerPage`:         `100`,
					`pageNum`:              `1`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Usage: `Human-readable label that represents the MongoDB database user account whose certificates you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`username`:     `[username]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`processId`:    `mongodb.example.com:27017`,
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
					`envelope`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Usage: `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`processId`:    `mongodb.example.com:27017`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`type`: {
				Usage: `Type of Federated Database Instances to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`name`:         `[name]`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`name`:         `[name]`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-11-13`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
					`envelope`:             `false`,
					`federationSettingsId`: `55fa922fb343282757d9554e`,
					`itemsPerPage`:         `100`,
					`pageNum`:              `1`,
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
					`envelope`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:           `false`,
					`fromDate`:           `2023-01-01`,
					`includeCount`:       `true`,
					`itemsPerPage`:       `100`,
					`orderBy`:            `desc`,
					`orgId`:              `4888442a3354817a7320eb61`,
					`pageNum`:            `1`,
					`pretty`:             `false`,
					`toDate`:             `2023-01-01`,
					`viewLinkedInvoices`: `true`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
					`envelope`:  `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Usage: `Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
			},
			},
		},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`providerName`: {
				Usage: `Cloud service provider to use for this VPC peering connection.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`providerName`: {
				Usage: `Cloud service provider that serves the desired network peering containers.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`createdBefore`: `2022-01-01T00:00:00Z`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`includeCount`:  `true`,
					`itemsPerPage`:  `100`,
					`pageNum`:       `1`,
					`pipelineName`:  `[pipelineName]`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`pipelineName`: `[pipelineName]`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`completedAfter`: `2022-01-01T00:00:00Z`,
					`envelope`:       `false`,
					`groupId`:        `32b6e34b3d91647abb20e7b8`,
					`includeCount`:   `true`,
					`itemsPerPage`:   `100`,
					`pageNum`:        `1`,
					`pipelineName`:   `[pipelineName]`,
					`pretty`:         `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Usage: `Email address of the user account invited to this project.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `project_ip_access_list_list`,
				Name:        `Return project IP access list`,
				Description: `Returns all access list entries from the specified project's IP access list.`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:     `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source:      `list_projects`,
				Name:        `Get a list of all projects`,
				Description: `Get a list of all projects inside of the organisation`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
				},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
					`envelope`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`:  `[clusterName]`,
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
					`envelope`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:     `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`:     `false`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`clientId`:     `mdb_sa_id_1234567890abcdef12345678`,
					`envelope`:     `false`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-08-05`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`clusterName`: `[clusterName]`,
					`envelope`:    `false`,
					`groupId`:     `32b6e34b3d91647abb20e7b8`,
					`pretty`:      `false`,
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
			`nLogs`: {
				Usage: `Maximum number of lines from the log to return.`,
			},
			`namespaces`: {
				Usage: `Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`tenantName`:   `[tenantName]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the stream instance.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2024-05-30`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`tenantName`:   `[tenantName]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`processId`:    `[processId]`,
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
			`pageNum`: {
				Usage: `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`itemsPerPage`: `100`,
					`orgId`:        `4888442a3354817a7320eb61`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
					`teamId`:       `[teamId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:     `false`,
					`groupId`:      `32b6e34b3d91647abb20e7b8`,
					`includeCount`: `true`,
					`itemsPerPage`: `100`,
					`pageNum`:      `1`,
					`pretty`:       `false`,
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
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
		Examples: nil,
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
					`envelope`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the user to be deleted.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
					`userId`:   `[userId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`apiUserId`: `[apiUserId]`,
					`envelope`:  `false`,
					`groupId`:   `32b6e34b3d91647abb20e7b8`,
					`pretty`:    `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`teamId`:   `[teamId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Usage: `Unique 24-hexadecimal string that identifies MongoDB Cloud user you want to remove from the specified project. To return a application user's ID using their application username, use the Get All application users in One Project endpoint.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2043-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
					`userId`:   `[userId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`orgId`:    `4888442a3354817a7320eb61`,
					`pretty`:   `false`,
					`teamId`:   `[teamId]`,
					`userId`:   `[userId]`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`cloudProvider`: `[cloudProvider]`,
					`endpointId`:    `[endpointId]`,
					`envelope`:      `false`,
					`groupId`:       `32b6e34b3d91647abb20e7b8`,
					`pretty`:        `false`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
	},
	`returnAllControlPlaneIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-11-15`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
				},
			},
			},
		},
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`limitName`:  `[limitName]`,
					`pretty`:     `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Usage: `Human-readable label that identifies the federated database instance for which you want to retrieve query limits.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`:   `false`,
					`groupId`:    `32b6e34b3d91647abb20e7b8`,
					`pretty`:     `false`,
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
					`envelope`:             `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
	},
	`toggleAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
	},
	`updateAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Usage: `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-01-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
					`groupId`:  `32b6e34b3d91647abb20e7b8`,
					`pretty`:   `false`,
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
				Usage: `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		Examples: nil,
	},
	`versionedExample`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`additionalInfo`: {
				Usage: `Show more info.`,
			},
			`envelope`: {
				Usage: `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		Examples: map[string][]metadatatypes.Example{
			`2023-02-01`: {{
				Source: `-`,

				Flags: map[string]string{
					`envelope`: `false`,
				},
			},
			},
		},
	},
}
