// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

import (
	"errors"
	"fmt"
)

func validateConfigOpts(opts *ConfigOpts) error {
	if opts.event == "" {
		return fmt.Errorf("--event flag is required")
	}
	if opts.notificationIntervalMin == 0 {
		return fmt.Errorf("--notificationIntervalMin is required")
	}
	if opts.notificationType == "" {
		return fmt.Errorf("--notificationType is required")
	}

	err := validateAlertSettingsTypes(opts)
	if err != nil {
		return err
	}
	return nil
}

//gocyclo:ignore
func validateAlertSettingsTypes(opts *ConfigOpts) error {
	switch opts.notificationType {
	case datadog:
		if opts.apiKey == "" || opts.notificationRegion == "" {
			return errors.New("--apiKey and --notificationRegion are required when --notificationType is DATADOG")
		}
	case email:
		if opts.notificationEmailAddress == "" {
			return errors.New("--notificationEmailAddress is required when --notificationType is EMAIL")
		}
	case microsoftTeams:
		if opts.notificationWebhookURL == "" {
			return errors.New("--notificationWebhookURL is required when --notificationType is MICROSOFT_TEAMS")
		}
	case opsGenie:
		if opts.apiKey == "" || opts.notificationRegion == "" {
			return errors.New("--apiKey and --notificationRegion are required when --notificationType is OPS_GENIE")
		}
	case pagerDuty:
		if opts.apiKey == "" {
			return errors.New("--apiKey is required when --notificationType is PAGER_DUTY")
		}
	case slack:
		if opts.notificationToken == "" || opts.notificationChannelName == "" {
			return errors.New("--notificationToken and --notificationChannelName are required when --notificationType is SLACK")
		}
	case sms:
		if opts.notificationMobileNumber == "" {
			return errors.New("--notificationMobileNumber is required when --notificationType is SMS")
		}
	case team:
		if opts.notificationTeamID == "" {
			return errors.New("--notificationTeamID is required when --notificationType is TEAM")
		}
	case user:
		if opts.notificationUsername == "" {
			return errors.New("--notificationUsername is required when --notificationType is USER")
		}
	case victor:
		if opts.apiKey == "" || opts.notificationVictorOpsRoutingKey == "" {
			return errors.New("--apiKey and --notificationVictorOpsRoutingKey are required when --notificationType is VICTOR_OPS")
		}
	case webhook:
		if opts.notificationWebhookURL == "" {
			return errors.New("--notificationWebhookURL is required when --notificationType is WEBHOOK")
		}
	}
	return nil
}
