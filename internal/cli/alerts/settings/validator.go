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

func (opts *ConfigOpts) validateConfigOpts() error {
	if opts.event == "" {
		return fmt.Errorf("--%s flag is required", flag.Event)
	}
	if opts.notificationIntervalMin == 0 {
		return fmt.Errorf("--%s is required", flag.NotificationIntervalMin)
	}
	if opts.notificationType == "" {
		return fmt.Errorf("--%s is required", flag.NotificationType)
	}

	return opts.validateAlertSettingsTypes()
}

func (opts *ConfigOpts) validateAlertSettingsTypes() error {
	switch opts.notificationType {
	case datadog:
		return opts.validateDatadog()
	case email:
		return opts.validateEmail()
	case microsoftTeams:
		return opts.validateMicrosoftTeams()
	case opsGenie:
		return opts.validateOpsGenie()
	case pagerDuty:
		return opts.validatePagerDuty()
	case slack:
		return opts.validateSlack()
	case sms:
		return opts.validateSMS()
	case team:
		return opts.validateTeams()
	case user:
		return opts.validateUser()
	case victor:
		return validateVictor(opts)
	case webhook:
		return opts.validateWebhook()
	}
	return nil
}

func (opts *ConfigOpts) validateDatadog() error {
	if opts.apiKey == "" || opts.notificationRegion == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is DATADOG", flag.APIKey, flag.NotificationRegion, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateEmail() error {
	if opts.notificationEmailAddress == "" {
		return fmt.Errorf("--%s is required when --%s is EMAIL", flag.NotificationEmailAddress, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateMicrosoftTeams() error {
	if opts.notificationWebhookURL == "" {
		return fmt.Errorf("--%s is required when --%s is MICROSOFT_TEAMS", flag.NotificationWebhookURL, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateOpsGenie() error {
	if opts.apiKey == "" || opts.notificationRegion == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is OPS_GENIE", flag.APIKey, flag.NotificationRegion, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validatePagerDuty() error {
	if opts.notificationServiceKey == "" || opts.notificationRegion == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is PAGER_DUTY", flag.NotificationServiceKey, flag.NotificationRegion, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateSlack() error {
	if opts.notificationToken == "" || opts.notificationChannelName == "" {
		return fmt.Errorf("--%s and --%s are required when --%s is SLACK", flag.NotificationToken, flag.NotificationChannelName, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateSMS() error {
	if opts.notificationMobileNumber == "" {
		return fmt.Errorf("--%s is required when --%s is SMS", flag.NotificationMobileNumber, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateTeams() error {
	if opts.notificationTeamID == "" {
		return fmt.Errorf("--%s is required when --%s is TEAM", flag.NotificationTeamID, flag.NotificationType)
	}
	return nil
}

func (opts *ConfigOpts) validateUser() error {
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

func (opts *ConfigOpts) validateWebhook() error {
	if opts.notificationWebhookURL == "" {
		return fmt.Errorf("--%s is required when --%s is WEBHOOK", flag.NotificationWebhookURL, flag.NotificationType)
	}
	return nil
}
