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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/advancedsettings"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/availableregions"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/connectionstring"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/indexes"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/onlinearchive"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/clusters/sampledata"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201006/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	labelKey           = "Infrastructure Tool"
	atlasCLILabelValue = "Atlas CLI"
	mongoCLILabelValue = "mongoCLI"
)

// MongoCLIBuilder is to split "mongocli atlas clusters" and "atlas clusters".
func MongoCLIBuilder() *cobra.Command {
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
		CreateBuilder(),
		WatchBuilder(),
		UpdateBuilder(),
		PauseBuilder(),
		StartBuilder(),
		DeleteBuilder(),
		LoadSampleDataBuilder(false),
		indexes.Builder(),
		search.Builder(),
		onlinearchive.Builder(),
		connectionstring.Builder(),
	)

	return cmd
}

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

func NewCLILabel() atlasv2.ComponentLabel {
	labelValue := atlasCLILabelValue
	if config.ToolName == config.MongoCLI {
		labelValue = mongoCLILabelValue
	}

	return atlasv2.ComponentLabel{
		Key:   pointer.Get(labelKey),
		Value: pointer.Get(labelValue),
	}
}

func AddLabel(out *atlasv2.AdvancedClusterDescription, l atlasv2.ComponentLabel) {
	if LabelExists(out.Labels, l) {
		return
	}

	out.Labels = append(out.Labels, l)
}

func LabelExists(labels []atlasv2.ComponentLabel, l atlasv2.ComponentLabel) bool {
	for _, v := range labels {
		if v.GetKey() == l.GetKey() && v.GetValue() == l.GetValue() {
			return true
		}
	}
	return false
}

func RemoveReadOnlyAttributes(out *atlasv2.AdvancedClusterDescription) {
	out.Id = nil
	out.CreateDate = nil
	out.StateName = nil
	out.MongoDBVersion = nil
	out.ConnectionStrings = nil
	isTenant := false

	for i, spec := range out.ReplicationSpecs {
		out.ReplicationSpecs[i].Id = nil
		for j, config := range spec.RegionConfigs {
			out.ReplicationSpecs[i].RegionConfigs[j].ProviderName = config.ProviderName
			if config.GetProviderName() == tenant {
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

func RemoveReadOnlyAttributesSharedCluster(out *atlas.Cluster) {
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

func SharedLabelExists(labels []atlas.Label, l atlasv2.ComponentLabel) bool {
	for _, v := range labels {
		if v.Key == l.GetKey() && v.Value == l.GetValue() {
			return true
		}
	}
	return false
}

func AddLabelSharedCluster(out *atlas.Cluster, l atlasv2.ComponentLabel) {
	if SharedLabelExists(out.Labels, l) {
		return
	}

	l2 := atlas.Label{
		Key:   l.GetKey(),
		Value: l.GetValue(),
	}
	out.Labels = append(out.Labels, l2)
}
