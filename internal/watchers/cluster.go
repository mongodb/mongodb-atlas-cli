package watchers

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
)

const (
	clusterDeleting  = "DELETING"
	clusterUpdating  = "UPDATING"
	clusterIdle      = "IDLE"
	clusterCreating  = "CREATING"
	clusterRepairing = "REPAIRING"

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

var ClusterUpdated = &StateTransition{
	StartState: pointer.Get(clusterUpdating),
	EndState:   pointer.Get(clusterIdle),
}

var ClusterUpgraded = &StateTransition{
	StartState:          pointer.Get(clusterUpdating),
	EndState:            pointer.Get(clusterIdle),
	RetryableErrorCodes: []string{clusterNotFound},
}

type AtlasClusterStatusDescriber struct {
	store       store.AtlasClusterDescriber
	projectID   string
	clusterName string
}

func (describer *AtlasClusterStatusDescriber) GetStatus() (string, error) {
	result, err := describer.store.AtlasCluster(describer.projectID, describer.clusterName)
	if result != nil && result.StateName != nil {
		return *result.StateName, err
	}

	return "", err
}

func NewAtlasClusterStatusDescriber(s store.AtlasClusterDescriber, projectID, clusterName string) *AtlasClusterStatusDescriber {
	return &AtlasClusterStatusDescriber{
		store:       s,
		projectID:   projectID,
		clusterName: clusterName,
	}
}
