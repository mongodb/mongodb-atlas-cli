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

type Compose interface {
	Render() (io.Reader, error)
	Down() error
	Up() error
	Logs() error
	Pause() error
	Unpause() error
	Start() error
}

type composeImpl struct {
	Name          string
	Port          string
	MongodVersion string
	BindIp        string
	Username      string
	Password      string
	KeyFile       string
}

type Option func(*composeImpl)

func New(name string, opt ...Option) Compose {
	c := &composeImpl{
		Name:          name,
		Port:          "27017",
		MongodVersion: "7.0",
		BindIp:        "127.0.0.1",
		KeyFile:       "keyfile",
	}

	for _, o := range opt {
		o(c)
	}

	return c
}

func WithPort(s string) Option {
	return func(c *composeImpl) {
		c.Port = s
	}
}

func WithMongodVersion(s string) Option {
	return func(c *composeImpl) {
		c.MongodVersion = s
	}
}

func WithBindIp(s string) Option {
	return func(c *composeImpl) {
		c.BindIp = s
	}
}

func WithKeyFile(s string) Option {
	return func(c *composeImpl) {
		c.KeyFile = s
	}
}

func WithUsername(s string) Option {
	return func(c *composeImpl) {
		c.Username = s
	}
}

func WithPassword(s string) Option {
	return func(c *composeImpl) {
		c.Password = s
	}
}

func (opt *composeImpl) Render() (io.Reader, error) {
	buf := bytes.NewBufferString("")
	if err := composeTemplate.Execute(buf, opt); err != nil {
		return nil, err
	}
	return buf, nil
}

func (opt *composeImpl) run(args ...string) error {
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

func (opt *composeImpl) Down() error {
	return opt.run("down", "-v")
}

func (opt *composeImpl) Up() error {
	return opt.run("up", "-d", "--wait")
}

func (opt *composeImpl) Logs() error {
	return opt.run("logs")
}

func (opt *composeImpl) Pause() error {
	return opt.run("pause")
}

func (opt *composeImpl) Unpause() error {
	return opt.run("unpause")
}

func (opt *composeImpl) Start() error {
	return opt.run("start")
}
