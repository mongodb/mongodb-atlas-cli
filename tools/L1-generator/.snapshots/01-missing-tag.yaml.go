// This code was autogenerated using `make gen-L1-commands`
// Don't make any manual changes to this file.
//
//nolint:revive,stylecheck
package L1

var Commands GroupedAndSortedCommands = GroupedAndSortedCommands{
	{
		Name:        `Clusters`,
		Description: ``,
		Commands: []Command{
			{
				OperationID: `createCluster`,
				Description: `Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters.`,
				RequestParameters: RequestParameters{
					URL: `/api/atlas/v2/groups/{groupId}/clusters`,
					QueryParameters: []Parameter{
						{
							Name:        `envelope`,
							Description: `Flag that indicates whether Application wraps the response in an envelope JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
							Required:    false,
						},
						{
							Name:        `pretty`,
							Description: `Flag that indicates whether the response body should be in the prettyprint format.`,
							Required:    false,
						},
					},
					URLParameters: []Parameter{
						{
							Name: `groupId`,
							Description: `Unique 24-hexadecimal digit string that identifies your project. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.


NOTE: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
							Required: true,
						},
					},
					Verb: POST,
				},
				Versions: []Version{
					{
						Version: `2023-01-01`,
						RequestContentTypes: []string{
							`json`,
						},
						ResponseContentTypes: []string{
							`json`,
						},
					},
					{
						Version: `2023-02-01`,
						RequestContentTypes: []string{
							`json`,
						},
						ResponseContentTypes: []string{
							`json`,
						},
					},
					{
						Version: `2024-08-05`,
						RequestContentTypes: []string{
							`json`,
						},
						ResponseContentTypes: []string{
							`json`,
						},
					},
					{
						Version: `2024-10-23`,
						RequestContentTypes: []string{
							`json`,
						},
						ResponseContentTypes: []string{
							`json`,
						},
					},
				},
			},
			{
				OperationID: `listClusters`,
				Description: `Returns the details for all clusters in the specific project to which you have access. Clusters contain a group of hosts that maintain the same data set. The response includes clusters with asymmetrically-sized shards. To use this resource, the requesting API Key must have the Project Read Only role. This feature is not  available for serverless clusters.`,
				RequestParameters: RequestParameters{
					URL: `/api/atlas/v2/groups/{groupId}/clusters`,
					QueryParameters: []Parameter{
						{
							Name:        `envelope`,
							Description: `Flag that indicates whether Application wraps the response in an envelope JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.`,
							Required:    false,
						},
						{
							Name:        `includeCount`,
							Description: `Flag that indicates whether the response returns the total number of items (totalCount) in the response.`,
							Required:    false,
						},
						{
							Name:        `itemsPerPage`,
							Description: `Number of items that the response returns per page.`,
							Required:    false,
						},
						{
							Name:        `pageNum`,
							Description: `Number of the page that displays the current set of the total objects that the response returns.`,
							Required:    false,
						},
						{
							Name:        `pretty`,
							Description: `Flag that indicates whether the response body should be in the prettyprint format.`,
							Required:    false,
						},
						{
							Name:        `includeDeletedWithRetainedBackups`,
							Description: `Flag that indicates whether to return Clusters with retain backups.`,
							Required:    false,
						},
					},
					URLParameters: []Parameter{
						{
							Name: `groupId`,
							Description: `Unique 24-hexadecimal digit string that identifies your project. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.


NOTE: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`,
							Required: true,
						},
					},
					Verb: GET,
				},
				Versions: []Version{
					{
						Version:             `2023-01-01`,
						RequestContentTypes: []string{},
						ResponseContentTypes: []string{
							`json`,
						},
					},
					{
						Version:             `2023-02-01`,
						RequestContentTypes: []string{},
						ResponseContentTypes: []string{
							`json`,
						},
					},
					{
						Version:             `2024-08-05`,
						RequestContentTypes: []string{},
						ResponseContentTypes: []string{
							`json`,
						},
					},
				},
			},
		},
	},
}
