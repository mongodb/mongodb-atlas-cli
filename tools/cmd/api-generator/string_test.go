// Copyright 2024 MongoDB Inc
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

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			`Flag that indicates whether the response body should be in the <a href="https://en.wikipedia.org/wiki/Prettyprint" target="_blank" rel="noopener noreferrer">prettyprint</a> format.`,
			`Flag that indicates whether the response body should be in the prettyprint format.`,
		},
		{
			`**Note**: This resource cannot be used to add a user invited via the deprecated [Invite One MongoDB Cloud User to Join One Project](#tag/Projects/operation/createProjectInvitation) endpoint.`,
			`Note: This resource cannot be used to add a user invited via the deprecated Invite One MongoDB Cloud User to Join One Project endpoint.`,
		},
		{
			`The delimiter that separates **databases.[n].collections.[n].dataSources.[n].path** segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to ` + "`" + `/` + "`" + `.`,
			`The delimiter that separates databases.[n].collections.[n].dataSources.[n].path segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to /.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output, err := Clean(tt.input)

			require.NoError(t, err)
			require.Equal(t, tt.output, output)
		})
	}
}

func TestSafeSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Access Tracking", "Access-Tracking"},
		{"Alert Configurations", "Alert-Configurations"},
		{"Alerts", "Alerts"},
		{"Atlas Search", "Atlas-Search"},
		{"Auditing", "Auditing"},
		{"AWS Clusters DNS", "AWS-Clusters-DNS"},
		{"Cloud Backups", "Cloud-Backups"},
		{"Cloud Migration Service", "Cloud-Migration-Service"},
		{"Cloud Provider Access", "Cloud-Provider-Access"},
		{"Cluster Outage Simulation", "Cluster-Outage-Simulation"},
		{"Clusters", "Clusters"},
		{"Collection Level Metrics", "Collection-Level-Metrics"},
		{"Custom Database Roles", "Custom-Database-Roles"},
		{"Data Federation", "Data-Federation"},
		{"Data Lake Pipelines", "Data-Lake-Pipelines"},
		{"Database Users", "Database-Users"},
		{"Encryption at Rest using Customer Key Management", "Encryption-at-Rest-using-Customer-Key-Management"},
		{"Events", "Events"},
		{"Federated Authentication", "Federated-Authentication"},
		{"Flex Clusters", "Flex-Clusters"},
		{"Flex Restore Jobs", "Flex-Restore-Jobs"},
		{"Flex Snapshots", "Flex-Snapshots"},
		{"Global Clusters", "Global-Clusters"},
		{"Internal", "Internal"},
		{"Invoices", "Invoices"},
		{"LDAP Configuration", "LDAP-Configuration"},
		{"Legacy Backup", "Legacy-Backup"},
		{"Maintenance Windows", "Maintenance-Windows"},
		{"MongoDB Cloud Users", "MongoDB-Cloud-Users"},
		{"Monitoring and Logs", "Monitoring-and-Logs"},
		{"Network Peering", "Network-Peering"},
		{"Online Archive", "Online-Archive"},
		{"OpenAPI", "OpenAPI"},
		{"Organizations", "Organizations"},
		{"Performance Advisor", "Performance-Advisor"},
		{"Private Endpoint Services", "Private-Endpoint-Services"},
		{"Programmatic API Keys", "Programmatic-API-Keys"},
		{"Project IP Access List", "Project-IP-Access-List"},
		{"Projects", "Projects"},
		{"Push-Based Log Export", "Push-Based-Log-Export"},
		{"Resource Policies", "Resource-Policies"},
		{"Rolling Index", "Rolling-Index"},
		{"Root", "Root"},
		{"Serverless Instances", "Serverless-Instances"},
		{"Serverless Private Endpoints", "Serverless-Private-Endpoints"},
		{"Service Accounts", "Service-Accounts"},
		{"Shared-Tier Restore Jobs", "Shared-Tier-Restore-Jobs"},
		{"Shared-Tier Snapshots", "Shared-Tier-Snapshots"},
		{"Streams", "Streams"},
		{"Teams", "Teams"},
		{"Test", "Test"},
		{"Third-Party Integrations", "Third-Party-Integrations"},
		{"X.509 Authentication", "X.509-Authentication"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := safeSlugify(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
