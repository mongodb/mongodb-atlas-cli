// Copyright 2021 MongoDB Inc
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

package convert

import (
	"strings"

	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	AdminDB             = "admin"
	ExternalAuthDB      = "$external"
	roleSep             = "@"
	scopeSep            = ":"
	defaultUserDatabase = "admin"
	defaultResourceType = "CLUSTER"
)

// BuildAtlasRoles converts the roles inside the array of string in an array of mongodbatlas.Role structs
// r contains roles in the format roleName@dbName
func BuildAtlasRoles(r []string) []atlas.Role {
	roles := make([]atlas.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := defaultUserDatabase
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}

// BuildOMRoles converts the roles inside the array of string in an array of opsmngr.Role structs
// r contains roles in the format roleName@dbName
func BuildOMRoles(r []string) []*opsmngr.Role {
	roles := make([]*opsmngr.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := defaultUserDatabase
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = &opsmngr.Role{
			Role:     roleName,
			Database: databaseName,
		}
	}
	return roles
}

// BuildAtlasScopes converts the scopes inside the array of string in an array of mongodbatlas.Scope structs
// r contains resources in the format resouceName@resourceType
func BuildAtlasScopes(r []string) []atlas.Scope {
	scopes := make([]atlas.Scope, len(r))
	for i, scopeP := range r {
		scope := strings.Split(scopeP, scopeSep)
		resourceType := defaultResourceType
		if len(scope) > 1 {
			resourceType = scope[1]
		}

		scopes[i] = atlas.Scope{
			Name: scope[0],
			Type: strings.ToUpper(resourceType),
		}
	}
	return scopes
}
