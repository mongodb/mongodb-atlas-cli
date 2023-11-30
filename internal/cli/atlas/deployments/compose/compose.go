package compose

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"os/exec"
	"text/template"
)

//go:embed compose.yaml.tmpl
var composeContent string
var composeTemplate *template.Template = template.Must(template.New("").Parse(composeContent))

type Compose struct {
	Name          string
	Port          string
	MongodVersion string
	BindIp        string
	Username      string
	Password      string
	KeyFile       string
}

func New(name string) *Compose {
	return &Compose{
		Name:          name,
		Port:          "27017",
		MongodVersion: "7.0",
		BindIp:        "127.0.0.1",
		KeyFile:       "keyfile",
	}
}

func (opt *Compose) Render() (io.Reader, error) {
	buf := bytes.NewBufferString("")
	if err := composeTemplate.Execute(buf, opt); err != nil {
		return nil, err
	}
	return buf, nil
}

func (opt *Compose) Run(args ...string) error {
	buf, err := opt.Render()
	if err != nil {
		return err
	}
	composeArgs := append([]string{"compose", "-f", "/dev/stdin"}, args...)
	cmd := exec.Command("docker", composeArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = buf
	cmd.Env = append(os.Environ(), "KEY_FILE="+opt.KeyFile)
	return cmd.Run()
}
