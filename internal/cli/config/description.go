// Copyright 2020 MongoDB Inc
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
package config

const (
	short         = "Configure a profile to store access settings for your MongoDB deployment."
	setShort      = "Configure specific properties of a profile."
	renameShort   = "Rename a profile."
	deleteShort   = "Delete a profile."
	listShort     = "List available profiles."
	describeShort = "Return a specific profile"
	long          = `Configure settings in a user profile.
All settings are optional. You can specify settings individually by running: 
  $ mongocli config set --help 

You can also use environment variables (MCLI_*) when running the tool.
To find out more, see the documentation: https://docs.mongodb.com/mongocli/stable/configure/environment-variables/.`
	setLong = `Configure specific properties of the profile.
Available properties include: %v.`
)
