// Copyright 2024 MongoDB Inc
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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Access to api resources.",
		Long:  `This command provides access to API resources specified in https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/.`,
	}

	cmd.AddCommand(
		aWSClustersDNSBuilder(),
		accessTrackingBuilder(),
		alertConfigurationsBuilder(),
		alertsBuilder(),
		atlasSearchBuilder(),
		auditingBuilder(),
		cloudBackupsBuilder(),
		cloudMigrationServiceBuilder(),
		cloudProviderAccessBuilder(),
		clusterOutageSimulationBuilder(),
		clustersBuilder(),
		collectionLevelMetricsBuilder(),
		customDatabaseRolesBuilder(),
		dataFederationBuilder(),
		dataLakePipelinesBuilder(),
		databaseUsersBuilder(),
		encryptionAtRestUsingCustomerKeyManagementBuilder(),
		eventsBuilder(),
		federatedAuthenticationBuilder(),
		globalClustersBuilder(),
		invoicesBuilder(),
		lDAPConfigurationBuilder(),
		legacyBackupBuilder(),
		maintenanceWindowsBuilder(),
		mongoDBCloudUsersBuilder(),
		monitoringAndLogsBuilder(),
		networkPeeringBuilder(),
		onlineArchiveBuilder(),
		organizationsBuilder(),
		performanceAdvisorBuilder(),
		privateEndpointServicesBuilder(),
		programmaticAPIKeysBuilder(),
		projectIPAccessListBuilder(),
		projectsBuilder(),
		pushBasedLogExportBuilder(),
		rollingIndexBuilder(),
		rootBuilder(),
		serverlessInstancesBuilder(),
		serverlessPrivateEndpointsBuilder(),
		sharedTierRestoreJobsBuilder(),
		sharedTierSnapshotsBuilder(),
		streamsBuilder(),
		teamsBuilder(),
		thirdPartyIntegrationsBuilder(),
		x509AuthenticationBuilder(),
	)

	return cmd
}
