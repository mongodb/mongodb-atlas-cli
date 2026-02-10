// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

var (
	errNoRegions    = errors.New("no regions available")
	errInvalidIndex = errors.New("invalid index")
)

const (
	// entities.
	eventsEntity                  = "events"
	clustersEntity                = "clusters"
	processesEntity               = "processes"
	metricsEntity                 = "metrics"
	searchEntity                  = "search"
	indexEntity                   = "index"
	nodesEntity                   = "nodes"
	datafederationEntity          = "datafederation"
	alertsEntity                  = "alerts"
	configEntity                  = "config"
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
	authEntity                    = "auth"
	streamsEntity                 = "streams"
	apiKeysEntity                 = "apikeys"
	apiKeyAccessListEntity        = "accessLists"
	usersEntity                   = "users"
	apiEntity                     = "api"
	serviceAccountsEntity         = "serviceAccounts"

	deletingState = "DELETING"

	maxRetryAttempts    = 10
	sleepTimeInSeconds  = 30
	clusterWatchTimeout = 30 * time.Minute // 30 minutes timeout for cluster watch operations

	// CLI Plugins System constants.
	examplePluginRepository = "mongodb/atlas-cli-plugin-example"
	examplePluginName       = "atlas-cli-plugin-example"

	// Cluster settings.
	e2eClusterTier       = "M10"
	e2eGovClusterTier    = "M20"
	e2eSharedClusterTier = "M2"
	e2eClusterProvider   = "AWS" // e2eClusterProvider preferred provider for e2e testing.

	// serviceAccount API constants.
	serviceAccountAPIVersion = "2024-08-05"
	secretExpiresAfterHours  = 8
)

// Backup compliance policy constants.
const (
	authorizedUserFirstName = "firstname"
	authorizedUserLastName  = "lastname"
	authorizedEmail         = "firstname.lastname@example.com"
)

func RandInt(maximum int64) (*big.Int, error) {
	return rand.Int(rand.Reader, big.NewInt(maximum))
}

// DeleteProjectWithRetry deletes a project with a retry backoff strategy.
func DeleteProjectWithRetry(t *testing.T, projectID string) {
	t.Helper()
	deleted := false
	backoff := 1
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		e := deleteProject(projectID)
		if e == nil || strings.Contains(e.Error(), "GROUP_NOT_FOUND") {
			t.Logf("project %q successfully deleted", projectID)
			deleted = true
			break
		}

		t.Logf("%d/%d attempts - trying again in %d seconds: unexpected error while deleting the project %q: %v", attempts, maxRetryAttempts, backoff, projectID, e)
		time.Sleep(time.Duration(backoff) * time.Second)
		backoff *= 2
	}
	if !deleted {
		t.Errorf("we could not delete the project %q", projectID)
	}
}

func RunAndGetStdOut(cmd *exec.Cmd) ([]byte, error) {
	cmd.Stderr = os.Stderr

	resp, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("%s (%w)", string(resp), err)
	}

	return resp, nil
}

func RunAndGetStdOutAndErr(cmd *exec.Cmd) ([]byte, error) {
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("%s (%w)", string(resp), err)
	}

	return resp, nil
}

func RunAndGetSeparateStdOutAndErr(cmd *exec.Cmd) ([]byte, []byte, error) {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()

	return stdOut.Bytes(), stdErr.Bytes(), err
}

func SplitOutput(cmd *exec.Cmd) (string, string, error) {
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	err := cmd.Run()
	return o.String(), e.String(), err
}

func deployClusterForProject(projectID, clusterName, tier, mDBVersion string, enableBackup bool) (string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", err
	}
	region, err := NewAvailableRegion(projectID, tier, e2eClusterProvider)
	if err != nil {
		return "", fmt.Errorf("failed to get available region for project %s, tier %s, provider %s: %w", projectID, tier, e2eClusterProvider, err)
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
		"-P",
		ProfileName(),
	}
	if enableBackup {
		args = append(args, "--backup")
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := RunAndGetStdOut(create); err != nil {
		return "", fmt.Errorf("error creating cluster %s in project %s: %w: %s", clusterName, projectID, err, string(resp))
	}

	watchArgs := []string{
		clustersEntity,
		"watch",
		clusterName,
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watch := exec.Command(cliPath, watchArgs...)
	watch.Env = os.Environ()
	if resp, err := RunAndGetStdOut(watch); err != nil {
		return "", fmt.Errorf("error watching cluster %s in project %s: %w: %s", clusterName, projectID, err, string(resp))
	}
	return region, nil
}

func E2eTier() string {
	tier := e2eClusterTier
	if IsGov() {
		tier = e2eGovClusterTier
	}
	return tier
}

func internalDeleteClusterForProject(projectID, clusterName string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	args := []string{
		clustersEntity,
		"delete",
		clusterName,
		"--force",
		"--watch",
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	deleteCmd := exec.Command(cliPath, args...)
	deleteCmd.Env = os.Environ()
	if resp, err := RunAndGetStdOutAndErr(deleteCmd); err != nil {
		return fmt.Errorf("error deleting cluster %s in project %s: %w: %s", clusterName, projectID, err, string(resp))
	}
	return nil
}

func WatchCluster(projectID, clusterName string) error {
	return WatchClusterWithTimeout(projectID, clusterName, clusterWatchTimeout)
}

func getClusterState(projectID, clusterName string) (string, string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", "", err
	}
	args := []string{
		clustersEntity,
		"describe",
		clusterName,
		"-o=json",
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", "", fmt.Errorf("failed to get cluster state: %w: %s", err, string(resp))
	}

	var cluster atlasClustersPinned.AdvancedClusterDescription
	if err := json.Unmarshal(resp, &cluster); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal cluster response: %w: %s", err, string(resp))
	}

	stateName := cluster.GetStateName()
	clusterID := cluster.GetId()
	return stateName, clusterID, nil
}

func WatchClusterWithTimeout(projectID, clusterName string, timeout time.Duration) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	watchArgs := []string{
		clustersEntity,
		"watch",
		clusterName,
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}

	// Create a context with timeout to prevent indefinite hanging
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	watchCmd := exec.CommandContext(ctx, cliPath, watchArgs...)
	watchCmd.Env = os.Environ()

	if resp, err := RunAndGetStdOut(watchCmd); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			// Get the actual cluster state to provide better debugging info
			stateName, clusterID, stateErr := getClusterState(projectID, clusterName)
			var stateInfo string
			if stateErr == nil {
				stateInfo = fmt.Sprintf("current state: %s, cluster ID: %s", stateName, clusterID)
			} else {
				stateInfo = fmt.Sprintf("failed to get cluster state: %v", stateErr)
			}
			return fmt.Errorf("timeout waiting for cluster %s in project %s after %v. %s. Cluster may be stuck in DELETING state or deletion is taking longer than expected", clusterName, projectID, timeout, stateInfo)
		}
		return fmt.Errorf("error waiting for cluster %s in project %s: %w: %s", clusterName, projectID, err, string(resp))
	}
	return nil
}

func removeTerminationProtectionFromCluster(projectID, clusterName string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	args := []string{
		clustersEntity,
		"update",
		clusterName,
		"--disableTerminationProtection",
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	updateCmd := exec.Command(cliPath, args...)
	updateCmd.Env = os.Environ()
	if resp, err := RunAndGetStdOut(updateCmd); err != nil {
		return fmt.Errorf("error removing termination protection from cluster %s in project %s: %w: %s", clusterName, projectID, err, string(resp))
	}

	return WatchCluster(projectID, clusterName)
}

func DeleteClusterForProject(projectID, clusterName string) error {
	if err := internalDeleteClusterForProject(projectID, clusterName); err != nil {
		if !strings.Contains(err.Error(), "CANNOT_TERMINATE_CLUSTER_WHEN_TERMINATION_PROTECTION_ENABLED") {
			return fmt.Errorf("failed to delete cluster %s in project %s: %w", clusterName, projectID, err)
		}

		// Try to remove termination protection and retry
		if err := removeTerminationProtectionFromCluster(projectID, clusterName); err != nil {
			return fmt.Errorf("failed to remove termination protection from cluster %s in project %s: %w", clusterName, projectID, err)
		}
		if err := internalDeleteClusterForProject(projectID, clusterName); err != nil {
			return fmt.Errorf("failed to delete cluster %s in project %s after removing termination protection: %w", clusterName, projectID, err)
		}
	}

	return nil
}

func NewAvailableRegion(projectID, tier, provider string) (string, error) {
	cliPath, err := AtlasCLIBin()
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
		"-P",
		ProfileName(),
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)

	if err != nil {
		return "", fmt.Errorf("error getting regions for project %s, tier %s, provider %s: %w: %s", projectID, tier, provider, err, string(resp))
	}

	var cloudProviders atlasv2.PaginatedApiAtlasProviderRegions
	err = json.Unmarshal(resp, &cloudProviders)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling regions response for project %s, tier %s, provider %s: %w: %s", projectID, tier, provider, err, string(resp))
	}

	if cloudProviders.GetTotalCount() == 0 || len(cloudProviders.GetResults()[0].GetInstanceSizes()) == 0 {
		return "", fmt.Errorf("%w: no regions available for project %s, tier %s, provider %s", errNoRegions, projectID, tier, provider)
	}

	return cloudProviders.GetResults()[0].GetInstanceSizes()[0].GetAvailableRegions()[0].GetName(), nil
}

func RandClusterName() (string, error) {
	return RandClusterNameWithPrefix("cluster")
}

func RandClusterNameWithPrefix(prefix string) (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
	if err != nil {
		return "", err
	}

	clusterName := fmt.Sprintf("%s-%d", prefix, n)
	if revision, ok := os.LookupEnv("revision"); ok {
		clusterName = fmt.Sprintf("%s-%v-%s", prefix, n, revision)
	}

	if len(clusterName) > 23 { //nolint:mnd // internal validation of cluster name
		clusterName = clusterName[:23]
	}

	if clusterName[len(clusterName)-1] == '-' {
		clusterName = clusterName[:len(clusterName)-1]
	}

	return clusterName, nil
}

func RandIdentityProviderName() (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("idp-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("idp-%v", n), nil
}

func RandTeamName() (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("team-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("team-%v", n), nil
}

func RandProjectName() (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
	if err != nil {
		return "", err
	}

	projectName := fmt.Sprintf("e2e-%v", n)
	if revision, ok := os.LookupEnv("revision"); ok {
		projectName = fmt.Sprintf("%v-%s", n, revision)
	}

	if len(projectName) > 23 { //nolint:mnd // internal validation of project name
		projectName = projectName[:23]
	}

	if projectName[len(projectName)-1] == '-' {
		projectName = projectName[:len(projectName)-1]
	}

	return projectName, nil
}

func RandUsername() (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
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
	if len(prefixedName) > 64 { //nolint:mnd // internal validation of team name
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
	if len(prefixedName) > 64 { //nolint:mnd // internal validation of project name
		return prefixedName[:64], nil
	}
	return prefixedName, nil
}

func RandEntityWithRevision(entity string) (string, error) {
	n, err := RandInt(1000) //nolint:mnd // RandInt is used to generate a random number
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
	atlasURL := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL")
	if atlasURL == "" {
		profile, err := ProfileData()
		if err != nil {
			return "", err
		}
		atlasURL = profile["ops_manager_url"]
	}
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

func TempConfigFolder(t *testing.T) string {
	t.Helper()

	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	t.Setenv("home", tmpDir)
	t.Setenv("XDG_CONFIG_HOME", tmpDir)
	t.Setenv("AppData", tmpDir)

	// Silence storage warning for e2e tests to reduce noise in the output.
	t.Setenv("MONGODB_ATLAS_SILENCE_STORAGE_WARNING", "true")

	dir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dir = filepath.Join(dir, "atlascli")

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return dir
}

func InitKeychain(t *testing.T) error {
	t.Helper()

	if runtime.GOOS == "darwin" {
		return InitKeychainMac(t)
	}

	return fmt.Errorf("keychain initialization not supported on %s", runtime.GOOS)
}

func InitKeychainMac(t *testing.T) error {
	t.Helper()

	// Run the following command to initialize the keychain
	//
	// Create the preferences directory, expected by the security command to exist
	// HOME=dir mkdir -p $dir/Library/Preferences
	//
	// Create the keychain:
	// HOME=dir /usr/bin/security create-keychain -p "" default.keychain-db
	//
	// Add the keychain to the search list:
	// HOME=dir /usr/bin/security list-keychains -d user -s default.keychain-db
	//
	// Set the default keychain:
	// HOME=dir /usr/bin/security default-keychain -s default.keychain-db
	//
	// Unlock the keychain:
	// HOME=dir /usr/bin/security unlock-keychain -p "" default.keychain-db
	//

	home := os.Getenv("HOME")

	if err := os.MkdirAll(filepath.Join(home, "Library", "Preferences"), os.ModePerm); err != nil {
		return fmt.Errorf("error creating preferences directory: %w", err)
	}

	createCmd := exec.Command("/usr/bin/security", "create-keychain", "-p", "", "default.keychain-db")
	createCmd.Env = os.Environ()
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("error creating keychain: %w", err)
	}

	listCmd := exec.Command("/usr/bin/security", "list-keychains", "-d", "user", "-s", "default.keychain-db")
	listCmd.Env = os.Environ()
	if err := listCmd.Run(); err != nil {
		return fmt.Errorf("error listing keychains: %w", err)
	}

	defaultCmd := exec.Command("/usr/bin/security", "default-keychain", "-s", "default.keychain-db")
	defaultCmd.Env = os.Environ()
	if err := defaultCmd.Run(); err != nil {
		return fmt.Errorf("error setting default keychain: %w", err)
	}

	unlockCmd := exec.Command("/usr/bin/security", "unlock-keychain", "-p", "", "default.keychain-db")
	unlockCmd.Env = os.Environ()
	if err := unlockCmd.Run(); err != nil {
		return fmt.Errorf("error unlocking keychain: %w", err)
	}

	return nil
}

func createProject(projectName string) (string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin", err)
	}
	args := []string{
		projectEntity,
		"create",
		projectName,
		"-o=json",
		"-P",
		ProfileName(),
	}
	if IsGov() {
		args = append(args, "--govCloudRegionsOnly")
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project atlasv2.Group
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.GetId(), nil
}

func listClustersForProject(t *testing.T, cliPath, projectID string) atlasClustersPinned.PaginatedAdvancedClusterDescription {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		clustersEntity,
		"list",
		"--projectId", projectID,
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var clusters atlasClustersPinned.PaginatedAdvancedClusterDescription
	require.NoError(t, json.Unmarshal(resp, &clusters))
	return clusters
}

func deleteAllClustersForProject(t *testing.T, cliPath, projectID string) {
	t.Helper()
	clusters := listClustersForProject(t, cliPath, projectID)
	if len(clusters.GetResults()) == 0 {
		t.Logf("no clusters found in project %s", projectID)
		return
	}

	t.Logf("found %d clusters to delete in project %s", len(clusters.GetResults()), projectID)
	var failedClusters []string
	for _, cluster := range clusters.GetResults() {
		func(clusterName, state string) {
			t.Run("delete cluster "+clusterName, func(t *testing.T) {
				t.Parallel()
				if state == deletingState {
					t.Logf("cluster %s is already in DELETING state, waiting for deletion to complete in project %s (timeout: %v)", clusterName, projectID, clusterWatchTimeout)
					err := WatchClusterWithTimeout(projectID, clusterName, clusterWatchTimeout)
					if err != nil {
						// Try to get current state for better debugging
						currentState, clusterID, stateErr := getClusterState(projectID, clusterName)
						if stateErr == nil {
							t.Errorf("failed to watch cluster %s (ID: %s) deletion in project %s: %v. Current state: %s", clusterName, clusterID, projectID, err, currentState)
						} else {
							t.Errorf("failed to watch cluster %s deletion in project %s: %v. Could not get current state: %v", clusterName, projectID, err, stateErr)
						}
						failedClusters = append(failedClusters, clusterName)
					} else {
						t.Logf("cluster %s successfully deleted in project %s", clusterName, projectID)
					}
					return
				}
				t.Logf("deleting cluster %s (state: %s) in project %s", clusterName, state, projectID)
				err := DeleteClusterForProject(projectID, clusterName)
				if err != nil {
					// Try to get current state for better debugging
					currentState, clusterID, stateErr := getClusterState(projectID, clusterName)
					if stateErr == nil {
						t.Errorf("failed to delete cluster %s (ID: %s) in project %s: %v. Current state: %s", clusterName, clusterID, projectID, err, currentState)
					} else {
						t.Errorf("failed to delete cluster %s in project %s: %v. Could not get current state: %v", clusterName, projectID, err, stateErr)
					}
					failedClusters = append(failedClusters, clusterName)
				} else {
					t.Logf("cluster %s successfully deleted in project %s", clusterName, projectID)
				}
			})
		}(cluster.GetName(), cluster.GetStateName())
	}
	if len(failedClusters) > 0 {
		t.Errorf("failed to delete %d clusters in project %s: %v", len(failedClusters), projectID, failedClusters)
	}
}

func deleteAllNetworkPeers(t *testing.T, cliPath, projectID, provider string) {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		networkingEntity,
		networkPeeringEntity,
		"list",
		"--provider",
		provider,
		"--projectId",
		projectID,
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Logf("available network peers for provider %s in project %s: %s", provider, projectID, string(resp))
	require.NoError(t, err, string(resp))
	var networkPeers []atlasv2.BaseNetworkPeeringConnectionSettings
	err = json.Unmarshal(resp, &networkPeers)
	require.NoError(t, err)

	if len(networkPeers) == 0 {
		t.Logf("no network peers found for provider %s in project %s", provider, projectID)
		return
	}

	t.Logf("found %d network peers to delete for provider %s in project %s", len(networkPeers), provider, projectID)
	var failedDeletions []string
	for _, peer := range networkPeers {
		peerID := peer.GetId()
		containerID := peer.GetContainerId()
		if containerID == "" {
			containerID = "unknown"
		}
		t.Logf("deleting network peer %s (container: %s) for provider %s in project %s", peerID, containerID, provider, projectID)
		cmd = exec.Command(cliPath,
			networkingEntity,
			networkPeeringEntity,
			"delete",
			peerID,
			"--projectId",
			projectID,
			"--force",
			"-P",
			ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err = RunAndGetStdOut(cmd)
		if err != nil {
			t.Errorf("failed to delete network peer %s (container: %s) for provider %s in project %s: %v, output: %s", peerID, containerID, provider, projectID, err, string(resp))
			failedDeletions = append(failedDeletions, peerID)
		} else {
			t.Logf("successfully deleted network peer %s for provider %s in project %s", peerID, provider, projectID)
		}
	}
	if len(failedDeletions) > 0 {
		t.Errorf("failed to delete %d network peers for provider %s in project %s: %v", len(failedDeletions), provider, projectID, failedDeletions)
	}
}

const sleep = 10 * time.Second

func deleteAllPrivateEndpoints(t *testing.T, cliPath, projectID, provider string) {
	t.Helper()

	privateEndpoints := listPrivateEndpointsByProject(t, cliPath, projectID, provider)
	if len(privateEndpoints) == 0 {
		t.Logf("no private endpoints found for provider %s in project %s", provider, projectID)
		return
	}

	t.Logf("found %d private endpoints to delete for provider %s in project %s", len(privateEndpoints), provider, projectID)
	for i, endpoint := range privateEndpoints {
		endpointID := endpoint.GetId()
		endpointName := endpoint.GetEndpointServiceName()
		t.Logf("deleting private endpoint %d/%d: ID=%s, Name=%s for provider %s in project %s", i+1, len(privateEndpoints), endpointID, endpointName, provider, projectID)
		deletePrivateEndpoint(t, cliPath, projectID, provider, endpointID)
	}

	done := false
	var remainingEndpoints []string
	for attempt := range 10 {
		privateEndpoints = listPrivateEndpointsByProject(t, cliPath, projectID, provider)
		if len(privateEndpoints) == 0 {
			t.Logf("all %s private endpoints successfully deleted after %d attempts", provider, attempt+1)
			done = true
			break
		}
		remainingEndpoints = make([]string, 0, len(privateEndpoints))
		for _, ep := range privateEndpoints {
			remainingEndpoints = append(remainingEndpoints, fmt.Sprintf("ID=%s, Name=%s, Status=%s", ep.GetId(), ep.GetEndpointServiceName(), ep.GetStatus()))
		}
		t.Logf("attempt %d/10: %d %s private endpoints still remaining in project %s: %v", attempt+1, len(privateEndpoints), provider, projectID, remainingEndpoints)
		time.Sleep(sleep)
	}

	if !done {
		t.Errorf("failed to clean all %s private endpoints in project %s after 10 attempts. Remaining endpoints: %v", provider, projectID, remainingEndpoints)
	}
}

func deleteAllStreams(t *testing.T, cliPath, projectID string) {
	t.Helper()

	streams := listStreamsByProject(t, cliPath, projectID)
	if streams.GetTotalCount() == 0 {
		t.Logf("no streams found in project %s", projectID)
		return
	}

	t.Logf("found %d streams to delete in project %s", streams.GetTotalCount(), projectID)
	var streamNames []string
	for i, stream := range streams.GetResults() {
		streamName := *stream.Name
		streamNames = append(streamNames, streamName)
		t.Logf("deleting stream %d/%d: %s in project %s", i+1, streams.GetTotalCount(), streamName, projectID)
		deleteStream(t, cliPath, projectID, streamName)
	}

	done := false
	var remainingStreams []string
	for attempt := range 10 {
		streams = listStreamsByProject(t, cliPath, projectID)
		if streams.GetTotalCount() == 0 {
			t.Logf("all streams successfully deleted after %d attempts in project %s", attempt+1, projectID)
			done = true
			break
		}
		remainingStreams = make([]string, 0, streams.GetTotalCount())
		for _, stream := range streams.GetResults() {
			remainingStreams = append(remainingStreams, *stream.Name)
		}
		t.Logf("attempt %d/10: %d streams still remaining in project %s: %v", attempt+1, streams.GetTotalCount(), projectID, remainingStreams)
		time.Sleep(sleep)
	}

	if !done {
		t.Errorf("failed to clean all streams in project %s after 10 attempts. Attempted to delete: %v. Remaining streams: %v", projectID, streamNames, remainingStreams)
	}
}

func listStreamsByProject(t *testing.T, cliPath, projectID string) *atlasv2.PaginatedApiStreamsTenant {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		streamsEntity,
		"instance",
		"list",
		"--projectId",
		projectID,
		"-o=json",
		"-P",
		ProfileName(),
	)

	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var streams *atlasv2.PaginatedApiStreamsTenant
	require.NoError(t, json.Unmarshal(resp, &streams))

	return streams
}

func deleteStream(t *testing.T, cliPath, projectID, streamID string) {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		streamsEntity,
		"instance",
		"delete",
		"--force",
		streamID,
		"--projectId",
		projectID,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		t.Logf("error deleting stream %s in project %s: %v, output: %s", streamID, projectID, err, string(resp))
	}
	require.NoError(t, err, "failed to delete stream %s in project %s: %s", streamID, projectID, string(resp))
}

func listPrivateEndpointsByProject(t *testing.T, cliPath, projectID, provider string) []atlasv2.EndpointService {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		privateEndpointsEntity,
		provider,
		"list",
		"--projectId",
		projectID,
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Log(string(resp))
	require.NoError(t, err, string(resp))
	var privateEndpoints []atlasv2.EndpointService
	require.NoError(t, json.Unmarshal(resp, &privateEndpoints))

	return privateEndpoints
}

func deletePrivateEndpoint(t *testing.T, cliPath, projectID, provider, endpointID string) {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		privateEndpointsEntity,
		provider,
		"delete",
		endpointID,
		"--projectId",
		projectID,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		t.Logf("error deleting private endpoint %s for provider %s in project %s: %v, output: %s", endpointID, provider, projectID, err, string(resp))
	}
	require.NoError(t, err, "failed to delete private endpoint %s for provider %s in project %s: %s", endpointID, provider, projectID, string(resp))
}

func DeleteTeam(teamID string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		teamsEntity,
		"delete",
		teamID,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}
	return nil
}

func deleteProject(projectID string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		projectEntity,
		"delete",
		projectID,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOutAndErr(cmd)
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}
	return nil
}

func listDataFederationsByProject(t *testing.T, cliPath, projectID string) []atlasv2.DataLakeTenant {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		datafederationEntity,
		"list",
		"--projectId", projectID,
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Log("available datafederations", string(resp))
	require.NoError(t, err, string(resp))

	var dataFederations []atlasv2.DataLakeTenant
	err = json.Unmarshal(resp, &dataFederations)
	require.NoError(t, err)

	return dataFederations
}

func deleteAllDataFederations(t *testing.T, cliPath, projectID string) {
	t.Helper()

	dataFederations := listDataFederationsByProject(t, cliPath, projectID)
	if len(dataFederations) == 0 {
		t.Logf("no data federations found in project %s", projectID)
		return
	}

	t.Logf("found %d data federations to delete in project %s", len(dataFederations), projectID)
	var federationNames []string
	for i, federation := range dataFederations {
		fedName := federation.GetName()
		federationNames = append(federationNames, fedName)
		t.Logf("deleting data federation %d/%d: %s in project %s", i+1, len(dataFederations), fedName, projectID)
		deleteDataFederationForProject(t, cliPath, projectID, fedName)
	}

	// Re-check to see if any remain after deletion attempts
	remainingFederations := listDataFederationsByProject(t, cliPath, projectID)
	if len(remainingFederations) > 0 {
		var remainingNames []string
		for _, fed := range remainingFederations {
			remainingNames = append(remainingNames, fed.GetName())
		}
		t.Errorf("failed to delete %d data federations in project %s. Attempted to delete: %v. Remaining: %v", len(remainingFederations), projectID, federationNames, remainingNames)
	} else {
		t.Logf("all data federations successfully deleted in project %s", projectID)
	}
}

func deleteDataFederationForProject(t *testing.T, cliPath, projectID, dataFedName string) {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		datafederationEntity,
		"delete",
		dataFedName,
		"--projectId", projectID,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func EnsureCluster(t *testing.T, cluster *atlasClustersPinned.AdvancedClusterDescription, clusterName, version string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.GetName())
	a.Equal(version, cluster.GetMongoDBMajorVersion())
	a.InDelta(diskSizeGB, cluster.GetDiskSizeGB(), 0.01) //nolint:mnd // ensure disk size is within 0.01 of expected value
	a.Equal(terminationProtection, cluster.GetTerminationProtectionEnabled())
}

func EnsureClusterLatest(t *testing.T, cluster *atlasv2.ClusterDescription20240805, clusterName, version string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.GetName())
	a.Equal(version, cluster.GetMongoDBMajorVersion())
	a.Equal(terminationProtection, cluster.GetTerminationProtectionEnabled())
	for _, repSpecs := range cluster.GetReplicationSpecs() {
		for _, config := range repSpecs.GetRegionConfigs() {
			electableSpecs := config.GetElectableSpecs()
			diskSize := electableSpecs.GetDiskSizeGB()
			a.InDelta(diskSizeGB, diskSize, 0.01) //nolint:mnd // ensure disk size is within 0.01 of expected value
		}
	}
}

func EnsureFlexCluster(t *testing.T, cluster *atlasv2.FlexClusterDescription20241113, clusterName string, diskSizeGB float64, terminationProtection bool) {
	t.Helper()
	a := assert.New(t)
	setting, ok := cluster.GetProviderSettingsOk()

	a.True(ok)
	a.Equal(clusterName, cluster.GetName())
	a.InDelta(diskSizeGB, setting.GetDiskSizeGB(), 0.01) //nolint:mnd // ensure disk size is within 0.01 of expected value
	a.Equal(terminationProtection, cluster.GetTerminationProtectionEnabled())
}

// CreateJSONFile creates a new JSON file at the specified path with the specified data
// and also registers its deletion on test cleanup.
func CreateJSONFile(t *testing.T, data any, path string) {
	t.Helper()
	jsonData, err := json.Marshal(data)
	require.NoError(t, err)
	const permission = 0600
	require.NoError(t, os.WriteFile(path, jsonData, permission))

	t.Cleanup(func() {
		require.NoError(t, os.Remove(path))
	})
}

func EnableCompliancePolicy(projectID string) error {
	cliPath, err := AtlasCLIBin()
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
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	output, outputErr := RunAndGetStdOut(cmd)
	if outputErr != nil {
		return fmt.Errorf("%w\n %s", outputErr, string(output))
	}
	return nil
}

func SetupCompliancePolicy(t *testing.T, projectID string, compliancePolicy *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error) {
	t.Helper()
	compliancePolicy.SetAuthorizedEmail(authorizedEmail)
	compliancePolicy.SetAuthorizedUserFirstName(authorizedUserFirstName)
	compliancePolicy.SetAuthorizedUserLastName(authorizedUserLastName)
	compliancePolicy.SetProjectId(projectID)

	n, err := RandInt(255) //nolint:mnd // RandInt is used to generate a random number
	if err != nil {
		return nil, fmt.Errorf("could not generate random int for setting up a compliance policy: %w", err)
	}
	randomPath := fmt.Sprintf("setup_compliance_policy_%d.json", n)
	CreateJSONFile(t, compliancePolicy, randomPath)

	cliPath, err := AtlasCLIBin()
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
		"-P",
		ProfileName(),
	)

	cmd.Env = os.Environ()
	resp, outputErr := RunAndGetStdOut(cmd)
	if outputErr != nil {
		return nil, fmt.Errorf("%w\n %s", outputErr, string(resp))
	}
	trimmedResponse := RemoveDotsFromWatching(resp)

	var result atlasv2.DataProtectionSettings20231001
	if err := json.Unmarshal(trimmedResponse, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// If we watch a command in a testing environment,
// the output has some dots in the beginning (depending on how long it took to finish) that need to be removed.
func RemoveDotsFromWatching(consoleOutput []byte) []byte {
	return []byte(strings.TrimLeft(string(consoleOutput), "."))
}

func DeleteAllPlugins(t *testing.T) {
	t.Helper()
	defaultPluginDir, err := plugin.GetDefaultPluginDirectory()
	require.NoError(t, err)

	err = os.RemoveAll(defaultPluginDir)
	require.NoError(t, err)
}

func InstallExamplePlugin(t *testing.T, cliPath string, version string) {
	t.Helper()
	// this is a test
	// #nosec G204
	cmd := exec.Command(cliPath,
		"plugin",
		"install",
		fmt.Sprintf("%s@%s", examplePluginRepository, version))
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func getFedSettingsID(t *testing.T, cliPath string) string {
	t.Helper()
	args := []string{federatedAuthenticationEntity,
		federationSettingsEntity,
		"describe",
		"-o=json",
		"-P",
		ProfileName(),
	}
	if orgID, set := os.LookupEnv("MONGODB_ATLAS_ORG_ID"); set {
		args = append(args, "--orgId", orgID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	var settings *atlasv2.OrgFederationSettings
	require.NoError(t, json.Unmarshal(resp, &settings))
	require.NotNil(t, settings.Id)

	return *settings.Id
}

func listIDPs(t *testing.T, cliPath string, fedSettingsID string) *atlasv2.PaginatedFederationIdentityProvider {
	t.Helper()
	cmd := exec.Command(cliPath, "federatedAuthentication", "federationSettings", "identityProvider", "list", "--federationSettingsId", fedSettingsID, "-o", "json", "-P", ProfileName()) //nolint:gosec // needed for e2e tests
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	var idps *atlasv2.PaginatedFederationIdentityProvider
	require.NoError(t, json.Unmarshal(resp, &idps))
	return idps
}

func deleteIDP(t *testing.T, cliPath string, id string, fedSettingsID string) {
	t.Helper()
	cmd := exec.Command(cliPath, federatedAuthenticationEntity, federationSettingsEntity, "identityProvider", "delete", id, "--federationSettingsId", fedSettingsID, "--force", "-P", ProfileName()) //nolint:gosec // needed for e2e tests
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func deleteAllIDPs(t *testing.T, cliPath string) {
	t.Helper()
	fedSettingsID := getFedSettingsID(t, cliPath)
	idps := listIDPs(t, cliPath, fedSettingsID)
	for _, idp := range *idps.Results {
		deleteIDP(t, cliPath, idp.Id, fedSettingsID)
	}
}

func CreateTeam(teamName string) (string, error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", err
	}
	username, _, err := OrgNUser(0)

	if err != nil {
		return "", err
	}
	cmd := exec.Command(cliPath,
		teamsEntity,
		"create",
		teamName,
		"--username",
		username,
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	var team atlasv2.Team
	if err := json.Unmarshal(resp, &team); err != nil {
		return "", err
	}

	return team.GetId(), nil
}

// OrgNUser returns the user at the position userIndex.
// We need to pass the userIndex because the command iam teams users add would not work
// if the user is already in the team.
func OrgNUser(n int) (username, userID string, err error) {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return "", "", err
	}
	cmd := exec.Command(cliPath,
		orgEntity,
		usersEntity,
		"list",
		"--limit",
		strconv.Itoa(n+1),
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", "", fmt.Errorf("error loading org users: %w (%s)", err, string(resp))
	}

	var users atlasv2.PaginatedOrgUser
	if err := json.Unmarshal(resp, &users); err != nil {
		return "", "", err
	}

	if len(users.GetResults()) <= n {
		return "", "", fmt.Errorf("%w: %d for %d users", errInvalidIndex, n, len(users.GetResults()))
	}

	return users.GetResults()[n].Username, users.GetResults()[n].GetId(), nil
}

func deleteKeys(t *testing.T, cliPath string, toDelete map[string]struct{}) {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		orgEntity,
		"apiKeys",
		"ls",
		"-o=json",
		"-P",
		ProfileName(),
	)

	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	var keys atlasv2.PaginatedApiApiUser
	err = json.Unmarshal(resp, &keys)
	require.NoError(t, err)

	uniqueKeysToDelete := make(map[string]struct{})
	for _, key := range keys.GetResults() {
		keyID := key.GetId()
		desc := key.GetDesc()

		if _, ok := toDelete[desc]; ok {
			uniqueKeysToDelete[keyID] = struct{}{}
		}
	}

	for keyID := range uniqueKeysToDelete {
		errs := []error{}
		t.Logf("Deleting key with ID: %s", keyID)
		cmd = exec.Command(cliPath,
			orgEntity,
			"apiKeys",
			"rm",
			keyID,
			"--force",
			"-P",
			ProfileName(),
		)
		cmd.Env = os.Environ()
		_, err = RunAndGetStdOutAndErr(cmd)
		if err != nil && !strings.Contains(err.Error(), "API_KEY_NOT_FOUND") {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			t.Errorf("unexpected errors while deleting keys: %v", errs)
		}
	}
}

func deleteOrgInvitations(t *testing.T, cliPath string) {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		orgEntity,
		invitationsEntity,
		"ls",
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Logf("%s\n", resp)
	require.NoError(t, err, string(resp))
	var invitations []atlasv2.OrganizationInvitation
	require.NoError(t, json.Unmarshal(resp, &invitations), string(resp))
	for _, i := range invitations {
		deleteOrgInvitation(t, cliPath, *i.Id)
	}
}

func deleteOrgTeams(t *testing.T, cliPath string) {
	t.Helper()

	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		teamsEntity,
		"ls",
		"-o=json",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	t.Logf("%s\n", resp)
	require.NoError(t, err, string(resp))
	var teams atlasv2.PaginatedTeam
	require.NoError(t, json.Unmarshal(resp, &teams), string(resp))
	for _, team := range teams.GetResults() {
		assert.NoError(t, DeleteTeam(team.GetId()))
	}
}

func deleteOrgInvitation(t *testing.T, cliPath string, id string) {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		orgEntity,
		invitationsEntity,
		"delete",
		id,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func DeleteOrgAPIKey(id string) error {
	cliPath, err := AtlasCLIBin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		orgEntity,
		apiKeysEntity,
		"rm",
		id,
		"--force",
		"-P",
		ProfileName(),
	)
	cmd.Env = os.Environ()
	return cmd.Run()
}

// createOrgServiceAccount creates a new organization service account.
func createOrgServiceAccount(cliPath, name string) (string, string, error) {
	payload := map[string]any{
		"description":             "Test service account",
		"name":                    name,
		"roles":                   []string{"ORG_OWNER"},
		"secretExpiresAfterHours": secretExpiresAfterHours,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	args := []string{
		apiEntity,
		serviceAccountsEntity,
		"createOrgServiceAccount",
		"--version",
		serviceAccountAPIVersion,
		"-P",
		ProfileName(),
		"--debug",
	}

	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = bytes.NewReader(payloadBytes)

	resp, err := RunAndGetStdOut(cmd)
	if err != nil {
		return "", "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var serviceAccount map[string]any
	if err := json.Unmarshal(resp, &serviceAccount); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Obtain client ID
	clientID, ok := serviceAccount["clientId"].(string)
	if !ok {
		return "", "", errors.New("clientId not found in response")
	}

	// Obtain client secret
	secrets, ok := serviceAccount["secrets"].([]any)
	if !ok || len(secrets) == 0 {
		return "", "", errors.New("secrets array not found or empty in response")
	}
	secret, ok := secrets[0].(map[string]any)
	if !ok {
		return "", "", errors.New("secret is not a valid object")
	}
	clientSecret, ok := secret["secret"].(string)
	if !ok {
		return "", "", errors.New("secret value not found in secret")
	}

	return clientID, clientSecret, nil
}

// deleteOrgServiceAccount deletes an organization service account.
func deleteOrgServiceAccount(t *testing.T, cliPath, clientID string) {
	t.Helper()

	args := []string{
		apiEntity,
		serviceAccountsEntity,
		"deleteOrgServiceAccount",
		"--clientId",
		clientID,
		"--version",
		serviceAccountAPIVersion,
		"-P",
		ProfileName(),
	}

	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()

	_, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, "failed to delete service account %s", clientID)
}
