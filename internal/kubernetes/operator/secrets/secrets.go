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

package secrets

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	TypeLabelKey      = "atlas.mongodb.com/type"
	CredLabelVal      = "credentials"
	PasswordField     = "password"
	CredPrivateAPIKey = "privateApiKey"
	CredPublicAPIKey  = "publicApiKey"
	CredOrgID         = "orgId"
)

func NewAtlasSecret(name, namespace string, data map[string][]byte) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta: v1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasResourceName(name),
			Namespace: namespace,
			Labels: map[string]string{
				TypeLabelKey: CredLabelVal,
			},
		},
		Data: data,
	}
}
