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

package operator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"reflect"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/datafederation"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/deployment"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/federatedauthentication"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/project"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/streamsprocessing"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	yamlSeparator        = "---\r\n"
	maxClusters          = 500
	DefaultClustersCount = 10
	InactiveStatus       = "INACTIVE"

	credentialSuffix = "-credentials"
)

type ConfigExporter struct {
	featureValidator        features.FeatureValidator
	dataProvider            store.OperatorGenericStore
	credsProvider           store.CredentialsGetter
	projectID               string
	clusterNames            []string
	flexClusterNames        []string
	targetNamespace         string
	operatorVersion         string
	includeSecretsData      bool
	orgID                   string
	dictionaryForAtlasNames map[string]string
	dataFederationNames     []string
	patcher                 Patcher
	independentResources    bool
}

type Patcher interface {
	Patch(obj runtime.Object) error
}

var (
	ErrServerless             = errors.New("serverless instance error")
	ErrFlex                   = errors.New("flex cluster error")
	ErrClusterNotFound        = errors.New("cluster not found")
	ErrNoCloudManagerClusters = errors.New("can not get 'advanced clusters' object")
)

func NewConfigExporter(dataProvider store.OperatorGenericStore, credsProvider store.CredentialsGetter, projectID, orgID string) *ConfigExporter {
	return &ConfigExporter{
		dataProvider:            dataProvider,
		credsProvider:           credsProvider,
		projectID:               projectID,
		clusterNames:            []string{},
		dataFederationNames:     []string{},
		targetNamespace:         "",
		includeSecretsData:      false,
		orgID:                   orgID,
		dictionaryForAtlasNames: resources.AtlasNameToKubernetesName(),
	}
}

func (e *ConfigExporter) WithClustersNames(clusters []string) *ConfigExporter {
	e.clusterNames = clusters
	return e
}

func (e *ConfigExporter) WithTargetNamespace(namespace string) *ConfigExporter {
	e.targetNamespace = namespace
	return e
}

func (e *ConfigExporter) WithSecretsData(enabled bool) *ConfigExporter {
	e.includeSecretsData = enabled
	return e
}

func (e *ConfigExporter) WithTargetOperatorVersion(version string) *ConfigExporter {
	e.operatorVersion = version
	return e
}

func (e *ConfigExporter) WithFeatureValidator(validator features.FeatureValidator) *ConfigExporter {
	e.featureValidator = validator
	return e
}

func (e *ConfigExporter) WithDataFederationNames(dataFederations []string) *ConfigExporter {
	e.dataFederationNames = dataFederations
	return e
}

func (e *ConfigExporter) WithPatcher(p Patcher) *ConfigExporter {
	e.patcher = p
	return e
}

func (e *ConfigExporter) WithIndependentResources(enabled bool) *ConfigExporter {
	e.independentResources = enabled
	return e
}
func (e *ConfigExporter) Run() (string, error) {
	// TODO: Add REST to OPERATOR entities matcher
	output := bytes.NewBufferString(yamlSeparator)
	var r []runtime.Object

	serializer := json.NewSerializerWithOptions(
		json.DefaultMetaFactory,
		scheme.Scheme,
		scheme.Scheme,
		json.SerializerOptions{Yaml: true, Pretty: true},
	)

	projectResources, projectName, err := e.exportProject()
	if err != nil {
		return "", err
	}
	r = append(r, projectResources...)

	deploymentsResources, err := e.exportDeployments(projectName)
	if err != nil {
		return "", err
	}
	r = append(r, deploymentsResources...)

	dataFederationResource, err := e.exportDataFederation(projectName)
	if err != nil {
		return "", err
	}
	r = append(r, dataFederationResource...)

	federatedAuthResource, err := e.exportAtlasFederatedAuth(projectName)
	if err != nil {
		return "", err
	}
	r = append(r, federatedAuthResource...)

	streamProcessingResources, err := e.exportAtlasStreamProcessing(projectName)
	if err != nil {
		return "", err
	}
	r = append(r, streamProcessingResources...)

	for _, res := range r {
		if e.patcher != nil {
			err = e.patcher.Patch(res)
			if err != nil {
				return "", fmt.Errorf("error patching %v: %w", res.GetObjectKind().GroupVersionKind(), err)
			}
		}

		err = serializer.Encode(res, output)
		if err != nil {
			return "", err
		}
		output.WriteString(yamlSeparator)
	}

	return output.String(), nil
}

//nolint:gocyclo
func (e *ConfigExporter) exportProject() ([]runtime.Object, string, error) {
	atlasProject, err := e.dataProvider.Project(e.projectID)
	if err != nil {
		return nil, "", err
	}
	if e.orgID != "" && e.orgID != atlasProject.OrgId {
		return nil, "", fmt.Errorf("the project %s (%s) is not part of the organization %q, "+
			"please confirm the arguments provided to the command or you are using the correct profile",
			atlasProject.GetName(), atlasProject.GetId(), e.orgID)
	}
	e.orgID = atlasProject.OrgId

	// Project
	projectData, err := project.BuildAtlasProject(&project.AtlasProjectBuildRequest{
		ProjectStore:    e.dataProvider,
		Project:         atlasProject,
		Validator:       e.featureValidator,
		OrgID:           e.orgID,
		ProjectID:       e.projectID,
		TargetNamespace: e.targetNamespace,
		IncludeSecret:   e.includeSecretsData,
		Dictionary:      e.dictionaryForAtlasNames,
		Version:         e.operatorVersion,
	})
	if err != nil {
		return nil, "", err
	}

	var r []runtime.Object //nolint:prealloc
	r = append(r, projectData.Project)
	for _, secret := range projectData.Secrets {
		r = append(r, secret)
	}

	// Teams
	for _, t := range projectData.Teams {
		r = append(r, t)
	}

	credentialsName := credentialsName(projectData.Project.Name)
	// Project secret with credentials
	r = append(r, project.BuildProjectNamedConnectionSecret(
		e.credsProvider,
		credentialsName,
		projectData.Project.Namespace,
		e.orgID,
		e.includeSecretsData,
		e.dictionaryForAtlasNames,
	))

	if e.featureValidator.IsResourceSupported(features.ResourceAtlasPrivateEndpoint) {
		privateEndpoints, err := project.BuildPrivateEndpointCustomResources(
			e.dataProvider,
			project.PrivateEndpointRequest{
				ProjectName:         projectData.Project.Name,
				ProjectID:           e.projectID,
				TargetNamespace:     e.targetNamespace,
				Version:             e.operatorVersion,
				Credentials:         credentialsName,
				IndependentResource: e.independentResources,
				Dictionary:          e.dictionaryForAtlasNames,
			},
		)
		if err != nil {
			return nil, "", err
		}

		for _, privateEndpoint := range privateEndpoints {
			r = append(r, &privateEndpoint)
		}
	}

	// Independent custom roles (AtlasCustomRole CR)
	if e.featureValidator.IsResourceSupported(features.ResourceAtlasCustomRole) {
		roles, err := project.BuildCustomRoles(e.dataProvider, project.CustomRolesRequest{
			ProjectID:       e.projectID,
			ProjectName:     projectData.Project.Name,
			TargetNamespace: e.targetNamespace,
			Version:         e.operatorVersion,
			Credentials:     credentialsName,
			IsIndependent:   e.independentResources,
			Dict:            e.dictionaryForAtlasNames,
		})
		if err != nil {
			return nil, "", err
		}

		for i := range len(roles) {
			r = append(r, &roles[i])
		}
	}

	// DB users
	usersData, relatedSecrets, err := dbusers.BuildDBUsers(
		e.dataProvider,
		e.projectID,
		projectData.Project.Name,
		e.targetNamespace,
		credentialsName,
		e.dictionaryForAtlasNames,
		e.operatorVersion,
		e.independentResources)
	if err != nil {
		return nil, "", err
	}
	for _, user := range usersData {
		r = append(r, user)
	}
	for _, s := range relatedSecrets {
		r = append(r, s)
	}

	// Backup Compliance Policy
	if projectData.BCP != nil {
		r = append(r, projectData.BCP)
	}

	return r, projectData.Project.Name, nil
}

func (e *ConfigExporter) exportDeployments(projectName string) ([]runtime.Object, error) {
	var result []runtime.Object

	if len(e.clusterNames) == 0 {
		clusters, flex, err := fetchClusterNames(e.dataProvider, e.projectID)
		if err != nil {
			return nil, err
		}
		e.clusterNames = clusters
		e.flexClusterNames = flex
	}

	credentials := credentialsName(projectName)
	for _, deploymentName := range e.clusterNames {
		// Try advanced cluster first
		if advancedCluster, err := deployment.BuildAtlasAdvancedDeployment(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, credentials, e.dictionaryForAtlasNames, e.operatorVersion, e.independentResources); err == nil {
			if advancedCluster != nil {
				// Append deployment to result
				result = append(result, advancedCluster.Deployment)
				// Append backup schedule
				if advancedCluster.BackupSchedule != nil {
					result = append(result, advancedCluster.BackupSchedule)
				}
				// Append backup policies (one)
				for _, policy := range advancedCluster.BackupPolicies {
					if policy != nil {
						result = append(result, policy)
					}
				}
			}
			continue
		}

		// Try serverless cluster next
		serverlessCluster, err := deployment.BuildServerlessDeployments(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion)
		if err == nil {
			if serverlessCluster != nil {
				result = append(result, serverlessCluster)
			}
			continue
		}
		return nil, fmt.Errorf("%w: %s(%s), e: %w", ErrServerless, deploymentName, e.projectID, err)
	}

	for _, deploymentName := range e.flexClusterNames {
		flexCluster, err := deployment.BuildFlexDeployments(e.dataProvider, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion)
		if err == nil {
			if flexCluster != nil {
				result = append(result, flexCluster)
			}
			continue
		}

		return nil, fmt.Errorf("%w: %s(%s), e: %w", ErrFlex, deploymentName, e.projectID, err)
	}

	return result, nil
}

func fetchClusterNames(clustersProvider store.AllClustersLister, projectID string) ([]string, []string, error) {
	flexResult := make(map[string]struct{}, DefaultClustersCount)
	flexClusters, err := clustersProvider.ListFlexClusters(&admin.ListFlexClustersApiParams{ItemsPerPage: pointer.Get(maxClusters)})
	if err != nil {
		return nil, nil, err
	}

	for _, cluster := range flexClusters.GetResults() {
		if reflect.ValueOf(cluster).IsZero() {
			continue
		}
		flexResult[cluster.GetName()] = struct{}{}
	}

	result := make([]string, 0, DefaultClustersCount)
	clusters, err := clustersProvider.ProjectClusters(projectID, &store.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, nil, err
	}

	if clusters == nil {
		return nil, nil, ErrNoCloudManagerClusters
	}

	for _, cluster := range clusters.GetResults() {
		if reflect.ValueOf(cluster).IsZero() {
			continue
		}

		if _, ok := flexResult[cluster.GetName()]; ok {
			continue
		}

		result = append(result, cluster.GetName())
	}

	serverlessInstances, err := clustersProvider.ServerlessInstances(projectID, &store.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, nil, err
	}

	if serverlessInstances == nil {
		return result, maps.Keys(flexResult), nil
	}

	for _, cluster := range serverlessInstances.GetResults() {
		if _, ok := flexResult[cluster.GetName()]; ok {
			continue
		}

		result = append(result, *cluster.Name)
	}

	return result, maps.Keys(flexResult), nil
}

func (e *ConfigExporter) exportDataFederation(projectName string) ([]runtime.Object, error) {
	nameList := e.dataFederationNames
	if len(nameList) == 0 {
		dataFederations, err := e.fetchDataFederationNames()
		if err != nil {
			return nil, err
		}
		nameList = dataFederations
	}
	result := make([]runtime.Object, 0, len(nameList))
	for _, name := range nameList {
		atlasDataFederations, err := datafederation.BuildAtlasDataFederation(e.dataProvider, name, e.projectID, projectName, e.operatorVersion, e.targetNamespace, e.dictionaryForAtlasNames)
		if err != nil {
			return nil, err
		}
		result = append(result, atlasDataFederations)
	}
	return result, nil
}

func (e *ConfigExporter) fetchDataFederationNames() ([]string, error) {
	dataFederations, err := e.dataProvider.DataFederationList(e.projectID)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(dataFederations))
	for _, obj := range dataFederations {
		name := obj.GetName()
		if reflect.ValueOf(name).IsZero() {
			continue
		}
		result = append(result, name)
	}
	return result, nil
}

func (e *ConfigExporter) exportAtlasStreamProcessing(projectName string) ([]runtime.Object, error) {
	if !e.featureValidator.IsResourceSupported(features.ResourceAtlasStreamInstance) ||
		!e.featureValidator.IsResourceSupported(features.ResourceAtlasStreamConnection) {
		return nil, nil
	}

	instancesList, err := e.dataProvider.ProjectStreams(&admin.ListStreamInstancesApiParams{GroupId: e.projectID})
	if err != nil {
		return nil, err
	}
	instances := instancesList.GetResults()
	result := make([]runtime.Object, 0, len(instances))

	for i := range instances {
		instance := instances[i]
		connectionsList, err := e.dataProvider.StreamsConnections(e.projectID, instance.GetName())
		if err != nil {
			return nil, err
		}

		akoInstance, akoConnections, akoSecrets, err := streamsprocessing.BuildAtlasStreamsProcessing(
			e.targetNamespace,
			e.operatorVersion,
			projectName,
			&instance,
			connectionsList.GetResults(),
			e.dictionaryForAtlasNames,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, akoInstance)

		for x := range akoConnections {
			result = append(result, akoConnections[x])
		}

		for x := range akoSecrets {
			result = append(result, akoSecrets[x])
		}
	}

	return result, nil
}

func (e *ConfigExporter) exportAtlasFederatedAuth(projectName string) ([]runtime.Object, error) {
	if !e.featureValidator.IsResourceSupported(features.ResourceAtlasFederatedAuth) {
		return nil, nil
	}
	result := make([]runtime.Object, 0)
	// Gets the FederationAuthSetting
	federatedAuthentificationSetting, err := e.dataProvider.FederationSetting(&admin.GetFederationSettingsApiParams{OrgId: e.orgID})
	if err != nil {
		if admin.IsErrorCode(err, "RESOURCE_NOT_FOUND") {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve federation settings: %w", err)
	}
	// Does not have an IdentityProvider set then no need to generate
	if !federatedAuthentificationSetting.HasIdentityProviderStatus() || federatedAuthentificationSetting.GetIdentityProviderStatus() == InactiveStatus {
		return nil, nil
	}
	// Does have an IdentityProvider and then we can generate the config
	federatedAuthentification, err := federatedauthentication.BuildAtlasFederatedAuth(&federatedauthentication.AtlasFederatedAuthBuildRequest{
		IncludeSecret:                e.includeSecretsData,
		IdentityProviderLister:       e.dataProvider,
		ConnectedOrgConfigsDescriber: e.dataProvider,
		IdentityProviderDescriber:    e.dataProvider,
		ProjectStore:                 e.dataProvider,
		ProjectID:                    e.projectID,
		OrgID:                        e.orgID,
		TargetNamespace:              e.targetNamespace,
		Version:                      e.operatorVersion,
		Dictionary:                   e.dictionaryForAtlasNames,
		ProjectName:                  projectName,
		FederatedSettings:            federatedAuthentificationSetting,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to export federated authentication: %w", err)
	}
	return append(result, federatedAuthentification), nil
}

func credentialsName(projectName string) string {
	return projectName + credentialSuffix
}
