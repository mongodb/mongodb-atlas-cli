package main

import (
	"embed"
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
	//go:embed templates/*.txt
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

	ErrGenerateCli = errors.New("error generating cli")
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

func (c *Command) TemplateFile() string {
	return c.Template + ".txt"
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
	return s.Template + ".txt"
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

func (c *Command) CommandPaths() []string {
	return strings.Split(c.CommandPath, " ")
}

func (c *Command) LastCommandPath() string {
	commandPaths := c.CommandPaths()

	return commandPaths[len(commandPaths)-1]
}

func (c *Command) FileName(basePath string) string {
	if len(c.SubCommands) > 0 {
		return filepath.Join(basePath, "internal", "cli", filepath.Join(c.CommandPaths()...), c.LastCommandPath()+".go")
	}

	return filepath.Join(basePath, "internal", "cli", filepath.Join(c.CommandPaths()...)+".go")

}

func fileExists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

func (cli *CLI) generateStore(store *Store) error {
	storeFile := filepath.Join(cli.basePath, "internal", "store", store.BaseFileName+".go")

	if !cli.overwrite && fileExists(storeFile) {
		fmt.Printf("File '%s' already present in disk, skipping\n", storeFile)
		return nil
	}

	err := os.MkdirAll(filepath.Dir(storeFile), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(storeFile)
	if err != nil {
		return err
	}

	defer f.Close()

	err = cli.templates.ExecuteTemplate(f, store.TemplateFile(), store)
	if err != nil {
		return err
	}

	return cleanupFile(storeFile, true)
}

func (cli *CLI) generateCommand(cmd *Command) error {
	commandFile := cmd.FileName(cli.basePath)

	if !cli.overwrite && fileExists(commandFile) {
		fmt.Printf("File '%s' already present in disk, skipping\n", commandFile)
	} else {
		err := os.MkdirAll(filepath.Dir(commandFile), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.Create(commandFile)
		if err != nil {
			return err
		}

		defer f.Close()

		err = cli.templates.ExecuteTemplate(f, cmd.TemplateFile(), cmd)
		if err != nil {
			return err
		}

		err = cleanupFile(commandFile, false)
		if err != nil {
			return err
		}
	}

	for _, childCmd := range cmd.SubCommands {
		err := cli.generateCommand(&childCmd)
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

	cli.templates, err = template.New("root").Funcs(templateFuncs).ParseFS(templateFolder, "templates/*.txt")
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
