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
	"math/rand"
	"os"
	"path"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	cli.OutputOpts
	filename string
	fs       afero.Fs
}

const createTemplate = `index created
`

type Index struct {
	CollectionName             string  `json:"collectionName"`
	LastObservedCollectionName string  `json:"lastObservedCollectionName"`
	Database                   string  `json:"database"`
	IndexID                    *string `json:"indexID,omitempty"`
	Mappings                   *struct {
		Dynamic *bool                             `json:"dynamic,omitempty"`
		Fields  map[string]map[string]interface{} `json:"fields,omitempty"`
	} `json:"mappings,omitempty"`
	Name           string  `json:"name"`
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty"`
	Status         *string `json:"status,omitempty"`
}

type SeachConfig struct {
	Address string `json:"address"`
	ID      struct {
		GroupID     string `json:"groupId"`
		ClusterName string `json:"clusterName"`
	} `json:"id"`
	ConnectionString string   `json:"connectionString"`
	HostnameRegex    string   `json:"hostnameRegex"`
	Indexes          []*Index `json:"indexes"`
	Analyzers        []any    `json:"analyzers"`
}

func mongotConfigPath() (string, error) {
	configHome, err := config.AtlasCLIConfigHome()
	if err != nil {
		return "", err
	}
	return path.Join(configHome, "mms-config.json"), nil
}

func (opts *CreateOpts) loadConfig() (*SeachConfig, error) {
	configPath, err := mongotConfigPath()
	if err != nil {
		return nil, err
	}
	var config SeachConfig
	if err := file.Load(opts.fs, configPath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (*CreateOpts) dumpConfig(config *SeachConfig) error {
	configPath, err := mongotConfigPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(configPath, data, os.ModePerm); err != nil {
		return err
	}

	podman.CopyFileToContainer(true, configPath, "mms", "/etc/mms/mms-config.json")
	return nil
}

var letterRunes = []rune("0123456789abcdef")

const objectIDLen = 24

func randObjectID() string {
	b := make([]rune, objectIDLen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] //nolint:gosec // ignore pseudo random is good enough
	}
	return string(b)
}

func (opts *CreateOpts) Run() error {
	var index Index
	if err := file.Load(opts.fs, opts.filename, &index); err != nil {
		return err
	}
	index.IndexID = pointer.Get(randObjectID())

	config, err := opts.loadConfig()
	if err != nil {
		return err
	}

	config.Indexes = append(config.Indexes, &index)

	err = opts.dumpConfig(config)
	if err != nil {
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

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.SearchFilename)
	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagRequired(flag.File)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
