package main

import (
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/tools/unused/internal/rules"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(rules.NotImportedFirstPass, rules.NotImported)
}
