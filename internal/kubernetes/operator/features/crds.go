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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/crds"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	LatestOperatorMajorVersion          = "2.4.0"
	maxDepth                            = 100
	ResourceVersion                     = "mongodb.com/atlas-resource-version"
	ResourceAtlasProject                = "atlasprojects"
	ResourceAtlasDeployment             = "atlasdeployments"
	ResourceAtlasDatabaseUser           = "atlasdatabaseusers"
	ResourceAtlasBackupSchedule         = "atlasbackupschedules"
	ResourceAtlasBackupPolicy           = "atlasbackuppolicies"
	ResourceAtlasTeam                   = "atlasteams"
	ResourceAtlasDataFederation         = "atlasdatafederations"
	ResourceAtlasFederatedAuth          = "atlasfederatedauths"
	ResourceAtlasStreamInstance         = "atlasstreaminstances"
	ResourceAtlasStreamConnection       = "atlasstreamconnections"
	ResourceAtlasBackupCompliancePolicy = "atlasbackupcompliancepolicies"
)

var (
	ErrVersionNotSupportedFmt    = "version '%s' is not supported"
	ErrDownloadResourceFailedFmt = "can not download resource '%s': %v"
	ErrDocumentIsEmpty           = errors.New("document is empty")
	ErrDocumentHasNoVersions     = errors.New("document contains no versions")
	ErrDocumentHasNoSchema       = errors.New("document contains no Schema")
	ErrDocumentHasNoSpec         = errors.New("document contains no Spec")

	versionsToResourcesMap = map[string][]resource{
		"2.2.0": {
			resource{ResourceAtlasDatabaseUser, NopPatcher()},
			resource{ResourceAtlasProject, NopPatcher()},
			resource{ResourceAtlasDeployment, NopPatcher()},
			resource{ResourceAtlasBackupSchedule, NopPatcher()},
			resource{ResourceAtlasBackupPolicy, NopPatcher()},
			resource{ResourceAtlasTeam, NopPatcher()},
			resource{ResourceAtlasDataFederation, NopPatcher()},
			resource{ResourceAtlasFederatedAuth, NopPatcher()},
		},
		"2.3.0": {
			resource{ResourceAtlasDatabaseUser, NopPatcher()},
			resource{ResourceAtlasProject, NopPatcher()},
			resource{ResourceAtlasDeployment, NopPatcher()},
			resource{ResourceAtlasBackupSchedule, NopPatcher()},
			resource{ResourceAtlasBackupPolicy, NopPatcher()},
			resource{ResourceAtlasTeam, NopPatcher()},
			resource{ResourceAtlasDataFederation, NopPatcher()},
			resource{ResourceAtlasFederatedAuth, NopPatcher()},
			resource{ResourceAtlasStreamInstance, NopPatcher()},
			resource{ResourceAtlasStreamConnection, NopPatcher()},
		},
		"2.4.0": {
			resource{ResourceAtlasDatabaseUser, NopPatcher()},
			resource{ResourceAtlasProject, NopPatcher()},
			resource{ResourceAtlasDeployment, NopPatcher()},
			resource{ResourceAtlasBackupSchedule, NopPatcher()},
			resource{ResourceAtlasBackupPolicy, NopPatcher()},
			resource{ResourceAtlasTeam, NopPatcher()},
			resource{ResourceAtlasDataFederation, NopPatcher()},
			resource{ResourceAtlasFederatedAuth, NopPatcher()},
			resource{ResourceAtlasStreamInstance, NopPatcher()},
			resource{ResourceAtlasStreamConnection, NopPatcher()},
			resource{ResourceAtlasBackupCompliancePolicy, NopPatcher()},
		},
	}
)

type resource struct {
	name    string
	patcher Patcher
}

func majorVersion(version string) string {
	v := semver.MustParse(version)
	return semver.New(v.Major(), v.Minor(), 0, "", "").String()
}

func GetResourcesForVersion(version string) ([]string, bool) {
	resources, ok := versionsToResourcesMap[majorVersion(version)]
	if !ok {
		return nil, false
	}
	result := make([]string, 0, len(resources))
	for i := range resources {
		result = append(result, resources[i].name)
	}
	return result, true
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
	resources map[string]*apiextensionsv1.JSONSchemaProps
	patchers  map[string]Patcher
}

func (a *AtlasCRDs) Patch(obj runtime.Object) error {
	// Despite marked as unsafe this pluralizer works well on our types.
	plural, _ := meta.UnsafeGuessKindToResource(obj.GetObjectKind().GroupVersionKind())

	crdSpec, ok := a.resources[plural.Resource]
	if !ok {
		return nil
	}
	patcher, ok := a.patchers[plural.Resource]
	if !ok {
		return nil
	}
	return patcher.Patch(crdSpec, obj)
}

func NewAtlasCRDs(crdProvider crds.AtlasOperatorCRDProvider, version string) (*AtlasCRDs, error) {
	resources, versionFound := versionsToResourcesMap[majorVersion(version)]
	if !versionFound {
		return nil, fmt.Errorf(ErrVersionNotSupportedFmt, version)
	}

	result := &AtlasCRDs{
		resources: map[string]*apiextensionsv1.JSONSchemaProps{},
		patchers:  map[string]Patcher{},
	}

	for _, resource := range resources {
		crd, err := crdProvider.GetAtlasOperatorResource(resource.name, version)
		if err != nil {
			return nil, fmt.Errorf(ErrDownloadResourceFailedFmt, resource, err)
		}
		// we only interested in the Spec section of a document
		root, err := getCRDRoot(crd)
		if err != nil {
			return nil, fmt.Errorf("failed to process CRD '%s:%s'. err: %w", resource, version, err)
		}
		result.resources[resource.name] = root
		result.patchers[resource.name] = resource.patcher
	}

	return result, nil
}

func (a *AtlasCRDs) IsResourceSupported(resourceName string) bool {
	_, ok := a.resources[resourceName]

	return ok
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

func pathExists(path string, data *apiextensionsv1.JSONSchemaProps) bool {
	parts := strings.Split(path, ".")
	if len(parts) == 0 || data == nil {
		return false
	}

	var lookup func(path []string, data *apiextensionsv1.JSONSchemaProps, depth int) bool
	lookup = func(path []string, data *apiextensionsv1.JSONSchemaProps, depth int) bool {
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

func getCRDRoot(document *apiextensionsv1.CustomResourceDefinition) (*apiextensionsv1.JSONSchemaProps, error) {
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
