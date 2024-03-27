package rules

import (
	"go/ast"
	"slices"
	"strings"
	"sync"

	"golang.org/x/tools/go/analysis"
)

var NotImportedFirstPass = &analysis.Analyzer{
	Name: "NotImportedFirstPass",
	Doc:  "Avoid exported methods that are not imported",
	Run:  notImportedRunFirstPass,
}

var Usages = map[string]map[string]bool{}
var mutex = &sync.RWMutex{}

func stripQuotes(s string) string {
	c, _ := strings.CutPrefix(s, `"`)
	c, _ = strings.CutSuffix(c, `"`)
	return c
}

func checkImportAlias(i *ast.ImportSpec, alias *ast.Ident) bool {
	if i.Name != nil {
		return i.Name == alias
	}

	s := strings.Split(stripQuotes(i.Path.Value), "/")
	return s[len(s)-1] == alias.Name
}

func process(file *ast.File, n ast.Node) {
	if ce, ok := n.(*ast.CallExpr); ok {
		if se, ok := ce.Fun.(*ast.SelectorExpr); ok {
			if alias, ok := se.X.(*ast.Ident); ok {
				i := slices.IndexFunc(file.Imports, func(i *ast.ImportSpec) bool {
					return checkImportAlias(i, alias)
				})
				if i >= 0 {
					pkgPath := stripQuotes(file.Imports[i].Path.Value)
					name := se.Sel.Name
					mutex.Lock()
					if Usages[pkgPath] == nil {
						Usages[pkgPath] = map[string]bool{}
					}
					Usages[pkgPath][name] = true
					mutex.Unlock()
				}
			}
		}
	}
}

func notImportedRunFirstPass(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			process(file, n)
			return true
		})
	}
	return nil, nil
}
