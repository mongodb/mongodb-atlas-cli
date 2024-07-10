// Copyright 2021 MongoDB Inc
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

package convert

import (
	"strings"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	resourceSep = "."
)

// BuildAtlasInheritedRoles converts the inherited roles inside the array of string in an array of atlas.InheritedRole structs
// r contains roles in the format roleName@dbName.
func BuildAtlasInheritedRoles(r []string) []atlasv2.DatabaseInheritedRole {
	roles := make([]atlasv2.DatabaseInheritedRole, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := defaultUserDatabase
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlasv2.DatabaseInheritedRole{
			Db:   databaseName,
			Role: roleName,
		}
	}
	return roles
}

// BuildAtlasActions converts the actions inside the array of string in an array of atlas.Action structs
// r contains roles in the format action[@dbName.collection].
func BuildAtlasActions(a []string) []atlasv2.DatabasePrivilegeAction {
	actions := make([]atlasv2.DatabasePrivilegeAction, len(a))
	for i, actionP := range a {
		resourceStruct := atlasv2.DatabasePermittedNamespaceResource{}
		action := strings.Split(actionP, roleSep)
		actionName := action[0]
		if len(action) > 1 {
			resource := strings.Split(action[1], resourceSep)
			resourceStruct.Db = resource[0]
			if len(resource) > 1 {
				resourceStruct.Collection = strings.Join(resource[1:], resourceSep)
			}
		} else {
			resourceStruct.Cluster = true
		}

		actions[i] = atlasv2.DatabasePrivilegeAction{
			Action:    actionName,
			Resources: &[]atlasv2.DatabasePermittedNamespaceResource{resourceStruct},
		}
	}
	return actions
}
