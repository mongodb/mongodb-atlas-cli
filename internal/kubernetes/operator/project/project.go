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

package project

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2project "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/project"
	akov2provider "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/provider"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8snames "k8s.io/apiserver/pkg/storage/names"
)

const (
	credSecretFormat                = "%s-credentials"
	MaxItems                        = 500
	featureAccessLists              = "projectIpAccessList"
	featureMaintenanceWindows       = "maintenanceWindow"
	featureIntegrations             = "integrations"
	featureNetworkPeering           = "networkPeers"
	featurePrivateEndpoints         = "privateEndpoints"
	featureEncryptionAtRest         = "encryptionAtRest"
	featureCloudProviderAccessRoles = "cloudProviderAccessRoles"
	featureProjectSettings          = "settings"
	featureAuditing                 = "auditing"
	featureAlertConfiguration       = "alertConfigurations"
	featureCustomRoles              = "customRoles"
	featureTeams                    = "teams"
	cidrException                   = "/32"
	datadogIntegrationType          = "DATADOG"
	newRelicIntegrationType         = "NEW_RELIC"
	opsGenieIntegrationType         = "OPS_GENIE"
	pagerDutyIntegrationType        = "PAGER_DUTY"
	victorOpsIntegrationType        = "VICTOR_OPS"
	webhookIntegrationType          = "WEBHOOK"
	microsoftTeamsIntegrationType   = "MICROSOFT_TEAMS"
	slackIntegrationType            = "SLACK"
	prometheusIntegrationType       = "PROMETHEUS"
)

type AtlasProjectBuildRequest struct {
	ProjectStore    store.OperatorProjectStore
	Project         *atlasv2.Group
	Validator       features.FeatureValidator
	OrgID           string
	ProjectID       string
	TargetNamespace string
	IncludeSecret   bool
	Dictionary      map[string]string
	Version         string
}

type AtlasProjectResult struct {
	Project *akov2.AtlasProject
	Secrets []*corev1.Secret
	Teams   []*akov2.AtlasTeam
}

func BuildAtlasProject(br *AtlasProjectBuildRequest) (*AtlasProjectResult, error) { //nolint:gocyclo
	projectResult := newAtlasProject(br.Project, br.Dictionary, br.TargetNamespace, br.Version)

	result := &AtlasProjectResult{
		Project: projectResult,
		Secrets: nil,
		Teams:   nil,
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureAccessLists) {
		ipAccessList, ferr := buildAccessLists(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.ProjectIPAccessList = ipAccessList
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureMaintenanceWindows) {
		maintenanceWindows, ferr := buildMaintenanceWindows(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.MaintenanceWindow = maintenanceWindows
	}

	secretRef := &akov2common.ResourceRefNamespaced{}
	if br.IncludeSecret {
		secretRef.Name = resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, br.Project.Name), br.Dictionary)
	}
	projectResult.Spec.ConnectionSecret = secretRef

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureIntegrations) {
		integrations, intSecrets, ferr := buildIntegrations(br.ProjectStore, br.ProjectID, br.TargetNamespace, true, br.Dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Integrations = integrations
		result.Secrets = intSecrets
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureNetworkPeering) {
		networkPeering, ferr := buildNetworkPeering(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.NetworkPeers = networkPeering
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featurePrivateEndpoints) {
		privateEndpoints, ferr := buildPrivateEndpoints(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.PrivateEndpoints = privateEndpoints
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureEncryptionAtRest) {
		encryptionAtRest, s, ferr := buildEncryptionAtRest(br.ProjectStore, br.ProjectID, br.Project.Name, br.TargetNamespace, br.Dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.EncryptionAtRest = encryptionAtRest
		result.Secrets = append(result.Secrets, s...)
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureCloudProviderAccessRoles) {
		cpa, ferr := buildCloudProviderAccessRoles(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.CloudProviderAccessRoles = cpa
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureProjectSettings) {
		projectSettings, ferr := buildProjectSettings(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Settings = projectSettings
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureAuditing) {
		auditing, ferr := buildAuditing(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Auditing = auditing
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureAlertConfiguration) {
		alertConfigurations, s, ferr := buildAlertConfigurations(br.ProjectStore, br.ProjectID, br.Project.Name, br.TargetNamespace, br.Dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.AlertConfigurations = alertConfigurations
		result.Secrets = append(result.Secrets, s...)
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureCustomRoles) {
		customRoles, ferr := buildCustomRoles(br.ProjectStore, br.ProjectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.CustomRoles = customRoles
	}

	if br.Validator.FeatureExist(features.ResourceAtlasProject, featureTeams) {
		teamsRefs, teams, ferr := buildTeams(br.ProjectStore, br.OrgID, br.ProjectID, br.Project.Name, br.TargetNamespace, br.Version, br.Dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Teams = teamsRefs
		result.Teams = teams
	}

	return result, nil
}

func newAtlasProject(project *atlasv2.Group, dictionary map[string]string, targetNamespace string, version string) *akov2.AtlasProject {
	return &akov2.AtlasProject{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(project.Name, dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: akov2.AtlasProjectSpec{
			Name:                          project.Name,
			ConnectionSecret:              nil,
			ProjectIPAccessList:           nil,
			PrivateEndpoints:              nil,
			CloudProviderAccessRoles:      nil,
			AlertConfigurations:           nil,
			AlertConfigurationSyncEnabled: false,
			NetworkPeers:                  nil,
			WithDefaultAlertsSettings:     project.GetWithDefaultAlertsSettings(),
			X509CertRef:                   nil, // not available for import
			Integrations:                  nil,
			EncryptionAtRest:              nil,
			Auditing:                      nil,
			Settings:                      nil,
			CustomRoles:                   nil,
			Teams:                         nil,
			RegionUsageRestrictions:       project.GetRegionUsageRestrictions(),
		},
		Status: akov2status.AtlasProjectStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}
}

const credentialSuffix = "-credentials"

func BuildProjectConnectionSecret(credsProvider store.CredentialsGetter, name, namespace, orgID string, includeCreds bool, dictionary map[string]string) *corev1.Secret {
	secret := secrets.NewAtlasSecretBuilder(name+credentialSuffix, namespace, dictionary).
		WithData(map[string][]byte{
			secrets.CredOrgID:         []byte(""),
			secrets.CredPublicAPIKey:  []byte(""),
			secrets.CredPrivateAPIKey: []byte(""),
		}).
		Build()
	if includeCreds {
		secret.Data = map[string][]byte{
			secrets.CredOrgID:         []byte(orgID),
			secrets.CredPublicAPIKey:  []byte(credsProvider.PublicAPIKey()),
			secrets.CredPrivateAPIKey: []byte(credsProvider.PrivateAPIKey()),
		}
	}
	return secret
}

func buildCustomRoles(crProvider store.DatabaseRoleLister, projectID string) ([]akov2.CustomRole, error) {
	dbRoles, err := crProvider.DatabaseRoles(projectID)
	if err != nil {
		return nil, err
	}
	if dbRoles == nil {
		return nil, nil
	}

	result := make([]akov2.CustomRole, 0, len(dbRoles))
	roles := dbRoles
	for rIdx := range roles {
		role := &roles[rIdx]

		inhRoles := make([]akov2.Role, 0, len(role.GetInheritedRoles()))
		for _, rl := range role.GetInheritedRoles() {
			inhRoles = append(inhRoles, akov2.Role{
				Name:     rl.Role,
				Database: rl.Db,
			})
		}

		actions := make([]akov2.Action, 0, len(role.GetActions()))
		for _, action := range role.GetActions() {
			r := make([]akov2.Resource, 0, len(action.GetResources()))
			for _, res := range action.GetResources() {
				r = append(r, akov2.Resource{
					Cluster:    pointer.Get(res.Cluster),
					Database:   pointer.Get(res.Db),
					Collection: pointer.Get(res.Collection),
				})
			}
			actions = append(actions, akov2.Action{
				Name:      action.Action,
				Resources: r,
			})
		}
		result = append(result, akov2.CustomRole{
			Name:           role.RoleName,
			InheritedRoles: inhRoles,
			Actions:        actions,
		})
	}
	return result, nil
}

func buildAccessLists(accessListProvider store.ProjectIPAccessListLister, projectID string) ([]akov2project.IPAccessList, error) {
	// pagination not required, max 200 entries can be configured via API
	accessLists, err := accessListProvider.ProjectIPAccessLists(projectID, &store.ListOptions{ItemsPerPage: MaxItems})
	if err != nil {
		return nil, err
	}

	result := make([]akov2project.IPAccessList, 0, len(accessLists.GetResults()))
	for _, list := range accessLists.GetResults() {
		if strings.HasSuffix(list.GetCidrBlock(), cidrException) && list.GetIpAddress() != "" {
			list.CidrBlock = pointer.Get("")
		}
		deleteAfterDate := ""
		if !list.GetDeleteAfterDate().IsZero() {
			deleteAfterDate = list.GetDeleteAfterDate().String()
		}
		result = append(result, akov2project.IPAccessList{
			AwsSecurityGroup: list.GetAwsSecurityGroup(),
			CIDRBlock:        list.GetCidrBlock(),
			Comment:          list.GetComment(),
			DeleteAfterDate:  deleteAfterDate,
			IPAddress:        list.GetIpAddress(),
		})
	}
	return result, nil
}

func buildMaintenanceWindows(mwProvider store.MaintenanceWindowDescriber, projectID string) (akov2project.MaintenanceWindow, error) {
	mw, err := mwProvider.MaintenanceWindow(projectID)
	if err != nil {
		return akov2project.MaintenanceWindow{}, err
	}

	return akov2project.MaintenanceWindow{
		DayOfWeek: mw.DayOfWeek,
		HourOfDay: mw.GetHourOfDay(),
		AutoDefer: mw.GetAutoDeferOnceEnabled(),
		StartASAP: mw.GetStartASAP(),
		Defer:     false,
	}, nil
}

func buildIntegrations(intProvider store.IntegrationLister, projectID, targetNamespace string, includeSecrets bool, dictionary map[string]string) ([]akov2project.Integration, []*corev1.Secret, error) { //nolint:gocyclo
	integrations, err := intProvider.Integrations(projectID)
	if err != nil {
		return nil, nil, err
	}
	result := make([]akov2project.Integration, 0, len(integrations.GetResults()))
	intSecrets := make([]*corev1.Secret, 0, len(integrations.GetResults()))

	for _, list := range integrations.GetResults() {
		iType := list.GetType()
		secret := secrets.NewAtlasSecretBuilder(
			fmt.Sprintf("%s-integration-%s", projectID, strings.ToLower(iType)),
			targetNamespace,
			dictionary,
		).WithData(map[string][]byte{secrets.PasswordField: []byte("")}).Build()

		integration := akov2project.Integration{
			Type: iType,
		}
		secretRef := akov2common.ResourceRefNamespaced{
			Name:      resources.NormalizeAtlasName(secret.Name, dictionary),
			Namespace: targetNamespace,
		}
		switch iType {
		case pagerDutyIntegrationType:
			integration.ServiceKeyRef = secretRef
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetServiceKey())
			}
		case slackIntegrationType:
			integration.TeamName = list.GetTeamName()
			integration.APITokenRef = secretRef
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetApiToken())
			}
		case datadogIntegrationType:
			integration.Region = list.GetRegion()
			integration.APIKeyRef = secretRef
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetApiKey())
			}
		case opsGenieIntegrationType:
			integration.Region = list.GetRegion()
			integration.APIKeyRef = secretRef
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetApiKey())
			}
		case webhookIntegrationType:
			integration.URL = list.GetUrl()
			integration.SecretRef = secretRef
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetSecret())
			}
		case microsoftTeamsIntegrationType:
			integration.MicrosoftTeamsWebhookURL = list.GetMicrosoftTeamsWebhookUrl()
		case prometheusIntegrationType:
			integration.UserName = list.GetUsername()
			integration.PasswordRef = secretRef
			integration.ServiceDiscovery = list.GetServiceDiscovery()
			integration.Enabled = list.GetEnabled()
			if includeSecrets {
				secret.Data[secrets.PasswordField] = []byte(list.GetPassword())
			}
		case victorOpsIntegrationType: // One more secret required
			integration.APIKeyRef = secretRef
			secret.Data[secrets.PasswordField] = []byte(list.GetApiKey())

			var routingKeyData string
			if includeSecrets {
				routingKeyData = list.GetRoutingKey()
			}
			if list.GetRoutingKey() != "" {
				// Secret with routing key
				routingSecret := secrets.NewAtlasSecretBuilder(
					fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
					targetNamespace,
					dictionary,
				).WithData(map[string][]byte{secrets.PasswordField: []byte(routingKeyData)}).Build()
				intSecrets = append(intSecrets, routingSecret)
			}
		case newRelicIntegrationType:
			integration.LicenseKeyRef = secretRef
			secret.Data[secrets.PasswordField] = []byte(list.GetLicenseKey())
			// Secrets with write and read tokens
			var writeToken, readToken string
			if includeSecrets {
				writeToken = list.GetWriteToken()
				readToken = list.GetReadToken()
			}
			writeTokenSecret := secrets.NewAtlasSecretBuilder(
				fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
				targetNamespace,
				dictionary,
			).WithData(map[string][]byte{secrets.PasswordField: []byte(writeToken)}).Build()
			readTokenSecret := secrets.NewAtlasSecretBuilder(
				fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
				targetNamespace,
				dictionary,
			).WithData(map[string][]byte{secrets.PasswordField: []byte(readToken)}).Build()
			intSecrets = append(intSecrets, writeTokenSecret, readTokenSecret)
		}
		result = append(result, integration)
		intSecrets = append(intSecrets, secret)
	}

	return result, intSecrets, nil
}

func buildPrivateEndpoints(peProvider store.PrivateEndpointLister, projectID string) ([]akov2.PrivateEndpoint, error) {
	var result []akov2.PrivateEndpoint
	for _, cloudProvider := range []akov2provider.ProviderName{akov2provider.ProviderAWS, akov2provider.ProviderGCP, akov2provider.ProviderAzure} {
		peList, err := peProvider.PrivateEndpoints(projectID, string(cloudProvider))
		if err != nil {
			return nil, err
		}
		for i := range peList {
			pe := &peList[i]
			result = append(result, akov2.PrivateEndpoint{
				Provider:          cloudProvider,
				Region:            pe.GetRegionName(),
				ID:                firstElementOrZeroValue(pe.GetInterfaceEndpoints()),
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         akov2.GCPEndpoints{},
			})
		}
	}
	return result, nil
}

func buildNetworkPeering(npProvider store.PeeringConnectionLister, projectID string) ([]akov2.NetworkPeer, error) {
	// pagination not required, max 25 entries per provider can be configured via API
	npListAWS, err := npProvider.PeeringConnections(projectID, &store.ContainersListOptions{
		ListOptions: store.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(akov2provider.ProviderAWS),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for AWS: %w", err)
	}

	npListGCP, err := npProvider.PeeringConnections(projectID, &store.ContainersListOptions{
		ListOptions: store.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(akov2provider.ProviderGCP),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for GCP: %w", err)
	}

	npListAzure, err := npProvider.PeeringConnections(projectID, &store.ContainersListOptions{
		ListOptions: store.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(akov2provider.ProviderAzure),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for Azure: %w", err)
	}

	result := make([]akov2.NetworkPeer, 0, len(npListAWS)+len(npListGCP)+len(npListAzure))

	for i := range npListAWS {
		np := npListAWS[i]
		result = append(result, convertNetworkPeer(np, akov2provider.ProviderAWS))
	}

	for i := range npListGCP {
		np := npListGCP[i]
		result = append(result, convertNetworkPeer(np, akov2provider.ProviderGCP))
	}

	for i := range npListAzure {
		np := npListAzure[i]
		result = append(result, convertNetworkPeer(np, akov2provider.ProviderAzure))
	}

	return result, nil
}

func convertNetworkPeer(np atlasv2.BaseNetworkPeeringConnectionSettings, providerName akov2provider.ProviderName) akov2.NetworkPeer {
	switch np.GetProviderName() {
	case "AWS":
		return convertAWSNetworkPeer(&np, providerName)
	case "GCP":
		return convertGCPNetworkPeer(&np, providerName)
	case "Azure":
		return convertAzureNetworkPeer(&np, providerName)
	default:
		return akov2.NetworkPeer{}
	}
}

func convertAWSNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName akov2provider.ProviderName) akov2.NetworkPeer {
	return akov2.NetworkPeer{
		AccepterRegionName:  np.GetAccepterRegionName(),
		AWSAccountID:        np.GetAwsAccountId(),
		ContainerRegion:     "",
		ContainerID:         np.ContainerId,
		ProviderName:        providerName,
		RouteTableCIDRBlock: np.GetRouteTableCidrBlock(),
		VpcID:               np.GetVpcId(),
	}
}

func convertAzureNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName akov2provider.ProviderName) akov2.NetworkPeer {
	return akov2.NetworkPeer{
		AzureDirectoryID:    np.GetAzureDirectoryId(),
		AzureSubscriptionID: np.GetAzureSubscriptionId(),
		ContainerRegion:     "",
		ContainerID:         np.GetContainerId(),
		ProviderName:        providerName,
		ResourceGroupName:   np.GetResourceGroupName(),
		VNetName:            np.GetVnetName(),
	}
}

func convertGCPNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName akov2provider.ProviderName) akov2.NetworkPeer {
	return akov2.NetworkPeer{
		GCPProjectID:    np.GetGcpProjectId(),
		ContainerRegion: "",
		ContainerID:     np.ContainerId,
		ProviderName:    providerName,
		NetworkName:     np.GetNetworkName(),
	}
}

func buildEncryptionAtRest(encProvider store.EncryptionAtRestDescriber, projectID, projectName, targetNamespace string, dictionary map[string]string) (*akov2.EncryptionAtRest, []*corev1.Secret, error) {
	data, err := encProvider.EncryptionAtRest(projectID)
	if err != nil {
		return nil, nil, err
	}

	ref := &akov2.EncryptionAtRest{
		AwsKms: akov2.AwsKms{
			Enabled: data.AwsKms.Enabled,
			Region:  data.AwsKms.GetRegion(),
			Valid:   data.AwsKms.Valid,
		},
		AzureKeyVault: akov2.AzureKeyVault{
			Enabled:           data.AzureKeyVault.Enabled,
			ClientID:          data.AzureKeyVault.GetClientID(),
			AzureEnvironment:  data.AzureKeyVault.GetAzureEnvironment(),
			ResourceGroupName: data.AzureKeyVault.GetResourceGroupName(),
			TenantID:          data.AzureKeyVault.GetTenantID(),
		},
		GoogleCloudKms: akov2.GoogleCloudKms{
			Enabled: data.GoogleCloudKms.Enabled,
		},
	}

	var ss []*corev1.Secret
	switch {
	case data.AwsKms.Enabled != nil && *data.AwsKms.Enabled:
		ref.AwsKms.SecretRef = akov2common.ResourceRefNamespaced{
			Name:      resources.NormalizeAtlasName(generateName("aws-credentials-"), dictionary),
			Namespace: targetNamespace,
		}

		ss = append(ss, secrets.NewAtlasSecretBuilder(ref.AwsKms.SecretRef.Name, ref.AwsKms.SecretRef.Namespace, dictionary).
			WithData(map[string][]byte{"CustomerMasterKeyID": []byte(""), "RoleID": []byte("")}).
			WithProjectLabels(projectID, projectName).
			Build())

	case data.AzureKeyVault.Enabled != nil && *data.AzureKeyVault.Enabled:
		ref.AzureKeyVault.SecretRef = akov2common.ResourceRefNamespaced{
			Name:      resources.NormalizeAtlasName(generateName("azure-credentials-"), dictionary),
			Namespace: targetNamespace,
		}

		ss = append(ss, secrets.NewAtlasSecretBuilder(ref.AzureKeyVault.SecretRef.Name, ref.AzureKeyVault.SecretRef.Namespace, dictionary).
			WithData(map[string][]byte{"SubscriptionID": []byte(""), "KeyVaultName": []byte(""), "KeyIdentifier": []byte(""), "Secret": []byte("")}).
			WithProjectLabels(projectID, projectName).
			Build())

	case data.GoogleCloudKms.Enabled != nil && *data.GoogleCloudKms.Enabled:
		ref.AzureKeyVault.SecretRef = akov2common.ResourceRefNamespaced{
			Name:      resources.NormalizeAtlasName(generateName("gcp-credentials-"), dictionary),
			Namespace: targetNamespace,
		}

		ss = append(ss, secrets.NewAtlasSecretBuilder(ref.GoogleCloudKms.SecretRef.Name, ref.GoogleCloudKms.SecretRef.Namespace, dictionary).
			WithData(map[string][]byte{"ServiceAccountKey": []byte(""), "KeyVersionResourceID": []byte("")}).
			WithProjectLabels(projectID, projectName).
			Build())
	}

	return ref, ss, nil
}

func buildCloudProviderAccessRoles(cpaProvider store.CloudProviderAccessRoleLister, projectID string) ([]akov2.CloudProviderAccessRole, error) {
	data, err := cpaProvider.CloudProviderAccessRoles(projectID)
	if err != nil {
		return nil, err
	}

	result := make([]akov2.CloudProviderAccessRole, 0, len(data.GetAwsIamRoles()))
	for _, cpa := range data.GetAwsIamRoles() {
		result = append(result, akov2.CloudProviderAccessRole{
			ProviderName:      cpa.ProviderName,
			IamAssumedRoleArn: cpa.GetIamAssumedRoleArn(),
		})
	}

	return result, nil
}

func buildProjectSettings(psProvider store.ProjectSettingsDescriber, projectID string) (*akov2.ProjectSettings, error) {
	data, err := psProvider.ProjectSettings(projectID)
	if err != nil {
		return nil, err
	}

	return &akov2.ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: data.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsDataExplorerEnabled:                       data.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 data.IsPerformanceAdvisorEnabled,
		IsRealtimePerformancePanelEnabled:           data.IsRealtimePerformancePanelEnabled,
		IsSchemaAdvisorEnabled:                      data.IsSchemaAdvisorEnabled,
	}, nil
}

func buildAuditing(auditingProvider store.AuditingDescriber, projectID string) (*akov2.Auditing, error) {
	data, err := auditingProvider.Auditing(projectID)
	if err != nil {
		return nil, err
	}

	return &akov2.Auditing{
		AuditAuthorizationSuccess: data.GetAuditAuthorizationSuccess(),
		AuditFilter:               data.GetAuditFilter(),
		Enabled:                   data.GetEnabled(),
	}, nil
}

func buildAlertConfigurations(acProvider store.AlertConfigurationLister, projectID, projectName, targetNamespace string, dictionary map[string]string) ([]akov2.AlertConfiguration, []*corev1.Secret, error) {
	data, err := acProvider.AlertConfigurations(&atlasv2.ListAlertConfigurationsApiParams{
		GroupId:      projectID,
		ItemsPerPage: pointer.Get(MaxItems),
	})
	if err != nil {
		return nil, nil, err
	}

	convertNotifications := func(atlasNotifications []atlasv2.AlertsNotificationRootForGroup) ([]akov2.Notification, []*corev1.Secret) {
		var (
			akoNotifications []akov2.Notification
			akoSecrets       []*corev1.Secret
		)

		for _, atlasNotification := range atlasNotifications {
			akoNotification := akov2.Notification{
				ChannelName:    atlasNotification.GetChannelName(),
				DatadogRegion:  atlasNotification.GetDatadogRegion(),
				DelayMin:       atlasNotification.DelayMin,
				EmailAddress:   atlasNotification.GetEmailAddress(),
				EmailEnabled:   atlasNotification.EmailEnabled,
				IntervalMin:    atlasNotification.GetIntervalMin(),
				MobileNumber:   atlasNotification.GetMobileNumber(),
				OpsGenieRegion: atlasNotification.GetOpsGenieRegion(),
				SMSEnabled:     atlasNotification.SmsEnabled,
				TeamID:         atlasNotification.GetTeamId(),
				TeamName:       atlasNotification.GetTeamName(),
				TypeName:       atlasNotification.GetTypeName(),
				Username:       atlasNotification.GetUsername(),
				Roles:          atlasNotification.GetRoles(),
			}

			if atlasNotification.TypeName != nil {
				switch *atlasNotification.TypeName {
				case pagerDutyIntegrationType:
					akoNotification.ServiceKeyRef = akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(generateName("service-key-"), dictionary),
						Namespace: targetNamespace,
					}

					akoSecrets = append(akoSecrets,
						secrets.NewAtlasSecretBuilder(akoNotification.ServiceKeyRef.Name, akoNotification.ServiceKeyRef.Namespace, dictionary).
							WithData(map[string][]byte{"ServiceKey": []byte("")}).
							WithProjectLabels(projectID, projectName).
							WithNotifierLabels(atlasNotification.NotifierId, pagerDutyIntegrationType).
							Build())

				case slackIntegrationType:
					akoNotification.APITokenRef = akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(generateName("api-token-"), dictionary),
						Namespace: targetNamespace,
					}

					akoSecrets = append(akoSecrets,
						secrets.NewAtlasSecretBuilder(akoNotification.APITokenRef.Name, akoNotification.APITokenRef.Namespace, dictionary).
							WithData(map[string][]byte{"APIToken": []byte("")}).
							WithProjectLabels(projectID, projectName).
							WithNotifierLabels(atlasNotification.NotifierId, slackIntegrationType).
							Build())

				case datadogIntegrationType:
					akoNotification.DatadogAPIKeyRef = akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(generateName("datadog-api-key-"), dictionary),
						Namespace: targetNamespace,
					}

					akoSecrets = append(akoSecrets,
						secrets.NewAtlasSecretBuilder(akoNotification.DatadogAPIKeyRef.Name, akoNotification.DatadogAPIKeyRef.Namespace, dictionary).
							WithData(map[string][]byte{"DatadogAPIKey": []byte("")}).
							WithProjectLabels(projectID, projectName).
							WithNotifierLabels(atlasNotification.NotifierId, datadogIntegrationType).
							Build())

				case opsGenieIntegrationType:
					akoNotification.OpsGenieAPIKeyRef = akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(generateName("ops-genie-api-key-"), dictionary),
						Namespace: targetNamespace,
					}

					akoSecrets = append(akoSecrets,
						secrets.NewAtlasSecretBuilder(akoNotification.OpsGenieAPIKeyRef.Name, akoNotification.OpsGenieAPIKeyRef.Namespace, dictionary).
							WithData(map[string][]byte{"OpsGenieAPIKey": []byte("")}).
							WithProjectLabels(projectID, projectName).
							WithNotifierLabels(atlasNotification.NotifierId, opsGenieIntegrationType).
							Build())

				case victorOpsIntegrationType:
					akoNotification.VictorOpsSecretRef = akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(generateName("victor-ops-credentials-"), dictionary),
						Namespace: targetNamespace,
					}

					akoSecrets = append(akoSecrets,
						secrets.NewAtlasSecretBuilder(akoNotification.VictorOpsSecretRef.Name, akoNotification.VictorOpsSecretRef.Namespace, dictionary).
							WithData(map[string][]byte{"VictorOpsAPIKey": []byte(""), "VictorOpsRoutingKey": []byte("")}).
							WithProjectLabels(projectID, projectName).
							WithNotifierLabels(atlasNotification.NotifierId, victorOpsIntegrationType).
							Build())
				}
			}

			akoNotifications = append(akoNotifications, akoNotification)
		}

		return akoNotifications, akoSecrets
	}

	var secretResults []*corev1.Secret
	results := make([]akov2.AlertConfiguration, 0, len(data.GetResults()))
	for _, alertConfig := range data.GetResults() {
		notifications, notificationSecrets := convertNotifications(alertConfig.GetNotifications())
		secretResults = append(secretResults, notificationSecrets...)

		results = append(results, akov2.AlertConfiguration{
			EventTypeName:   alertConfig.GetEventTypeName(),
			Enabled:         alertConfig.GetEnabled(),
			Matchers:        convertMatchers(alertConfig.GetMatchers()),
			MetricThreshold: convertMetricThreshold(alertConfig.MetricThreshold),
			Threshold:       convertThreshold(alertConfig.Threshold),
			Notifications:   notifications,
		})
	}
	return results, secretResults, nil
}

func convertMatchers(atlasMatcher []atlasv2.StreamsMatcher) []akov2.Matcher {
	res := make([]akov2.Matcher, 0, len(atlasMatcher))
	for _, m := range atlasMatcher {
		matcher, err := toMatcher(m)
		if err != nil {
			_, _ = log.Warningf("Skipping matcher %v, conversion failed: %v\n", m, err.Error())
			continue
		}
		res = append(res, matcher)
	}
	return res
}

func toMatcher(m atlasv2.StreamsMatcher) (akov2.Matcher, error) {
	var matcher akov2.Matcher

	matcher.FieldName = m.GetFieldName()
	matcher.Operator = m.GetOperator()
	matcher.Value = m.GetValue()

	if matcher.FieldName == "" && matcher.Operator == "" && matcher.Value == "" {
		return matcher, errors.New("matcher is empty")
	}
	if matcher.FieldName == "" {
		return matcher, errors.New("matcher's fieldName is not set")
	}
	if matcher.Operator == "" {
		return matcher, errors.New("matcher's operator is not set")
	}
	if matcher.Value == "" {
		return matcher, errors.New("matcher's value is not set")
	}
	return matcher, nil
}

func convertMetricThreshold(atlasMT *atlasv2.ServerlessMetricThreshold) *akov2.MetricThreshold {
	if atlasMT == nil {
		return &akov2.MetricThreshold{}
	}
	return &akov2.MetricThreshold{
		MetricName: atlasMT.MetricName,
		Operator:   atlasMT.GetOperator(),
		Threshold:  fmt.Sprintf("%f", atlasMT.GetThreshold()),
		Units:      atlasMT.GetUnits(),
		Mode:       atlasMT.GetMode(),
	}
}

func convertThreshold(atlasT *atlasv2.GreaterThanRawThreshold) *akov2.Threshold {
	if atlasT == nil {
		return &akov2.Threshold{}
	}
	return &akov2.Threshold{
		Operator:  atlasT.GetOperator(),
		Units:     atlasT.GetUnits(),
		Threshold: strconv.Itoa(atlasT.GetThreshold()),
	}
}

func generateName(base string) string {
	return k8snames.SimpleNameGenerator.GenerateName(base)
}

func buildTeams(teamsProvider store.OperatorTeamsStore, orgID, projectID, projectName, targetNamespace, version string, dictionary map[string]string) ([]akov2.Team, []*akov2.AtlasTeam, error) {
	projectTeams, err := teamsProvider.ProjectTeams(projectID, nil)
	if err != nil {
		return nil, nil, err
	}

	fetchUsers := func(teamID string) ([]string, error) {
		users, err := teamsProvider.TeamUsers(orgID, teamID)
		if err != nil {
			return nil, err
		}
		result := make([]string, 0, len(users.GetResults()))
		for _, user := range users.GetResults() {
			result = append(result, user.Username)
		}
		return result, nil
	}

	convertRoleNames := func(input []string) []akov2.TeamRole {
		if len(input) == 0 {
			return nil
		}
		result := make([]akov2.TeamRole, 0, len(input))
		for i := range input {
			result = append(result, akov2.TeamRole(input[i]))
		}
		return result
	}

	convertUserNames := func(input []string) []akov2.TeamUser {
		if len(input) == 0 {
			return nil
		}

		result := make([]akov2.TeamUser, 0, len(input))
		for i := range input {
			result = append(result, akov2.TeamUser(input[i]))
		}
		return result
	}

	teamsRefs := make([]akov2.Team, 0, len(projectTeams.GetResults()))
	atlasTeamCRs := make([]*akov2.AtlasTeam, 0, len(projectTeams.GetResults()))

	for _, teamRef := range projectTeams.GetResults() {
		teamID := teamRef.GetTeamId()

		team, err := teamsProvider.TeamByID(orgID, teamID)
		if err != nil {
			return nil, nil, fmt.Errorf("team id: %s is assigned to project %s (id: %s) but not found. %w",
				teamID, projectName, projectID, err)
		}

		teamName := team.GetName()
		crName := resources.NormalizeAtlasName(fmt.Sprintf("%s-team-%s", projectName, teamName), dictionary)
		teamsRefs = append(teamsRefs, akov2.Team{
			TeamRef: akov2common.ResourceRefNamespaced{
				Name:      crName,
				Namespace: targetNamespace,
			},
			Roles: convertRoleNames(teamRef.GetRoleNames()),
		})

		users, err := fetchUsers(team.GetId())
		if err != nil {
			return nil, nil, err
		}

		atlasTeamCRs = append(atlasTeamCRs, &akov2.AtlasTeam{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasTeam",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      crName,
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: akov2.TeamSpec{
				Name:      team.GetName(),
				Usernames: convertUserNames(users),
			},
			Status: akov2status.TeamStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		})
	}

	return teamsRefs, atlasTeamCRs, nil
}

func firstElementOrZeroValue[T any](collection []T) T {
	var item T

	if len(collection) > 0 {
		item = collection[0]
	}

	return item
}
