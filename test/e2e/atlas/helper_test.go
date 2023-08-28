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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	eventsEntity                 = "events"
	clustersEntity               = "clusters"
	processesEntity              = "processes"
	metricsEntity                = "metrics"
	searchEntity                 = "search"
	indexEntity                  = "index"
	datalakeEntity               = "datalake"
	datafederationEntity         = "datafederation"
	datalakePipelineEntity       = "datalakepipeline"
	alertsEntity                 = "alerts"
	configEntity                 = "settings"
	dbusersEntity                = "dbusers"
	certsEntity                  = "certs"
	privateEndpointsEntity       = "privateendpoints"
	queryLimitsEntity            = "querylimits"
	onlineArchiveEntity          = "onlineArchives"
	projectEntity                = "project"
	orgEntity                    = "org"
	invitationsEntity            = "invitations"
	maintenanceEntity            = "maintenanceWindows"
	integrationsEntity           = "integrations"
	securityEntity               = "security"
	ldapEntity                   = "ldap"
	awsEntity                    = "aws"
	azureEntity                  = "azure"
	gcpEntity                    = "gcp"
	customDNSEntity              = "customDns"
	logsEntity                   = "logs"
	cloudProvidersEntity         = "cloudProviders"
	accessRolesEntity            = "accessRoles"
	customDBRoleEntity           = "customDbRoles"
	regionalModeEntity           = "regionalModes"
	serverlessEntity             = "serverless"
	liveMigrationsEntity         = "liveMigrations"
	auditingEntity               = "auditing"
	accessLogsEntity             = "accessLogs"
	accessListEntity             = "accessList"
	performanceAdvisorEntity     = "performanceAdvisor"
	slowQueryLogsEntity          = "slowQueryLogs"
	namespacesEntity             = "namespaces"
	networkingEntity             = "networking"
	networkPeeringEntity         = "peering"
	suggestedIndexesEntity       = "suggestedIndexes"
	slowOperationThresholdEntity = "slowOperationThreshold"
	tierM10                      = "M10"
	tierM2                       = "M2"
	diskSizeGB40                 = "40"
	diskSizeGB30                 = "30"
	projectsEntity               = "projects"
	settingsEntity               = "settings"
	backupsEntity                = "backups"
	exportsEntity                = "exports"
	bucketsEntity                = "buckets"
	jobsEntity                   = "jobs"
	snapshotsEntity              = "snapshots"
	restoresEntity               = "restores"
	compliancepolicyEntity       = "compliancepolicy"
	policiesEntity               = "policies"
	teamsEntity                  = "teams"
	setupEntity                  = "setup"
)

// AlertConfig constants.
const (
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	intervalMin   = 5
	delayMin      = 0
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
	e2eMDBVer            = "4.4"
	e2eSharedMDBVer      = "6.0"
)

// Backup compliance policy constants.
const (
	authorizedEmail = "firstname.lastname@example.com"
)

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
	if resp, err := create.CombinedOutput(); err != nil {
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
	if resp, err := watch.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error watching serverless instance %w: %s", err, string(resp))
	}
	return clusterName, nil
}

func deleteServerlessInstanceForProject(projectID, clusterName string) error {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return err
	}
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
	if resp, err := deleteCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error deleting serverless instance %w: %s", err, string(resp))
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
	// this command will fail with 404 once the cluster is deleted
	// we just need to wait for this to close the project
	_ = watchCmd.Run()
	return nil
}

func deployClusterForProject(projectID, tier string, enableBackup bool) (string, string, error) {
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
		"--mdbVersion", e2eMDBVer,
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
	if resp, err := create.CombinedOutput(); err != nil {
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
	if resp, err := watch.CombinedOutput(); err != nil {
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

func deleteClusterForProject(projectID, clusterName string) error {
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
	if resp, err := deleteCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error deleting cluster %w: %s", err, string(resp))
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
	// this command will fail with 404 once the cluster is deleted
	// we just need to wait for this to close the project
	_ = watchCmd.Run()
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
	if resp, err := deleteCmd.CombinedOutput(); err != nil {
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
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("error getting regions %w: %s", err, string(resp))
	}

	var cloudProviders mongodbatlas.CloudProviders
	err = json.Unmarshal(resp, &cloudProviders)
	if err != nil {
		return "", fmt.Errorf("error unmashaling response %w: %s", err, string(resp))
	}

	if len(cloudProviders.Results) == 0 || len(cloudProviders.Results[0].InstanceSizes) == 0 {
		return "", errNoRegions
	}

	return cloudProviders.Results[0].InstanceSizes[0].AvailableRegions[0].Name, nil
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
	services := thirdPartyIntegrations.Results
	for i := range services {
		iType := getIntegrationType(services[i])
		if iType == name {
			return true
		}
	}
	return false
}

func getIntegrationType(val atlasv2.ThridPartyIntegration) string {
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
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var users mongodbatlas.AtlasUsersResponse
	if err := json.Unmarshal(resp, &users); err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}
	if len(users.Results) == 0 {
		return "", fmt.Errorf("no users found")
	}

	return users.Results[0].Username, nil
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
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var team mongodbatlas.Team
	if err := json.Unmarshal(resp, &team); err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	return team.ID, nil
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
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project mongodbatlas.Project
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.ID, nil
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
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project mongodbatlas.Project
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.ID, nil
}

func deleteClustersForProject(t *testing.T, cliPath, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	t.Log(string(resp))
	require.NoError(t, err)
	var clusters mongodbatlas.AdvancedClustersResponse
	require.NoError(t, json.Unmarshal(resp, &clusters))
	for _, cluster := range clusters.Results {
		if cluster.StateName == "DELETING" {
			continue
		}
		assert.NoError(t, deleteClusterForProject(projectID, cluster.Name))
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
	resp, err := cmd.CombinedOutput()
	t.Log(string(resp))
	require.NoError(t, err)
	var pipelines []atlasv2.DataLakeIngestionPipeline
	require.NoError(t, json.Unmarshal(resp, &pipelines))
	for _, p := range pipelines {
		assert.NoError(t, deleteDatalakeForProject(cliPath, projectID, p.GetId()))
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
	resp, err := cmd.CombinedOutput()
	t.Log("available network peers", string(resp))
	require.NoError(t, err)
	var networkPeers []mongodbatlas.Peer
	err = json.Unmarshal(resp, &networkPeers)
	require.NoError(t, err)
	for _, peer := range networkPeers {
		peerID := peer.ID
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
		resp, err = cmd.CombinedOutput()
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

	clear := false
	for attempt := 0; attempt < 10; attempt++ {
		privateEndpoints = listPrivateEndpointsByProject(t, cliPath, projectID, provider)
		if len(privateEndpoints) == 0 {
			t.Logf("all %s private endpoints successfully deleted", provider)
			clear = true
			break
		}
		time.Sleep(sleep)
	}

	require.True(t, clear, "failed to clean all private endpoints")
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
	resp, err := cmd.CombinedOutput()
	t.Log(string(resp))
	require.NoError(t, err)
	var privateEndpoints []atlasv2.EndpointService
	err = json.Unmarshal(resp, &privateEndpoints)
	require.NoError(t, err)

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
	resp, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(resp))
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
	resp, err := cmd.CombinedOutput()
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
	resp, err := cmd.CombinedOutput()
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
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}

	return nil
}

func ensureCluster(t *testing.T, cluster *atlasv2.AdvancedClusterDescription, clusterName, version string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.GetName())
	a.Equal(version, cluster.GetMongoDBMajorVersion())
	a.Equal(diskSizeGB, cluster.GetDiskSizeGB())
	a.Equal(terminationProtection, cluster.GetTerminationProtectionEnabled())
}

func ensureSharedCluster(t *testing.T, cluster *mongodbatlas.Cluster, clusterName, tier string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.Name)
	a.Equal(e2eSharedMDBVer, cluster.MongoDBMajorVersion)
	if cluster.ProviderSettings != nil {
		a.Equal(tier, cluster.ProviderSettings.InstanceSizeName)
	}
	a.Equal(diskSizeGB, *cluster.DiskSizeGB)
	a.Equal(terminationProtection, *cluster.TerminationProtectionEnabled)
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
func createJSONFile(t *testing.T, data interface{}, path string) {
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
		compliancepolicyEntity,
		"enable",
		"--projectId",
		projectID,
		"--authorizedEmail",
		authorizedEmail,
		"-o=json",
		"--force",
		"--watch", // avoiding HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
	)
	cmd.Env = os.Environ()
	output, outputErr := cmd.CombinedOutput()
	if outputErr != nil {
		return fmt.Errorf("%w\n %s", outputErr, string(output))
	}
	return nil
}

func setupCompliancePolicy(t *testing.T, projectID string, compliancePolicy *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error) {
	t.Helper()
	compliancePolicy.SetAuthorizedEmail(authorizedEmail)
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
		compliancepolicyEntity,
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
	resp, outputErr := cmd.CombinedOutput()
	if outputErr != nil {
		return nil, fmt.Errorf("%w\n %s", outputErr, string(resp))
	}
	trimmedResponse := removeDotsFromWatching(resp)

	var result atlasv2.DataProtectionSettings
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
