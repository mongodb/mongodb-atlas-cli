package options

import (
	"bytes"
	_ "embed"
	"io"
	"text/template"
)

//go:embed compose.yaml.tmpl
var composeContent string
var composeTemplate *template.Template = template.Must(template.New("").Parse(composeContent))

type ComposeDefinitionOptions struct {
	Name          string
	Port          string
	MongodVersion string
	BindIp        string
	Username      string
	Password      string
}

func ComposeDefinition(options *ComposeDefinitionOptions) (io.Reader, error) {
	buf := bytes.NewBufferString("")
	if err := composeTemplate.Execute(buf, options); err != nil {
		return nil, err
	}
	return buf, nil
}
