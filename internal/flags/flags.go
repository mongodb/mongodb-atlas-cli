// Copyright 2020 MongoDB Inc
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

package flags

const (
	Service                         = "service"                         // Service flag to set service
	Profile                         = "profile"                         // Profile flag to use a profile
	ProfileShort                    = "p"                               // ProfileShort flag to use a profile
	OrgID                           = "orgId"                           // OrgID flag to use an Organization ID
	ProjectID                       = "projectId"                       // ProjectID flag to use a project ID
	Provider                        = "provider"                        // Provider flag to set the cloud provider
	Region                          = "region"                          // Region flag
	RegionShort                     = "r"                               // RegionShort flag
	Members                         = "members"                         // Members flag
	MembersShort                    = "m"                               // MembersShort flag
	InstanceSize                    = "instanceSize"                    // InstanceSize flag
	DiskSizeGB                      = "diskSizeGB"                      // DiskSizeGB flag
	MDBVersion                      = "mdbVersion"                      // MDBVersion flag
	Backup                          = "backup"                          // Backup flag
	Username                        = "username"                        // Username flag
	Password                        = "password"                        // Password flag
	Email                           = "email"                           // Email flag
	FirstName                       = "firstName"                       // FirstName flag
	LastName                        = "lastName"                        // LastName flag
	Role                            = "role"                            // Role flag
	Type                            = "type"                            // Type flag
	Comment                         = "comment"                         // Comment flag
	Page                            = "page"                            // Page flag
	Limit                           = "limit"                           // Limit flag
	File                            = "file"                            // File flag
	FileShort                       = "f"                               // File flag
	Force                           = "force"                           // Force flag
	WhitelistIP                     = "whitelistIp"                     // WhitelistIP flag
	Event                           = "event"                           // EventTypeName flag
	Enabled                         = "enabled"                         // Enabled flag
	MatcherFieldName                = "matcherFieldName"                // MatcherFieldName flag
	MatcherOperator                 = "matcherOperator"                 // MatcherOperator flag
	MatcherValue                    = "matcherValue"                    // MatcherValue flag
	MetricThresholdMetricName       = "metricThresholdMetricName"       // MetricThresholdMetricName flag
	MetricThresholdOperator         = "metricThresholdOperator"         // MetricThresholdOperator flag
	MetricThresholdThreshold        = "metricThresholdThreshold"        // MetricThresholdThreshold flag
	MetricThresholdUnits            = "metricThresholdUnits"            // MetricThresholdUnits flag
	MetricThresholdMode             = "metricThresholdMode"             // MetricThresholdMode flag
	Token                           = "token"                           // Token flag
	NotificationChannelName         = "notificationsChannelName"        // NotificationChannelName flag
	APIKey                          = "apiKey"                          // APIKey flag
	NotificationRegion              = "notificationRegion"              // NotificationRegion flag
	NotificationDelayMin            = "notificationDelayMin"            // NotificationDelayMin flag
	NotificationEmailAddress        = "notificationEmailAddress"        // NotificationEmailAddress flag
	NotificationEmailEnabled        = "notificationEmailEnabled"        // NotificationEmailEnabled flag
	NotificationFlowName            = "notificationFlowName"            // NotificationFlowName flag
	NotificationIntervalMin         = "notificationIntervalMin"         // NotificationIntervalMin flag
	NotificationMobileNumber        = "notificationMobileNumber"        // NotificationMobileNumber flag
	NotificationOrgName             = "notificationOrgName"             // NotificationsOrgName flag
	NotificationServiceKey          = "notificationServiceKey"          // NotificationsServiceKey flag
	NotificationSmsEnabled          = "notificationSmsEnabled"          // NotificationsSmsEnabled flag
	NotificationTeamID              = "notificationTeamId"              // NotificationTeamID flag
	NotificationTypeName            = "notificationTypeName"            // NotificationTypeName flag
	NotificationUsername            = "notificationUsername"            // NotificationUsername flag
	NotificationVictorOpsRoutingKey = "notificationVictorOpsRoutingKey" // NotificationVictorOpsRoutingKey flag
)
