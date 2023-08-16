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

package searchindexes

import (
	"encoding/json"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	cli.OutputOpts
	filename string
	fs       afero.Fs
	debug    bool
}

const createTemplate = `index created
`

type Index struct {
	CollectionName             string  `json:"collectionName,omitempty"`
	LastObservedCollectionName string  `json:"lastObservedCollectionName,omitempty"`
	Database                   string  `json:"database,omitempty"`
	IndexID                    *string `json:"indexID,omitempty"`
	Mappings                   *struct {
		Dynamic *bool                             `json:"dynamic,omitempty"`
		Fields  map[string]map[string]interface{} `json:"fields,omitempty"`
	} `json:"mappings,omitempty"`
	Name           string  `json:"name"`
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty"`
	Status         *string `json:"status,omitempty"`
}

func connString(db string, port int) string {
	return fmt.Sprintf("mongodb://localhost:%d/%s", port, db)
}

func (opts *CreateOpts) Run() error {
	var index Index
	if err := file.Load(opts.fs, opts.filename, &index); err != nil {
		return err
	}

	collectionName := index.CollectionName
	database := index.Database
	indexName := index.Name

	// todo: instead of cleaning fields not part of the index definition, create a separate struct
	// ref: https://www.mongodb.com/docs/manual/reference/method/db.collection.createSearchIndex/
	index.CollectionName = ""
	index.Database = ""

	serializedIndex, err := json.Marshal(index)
	if err != nil {
		return err
	}

	if opts.debug {
		fmt.Println("creating index: ", string(serializedIndex))
	}

	idxCommand := fmt.Sprintf("db.%s.createSearchIndex('%s', %s)", collectionName, indexName, string(serializedIndex))
	if err = mongosh.Exec(opts.debug, connString(database, 37017), "--eval", idxCommand); err != nil {
		return err
	}

	return opts.Print(createTemplate)
}

func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{fs: afero.NewOsFs()}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new index.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.SearchFilename)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagRequired(flag.File)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
