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
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	cobraCommandTypeName = "*github.com/spf13/cobra.Command"
)

var stringDelimiters = []rune{'`', '\'', '"'}

type CommandBuilderFunc struct {
	pkg     *packages.Package
	funDecl *ast.FuncDecl
}

type CommandBuilderInfo struct {
	Pkg               *packages.Package
	CommandOptsStruct *ast.TypeSpec
	TemplateIdent     *ast.Ident
	TemplateType      *NamedStructInfo
	TemplateValue     string
}

func LoadCommandBuilderInfos(p []*packages.Package) []*CommandBuilderInfo {
	commandBuilderFuncDecls := findCommandBuilderFuncDecl(p)

	builderFuncs := make([]*CommandBuilderInfo, 0)

	for _, commandBuilderFunc := range commandBuilderFuncDecls {
		fun, err := newCmdBuilderInfo(commandBuilderFunc)
		if err != nil {
			// the method returns an error when the command is not compatible
			// if we get an error it means we can ignore this function
			continue
		}

		builderFuncs = append(builderFuncs, fun)
	}

	return builderFuncs
}

// Iterate through all to find all command builder functions.
func findCommandBuilderFuncDecl(p []*packages.Package) []*CommandBuilderFunc {
	builderFuncs := make([]*CommandBuilderFunc, 0)

	// Loop through all declarations in every package, file
	for _, pkg := range p {
		for _, file := range pkg.Syntax {
			// Loop through all declarations, this can be variable, function, ... declarations
			for _, declaration := range file.Decls {
				// We only care about function declarations
				funcDecl, ok := declaration.(*ast.FuncDecl)

				if !ok {
					continue
				}

				// Look up the signature of the function that's being declared
				signature, ok := pkg.TypesInfo.Defs[funcDecl.Name].Type().(*types.Signature)
				if !ok {
					continue
				}

				// We're searching for functions with a signature like this:
				// func FUNCTIONNAME(OPTIONAL) *cobra.Command {
				// First verify that the function only returns 1 object
				results := signature.Results()
				if results.Len() != 1 {
					continue
				}

				// Verify that the return type = *cobra.Command
				returnType := results.At(0)
				if returnType.Type().String() != cobraCommandTypeName {
					continue
				}

				// We got a match, add it ot the list
				builderFuncs = append(builderFuncs, &CommandBuilderFunc{
					pkg:     pkg,
					funDecl: funcDecl,
				})
			}
		}
	}

	return builderFuncs
}

// Convert the command builder function into something we can use to test our templates
//
// Example code I'll refer to in the commands to make things easier to understand
// const listTemplate = `{{range .}}{{.childprop}}`
//
//	type ListOpts struct {
//		cli.OutputOpts
//		/* other fields */
//	}
//
//	func (opts *ListOpts) Run() error {
//	 /* other code */
//		return opts.Print(r)
//	}
//
//	func ListBuilder() *cobra.Command {
//		opts := &ListOpts{}
//	    opts.template = listTemplate // OPTION 1
//
//		cmd := &cobra.Command{
//			PreRunE: func(cmd *cobra.Command, args []string) error {
//				return opts.PreRunE(
//					/* other code */
//					opts.InitOutput(cmd.OutOrStdout(), listTemplate), // OPTION 2
//				)
//			},
//			/* other fields */
//		}
//	 /* other code */
//	}
func newCmdBuilderInfo(commandBuilderFunc *CommandBuilderFunc) (*CommandBuilderInfo, error) {
	pkg, fun := commandBuilderFunc.pkg, commandBuilderFunc.funDecl

	// Try to find the opts variable in our command
	// This will return:
	// - ident pointing to the opts variable, in the example this is the "opts" variable in the ListBuilder function
	// - struct info about the opts variable, in the example this would be the struct info about `ListOpts`
	optsVariableIdent, commandOptsStruct, err := getCommandOpts(fun)
	if err != nil {
		return nil, errors.New("could not determine command opts")
	}

	// Find the ident pointing to the template constant.
	// In the example this would be `listTemplate`
	templateIdent, err := getRelatedTemplate(pkg, fun, optsVariableIdent)
	if err != nil {
		return nil, fmt.Errorf("err getting template ident: %w, %v", err, pkg.Fset.Position(fun.Pos()))
	}

	// Find the template type that's being passed to the opts.run method
	// In the example this would be the type of the variable `r` we pass to `opts.Print` in `opts.Run`
	relatedTemplateType, err := getRelatedTemplateType(pkg, commandOptsStruct)
	if err != nil {
		return nil, errors.New("could not determine related template type")
	}

	// Get the string value of the template
	templateValue, err := getStringIdentValue(templateIdent)
	if err != nil {
		return nil, fmt.Errorf("could not find template value: %w, %v", err, pkg.Fset.Position(fun.Pos()))
	}

	return &CommandBuilderInfo{
		Pkg:               pkg,
		CommandOptsStruct: commandOptsStruct,
		TemplateIdent:     templateIdent,
		TemplateType:      relatedTemplateType,
		TemplateValue:     templateValue,
	}, nil
}

// Find to the opts struct which is used in this command
//
// Example:
// In `internal/cli/atlas/accesslists/list.go`, this would be `ListOpts`.
//
//nolint:gocyclo
func getCommandOpts(fun *ast.FuncDecl) (*ast.Ident, *ast.TypeSpec, error) {
	// Search for all asignment statements which assign a variable which implements the OutputOpts interface
	for _, stmt := range fun.Body.List {
		// We're only interested in assign statements
		assignStatement, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		// Look at the type that's being assigned
		// We're looking for an assignment that looks like this:
		// - opts := &ListOpts{}
		// - opts := ListOpts{}
		// - opts := new(ListOpts)
		for i, rhs := range assignStatement.Rhs {
			// This variable will hold the Opts type ident
			var maybeTypeIdent any

			switch rhs := rhs.(type) {
			// This case handles: opts := &ListOpts{}
			case *ast.UnaryExpr:
				// We are interested in the ListOpts{} part
				x, ok := rhs.X.(*ast.CompositeLit)
				if !ok {
					continue
				}

				// Get the type ident
				maybeTypeIdent = x.Type

			// This case handles: opts := ListOpts{}
			case *ast.CompositeLit:
				// Get the type ident
				maybeTypeIdent = rhs.Type

			// This case handles: opts := new(ListOpts)
			case *ast.CallExpr:
				// Verify that the call that's being done is to the "new" function
				funcName, ok := rhs.Fun.(*ast.Ident)
				if !ok || funcName.Name != "new" {
					continue
				}

				// Get the arguments of the function and verify that there's only one
				args := rhs.Args
				if len(args) != 1 {
					continue
				}

				// Get the type of the first argument
				maybeTypeIdent = args[0]
			}

			// Try to convert the "maybeTypeIdent" to an actual ident, if it fails: continue
			typeIdent, ok := maybeTypeIdent.(*ast.Ident)
			if !ok {
				continue
			}

			// Make sure we can read how the type is declared, if not: continue
			if typeIdent.Obj == nil || typeIdent.Obj.Decl == nil {
				continue
			}

			// Make that the declaration is a type specification, if not: continue
			typeSpec, ok := typeIdent.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Make that the type which is being declared is a struct, if not: continue
			typeStructType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Get the ident of the variable which the opts are assigned to, if not found: continue
			leftHandIdent, ok := assignStatement.Lhs[i].(*ast.Ident)
			if !ok {
				continue
			}

			// Verify that the struct has "cli.OutputOpts" as a composite field, if not: continue
			if !structHasComposite(typeStructType, "cli", "OutputOpts") {
				continue
			}

			// When all succeeds we return the:
			// - ident which points to the variable (opts in opts:=&ListOpts{})
			// - struct type information about the opt variable
			return leftHandIdent, typeSpec, nil
		}
	}

	return nil, nil, errors.New("no command opts found")
}

func getRelatedTemplate(pkg *packages.Package, fun *ast.FuncDecl, argsIdent *ast.Ident) (*ast.Ident, error) {
	templateIdents := make([]*ast.Ident, 0)
	var err error

	ast.Inspect(fun, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.CallExpr:
			// Get the ident to the template based on a method call to opts.InitOutput.
			ident, e := getRelatedTemplateFromCallExpr(pkg, t, argsIdent)

			if e != nil {
				err = nil
			}

			if ident != nil {
				templateIdents = append(templateIdents, ident)

				// stop looking deeper in this specific node (call node)
				return false
			}

		case *ast.AssignStmt:
			// Get the idents to the template based on an asignment.
			// Example: opt.Template = template;
			idents, e := getRelatedTemplateFromAssignStmt(pkg, t, argsIdent)

			if e != nil {
				err = nil
			}

			templateIdents = append(templateIdents, idents...)
		}

		return true
	})

	templateIdentsLen := len(templateIdents)

	// When there's not exactly one tempate, throw and error
	if templateIdentsLen != 1 {
		// If there's an error, throw that one
		if err != nil {
			fmt.Println(err)
		}

		// If there's no error, show the templates that are not matching
		return nil, fmt.Errorf("expected 1 template, got %v", templateIdentsLen)
	}

	// Return the ident
	return templateIdents[0], nil
}

// Get the ident to the template based on a method call to opts.InitOutput.
func getRelatedTemplateFromCallExpr(pkg *packages.Package, callExpr *ast.CallExpr, argsIdent *ast.Ident) (*ast.Ident, error) {
	args := callExpr.Args
	funExpr, ok := callExpr.Fun.(*ast.SelectorExpr)

	// make sure the function is called `InitOutput`
	// make sure the function takes 2 arguments
	if !ok || funExpr.Sel.Name != "InitOutput" || len(args) != 2 {
		return nil, nil
	}

	// try to get X as an ident, if that fails, return nil
	funIdent, ok := funExpr.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	// make sure that the ident matches the argument ident
	// we're basically testing that we call InitOutput on the variable `argsIdent`
	if funIdent.Obj != argsIdent.Obj {
		return nil, nil
	}

	// Take the second argument that's passed to the method
	// that is the argument that contains the template
	switch templateArg := args[1].(type) {
	// If it's an ident (variable/constant), return the ident
	case *ast.Ident:
		return templateArg, nil

	// If there's a string literal that's not a template ignore it
	case *ast.BasicLit:
		if !(strings.Contains(templateArg.Value, "{{") && strings.Contains(templateArg.Value, "}}")) {
			return nil, nil
		}
	}

	// Any other argument is not supported, return an error
	return nil, fmt.Errorf("unsupported argument in package: %v", pkg.Fset.Position(args[1].Pos()))
}

// Get the ident to the template based on a method call to opts.InitOutput.
func getRelatedTemplateFromAssignStmt(pkg *packages.Package, t *ast.AssignStmt, argsIdent *ast.Ident) ([]*ast.Ident, error) {
	templateIdents := make([]*ast.Ident, 0)
	var err error

	// An asignment can be multiple things being assigned to multiple variables
	// we only care what's being assigned to arg.Template
	for n, lhs := range t.Lhs {
		// Make sure the left hand side of the assignment is a selector expression (aka variable.field)
		// Verify that the field that's being selected is called "Template"
		selectorExpr, ok := lhs.(*ast.SelectorExpr)
		if !ok || selectorExpr.Sel == nil || selectorExpr.Sel.Name != "Template" {
			continue
		}

		// Make sure that X is an ident
		lhsIdent, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			continue
		}

		// Make sure that the variable that we're selecting the template field from is the argument ident we're searching for
		if lhsIdent.Obj != argsIdent.Obj {
			continue
		}

		// Get the matching right hand side and make sure it's an ident
		rhs := t.Rhs[n]
		rhsIdent, ok := rhs.(*ast.Ident)
		if !ok {
			err = fmt.Errorf("unsupported argument in package: %v", pkg.Fset.Position(rhs.Pos()))
			continue
		}

		// In this case, add the ident
		templateIdents = append(templateIdents, rhsIdent)
	}

	return templateIdents, err
}

// Look for all methods on commandOptsStruct which call `Run` and get the type of that parameter.
//
//nolint:gocyclo
func getRelatedTemplateType(pkg *packages.Package, commandOptsStruct *ast.TypeSpec) (*NamedStructInfo, error) {
	// Find all methods on commands
	methods := getStructMethods(pkg, commandOptsStruct)
	errs := make([]error, 0)
	var s *NamedStructInfo

	// Find the run method
	for _, method := range methods {
		ast.Inspect(method, func(n ast.Node) bool {
			// From here, verify that we're calling commandOptsStruct.Run([argument])
			// Make sure that we're inspecting a call expression
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Make sure that we call a function with the name print
			funSelectorExpression, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok || funSelectorExpression.Sel.Name != "Print" {
				return true
			}

			// Make sure that it's a method call
			objIdent, ok := funSelectorExpression.X.(*ast.Ident)
			if !ok || objIdent.Obj == nil || objIdent.Obj.Decl == nil {
				return true
			}

			// Make sure that the object that it's being called on is a field
			objField, ok := objIdent.Obj.Decl.(*ast.Field)
			if !ok {
				return true
			}

			// Make sure that the type of the object is a pointer
			objFieldStartExpr, ok := objField.Type.(*ast.StarExpr)
			if !ok {
				return true
			}

			// Convert the pointer type is a ident
			objFieldStartExprIdent, ok := objFieldStartExpr.X.(*ast.Ident)
			if !ok {
				return true
			}

			// Make sure the type of the object mathes the type of our commandOptsStruct
			if objFieldStartExprIdent.Obj != commandOptsStruct.Name.Obj {
				return true
			}

			// From here: Extract the argument type
			// Make sure that the method is passed one argument
			if len(callExpr.Args) != 1 {
				return true
			}

			switch arg := callExpr.Args[0].(type) {
			case *ast.Ident:
				// Make sure the declaration is not null
				if arg.Obj == nil || arg.Obj.Decl == nil {
					errs = append(errs, errors.New("found ident but obj declaration is nil"))
					return true
				}

				// Make sure that the declaration is assign statement
				argAssignStmt, ok := arg.Obj.Decl.(*ast.AssignStmt)
				if !ok {
					errs = append(errs, errors.New("found ident but obj declaration is not of the type *ast.AssignStmt"))
					return true
				}

				// For now we only support assignments which come from a function returning a tuple (typeWeNeed, error)
				argType, err := getReturnTypeOfMethodReturningErrorTuple(pkg, argAssignStmt)
				if err != nil {
					errs = append(errs, fmt.Errorf("failed to find return type: %w", err))
					return true
				}

				s = argType
				return false
			default:
				errs = append(errs, errors.New("expression type not supported"))
				return true
			}
		})
	}

	// If there's no return type and we do have an error, return the error
	if s == nil && len(errs) > 0 {
		return nil, errs[0]
	}

	return s, nil
}
