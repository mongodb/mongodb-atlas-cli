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
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func (opts *UpdateOpts) askProjectOptions() (string, error) {
	res, err := opts.store.Projects(nil)
	if err != nil {
		return "", fmt.Errorf("unable to fetch projects: %w", err)
	}
	projects := res.GetResults()
	if len(projects) == 0 {
		return "", errors.New("you do not have any projects set up")
	}
	q := newProjectSelectionQuestion(projects)
	var projectID string
	if err := telemetry.TrackAskOne(q, &projectID); err != nil {
		return "", err
	}
	return projectID, nil
}

func extractNamesAndIDs(groups []atlasv2.Group) ([]string, []string) {
	var IDs []string
	var names []string
	for _, group := range groups {
		names = append(names, group.GetName())
		IDs = append(IDs, group.GetId())
	}
	return names, IDs
}

func newProjectSelectionQuestion(projects []atlasv2.Group) survey.Prompt {
	names, ids := extractNamesAndIDs(projects)
	return &survey.Select{
		Message: "Choose a project:",
		Options: ids,
		Description: func(_ string, i int) string {
			return names[i]
		},
		Filter: func(filter string, _ string, i int) bool {
			filter = strings.ToLower(filter)
			return strings.HasPrefix(strings.ToLower(names[i]), filter) || strings.HasPrefix(ids[i], filter)
		},
	}
}
