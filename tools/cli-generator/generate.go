package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	//go:embed spec.yaml
	spec []byte
	//go:embed commandListTemplate.txt
	commandListTemplateContent string
	//go:embed commandDescribeTemplate.txt
	commandDescribeTemplateContent string
	//go:embed commandDeleteTemplate.txt
	commandDeleteTemplateContent string
	//go:embed commandCreateTemplate.txt
	commandCreateTemplateContent string
	//go:embed commandUpdateTemplate.txt
	commandUpdateTemplateContent string
	//go:embed commandParentTemplate.txt
	commandParentTemplateContent string
	//go:embed store.txt
	storeTemplateContent string

	templateFuncs = template.FuncMap{
		"Year": func() int {
			return time.Now().Year()
		},
		"Now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"Join": strings.Join,
	}

	ErrGenerateCli        = errors.New("error generating cli")
	ErrCommandTypeUnknown = errors.New("command type is unknown")
)

type CLI struct {
	Commands []Command
	Stores   []Store

	basePath                string
	commandListTemplate     *template.Template
	commandDescribeTemplate *template.Template
	commandDeleteTemplate   *template.Template
	commandCreateTemplate   *template.Template
	commandUpdateTemplate   *template.Template
	commandParentTemplate   *template.Template
	storeTemplate           *template.Template
}

type Command struct {
	CommandPath   string    `yaml:"command_path,omitempty"`
	PackageName   string    `yaml:"package_name,omitempty"`
	Template      string    `yaml:"template,omitempty"`
	Description   string    `yaml:"description,omitempty"`
	StoreName     string    `yaml:"store_name,omitempty"`
	StoreMethod   string    `yaml:"store_method,omitempty"`
	IDDescription string    `yaml:"id_description,omitempty"`
	IDName        string    `yaml:"id_name,omitempty"`
	Example       string    `yaml:"example,omitempty"`
	RequestType   string    `yaml:"request_type,omitempty"`
	SubCommands   []Command `yaml:"sub_commands,omitempty"`
}

type Store struct {
	BaseFileName string          `yaml:"base_file_name,omitempty"`
	Lister       *StoreInterface `yaml:"lister,omitempty"`
	Describer    *StoreInterface `yaml:"describer,omitempty"`
	Creator      *StoreInterface `yaml:"creator,omitempty"`
	Updater      *StoreInterface `yaml:"updater,omitempty"`
	Deleter      *StoreInterface `yaml:"deleter,omitempty"`
}

func (s *Store) InterfaceNames() string {
	names := []string{}
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

type CommandType int

const (
	CommandTypeList CommandType = iota
	CommandTypeDescribe
	CommandTypeDelete
	CommandTypeCreate
	CommandTypeUpdate
	CommandTypeParent
)

func (c *Command) CommandPaths() []string {
	return strings.Split(c.CommandPath, " ")
}

func (c *Command) LastCommandPath() string {
	commandPaths := c.CommandPaths()

	return commandPaths[len(commandPaths)-1]
}

func (c *Command) CommandType() CommandType {
	switch c.LastCommandPath() {
	case "list":
		return CommandTypeList
	case "describe":
		return CommandTypeDescribe
	case "delete":
		return CommandTypeDelete
	case "create":
		return CommandTypeCreate
	case "update":
		return CommandTypeUpdate
	}

	return CommandTypeParent
}

func (c *Command) FileName(basePath string) string {
	if c.CommandType() == CommandTypeParent {
		return filepath.Join(basePath, "internal", "cli", filepath.Join(c.CommandPaths()...), c.LastCommandPath()+".go")
	}

	return filepath.Join(basePath, "internal", "cli", filepath.Join(c.CommandPaths()...)+".go")

}

func (cli *CLI) generateStore(store *Store) error {
	storeFile := filepath.Join(cli.basePath, "internal", "store", store.BaseFileName+".go")

	err := os.MkdirAll(filepath.Dir(storeFile), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(storeFile)
	if err != nil {
		return err
	}

	defer f.Close()

	err = cli.storeTemplate.Execute(f, store)
	if err != nil {
		return err
	}

	return cleanupFile(storeFile, true)
}

func (cli *CLI) template(cmd *Command) *template.Template {
	switch cmd.CommandType() {
	case CommandTypeList:
		return cli.commandListTemplate
	case CommandTypeDescribe:
		return cli.commandDescribeTemplate
	case CommandTypeDelete:
		return cli.commandDeleteTemplate
	case CommandTypeCreate:
		return cli.commandCreateTemplate
	case CommandTypeUpdate:
		return cli.commandUpdateTemplate
	default:
		return cli.commandParentTemplate
	}
}

func (cli *CLI) generateCommand(cmd *Command) error {
	commandFile := cmd.FileName(cli.basePath)

	err := os.MkdirAll(filepath.Dir(commandFile), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(commandFile)
	if err != nil {
		return err
	}

	defer f.Close()

	tpl := cli.template(cmd)
	err = tpl.Execute(f, cmd)
	if err != nil {
		return err
	}

	err = cleanupFile(commandFile, false)
	if err != nil {
		return err
	}

	for _, childCmd := range cmd.SubCommands {
		err = cli.generateCommand(&childCmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func cleanupFile(filePath string, generateMocks bool) error {
	execCmd := exec.Command("goimports", "-w", filePath)
	_, err := execCmd.Output()
	if err != nil {
		return err
	}

	execCmd = exec.Command("gofmt", "-w", "-s", filePath)
	_, err = execCmd.Output()
	if err != nil {
		return err
	}

	if generateMocks {
		execCmd = exec.Command("go", "generate", filePath)
		_, err = execCmd.Output()
		return err
	}

	return nil
}

func runMake() error {
	execCmd := exec.Command("make", "gen-docs")
	_, err := execCmd.Output()
	if err != nil {
		return err
	}
	return err
}

func newCli() (*CLI, error) {
	var cli CLI
	err := yaml.Unmarshal(spec, &cli)
	if err != nil {
		return nil, err
	}
	cli.basePath, err = filepath.Abs("./")
	if err != nil {
		return nil, err
	}

	cli.commandListTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandListTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.commandDescribeTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandDescribeTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.commandDeleteTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandDeleteTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.commandCreateTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandCreateTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.commandUpdateTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandUpdateTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.commandParentTemplate, err = template.New("").Funcs(templateFuncs).Parse(commandParentTemplateContent)
	if err != nil {
		return nil, err
	}

	cli.storeTemplate, err = template.New("").Funcs(templateFuncs).Parse(storeTemplateContent)
	if err != nil {
		return nil, err
	}

	return &cli, nil
}

func (cli *CLI) generateCli() error {
	for _, store := range cli.Stores {
		err := cli.generateStore(&store)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrGenerateCli, err)
		}
	}

	for _, cmd := range cli.Commands {
		err := cli.generateCommand(&cmd)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrGenerateCli, err)
		}
	}

	err := runMake()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrGenerateCli, err)
	}

	return nil
}
