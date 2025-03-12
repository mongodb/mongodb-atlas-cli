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

var metadata = map[string]metadatatypes.Metadata{
	`acceptVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Example: ``,
				Usage:   `The VPC Peering Connection id.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`acknowledgeAlert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addAllTeamsToProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addOrganizationRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addProjectApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addProjectRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addProjectServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addProjectUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addTeamUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addUserToProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`addUserToTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team to add the MongoDB Cloud user to.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`authorizeCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`cancelBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreJobId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job to remove.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createApiKeyAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createAtlasSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Label that identifies the cluster to create Search Nodes for.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createAtlasSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection on which to create an Atlas Search index.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createAtlasSearchIndexDeprecated`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection on which to create an Atlas Search index.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createBackupExportJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2024-08-05`: {{
				Name:        `Cluster`,
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
			},
			},
			`2024-10-23`: {{
				Name:        `Cluster`,
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
			},
			},
		},
	},
	`createCostExplorerQueryProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createCustomDatabaseRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createCustomZoneMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createDataFederationPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `AWS IAM Authentication`,
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
			}, {
				Name:        `LDAP Authentication`,
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
			}, {
				Name:        `OIDC Workforce Federated Authentication`,
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
			}, {
				Name:        `OIDC Workload Federated Authentication`,
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
			}, {
				Name:        `SCRAM-SHA Authentication`,
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
			}, {
				Name:        `X509 Authentication`,
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
			},
			},
		},
	},
	`createDatabaseUserCertificate`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: ``,
				Usage:   `Human-readable label that represents the MongoDB database user account for whom to create a certificate.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createEncryptionAtRestPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cloud provider for the private endpoint to create.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `AWS`,
				Description: `AWS`,
				Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592"
}`,
			},
			},
			`2024-05-30`: {{
				Name:        `AWS`,
				Description: `AWS`,
				Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592"
}`,
			}, {
				Name:        `Azure`,
				Description: `Azure`,
				Value: `{
  "cloudProvider": "AZURE",
  "roleId": "668c5f0ed436263134491592",
  "serviceUrl": "https://examplestorageaccount.blob.core.windows.net/examplecontainer"
}`,
			},
			},
		},
	},
	`createFederatedDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`skipRoleValidation`: {
				Example: ``,
				Usage:   `Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createFlexBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster whose snapshot you want to restore.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createLegacyBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createLinkToken`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createManagedNamespace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createOneDataFederationQueryLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createOrganization`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createOrganizationInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createOrganizationUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPeeringContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint.`,
			},
			`endpointServiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPrivateEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`projectOwnerId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user to whom to grant the Project Owner role on the specified project. If you set this parameter, it overrides the default value of the oldest Organization Owner.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
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
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProjectInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProjectIpAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProjectServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProjectServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createProjectServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPushBasedLogConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createPushMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createRollingIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster on which MongoDB Cloud creates an index.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `2dspere Index`,
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
			}, {
				Name:        `Partial Index`,
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
			}, {
				Name:        `Sparse Index`,
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
			},
			},
		},
	},
	`createServerlessBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance whose snapshot you want to restore.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createServerlessPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createSharedClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createStreamInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createStreamInstanceWithSampleConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createThirdPartyIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`integrationType`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`createUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`cutoverMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`liveMigrationId`: {
				Example: `6296fb4c7c7aa997cf94e9a8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the migration.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deauthorizeCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cloud provider of the role to deauthorize.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deferMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAllBackupSchedules`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAllCustomZoneMappings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteApiKeyAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Example: `192.0.2.0%2F24`,
				Usage:   `One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAtlasSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Label that identifies the cluster to delete.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAtlasSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAtlasSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`collectionName`: {
				Example: ``,
				Usage:   `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Example: ``,
				Usage:   `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Example: ``,
				Usage:   `Name of the Atlas Search index to delete.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteAtlasSearchIndexDeprecated`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the database and collection with one or more Application Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`retainBackups`: {
				Example: ``,
				Usage:   `Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteCustomDatabaseRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteDataFederationPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Example: ``,
				Usage:   `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: `SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName`,
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
		RequestBodyExamples: nil,
	},
	`deleteExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportBucketId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal character string that identifies the Export Bucket.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteFederatedDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance to remove.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteFederationApp`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the identity provider to connect.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteLdapConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteLegacySnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteLinkToken`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteManagedNamespace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`collection`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the collection associated with the managed namespace.`,
			},
			`db`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the database that contains the collection.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteOneDataFederationInstanceQueryLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the online archive to delete.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteOrganization`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteOrganizationInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePeeringContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePipelineRunDataset`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pipelineRunId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint.`,
			},
			`endpointId`: {
				Example: ``,
				Usage:   `Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`,
			},
			`endpointServiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePrivateEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint service.`,
			},
			`endpointServiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionId`: {
				Example: ``,
				Usage:   `Unique ID that identifies the Private Link connection.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `Delete a project`,
				Description: `Deletes an existing project`,
			},
			},
		},
	},
	`deleteProjectInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProjectIpAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Example: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
				Usage: `Access list entry that you want to remove from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (` + "`" + `/` + "`" + `) with its URL-encoded value (` + "`" + `%2F` + "`" + `). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:

- how your application established the connection,
- how MongoDB Cloud or the driver using the address behaves, and
- which protocol (like TCP or UDP) the connection uses.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProjectLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProjectServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProjectServiceAccountAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`ipAddress`: {
				Example: `192.0.2.0%2F24`,
				Usage:   `One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteProjectServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`secretId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the secret.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deletePushBasedLogConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteReplicaSetBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteServerlessPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteServiceAccountAccessListEntry`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Example: `192.0.2.0%2F24`,
				Usage:   `One IP address or multiple IP addresses represented as one CIDR block. When specifying a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteServiceAccountSecret`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`secretId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the secret.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteShardedClusterBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream connection.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteStreamInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance to delete.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processorName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team that you want to delete.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteThirdPartyIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`integrationType`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`deleteVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Example: ``,
				Usage:   `The VPC Peering Connection id.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`disableCustomerManagedX509`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`disableDataProtectionSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`disablePeering`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`disableSlowOperationThresholding`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadFederatedDatabaseQueryLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Example: ``,
				Usage:   `Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Example: ``,
				Usage:   `Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance for which you want to download query logs.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadFlexBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadInvoiceCsv`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadOnlineArchiveQueryLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveOnly`: {
				Example: ``,
				Usage:   `Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.`,
			},
			`endDate`: {
				Example: ``,
				Usage:   `Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Example: ``,
				Usage:   `Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadSharedClusterBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`downloadStreamTenantAuditLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Example: ``,
				Usage:   `Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`startDate`: {
				Example: ``,
				Usage:   `Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`enableSlowOperationThresholding`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`endOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that is undergoing outage simulation.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAccountDetails`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `One of "aws", "azure" or "gcp".`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`regionName`: {
				Example: ``,
				Usage:   `The cloud provider specific region name, i.e. "US_EAST_1" for cloud provider "aws".`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getActiveVpcPeeringConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAlert`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getApiKeyAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`ipAddress`: {
				Example: `192.0.2.0%2F24`,
				Usage:   `One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getApiVersions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`env`: {
				Example: ``,
				Usage:   `The environment to get the versions from. If not provided, it returnsthe versions for the given MongoDB URL. (E.g. prod for cloud.mongodb.com)`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasResourcePolicies`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Label that identifies the cluster to return the Search Nodes for.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Example: ``,
				Usage:   `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Example: ``,
				Usage:   `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Example: ``,
				Usage:   `Name of the Atlas Search index to return.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAtlasSearchIndexDeprecated`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAuditingConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getAwsCustomDns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getBackupExportJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal character string that identifies the Export Job.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the restore jobs you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreJobId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getBackupSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCloudProviderAccessRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getClusterAdvancedConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getClusterStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCollStatsLatencyNamespaceClusterMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to retrieve metrics for.`,
			},
			`clusterView`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster topology to retrieve metrics for.`,
			},
			`collectionName`: {
				Example: `mycoll`,
				Usage:   `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Example: `mydb`,
				Usage:   `Human-readable label that identifies the database.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Example: ``,
				Usage:   `List that contains the metrics that you want to retrieve for the associated data series. If you don't set this parameter, this resource returns data series for all Coll Stats Latency metrics.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCollStatsLatencyNamespaceHostMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Example: `mycoll`,
				Usage:   `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Example: `mydb`,
				Usage:   `Human-readable label that identifies the database.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Example: ``,
				Usage:   `List that contains the metrics that you want to retrieve for the associated data series. If you don't set this parameter, this resource returns data series for all Coll Stats Latency metrics.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCollStatsLatencyNamespaceMetrics`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCollStatsLatencyNamespacesForCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`clusterView`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster topology to retrieve metrics for.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCollStatsLatencyNamespacesForHost`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCostExplorerQueryProcess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`token`: {
				Example: `4ABBE973862346D40F3AE859D4BE96E0F895764EB14EAB039E7B82F9D638C05C`,
				Usage:   `Unique 64 digit string that identifies the Cost Explorer query.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getCustomDatabaseRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getDataFederationPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getDataProtectionSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the database that the specified MongoDB process serves.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getDatabaseMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the database that the specified MongoDB process serves.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Example: ``,
				Usage:   `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Example: ``,
				Usage:   `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: `SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName`,
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
		RequestBodyExamples: nil,
	},
	`getDiskMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Example: ``,
				Usage:   `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`partitionName`: {
				Example: ``,
				Usage:   `Human-readable label of the disk or partition to which the measurements apply.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getEncryptionAtRest`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getEncryptionAtRestPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cloud provider of the private endpoint.`,
			},
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getEncryptionAtRestPrivateEndpointsForCloudProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cloud provider for the private endpoints to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getExportBucket`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`exportBucketId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal character string that identifies the Export Bucket.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getFederatedDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Federated Database to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getFederationSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getFlexBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`snapshotId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getFlexBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreJobId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getHostLogs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endDate`: {
				Example: ``,
				Usage:   `Specifies the date and time for the ending point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours after the start date. If the start date is also unspecified, the value will default to the time of the request.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`hostName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the host that stores the log files that you want to download.`,
			},
			`logName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project.`,
			},
			`startDate`: {
				Example: ``,
				Usage:   `Specifies the date and time for the starting point of the range of log messages to retrieve, in the number of seconds that have elapsed since the UNIX epoch. This value will default to 24 hours prior to the end date. If the end date is also unspecified, the value will default to 24 hours prior to the time of the request.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getHostMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`m`: {
				Example: ``,
				Usage:   `One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for ` + "`" + `m` + "`" + `, repeat the ` + "`" + `m` + "`" + ` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getIdentityProviderMetadata`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Example: `c2777a9eca931f29fc2f`,
				Usage:   `Legacy 20-hexadecimal digit string that identifies the identity provider. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getIndexMetrics`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Example: `mycoll`,
				Usage:   `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Example: `mydb`,
				Usage:   `Human-readable label that identifies the database.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Example: `myindex`,
				Usage:   `Human-readable label that identifies the index.`,
			},
			`metrics`: {
				Example: ``,
				Usage:   `List that contains the measurements that MongoDB Atlas reports for the associated data series.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getInvoice`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLdapConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLdapConfigurationStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`requestId`: {
				Example: ``,
				Usage:   `Unique string that identifies the request to verify an <abbr title="Lightweight Directory Access Protocol">LDAP</abbr> configuration.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLegacyBackupCheckpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`checkpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the checkpoint.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLegacyBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`jobId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLegacySnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getLegacySnapshotSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getManagedNamespace`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getManagedSlowMs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Example: ``,
				Usage:   `List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don't set this parameter, this resource returns all hardware and status metrics for the associated data series.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the online archive to return.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOpenApiInfo`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOrganization`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOrganizationEvent`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.`,
			},
			`includeRaw`: {
				Example: ``,
				Usage:   `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOrganizationInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOrganizationSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOrganizationUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that is undergoing outage simulation.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPeeringContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPinnedNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to retrieve pinned namespaces for.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPipelineRun`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pipelineRunId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint.`,
			},
			`endpointId`: {
				Example: ``,
				Usage:   `Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`,
			},
			`endpointServiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPrivateEndpointService`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint service.`,
			},
			`endpointServiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPrivateLinkConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionId`: {
				Example: ``,
				Usage:   `Unique ID that identifies the Private Link connection.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `Get a project`,
				Description: `Get a project using a project id`,
			},
			},
		},
	},
	`getProjectByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this project.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectEvent`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeRaw`: {
				Example: ``,
				Usage:   `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectIpAccessListStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Example: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
				Usage:   `Network address or cloud provider security construct that identifies which project access list entry to be verified.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectIpList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`entryValue`: {
				Example: `IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8`,
				Usage:   `Access list entry that you want to return from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (` + "`" + `/` + "`" + `) with its URL-encoded value (` + "`" + `%2F` + "`" + `).`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectLtsVersions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Filter results to only one cloud provider.`,
			},
			`defaultStatus`: {
				Example: ``,
				Usage:   `Filter results to only the default values per tier. This value must be DEFAULT.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceSize`: {
				Example: `M10`,
				Usage:   `Filter results to only one instance size.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getProjectUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPushBasedLogConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getPushMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`liveMigrationId`: {
				Example: `6296fb4c7c7aa997cf94e9a8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the migration.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getRegionalizedPrivateEndpointSetting`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getReplicaSetBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getResourcesNonCompliant`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getSampleDatasetLoadStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`sampleDatasetId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the loaded sample dataset.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServerlessAutoIndexing`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServerlessBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`snapshotId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServerlessBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreJobId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServerlessPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the tenant endpoint.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getShardedClusterBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getSharedClusterBackup`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getSharedClusterBackupRestoreJob`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`restoreId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the restore job to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream connection to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getStreamInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeConnections`: {
				Example: ``,
				Usage:   `Flag to indicate whether connections information should be included in the stream instance.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processorName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getSystemStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getTeamById`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team whose information you want to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getTeamByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamName`: {
				Example: ``,
				Usage:   `Name of the team whose information you want to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getThirdPartyIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`integrationType`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this user.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getUserByUsername`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userName`: {
				Example: ``,
				Usage:   `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getValidationStatus`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`validationId`: {
				Example: `507f1f77bcf86cd799439011`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the validation job.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`getVpcPeeringConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`requesterAccountId`: {
				Example: ``,
				Usage:   `The Account ID of the VPC Peering connection/s.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`grantMongoDbEmployeeAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAccessLogsByClusterName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`authResult`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the successful authentication attempts only.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`ipAddress`: {
				Example: ``,
				Usage:   `One Internet Protocol address that attempted to authenticate with the database.`,
			},
			`nLogs`: {
				Example: ``,
				Usage:   `Maximum number of lines from the log to return.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAccessLogsByHostname`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`authResult`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the successful authentication attempts only.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses UNIX epoch time in milliseconds.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`hostname`: {
				Example: ``,
				Usage:   `Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.`,
			},
			`ipAddress`: {
				Example: ``,
				Usage:   `One Internet Protocol address that attempted to authenticate with the database.`,
			},
			`nLogs`: {
				Example: ``,
				Usage:   `Maximum number of lines from the log to return.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses UNIX epoch time in milliseconds.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAlertConfigurationMatchersFieldNames`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAlertConfigurations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAlertConfigurationsByAlertId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert. Use the [/alerts](#tag/Alerts/operation/listAlerts) endpoint to retrieve all alerts to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAlerts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`status`: {
				Example: ``,
				Usage:   `Status of the alerts to return. Omit to return all alerts in all statuses.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAlertsByAlertConfigurationId`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listApiKeyAccessListsEntries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAtlasProcesses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAtlasSearchIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Example: ``,
				Usage:   `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Example: ``,
				Usage:   `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAtlasSearchIndexesCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listAtlasSearchIndexesDeprecated`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection with one or more Atlas Search indexes.`,
			},
			`collectionName`: {
				Example: ``,
				Usage:   `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listBackupExportJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the restore jobs you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listCloudProviderAccessRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listCloudProviderRegions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`providers`: {
				Example: ``,
				Usage:   `Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters.`,
			},
			`tier`: {
				Example: ``,
				Usage:   `Cluster tier for which to retrieve the regions.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listClusterSuggestedIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`namespaces`: {
				Example: ``,
				Usage:   `Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`processIds`: {
				Example: ``,
				Usage:   `ProcessIds from which to retrieve suggested indexes. A processId is a combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests. To include multiple processIds, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each processId.`,
			},
			`since`: {
				Example: ``,
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **until** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`until`: {
				Example: ``,
				Usage: `Date and time up until which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you specify the **until** parameter, you must specify the **since** parameter.
- If you specify neither the **since** nor the **until** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listClusters`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeDeletedWithRetainedBackups`: {
				Example: ``,
				Usage:   `Flag that indicates whether to return Clusters with retain backups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listClustersForAllProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listConnectedOrgConfigs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listCustomDatabaseRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDataFederationPrivateEndpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDatabaseUserCertificates`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: ``,
				Usage:   `Human-readable label that represents the MongoDB database user account whose certificates you want to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDatabaseUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDatabases`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDiskMeasurements`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`partitionName`: {
				Example: ``,
				Usage:   `Human-readable label of the disk or partition to which the measurements apply.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDiskPartitions`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: `mongodb.example.com:27017`,
				Usage:   `Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listDropIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listEventTypes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listExportBuckets`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listFederatedDatabases`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`type`: {
				Example: ``,
				Usage:   `Type of Federated Database Instances to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listFlexBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listFlexBackups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listFlexClusters`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listIdentityProviders`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`idpType`: {
				Example: ``,
				Usage:   `The types of the target identity providers.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`protocol`: {
				Example: ``,
				Usage:   `The protocols of the target identity providers.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listIndexMetrics`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`collectionName`: {
				Example: `mycoll`,
				Usage:   `Human-readable label that identifies the collection.`,
			},
			`databaseName`: {
				Example: `mydb`,
				Usage:   `Human-readable label that identifies the database.`,
			},
			`end`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`granularity`: {
				Example: `PT1M`,
				Usage:   `Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`metrics`: {
				Example: ``,
				Usage:   `List that contains the measurements that MongoDB Atlas reports for the associated data series.`,
			},
			`period`: {
				Example: `PT10H`,
				Usage:   `Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`start`: {
				Example: ``,
				Usage:   `Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listInvoices`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`fromDate`: {
				Example: `2023-01-01`,
				Usage:   `Retrieve the invoices the startDates of which are greater than or equal to the fromDate. If omit, the invoices return will go back to earliest startDate.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orderBy`: {
				Example: ``,
				Usage:   `Field used to order the returned invoices by. Use in combination of sortBy parameter to control the order of the result.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`sortBy`: {
				Example: ``,
				Usage:   `Field used to sort the returned invoices by. Use in combination with orderBy parameter to control the order of the result.`,
			},
			`statusNames`: {
				Example: ``,
				Usage:   `Statuses of the invoice to be retrieved. Omit to return invoices of all statuses.`,
			},
			`toDate`: {
				Example: `2023-01-01`,
				Usage:   `Retrieve the invoices the endDates of which are smaller than or equal to the toDate. If omit, the invoices return will go further to latest endDate.`,
			},
			`viewLinkedInvoices`: {
				Example: ``,
				Usage:   `Flag that indicates whether to return linked invoices in the linkedInvoices field.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listLegacyBackupCheckpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listLegacyBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`batchId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can't be part of a batch.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listLegacySnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`completed`: {
				Example: ``,
				Usage:   `Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listMetricTypes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`processId`: {
				Example: `my.host.name.com:27017`,
				Usage:   `Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOnlineArchives`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizationEvents`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventType`: {
				Example: ``,
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeRaw`: {
				Example: ``,
				Usage:   `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`maxDate`: {
				Example: ``,
				Usage:   `Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Example: ``,
				Usage:   `Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizationInvitations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: ``,
				Usage:   `Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizationProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizationTeams`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizationUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listOrganizations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPeeringConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`providerName`: {
				Example: ``,
				Usage:   `Cloud service provider to use for this VPC peering connection.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPeeringContainerByCloudProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`providerName`: {
				Example: ``,
				Usage:   `Cloud service provider that serves the desired network peering containers.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPeeringContainers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPendingInvoices`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPipelineRuns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`createdBefore`: {
				Example: `2022-01-01T00:00:00Z`,
				Usage:   `If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPipelineSchedules`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPipelineSnapshots`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`completedAfter`: {
				Example: `2022-01-01T00:00:00Z`,
				Usage:   `Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPipelines`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPrivateEndpointServices`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Cloud service provider that manages this private endpoint service.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listPrivateLinkConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectApiKeys`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectEvents`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterNames`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`eventType`: {
				Example: ``,
				Usage: `Category of incident recorded at this moment in time.

**IMPORTANT**: The complete list of event type values changes frequently.`,
			},
			`excludedEventType`: {
				Example: ``,
				Usage: `Category of event that you would like to exclude from query results, such as CLUSTER_CREATED

**IMPORTANT**: Event type names change frequently. Verify that you specify the event type correctly by checking the complete list of event types.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeRaw`: {
				Example: ``,
				Usage:   `Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`maxDate`: {
				Example: ``,
				Usage:   `Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`minDate`: {
				Example: ``,
				Usage:   `Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectInvitations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: ``,
				Usage:   `Email address of the user account invited to this project.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectIpAccessLists`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectServiceAccounts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectTeams`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjectUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`flattenTeams`: {
				Example: ``,
				Usage:   `Flag that indicates whether the returned list should include users who belong to a team with a role in this project. You might not have assigned the individual users a role in this project. If ` + "`" + `"flattenTeams" : false` + "`" + `, this resource returns only users with a role in the project.  If ` + "`" + `"flattenTeams" : true` + "`" + `, this resource returns both users with roles in the project and users who belong to teams with roles in the project.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`includeOrgUsers`: {
				Example: ``,
				Usage:   `Flag that indicates whether the returned list should include users with implicit access to the project, the Organization Owner or Organization Read Only role. You might not have assigned the individual users a role in this project. If ` + "`" + `"includeOrgUsers": false` + "`" + `, this resource returns only users with a role in the project. If ` + "`" + `"includeOrgUsers": true` + "`" + `, this resource returns both users with roles in the project and users who have implicit access to the project through their organization role.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
				Name:        `Get a list of all projects`,
				Description: `Get a list of all projects inside of the organisation`,
			},
			},
		},
	},
	`listReplicaSetBackups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listRoleMappings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSchemaAdvice`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServerlessBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServerlessBackups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServerlessInstances`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServerlessPrivateEndpoints`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServiceAccountAccessList`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServiceAccountProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listServiceAccounts`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listShardedClusterBackups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSharedClusterBackupRestoreJobs`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSharedClusterBackups`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSlowQueries`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Example: ``,
				Usage: `Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`nLogs`: {
				Example: ``,
				Usage:   `Maximum number of lines from the log to return.`,
			},
			`namespaces`: {
				Example: ``,
				Usage:   `Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: ``,
				Usage:   `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Example: ``,
				Usage: `Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSlowQueryNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Example: ``,
				Usage: `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: ``,
				Usage:   `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Example: ``,
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSourceProjects`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listStreamConnections`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listStreamInstances`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listStreamProcessors`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listSuggestedIndexes`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`duration`: {
				Example: ``,
				Usage: `Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.

- If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time.
- If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`nExamples`: {
				Example: ``,
				Usage:   `Maximum number of example queries that benefit from the suggested index.`,
			},
			`nIndexes`: {
				Example: ``,
				Usage:   `Number that indicates the maximum indexes to suggest.`,
			},
			`namespaces`: {
				Example: ``,
				Usage:   `Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as ` + "`" + `.` + "`" + `: ` + "`" + `<database>.<collection>` + "`" + `. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (` + "`" + `&` + "`" + `) between each namespace. Omit this parameter to return results for all namespaces.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processId`: {
				Example: ``,
				Usage:   `Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (` + "`" + `mongod` + "`" + ` or ` + "`" + `mongos` + "`" + `). The port must be the IANA port on which the MongoDB process listens for requests.`,
			},
			`since`: {
				Example: ``,
				Usage: `Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).

- If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time.
- If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listTeamUsers`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`listThirdPartyIntegrations`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`loadSampleDataset`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster into which you load the sample dataset.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`migrateProjectToAnotherOrg`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`modifyStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processorName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`pausePipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`pinFeatureCompatibilityVersion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`pinNamespacesPatch`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`pinNamespacesPut`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to pin namespaces to.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`queryLineItemsFromSingleInvoice`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invoiceId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`rejectVpcPeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`id`: {
				Example: ``,
				Usage:   `The VPC Peering Connection id.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeOrganizationRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeOrganizationUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the user to be deleted.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeProjectApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeProjectRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the project. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Project resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeProjectTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeProjectUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal string that identifies MongoDB Cloud user you want to remove from the specified project. To return a application user's ID using their application username, use the Get All application users in One Project endpoint.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeTeamUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`removeUserFromTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team to remove the MongoDB user from.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`renameTeam`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team that you want to rename.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`requestEncryptionAtRestPrivateEndpointDeletion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`cloudProvider`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cloud provider of the private endpoint to delete.`,
			},
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the private endpoint to delete.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`resetMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`resumePipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`returnAllControlPlaneIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`returnAllIpAddresses`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`returnFederatedDatabaseQueryLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance to which the query limit applies.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`returnFederatedDatabaseQueryLimits`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance for which you want to retrieve query limits.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`revokeJwksFromIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the identity provider to connect.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`revokeMongoDbEmployeeAccess`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`saveLdapConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`setProjectLimit`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`limitName`: {
				Example: ``,
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
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`setServerlessAutoIndexing`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`enable`: {
				Example: ``,
				Usage:   `Value that we want to set for the Serverless Auto Indexing toggle.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`startOutageSimulation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to undergo an outage simulation.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`startStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processorName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`stopStreamProcessor`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`processorName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream processor.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`takeSnapshot`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`testFailover`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`toggleAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`toggleAwsCustomDns`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`toggleMaintenanceAutoDefer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`toggleRegionalizedPrivateEndpointSetting`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`triggerSnapshotIngestion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`unpinFeatureCompatibilityVersion`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies this cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`unpinNamespaces`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster to unpin namespaces from.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAlertConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`alertConfigId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the alert configuration. Use the [/alertConfigs](#tag/Alert-Configurations/operation/listAlertConfigurations) endpoint to retrieve all alert configurations to which the authenticated user has access.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateApiKey`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateApiKeyRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`apiUserId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`resourcePolicyId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies an atlas resource policy.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAtlasSearchDeployment`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Label that identifies the cluster to update the Search Nodes for.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAtlasSearchIndex`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection whose Atlas Search index you want to update.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAtlasSearchIndexByName`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection whose Atlas Search index you want to update.`,
			},
			`collectionName`: {
				Example: ``,
				Usage:   `Name of the collection that contains one or more Atlas Search indexes.`,
			},
			`databaseName`: {
				Example: ``,
				Usage:   `Label that identifies the database that contains the collection with one or more Atlas Search indexes.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexName`: {
				Example: ``,
				Usage:   `Name of the Atlas Search index to update.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAtlasSearchIndexDeprecated`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Name of the cluster that contains the collection whose Atlas Search index to update.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`indexId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateAuditingConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateBackupSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateClusterAdvancedConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateConnectedOrgConfig`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`orgId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateCustomDatabaseRole`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`roleName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateDataProtectionSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`overwriteBackupPolicies`: {
				Example: ``,
				Usage:   `Flag that indicates whether to overwrite non complying backup policies with the new data protection settings or not.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateDatabaseUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`databaseName`: {
				Example: ``,
				Usage:   `The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be ` + "`" + `$external` + "`" + `. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be ` + "`" + `admin` + "`" + `.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`username`: {
				Example: `SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName`,
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
		RequestBodyExamples: nil,
	},
	`updateEncryptionAtRest`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateFederatedDatabase`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`skipRoleValidation`: {
				Example: ``,
				Usage:   `Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the federated database instance to update.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the flex cluster.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateIdentityProvider`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`identityProviderId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique string that identifies the identity provider to connect. If using an API version before 11-15-2023, use the legacy 20-hexadecimal digit id. This id can be found within the Federation Management Console > Identity Providers tab by clicking the info icon in the IdP ID row of a configured identity provider. For all other versions, use the 24-hexadecimal digit id.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateLegacySnapshotRetention`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateLegacySnapshotSchedule`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster with the snapshot you want to return.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateMaintenanceWindow`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOnlineArchive`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`archiveId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the online archive to update.`,
			},
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganization`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganizationInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganizationInvitationById`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganizationRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the user to modify.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganizationSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateOrganizationUser`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the pending or active user in the organization. If you need to lookup a user's userId or verify a user's status in the organization, use the Return All MongoDB Cloud Users in One Organization resource and filter by username.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updatePeeringConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`peerId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updatePeeringContainer`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`containerId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updatePipeline`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pipelineName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the Data Lake Pipeline.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: map[string][]metadatatypes.RequestBodyExample{
			`2023-01-01`: {{
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
			},
			},
		},
	},
	`updateProjectInvitation`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateProjectInvitationById`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`invitationId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the invitation.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateProjectRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`userId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the user to modify.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateProjectServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateProjectSettings`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updatePushBasedLogConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateRoleMapping`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`federationSettingsId`: {
				Example: `55fa922fb343282757d9554e`,
				Usage:   `Unique 24-hexadecimal digit string that identifies your federation.`,
			},
			`id`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateServerlessInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`name`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateServerlessPrivateEndpoint`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`endpointId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`instanceName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateServiceAccount`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clientId`: {
				Example: `mdb_sa_id_1234567890abcdef12345678`,
				Usage:   `The Client ID of the Service Account.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateSnapshotRetention`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`clusterName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the cluster.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`snapshotId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the desired snapshot.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateStreamConnection`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`connectionName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream connection.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateStreamInstance`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`tenantName`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the stream instance to update.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateTeamRoles`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
			`teamId`: {
				Example: ``,
				Usage:   `Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`updateThirdPartyIntegration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`includeCount`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`,
			},
			`integrationType`: {
				Example: ``,
				Usage:   `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`,
			},
			`itemsPerPage`: {
				Example: ``,
				Usage:   `Number of items that the response returns per page.`,
			},
			`pageNum`: {
				Example: ``,
				Usage:   `Number of the page that displays the current set of the total objects that the response returns.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`upgradeFlexCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`upgradeSharedCluster`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`upgradeSharedClusterToServerless`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`validateAtlasResourcePolicy`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`orgId`: {
				Example: `4888442a3354817a7320eb61`,
				Usage:   `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`validateMigration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`verifyConnectViaPeeringOnlyModeForOneProject`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`verifyLdapConfiguration`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
			`groupId`: {
				Example: `32b6e34b3d91647abb20e7b8`,
				Usage: `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
			},
			`pretty`: {
				Example: ``,
				Usage:   `Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			},
		},
		RequestBodyExamples: nil,
	},
	`versionedExample`: {
		Parameters: map[string]metadatatypes.ParameterMetadata{
			`additionalInfo`: {
				Example: ``,
				Usage:   `Show more info.`,
			},
			`envelope`: {
				Example: ``,
				Usage:   `Flag that indicates whether Application wraps the response in an ` + "`" + `envelope` + "`" + ` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
			},
		},
		RequestBodyExamples: nil,
	},
}
