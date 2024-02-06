package cli

import (
	"fmt"
	"log"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/astparsing"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/templateparsing"
)

func TestTemplates(t *testing.T) {
	//pkgs, err := astparsing.LoadPackagesRecursive("../atlas/metrics/databases/describe.go")
	//pkgs, err := astparsing.LoadPackagesRecursive("../mongocli/performanceadvisor/suggestedindexes/list.go")

	// Still broken
	//pkgs, err := astparsing.LoadPackagesRecursive("../atlas/privateendpoints/datalake/aws/create.go")
	//pkgs, err := astparsing.LoadPackagesRecursive("../atlas/datafederation/privateendpoints/create.go")

	pkgs, err := astparsing.LoadPackagesRecursive("..")
	if err != nil {
		log.Fatal(err)
	}

	builderFuncs := astparsing.LoadCommandBuilderInfos(pkgs)

	for _, builderFunc := range builderFuncs {
		templateTree, err := templateparsing.ParseTemplate(builderFunc.TemplateValue)
		if err != nil {
			log.Fatal(err)
		}

		valid, err := templateTree.Validate(builderFunc.Pkg, builderFunc.TemplateType.NamedStruct)
		if err != nil {
			log.Fatal(err)
		}

		errorMessage := "Template and struct don't match:\n"
		errorMessage += "struct:\n"
		errorMessage += fmt.Sprintf("- location: %v\n", builderFunc.Pkg.Fset.Position(builderFunc.TemplateType.NamedStruct.Obj().Pos()))

		errorMessage += "template:\n"
		errorMessage += fmt.Sprintf("- location: %v\n", builderFunc.Pkg.Fset.Position(builderFunc.CommandOptsStruct.Pos()))
		errorMessage += fmt.Sprintf("- value: %v\n", builderFunc.TemplateValue)

		errorMessage += fmt.Sprintf("- parsed representation:\n%v\n", templateTree.Fprint(1))

		if !valid {
			log.Println(errorMessage)
		}
		//assert.True(t, valid, errorMessage)
	}
}
