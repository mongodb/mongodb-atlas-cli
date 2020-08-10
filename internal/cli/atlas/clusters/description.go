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
package clusters

const (
	Clusters          = "Manage clusters for your project."
	ClustersLong      = "The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters."
	CreateCluster     = "Create a MongoDB cluster for your project."
	DeleteCluster     = "Delete a cluster from your project."
	DescribeCluster   = "Describe a cluster."
	ListClusters      = "List clusters for your project."
	UpdateCluster     = "Update a MongoDB cluster."
	PauseCluster      = "Pause a running MongoDB cluster in Atlas."
	StartCluster      = "Start a paused MongoDB cluster in Atlas."
	WatchCluster      = "Watch for a cluster to be available."
	Indexes           = "Manage cluster rolling indexes for your project."
	CreateIndex       = "Create a rolling index for your MongoDB cluster."
	CreateClusterLong = `You can create MongoDB clusters using this command.
To quickest way to get started is to just specify a name for your cluster and cloud provider and region to deploy, 
this will create a 3 member replica set with the latest available mongodb server version available.
Some of the cluster configuration options are available via flags but for full control of your deployment you can provide a config file.`
)
