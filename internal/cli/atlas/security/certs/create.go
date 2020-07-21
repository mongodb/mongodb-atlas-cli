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

package certs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type SaveOpts struct {
	cli.GlobalOpts
	store    store.X509CertificateStore
	casPath string
}

func (opts *SaveOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *SaveOpts) Run() error {
	index, err := opts.store.X509Certificates()
	if err != nil {
		return err
	}
	result, err := opts.store.CreateSearchIndexes(opts.ConfigProjectID(), opts.clusterName, index)
	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *SaveOpts) newSearchIndex() (*atlas.SearchIndex, error) {
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

// mongocli atlas security certs create --projectId projectId --username dbUser
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: description.CreateSearchIndexes,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !opts.dynamic && len(opts.fields) == 0 {
				return errors.New("you need to specify fields for the index or use a dynamic index")
			}
			if opts.dynamic && len(opts.fields) > 0 {
				return errors.New("you can't specify fields and dynamic at the same time")
			}
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.CASFilePath, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Collection)

	return cmd
}
