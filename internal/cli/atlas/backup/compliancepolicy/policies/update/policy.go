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

package update

import (
	"errors"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

var policiesListTemplate = `
Your existing policies
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

func (opts *UpdateOpts) askPolicyOptions(compliancePolicy *atlasv2.DataProtectionSettings) (*atlasv2.DiskBackupApiPolicyItem, error) {

	policyItems := existingPolicyItems(compliancePolicy)
	if len(policyItems) == 0 {
		return nil, errors.New("no policy items found")
	}
	opts.printExistingPolicies(compliancePolicy)

	q := newPolicyItemSelectionQuestion(policyItems)
	var policyID string
	if err := telemetry.TrackAskOne(q, &policyID); err != nil {
		return nil, err
	}

	policyItem := opts.FindPolicyItemByID(policyItems, policyID)
	if policyItem == nil {
		return nil, errors.New("could not find the specified policy item")
	}
	return policyItem, nil
}

func (opts *UpdateOpts) printExistingPolicies(compliancePolicy *atlasv2.DataProtectionSettings) {
	opts.Template = policiesListTemplate
	opts.Print(compliancePolicy)
}

func existingPolicyItems(compliancePolicy *atlasv2.DataProtectionSettings) []atlasv2.DiskBackupApiPolicyItem {
	var items []atlasv2.DiskBackupApiPolicyItem
	scheduledPolicyItems, ok := compliancePolicy.GetScheduledPolicyItemsOk()
	if ok {
		items = append(items, scheduledPolicyItems...)
	}

	onDemandItem, ok := compliancePolicy.GetOnDemandPolicyItemOk()
	if ok {
		items = append(items, *onDemandItem)
	}
	return items
}

func extractIDs(items []atlasv2.DiskBackupApiPolicyItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.GetId())
	}
	return ids
}

func newPolicyItemSelectionQuestion(policyItems []atlasv2.DiskBackupApiPolicyItem) survey.Prompt {
	ids := extractIDs(policyItems)

	return &survey.Select{
		Message: "Select a policy to edit",
		Options: ids,
		Filter: func(filter string, _ string, i int) bool {
			filter = strings.ToLower(filter)
			return strings.HasPrefix(ids[i], filter)
		},
	}
}

func (opts *UpdateOpts) FindPolicyItemByID(items []atlasv2.DiskBackupApiPolicyItem, id string) *atlasv2.DiskBackupApiPolicyItem {
	for _, item := range items {
		if item.GetId() == id {
			return &item
		}
	}
	return nil
}
