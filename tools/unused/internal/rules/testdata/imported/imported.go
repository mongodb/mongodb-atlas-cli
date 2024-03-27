package imported

import "github.com/fmenezes/mongodb-atlas-cli/atlascli/tools/unused/internal/rules/testdata/notimported"

func More() { // want `Exported func but never imported`
	notimported.ExportedAndImported()
}
