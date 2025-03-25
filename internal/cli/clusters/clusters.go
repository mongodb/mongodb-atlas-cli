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
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/advancedsettings"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/availableregions"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/connectionstring"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/indexes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/onlinearchive"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/sampledata"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var errFailedToLoadClusterFileMessage = errors.New("failed to parse JSON file")

const (
	cannotUseFlexWithClusterApisErrorCode = "CANNOT_USE_FLEX_CLUSTER_IN_CLUSTER_API"
	deprecateMessageSharedTier            = "Deprecation note: the M2 and M5 tiers are now deprecated ('%s' was selected); when selecting M2 or M5, a FLEX tier will be created instead. For the migration guide, visit: https://dochub.mongodb.org/core/flex-migration.\n"
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

func addTags(out *atlasClustersPinned.AdvancedClusterDescription, tags map[string]string) {
	resourceTagsAtlasV2 := newResourceTags(tags)
	if resourceTagsAtlasV2 == nil {
		return
	}

	resourceTags := make([]atlasClustersPinned.ResourceTag, len(*resourceTagsAtlasV2))
	for i, v := range *resourceTagsAtlasV2 {
		resourceTags[i] = atlasClustersPinned.ResourceTag{
			Key:   v.Key,
			Value: v.Value,
		}
	}

	out.Tags = &resourceTags
}

func newResourceTags(tags map[string]string) *[]atlasv2.ResourceTag {
	if len(tags) == 0 {
		return nil
	}
	t := make([]atlasv2.ResourceTag, len(tags))
	i := 0
	for k, v := range tags {
		if k == "" || v == "" {
			continue
		}
		key, value := k, v
		tag := atlasv2.ResourceTag{
			Key:   key,
			Value: value,
		}
		t[i] = tag
		i++
	}

	return &t
}

func removeReadOnlyAttributes(out *atlasClustersPinned.AdvancedClusterDescription) {
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
