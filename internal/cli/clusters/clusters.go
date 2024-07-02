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

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/advancedsettings"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/availableregions"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/connectionstring"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/indexes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/onlinearchive"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/sampledata"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func Builder() *cobra.Command {
	const use = "clusters"
	cmd := &cobra.Command{
		Use:        use,
		Aliases:    cli.GenerateAliases(use),
		SuggestFor: []string{"replicasets"},
		Short:      "Manage clusters for your project.",
		Long:       `The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters.`,
	}
	cmd.AddCommand(
		ListBuilder(),
		DescribeBuilder(),
		advancedsettings.Builder(),
		CreateBuilder(),
		WatchBuilder(),
		UpdateBuilder(),
		PauseBuilder(),
		StartBuilder(),
		DeleteBuilder(),
		LoadSampleDataBuilder(true),
		UpgradeBuilder(),
		FailoverBuilder(),
		indexes.Builder(),
		search.Builder(),
		onlinearchive.Builder(),
		connectionstring.Builder(),
		availableregions.Builder(),
		sampledata.Builder(),
	)

	return cmd
}

func addTags(out *atlasv2.AdvancedClusterDescription, tags map[string]string) {
	if len(tags) > 0 {
		var t []atlasv2.ResourceTag
		for k, v := range tags {
			if k == "" || v == "" {
				continue
			}
			key, value := k, v
			tag := atlasv2.ResourceTag{
				Key:   key,
				Value: value,
			}
			t = append(t, tag)
		}
		out.Tags = &t
	}
}

func removeReadOnlyAttributes(out *atlasv2.AdvancedClusterDescription) {
	out.Id = nil
	out.CreateDate = nil
	out.StateName = nil
	out.MongoDBVersion = nil
	out.ConnectionStrings = nil
	isTenant := false

	for i, spec := range out.GetReplicationSpecs() {
		(*out.ReplicationSpecs)[i].Id = nil
		for _, c := range spec.GetRegionConfigs() {
			if c.GetProviderName() == tenant {
				isTenant = true
				break
			}
		}
	}

	if isTenant {
		out.BiConnector = nil
		out.EncryptionAtRestProvider = nil
		out.DiskSizeGB = nil
		out.MongoDBMajorVersion = nil
		out.PitEnabled = nil
		out.BackupEnabled = nil
	}
}

func removeReadOnlyAttributesSharedCluster(out *atlas.Cluster) {
	out.ID = ""
	out.CreateDate = ""
	out.StateName = ""
	out.MongoDBVersion = ""
	out.ConnectionStrings = nil
	out.ReplicationSpec = nil
	out.MongoURI = ""
	out.MongoURIUpdated = ""
	out.MongoURIWithOptions = ""
	if out.ProviderSettings != nil {
		out.ProviderSettings.AutoScaling = nil
	}

	for _, spec := range out.ReplicationSpecs {
		spec.ID = ""
	}
}
