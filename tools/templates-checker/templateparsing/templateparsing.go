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

package templateparsing

import (
	"errors"
	"fmt"
	"reflect"
	"text/template"
	"text/template/parse"
)

var templateFuncs = template.FuncMap{
	// Defined in our codebase, expand if required.
	// The template parser in the standard library will fail if the function is not defined here
	"Year":              nop,
	"Now":               nop,
	"Join":              nop,
	"valueOrEmptySlice": nop,

	// BUILDIN
	"and":      nop,
	"call":     nop,
	"html":     nop,
	"index":    nop,
	"slice":    nop,
	"js":       nop,
	"len":      nop,
	"not":      nop,
	"or":       nop,
	"print":    nop,
	"printf":   nop,
	"println":  nop,
	"urlquery": nop,
	"eq":       nop, // ==
	"ge":       nop, // >=
	"gt":       nop, // >
	"le":       nop, // <=
	"lt":       nop, // <
	"ne":       nop, // !=
}

func nop() {}

type TemplateCallTree struct {
	listType   *TemplateCallTree
	structType *TemplateCallTreeEntry
}

func NewTemplateCallTree() *TemplateCallTree {
	return &TemplateCallTree{
		listType:   nil,
		structType: nil,
	}
}

func (c *TemplateCallTree) String() string {
	return c.Fprint(0)
}

func (c *TemplateCallTree) Fprint(depth int) string {
	out := ""

	if c.listType != nil {
		out += c.listType.Fprint(depth)
	}

	if c.structType != nil {
		out += c.structType.Fprint(depth)
	}

	return out
}

func (c *TemplateCallTree) List() *TemplateCallTree {
	if c.listType == nil {
		c.listType = NewTemplateCallTree()
	}

	return c.listType
}

func (c *TemplateCallTree) IsValid() bool {
	switch {
	case c.listType == nil && c.structType == nil:
		return true
	case c.listType != nil && c.structType == nil:
		return c.listType.IsValid()
	case c.listType == nil && c.structType != nil:
		return c.structType.IsValid()
	default:
		return false
	}
}

func (c *TemplateCallTree) Struct() *TemplateCallTreeEntry {
	if c.structType == nil {
		c.structType = NewTemplateCallTreeEntry()
	}

	return c.structType
}

type TemplateCallTreeEntry struct {
	fields map[string]*TemplateCallTree
}

func NewTemplateCallTreeEntry() *TemplateCallTreeEntry {
	return &TemplateCallTreeEntry{
		fields: make(map[string]*TemplateCallTree),
	}
}

func (c *TemplateCallTreeEntry) IsValid() bool {
	for _, t := range c.fields {
		if !t.IsValid() {
			return false
		}
	}

	return true
}

func (c *TemplateCallTreeEntry) Fprint(depth int) string {
	out := ""

	for key, value := range c.fields {
		if value.listType != nil {
			out += ident(fmt.Sprintf("- []%v:\n", key), depth)
			out += value.listType.Fprint(depth + 1)
		}
		if value.structType != nil {
			out += ident(fmt.Sprintf("- %v:\n", key), depth)
			out += value.structType.Fprint(depth + 1)
		}
		if value.listType == nil && value.structType == nil {
			out += ident(fmt.Sprintf("- %v\n", key), depth)
		}
	}

	return out
}

const spacesPerDepth = 4

func ident(value string, depth int) string {
	return fmt.Sprintf("%*s%s", depth*spacesPerDepth, "", value)
}

func ParseTemplate(t string) (*TemplateCallTree, error) {
	parseTree, err := parse.Parse("templ", t, "{{", "}}", templateFuncs)
	if err != nil {
		return nil, err
	}

	templ := parseTree["templ"]
	root := NewTemplateCallTree()

	if err := buildTreeRecursive(root, templ.Root); err != nil {
		return nil, err
	}

	return root, nil
}

type TreeBuilder func(*TemplateCallTree, parse.Node) error

func buildTreeRecursive(root *TemplateCallTree, node parse.Node) error {
	var nextInner TreeBuilder

	nextInner = buildTreeFunc(func(r *TemplateCallTree, n parse.Node) error {
		return nextInner(r, n)
	})

	return nextInner(root, node)
}

//nolint:gocyclo
func buildTreeFunc(next TreeBuilder) TreeBuilder {
	return func(root *TemplateCallTree, node parse.Node) error {
		if IsNil(node) {
			return nil
		}

		buildAndMergeTrees := func(nodes ...parse.Node) error {
			for _, node := range nodes {
				if err := next(root, node); err != nil {
					return err
				}
			}

			return nil
		}

		switch node := node.(type) {
		case *parse.ListNode:
			return buildAndMergeTrees(node.Nodes...)
		case *parse.RangeNode:
			mainRoot, err := buildRangeNodeMainRoot(root, node)
			if err != nil {
				return err
			}

			if err := next(mainRoot, node.List); err != nil {
				return err
			}

			if err := next(root, node.ElseList); err != nil {
				return err
			}
		case *parse.BranchNode:
			return buildAndMergeTrees(node.Pipe, node.List, node.ElseList)
		case *parse.PipeNode:
			indexCommandFields, err := parseIndexCommand(node)
			if err == nil {
				return buildTreeFunc(func(r *TemplateCallTree, _ parse.Node) error {
					return next(r.List(), nil)
				})(root, indexCommandFields)
			}

			for _, n := range node.Cmds {
				if err := buildAndMergeTrees(n); err != nil {
					return err
				}
			}

			for _, n := range node.Decl {
				if err := buildAndMergeTrees(n); err != nil {
					return err
				}
			}
		case *parse.ActionNode:
			return next(root, node.Pipe)
		case *parse.IfNode:
			return buildAndMergeTrees(&node.BranchNode)
		case *parse.CommandNode:
			return buildAndMergeTrees(node.Args...)
		case *parse.FieldNode:
			// Build the structure from a string array
			return next(buildFieldStructure(root, node.Ident), nil)
		case *parse.ChainNode:
			// Get the result of the first part of the chain and then chain the fields to it
			// Example: (index .field1.field2).chain1.chain2
			return buildTreeFunc(func(newRoot *TemplateCallTree, _ parse.Node) error {
				return next(buildFieldStructure(newRoot, node.Field), nil)
			})(root, node.Node)
		case *parse.DotNode:
			// NOOP outside of range
			return nil
		case *parse.TextNode:
			// we don't care about text nodes (plain text)
			return nil
		case *parse.StringNode:
			// we don't care about string nodes (string constants)
			return nil
		case *parse.IdentifierNode:
			// we don't care about identifier nodes, they always contain function names
			return nil
		case *parse.VariableNode:
			// this is the left side of an assignment, we don't care about variable names
			return nil
		case *parse.NumberNode:
			// we don't care about number constants nodes (string constants)
			return nil
		default:
			return errors.New("unsupported node type")
		}

		return nil
	}
}

func buildFieldStructure(root *TemplateCallTree, fields []string) *TemplateCallTree {
	c := root

	for _, ident := range fields {
		if c.Struct().fields[ident] == nil {
			t := NewTemplateCallTree()
			c.Struct().fields[ident] = t
			c = t
		} else {
			c = c.Struct().fields[ident]
		}
	}

	return c
}

func buildRangeNodeMainRoot(root *TemplateCallTree, node *parse.RangeNode) (*TemplateCallTree, error) {
	identifiers, err := pipelineToIdentifiers(node.Pipe)
	if err != nil {
		return nil, err
	}

	newRoot := root
	for _, identifier := range identifiers {
		r := NewTemplateCallTree()
		newRoot.Struct().fields[identifier] = r
		newRoot = r
	}

	return newRoot.List(), nil
}

func pipelineToIdentifiers(pipeline *parse.PipeNode) ([]string, error) {
	const noFuncSliceCmdCount = 1
	const funcWith1ArgCmdCount = 2

	if len(pipeline.Cmds) != 1 {
		return nil, errors.New("unsupported number of cmds, expected 1")
	}

	cmd := pipeline.Cmds[0]

	var arg parse.Node
	switch len(cmd.Args) {
	case noFuncSliceCmdCount:
		arg = cmd.Args[0]
	case funcWith1ArgCmdCount:
		funcIdent, ok := cmd.Args[0].(*parse.IdentifierNode)
		if !ok || funcIdent.Ident != "valueOrEmptySlice" {
			return nil, errors.New("only supporting calls to 'valueOrEmptySlice'")
		}

		arg = cmd.Args[1]
	default:
		return nil, errors.New("unsupported number of cmd args, expected 1. Or 2 when 'valueOrEmptySlice' is used")
	}

	switch arg := arg.(type) {
	case *parse.DotNode:
		return make([]string, 0), nil
	case *parse.FieldNode:
		return arg.Ident, nil
	default:
		return nil, errors.New("unsupported node type")
	}
}

func IsNil(i any) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func parseIndexCommand(pipeNode *parse.PipeNode) (*parse.FieldNode, error) {
	const indexCmdNumber = 1
	const indexCmdArgsNumber = 3

	if len(pipeNode.Decl) != 0 {
		return nil, errors.New("expected 0 decls")
	}

	if len(pipeNode.Cmds) != indexCmdNumber {
		return nil, errors.New("expected 3 cmds")
	}

	cmd := pipeNode.Cmds[0]
	if len(cmd.Args) != indexCmdArgsNumber {
		return nil, errors.New("expected 3 cmds")
	}

	indexIdentifier, identifierOk := cmd.Args[0].(*parse.IdentifierNode)
	fieldNode, fieldNodeOk := cmd.Args[1].(*parse.FieldNode)
	_, numberNodeOk := cmd.Args[2].(*parse.NumberNode)

	if !identifierOk || !fieldNodeOk || !numberNodeOk {
		return nil, errors.New("not the required arguments")
	}

	if indexIdentifier.Ident != "index" {
		return nil, errors.New("not a call to index")
	}

	return fieldNode, nil
}
