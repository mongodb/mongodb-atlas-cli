package cli

import (
	"fmt"
	"log"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/astparsing"
)

func TestTemplates(t *testing.T) {
	if err := inner(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// use later
	println(t)
}

func inner() error {
	pkgs, err := astparsing.LoadPackagesRecursive("../atlas/accesslists/list.go")
	// pkgs, err := astparsing.LoadPackagesRecursive("..")
	if err != nil {
		return err
	}

	builderFuncs := astparsing.LoadCommandBuilderInfos(pkgs)

	for _, builderFunc := range builderFuncs {
		fmt.Printf("%v\n", builderFunc.TemplateValue)
	}

	return nil
}
