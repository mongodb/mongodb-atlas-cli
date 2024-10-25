//nolint:revive,stylecheck
package L1

import (
	"fmt"
	"strings"
)

type GroupedAndSortedCommands []Group

type Group struct {
	Name        string
	Description string
	Commands    []Command
}

type Command struct {
	OperationID       string
	Description       string
	RequestParameters RequestParameters
	Versions          []Version
}

type RequestParameters struct {
	URL             string
	QueryParameters []Parameter
	URLParameters   []Parameter
	Verb            HTTPVerb
}

type Version struct {
	Version              string
	RequestContentTypes  []string
	ResponseContentTypes []string
}

type Parameter struct {
	Name        string
	Description string
	Required    bool
}

type HTTPVerb string

const (
	DELETE HTTPVerb = "DELETE"
	GET    HTTPVerb = "GET"
	POST   HTTPVerb = "POST"
)

func ToHTTPVerb(s string) (HTTPVerb, error) {
	verb := HTTPVerb(strings.ToUpper(s))

	switch verb {
	case DELETE, GET, POST:
		return verb, nil
	default:
		return "", fmt.Errorf("invalid HTTP verb: %s", s)
	}
}
