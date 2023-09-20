// Copyright 2023 MongoDB Inc
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
	"time"
	"net"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/auth"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
)

const (
	cloudGovServiceURL    = "https://cloud.mongodbgov.com/"
	timeout               = 5 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 5
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
)

func httpClient(username, password string, accessToken *auth.Token) (*http.Client, error) {
	httpTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		Proxy:                 http.ProxyFromEnvironment,
		IdleConnTimeout:       idleConnTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
	}

	if username == "" && password == "" && accessToken == nil {
		return &http.Client{Transport: httpTransport}, nil
	}
	if username != "" && password != "" {
		t := &digest.Transport{
			Username: username,
			Password: password,
		}
		t.Transport = httpTransport
		return t.Client()
	}
	tr := &Transport{
		token: accessToken,
		base:  httpTransport,
	}

	return &http.Client{Transport: tr}, nil
}

func newClientWithAuth() (*admin.APIClient, error) {
	profile := config.Default()

	var authToken *auth.Token

	username := profile.PublicAPIKey()
	password := profile.PrivateAPIKey()

	if username == "" && password == "" {
		var err error
		authToken, err = profile.Token()
		if err != nil {
			return nil, err
		}
	}

	baseURL := profile.OpsManagerURL()
	if baseURL == "" && profile.Service() == config.CloudGovService {
		baseURL = cloudGovServiceURL
	}

	client, err := httpClient(username, password, authToken)
	if err != nil {
		return nil, err
	}

	opts := []admin.ClientModifier{
		admin.UseHTTPClient(client),
		admin.UseUserAgent(config.UserAgent),
		admin.UseDebug(log.IsDebugLevel())}

	if baseURL != "" {
		opts = append(opts, admin.UseBaseURL(baseURL))
	}
	c, err := admin.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Transport struct {
	token *auth.Token
	base  http.RoundTripper
}

func (tr *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.token.SetAuthHeader(req)
	return tr.base.RoundTrip(req)
}

func convertTime(s *string) *time.Time {
	if s == nil {
		return nil
	}

	r, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil
	}
	return &r
}

func Builder() *cobra.Command {
	const use = "api"
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
		legacyBackupRestoreJobsBuilder(),
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
