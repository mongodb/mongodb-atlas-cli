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
	"testing"
	"time"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	alertsEntity                 = "alerts"
	configEntity                 = "settings"
	dbusersEntity                = "dbusers"
	certsEntity                  = "certs"
	privateEndpointsEntity       = "privateendpoints"
	onlineArchiveEntity          = "onlineArchives"
	projectEntity                = "project"
	orgEntity                    = "org"
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
	accessLogsEntity             = "accessLogs"
	accessListEntity             = "accessList"
	performanceAdvisorEntity     = "performanceAdvisor"
	slowQueryLogsEntity          = "slowQueryLogs"
	namespacesEntity             = "namespaces"
	suggestedIndexesEntity       = "suggestedIndexes"
	slowOperationThresholdEntity = "slowOperationThreshold"
	tierM30                      = "M30"
	tierM10                      = "M10"
	tierM2                       = "M2"
	diskSizeGB40                 = "40"
	diskSizeGB30                 = "30"
	projectsEntity               = "projects"
	settingsEntity               = "settings"
	backupsEntity                = "backups"
	exportsEntity                = "exports"
	bucketsEntity                = "buckets"
)

// Cluster settings.
const (
	e2eClusterTier     = "M30"
	e2eClusterProvider = "AWS" // e2eClusterProvider preferred provider for e2e testing.
	e2eMDBVer          = "4.4"
	e2eSharedMDBVer    = "5.0"
)

const (
	defaultTimeout   = 30 * time.Minute
	defaultTick      = 1 * time.Minute
	ResourceNotFound = "RESOURCE_NOT_FOUND"
)

type K8SHelper struct {
	atlasClient *mongodbatlas.Client
	project     *mongodbatlas.Project
	clusters    map[string]*mongodbatlas.Cluster
	orgID       string
	t           *testing.T
}

func NewK8sHelper(t *testing.T) *K8SHelper {
	t.Helper()

	// These ENV variables MUST be set before running the test
	orgID := os.Getenv("MCLI_ORG_ID")
	privateKey := os.Getenv("MCLI_PRIVATE_API_KEY")
	publicKey := os.Getenv("MCLI_PUBLIC_API_KEY")
	opsManagerURL := os.Getenv("MCLI_OPS_MANAGER_URL")
	transport := digest.NewTransport(publicKey, privateKey)
	tClient, err := transport.Client()
	require.NoError(t, err, "can not create transport for atlas client")

	atlasClient := mongodbatlas.NewClient(tClient)
	atlasClient.BaseURL, err = url.Parse(opsManagerURL)
	require.NoError(t, err, "can not create atlas client")

	return &K8SHelper{
		atlasClient: atlasClient,
		t:           t,
		orgID:       orgID,
		clusters:    map[string]*mongodbatlas.Cluster{},
	}
}

func (kh *K8SHelper) NewProject(project *mongodbatlas.Project) {
	kh.t.Helper()

	createdProject, resp, err := kh.atlasClient.Projects.Create(context.Background(), project, &mongodbatlas.CreateProjectOptions{})
	require.Nil(kh.t, err, "error while creating project", err, resp.Status)
	kh.project = createdProject

	assert.Eventually(kh.t, func() bool {
		_, _, err := kh.atlasClient.Projects.GetOneProject(context.Background(), kh.project.ID)
		return assert.NoError(kh.t, err, "waiting for project to be created", kh.project.ID)
	}, defaultTimeout, defaultTick, "project should've been created")

	kh.t.Cleanup(func() {
		assert.Eventually(kh.t, func() bool {
			_, err := kh.atlasClient.Projects.Delete(context.Background(), kh.project.ID)
			if err != nil {
				var apiError *mongodbatlas.ErrorResponse
				if errors.As(err, &apiError) && (apiError.ErrorCode == "GROUP_NOT_FOUND" || apiError.ErrorCode == ResourceNotFound) {
					return true
				}
				return false
			}
			return true
		}, defaultTimeout, defaultTick, "project should've been deleted", project.ID)
	})
}

func (kh *K8SHelper) NewCluster(cluster *mongodbatlas.Cluster) {
	kh.t.Helper()

	require.True(kh.t, kh.project != nil, "can not create cluster without a project")
	createdCluster, _, err := kh.atlasClient.Clusters.Create(context.Background(), kh.project.ID, cluster)
	require.Nil(kh.t, err, "error while creating cluster")

	kh.clusters[cluster.Name] = createdCluster

	assert.Eventually(kh.t, func() bool {
		_, _, err := kh.atlasClient.Clusters.Get(context.Background(), kh.project.ID, cluster.Name)
		return assert.NoError(kh.t, err, "waiting for cluster to be created", cluster.Name)
	}, defaultTimeout, defaultTick, "cluster should've been created")

	kh.t.Cleanup(func() {
		assert.Eventually(kh.t, func() bool {
			_, err := kh.atlasClient.Clusters.Delete(context.Background(), kh.project.ID, cluster.Name)
			if !assert.NoError(kh.t, err, "can not delete test cluster", cluster.Name) {
				return false
			}
			_, _, err = kh.atlasClient.Clusters.Get(context.Background(), kh.project.ID, cluster.Name)
			if err != nil {
				var apiError *mongodbatlas.ErrorResponse
				if errors.As(err, &apiError) && (apiError.ErrorCode == "CLUSTER_NOT_FOUND" || apiError.ErrorCode == ResourceNotFound) {
					return true
				}
				return false
			}
			return true
		}, defaultTimeout, defaultTick, "cluster should've been deleted", cluster.Name, cluster.ID)
	})
}

func (kh *K8SHelper) NewServerlessInstance(instance *mongodbatlas.ServerlessCreateRequestParams) {
	kh.t.Helper()
	require.True(kh.t, kh.project != nil, "can not create serverless instance without a project")
	createdInstance, _, err := kh.atlasClient.ServerlessInstances.Create(context.Background(), kh.project.ID, instance)
	require.Nil(kh.t, err, "error while creating serverless instance")

	kh.clusters[createdInstance.Name] = createdInstance

	assert.Eventually(kh.t, func() bool {
		_, _, err := kh.atlasClient.ServerlessInstances.Get(context.Background(), kh.project.ID, createdInstance.Name)
		return assert.NoError(kh.t, err, "waiting for serverless instance to be created", createdInstance.Name)
	}, defaultTimeout, defaultTick, "serverless instance should've been created")

	kh.t.Cleanup(func() {
		assert.Eventually(kh.t, func() bool {
			_, err := kh.atlasClient.ServerlessInstances.Delete(context.Background(), kh.project.ID, createdInstance.Name)
			if !assert.NoError(kh.t, err, "can not delete test cluster", createdInstance.Name) {
				return false
			}
			_, _, err = kh.atlasClient.ServerlessInstances.Get(context.Background(), kh.project.ID, createdInstance.Name)
			if err != nil {
				var apiError *mongodbatlas.ErrorResponse
				if errors.As(err, &apiError) && (apiError.ErrorCode == "SERVERLESS_INSTANCE_NOT_FOUND" || apiError.ErrorCode == ResourceNotFound) {
					return true
				}
				return false
			}
			return true
		}, defaultTimeout, defaultTick, "serverless instance should've been deleted", createdInstance.Name, createdInstance.ID)
	})
}

func deployClusterForProject(projectID string) (string, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return "", err
	}
	clusterName, err := RandClusterName()
	if err != nil {
		return "", err
	}
	region, err := newAvailableRegion(projectID, e2eClusterTier, e2eClusterProvider)
	if err != nil {
		return "", err
	}
	args := []string{
		clustersEntity,
		"create",
		clusterName,
		"--mdbVersion", e2eMDBVer,
		"--region", region,
		"--tier", e2eClusterTier,
		"--provider", e2eClusterProvider,
		"--diskSizeGB=30",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := create.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error creating cluster %w: %s", err, string(resp))
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
		return "", fmt.Errorf("error watching cluster %w: %s", err, string(resp))
	}
	return clusterName, nil
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

func integrationExists(name string, thirdPartyIntegrations mongodbatlas.ThirdPartyIntegrations) bool {
	services := thirdPartyIntegrations.Results
	for i := range services {
		if services[i].Type == name {
			return true
		}
	}
	return false
}

func Gov() bool {
	return os.Getenv("MCLI_SERVICE") == "cloudgov"
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
	if Gov() {
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

func ensureCluster(t *testing.T, cluster *mongodbatlas.AdvancedCluster, clusterName, version string, diskSizeGB float64) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.Name)
	a.Equal(version, cluster.MongoDBMajorVersion)
	a.Equal(diskSizeGB, *cluster.DiskSizeGB)
}

func ensureSharedCluster(t *testing.T, cluster *mongodbatlas.Cluster, clusterName, version, tier string, diskSizeGB float64) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.Name)
	a.Equal(version, cluster.MongoDBMajorVersion)
	if cluster.ProviderSettings != nil {
		a.Equal(tier, cluster.ProviderSettings.InstanceSizeName)
	}
	a.Equal(diskSizeGB, *cluster.DiskSizeGB)
}
