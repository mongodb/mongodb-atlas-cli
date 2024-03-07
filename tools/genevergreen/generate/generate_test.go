// Copyright 2022 MongoDB Inc
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

package generate

import (
	"testing"

	"github.com/evergreen-ci/shrub"
	"github.com/stretchr/testify/assert"
)

func TestPublishSnapshotTasks(t *testing.T) {
	t.Run(mongocli, func(t *testing.T) {
		c := &shrub.Configuration{}
		PublishSnapshotTasks(c, mongocli)
		assert.Len(t, c.Tasks, 34)
		assert.Len(t, c.Variants, 2)
	})
}

func TestPublishStableTasks(t *testing.T) {
	t.Run(mongocli, func(t *testing.T) {
		c := &shrub.Configuration{}
		PublishStableTasks(c, mongocli)
		assert.Len(t, c.Variants, 4)
		assert.Len(t, c.Tasks, 136)
	})
}
