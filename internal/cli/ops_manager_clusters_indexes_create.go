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

package cli

import (
	"fmt"
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/messages"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerClustersIndexesCreateOpts struct {
	globalOpts
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

func (opts *opsManagerClustersIndexesCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerClustersIndexesCreateOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())
	if err != nil {
		return err
	}

	index, err := opts.newIndex()
	if err != nil {
		return err
	}

	current.IndexConfigs = index

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

func (opts *opsManagerClustersIndexesCreateOpts) newIndex() ([]*om.IndexConfigs, error) {
	keys, err := opts.indexKeys()
	if err != nil {
		return nil, err
	}

	i := new(om.IndexConfigs)
	i.DBName = opts.db
	i.CollectionName = opts.collection
	i.RSName = opts.rsName
	i.Key = keys
	i.Options = opts.newIndexOptions()

	if opts.locale != "" {
		i.Collation = opts.newCollationOptions()
	}

	return []*om.IndexConfigs{i}, nil
}

func (opts *opsManagerClustersIndexesCreateOpts) newIndexOptions() *atlas.IndexOptions {
	return &atlas.IndexOptions{
		Background: opts.background,
		Unique:     opts.unique,
		Sparse:     opts.sparse,
		Name:       opts.name,
	}
}

func (opts *opsManagerClustersIndexesCreateOpts) newCollationOptions() *atlas.CollationOptions {
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

//// indexKeys  takes a slice of values formatted as key:vale and returns an array of slice [key]:value
func (opts *opsManagerClustersIndexesCreateOpts) indexKeys() ([][]string, error) {
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
func OpsManagerClustersIndexesCreateBuilder() *cobra.Command {
	opts := &opsManagerClustersIndexesCreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateCluster,
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.name = args[0]
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.db, flags.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.rsName, flags.RSName, "", usage.RSName)
	cmd.Flags().StringVar(&opts.collection, flags.CollectionName, "", usage.Collection)
	cmd.Flags().StringArrayVar(&opts.keys, flags.Key, nil, usage.Key)
	cmd.Flags().StringVar(&opts.locale, flags.Locale, "", usage.Locale)
	cmd.Flags().StringVar(&opts.caseFirst, flags.CaseFirst, "", usage.CaseFirst)
	cmd.Flags().StringVar(&opts.alternate, flags.Alternate, "", usage.Alternate)
	cmd.Flags().StringVar(&opts.maxVariable, flags.MaxVariable, "", usage.MaxVariable)
	cmd.Flags().BoolVar(&opts.caseLevel, flags.CaseLevel, false, usage.CaseLevel)
	cmd.Flags().BoolVar(&opts.numericOrdering, flags.NumericOrdering, false, usage.NumericOrdering)
	cmd.Flags().BoolVar(&opts.normalization, flags.Normalization, false, usage.Normalization)
	cmd.Flags().BoolVar(&opts.backwards, flags.Backwards, false, usage.Backwards)
	cmd.Flags().IntVar(&opts.strength, flags.Strength, 0, usage.Strength)
	cmd.Flags().BoolVar(&opts.unique, flags.Unique, false, usage.Unique)
	cmd.Flags().BoolVar(&opts.sparse, flags.Sparse, false, usage.Sparse)
	cmd.Flags().BoolVar(&opts.background, flags.Background, false, usage.Background)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.RSName)
	_ = cmd.MarkFlagRequired(flags.Database)
	_ = cmd.MarkFlagRequired(flags.CollectionName)
	_ = cmd.MarkFlagRequired(flags.Key)

	return cmd
}
