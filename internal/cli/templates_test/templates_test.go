package cli

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/astparsing"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/templateparsing"
	"github.com/stretchr/testify/assert"
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
	templateValidationErrors := make([]string, 0)

	for _, builderFunc := range builderFuncs {
		templateTree, err := templateparsing.ParseTemplate(builderFunc.TemplateValue)
		if err != nil {
			log.Fatal(err)
		}

		validationResult, err := templateTree.Validate(builderFunc.Pkg, builderFunc.TemplateType.NamedStruct)
		if err != nil {
			log.Fatal(err)
		}

		errorMessage := "Template and struct don't match:\n"

		errorMessage += "Error messages:\n"
		for _, message := range validationResult.ErrorMessages() {
			errorMessage += fmt.Sprintf("- %v\n", message)
		}

		errorMessage += "\nStruct:\n"
		errorMessage += fmt.Sprintf("- location: %v\n", builderFunc.Pkg.Fset.Position(builderFunc.TemplateType.NamedStruct.Obj().Pos()))

		errorMessage += "\nTemplate:\n"
		errorMessage += fmt.Sprintf("- location: %v\n", builderFunc.Pkg.Fset.Position(builderFunc.CommandOptsStruct.Pos()))
		errorMessage += fmt.Sprintf("- value: %v\n", builderFunc.TemplateValue)

		if !validationResult.IsValid() {
			templateValidationErrors = append(templateValidationErrors, errorMessage)
		}
	}

	assert.True(t, len(templateValidationErrors) == 0, strings.Join(templateValidationErrors, "\n\n"))
}
