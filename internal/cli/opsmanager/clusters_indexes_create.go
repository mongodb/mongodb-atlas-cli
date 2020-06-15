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

package opsmanager

import (
	"fmt"
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/atmcfg"
	"go.mongodb.org/ops-manager/opsmngr"
)

type ClustersIndexesCreateOpts struct {
	cli.GlobalOpts
	name            string
	db              string
	collection      string
	rsName          string
	locale          string
	caseFirst       string
	alternate       string
	maxVariable     string
	strength        int
	caseLevel       bool
	numericOrdering bool
	normalization   bool
	backwards       bool
	unique          bool
	sparse          bool
	background      bool
	keys            []string
	store           store.AutomationPatcher
}

func (opts *ClustersIndexesCreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ClustersIndexesCreateOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	index, err := opts.newIndex()
	if err != nil {
		return err
	}

	err = atmcfg.AddIndexConfig(current, index)
	if err != nil {
		return err
	}

	if err := opts.store.UpdateAutomationConfig(opts.ConfigProjectID(), current); err != nil {
		return err
	}

	fmt.Print(cli.DeploymentStatus(config.OpsManagerURL(), opts.ConfigProjectID()))

	return nil
}

func (opts *ClustersIndexesCreateOpts) newIndex() (*opsmngr.IndexConfig, error) {
	keys, err := opts.indexKeys()
	if err != nil {
		return nil, err
	}

	i := new(opsmngr.IndexConfig)
	i.DBName = opts.db
	i.CollectionName = opts.collection
	i.RSName = opts.rsName
	i.Key = keys
	i.Options = opts.newIndexOptions()

	if opts.locale != "" {
		i.Collation = opts.newCollationOptions()
	}

	return i, nil
}

func (opts *ClustersIndexesCreateOpts) newIndexOptions() *atlas.IndexOptions {
	return &atlas.IndexOptions{
		Background: opts.background,
		Unique:     opts.unique,
		Sparse:     opts.sparse,
		Name:       opts.name,
	}
}

func (opts *ClustersIndexesCreateOpts) newCollationOptions() *atlas.CollationOptions {
	return &atlas.CollationOptions{
		Locale:          opts.locale,
		CaseLevel:       opts.caseLevel,
		CaseFirst:       opts.caseFirst,
		Strength:        opts.strength,
		NumericOrdering: opts.numericOrdering,
		Alternate:       opts.alternate,
		MaxVariable:     opts.maxVariable,
		Normalization:   opts.normalization,
		Backwards:       opts.backwards,
	}
}

// indexKeys takes a slice of values formatted as key:vale and returns an array of slice [[key, value][key, value]]
func (opts *ClustersIndexesCreateOpts) indexKeys() ([][]string, error) {
	propertiesList := make([][]string, len(opts.keys))
	for i, key := range opts.keys {
		value := strings.Split(key, ":")
		if len(value) != 2 {
			return nil, fmt.Errorf("unexpected key format: %s", key)
		}
		values := []string{value[0], value[1]}
		propertiesList[i] = values
	}

	return propertiesList, nil
}

// mongocli cloud-manager cluster(s) index(es) create [name]  --rsName rsName --dbName dbName [--key field:type] --projectId projectId
// --locale locale --caseFirst caseFirst --alternate alternate --maxVariable maxVariable --strength strength --caseLevel caseLevel --numericOrdering numericOrdering
// --normalization normalization --backwards backwards --unique unique --sparse sparse --background background
func ClustersIndexesCreateBuilder() *cobra.Command {
	opts := &ClustersIndexesCreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateIndex,
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.name = args[0]
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.db, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.rsName, flag.RSName, "", usage.RSName)
	cmd.Flags().StringVar(&opts.collection, flag.CollectionName, "", usage.Collection)
	cmd.Flags().StringArrayVar(&opts.keys, flag.Key, nil, usage.Key)
	cmd.Flags().StringVar(&opts.locale, flag.Locale, "", usage.Locale)
	cmd.Flags().StringVar(&opts.caseFirst, flag.CaseFirst, "", usage.CaseFirst)
	cmd.Flags().StringVar(&opts.alternate, flag.Alternate, "", usage.Alternate)
	cmd.Flags().StringVar(&opts.maxVariable, flag.MaxVariable, "", usage.MaxVariable)
	cmd.Flags().BoolVar(&opts.caseLevel, flag.CaseLevel, false, usage.CaseLevel)
	cmd.Flags().BoolVar(&opts.numericOrdering, flag.NumericOrdering, false, usage.NumericOrdering)
	cmd.Flags().BoolVar(&opts.normalization, flag.Normalization, false, usage.Normalization)
	cmd.Flags().BoolVar(&opts.backwards, flag.Backwards, false, usage.Backwards)
	cmd.Flags().IntVar(&opts.strength, flag.Strength, 0, usage.Strength)
	cmd.Flags().BoolVar(&opts.unique, flag.Unique, false, usage.Unique)
	cmd.Flags().BoolVar(&opts.sparse, flag.Sparse, false, usage.Sparse)
	cmd.Flags().BoolVar(&opts.background, flag.Background, false, usage.Background)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.RSName)
	_ = cmd.MarkFlagRequired(flag.Database)
	_ = cmd.MarkFlagRequired(flag.CollectionName)
	_ = cmd.MarkFlagRequired(flag.Key)

	_ = cmd.Flags().MarkHidden(flag.Background) // Deprecated

	return cmd
}
