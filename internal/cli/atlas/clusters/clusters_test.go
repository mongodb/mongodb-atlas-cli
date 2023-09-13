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

//go:build unit

package clusters

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201007/admin"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		18,
		[]string{},
	)
}

func TestMongoCLIBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		MongoCLIBuilder(),
		13,
		[]string{},
	)
}

func TestAddLabel(t *testing.T) {
	type args struct {
		out   *atlasv2.AdvancedClusterDescription
		label atlasv2.ComponentLabel
	}
	tests := []struct {
		name     string
		args     args
		wantsAdd bool
	}{
		{
			name: "adds",
			args: args{
				out: &atlasv2.AdvancedClusterDescription{
					Labels: []atlasv2.ComponentLabel{},
				},
				label: atlasv2.ComponentLabel{Key: pointer.Get("test"), Value: pointer.Get("test")},
			},
			wantsAdd: true,
		},
		{
			name: "doesn't adds",
			args: args{
				out: &atlasv2.AdvancedClusterDescription{
					Labels: []atlasv2.ComponentLabel{{Key: pointer.Get("test"), Value: pointer.Get("test")}},
				},
				label: atlasv2.ComponentLabel{Key: pointer.Get("test"), Value: pointer.Get("test")},
			},
			wantsAdd: true,
		},
	}
	for _, tt := range tests {
		name := tt.name
		args := tt.args
		wantsAdd := tt.wantsAdd
		t.Run(name, func(t *testing.T) {
			AddLabel(args.out, args.label)
			if exists := LabelExists(args.out.Labels, args.label); exists != wantsAdd {
				t.Errorf("wants to add %v, got %v\n", wantsAdd, exists)
			}
		})
	}
}

func TestLabelExists(t *testing.T) {
	type args struct {
		labels []atlasv2.ComponentLabel
		l      atlasv2.ComponentLabel
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "label doesn't exist",
			args: args{
				labels: []atlasv2.ComponentLabel{},
				l: atlasv2.ComponentLabel{
					Key:   pointer.Get("test"),
					Value: pointer.Get("test"),
				},
			},
			want: false,
		},
		{
			name: "label exist",
			args: args{
				labels: []atlasv2.ComponentLabel{
					{
						Key:   pointer.Get("test"),
						Value: pointer.Get("test"),
					},
				},
				l: atlasv2.ComponentLabel{
					Key:   pointer.Get("test"),
					Value: pointer.Get("test"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		name := tt.name
		args := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			if got := LabelExists(args.labels, args.l); got != want {
				t.Errorf("LabelExists() = %v, want %v", got, want)
			}
		})
	}
}

func TestRemoveReadOnlyAttributes(t *testing.T) {
	var (
		id        = "Test"
		testVar   = "test"
		specID    = "22"
		shards    = 2
		zone      = "1"
		timeStamp = time.Now()
	)
	tests := []struct {
		name string
		args *atlasv2.AdvancedClusterDescription
		want *atlasv2.AdvancedClusterDescription
	}{
		{
			name: "One AdvancedReplicationSpec",
			args: &atlasv2.AdvancedClusterDescription{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasv2.AdvancedClusterDescription{
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
		{
			name: "More AdvancedReplicationSpecs",
			args: &atlasv2.AdvancedClusterDescription{
				Id:             &id,
				MongoDBVersion: &testVar,
				StateName:      &testVar,
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						Id:        &specID,
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
				CreateDate: &timeStamp,
			},
			want: &atlasv2.AdvancedClusterDescription{
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
		{
			name: "Nothing to remove",
			args: &atlasv2.AdvancedClusterDescription{
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
			want: &atlasv2.AdvancedClusterDescription{
				ReplicationSpecs: []atlasv2.ReplicationSpec{
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
					{
						NumShards: &shards,
						ZoneName:  &zone,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		name := tt.name
		arg := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			RemoveReadOnlyAttributes(arg)
			if diff := deep.Equal(arg, want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
