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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	TypeLabelKey         = "atlas.mongodb.com/type"
	ProjectIDLabelKey    = "atlas.mongodb.com/project-id"
	ProjectNameLabelKey  = "atlas.mongodb.com/project-name"
	NotifierIDLabelKey   = "atlas.mongodb.com/notifier-id"
	NotifierNameLabelKey = "atlas.mongodb.com/notifier-type-name"
	CredLabelVal         = "credentials"
	UsernameField        = "username"
	PasswordField        = "password"
	CertificateField     = "certificate"
	CredPrivateAPIKey    = "privateApiKey"
	CredPublicAPIKey     = "publicApiKey"
	CredOrgID            = "orgId"
)

type AtlasSecretBuilder func() (*corev1.Secret, map[string]string)

func NewAtlasSecretBuilder(name, namespace string, dictionary map[string]string) AtlasSecretBuilder {
	return func() (*corev1.Secret, map[string]string) {
		secret := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(name, dictionary),
				Namespace: namespace,
				Labels: map[string]string{
					TypeLabelKey: CredLabelVal,
				},
			},
		}
		return secret, dictionary
	}
}

func (a AtlasSecretBuilder) WithData(data map[string][]byte) AtlasSecretBuilder {
	return func() (*corev1.Secret, map[string]string) {
		s, d := a()
		s.Data = data
		return s, d
	}
}

func (a AtlasSecretBuilder) WithProjectLabels(id, name string) AtlasSecretBuilder {
	return func() (*corev1.Secret, map[string]string) {
		s, d := a()
		s.Labels[ProjectIDLabelKey] = resources.NormalizeAtlasName(id, d)
		s.Labels[ProjectNameLabelKey] = resources.NormalizeAtlasName(name, d)
		return s, d
	}
}

func (a AtlasSecretBuilder) WithNotifierLabels(id *string, typeName string) AtlasSecretBuilder {
	return func() (*corev1.Secret, map[string]string) {
		s, d := a()
		if id == nil {
			return s, d
		}
		s.Labels[NotifierIDLabelKey] = resources.NormalizeAtlasName(*id, d)
		s.Labels[NotifierNameLabelKey] = typeName // don't normalize type name, as it is already a short form
		return s, d
	}
}

func (a AtlasSecretBuilder) Build() *corev1.Secret {
	secret, _ := a()
	return secret
}
