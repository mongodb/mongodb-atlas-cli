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
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

const DefaultAnalyzer = "lucene.standard"
const deprecatedFlagMessage = "please use --file instead"

type IndexOpts struct {
	Name           string
	DBName         string
	Collection     string
	Analyzer       string
	SearchAnalyzer string
	Dynamic        bool
	fields         []string
	Filename       string
	Fs             afero.Fs
}

func (opts *IndexOpts) validateOpts() error {
	if opts.Filename == "" {
		if !opts.Dynamic && len(opts.fields) == 0 {
			return errors.New("specify the fields to index for a static index or specify a dynamic index")
		}
		if opts.Dynamic && len(opts.fields) > 0 {
			return errors.New("do not specify --fields and --dynamic at the same time")
		}
	} else {
		if opts.Name != "" {
			return errors.New("do not specify --indexName and --file at the same time")
		}
		if opts.DBName != "" {
			return errors.New("do not specify --db and --file at the same time")
		}
		if opts.Collection != "" {
			return errors.New("do not specify --collection and --file at the same time")
		}
		if opts.Analyzer != DefaultAnalyzer {
			return errors.New("do not specify --analyzer and --file at the same time")
		}
		if opts.SearchAnalyzer != DefaultAnalyzer {
			return errors.New("do not specify --searchAnalyzer and --file at the same time")
		}
		if opts.Dynamic {
			return errors.New("do not specify --dynamic and --file at the same time")
		}
		if len(opts.fields) > 0 {
			return errors.New("do not specify --fields and --file at the same time")
		}
	}
	return nil
}

func (opts *IndexOpts) NewSearchIndex() (*atlasv2.ClusterSearchIndex, error) {
	if len(opts.Filename) > 0 {
		index := &atlasv2.ClusterSearchIndex{}
		if err := file.Load(opts.Fs, opts.Filename, index); err != nil {
			return nil, err
		}
		return index, nil
	}

	f, err := opts.indexFields()
	if err != nil {
		return nil, err
	}
	i := &atlasv2.ClusterSearchIndex{
		Analyzer:       &opts.Analyzer,
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  f,
		},
		Name:           opts.Name,
		SearchAnalyzer: &opts.SearchAnalyzer,
	}
	return i, nil
}

// indexFieldParts index field should be fieldName:analyzer:fieldType.
const indexFieldParts = 2

func (opts *IndexOpts) indexFields() (map[string]interface{}, error) {
	if len(opts.fields) == 0 {
		return nil, nil
	}
	fields := make(map[string]interface{})
	for _, p := range opts.fields {
		f := strings.Split(p, ":")
		if len(f) != indexFieldParts {
			return nil, fmt.Errorf("partition should be fieldName:fieldType, got: %s", p)
		}
		fields[f[0]] = map[string]interface{}{
			"type": f[1],
		}
	}
	return fields, nil
}
