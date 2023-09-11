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
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	operatorProject "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/provider"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201006/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

var (
	ErrAtlasProject  = errors.New("can not get 'atlas project' resource")
	ErrTeamsAssigned = errors.New("can not get 'atlas assigned teams' resource")
	ErrTeamUsers     = errors.New("can not get 'users' objects")
)

type AtlasProjectResult struct {
	Project *atlasV1.AtlasProject
	Secrets []*corev1.Secret
	Teams   []*atlasV1.AtlasTeam
}

func BuildAtlasProject(projectStore store.AtlasOperatorProjectStore, validator features.FeatureValidator, orgID, projectID, targetNamespace string, includeSecret bool, dictionary map[string]string, version string) (*AtlasProjectResult, error) {
	data, err := projectStore.Project(projectID)
	if err != nil {
		return nil, err
	}

	project, ok := data.(*atlas.Project)
	if !ok {
		return nil, ErrAtlasProject
	}

	projectResult := &atlasV1.AtlasProject{
		TypeMeta: v1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(project.Name, dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: version,
			},
		},
		Spec: atlasV1.AtlasProjectSpec{
			Name:                          project.Name,
			ConnectionSecret:              nil,
			ProjectIPAccessList:           nil,
			PrivateEndpoints:              nil,
			CloudProviderAccessRoles:      nil,
			AlertConfigurations:           nil,
			AlertConfigurationSyncEnabled: false,
			NetworkPeers:                  nil,
			WithDefaultAlertsSettings:     pointer.GetOrDefault(project.WithDefaultAlertsSettings, false),
			X509CertRef:                   nil, // not available for import
			Integrations:                  nil,
			EncryptionAtRest:              nil,
			Auditing:                      nil,
			Settings:                      nil,
			CustomRoles:                   nil,
			Teams:                         nil,
		},
		Status: status.AtlasProjectStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}

	result := &AtlasProjectResult{
		Project: projectResult,
		Secrets: nil,
		Teams:   nil,
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureAccessLists) {
		ipAccessList, ferr := buildAccessLists(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.ProjectIPAccessList = ipAccessList
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureMaintenanceWindows) {
		maintenanceWindows, ferr := buildMaintenanceWindows(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.MaintenanceWindow = maintenanceWindows
	}

	secretRef := &common.ResourceRefNamespaced{}
	if includeSecret {
		secretRef.Name = resources.NormalizeAtlasName(fmt.Sprintf(credSecretFormat, project.Name), dictionary)
	}
	projectResult.Spec.ConnectionSecret = secretRef

	if validator.FeatureExist(features.ResourceAtlasProject, featureIntegrations) {
		integrations, intSecrets, ferr := buildIntegrations(projectStore, projectID, targetNamespace, true, dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Integrations = integrations
		result.Secrets = intSecrets
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureNetworkPeering) {
		networkPeering, ferr := buildNetworkPeering(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.NetworkPeers = networkPeering
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featurePrivateEndpoints) {
		privateEndpoints, ferr := buildPrivateEndpoints(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.PrivateEndpoints = privateEndpoints
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureEncryptionAtRest) {
		encryptionAtRest, ferr := buildEncryptionAtRest(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.EncryptionAtRest = encryptionAtRest
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureCloudProviderAccessRoles) {
		cpa, ferr := buildCloudProviderAccessRoles(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.CloudProviderAccessRoles = cpa
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureProjectSettings) {
		projectSettings, ferr := buildProjectSettings(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Settings = projectSettings
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureAuditing) {
		auditing, ferr := buildAuditing(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Auditing = auditing
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureAlertConfiguration) {
		alertConfigurations, ferr := buildAlertConfigurations(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.AlertConfigurations = alertConfigurations
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureCustomRoles) {
		customRoles, ferr := buildCustomRoles(projectStore, projectID)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.CustomRoles = customRoles
	}

	if validator.FeatureExist(features.ResourceAtlasProject, featureTeams) {
		teamsRefs, teams, ferr := buildTeams(projectStore, orgID, projectID, project.Name, targetNamespace, version, dictionary)
		if ferr != nil {
			return nil, ferr
		}
		projectResult.Spec.Teams = teamsRefs
		result.Teams = teams
	}

	return result, err
}

func BuildProjectConnectionSecret(credsProvider store.CredentialsGetter, name, namespace, orgID string, includeCreds bool, dictionary map[string]string) *corev1.Secret {
	secret := secrets.NewAtlasSecret(fmt.Sprintf("%s-credentials", name), namespace, map[string][]byte{
		secrets.CredOrgID:         []byte(""),
		secrets.CredPublicAPIKey:  []byte(""),
		secrets.CredPrivateAPIKey: []byte(""),
	},
		dictionary)
	if includeCreds {
		secret.Data = map[string][]byte{
			secrets.CredOrgID:         []byte(orgID),
			secrets.CredPublicAPIKey:  []byte(credsProvider.PublicAPIKey()),
			secrets.CredPrivateAPIKey: []byte(credsProvider.PrivateAPIKey()),
		}
	}
	return secret
}

func buildCustomRoles(crProvider store.DatabaseRoleLister, projectID string) ([]atlasV1.CustomRole, error) {
	dbRoles, err := crProvider.DatabaseRoles(projectID)
	if err != nil {
		return nil, err
	}
	if dbRoles == nil {
		return nil, nil
	}

	result := make([]atlasV1.CustomRole, 0, len(dbRoles))
	roles := dbRoles
	for rIdx := range roles {
		role := &roles[rIdx]

		inhRoles := make([]atlasV1.Role, 0, len(role.InheritedRoles))
		for inhRIdx := range role.InheritedRoles {
			rl := &role.InheritedRoles[inhRIdx]
			inhRoles = append(inhRoles, atlasV1.Role{
				Name:     rl.Role,
				Database: rl.Db,
			})
		}

		actions := make([]atlasV1.Action, 0, len(role.Actions))
		for aIdx := range role.Actions {
			action := &role.Actions[aIdx]

			resources := make([]atlasV1.Resource, 0, len(action.Resources))
			for resIdx := range action.Resources {
				res := &action.Resources[resIdx]
				resources = append(resources, atlasV1.Resource{
					Cluster:    &res.Cluster,
					Database:   &res.Db,
					Collection: &res.Collection,
				})
			}
			actions = append(actions, atlasV1.Action{
				Name:      action.Action,
				Resources: resources,
			})
		}
		result = append(result, atlasV1.CustomRole{
			Name:           role.RoleName,
			InheritedRoles: inhRoles,
			Actions:        actions,
		})
	}
	return result, nil
}

func buildAccessLists(accessListProvider store.ProjectIPAccessListLister, projectID string) ([]operatorProject.IPAccessList, error) {
	// pagination not required, max 200 entries can be configured via API
	accessLists, err := accessListProvider.ProjectIPAccessLists(projectID, &atlas.ListOptions{ItemsPerPage: MaxItems})
	if err != nil {
		return nil, err
	}

	var result []operatorProject.IPAccessList
	for _, list := range accessLists.Results {
		if strings.HasSuffix(list.GetCidrBlock(), cidrException) && list.GetIpAddress() != "" {
			list.CidrBlock = pointer.Get("")
		}
		deleteAfterDate := ""
		if !list.GetDeleteAfterDate().IsZero() {
			deleteAfterDate = list.GetDeleteAfterDate().String()
		}
		result = append(result, operatorProject.IPAccessList{
			AwsSecurityGroup: list.GetAwsSecurityGroup(),
			CIDRBlock:        list.GetCidrBlock(),
			Comment:          list.GetComment(),
			DeleteAfterDate:  deleteAfterDate,
			IPAddress:        list.GetIpAddress(),
		})
	}
	return result, nil
}

func buildMaintenanceWindows(mwProvider store.MaintenanceWindowDescriber, projectID string) (operatorProject.MaintenanceWindow, error) {
	mw, err := mwProvider.MaintenanceWindow(projectID)
	if err != nil {
		return operatorProject.MaintenanceWindow{}, err
	}

	return operatorProject.MaintenanceWindow{
		DayOfWeek: mw.DayOfWeek,
		HourOfDay: mw.HourOfDay,
		AutoDefer: pointer.GetOrDefault(mw.AutoDeferOnceEnabled, false),
		StartASAP: pointer.GetOrDefault(mw.StartASAP, false),
		Defer:     false,
	}, nil
}

func buildIntegrations(intProvider store.IntegrationLister, projectID, targetNamespace string, includeSecrets bool, dictionary map[string]string) ([]operatorProject.Integration, []*corev1.Secret, error) {
	integrations, err := intProvider.Integrations(projectID)
	if err != nil {
		return nil, nil, err
	}
	var result []operatorProject.Integration
	var intSecrets []*corev1.Secret

	for _, list := range integrations.Results {
		iType := list.GetType()
		secret := secrets.NewAtlasSecret(fmt.Sprintf("%s-integration-%s", projectID, strings.ToLower(iType)),
			targetNamespace, map[string][]byte{secrets.PasswordField: []byte("")}, dictionary)

		integration := operatorProject.Integration{
			Type: iType,
		}
		secretRef := common.ResourceRefNamespaced{
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
				routingSecret := secrets.NewAtlasSecret(fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
					targetNamespace,
					map[string][]byte{secrets.PasswordField: []byte(routingKeyData)}, dictionary)
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
			writeTokenSecret := secrets.NewAtlasSecret(fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
				targetNamespace,
				map[string][]byte{secrets.PasswordField: []byte(writeToken)}, dictionary)
			readTokenSecret := secrets.NewAtlasSecret(fmt.Sprintf("%s-integration-%s-routing-key", projectID, strings.ToLower(iType)),
				targetNamespace,
				map[string][]byte{secrets.PasswordField: []byte(readToken)},
				dictionary,
			)
			intSecrets = append(intSecrets, writeTokenSecret, readTokenSecret)
		}
		result = append(result, integration)
		intSecrets = append(intSecrets, secret)
	}

	return result, intSecrets, nil
}

func buildPrivateEndpoints(peProvider store.PrivateEndpointLister, projectID string) ([]atlasV1.PrivateEndpoint, error) {
	var result []atlasV1.PrivateEndpoint
	for _, cloudProvider := range []provider.ProviderName{provider.ProviderAWS, provider.ProviderGCP, provider.ProviderAzure} {
		peList, err := peProvider.PrivateEndpoints(projectID, string(cloudProvider))
		if err != nil {
			return nil, err
		}
		for i := range peList {
			pe := &peList[i]
			result = append(result, atlasV1.PrivateEndpoint{
				Provider:          cloudProvider,
				IP:                "",
				GCPProjectID:      "",
				EndpointGroupName: "",
				Endpoints:         atlasV1.GCPEndpoints{},
				ID:                pe.GetId(),
				Region:            pe.GetRegionName(),
			})
		}
	}
	return result, nil
}

func buildNetworkPeering(npProvider store.PeeringConnectionLister, projectID string) ([]atlasV1.NetworkPeer, error) {
	// pagination not required, max 25 entries per provider can be configured via API
	npListAWS, err := npProvider.PeeringConnections(projectID, &atlas.ContainersListOptions{
		ListOptions: atlas.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(provider.ProviderAWS),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for AWS: %w", err)
	}

	npListGCP, err := npProvider.PeeringConnections(projectID, &atlas.ContainersListOptions{
		ListOptions: atlas.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(provider.ProviderGCP),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for GCP: %w", err)
	}

	npListAzure, err := npProvider.PeeringConnections(projectID, &atlas.ContainersListOptions{
		ListOptions: atlas.ListOptions{
			ItemsPerPage: MaxItems,
		},
		ProviderName: string(provider.ProviderAzure),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting network peering connections for Azure: %w", err)
	}

	result := make([]atlasV1.NetworkPeer, 0, len(npListAWS)+len(npListGCP)+len(npListAzure))

	for i := range npListAWS {
		np := npListAWS[i]
		result = append(result, convertNetworkPeer(np, provider.ProviderAWS))
	}

	for i := range npListGCP {
		np := npListGCP[i]
		result = append(result, convertNetworkPeer(np, provider.ProviderGCP))
	}

	for i := range npListAzure {
		np := npListAzure[i]
		result = append(result, convertNetworkPeer(np, provider.ProviderAzure))
	}

	return result, nil
}

func convertNetworkPeer(np atlasv2.BaseNetworkPeeringConnectionSettings, providerName provider.ProviderName) atlasV1.NetworkPeer {
	switch np.GetProviderName() {
	case "AWS":
		return convertAWSNetworkPeer(&np, providerName)
	case "GCP":
		return convertGCPNetworkPeer(&np, providerName)
	case "Azure":
		return convertAzureNetworkPeer(&np, providerName)
	default:
		return atlasV1.NetworkPeer{}
	}
}

func convertAWSNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName provider.ProviderName) atlasV1.NetworkPeer {
	return atlasV1.NetworkPeer{
		AccepterRegionName:  np.GetAccepterRegionName(),
		AWSAccountID:        np.GetAwsAccountId(),
		ContainerRegion:     "",
		ContainerID:         np.ContainerId,
		ProviderName:        providerName,
		RouteTableCIDRBlock: np.GetRouteTableCidrBlock(),
		VpcID:               np.GetVpcId(),
	}
}

func convertAzureNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName provider.ProviderName) atlasV1.NetworkPeer {
	return atlasV1.NetworkPeer{
		AzureDirectoryID:    np.GetAzureDirectoryId(),
		AzureSubscriptionID: np.GetAzureSubscriptionId(),
		ContainerRegion:     "",
		ContainerID:         np.GetContainerId(),
		ProviderName:        providerName,
		ResourceGroupName:   np.GetResourceGroupName(),
		VNetName:            np.GetVnetName(),
	}
}

func convertGCPNetworkPeer(np *atlasv2.BaseNetworkPeeringConnectionSettings, providerName provider.ProviderName) atlasV1.NetworkPeer {
	return atlasV1.NetworkPeer{
		GCPProjectID:    np.GetGcpProjectId(),
		ContainerRegion: "",
		ContainerID:     np.ContainerId,
		ProviderName:    providerName,
		NetworkName:     np.GetNetworkName(),
	}
}

func buildEncryptionAtRest(encProvider store.EncryptionAtRestDescriber, projectID string) (*atlasV1.EncryptionAtRest, error) {
	data, err := encProvider.EncryptionAtRest(projectID)
	if err != nil {
		return nil, err
	}

	return &atlasV1.EncryptionAtRest{
		AwsKms: atlasV1.AwsKms{
			Enabled:             data.AwsKms.Enabled,
			AccessKeyID:         data.AwsKms.GetAccessKeyID(),
			SecretAccessKey:     data.AwsKms.GetSecretAccessKey(),
			CustomerMasterKeyID: data.AwsKms.GetCustomerMasterKeyID(),
			Region:              data.AwsKms.GetRegion(),
			RoleID:              data.AwsKms.GetRoleId(),
			Valid:               data.AwsKms.Valid,
		},
		AzureKeyVault: atlasV1.AzureKeyVault{
			Enabled:           data.AzureKeyVault.Enabled,
			ClientID:          data.AzureKeyVault.GetClientID(),
			AzureEnvironment:  data.AzureKeyVault.GetAzureEnvironment(),
			SubscriptionID:    data.AzureKeyVault.GetSubscriptionID(),
			ResourceGroupName: data.AzureKeyVault.GetResourceGroupName(),
			KeyVaultName:      data.AzureKeyVault.GetKeyVaultName(),
			KeyIdentifier:     data.AzureKeyVault.GetKeyIdentifier(),
			Secret:            data.AzureKeyVault.GetSecret(),
			TenantID:          data.AzureKeyVault.GetTenantID(),
		},
		GoogleCloudKms: atlasV1.GoogleCloudKms{
			Enabled:              data.GoogleCloudKms.Enabled,
			ServiceAccountKey:    data.GoogleCloudKms.GetServiceAccountKey(),
			KeyVersionResourceID: data.GoogleCloudKms.GetKeyVersionResourceID(),
		},
	}, nil
}

func buildCloudProviderAccessRoles(cpaProvider store.CloudProviderAccessRoleLister, projectID string) ([]atlasV1.CloudProviderAccessRole, error) {
	data, err := cpaProvider.CloudProviderAccessRoles(projectID)
	if err != nil {
		return nil, err
	}

	var result []atlasV1.CloudProviderAccessRole
	for i := range data.AwsIamRoles {
		cpa := &data.AwsIamRoles[i]
		result = append(result, atlasV1.CloudProviderAccessRole{
			ProviderName:      cpa.ProviderName,
			IamAssumedRoleArn: cpa.GetIamAssumedRoleArn(),
		})
	}

	return result, nil
}

func buildProjectSettings(psProvider store.ProjectSettingsDescriber, projectID string) (*atlasV1.ProjectSettings, error) {
	data, err := psProvider.ProjectSettings(projectID)
	if err != nil {
		return nil, err
	}

	return &atlasV1.ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: data.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsDataExplorerEnabled:                       data.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 data.IsPerformanceAdvisorEnabled,
		IsRealtimePerformancePanelEnabled:           data.IsRealtimePerformancePanelEnabled,
		IsSchemaAdvisorEnabled:                      data.IsSchemaAdvisorEnabled,
	}, nil
}

func buildAuditing(auditingProvider store.AuditingDescriber, projectID string) (*atlasV1.Auditing, error) {
	data, err := auditingProvider.Auditing(projectID)
	if err != nil {
		return nil, err
	}

	return &atlasV1.Auditing{
		AuditAuthorizationSuccess: &data.AuditAuthorizationSuccess,
		AuditFilter:               data.AuditFilter,
		Enabled:                   &data.Enabled,
	}, nil
}

func buildAlertConfigurations(acProvider store.AlertConfigurationLister, projectID string) ([]atlasV1.AlertConfiguration, error) {
	data, err := acProvider.AlertConfigurations(projectID, &atlas.ListOptions{
		ItemsPerPage: MaxItems,
	})
	if err != nil {
		return nil, err
	}
	var result []atlasV1.AlertConfiguration

	convertMatchers := func(atlasMatcher []atlas.Matcher) []atlasV1.Matcher {
		var res []atlasV1.Matcher
		for _, m := range atlasMatcher {
			res = append(res, atlasV1.Matcher{
				FieldName: m.FieldName,
				Operator:  m.Operator,
				Value:     m.Value,
			})
		}
		return res
	}

	convertMetricThreshold := func(atlasMT *atlas.MetricThreshold) *atlasV1.MetricThreshold {
		if atlasMT == nil {
			return &atlasV1.MetricThreshold{}
		}
		return &atlasV1.MetricThreshold{
			MetricName: atlasMT.MetricName,
			Operator:   atlasMT.Operator,
			Threshold:  fmt.Sprintf("%f", atlasMT.Threshold),
			Units:      atlasMT.Units,
			Mode:       atlasMT.Mode,
		}
	}

	convertThreshold := func(atlasT *atlas.Threshold) *atlasV1.Threshold {
		if atlasT == nil {
			return &atlasV1.Threshold{}
		}
		return &atlasV1.Threshold{
			Operator:  atlasT.Operator,
			Units:     atlasT.Units,
			Threshold: fmt.Sprintf("%f", atlasT.Threshold),
		}
	}

	convertNotifications := func(atlasN []atlas.Notification) []atlasV1.Notification {
		var res []atlasV1.Notification
		for i := range atlasN {
			n := &atlasN[i]
			res = append(res, atlasV1.Notification{
				APIToken:            n.APIToken,
				ChannelName:         n.ChannelName,
				DatadogAPIKey:       n.DatadogAPIKey,
				DatadogRegion:       n.DatadogRegion,
				DelayMin:            n.DelayMin,
				EmailAddress:        n.EmailAddress,
				EmailEnabled:        n.EmailEnabled,
				FlowdockAPIToken:    n.FlowdockAPIToken,
				FlowName:            n.FlowName,
				IntervalMin:         n.IntervalMin,
				MobileNumber:        n.MobileNumber,
				OpsGenieAPIKey:      n.OpsGenieAPIKey,
				OpsGenieRegion:      n.OpsGenieRegion,
				OrgName:             n.OrgName,
				ServiceKey:          n.ServiceKey,
				SMSEnabled:          n.SMSEnabled,
				TeamID:              n.TeamID,
				TeamName:            n.TeamName,
				TypeName:            n.TypeName,
				Username:            n.Username,
				VictorOpsAPIKey:     n.VictorOpsAPIKey,
				VictorOpsRoutingKey: n.VictorOpsRoutingKey,
				Roles:               n.Roles,
			})
		}
		return res
	}

	for i := range data {
		alertConfig := &data[i]
		result = append(result, atlasV1.AlertConfiguration{
			EventTypeName:   alertConfig.EventTypeName,
			Enabled:         pointer.GetOrDefault(alertConfig.Enabled, false),
			Matchers:        convertMatchers(alertConfig.Matchers),
			MetricThreshold: convertMetricThreshold(alertConfig.MetricThreshold),
			Threshold:       convertThreshold(alertConfig.Threshold),
			Notifications:   convertNotifications(alertConfig.Notifications),
		})
	}

	return result, nil
}

func buildTeams(teamsProvider store.AtlasOperatorTeamsStore, orgID, projectID, projectName, targetNamespace, version string, dictionary map[string]string) ([]atlasV1.Team, []*atlasV1.AtlasTeam, error) {
	pt, err := teamsProvider.ProjectTeams(projectID)
	if err != nil {
		return nil, nil, err
	}

	projectTeams, ok := pt.(*atlas.TeamsAssigned)
	if !ok {
		return nil, nil, ErrTeamsAssigned
	}

	fetchUsers := func(teamID string) ([]string, error) {
		assignedUsers, err := teamsProvider.TeamUsers(orgID, teamID)
		if err != nil {
			return nil, err
		}
		users, ok := assignedUsers.([]atlas.AtlasUser)
		if !ok {
			return nil, ErrTeamUsers
		}
		result := make([]string, 0, len(users))
		for i := range users {
			result = append(result, users[i].Username)
		}
		return result, nil
	}

	convertRoleNames := func(input []string) []atlasV1.TeamRole {
		if len(input) == 0 {
			return nil
		}
		result := make([]atlasV1.TeamRole, 0, len(input))
		for i := range input {
			result = append(result, atlasV1.TeamRole(input[i]))
		}
		return result
	}

	convertUserNames := func(input []string) []atlasV1.TeamUser {
		if len(input) == 0 {
			return nil
		}

		result := make([]atlasV1.TeamUser, 0, len(input))
		for i := range input {
			result = append(result, atlasV1.TeamUser(input[i]))
		}
		return result
	}

	teamsRefs := make([]atlasV1.Team, 0, len(projectTeams.Results))
	atlasTeamCRs := make([]*atlasV1.AtlasTeam, 0, len(projectTeams.Results))

	for i := range projectTeams.Results {
		teamRef := projectTeams.Results[i]

		if teamRef == nil {
			continue
		}

		team, err := teamsProvider.TeamByID(orgID, teamRef.TeamID)
		if err != nil {
			return nil, nil, fmt.Errorf("team id: %s is assigned to project %s (id: %s) but not found. %w",
				teamRef.TeamID, projectName, projectID, err)
		}

		crName := resources.NormalizeAtlasName(fmt.Sprintf("%s-team-%s", projectName, team.Name), dictionary)
		teamsRefs = append(teamsRefs, atlasV1.Team{
			TeamRef: common.ResourceRefNamespaced{
				Name:      crName,
				Namespace: targetNamespace,
			},
			Roles: convertRoleNames(teamRef.RoleNames),
		})

		users, err := fetchUsers(team.ID)
		if err != nil {
			return nil, nil, err
		}

		atlasTeamCRs = append(atlasTeamCRs, &atlasV1.AtlasTeam{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasTeam",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      crName,
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: version,
				},
			},
			Spec: atlasV1.TeamSpec{
				Name:      team.Name,
				Usernames: convertUserNames(users),
			},
			Status: status.TeamStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		})
	}

	return teamsRefs, atlasTeamCRs, nil
}
