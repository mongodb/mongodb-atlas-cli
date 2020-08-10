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
package logs

const (
	short        = "Download host logs for your project."
	download     = "Download a host mongodb logs."
	downloadLong = `To download a log you need the name of the host where the log files that you want to download are stored.
To see the hostnames of your Atlas cluster, visit the cluster overview page in the Atlas UI.
The name of the log file must be one of: mongodb.gz, mongos.gz, mongodb-audit-log.gz, mongos-audit-log.gz`
)
