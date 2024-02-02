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

const (
	cobraCommandTypeName             = "*github.com/spf13/cobra.Command"
	testingTypeName                  = "*testing.T"
	testPrefix                       = "Test"
	verifyOutputTemplateMod          = "test"
	verifyOutputTemplateName         = "VerifyOutputTemplate"
	verifyOutputTemplateNumberOfArgs = 3
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

type NamedStructInfo struct {
	namedStruct *types.Named
	structInfo  *types.Struct
}

func LoadCommandBuilderInfos(packages []*packages.Package) []*CommandBuilderInfo {
	commandBuilderFuncDecls := findCommandBuilderFuncDecl(packages)

	builderFuncs := make([]*CommandBuilderInfo, 0)

	for _, commandBuilderFunc := range commandBuilderFuncDecls {
		fun, err := newCmdBuilderInfo(commandBuilderFunc)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		builderFuncs = append(builderFuncs, fun)
	}

	return builderFuncs
}

// Iterate through all to find all command builder functions
func findCommandBuilderFuncDecl(packages []*packages.Package) []*CommandBuilderFunc {
	builderFuncs := make([]*CommandBuilderFunc, 0)

	// Loop through all declarations in every package, file
	for _, pkg := range packages {
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

	relatedTemplateType, err := getRelatedTemplateType(pkg, commandOptsStruct)
	if err != nil {
		return nil, errors.New("could not determine related template type")
	}

	templateValue, err := getStringIdentValue(pkg, templateIdent)
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
// In `internal/cli/atlas/accesslists/list.go`, this would be `ListOpts`
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
			if structHasComposite(typeStructType, "cli", "OutputOpts") {
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

//nolint:gocyclo
func getRelatedTemplate(pkg *packages.Package, fun *ast.FuncDecl, argsIdent *ast.Ident) (*ast.Ident, error) {
	templateIdents := make([]*ast.Ident, 0)
	var err error

	ast.Inspect(fun, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.CallExpr:
			args := t.Args
			funExpr, ok := t.Fun.(*ast.SelectorExpr)
			if !ok || funExpr.Sel.Name != "InitOutput" || len(args) != 2 {
				return true
			}

			funIdent, ok := funExpr.X.(*ast.Ident)
			if !ok {
				return true
			}

			if funIdent.Obj != argsIdent.Obj {
				return true
			}

			ok = false

			switch templateArg := args[1].(type) {
			case *ast.Ident:
				templateIdents = append(templateIdents, templateArg)
				ok = true
			case *ast.BasicLit:
				if !(strings.Contains(templateArg.Value, "{{") && strings.Contains(templateArg.Value, "}}")) {
					ok = true
				}
			}

			if !ok {
				err = fmt.Errorf("unsupported argument in package: %v", pkg.Fset.Position(args[1].Pos()))
				return true
			}

		case *ast.AssignStmt:
			for n, lhs := range t.Lhs {
				selectorExpr, ok := lhs.(*ast.SelectorExpr)
				if !ok || selectorExpr.Sel == nil || selectorExpr.Sel.Name != "Template" {
					continue
				}

				lhsIdent, ok := selectorExpr.X.(*ast.Ident)
				if !ok {
					continue
				}

				if lhsIdent.Obj != argsIdent.Obj {
					continue
				}

				rhs := t.Rhs[n]
				rhsIdent, ok := rhs.(*ast.Ident)
				if !ok {
					err = fmt.Errorf("unsupported argument in package: %v", pkg.Fset.Position(rhs.Pos()))
					continue
				}

				templateIdents = append(templateIdents, rhsIdent)
			}
		}

		return true
	})

	templateIdentsLen := len(templateIdents)
	if templateIdentsLen != 1 {
		if err != nil {
			fmt.Println(err)
		}

		return nil, fmt.Errorf("expected 1 template, got %v", templateIdentsLen)
	}

	return templateIdents[0], nil
}

func getRelatedTemplateFromCallExpr(*ast.CallExpr) (*ast.Ident, error) {

}

// Look for all methods on commandOptsStruct which call `Run` and get the type of that parameter
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
				if arg.Obj == nil || arg.Obj.Decl == nil {
					errs = append(errs, fmt.Errorf("found ident but obj declaration is nil"))
					return true
				}

				argAssignStmt, ok := arg.Obj.Decl.(*ast.AssignStmt)
				if !ok {
					errs = append(errs, fmt.Errorf("found ident but obj declaration is not of the type *ast.AssignStmt"))
					return true
				}

				argType, err := getReturnTypeOfMethodReturningErrorTuple(pkg, argAssignStmt)
				if err != nil {
					errs = append(errs, fmt.Errorf("failed to find return type: %v", err))
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

	if s == nil && len(errs) > 0 {
		return nil, errs[0]
	}

	return s, nil
}

func getStringIdentValue(pkg *packages.Package, ident *ast.Ident) (string, error) {

	if ident.Obj == nil || ident.Obj.Decl == nil {
		return "", errors.New("ident.Obj.Decl == nil")
	}

	valueSpec, ok := ident.Obj.Decl.(*ast.ValueSpec)
	if !ok {
		return "", errors.New("ident.Obj.Decl is not a *ast.ValueSpec")
	}

	if len(valueSpec.Values) != 1 {
		return "", errors.New("ident.Obj.Decl.Values expecting exactly 1 value")
	}

	basicLit, ok := valueSpec.Values[0].(*ast.BasicLit)
	if !ok {
		return "", errors.New("ident.Obj.Decl.Values[0] is not an *ast.BasicLit")
	}

	if basicLit.Kind != token.STRING {
		return "", errors.New("ident.Obj.Decl.Values[0] is not a string literal")
	}

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
		if results.Len() != 2 {
			return nil, errors.New("expecting 2 return parameters")
		}

		returnType, err := getUnderlyingStruct(results.At(0).Type())
		if err != nil {
			return nil, err
		}

		return returnType, nil
	}

	return nil, nil
}

func getUnderlyingStruct(v types.Type) (*NamedStructInfo, error) {
	switch returnType := v.(type) {
	case *types.Pointer:
		return getUnderlyingStruct(returnType.Elem())
	case *types.Named:
		underlyingStruct, ok := returnType.Underlying().(*types.Struct)
		if !ok {
			return nil, errors.New("underlying type of named type is not a struct")
		}

		return &NamedStructInfo{
			namedStruct: returnType,
			structInfo:  underlyingStruct,
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
