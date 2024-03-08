package imported

import alias "github.com/fmenezes/mongodb-atlas-cli/atlascli/tools/unused/internal/rules/testdata/notimported"

func Again() { // want `Exported func but never imported`
	alias.ExportedAndImported()
}
