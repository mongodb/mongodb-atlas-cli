package rules_test

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/tools/unused/internal/rules"
	"golang.org/x/tools/go/analysis/analysistest"
)

type ignoreErr struct {
}

func (ignoreErr) Errorf(format string, args ...interface{}) {

}

func TestNotImported(t *testing.T) {
	testDataDir := analysistest.TestData()

	analysistest.Run(ignoreErr{}, testDataDir, rules.NotImportedFirstPass, "./notimported", "./imported")
	analysistest.Run(t, testDataDir, rules.NotImported, "./notimported", "./imported")
}
