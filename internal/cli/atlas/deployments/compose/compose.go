package compose

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"os"
	"os/exec"
	"text/template"
)

//go:embed compose.yaml.tmpl
var composeContent string
var composeTemplate = template.Must(template.New("").Parse(composeContent))

type Compose interface {
	Render() (io.Reader, error)
	Down(context.Context) error
	Up(context.Context) error
	Logs(context.Context) error
	Pause(context.Context) error
	Unpause(context.Context) error
	Start(context.Context) error
}

type composeImpl struct {
	Name          string
	Port          string
	MongodVersion string
	BindIP        string
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
		BindIP:        "127.0.0.1",
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

func WithBindIP(s string) Option {
	return func(c *composeImpl) {
		c.BindIP = s
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

func (opt *composeImpl) run(ctx context.Context, args ...string) error {
	buf, err := opt.Render()
	if err != nil {
		return err
	}
	composeArgs := append([]string{"compose", "-f", "/dev/stdin"}, args...)
	cmd := exec.CommandContext(ctx, "docker", composeArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = buf
	cmd.Env = append(os.Environ(), "KEY_FILE="+opt.KeyFile)
	return cmd.Run()
}

func (opt *composeImpl) Down(ctx context.Context) error {
	return opt.run(ctx, "down", "-v")
}

func (opt *composeImpl) Up(ctx context.Context) error {
	return opt.run(ctx, "up", "-d", "--wait")
}

func (opt *composeImpl) Logs(ctx context.Context) error {
	return opt.run(ctx, "logs")
}

func (opt *composeImpl) Pause(ctx context.Context) error {
	return opt.run(ctx, "pause")
}

func (opt *composeImpl) Unpause(ctx context.Context) error {
	return opt.run(ctx, "unpause")
}

func (opt *composeImpl) Start(ctx context.Context) error {
	return opt.run(ctx, "start")
}
