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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/cobra"
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
		Long:       "The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters.",
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
		LoadSampleDataBuilder(),
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
		Long:       "The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters.",
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
		LoadSampleDataBuilder(),
		UpgradeBuilder(),
		indexes.Builder(),
		search.Builder(),
		onlinearchive.Builder(),
		connectionstring.Builder(),
		availableregions.Builder(),
	)

	return cmd
}

func NewCLILabel() atlas.Label {
	labelValue := atlasCLILabelValue
	if config.ToolName == config.MongoCLI {
		labelValue = mongoCLILabelValue
	}

	return atlas.Label{
		Key:   labelKey,
		Value: labelValue,
	}
}

func AddLabel(out *atlas.AdvancedCluster, l atlas.Label) {
	if LabelExists(out.Labels, l) {
		return
	}

	out.Labels = append(out.Labels, l)
}

func LabelExists(labels []atlas.Label, l atlas.Label) bool {
	for _, v := range labels {
		if v.Key == l.Key && v.Value == l.Value {
			return true
		}
	}
	return false
}

func RemoveReadOnlyAttributes(out *atlas.AdvancedCluster) {
	out.ID = ""
	out.CreateDate = ""
	out.StateName = ""
	out.MongoDBVersion = ""
	out.ConnectionStrings = nil
	isTenant := false
	for _, spec := range out.ReplicationSpecs {
		spec.ID = ""
		for _, c := range spec.RegionConfigs {
			if c.ProviderName == tenant {
				isTenant = true
				break
			}
		}
	}
	if isTenant {
		out.BiConnector = nil
		out.EncryptionAtRestProvider = ""
		out.DiskSizeGB = nil
		out.MongoDBMajorVersion = ""
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

func AddLabelSharedCluster(out *atlas.Cluster, l atlas.Label) {
	if LabelExists(out.Labels, l) {
		return
	}

	out.Labels = append(out.Labels, l)
}
