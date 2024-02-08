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

package astparsing

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"
)

type NamedStructInfo struct {
	NamedStruct *types.Named
	StructInfo  *types.Struct
}

// Checks if a struct has a composite field
//
// Example:
//
//	type Foo struct {
//	    bar.Baz
//	}
//
// structHasComposite.
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

// Get the string value of an ident.
func getStringIdentValue(ident *ast.Ident) (string, error) {
	// Make sure the declaration is not nil
	if ident.Obj == nil || ident.Obj.Decl == nil {
		return "", errors.New("ident.Obj.Decl == nil")
	}

	// Make sure the declaration is of the type valueSpec
	valueSpec, ok := ident.Obj.Decl.(*ast.ValueSpec)
	if !ok {
		return "", errors.New("ident.Obj.Decl is not a *ast.ValueSpec")
	}

	// Make sure we're only receiving one value spec
	if len(valueSpec.Values) != 1 {
		return "", errors.New("ident.Obj.Decl.Values expecting exactly 1 value")
	}

	// Make sure that's a literal assignment (aka const)
	basicLit, ok := valueSpec.Values[0].(*ast.BasicLit)
	if !ok {
		return "", errors.New("ident.Obj.Decl.Values[0] is not an *ast.BasicLit")
	}

	// Verify that it's a string literal being assigned
	if basicLit.Kind != token.STRING {
		return "", errors.New("ident.Obj.Decl.Values[0] is not a string literal")
	}

	// Get rid of the quotes around the string, can be `, ' or "
	templateValue := strings.TrimFunc(basicLit.Value, func(r rune) bool {
		return slices.Contains(stringDelimiters, r)
	})

	return templateValue, nil
}

func getReturnTypeOfMethodReturningErrorTuple(pkg *packages.Package, argAssignStmt *ast.AssignStmt) (*NamedStructInfo, error) {
	// For now support something in the form of:
	// args, err := foo.bar.Method(args...)
	if len(argAssignStmt.Lhs) == 2 && len(argAssignStmt.Rhs) == 1 {
		callExpr, ok := argAssignStmt.Rhs[0].(*ast.CallExpr)
		if !ok {
			return nil, errors.New("not a call expression statement")
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil, errors.New("not a call selector expression statement")
		}

		selection, ok := pkg.TypesInfo.Selections[selectorExpr]
		if !ok {
			return nil, errors.New("selection not found")
		}

		methodSignature, ok := selection.Type().(*types.Signature)
		if !ok {
			return nil, errors.New("was expecting to find a method signature")
		}

		results := methodSignature.Results()
		const numberOfExpectedResults = 2
		if results.Len() != numberOfExpectedResults {
			return nil, errors.New("expecting 2 return parameters")
		}

		returnType, err := GetUnderlyingStruct(results.At(0).Type())
		if err != nil {
			return nil, err
		}

		return returnType, nil
	}

	return nil, nil
}

func GetUnderlyingStruct(v types.Type) (*NamedStructInfo, error) {
	switch returnType := v.(type) {
	case *types.Pointer:
		return GetUnderlyingStruct(returnType.Elem())
	case *types.Named:
		underlyingStruct, ok := returnType.Underlying().(*types.Struct)
		if !ok {
			return nil, errors.New("underlying type of named type is not a struct")
		}

		return &NamedStructInfo{
			NamedStruct: returnType,
			StructInfo:  underlyingStruct,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v",
			returnType.String())
	}
}

func getStructMethods(pkg *packages.Package, commandOptsStruct *ast.TypeSpec) []*ast.FuncDecl {
	methods := make([]*ast.FuncDecl, 0)

	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok || funcDecl.Recv == nil || funcDecl.Recv.List == nil {
				return true
			}

			receivers := funcDecl.Recv.List
			if len(receivers) != 1 {
				return true
			}

			receiver, ok := receivers[0].Type.(*ast.StarExpr)
			if !ok {
				return true
			}

			receiverTypeIdent, ok := receiver.X.(*ast.Ident)
			if !ok {
				return true
			}

			if receiverTypeIdent.Obj == commandOptsStruct.Name.Obj {
				methods = append(methods, funcDecl)
			}

			return true
		})
	}

	return methods
}
