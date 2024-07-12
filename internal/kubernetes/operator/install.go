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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	roleOrgGroupCreator = "ORG_GROUP_CREATOR"
	roleProjectOwner    = "GROUP_OWNER"
)

type Install struct {
	installResources Installer
	atlasStore       store.OperatorGenericStore
	credStore        store.CredentialsGetter
	featureValidator features.FeatureValidator
	kubectl          *kubernetes.KubeCtl

	featureDeletionProtection    bool
	featureSubDeletionProtection bool
	version                      string
	namespace                    string
	watch                        []string
	projectName                  string
	importResources              bool
	atlasGov                     bool
}

func (i *Install) WithNamespace(namespace string) *Install {
	i.namespace = namespace

	return i
}

func (i *Install) WithWatchNamespaces(namespaces []string) *Install {
	i.watch = namespaces

	return i
}

func (i *Install) WithWatchProjectName(name string) *Install {
	i.projectName = name

	return i
}

func (i *Install) WithImportResources(flag bool) *Install {
	i.importResources = flag

	return i
}

func (i *Install) WithResourceDeletionProtection(flag bool) *Install {
	i.featureDeletionProtection = flag

	return i
}

func (i *Install) WithSubResourceDeletionProtection(flag bool) *Install {
	i.featureSubDeletionProtection = flag

	return i
}

func (i *Install) WithAtlasGov(flag bool) *Install {
	i.atlasGov = flag

	return i
}

func (i *Install) Run(ctx context.Context, orgID string) error {
	keys, err := i.generateKeys(orgID)
	if err != nil {
		return err
	}

	if err = i.installResources.InstallCRDs(ctx, i.version, len(i.watch) > 0); err != nil {
		return err
	}

	if err = i.installResources.InstallConfiguration(ctx, &InstallConfig{
		Version:                              i.version,
		Namespace:                            i.namespace,
		Watch:                                i.watch,
		ResourceDeletionProtectionEnabled:    i.featureDeletionProtection,
		SubResourceDeletionProtectionEnabled: i.featureSubDeletionProtection,
		AtlasGov:                             i.atlasGov,
	}); err != nil {
		return err
	}

	if err = i.installResources.InstallCredentials(
		ctx,
		i.namespace,
		orgID,
		keys.GetPublicKey(),
		keys.GetPrivateKey(),
		i.projectName); err != nil {
		return err
	}

	if i.importResources {
		if err = i.importAtlasResources(orgID, keys.GetId()); err != nil {
			return err
		}

		if err = i.ensureCredentialsAssignment(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (i *Install) ensureProject(orgID, projectName string) (*admin.Group, error) {
	project, err := i.atlasStore.ProjectByName(projectName)

	if err == nil {
		return project, nil
	}

	project, err = i.atlasStore.CreateProject(&admin.CreateProjectApiParams{
		Group: &admin.Group{
			Name:                      projectName,
			OrgId:                     orgID,
			RegionUsageRestrictions:   pointer.Get(""),
			WithDefaultAlertsSettings: pointer.Get(true),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return project, nil
}

func (i *Install) generateKeys(orgID string) (*admin.ApiKeyUserDetails, error) {
	if i.projectName == "" {
		input := &admin.CreateAtlasOrganizationApiKey{
			Desc: credentialsGlobalName,
			Roles: []string{
				roleOrgGroupCreator,
			},
		}
		keys, err := i.atlasStore.CreateOrganizationAPIKey(orgID, input)
		if err != nil {
			return nil, fmt.Errorf("failed to generate org keys: %w", err)
		}

		return keys, nil
	}

	project, err := i.ensureProject(orgID, i.projectName)
	if err != nil {
		return nil, err
	}

	input := &admin.CreateAtlasProjectApiKey{
		Desc: fmt.Sprintf(credentialsProjectScopedName, resources.NormalizeAtlasName(i.projectName, resources.AtlasNameToKubernetesName())),
		Roles: []string{
			roleProjectOwner,
		},
	}
	keys, err := i.atlasStore.CreateProjectAPIKey(project.GetId(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to generate project keys: %w", err)
	}

	return keys, nil
}

func (i *Install) importAtlasResources(orgID, apiKeyID string) error {
	projectsIDs := make([]string, 0)

	if i.projectName != "" {
		project, err := i.ensureProject(orgID, i.projectName)
		if err != nil {
			return err
		}

		projectsIDs = append(projectsIDs, project.GetId())
	} else {
		projectsData, err := i.atlasStore.GetOrgProjects(orgID, &store.ListOptions{})
		if err != nil {
			return fmt.Errorf("unable to retrieve list of projects: %w", err)
		}

		for _, project := range projectsData.GetResults() {
			projectsIDs = append(projectsIDs, project.GetId())
		}
	}

	crdVersion, err := features.CRDCompatibleVersion(i.version)
	if err != nil {
		return err
	}

	for _, projectID := range projectsIDs {
		err = i.atlasStore.AssignProjectAPIKey(
			projectID,
			apiKeyID,
			&admin.UpdateAtlasProjectApiKey{
				Roles: &[]string{roleProjectOwner},
			},
		)
		if err != nil {
			return fmt.Errorf("failed to assign api key to project %s: %w", projectID, err)
		}

		exporter := NewConfigExporter(i.atlasStore, i.credStore, projectID, orgID).
			WithTargetNamespace(i.namespace).
			WithTargetOperatorVersion(crdVersion).
			WithFeatureValidator(i.featureValidator).
			WithSecretsData(false)

		err = NewConfigApply(
			NewConfigApplyParams{
				OrgID:     orgID,
				ProjectID: projectID,
				KubeCtl:   i.kubectl,
				Exporter:  exporter,
			},
		).WithTargetOperatorVersion(crdVersion).
			WithNamespace(i.namespace).
			Run()

		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Install) ensureCredentialsAssignment(ctx context.Context) error {
	projects := &akov2.AtlasProjectList{}
	err := i.kubectl.List(ctx, projects, client.InNamespace(i.namespace))
	if err != nil {
		return errors.New("failed to list projects")
	}

	for index := range projects.Items {
		var connectionSecret *akov2common.ResourceRefNamespaced
		project := projects.Items[index]

		if i.projectName != "" {
			if project.Spec.ConnectionSecret != nil && project.Spec.ConnectionSecret.Name != "" {
				err = i.deleteSecret(ctx, *project.Spec.ConnectionSecret.GetObject(project.Namespace))
				if err != nil {
					return fmt.Errorf("failed to cleanup secret for project %s", project.Name)
				}
			}

			connectionSecret = &akov2common.ResourceRefNamespaced{
				Name:      fmt.Sprintf(credentialsProjectScopedName, project.Name),
				Namespace: i.namespace,
			}
		}

		project.Spec.ConnectionSecret = connectionSecret

		if err = i.kubectl.Update(ctx, &project); err != nil {
			return fmt.Errorf("failed to update atlas project %s", i.projectName)
		}
	}

	return nil
}

func (i *Install) deleteSecret(ctx context.Context, key client.ObjectKey) error {
	secret := &corev1.Secret{}
	err := i.kubectl.Get(ctx, key, secret)
	if err != nil {
		return fmt.Errorf("failed to get secret %s", key)
	}

	return i.kubectl.Delete(ctx, secret)
}

func NewInstall(
	installer Installer,
	atlasStore store.OperatorGenericStore,
	credStore store.CredentialsGetter,
	featureValidator features.FeatureValidator,
	kubectl *kubernetes.KubeCtl,
	version string,
) *Install {
	return &Install{
		installResources: installer,
		atlasStore:       atlasStore,
		credStore:        credStore,
		featureValidator: featureValidator,
		kubectl:          kubectl,
		version:          version,
	}
}
