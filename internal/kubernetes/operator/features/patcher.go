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

package features

import (
	"fmt"
	"strings"

	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Patcher is the type that is able to patch Kubernetes objects using a CRD specification.
type Patcher interface {
	Patch(crdSpec *apiextensionsv1.JSONSchemaProps, obj runtime.Object) error
}

// PatcherFunc is a convenience function wrapper around Patcher.
type PatcherFunc func(crdSpec *apiextensionsv1.JSONSchemaProps, obj runtime.Object) error

func (pf PatcherFunc) Patch(crdSpec *apiextensionsv1.JSONSchemaProps, obj runtime.Object) error {
	return pf(crdSpec, obj)
}

// NopPatcher does not patch anything.
func NopPatcher() Patcher {
	return PatcherFunc(func(*apiextensionsv1.JSONSchemaProps, runtime.Object) error {
		return nil
	})
}

// UnknownBackupPolicyFrequencyTypesPruner removes backup policy items from a backup policy
// with unknown frequency types.
// It inspects the CRD definition to determine supported frequency types.
func UnknownBackupPolicyFrequencyTypesPruner(crdSpec *apiextensionsv1.JSONSchemaProps, obj runtime.Object) error {
	// we are not defensive here as this function assumes the invariant
	// of a stable CRD definition for a given version of Kubernetes Atlas Operator.
	frequencyTypePropsEnum := crdSpec.Properties["items"].Items.Schema.Properties["frequencyType"].Enum

	knownFrequencyTypes := make(map[string]struct{})
	for i := range frequencyTypePropsEnum {
		knownFrequencyType := strings.Trim(string(frequencyTypePropsEnum[i].Raw), `"`)
		knownFrequencyTypes[knownFrequencyType] = struct{}{}
	}

	policy, ok := obj.(*akov2.AtlasBackupPolicy)
	if !ok || policy == nil {
		return fmt.Errorf("invalid object: %T: %v", obj, obj)
	}

	prunedItems := make([]akov2.AtlasBackupPolicyItem, 0, len(policy.Spec.Items))
	for i := range policy.Spec.Items {
		frequencyType := policy.Spec.Items[i].FrequencyType
		if _, ok := knownFrequencyTypes[frequencyType]; ok {
			prunedItems = append(prunedItems, policy.Spec.Items[i])
		}
	}
	policy.Spec.Items = prunedItems

	return nil
}
