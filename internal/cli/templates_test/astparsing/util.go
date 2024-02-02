package astparsing

import "go/ast"

// Checks if a struct has a composite field
//
// Example:
//
//	type Foo struct {
//	    bar.Baz
//	}
//
// structHasComposite
func structHasComposite(structType *ast.StructType, module string, name string) bool {
	for _, field := range structType.Fields.List {
		if len(field.Names) != 0 {
			continue
		}

		typeSelector, ok := field.Type.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		packageIdent, ok := typeSelector.X.(*ast.Ident)
		if !ok {
			continue
		}

		if packageIdent.Name != module || typeSelector.Sel.Name != name {
			continue
		}

		return true
	}

	return false
}
