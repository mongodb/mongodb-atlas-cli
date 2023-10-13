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

package features

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/crds"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

const (
	LatestOperatorMajorVersion  = "1.9.0"
	maxDepth                    = 100
	ResourceVersion             = "app.kubernetes.io/version"
	ResourceAtlasProject        = "atlasprojects"
	ResourceAtlasDeployment     = "atlasdeployments"
	ResourceAtlasDatabaseUser   = "atlasdatabaseusers"
	ResourceAtlasBackupSchedule = "atlasbackupschedules"
	ResourceAtlasBackupPolicy   = "atlasbackuppolicies"
	ResourceAtlasTeam           = "atlasteams"
	ResourceAtlasDataFederation = "atlasdatafederations"
	ResourceAtlasFederatedAuth  = "atlasfederatedauths"
)

var (
	ErrVersionNotSupportedFmt    = "version '%s' is not supported"
	ErrDownloadResourceFailedFmt = "can not download resource '%s': %v"
	ErrDocumentIsEmpty           = errors.New("document is empty")
	ErrDocumentHasNoVersions     = errors.New("document contains no versions")
	ErrDocumentHasNoSchema       = errors.New("document contains no Schema")
	ErrDocumentHasNoSpec         = errors.New("document contains no Spec")

	versionsToResourcesMap = map[string][]string{
		"1.7.0": {
			ResourceAtlasDatabaseUser,
			ResourceAtlasProject,
			ResourceAtlasDeployment,
			ResourceAtlasBackupSchedule,
			ResourceAtlasBackupPolicy,
			ResourceAtlasTeam,
		},
		"1.8.0": {
			ResourceAtlasDatabaseUser,
			ResourceAtlasProject,
			ResourceAtlasDeployment,
			ResourceAtlasBackupSchedule,
			ResourceAtlasBackupPolicy,
			ResourceAtlasTeam,
			ResourceAtlasDataFederation,
		},
		"1.9.0": {
			ResourceAtlasDatabaseUser,
			ResourceAtlasProject,
			ResourceAtlasDeployment,
			ResourceAtlasBackupSchedule,
			ResourceAtlasBackupPolicy,
			ResourceAtlasTeam,
			ResourceAtlasDataFederation,
			ResourceAtlasFederatedAuth,
		},
	}
)

func GetResourcesForVersion(version string) ([]string, bool) {
	v := semver.MustParse(version)
	majorVersion := semver.New(v.Major(), v.Minor(), 0, "", "").String()
	resources, ok := versionsToResourcesMap[majorVersion]
	return resources, ok
}

func SupportedVersions() []string {
	result := make([]string, 0, len(versionsToResourcesMap))
	for version := range versionsToResourcesMap {
		result = append(result, version)
	}
	return result
}

func CRDCompatibleVersion(operatorVersion string) (string, error) {
	operatorVersionSem, err := semver.NewVersion(operatorVersion)
	if err != nil {
		return "", fmt.Errorf("operator version %s is invalid", operatorVersion)
	}

	latestCRDVersionSem, err := semver.NewVersion(LatestOperatorMajorVersion)
	if err != nil {
		return "", fmt.Errorf("CRD version %s is invalid", LatestOperatorMajorVersion)
	}

	if operatorVersionSem.GreaterThan(latestCRDVersionSem) {
		return LatestOperatorMajorVersion, nil
	}

	return semver.New(
		operatorVersionSem.Major(),
		operatorVersionSem.Minor(),
		0,
		"",
		"").String(), nil
}

type AtlasCRDs struct {
	resources map[string]*apiextensions.JSONSchemaProps
}

func NewAtlasCRDs(crdProvider crds.AtlasOperatorCRDProvider, version string) (*AtlasCRDs, error) {
	resources, versionFound := GetResourcesForVersion(version)
	if !versionFound {
		return nil, fmt.Errorf(ErrVersionNotSupportedFmt, version)
	}

	result := &AtlasCRDs{resources: map[string]*apiextensions.JSONSchemaProps{}}

	for _, resource := range resources {
		crd, err := crdProvider.GetAtlasOperatorResource(resource, version)
		if err != nil {
			return nil, fmt.Errorf(ErrDownloadResourceFailedFmt, resource, err)
		}
		// we only interested in the Spec section of a document
		root, err := getCRDRoot(crd)
		if err != nil {
			return nil, fmt.Errorf("failed to process CRD '%s:%s'. err: %w", resource, version, err)
		}
		result.resources[resource] = root
	}

	return result, nil
}

// FeatureExist
// resourceName: one of SupportedResources
// featurePath: dot-separated string - path in CRD spec to check.
func (a *AtlasCRDs) FeatureExist(resourceName, featurePath string) bool {
	if res, ok := a.resources[resourceName]; ok {
		if pathExists(featurePath, res) {
			return true
		}
	}
	return false
}

func pathExists(path string, data *apiextensions.JSONSchemaProps) bool {
	parts := strings.Split(path, ".")
	if len(parts) == 0 || data == nil {
		return false
	}

	var lookup func(path []string, data *apiextensions.JSONSchemaProps, depth int) bool
	lookup = func(path []string, data *apiextensions.JSONSchemaProps, depth int) bool {
		if len(path) == 0 {
			return true
		}

		if depth == 0 || data == nil {
			return false
		}

		if props, ok := data.Properties[path[0]]; ok {
			return lookup(path[1:], &props, depth-1)
		} else if data.Items != nil {
			if len(data.Items.JSONSchemas) == 0 {
				return lookup(path, data.Items.Schema, depth-1)
			}
			for i := 0; i < len(data.Items.JSONSchemas); i++ {
				if lookup(path, &data.Items.JSONSchemas[i], depth-1) {
					return true
				}
			}
			return false
		}
		return false
	}

	return lookup(parts, data, maxDepth)
}

func getCRDRoot(document *apiextensions.CustomResourceDefinition) (*apiextensions.JSONSchemaProps, error) {
	if document == nil {
		return nil, ErrDocumentIsEmpty
	}

	if len(document.Spec.Versions) == 0 {
		return nil, ErrDocumentHasNoVersions
	}

	// There is only one version of Atlas CRDs atm
	if document.Spec.Versions[0].Schema == nil || document.Spec.Versions[0].Schema.OpenAPIV3Schema == nil {
		return nil, ErrDocumentHasNoSchema
	}

	specs, ok := document.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties["spec"]
	if !ok {
		return nil, ErrDocumentHasNoSpec
	}

	return &specs, nil
}
