// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rest was mostly inspired by github.com/spf13/cobra/doc
// but with some changes to match the expected formats and styles of our writers and tools.
package rest

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/mongodb/mongocli/internal/search"
	"github.com/spf13/cobra"
)

const (
	separator        = "-"
	defaultExtension = ".txt"
)

// GenReSTTree generates the docs for the full tree of commands
func GenReSTTree(cmd *cobra.Command, dir string) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenReSTTree(c, dir); err != nil {
			return err
		}
	}

	basename := strings.ReplaceAll(cmd.CommandPath(), " ", separator) + defaultExtension
	filename := filepath.Join(dir, basename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := GenReSTCustom(cmd, f); err != nil {
		return err
	}
	return nil
}

const toc = `
.. default-domain:: mongodb

.. contents:: On this page
   :local:
   :backlinks: none
   :depth: 1
   :class: singlecol
`

const syntaxHeader = `Syntax
------

.. code-block::
`

const examplesHeader = `Examples
--------

.. code-block::
`

const tocHeader = `
.. toctree::
   :titlesonly:
`

// GenReSTCustom creates custom reStructured Text output.
// Adapted from github.com/spf13/cobra/doc to match MongoDB tooling and style
func GenReSTCustom(cmd *cobra.Command, w io.Writer) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	name := cmd.CommandPath()

	ref := strings.ReplaceAll(name, " ", separator)

	buf.WriteString(".. _" + ref + ":\n\n")
	buf.WriteString(strings.Repeat("=", len(name)) + "\n")
	buf.WriteString(name + "\n")
	buf.WriteString(strings.Repeat("=", len(name)) + "\n")
	buf.WriteString(toc)
	buf.WriteString("\n" + cmd.Short + "\n")
	if cmd.Long != "" {
		buf.WriteString("\n" + cmd.Long + "\n")
	}
	buf.WriteString("\n")

	if cmd.Runnable() {
		buf.WriteString(syntaxHeader)
		buf.WriteString(fmt.Sprintf("\n   %s\n\n", strings.ReplaceAll(cmd.UseLine(), "[flags]", "[options]")))
	}
	printArgsReST(buf, cmd)
	printOptionsReST(buf, cmd)

	if len(cmd.Example) > 0 {
		buf.WriteString(examplesHeader)
		buf.WriteString(fmt.Sprintf("\n%s\n\n", indentString(cmd.Example, " ")))
	}

	if hasRelatedCommands(cmd) {
		buf.WriteString("Related Commands\n")
		buf.WriteString("--------\n\n")

		children := cmd.Commands()
		sort.Sort(byName(children))

		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			cname := name + " " + child.Name()
			ref = strings.ReplaceAll(cname, " ", separator)
			buf.WriteString(fmt.Sprintf("* :ref:`%s` - %s\n", ref, child.Short))
		}
		buf.WriteString("\n")
	}
	if _, ok := cmd.Annotations["toc"]; ok || !cmd.Runnable() {
		buf.WriteString(tocHeader)
		buf.WriteString("\n")
		children := cmd.Commands()
		sort.Sort(byName(children))

		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			ref = strings.ReplaceAll(child.Name(), " ", separator)
			buf.WriteString(fmt.Sprintf("   /reference/%s\n", ref))
		}
		buf.WriteString("\n")
	}

	if !cmd.DisableAutoGenTag {
		buf.WriteString("*Auto generated by MongoDB CLI on " + time.Now().Format("2-Jan-2006") + "*\n")
	}
	_, err := buf.WriteTo(w)
	return err
}

// Test to see if we have a reason to print See Also information in docs
// Basically this is a test for a parent command or a subcommand which is
// both not deprecated and not the autogenerated help command.
func hasRelatedCommands(cmd *cobra.Command) bool {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		return true
	}
	return false
}

type byName []*cobra.Command

func (s byName) Len() int           { return len(s) }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }

const optionsHeader = `.. list-table::
   :header-rows: 1
   :widths: 20 10 10 60

   * - Name
     - Type
     - Required
     - Description
`

func printArgsReST(buf *bytes.Buffer, cmd *cobra.Command) {
	if args, ok := cmd.Annotations["args"]; ok {
		buf.WriteString("Arguments\n")
		buf.WriteString("---------\n\n")
		buf.WriteString(optionsHeader)
		var requiredSlice []string
		if requiredArgs, hasRequired := cmd.Annotations["requiredArgs"]; hasRequired {
			requiredSlice = strings.Split(requiredArgs, ",")
		}

		for _, arg := range strings.Split(args, ",") {
			required := search.StringInSlice(requiredSlice, arg)
			description := cmd.Annotations[arg+"Desc"]
			line := fmt.Sprintf("   * - %s\n     - string\n     - %v\n     - %s", arg, required, description)
			buf.WriteString(line)
		}
		buf.WriteString("\n\n")
	}
}

func printOptionsReST(buf *bytes.Buffer, cmd *cobra.Command) {
	flags := cmd.NonInheritedFlags()
	if flags.HasAvailableFlags() {
		buf.WriteString("Options\n")
		buf.WriteString("-------\n\n")
		buf.WriteString(optionsHeader)
		buf.WriteString(indentString(FlagUsages(flags), " "))
		buf.WriteString("\n")
	}

	parentFlags := cmd.InheritedFlags()
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("Inherited Options\n")
		buf.WriteString("-----------------\n\n")
		buf.WriteString(optionsHeader)
		buf.WriteString(indentString(FlagUsages(parentFlags), " "))
		buf.WriteString("\n")
	}
}

// adapted from: https://github.com/kr/text/blob/main/indent.go
func indentString(s, p string) string {
	var res []byte
	b := []byte(s)
	prefix := []byte(p)
	bol := true
	for _, c := range b {
		if bol && c != '\n' {
			res = append(res, prefix...)
		}
		res = append(res, c)
		bol = c == '\n'
	}
	return string(res)
}
