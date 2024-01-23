// Copyright 2022 MongoDB Inc
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
	"reflect"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/datafederation"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/deployment"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/project"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115004/admin"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	yamlSeparator        = "---\r\n"
	maxClusters          = 500
	DefaultClustersCount = 10
)

type ConfigExporter struct {
	featureValidator        features.FeatureValidator
	dataProvider            atlas.OperatorGenericStore
	credsProvider           store.CredentialsGetter
	projectID               string
	clusterNames            []string
	targetNamespace         string
	operatorVersion         string
	includeSecretsData      bool
	orgID                   string
	dictionaryForAtlasNames map[string]string
	dataFederationNames     []string
}

var (
	ErrClusterNotFound        = errors.New("cluster not found")
	ErrNoCloudManagerClusters = errors.New("can not get 'advanced clusters' object")
)

func NewConfigExporter(dataProvider atlas.OperatorGenericStore, credsProvider store.CredentialsGetter, projectID, orgID string) *ConfigExporter {
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

	for _, res := range r {
		err = serializer.Encode(res, output)
		if err != nil {
			return "", err
		}
		output.WriteString(yamlSeparator)
	}

	return output.String(), nil
}

func (e *ConfigExporter) exportProject() ([]runtime.Object, string, error) {
	// Project
	projectData, err := project.BuildAtlasProject(
		e.dataProvider,
		e.featureValidator,
		e.orgID, e.projectID,
		e.targetNamespace,
		e.includeSecretsData,
		e.dictionaryForAtlasNames,
		e.operatorVersion)
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

	// Project secret with credentials
	r = append(r, project.BuildProjectConnectionSecret(
		e.credsProvider,
		projectData.Project.Name,
		projectData.Project.Namespace,
		e.orgID,
		e.includeSecretsData,
		e.dictionaryForAtlasNames,
	))

	// DB users
	usersData, relatedSecrets, err := dbusers.BuildDBUsers(
		e.dataProvider,
		e.projectID,
		projectData.Project.Name,
		e.targetNamespace,
		e.dictionaryForAtlasNames,
		e.operatorVersion)
	if err != nil {
		return nil, "", err
	}
	for _, user := range usersData {
		r = append(r, user)
	}
	for _, s := range relatedSecrets {
		r = append(r, s)
	}

	return r, projectData.Project.Name, nil
}

func (e *ConfigExporter) exportDeployments(projectName string) ([]runtime.Object, error) {
	var result []runtime.Object

	if len(e.clusterNames) == 0 {
		clusters, err := fetchClusterNames(e.dataProvider, e.projectID)
		if err != nil {
			return nil, err
		}
		e.clusterNames = clusters
	}

	for _, deploymentName := range e.clusterNames {
		// Try advanced cluster first
		if advancedCluster, err := deployment.BuildAtlasAdvancedDeployment(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion); err == nil {
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
		if serverlessCluster, err := deployment.BuildServerlessDeployments(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion); err == nil {
			if serverlessCluster != nil {
				result = append(result, serverlessCluster)
			}
			continue
		}
		return nil, fmt.Errorf("%w: %s(%s)", ErrClusterNotFound, deploymentName, e.projectID)
	}
	return result, nil
}

func fetchClusterNames(clustersProvider atlas.AllClustersLister, projectID string) ([]string, error) {
	result := make([]string, 0, DefaultClustersCount)
	response, err := clustersProvider.ProjectClusters(projectID, &atlas.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, err
	}

	if clusters, ok := response.(*atlasv2.PaginatedAdvancedClusterDescription); ok {
		if clusters == nil {
			return nil, ErrNoCloudManagerClusters
		}

		for i := range clusters.Results {
			cluster := clusters.Results[i]
			if reflect.ValueOf(cluster).IsZero() {
				continue
			}
			result = append(result, cluster.GetName())
		}
	}

	serverlessInstances, err := clustersProvider.ServerlessInstances(projectID, &atlas.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, err
	}

	if serverlessInstances == nil {
		return result, nil
	}

	for i := range serverlessInstances.Results {
		cluster := serverlessInstances.Results[i]
		result = append(result, *cluster.Name)
	}

	return result, nil
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
