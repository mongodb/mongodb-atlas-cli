package watcher

import (
	"errors"

	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	active = "ACTIVE"
)

// CompliancePolicyWatcherFactory allows to init a compliance policy watcher in a functional way.
//
// Parameters
//
//   - projectID: The ID of the project for which the compliance policy details need to be fetched.
//   - store: An implementation of the CompliancePolicyDescriber interface, which is used to describe/fetch the compliance policy.
//   - policy: A pointer to a DataProtectionSettings object which will be updated with the fetched details.
func CompliancePolicyWatcherFactory(projectID string, store store.CompliancePolicyDescriber, policy *atlasv2.DataProtectionSettings) func() (bool, error) {
	return func() (bool, error) {
		res, err := store.DescribeCompliancePolicy(projectID)
		if err != nil {
			return false, err
		}
		policy = res
		if res.GetState() == "" {
			return false, errors.New("could not access State field")
		}
		return (res.GetState() == active), nil
	}
}
