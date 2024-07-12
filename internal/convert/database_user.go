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
	"time"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	AdminDB             = "admin"
	ExternalAuthDB      = "$external"
	roleSep             = "@"
	scopeSep            = ":"
	collectionSep       = "."
	defaultUserDatabase = "admin"
	userLdapAuthType    = "USER"
	defaultResourceType = "CLUSTER"
)

// BuildAtlasRoles converts the roles inside the array of string in an array of mongodbatlas.DatabaseUserRole structs.
// r contains roles in the format roleName@dbName.
func BuildAtlasRoles(r []string) []atlasv2.DatabaseUserRole {
	roles := make([]atlasv2.DatabaseUserRole, len(r))
	for i, roleP := range r {
		roleName, databaseName := splitRoleAndDBName(roleP)
		dbCollection := strings.Split(databaseName, collectionSep)
		databaseName = dbCollection[0]
		roles[i] = atlasv2.DatabaseUserRole{
			RoleName:       roleName,
			DatabaseName:   databaseName,
			CollectionName: buildCollectionName(dbCollection),
		}
	}
	return roles
}

func buildCollectionName(dbCollection []string) *string {
	if len(dbCollection) > 1 {
		c := strings.Join(dbCollection[1:], ".")
		return &c
	}
	return nil
}

func ParseDeleteAfter(deleteAfter string) *time.Time {
	deleteAfterDate, err := ParseTimestamp(deleteAfter)
	if err == nil {
		return &deleteAfterDate
	}
	return nil
}

func splitRoleAndDBName(roleAndDBNAme string) (role, dbName string) {
	rd := strings.Split(roleAndDBNAme, roleSep)
	dbName = defaultUserDatabase
	role = rd[0]
	if len(rd) > 1 {
		dbName = rd[1]
	}
	return
}

// BuildAtlasScopes converts the scopes inside the array of string in an array of mongodbatlas.Scope structs.
// r contains resources in the format resourceName:resourceType.
func BuildAtlasScopes(r []string) []atlasv2.UserScope {
	scopes := make([]atlasv2.UserScope, len(r))
	for i, scopeP := range r {
		scope := strings.Split(scopeP, scopeSep)
		resourceType := defaultResourceType
		if len(scope) > 1 {
			resourceType = scope[1]
		}

		scopes[i] = atlasv2.UserScope{
			Name: scope[0],
			Type: strings.ToUpper(resourceType),
		}
	}
	return scopes
}

// GetAuthDB determines the authentication database based on the type of user.
// LDAP, X509 and AWSIAM should all use $external.
// SCRAM-SHA should use admin.
func GetAuthDB(user *atlasv2.CloudDatabaseUser) string {
	// base documentation https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/database_user
	_, isX509 := adminX509Type[user.GetX509Type()]
	_, isIAM := awsIAMType[user.GetAwsIAMType()]

	// just USER is external
	isLDAP := user.LdapAuthType != nil && *user.LdapAuthType == userLdapAuthType

	if isX509 || isIAM || isLDAP {
		return ExternalAuthDB
	}
	return defaultUserDatabase
}

var adminX509Type = map[string]struct{}{
	"MANAGED":  {},
	"CUSTOMER": {},
}

var awsIAMType = map[string]struct{}{
	"USER": {},
	"ROLE": {},
}
