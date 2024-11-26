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

package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/tangzero/inflector"
	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed spec.yaml
	spec []byte
	//go:embed templates/*.gotmpl
	templateFolder embed.FS

	templateFuncs = template.FuncMap{
		"Year": func() int {
			return time.Now().Year()
		},
		"Now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"Join": strings.Join,
	}

	ErrGenerateCli     = errors.New("error generating cli")
	ErrGenerateStore   = errors.New("error generating store")
	ErrGenerateCommand = errors.New("error generating command")
)

type CLI struct {
	Commands []Command
	Stores   []Store

	overwrite bool
	basePath  string
	templates *template.Template
}

type Command struct {
	CommandPath    string    `yaml:"command_path,omitempty"`
	PackageName    string    `yaml:"package_name,omitempty"`
	OutputTemplate string    `yaml:"output_template,omitempty"`
	Description    string    `yaml:"description,omitempty"`
	StoreName      string    `yaml:"store_name,omitempty"`
	StoreMethod    string    `yaml:"store_method,omitempty"`
	Template       string    `yaml:"template,omitempty"`
	IDDescription  string    `yaml:"id_description,omitempty"`
	IDName         string    `yaml:"id_name,omitempty"`
	Example        string    `yaml:"example,omitempty"`
	RequestType    string    `yaml:"request_type,omitempty"`
	SubCommands    []Command `yaml:"sub_commands,omitempty"`
}

const (
	gotmplExt     = ".gotmpl"
	gotmplTestExt = "_test" + gotmplExt
	goExt         = ".go"
)

func (c *Command) TemplateFile() string {
	return c.Template + gotmplExt
}

func (c *Command) TemplateUnitTestFile() string {
	return c.Template + gotmplTestExt
}

type Store struct {
	BaseFileName string          `yaml:"base_file_name,omitempty"`
	Template     string          `yaml:"template,omitempty"`
	Lister       *StoreInterface `yaml:"lister,omitempty"`
	Describer    *StoreInterface `yaml:"describer,omitempty"`
	Creator      *StoreInterface `yaml:"creator,omitempty"`
	Updater      *StoreInterface `yaml:"updater,omitempty"`
	Deleter      *StoreInterface `yaml:"deleter,omitempty"`
}

func (s *Store) TemplateFile() string {
	return s.Template + gotmplExt
}

func (s *Store) InterfaceNames() string {
	var names []string
	if s.Lister != nil {
		names = append(names, s.Lister.Name)
	}
	if s.Describer != nil {
		names = append(names, s.Describer.Name)
	}
	if s.Creator != nil {
		names = append(names, s.Creator.Name)
	}
	if s.Updater != nil {
		names = append(names, s.Updater.Name)
	}
	if s.Deleter != nil {
		names = append(names, s.Deleter.Name)
	}
	return strings.Join(names, ",")
}

type StoreInterface struct {
	Name       string `yaml:"name,omitempty"`
	Method     string `yaml:"method,omitempty"`
	SDKMethod  string `yaml:"sdk_method,omitempty"`
	ArgName    string `yaml:"arg_name,omitempty"`
	ArgType    string `yaml:"arg_type,omitempty"`
	ReturnType string `yaml:"return_type,omitempty"`
}

func (c *Command) CommandPaths() []string {
	return strings.Split(c.CommandPath, " ")
}

func (c *Command) LastCommandPath() string {
	commandPaths := c.CommandPaths()

	return commandPaths[len(commandPaths)-1]
}

func (c *Command) baseFileName(basePath string) string {
	internalPath := strings.ToLower(filepath.Join(c.CommandPaths()...))

	if len(c.SubCommands) > 0 {
		internalPath = filepath.Join(internalPath, inflector.Underscorize(c.LastCommandPath()))
	}

	return filepath.Join(basePath, "internal", "cli", internalPath)
}

func (c *Command) FileName(basePath string) string {
	return c.baseFileName(basePath) + goExt
}

func (c *Command) UnitTestFileName(basePath string) string {
	return c.baseFileName(basePath) + gotmplTestExt
}

func fileExists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

func (cli *CLI) generateStore(store *Store) error {
	storeFile := filepath.Join(cli.basePath, "internal", "store", "atlas", store.BaseFileName+goExt)

	fileCreated, err := cli.generateFile(storeFile, store.TemplateFile(), store)
	if err != nil {
		return err
	}
	if !fileCreated {
		return nil
	}
	return goGenerate(storeFile)
}

func (cli *CLI) generateFile(file string, templateFile string, data any) (bool, error) {
	if !cli.overwrite && fileExists(file) {
		_, _ = fmt.Fprintf(os.Stderr, "File %q already present in disk, skipping\n", file)
		return false, nil
	}
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return false, err
	}

	f, err := os.Create(file)
	if err != nil {
		return false, err
	}

	defer f.Close()

	err = cli.templates.ExecuteTemplate(f, templateFile, data)
	if err != nil {
		return false, err
	}

	err = cleanupFile(file)
	if err != nil {
		return true, err
	}

	return true, nil
}

func (cli *CLI) generateCommand(cmd *Command) error {
	_, err := cli.generateFile(cmd.FileName(cli.basePath), cmd.TemplateFile(), cmd)
	if err != nil {
		return err
	}

	_, err = cli.generateFile(cmd.UnitTestFileName(cli.basePath), cmd.TemplateUnitTestFile(), cmd)
	if err != nil {
		return err
	}

	for i := range cmd.SubCommands {
		if err := cli.generateCommand(&cmd.SubCommands[i]); err != nil {
			return err
		}
	}

	return nil
}

const filePermissions = 0o600

func cleanupFile(filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	r, err := imports.Process(filePath, b, nil)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, r, filePermissions)
}

func goGenerate(filePath string) error {
	execCmd := exec.Command("go", "generate", filePath)
	_, err := execCmd.Output()
	return err
}

func runMake() error {
	execCmd := exec.Command("make", "gen-docs")
	_, err := execCmd.Output()
	return err
}

func newCli(overwrite bool) (*CLI, error) {
	cli := CLI{overwrite: overwrite}
	err := yaml.Unmarshal(spec, &cli)
	if err != nil {
		return nil, err
	}
	cli.basePath, err = filepath.Abs("./")
	if err != nil {
		return nil, err
	}

	cli.templates, err = template.New("root").Funcs(templateFuncs).ParseFS(templateFolder, "templates/*.gotmpl")
	if err != nil {
		return nil, err
	}

	return &cli, nil
}

func (cli *CLI) generateCli() {
	for i := range cli.Stores {
		if err := cli.generateStore(&cli.Stores[i]); err != nil {
			log.Printf("%s: %s\n", ErrGenerateStore, err)
		}
	}

	for i := range cli.Commands {
		if err := cli.generateCommand(&cli.Commands[i]); err != nil {
			log.Printf("%s: %s\n", ErrGenerateCommand, err)
		}
	}

	if err := runMake(); err != nil {
		log.Printf("%s: %s\n", ErrGenerateCli, err)
	}
}
