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
            "instanceSize": "M30",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M30",
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
            "instanceSize": "M10",
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
            "instanceSize": "M10",
            "nodeCount": 3
          },
          "priority": 7,
          "providerName": "AWS",
          "readOnlySpecs": {
            "diskSizeGB": 10,
            "instanceSize": "M10",
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
	"getCluster": {
		ParameterExample: map[string]string{
			"groupId": "32b6e34b3d91647abb20e7b8",
		},
		RequestBodyExamples: map[string][]RequestBodyExample{},
	},
}
