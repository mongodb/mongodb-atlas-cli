package output

import (
	"github.com/mongodb/mongocli/internal/json"
	"os"
	"text/template"
)

type Config interface {
	Output() string
}

const jsonFormat = "json"

// Print outputs v to os.Stdout while handling configured formats,
// if the optional t is given then it's processed as a go-template
func Print(c Config, t string, v interface{}) error {
	if c.Output() == jsonFormat {
		return json.PrettyPrint(v)
	}
	if t != ""{
		tmpl, err := template.New("test").Parse(t)
		if err != nil {
			return err
		}
		return tmpl.Execute(os.Stdout, v)
	}
	return nil
}
