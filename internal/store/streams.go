// Copyright 2023 MongoDB Inc
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
	"go.mongodb.org/atlas/mongodbatlas"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

// Stub types until I get the actual generated types
type StreamProcessorInstance struct {
	ID                string                         `json:"id,omitempty"`
	Name              string                         `json:"name,omitempty"`
	GroupID           string                         `json:"groupId,omitempty"`
	DataProcessRegion mongodbatlas.DataProcessRegion `json:"dataProcessRegion,omitempty"`
	Created           string                         `json:"created,omitempty"`
	LastUpdated       string                         `json:"lastUpdated,omitempty"`
	SRV               string                         `json:"standardSrv,omitempty"`
}

type StreamsList struct {
	Results []StreamProcessorInstance
}

//go:generate mockgen -destination=../mocks/mock_streams.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store StreamsLister,StreamsDescriber,StreamsCreator,StreamsDeleter,StreamsUpdater

type StreamsLister interface {
	ProjectStreams(string, *atlas.ListOptions) (interface{}, error)
}

type StreamsDescriber interface {
	AtlasStream(string, string) (*StreamProcessorInstance, error)
}

type StreamsCreator interface {
	CreateStream(*StreamProcessorInstance) (*StreamProcessorInstance, error)
}

type StreamsDeleter interface {
	DeleteStream(string, string) error
}

type StreamsUpdater interface {
	UpdateStream(string, string, *StreamProcessorInstance) (*StreamProcessorInstance, error)
}

// ProjectStreams encapsulate the logic to list all the streams of a given project
// Probably shouldn't be interface{}, should be something with Results & the list properties
func (s *Store) ProjectStreams(projectID string, opts *atlas.ListOptions) (interface{}, error) {
	// We get
	// opts.PageNum
	// opts.ItemsPerPage
	// projectId

	streamProcessor := new(StreamProcessorInstance)
	streamProcessor.ID = "SampleId"
	streamProcessor.Name = "Processor1"
	streamProcessor.DataProcessRegion = mongodbatlas.DataProcessRegion{CloudProvider: "AWS", Region: "US-EAST-1"}

	streamProcessor2 := new(StreamProcessorInstance)
	streamProcessor2.ID = "SampleId2"
	streamProcessor2.Name = "Processor2"
	streamProcessor2.DataProcessRegion = mongodbatlas.DataProcessRegion{CloudProvider: "AZURE", Region: "SYDNEY"}

	return &StreamsList{Results: []StreamProcessorInstance{*streamProcessor, *streamProcessor2}}, nil
}

func (s *Store) AtlasStream(projectId, name string) (*StreamProcessorInstance, error) {
	streamProcessor := new(StreamProcessorInstance)
	streamProcessor.ID = "SampleId"
	streamProcessor.Name = name
	streamProcessor.DataProcessRegion = mongodbatlas.DataProcessRegion{CloudProvider: "AWS", Region: "US-EAST-1"}

	return streamProcessor, nil
}

func (s *Store) CreateStream(processor *StreamProcessorInstance) (*StreamProcessorInstance, error) {
	return processor, nil
}

func (s *Store) DeleteStream(projectId, name string) error {
	return nil
}

func (s *Store) UpdateStream(projectID, name string, processor *StreamProcessorInstance) (*StreamProcessorInstance, error) {
	return processor, nil
}
