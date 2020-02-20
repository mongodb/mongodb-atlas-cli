package convert

import (
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

//Public constants
const (
	AdminDB = "admin"
)

//Private constants
const (
	roleSep = "@"
)

// BuildRoles converts the roles inside the array of string in an array of Atlas.Role Objects
// r contains roles in the format roleName@dbName
func BuildRoles(r []string) []atlas.Role {
	roles := make([]atlas.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := AdminDB
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
