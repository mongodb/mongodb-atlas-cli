package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

func main() {
	if err := inner(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

const (
	commandTypeName                  = "*github.com/spf13/cobra.Command"
	testingTypeName                  = "*testing.T"
	testPrefix                       = "Test"
	verifyOutputTemplateMod          = "test"
	verifyOutputTemplateName         = "VerifyOutputTemplate"
	verifyOutputTemplateNumberOfArgs = 3
)

func inner() error {
	conf := &packages.Config{
		Mode: packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		Tests: true,
	}

	sourcePath, err := filepath.Abs("../../internal/cli/")
	if err != nil {
		return err
	}

	filesPerDirectory, err := getSourceFilesPerDirectory(sourcePath)
	if err != nil {
		return err
	}

	pkgs := make([]*packages.Package, 0)

	for directory, files := range filesPerDirectory {
		fmt.Printf("Loading dir: %v\n", directory)
		pkgsInDir, err := packages.Load(conf, files...)

		if err != nil {
			return err
		}

		pkgs = append(pkgs, pkgsInDir...)
	}

	builderFuncs := findBuilderFuncs(pkgs)

	for _, builderFunc := range builderFuncs {
		p := builderFunc.pkg

		if !strings.HasSuffix(p.ID, ".test]") {
			continue
		}

		templateGotTested := builderFunc.isTested()

		fmt.Print("[")
		if !templateGotTested {
			fmt.Print("N")
		}

		fmt.Printf("OK] %v\t\t%v", builderFunc.templateIdent, builderFunc.pkg.Fset.Position(builderFunc.templateIdent.NamePos))
		fmt.Println()
	}

	return nil
}

type BuilderFunc struct {
	pkg               *packages.Package
	commandOptsStruct *ast.TypeSpec
	templateIdent     *ast.Ident
}

//nolint:gocyclo
func (builderFunc *BuilderFunc) isTested() bool {
	p := builderFunc.pkg
	templateGotTested := false

	for _, f := range p.Syntax {
		for _, n := range f.Decls {
			fun, ok := n.(*ast.FuncDecl)
			if !ok || !strings.HasPrefix(fun.Name.Name, testPrefix) {
				continue
			}

			signature, ok := p.TypesInfo.Defs[fun.Name].Type().(*types.Signature)
			if !ok {
				continue
			}

			params := signature.Params()
			if params.Len() == 0 {
				continue
			}

			firstParameter := params.At(0).Type()
			if firstParameter.String() != testingTypeName {
				continue
			}

			ast.Inspect(fun.Body, func(n ast.Node) bool {
				callExpr, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}

				selector, ok := callExpr.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				ident, ok := selector.X.(*ast.Ident)
				if !ok {
					return true
				}

				if ident.Name != verifyOutputTemplateMod ||
					selector.Sel.Name != verifyOutputTemplateName {
					return true
				}

				if len(callExpr.Args) != verifyOutputTemplateNumberOfArgs {
					return true
				}

				templateIdent, ok := callExpr.Args[1].(*ast.Ident)
				if !ok {
					return true
				}

				templateInstance := p.TypesInfo.Instances[templateIdent]
				builderFuncTemplateInstance := p.TypesInfo.Instances[builderFunc.templateIdent]

				if templateInstance != builderFuncTemplateInstance {
					return true
				}

				templateGotTested = true

				return true
			})
		}
	}

	return templateGotTested
}

func getSourceFilesPerDirectory(sourcePath string) (sourceFiles map[string][]string, err error) {
	sourceFiles = make(map[string][]string)

	err = filepath.WalkDir(sourcePath, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(s); ext == ".go" {
			dir := filepath.Dir(s)

			if sourceFiles[dir] == nil {
				sourceFiles[dir] = make([]string, 0)
			}

			sourceFiles[dir] = append(sourceFiles[dir], s)
		}
		return nil
	})

	return
}

func findBuilderFuncs(packages []*packages.Package) []*BuilderFunc {
	builderFuncs := make([]*BuilderFunc, 0)

	for _, p := range packages {
		if !strings.HasSuffix(p.ID, ".test]") {
			continue
		}

		for _, f := range p.Syntax {
			for _, n := range f.Decls {
				t, ok := n.(*ast.FuncDecl)

				if !ok {
					continue
				}

				fun, err := parseCmdBuilderFunc(p, t)
				if err != nil {
					// fmt.Printf("Error: %v\n", err)
					continue
				}

				builderFuncs = append(builderFuncs, fun)
			}
		}
	}

	return builderFuncs
}

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
		for i, rhs := range assignStatement.Rhs {
			var maybeTypeIdent any

			switch rhs := rhs.(type) {
			case *ast.UnaryExpr:
				x, ok := rhs.X.(*ast.CompositeLit)
				if !ok {
					continue
				}

				maybeTypeIdent = x.Type
			case *ast.CompositeLit:
				maybeTypeIdent = rhs.Type
			case *ast.CallExpr:
				funcName, ok := rhs.Fun.(*ast.Ident)
				if !ok || funcName.Name != "new" {
					continue
				}

				args := rhs.Args
				if len(args) != 1 {
					continue
				}

				maybeTypeIdent = args[0]
			}

			typeIdent, ok := maybeTypeIdent.(*ast.Ident)
			if !ok {
				continue
			}

			if typeIdent.Obj == nil || typeIdent.Obj.Decl == nil {
				continue
			}

			typeSpec, ok := typeIdent.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				continue
			}

			typeStructType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			leftHandIdent, ok := assignStatement.Lhs[i].(*ast.Ident)
			if !ok {
				continue
			}

			for _, field := range typeStructType.Fields.List {
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

				if packageIdent.Name != "cli" || typeSelector.Sel.Name != "OutputOpts" {
					continue
				}

				return leftHandIdent, typeSpec, nil
			}
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

func parseCmdBuilderFunc(pkg *packages.Package, fun *ast.FuncDecl) (*BuilderFunc, error) {
	signature, ok := pkg.TypesInfo.Defs[fun.Name].Type().(*types.Signature)
	if !ok {
		return nil, errors.New("not a signature")
	}

	results := signature.Results()
	if results.Len() != 1 {
		return nil, errors.New("expecting 1 return parameter")
	}

	returnType := results.At(0)
	if returnType.Type().String() != commandTypeName {
		return nil, errors.New("command builder should return a command")
	}

	optsVariable, commandOptsStruct, err := getCommandOpts(fun)
	if err != nil {
		return nil, errors.New("could not determine command opts")
	}

	templateIdent, err := getRelatedTemplate(pkg, fun, optsVariable)
	if err != nil {
		return nil, fmt.Errorf("err getting template ident: %w, %v", err, pkg.Fset.Position(fun.Pos()))
	}

	return &BuilderFunc{
		pkg:               pkg,
		commandOptsStruct: commandOptsStruct,
		templateIdent:     templateIdent,
	}, nil
}
