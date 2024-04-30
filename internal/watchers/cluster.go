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

package watchers

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
)

const (
	clusterDeleting = "DELETING"
	clusterUpdating = "UPDATING"
	clusterIdle     = "IDLE"
	clusterCreating = "CREATING"

	clusterNotFound = "CLUSTER_NOT_FOUND"
)

var ClusterDeleted = &StateTransition{
	StartState:   pointer.Get(clusterDeleting),
	EndErrorCode: pointer.Get(clusterNotFound),
}

var ClusterCreated = &StateTransition{
	StartState: pointer.Get(clusterCreating),
	EndState:   pointer.Get(clusterIdle),
}

type AtlasClusterStateDescriber struct {
	store       store.ClusterDescriber
	projectID   string
	clusterName string
}

func (describer *AtlasClusterStateDescriber) GetState() (string, error) {
	result, err := describer.store.AtlasCluster(describer.projectID, describer.clusterName)
	if result != nil && result.StateName != nil {
		return *result.StateName, err
	}

	return "", err
}

func NewAtlasClusterStateDescriber(s store.ClusterDescriber, projectID, clusterName string) *AtlasClusterStateDescriber {
	return &AtlasClusterStateDescriber{
		store:       s,
		projectID:   projectID,
		clusterName: clusterName,
	}
}
