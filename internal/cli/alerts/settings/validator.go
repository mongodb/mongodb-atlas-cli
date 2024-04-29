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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
)

func validateConfigOpts(opts *ConfigOpts) error {
	if opts.event == "" {
		return fmt.Errorf("--%s flag is required", flag.Event)
	}
	if opts.notificationIntervalMin == 0 {
		return fmt.Errorf("--%s is required", flag.NotificationIntervalMin)
	}
	if opts.notificationType == "" {
		return fmt.Errorf("--%s is required", flag.NotificationType)
	}

	return validateAlertSettingsTypes(opts)
}

func validateAlertSettingsTypes(opts *ConfigOpts) error {
	switch opts.notificationType {
	case datadog:
		return validateDatadog(opts)
	case email:
		return validateEmail(opts)
	case microsoftTeams:
		return validateMicrosoftTeams(opts)
	case opsGenie:
		return validateOpsGenie(opts)
	case pagerDuty:
		return validatePagerDuty(opts)
	case slack:
		return validateSlack(opts)
	case sms:
		return validateSMS(opts)
	case team:
		return validateTeams(opts)
	case user:
		return validateUser(opts)
	case victor:
		return validateVictor(opts)
	case webhook:
		return validateWebhook(opts)
	}
	return nil
}

func validateDatadog(opts *ConfigOpts) error {
	if opts.apiKey == "" || opts.notificationRegion == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is DATADOG", flag.APIKey, flag.NotificationRegion, flag.NotificationType)
	}
	return nil
}

func validateEmail(opts *ConfigOpts) error {
	if opts.notificationEmailAddress == "" {
		return fmt.Errorf("--%s is required when --%s is EMAIL", flag.NotificationEmailAddress, flag.NotificationType)
	}
	return nil
}

func validateMicrosoftTeams(opts *ConfigOpts) error {
	if opts.notificationWebhookURL == "" {
		return fmt.Errorf("--%s is required when --%s is MICROSOFT_TEAMS", flag.NotificationWebhookURL, flag.NotificationType)
	}
	return nil
}

func validateOpsGenie(opts *ConfigOpts) error {
	if opts.apiKey == "" || opts.notificationRegion == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is OPS_GENIE", flag.APIKey, flag.NotificationRegion, flag.NotificationType)
	}
	return nil
}

func validatePagerDuty(opts *ConfigOpts) error {
	if opts.apiKey == "" {
		return fmt.Errorf("--%s is required when --%s is PAGER_DUTY", flag.APIKey, flag.NotificationType)
	}
	return nil
}

func validateSlack(opts *ConfigOpts) error {
	if opts.notificationToken == "" || opts.notificationChannelName == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is SLACK", flag.NotificationToken, flag.NotificationChannelName, flag.NotificationType)
	}
	return nil
}

func validateSMS(opts *ConfigOpts) error {
	if opts.notificationMobileNumber == "" {
		return fmt.Errorf("--%s is required when --%s is SMS", flag.NotificationMobileNumber, flag.NotificationType)
	}
	return nil
}

func validateTeams(opts *ConfigOpts) error {
	if opts.notificationTeamID == "" {
		return fmt.Errorf("--%s is required when --%s is TEAM", flag.NotificationTeamID, flag.NotificationType)
	}
	return nil
}

func validateUser(opts *ConfigOpts) error {
	if opts.notificationUsername == "" {
		return fmt.Errorf("--%s is required when --%s is USER", flag.NotificationUsername, flag.NotificationType)
	}
	return nil
}

func validateVictor(opts *ConfigOpts) error {
	if opts.apiKey == "" || opts.notificationVictorOpsRoutingKey == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is VICTOR_OPS", flag.APIKey, flag.NotificationVictorOpsRoutingKey, flag.NotificationType)
	}
	return nil
}

func validateWebhook(opts *ConfigOpts) error {
	if opts.notificationWebhookURL == "" {
		return fmt.Errorf("--%s is required when --%s is WEBHOOK", flag.NotificationWebhookURL, flag.NotificationType)
	}
	return nil
}
