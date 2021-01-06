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
package dbusers

const (
	dbUsers        = "Manage database users for your project."
	createDBUser   = "Create a database user for your project."
	deleteDBUser   = "Delete a database user for your project."
	listDBUsers    = "List Atlas database users for your project."
	describeDBUser = "Return a single Atlas database user for your project."
	updateDBUser   = "Update a database user for your project."
	dbUsersLong    = `The dbusers command retrieves, creates and modifies the MongoDB database users in your project.
Each user has a set of roles that provide access to the project’s databases. 
A user’s roles apply to all the clusters in the project.`
)
