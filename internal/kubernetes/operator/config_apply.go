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

package operator

import (
	"context"
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	akov1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigApply struct {
	OrgID        string
	ProjectID    string
	ClusterNames []string

	Namespace string
	Version   string

	kubeCtl  *kubernetes.KubeCtl
	exporter *ConfigExporter
}

type NewConfigApplyParams struct {
	OrgID     string
	ProjectID string

	KubeCtl  *kubernetes.KubeCtl
	Exporter *ConfigExporter
}

func NewConfigApply(params NewConfigApplyParams) *ConfigApply {
	return &ConfigApply{
		OrgID:     params.OrgID,
		ProjectID: params.ProjectID,
		kubeCtl:   params.KubeCtl,
		exporter:  params.Exporter,
	}
}

func (apply *ConfigApply) WithTargetOperatorVersion(version string) *ConfigApply {
	apply.Version = version

	return apply
}

func (apply *ConfigApply) WithNamespace(namespace string) *ConfigApply {
	apply.Namespace = namespace

	return apply
}

func (apply *ConfigApply) Run() error {
	ProjectResources, projectName, err := apply.exporter.exportProject()
	if err != nil {
		return err
	}

	DeploymentResources, err := apply.exporter.exportDeployments(projectName)
	if err != nil {
		return err
	}

	sortedResources := sortResources(ProjectResources, DeploymentResources, apply.Version)

	for _, objects := range sortedResources {
		for _, object := range objects {
			ctrlObj, ok := object.(client.Object)
			if !ok {
				return errors.New("unable to apply resource")
			}

			err = apply.kubeCtl.Create(context.Background(), ctrlObj, &client.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func sortResources(projectResources, deploymentResources []runtime.Object, version string) [][]runtime.Object {
	resources, versionFound := features.GetResourcesForVersion(version)
	if !versionFound {
		return nil
	}

	sortedResources := make([][]runtime.Object, len(resources)+1)

	for _, resource := range projectResources {
		if _, ok := resource.(*corev1.Secret); ok {
			sortedResources[0] = append(sortedResources[0], resource)
		}

		if _, ok := resource.(*akov1.AtlasTeam); ok {
			sortedResources[1] = append(sortedResources[1], resource)
		}

		if _, ok := resource.(*akov1.AtlasProject); ok {
			sortedResources[2] = append(sortedResources[2], resource)
		}

		if _, ok := resource.(*akov1.AtlasDatabaseUser); ok {
			sortedResources[3] = append(sortedResources[3], resource)
		}
	}

	for _, resource := range deploymentResources {
		if _, ok := resource.(*akov1.AtlasBackupPolicy); ok {
			sortedResources[4] = append(sortedResources[4], resource)
		}

		if _, ok := resource.(*akov1.AtlasBackupSchedule); ok {
			sortedResources[5] = append(sortedResources[5], resource)
		}

		if _, ok := resource.(*akov1.AtlasDeployment); ok {
			sortedResources[6] = append(sortedResources[6], resource)
		}
	}

	return sortedResources
}
