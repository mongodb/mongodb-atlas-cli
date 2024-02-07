// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e || generic

package e2e

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e/templates_test/astparsing"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e/templates_test/templateparsing"
	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	pkgs, err := astparsing.LoadPackagesRecursive("../../../internal/cli/")
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

	assert.Empty(t, templateValidationErrors, strings.Join(templateValidationErrors, "\n\n"))
}
