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

package api

var EndpointExamples = map[string]Examples{
	"acceptVpcPeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"acknowledgeAlert": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addAllTeamsToProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addOrganizationRole": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addProjectApiKey": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addProjectRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addProjectServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addProjectUser": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addTeamUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addUserToProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"addUserToTeam": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"authorizeCloudProviderAccessRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"cancelBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createAlertConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createApiKey": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createApiKeyAccessList": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createAtlasResourcePolicy": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createAtlasSearchDeployment": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createAtlasSearchIndex": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createAtlasSearchIndexDeprecated": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createBackupExportJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createCloudProviderAccessRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{
			"2024-08-05": {
				{
					Name:        "Cluster",
					Description: "Cluster",
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
			"2024-10-23": {
				{
					Name:        "Cluster",
					Description: "Cluster",
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
				{
					Name:        "Cluster",
					Description: "Cluster",
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
	"createCostExplorerQueryProcess": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createCustomDatabaseRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createCustomZoneMapping": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createDataFederationPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createDatabaseUser": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{
			"2023-01-01": {
				{
					Name:        "AWS IAM Authentication",
					Description: "AWS IAM Authentication",
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
				},
				{
					Name:        "LDAP Authentication",
					Description: "LDAP Authentication",
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
				},
				{
					Name:        "OIDC Workforce Federated Authentication",
					Description: "OIDC Workforce Federated Authentication",
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
				},
				{
					Name:        "OIDC Workload Federated Authentication",
					Description: "OIDC Workload Federated Authentication",
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
				},
				{
					Name:        "SCRAM-SHA Authentication",
					Description: "SCRAM-SHA Authentication",
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
				},
				{
					Name:        "X509 Authentication",
					Description: "X509 Authentication",
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
	"createDatabaseUserCertificate": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createEncryptionAtRestPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createExportBucket": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{
			"2023-01-01": {
				{
					Name:        "AWS",
					Description: "AWS",
					Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592"
}`,
				},
			},
			"2024-05-30": {
				{
					Name:        "AWS",
					Description: "AWS",
					Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592"
}`,
				},
				{
					Name:        "AWS",
					Description: "AWS",
					Value: `{
  "bucketName": "export-bucket",
  "cloudProvider": "AWS",
  "iamRoleId": "668c5f0ed436263134491592"
}`,
				},
				{
					Name:        "Azure",
					Description: "Azure",
					Value: `{
  "cloudProvider": "AZURE",
  "roleId": "668c5f0ed436263134491592",
  "serviceUrl": "https://examplestorageaccount.blob.core.windows.net/examplecontainer"
}`,
				},
			},
		},
	},
	"createFederatedDatabase": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createFlexBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createFlexCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createIdentityProvider": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createLegacyBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createLinkToken": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createManagedNamespace": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createOneDataFederationQueryLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createOnlineArchive": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createOrganization": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createOrganizationInvitation": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createOrganizationUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPeeringContainer": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPrivateEndpointService": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPrivateLinkConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProject": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectApiKey": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectInvitation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectIpAccessList": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectServiceAccount": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectServiceAccountAccessList": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createProjectServiceAccountSecret": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPushBasedLogConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createPushMigration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createRoleMapping": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"orgId":                "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createRollingIndex": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{
			"2023-01-01": {
				{
					Name:        "2dspere Index",
					Description: "2dspere Index",
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
				},
				{
					Name:        "Partial Index",
					Description: "Partial Index",
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
				},
				{
					Name:        "Sparse Index",
					Description: "Sparse Index",
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
	"createServerlessBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createServerlessInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createServerlessPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createServiceAccount": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createServiceAccountAccessList": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createServiceAccountSecret": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createSharedClusterBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createStreamConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createStreamInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createStreamInstanceWithSampleConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createTeam": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createThirdPartyIntegration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"createUser": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"cutoverMigration": {
		ParameterExample: map[string]string{
			"groupId":         "32b6e34b3d91647abb20e7b8",
			"liveMigrationId": "6296fb4c7c7aa997cf94e9a8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deauthorizeCloudProviderAccessRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deferMaintenanceWindow": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAlertConfiguration": {
		ParameterExample: map[string]string{
			"alertConfigId": "32b6e34b3d91647abb20e7b8",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAllBackupSchedules": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAllCustomZoneMappings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteApiKey": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteApiKeyAccessListEntry": {
		ParameterExample: map[string]string{
			"ipAddress": "192.0.2.0%2F24",
			"orgId":     "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAtlasResourcePolicy": {
		ParameterExample: map[string]string{
			"orgId":            "4888442a3354817a7320eb61",
			"resourcePolicyId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAtlasSearchDeployment": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAtlasSearchIndex": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAtlasSearchIndexByName": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteAtlasSearchIndexDeprecated": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteCustomDatabaseRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteDataFederationPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteDatabaseUser": {
		ParameterExample: map[string]string{
			"groupId":  "32b6e34b3d91647abb20e7b8",
			"username": "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteExportBucket": {
		ParameterExample: map[string]string{
			"exportBucketId": "32b6e34b3d91647abb20e7b8",
			"groupId":        "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteFederatedDatabase": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteFederationApp": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteFlexCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteIdentityProvider": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"identityProviderId":   "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteLdapConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteLegacySnapshot": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteLinkToken": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteManagedNamespace": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteOneDataFederationInstanceQueryLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteOnlineArchive": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteOrganization": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteOrganizationInvitation": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePeeringContainer": {
		ParameterExample: map[string]string{
			"containerId": "32b6e34b3d91647abb20e7b8",
			"groupId":     "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePipelineRunDataset": {
		ParameterExample: map[string]string{
			"groupId":       "32b6e34b3d91647abb20e7b8",
			"pipelineRunId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePrivateEndpointService": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePrivateLinkConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectInvitation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectIpAccessList": {
		ParameterExample: map[string]string{
			"entryValue": "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8",
			"groupId":    "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectServiceAccountAccessListEntry": {
		ParameterExample: map[string]string{
			"clientId":  "mdb_sa_id_1234567890abcdef12345678",
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"ipAddress": "192.0.2.0%2F24",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteProjectServiceAccountSecret": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deletePushBasedLogConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteReplicaSetBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteRoleMapping": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"id":                   "32b6e34b3d91647abb20e7b8",
			"orgId":                "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteServerlessInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteServerlessPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteServiceAccountAccessListEntry": {
		ParameterExample: map[string]string{
			"clientId":  "mdb_sa_id_1234567890abcdef12345678",
			"ipAddress": "192.0.2.0%2F24",
			"orgId":     "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteServiceAccountSecret": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteShardedClusterBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteStreamConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteStreamInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteTeam": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteThirdPartyIntegration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"deleteVpcPeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"disableCustomerManagedX509": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"disableDataProtectionSettings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"disablePeering": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"disableSlowOperationThresholding": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadFederatedDatabaseQueryLogs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadFlexBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadInvoiceCsv": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadOnlineArchiveQueryLogs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadSharedClusterBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"downloadStreamTenantAuditLogs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"enableSlowOperationThresholding": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"endOutageSimulation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAccountDetails": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getActiveVpcPeeringConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAlert": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAlertConfiguration": {
		ParameterExample: map[string]string{
			"alertConfigId": "32b6e34b3d91647abb20e7b8",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getApiKey": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getApiKeyAccessList": {
		ParameterExample: map[string]string{
			"ipAddress": "192.0.2.0%2F24",
			"orgId":     "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getApiVersions": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasProcess": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasResourcePolicies": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasResourcePolicy": {
		ParameterExample: map[string]string{
			"orgId":            "4888442a3354817a7320eb61",
			"resourcePolicyId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasSearchDeployment": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasSearchIndex": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasSearchIndexByName": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAtlasSearchIndexDeprecated": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAuditingConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getAwsCustomDns": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getBackupExportJob": {
		ParameterExample: map[string]string{
			"exportId": "32b6e34b3d91647abb20e7b8",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getBackupSchedule": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCloudProviderAccessRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getClusterAdvancedConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getClusterStatus": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCollStatsLatencyNamespaceClusterMeasurements": {
		ParameterExample: map[string]string{
			"collectionName": "mycoll",
			"databaseName":   "mydb",
			"groupId":        "32b6e34b3d91647abb20e7b8",
			"period":         "PT10H",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCollStatsLatencyNamespaceHostMeasurements": {
		ParameterExample: map[string]string{
			"collectionName": "mycoll",
			"databaseName":   "mydb",
			"groupId":        "32b6e34b3d91647abb20e7b8",
			"period":         "PT10H",
			"processId":      "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCollStatsLatencyNamespaceMetrics": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCollStatsLatencyNamespacesForCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
			"period":  "PT10H",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCollStatsLatencyNamespacesForHost": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"period":    "PT10H",
			"processId": "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getConnectedOrgConfig": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"orgId":                "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCostExplorerQueryProcess": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
			"token": "4ABBE973862346D40F3AE859D4BE96E0F895764EB14EAB039E7B82F9D638C05C",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getCustomDatabaseRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDataFederationPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDataProtectionSettings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDatabase": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDatabaseMeasurements": {
		ParameterExample: map[string]string{
			"granularity": "PT1M",
			"groupId":     "32b6e34b3d91647abb20e7b8",
			"period":      "PT10H",
			"processId":   "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDatabaseUser": {
		ParameterExample: map[string]string{
			"groupId":  "32b6e34b3d91647abb20e7b8",
			"username": "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getDiskMeasurements": {
		ParameterExample: map[string]string{
			"granularity": "PT1M",
			"groupId":     "32b6e34b3d91647abb20e7b8",
			"period":      "PT10H",
			"processId":   "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getEncryptionAtRest": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getEncryptionAtRestPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getEncryptionAtRestPrivateEndpointsForCloudProvider": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getExportBucket": {
		ParameterExample: map[string]string{
			"exportBucketId": "32b6e34b3d91647abb20e7b8",
			"groupId":        "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getFederatedDatabase": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getFederationSettings": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getFlexBackup": {
		ParameterExample: map[string]string{
			"groupId":    "32b6e34b3d91647abb20e7b8",
			"snapshotId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getFlexBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getFlexCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getHostLogs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getHostMeasurements": {
		ParameterExample: map[string]string{
			"granularity": "PT1M",
			"groupId":     "32b6e34b3d91647abb20e7b8",
			"period":      "PT10H",
			"processId":   "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getIdentityProvider": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"identityProviderId":   "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getIdentityProviderMetadata": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"identityProviderId":   "c2777a9eca931f29fc2f",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getIndexMetrics": {
		ParameterExample: map[string]string{
			"collectionName": "mycoll",
			"databaseName":   "mydb",
			"granularity":    "PT1M",
			"groupId":        "32b6e34b3d91647abb20e7b8",
			"indexName":      "myindex",
			"period":         "PT10H",
			"processId":      "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getInvoice": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLdapConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLdapConfigurationStatus": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLegacyBackupCheckpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLegacyBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLegacySnapshot": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getLegacySnapshotSchedule": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getMaintenanceWindow": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getManagedNamespace": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getManagedSlowMs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getMeasurements": {
		ParameterExample: map[string]string{
			"granularity": "PT1M",
			"groupId":     "32b6e34b3d91647abb20e7b8",
			"period":      "PT10H",
			"processId":   "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOnlineArchive": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOpenApiInfo": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOrganization": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOrganizationEvent": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOrganizationInvitation": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOrganizationSettings": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOrganizationUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getOutageSimulation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPeeringContainer": {
		ParameterExample: map[string]string{
			"containerId": "32b6e34b3d91647abb20e7b8",
			"groupId":     "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPinnedNamespaces": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPipelineRun": {
		ParameterExample: map[string]string{
			"groupId":       "32b6e34b3d91647abb20e7b8",
			"pipelineRunId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPrivateEndpointService": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPrivateLinkConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectByName": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectEvent": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectInvitation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectIpAccessListStatus": {
		ParameterExample: map[string]string{
			"entryValue": "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8",
			"groupId":    "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectIpList": {
		ParameterExample: map[string]string{
			"entryValue": "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8",
			"groupId":    "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectLtsVersions": {
		ParameterExample: map[string]string{
			"groupId":      "32b6e34b3d91647abb20e7b8",
			"instanceSize": "M10",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectSettings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getProjectUser": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPushBasedLogConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getPushMigration": {
		ParameterExample: map[string]string{
			"groupId":         "32b6e34b3d91647abb20e7b8",
			"liveMigrationId": "6296fb4c7c7aa997cf94e9a8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getRegionalizedPrivateEndpointSetting": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getReplicaSetBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getResourcesNonCompliant": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getRoleMapping": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"id":                   "32b6e34b3d91647abb20e7b8",
			"orgId":                "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getSampleDatasetLoadStatus": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServerlessAutoIndexing": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServerlessBackup": {
		ParameterExample: map[string]string{
			"groupId":    "32b6e34b3d91647abb20e7b8",
			"snapshotId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServerlessBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServerlessInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServerlessPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getShardedClusterBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getSharedClusterBackup": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getSharedClusterBackupRestoreJob": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getStreamConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getStreamInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getSystemStatus": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getTeamById": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getTeamByName": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getThirdPartyIntegration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getUser": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getUserByUsername": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getValidationStatus": {
		ParameterExample: map[string]string{
			"groupId":      "32b6e34b3d91647abb20e7b8",
			"validationId": "507f1f77bcf86cd799439011",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"getVpcPeeringConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"grantMongoDbEmployeeAccess": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAccessLogsByClusterName": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAccessLogsByHostname": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAlertConfigurationMatchersFieldNames": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAlertConfigurations": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAlertConfigurationsByAlertId": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAlerts": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAlertsByAlertConfigurationId": {
		ParameterExample: map[string]string{
			"alertConfigId": "32b6e34b3d91647abb20e7b8",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listApiKeyAccessListsEntries": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listApiKeys": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAtlasProcesses": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAtlasSearchIndexes": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAtlasSearchIndexesCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listAtlasSearchIndexesDeprecated": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listBackupExportJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listBackupRestoreJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listCloudProviderAccessRoles": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listCloudProviderRegions": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listClusterSuggestedIndexes": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listClusters": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listClustersForAllProjects": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listConnectedOrgConfigs": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listCustomDatabaseRoles": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDataFederationPrivateEndpoints": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDatabaseUserCertificates": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDatabaseUsers": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDatabases": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDiskMeasurements": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDiskPartitions": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "mongodb.example.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listDropIndexes": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listEventTypes": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listExportBuckets": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listFederatedDatabases": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listFlexBackupRestoreJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listFlexBackups": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listFlexClusters": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listIdentityProviders": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listIndexMetrics": {
		ParameterExample: map[string]string{
			"collectionName": "mycoll",
			"databaseName":   "mydb",
			"granularity":    "PT1M",
			"groupId":        "32b6e34b3d91647abb20e7b8",
			"period":         "PT10H",
			"processId":      "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listInvoices": {
		ParameterExample: map[string]string{
			"fromDate": "2023-01-01",
			"orgId":    "4888442a3354817a7320eb61",
			"toDate":   "2023-01-01",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listLegacyBackupCheckpoints": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listLegacyBackupRestoreJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listLegacySnapshots": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listMetricTypes": {
		ParameterExample: map[string]string{
			"groupId":   "32b6e34b3d91647abb20e7b8",
			"processId": "my.host.name.com:27017",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOnlineArchives": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizationEvents": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizationInvitations": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizationProjects": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizationTeams": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizationUsers": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listOrganizations": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPeeringConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPeeringContainerByCloudProvider": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPeeringContainers": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPendingInvoices": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPipelineRuns": {
		ParameterExample: map[string]string{
			"createdBefore": "2022-01-01T00:00:00Z",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPipelineSchedules": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPipelineSnapshots": {
		ParameterExample: map[string]string{
			"completedAfter": "2022-01-01T00:00:00Z",
			"groupId":        "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPipelines": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPrivateEndpointServices": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listPrivateLinkConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectApiKeys": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectEvents": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectInvitations": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectIpAccessLists": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectLimits": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectServiceAccountAccessList": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectServiceAccounts": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectTeams": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjectUsers": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listProjects": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listReplicaSetBackups": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listRoleMappings": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"orgId":                "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSchemaAdvice": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServerlessBackupRestoreJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServerlessBackups": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServerlessInstances": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServerlessPrivateEndpoints": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServiceAccountAccessList": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServiceAccountProjects": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listServiceAccounts": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listShardedClusterBackups": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSharedClusterBackupRestoreJobs": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSharedClusterBackups": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSlowQueries": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSlowQueryNamespaces": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSourceProjects": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listStreamConnections": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listStreamInstances": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listStreamProcessors": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listSuggestedIndexes": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listTeamUsers": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"listThirdPartyIntegrations": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"loadSampleDataset": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"migrateProjectToAnotherOrg": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"modifyStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"pausePipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"pinFeatureCompatibilityVersion": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"pinNamespacesPatch": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"pinNamespacesPut": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"queryLineItemsFromSingleInvoice": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"rejectVpcPeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeConnectedOrgConfig": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"orgId":                "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeOrganizationRole": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeOrganizationUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeProjectApiKey": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeProjectRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeProjectTeam": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeProjectUser": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeTeamUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"removeUserFromTeam": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"renameTeam": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"requestEncryptionAtRestPrivateEndpointDeletion": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"resetMaintenanceWindow": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"resumePipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"returnAllControlPlaneIpAddresses": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"returnAllIpAddresses": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"returnFederatedDatabaseQueryLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"returnFederatedDatabaseQueryLimits": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"revokeJwksFromIdentityProvider": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"identityProviderId":   "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"revokeMongoDbEmployeeAccess": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"saveLdapConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"setProjectLimit": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"setServerlessAutoIndexing": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"startOutageSimulation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"startStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"stopStreamProcessor": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"takeSnapshot": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"testFailover": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"toggleAlertConfiguration": {
		ParameterExample: map[string]string{
			"alertConfigId": "32b6e34b3d91647abb20e7b8",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"toggleAwsCustomDns": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"toggleMaintenanceAutoDefer": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"toggleRegionalizedPrivateEndpointSetting": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"triggerSnapshotIngestion": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"unpinFeatureCompatibilityVersion": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"unpinNamespaces": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAlertConfiguration": {
		ParameterExample: map[string]string{
			"alertConfigId": "32b6e34b3d91647abb20e7b8",
			"groupId":       "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateApiKey": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateApiKeyRoles": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAtlasResourcePolicy": {
		ParameterExample: map[string]string{
			"orgId":            "4888442a3354817a7320eb61",
			"resourcePolicyId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAtlasSearchDeployment": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAtlasSearchIndex": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAtlasSearchIndexByName": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAtlasSearchIndexDeprecated": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateAuditingConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateBackupSchedule": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateClusterAdvancedConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateConnectedOrgConfig": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"orgId":                "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateCustomDatabaseRole": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateDataProtectionSettings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateDatabaseUser": {
		ParameterExample: map[string]string{
			"groupId":  "32b6e34b3d91647abb20e7b8",
			"username": "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs or x.509/LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM or OIDC: IdPIdentifier/IdPGroupName",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateEncryptionAtRest": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateFederatedDatabase": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateFlexCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateIdentityProvider": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"identityProviderId":   "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateLegacySnapshotRetention": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateLegacySnapshotSchedule": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateMaintenanceWindow": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOnlineArchive": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganization": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganizationInvitation": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganizationInvitationById": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganizationRoles": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganizationSettings": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateOrganizationUser": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updatePeeringConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updatePeeringContainer": {
		ParameterExample: map[string]string{
			"containerId": "32b6e34b3d91647abb20e7b8",
			"groupId":     "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updatePipeline": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProjectInvitation": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProjectInvitationById": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProjectRoles": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProjectServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"groupId":  "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateProjectSettings": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updatePushBasedLogConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateRoleMapping": {
		ParameterExample: map[string]string{
			"federationSettingsId": "55fa922fb343282757d9554e",
			"id":                   "32b6e34b3d91647abb20e7b8",
			"orgId":                "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateServerlessInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateServerlessPrivateEndpoint": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateServiceAccount": {
		ParameterExample: map[string]string{
			"clientId": "mdb_sa_id_1234567890abcdef12345678",
			"orgId":    "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateSnapshotRetention": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateStreamConnection": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateStreamInstance": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateTeamRoles": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"updateThirdPartyIntegration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"upgradeFlexCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"upgradeSharedCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"validateAtlasResourcePolicy": {
		ParameterExample: map[string]string{
			"orgId": "4888442a3354817a7320eb61",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"validateMigration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"verifyConnectViaPeeringOnlyModeForOneProject": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"verifyLdapConfiguration": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
	"versionedExample": {
		ParameterExample:    map[string]string{},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
}
