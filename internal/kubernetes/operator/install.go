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

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/util/toptr"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	roleOrgGroupCreator       = "ORG_GROUP_CREATOR"
	roleProjectOwner          = "GROUP_OWNER"
	atlasErrorProjectNotFound = "GROUP_NAME_NOT_FOUND"
)

type Install struct {
	installResources Installer
	atlasStore       store.AtlasOperatorGenericStore
	credStore        store.CredentialsGetter
	featureValidator features.FeatureValidator
	kubectl          *kubernetes.KubeCtl

	version         string
	namespace       string
	watch           []string
	projectName     string
	importResources bool
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

func (i *Install) Run(ctx context.Context, orgID string) error {
	keys, err := i.generateKeys(orgID)
	if err != nil {
		return err
	}

	if err = i.installResources.InstallCRDs(ctx, i.version, len(i.watch) > 0); err != nil {
		return err
	}

	if err = i.installResources.InstallConfiguration(ctx, i.version, i.namespace, i.watch); err != nil {
		return err
	}

	if err = i.installResources.InstallCredentials(ctx, i.namespace, orgID, keys.PublicKey, keys.PrivateKey, i.projectName == ""); err != nil {
		return err
	}

	if i.importResources {
		if err = i.importAtlasResources(orgID, keys.ID); err != nil {
			return err
		}
	}

	return nil
}

func (i *Install) ensureProject(orgID, projectName string) (*mongodbatlas.Project, error) {
	data, err := i.atlasStore.ProjectByName(projectName)
	if err == nil {
		project, ok := data.(*mongodbatlas.Project)
		if !ok {
			return nil, fmt.Errorf("failed to decode project: %w", err)
		}

		return project, nil
	}

	var apiError *mongodbatlas.ErrorResponse
	errors.As(err, &apiError)

	if apiError.ErrorCode != atlasErrorProjectNotFound {
		return nil, fmt.Errorf("failed to retrieve project: %w", err)
	}

	data, err = i.atlasStore.CreateProject(
		projectName,
		orgID,
		"",
		toptr.MakePtr(true),
		&mongodbatlas.CreateProjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	project, ok := data.(*mongodbatlas.Project)
	if !ok {
		return nil, fmt.Errorf("failed to decode created project: %w", err)
	}

	return project, nil
}

func (i *Install) generateKeys(orgID string) (*mongodbatlas.APIKey, error) {
	if i.projectName == "" {
		input := &mongodbatlas.APIKeyInput{
			Desc: "mongodb-atlas-operator-api-key",
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

	input := &mongodbatlas.APIKeyInput{
		Desc: "mongodb-atlas-project-key",
		Roles: []string{
			roleProjectOwner,
		},
	}
	keys, err := i.atlasStore.CreateProjectAPIKey(project.ID, input)
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

		projectsIDs = append(projectsIDs, project.ID)
	} else {
		projectsData, err := i.atlasStore.GetOrgProjects(orgID, &mongodbatlas.ProjectsListOptions{})
		if err != nil {
			return fmt.Errorf("unable to retrieve list of projects: %w", err)
		}

		projects, ok := projectsData.(*mongodbatlas.Projects)
		if !ok {
			return fmt.Errorf("unable to decode list of projects")
		}

		for _, project := range projects.Results {
			projectsIDs = append(projectsIDs, project.ID)
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
			&mongodbatlas.AssignAPIKey{
				Roles: []string{roleProjectOwner},
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

func NewInstall(
	installer Installer,
	atlasStore store.AtlasOperatorGenericStore,
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
