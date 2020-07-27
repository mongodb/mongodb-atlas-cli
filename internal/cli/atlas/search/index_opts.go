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

package search

import (
	"fmt"
	"strings"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type IndexOpts struct {
	name           string
	dbName         string
	collection     string
	analyzer       string
	searchAnalyzer string
	dynamic        bool
	fields         []string
}

func (opts *IndexOpts) newSearchIndex() (*atlas.SearchIndex, error) {
	f, err := opts.indexFields()
	if err != nil {
		return nil, err
	}
	i := &atlas.SearchIndex{
		Analyzer:       opts.analyzer,
		CollectionName: opts.collection,
		Database:       opts.dbName,
		Mappings: &atlas.IndexMapping{
			Dynamic: opts.dynamic,
			Fields:  &f,
		},
		Name:           opts.name,
		SearchAnalyzer: opts.searchAnalyzer,
	}
	return i, nil
}

// indexFieldParts index field should be fieldName:analyzer:fieldType
const indexFieldParts = 3

func (opts *IndexOpts) indexFields() (map[string]atlas.IndexField, error) {
	if len(opts.fields) == 0 {
		return nil, nil
	}
	fields := make(map[string]atlas.IndexField, len(opts.fields))
	for _, p := range opts.fields {
		f := strings.Split(p, ":")
		if len(f) != indexFieldParts {
			return nil, fmt.Errorf("partition should be fieldName:analyzer:fieldType, got: %s", p)
		}
		fields[f[0]] = atlas.IndexField{
			Analyzer: f[1],
			Type:     f[2],
		}
	}
	return fields, nil
}
