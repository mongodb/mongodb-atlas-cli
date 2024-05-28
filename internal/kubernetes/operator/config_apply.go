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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
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

	streamsResources, err := apply.exporter.exportAtlasStreamProcessing(projectName)
	if err != nil {
		return err
	}

	sortedResources := sortResources(ProjectResources, DeploymentResources, streamsResources, apply.Version)

	for _, objects := range sortedResources {
		for _, object := range objects {
			if apply.exporter.patcher != nil {
				err = apply.exporter.patcher.Patch(object)
				if err != nil {
					return fmt.Errorf("error patching %v: %w", object.GetObjectKind().GroupVersionKind(), err)
				}
			}

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

func sortResources(
	projectResources, deploymentResources, streamsResources []runtime.Object,
	version string,
) [][]runtime.Object {
	resources, versionFound := features.GetResourcesForVersion(version)
	if !versionFound {
		return nil
	}

	sortedResources := make([][]runtime.Object, len(resources)+1)

	for _, resource := range projectResources {
		if _, ok := resource.(*corev1.Secret); ok {
			sortedResources[0] = append(sortedResources[0], resource)
		}

		if _, ok := resource.(*akov2.AtlasTeam); ok {
			sortedResources[1] = append(sortedResources[1], resource)
		}

		if _, ok := resource.(*akov2.AtlasProject); ok {
			sortedResources[2] = append(sortedResources[2], resource)
		}

		if _, ok := resource.(*akov2.AtlasDatabaseUser); ok {
			sortedResources[3] = append(sortedResources[3], resource)
		}
	}

	for _, resource := range deploymentResources {
		if _, ok := resource.(*akov2.AtlasBackupPolicy); ok {
			sortedResources[4] = append(sortedResources[4], resource)
		}

		if _, ok := resource.(*akov2.AtlasBackupSchedule); ok {
			sortedResources[5] = append(sortedResources[5], resource)
		}

		if _, ok := resource.(*akov2.AtlasDeployment); ok {
			sortedResources[6] = append(sortedResources[6], resource)
		}
	}

	for _, resource := range streamsResources {
		if _, ok := resource.(*corev1.Secret); ok {
			sortedResources[0] = append(sortedResources[0], resource)
		}

		if _, ok := resource.(*akov2.AtlasStreamConnection); ok {
			sortedResources[7] = append(sortedResources[7], resource)
		}

		if _, ok := resource.(*akov2.AtlasStreamInstance); ok {
			sortedResources[8] = append(sortedResources[8], resource)
		}
	}

	return sortedResources
}
