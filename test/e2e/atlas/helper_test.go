// Copyright 2021 MongoDB Inc
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
//go:build e2e || atlas

package atlas_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	eventsEntity                  = "events"
	clustersEntity                = "clusters"
	processesEntity               = "processes"
	metricsEntity                 = "metrics"
	searchEntity                  = "search"
	indexEntity                   = "index"
	nodesEntity                   = "nodes"
	datafederationEntity          = "datafederation"
	datalakePipelineEntity        = "datalakepipeline"
	alertsEntity                  = "alerts"
	configEntity                  = "settings"
	dbusersEntity                 = "dbusers"
	certsEntity                   = "certs"
	privateEndpointsEntity        = "privateendpoints"
	queryLimitsEntity             = "querylimits"
	onlineArchiveEntity           = "onlineArchives"
	projectEntity                 = "project"
	orgEntity                     = "org"
	invitationsEntity             = "invitations"
	maintenanceEntity             = "maintenanceWindows"
	integrationsEntity            = "integrations"
	securityEntity                = "security"
	ldapEntity                    = "ldap"
	awsEntity                     = "aws"
	azureEntity                   = "azure"
	gcpEntity                     = "gcp"
	customDNSEntity               = "customDns"
	logsEntity                    = "logs"
	cloudProvidersEntity          = "cloudProviders"
	accessRolesEntity             = "accessRoles"
	customDBRoleEntity            = "customDbRoles"
	regionalModeEntity            = "regionalModes"
	serverlessEntity              = "serverless"
	liveMigrationsEntity          = "liveMigrations"
	auditingEntity                = "auditing"
	accessLogsEntity              = "accessLogs"
	accessListEntity              = "accessList"
	performanceAdvisorEntity      = "performanceAdvisor"
	slowQueryLogsEntity           = "slowQueryLogs"
	namespacesEntity              = "namespaces"
	networkingEntity              = "networking"
	networkPeeringEntity          = "peering"
	suggestedIndexesEntity        = "suggestedIndexes"
	slowOperationThresholdEntity  = "slowOperationThreshold"
	tierM10                       = "M10"
	tierM0                        = "M0"
	tierM2                        = "M2"
	diskSizeGB40                  = "40"
	diskSizeGB30                  = "30"
	projectsEntity                = "projects"
	settingsEntity                = "settings"
	backupsEntity                 = "backups"
	exportsEntity                 = "exports"
	bucketsEntity                 = "buckets"
	jobsEntity                    = "jobs"
	snapshotsEntity               = "snapshots"
	restoresEntity                = "restores"
	compliancePolicyEntity        = "compliancepolicy"
	policiesEntity                = "policies"
	teamsEntity                   = "teams"
	setupEntity                   = "setup"
	deploymentEntity              = "deployments"
	federatedAuthenticationEntity = "federatedAuthentication"
	federationSettingsEntity      = "federationSettings"
	identityProviderEntity        = "identityProvider"
	connectedOrgsConfigsEntity    = "connectedOrgConfigs"
	deletingState                 = "DELETING"
	authEntity                    = "auth"
	streamsEntity                 = "streams"
)

// AlertConfig constants.
const (
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	intervalMin   = 5
	delayMin      = 0
)

// Auth constants.
const (
	whoami = "whoami"
)

// Integration constants.
const (
	datadogEntity   = "DATADOG"
	opsGenieEntity  = "OPS_GENIE"
	pagerDutyEntity = "PAGER_DUTY"
	victorOpsEntity = "VICTOR_OPS"
	webhookEntity   = "WEBHOOK"
)

// Cluster settings.
const (
	e2eClusterTier       = "M10"
	e2eGovClusterTier    = "M20"
	e2eSharedClusterTier = "M2"
	e2eClusterProvider   = "AWS" // e2eClusterProvider preferred provider for e2e testing.
)

// Backup compliance policy constants.
const (
	authorizedUserFirstName = "firstname"
	authorizedUserLastName  = "lastname"
	authorizedEmail         = "firstname.lastname@example.com"
)

// Local Development constants.
const (
	collectionName  = "myCol"
	databaseName    = "myDB"
	searchIndexName = "indexTest"
	vectorSearchDB  = "sample_mflix"
	vectorSearchCol = "embedded_movies"
)

func splitOutput(cmd *exec.Cmd) (string, string, error) {
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	err := cmd.Run()
	return o.String(), e.String(), err
}

func deployServerlessInstanceForProject(projectID string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}
	clusterName, err := RandClusterName()
	if err != nil {
		return "", err
	}
	tier := e2eTier()
	region, err := newAvailableRegion(projectID, tier, e2eClusterProvider)
	if err != nil {
		return "", err
	}
	args := []string{
		serverlessEntity,
		"create",
		clusterName,
		"--region", region,
		"--provider", e2eClusterProvider,
	}

	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(create); err != nil {
		return "", fmt.Errorf("error creating serverless instance %w: %s", err, string(resp))
	}

	watchArgs := []string{
		serverlessEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watch := exec.Command(cliPath, watchArgs...)
	watch.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(watch); err != nil {
		return "", fmt.Errorf("error watching serverless instance %w: %s", err, string(resp))
	}
	return clusterName, nil
}

func watchServerlessInstanceForProject(projectID, clusterName string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}

	watchArgs := []string{
		serverlessEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watchCmd := exec.Command(cliPath, watchArgs...)
	watchCmd.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(watchCmd); err != nil {
		return fmt.Errorf("error watching serverless instance %w: %s", err, string(resp))
	}
	return nil
}

func deleteServerlessInstanceForProject(t *testing.T, cliPath, projectID, clusterName string) {
	t.Helper()

	args := []string{
		serverlessEntity,
		"delete",
		clusterName,
		"--force",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	deleteCmd := exec.Command(cliPath, args...)
	deleteCmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(deleteCmd)
	require.NoError(t, err, string(resp))

	_ = watchServerlessInstanceForProject(projectID, clusterName)
}

func deployClusterForProject(projectID, tier, mDBVersion string, enableBackup bool) (string, string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", "", err
	}
	clusterName, err := RandClusterName()
	if err != nil {
		return "", "", err
	}
	region, err := newAvailableRegion(projectID, tier, e2eClusterProvider)
	if err != nil {
		return "", "", err
	}
	args := []string{
		clustersEntity,
		"create",
		clusterName,
		"--mdbVersion", mDBVersion,
		"--region", region,
		"--tier", tier,
		"--provider", e2eClusterProvider,
		"--diskSizeGB=30",
	}
	if enableBackup {
		args = append(args, "--backup")
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(create); err != nil {
		return "", "", fmt.Errorf("error creating cluster %w: %s", err, string(resp))
	}

	watchArgs := []string{
		clustersEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watch := exec.Command(cliPath, watchArgs...)
	watch.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(watch); err != nil {
		return "", "", fmt.Errorf("error watching cluster %w: %s", err, string(resp))
	}
	return clusterName, region, nil
}

func e2eTier() string {
	tier := e2eClusterTier
	if IsGov() {
		tier = e2eGovClusterTier
	}
	return tier
}

func internalDeleteClusterForProject(projectID, clusterName string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
	args := []string{
		clustersEntity,
		"delete",
		clusterName,
		"--force",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	deleteCmd := exec.Command(cliPath, args...)
	deleteCmd.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(deleteCmd); err != nil {
		return fmt.Errorf("error deleting cluster %w: %s", err, string(resp))
	}

	// this command will fail with 404 once the cluster is deleted
	// we just need to wait for this to close the project
	_ = watchCluster(projectID, clusterName)
	return nil
}

func watchCluster(projectID, clusterName string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
	watchArgs := []string{
		clustersEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watchCmd := exec.Command(cliPath, watchArgs...)
	watchCmd.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(watchCmd); err != nil {
		return fmt.Errorf("error waiting for cluster %w: %s", err, string(resp))
	}
	return nil
}

func removeTerminationProtectionFromCluster(projectID, clusterName string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
	args := []string{
		clustersEntity,
		"update",
		clusterName,
		"--disableTerminationProtection",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	updateCmd := exec.Command(cliPath, args...)
	updateCmd.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(updateCmd); err != nil {
		return fmt.Errorf("error updating cluster %w: %s", err, string(resp))
	}

	return watchCluster(projectID, clusterName)
}

func deleteClusterForProject(projectID, clusterName string) error {
	if err := internalDeleteClusterForProject(projectID, clusterName); err != nil {
		if !strings.Contains(err.Error(), "CANNOT_TERMINATE_CLUSTER_WHEN_TERMINATION_PROTECTION_ENABLED") {
			return err
		}
		if err := removeTerminationProtectionFromCluster(projectID, clusterName); err != nil {
			return err
		}
		return internalDeleteClusterForProject(projectID, clusterName)
	}

	return nil
}

func deleteDatalakeForProject(cliPath, projectID, id string) error {
	args := []string{
		datalakePipelineEntity,
		"delete",
		id,
		"--force",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	deleteCmd := exec.Command(cliPath, args...)
	deleteCmd.Env = os.Environ()
	if resp, err := e2e.RunAndGetStdOut(deleteCmd); err != nil {
		return fmt.Errorf("error deleting datalake %w: %s", err, string(resp))
	}
	return nil
}

var errNoRegions = errors.New("no regions available")

func newAvailableRegion(projectID, tier, provider string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}
	args := []string{
		clustersEntity,
		"availableRegions",
		"ls",
		"--provider", provider,
		"--tier", tier,
		"-o=json",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)

	if err != nil {
		return "", fmt.Errorf("error getting regions %w: %s", err, string(resp))
	}

	var cloudProviders atlasv2.PaginatedApiAtlasProviderRegions
	err = json.Unmarshal(resp, &cloudProviders)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response %w: %s", err, string(resp))
	}

	if cloudProviders.GetTotalCount() == 0 || len(cloudProviders.GetResults()[0].GetInstanceSizes()) == 0 {
		return "", errNoRegions
	}

	return cloudProviders.GetResults()[0].GetInstanceSizes()[0].GetAvailableRegions()[0].GetName(), nil
}

func RandClusterName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("cluster-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("cluster-%v", n), nil
}

func RandIdentityProviderName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("idp-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("idp-%v", n), nil
}

func RandTeamName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("team-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("team-%v", n), nil
}

func RandProjectName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("%v-%s", n, revision), nil
	}
	return fmt.Sprintf("e2e-%v", n), nil
}

func RandUsername() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user-%v", n), nil
}

func RandTeamNameWithPrefix(prefix string) (string, error) {
	name, err := RandTeamName()
	if err != nil {
		return "", err
	}
	prefixedName := fmt.Sprintf("%s-%s", prefix, name)
	if len(prefixedName) > 64 {
		return prefixedName[:64], nil
	}
	return prefixedName, nil
}

func RandProjectNameWithPrefix(prefix string) (string, error) {
	name, err := RandProjectName()
	if err != nil {
		return "", err
	}
	prefixedName := fmt.Sprintf("%s-%s", prefix, name)
	if len(prefixedName) > 64 {
		return prefixedName[:64], nil
	}
	return prefixedName, nil
}

func RandEntityWithRevision(entity string) (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("%s-%v-%s", entity, n, revision), nil
	}
	return fmt.Sprintf("%s-%v", entity, n), nil
}

func MongoDBMajorVersion() (string, error) {
	atlasClient := mongodbatlas.NewClient(nil)
	atlasURL := os.Getenv("MCLI_OPS_MANAGER_URL")
	baseURL, err := url.Parse(atlasURL)
	if err != nil {
		return "", err
	}
	atlasClient.BaseURL = baseURL
	version, _, err := atlasClient.DefaultMongoDBMajorVersion.Get(context.Background())
	if err != nil {
		return "", err
	}

	return version, nil
}

func integrationExists(name string, thirdPartyIntegrations atlasv2.PaginatedIntegration) bool {
	services := thirdPartyIntegrations.GetResults()
	for i := range services {
		iType := getIntegrationType(services[i])
		if iType == name {
			return true
		}
	}
	return false
}

func getIntegrationType(val atlasv2.ThirdPartyIntegration) string {
	return val.GetType()
}

func IsGov() bool {
	return os.Getenv("MCLI_SERVICE") == "cloudgov"
}

func getFirstOrgUser() (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}
	args := []string{
		orgEntity,
		"users",
		"list",
		"-o=json",
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var users atlasv2.PaginatedAppUser
	if err := json.Unmarshal(resp, &users); err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}
	if users.GetTotalCount() == 0 {
		return "", errors.New("no users found")
	}

	return users.GetResults()[0].Username, nil
}

func createTeam(teamName, userName string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin", err)
	}
	args := []string{
		teamsEntity,
		"create",
		teamName,
		"--username",
		userName,
		"-o=json",
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var team atlasv2.Team
	if err := json.Unmarshal(resp, &team); err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	return team.GetId(), nil
}

func createProject(projectName string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin", err)
	}
	args := []string{
		projectEntity,
		"create",
		projectName,
		"-o=json",
	}
	if IsGov() {
		args = append(args, "--govCloudRegionsOnly")
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project atlasv2.Group
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.GetId(), nil
}

func createProjectWithoutAlertSettings(projectName string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin", err)
	}
	args := []string{
		projectEntity,
		"create",
		projectName,
		"-o=json",
		"--withoutDefaultAlertSettings",
	}
	if IsGov() {
		args = append(args, "--govCloudRegionsOnly")
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project atlasv2.Group
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.GetId(), nil
}

func listClustersForProject(t *testing.T, cliPath, projectID string) atlasv2.PaginatedAdvancedClusterDescription {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var clusters atlasv2.PaginatedAdvancedClusterDescription
	require.NoError(t, json.Unmarshal(resp, &clusters))
	return clusters
}

func deleteAllClustersForProject(t *testing.T, cliPath, projectID string) {
	t.Helper()
	clusters := listClustersForProject(t, cliPath, projectID)
	for _, cluster := range clusters.GetResults() {
		func(clusterName, state string) {
			t.Run("delete cluster "+clusterName, func(t *testing.T) {
				t.Parallel()
				if state == deletingState {
					_ = watchCluster(projectID, clusterName)
					return
				}
				assert.NoError(t, deleteClusterForProject(projectID, clusterName))
			})
		}(cluster.GetName(), cluster.GetStateName())
	}
}

func deleteDatapipelinesForProject(t *testing.T, cliPath, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		datalakePipelineEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var pipelines []atlasv2.DataLakeIngestionPipeline
	require.NoError(t, json.Unmarshal(resp, &pipelines))
	for _, p := range pipelines {
		assert.NoError(t, deleteDatalakeForProject(cliPath, projectID, p.GetName()))
	}
}

func deleteAllNetworkPeers(t *testing.T, cliPath, projectID, provider string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		networkingEntity,
		networkPeeringEntity,
		"list",
		"--provider",
		provider,
		"--projectId",
		projectID,
		"-o=json",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	t.Log("available network peers", string(resp))
	require.NoError(t, err, string(resp))
	var networkPeers []atlasv2.BaseNetworkPeeringConnectionSettings
	err = json.Unmarshal(resp, &networkPeers)
	require.NoError(t, err)
	for _, peer := range networkPeers {
		peerID := peer.GetId()
		cmd = exec.Command(cliPath,
			networkingEntity,
			networkPeeringEntity,
			"delete",
			peerID,
			"--projectId",
			projectID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err = e2e.RunAndGetStdOut(cmd)
		assert.NoError(t, err, string(resp))
	}
}

const sleep = 10 * time.Second

func deleteAllPrivateEndpoints(t *testing.T, cliPath, projectID, provider string) {
	t.Helper()

	privateEndpoints := listPrivateEndpointsByProject(t, cliPath, projectID, provider)
	for _, endpoint := range privateEndpoints {
		deletePrivateEndpoint(t, cliPath, projectID, provider, endpoint.GetId())
	}

	done := false
	for attempt := 0; attempt < 10; attempt++ {
		privateEndpoints = listPrivateEndpointsByProject(t, cliPath, projectID, provider)
		if len(privateEndpoints) == 0 {
			t.Logf("all %s private endpoints successfully deleted", provider)
			done = true
			break
		}
		time.Sleep(sleep)
	}

	require.True(t, done, "failed to clean all private endpoints")
}

func listPrivateEndpointsByProject(t *testing.T, cliPath, projectID, provider string) []atlasv2.EndpointService {
	t.Helper()
	cmd := exec.Command(cliPath,
		privateEndpointsEntity,
		provider,
		"list",
		"--projectId",
		projectID,
		"-o=json",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var privateEndpoints []atlasv2.EndpointService
	require.NoError(t, json.Unmarshal(resp, &privateEndpoints))

	return privateEndpoints
}

func deletePrivateEndpoint(t *testing.T, cliPath, projectID, provider, endpointID string) {
	t.Helper()

	cmd := exec.Command(cliPath,
		privateEndpointsEntity,
		provider,
		"delete",
		endpointID,
		"--projectId",
		projectID,
		"--force",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func deleteTeam(teamID string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		teamsEntity,
		"delete",
		teamID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}
	return nil
}

func deleteProject(projectID string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		projectEntity,
		"delete",
		projectID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}
	return nil
}

func createDBUserWithCert(projectID, username string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}

	cmd := exec.Command(cliPath,
		dbusersEntity,
		"create",
		"readAnyDatabase",
		"--username", username,
		"--x509Type", "MANAGED",
		"--projectId", projectID)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}

	return nil
}

func createDataFederationForProject(projectID string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	dataFederationName := fmt.Sprintf("e2e-data-federation-%v", n)

	cmd := exec.Command(cliPath,
		datafederationEntity,
		"create",
		dataFederationName,
		"--projectId", projectID,
		"--region", "DUBLIN_IRL")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	return dataFederationName, nil
}

func listDataFederationsByProject(t *testing.T, cliPath, projectID string) []atlasv2.DataLakeTenant {
	t.Helper()

	cmd := exec.Command(cliPath,
		datafederationEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	t.Log("available datafederations", string(resp))
	require.NoError(t, err, string(resp))

	var dataFederations []atlasv2.DataLakeTenant
	err = json.Unmarshal(resp, &dataFederations)
	require.NoError(t, err)

	return dataFederations
}

func listServerlessByProject(t *testing.T, cliPath, projectID string) *atlasv2.PaginatedServerlessInstanceDescription {
	t.Helper()

	cmd := exec.Command(cliPath,
		serverlessEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	var serverlessInstances *atlasv2.PaginatedServerlessInstanceDescription
	err = json.Unmarshal(resp, &serverlessInstances)
	require.NoError(t, err)

	return serverlessInstances
}

func deleteAllDataFederations(t *testing.T, cliPath, projectID string) {
	t.Helper()

	dataFederations := listDataFederationsByProject(t, cliPath, projectID)
	for _, federation := range dataFederations {
		deleteDataFederationForProject(t, cliPath, projectID, federation.GetName())
	}
	t.Log("all datafederations successfully deleted")
}

func deleteAllServerlessInstances(t *testing.T, cliPath, projectID string) {
	t.Helper()

	serverlessInstances := listServerlessByProject(t, cliPath, projectID)
	for _, serverless := range serverlessInstances.GetResults() {
		func(serverlessInstance, state string) {
			t.Run(fmt.Sprintf("delete serverless instance %s\n", serverlessInstance), func(t *testing.T) {
				t.Parallel()
				if state == deletingState {
					_ = watchServerlessInstanceForProject(projectID, serverlessInstance)
					return
				}
				deleteServerlessInstanceForProject(t, cliPath, projectID, serverlessInstance)
			})
		}(serverless.GetName(), serverless.GetStateName())
	}

	t.Log("all serverless instances successfully deleted")
}

func deleteDataFederationForProject(t *testing.T, cliPath, projectID, dataFedName string) {
	t.Helper()

	cmd := exec.Command(cliPath,
		datafederationEntity,
		"delete",
		dataFedName,
		"--projectId", projectID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func ensureCluster(t *testing.T, cluster *atlasv2.AdvancedClusterDescription, clusterName, version string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.GetName())
	a.Equal(version, cluster.GetMongoDBMajorVersion())
	a.InDelta(diskSizeGB, cluster.GetDiskSizeGB(), 0.01)
	a.Equal(terminationProtection, cluster.GetTerminationProtectionEnabled())
}

func compareStingsWithHiddenPart(expectedSting, actualString string, specialChar uint8) bool {
	if len(expectedSting) != len(actualString) {
		return false
	}
	for i := 0; i < len(expectedSting); i++ {
		if expectedSting[i] != actualString[i] && actualString[i] != specialChar {
			return false
		}
	}
	return true
}

// createJSONFile creates a new JSON file at the specified path with the specified data
// and also registers its deletion on test cleanup.
func createJSONFile(t *testing.T, data any, path string) {
	t.Helper()
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Error marshaling to JSON: %v", err)
		return
	}

	err = os.WriteFile(path, jsonData, 0600)
	if err != nil {
		t.Errorf("Error writing JSON to file: %v", err)
		return
	}

	t.Cleanup(func() {
		if err := os.Remove(path); err != nil {
			t.Errorf("Error deleting file: %v", err)
		}
	})
}

func enableCompliancePolicy(projectID string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return fmt.Errorf("%w: invalid bin", err)
	}
	cmd := exec.Command(cliPath,
		backupsEntity,
		compliancePolicyEntity,
		"enable",
		"--projectId",
		projectID,
		"--authorizedEmail",
		authorizedEmail,
		"--authorizedUserFirstName",
		authorizedUserFirstName,
		"--authorizedUserLastName",
		authorizedUserLastName,
		"-o=json",
		"--force",
		"--watch", // avoiding HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
	)
	cmd.Env = os.Environ()
	output, outputErr := e2e.RunAndGetStdOut(cmd)
	if outputErr != nil {
		return fmt.Errorf("%w\n %s", outputErr, string(output))
	}
	return nil
}

func setupCompliancePolicy(t *testing.T, projectID string, compliancePolicy *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error) {
	t.Helper()
	compliancePolicy.SetAuthorizedEmail(authorizedEmail)
	compliancePolicy.SetAuthorizedUserFirstName(authorizedUserFirstName)
	compliancePolicy.SetAuthorizedUserLastName(authorizedUserLastName)
	compliancePolicy.SetProjectId(projectID)

	n, err := e2e.RandInt(255)
	if err != nil {
		return nil, fmt.Errorf("could not generate random int for setting up a compliance policy: %w", err)
	}
	randomPath := fmt.Sprintf("setup_compliance_policy_%d.json", n)
	createJSONFile(t, compliancePolicy, randomPath)

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return nil, fmt.Errorf("%w: invalid bin", err)
	}
	cmd := exec.Command(cliPath,
		backupsEntity,
		compliancePolicyEntity,
		"setup",
		"--projectId",
		projectID,
		"-o=json",
		"--force",
		"--file",
		randomPath,
		"--watch", // avoiding HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
	)

	cmd.Env = os.Environ()
	resp, outputErr := e2e.RunAndGetStdOut(cmd)
	if outputErr != nil {
		return nil, fmt.Errorf("%w\n %s", outputErr, string(resp))
	}
	trimmedResponse := removeDotsFromWatching(resp)

	var result atlasv2.DataProtectionSettings20231001
	if err := json.Unmarshal(trimmedResponse, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// If we watch a command in a testing environment,
// the output has some dots in the beginning (depending on how long it took to finish) that need to be removed.
func removeDotsFromWatching(consoleOutput []byte) []byte {
	return []byte(strings.TrimLeft(string(consoleOutput), "."))
}

func createStreamsInstance(t *testing.T, projectID, name string) (string, error) {
	t.Helper()

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	instanceName := fmt.Sprintf("e2e-%s-%v", name, n)

	cmd := exec.Command(
		cliPath,
		streamsEntity,
		"instance",
		"create",
		instanceName,
		"--projectId", projectID,
		"--provider", "AWS",
		"--region", "VIRGINIA_USA",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	return instanceName, nil
}

func deleteStreamsInstance(t *testing.T, projectID, name string) error {
	t.Helper()

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}

	cmd := exec.Command(
		cliPath,
		streamsEntity,
		"instance",
		"delete",
		name,
		"--projectId", projectID,
		"--force",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}

	return nil
}

func createStreamsConnection(t *testing.T, projectID, instanceName, name string) (string, error) {
	t.Helper()

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	connectionName := fmt.Sprintf("e2e-%s-%v", name, n)

	cmd := exec.Command(
		cliPath,
		streamsEntity,
		"connection",
		"create",
		connectionName,
		"--file", "data/create_streams_connection_test.json",
		"--instance", instanceName,
		"--projectId", projectID,
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	return connectionName, nil
}

func deleteStreamsConnection(t *testing.T, projectID, instanceName, name string) error {
	t.Helper()

	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}

	cmd := exec.Command(
		cliPath,
		streamsEntity,
		"connection",
		"delete",
		name,
		"--instance", instanceName,
		"--projectId", projectID,
		"--force",
	)
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}

	return nil
}
