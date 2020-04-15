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

package store

import (
	"context"
	"fmt"
	"io"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
)

type LogsDownloader interface {
	DownloadLog(string, string, string, io.Writer, *atlas.DateRangetOptions) error
}

// ProcessDisks encapsulate the logic to manage different cloud providers
func (s *Store) DownloadLog(groupID, host, name string, out io.Writer, opts *atlas.DateRangetOptions) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Logs.Get(context.Background(), groupID, host, name, out, opts)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
