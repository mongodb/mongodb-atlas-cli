package rules

import (
	"golang.org/x/tools/go/analysis"
)

var NotImported = &analysis.Analyzer{
	Name: "NotImported",
	Doc:  "Avoid exported methods that are not imported",
	Run:  notImportedRun,
}

func notImportedRun(pass *analysis.Pass) (interface{}, error) {
	for _, name := range pass.Pkg.Scope().Names() {
		obj := pass.Pkg.Scope().Lookup(name)
		if !obj.Exported() {
			continue
		}
		mutex.Lock()
		ok := Usages[pass.Pkg.Path()][name]
		mutex.Unlock()
		if !ok {
			pass.Report(analysis.Diagnostic{
				Pos:     obj.Pos(),
				Message: "Exported func but never imported",
			})
		}
	}
	return nil, nil
}
