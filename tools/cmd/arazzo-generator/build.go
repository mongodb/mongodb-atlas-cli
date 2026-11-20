// Copyright 2026 MongoDB Inc
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

package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	transformPreviousMajorVersion = "previousMajorVersion"

	cleanupKindRun   = "run"
	cleanupKindRetry = "retry"
)

// buildTest resolves a workflow into a renderable test model.
func buildTest(workflow Workflow, index map[string]commandInfo) (*testModel, error) {
	steps, watches, err := mergeWatchSteps(workflow.Steps, index)
	if err != nil {
		return nil, err
	}

	refs := make(map[string]bool)
	for _, step := range steps {
		step.refs(refs)
	}

	inputs, optionalInputs, err := buildInputs(workflow.Inputs, refs)
	if err != nil {
		return nil, err
	}

	outputUsage := computeOutputUsage(steps, refs)
	stepModels, varDecls, err := buildSteps(steps, index, watches, optionalInputs, outputUsage)
	if err != nil {
		return nil, err
	}

	model := &testModel{
		TestName: "Test" + strcase.ToCamel(workflow.WorkflowID),
		Inputs:   inputs,
		VarDecls: varDecls,
		Steps:    stepModels,
	}
	for _, step := range stepModels {
		if len(step.Assertions) > 0 || len(step.DerivedStmts) > 0 {
			model.NeedsRequire = true
		}
	}

	return model, nil
}

func stepOutputRef(stepID, name string) string {
	return stepsPrefix + stepID + outputsInfix + name
}

// computeOutputUsage marks which step outputs and derived outputs later steps
// reference, so the generator only emits variables that are actually used.
func computeOutputUsage(steps []Step, refs map[string]bool) map[string]bool {
	usage := make(map[string]bool)
	for _, step := range steps {
		for name := range step.Outputs {
			usage[stepOutputRef(step.StepID, name)] = false
		}
		for name := range step.XDerivedOutputs {
			usage[stepOutputRef(step.StepID, name)] = false
		}
	}
	for _, step := range steps {
		for _, derived := range step.XDerivedOutputs {
			for _, part := range strings.Split(derived.Template, "{") {
				if expr, found := strings.CutSuffix(part, "}"); found {
					usage[strings.TrimSpace(expr)] = true
				}
			}
		}
	}
	for ref := range refs {
		if strings.HasPrefix(ref, stepsPrefix) {
			usage[ref] = true
		}
	}

	return usage
}

func buildSteps(
	steps []Step,
	index map[string]commandInfo,
	watches map[string]int,
	optionalInputs map[string]bool,
	outputUsage map[string]bool,
) ([]stepModel, []string, error) {
	cleanupsByProducer, err := buildCleanups(steps, index, watches)
	if err != nil {
		return nil, nil, err
	}

	stepModels := make([]stepModel, 0, len(steps))
	varDecls := make([]string, 0)
	for _, step := range steps {
		model, err := buildStep(step, index, watches, optionalInputs, outputUsage)
		if err != nil {
			return nil, nil, err
		}
		model.Cleanups = cleanupsByProducer[model.ID]
		stepModels = append(stepModels, model)

		varDecls = appendUsedOutputDecls(varDecls, step, outputUsage)
	}
	sort.Strings(varDecls)

	return stepModels, varDecls, nil
}

func appendUsedOutputDecls(varDecls []string, step Step, outputUsage map[string]bool) []string {
	for _, name := range sortedKeys(step.Outputs) {
		if outputUsage[stepOutputRef(step.StepID, name)] {
			varDecls = append(varDecls, sanitizeIdent(step.StepID+"_"+name)+" string")
		}
	}
	for _, name := range sortedKeys(step.XDerivedOutputs) {
		if outputUsage[stepOutputRef(step.StepID, name)] {
			varDecls = append(varDecls, sanitizeIdent(step.StepID+"_"+name)+" string")
		}
	}

	return varDecls
}

// buildCleanups turns steps marked with x-cleanup into registrations that run
// the same command best-effort when the test finishes. Each registration is
// attached to the latest step that produces the inputs it references, so the
// inputs are always set and Go's LIFO cleanup ordering releases dependent
// resources first.
func buildCleanups(steps []Step, index map[string]commandInfo, watches map[string]int) (map[string][]cleanupModel, error) {
	cleanupsByProducer := make(map[string][]cleanupModel)
	for _, step := range steps {
		if step.XCleanup == "" {
			continue
		}
		cleanup, producer, err := buildCleanup(step, steps, index, watches)
		if err != nil {
			return nil, err
		}
		cleanupsByProducer[producer] = append(cleanupsByProducer[producer], cleanup)
	}

	return cleanupsByProducer, nil
}

func buildCleanup(step Step, steps []Step, index map[string]commandInfo, watches map[string]int) (cleanupModel, string, error) {
	kind := cleanupKindRun
	switch step.XCleanup {
	case cleanupKindRetry:
		kind = cleanupKindRetry
	case cleanupKindRun:
	default:
		return cleanupModel{}, "", fmt.Errorf("step %q: unsupported x-cleanup %q", step.StepID, step.XCleanup)
	}

	info, ok := index[step.OperationID]
	if !ok {
		return cleanupModel{}, "", fmt.Errorf("step %q: no `atlas api` command for operationId %q", step.StepID, step.OperationID)
	}
	args, err := info.renderArgs(step)
	if err != nil {
		return cleanupModel{}, "", err
	}
	if timeout, watched := watches[step.StepID]; watched {
		args += `, "--watch", "--watchTimeout", ` + strconv.Quote(strconv.Itoa(timeout))
	}

	producer, err := latestProducer(step, steps)
	if err != nil {
		return cleanupModel{}, "", fmt.Errorf("step %q: %w", step.StepID, err)
	}

	return cleanupModel{Kind: kind, Args: args}, producer, nil
}

// latestProducer finds the last step whose outputs the given step references,
// so a cleanup registration lands after all of its inputs are assigned.
func latestProducer(step Step, steps []Step) (string, error) {
	refs := make(map[string]bool)
	step.refs(refs)

	producerIndex := -1
	producerID := ""
	for i, candidate := range steps {
		for ref := range refs {
			if strings.HasPrefix(ref, stepsPrefix) && strings.HasPrefix(ref, stepsPrefix+candidate.StepID+outputsInfix) && i > producerIndex {
				producerIndex = i
				producerID = candidate.StepID
			}
		}
	}
	if producerIndex == -1 {
		return "", errors.New("x-cleanup requires at least one reference to another step's outputs")
	}

	return producerID, nil
}

// buildInputs renders the top-of-function statements that bind and derive the
// workflow inputs, and reports which derived inputs are conditionally set.
func buildInputs(schema InputsSchema, refs map[string]bool) ([]string, map[string]bool, error) {
	names := sortedKeys(schema.Properties)
	for _, name := range names {
		if property := schema.Properties[name]; property.XBinding != "" && property.XDerived != nil {
			return nil, nil, fmt.Errorf("input %q declares both x-binding and x-derived", name)
		}
	}

	bound, derived := classifyInputs(names, schema.Properties)
	derivedFrom, templateInputs, err := collectDerivedUsage(derived)
	if err != nil {
		return nil, nil, err
	}

	used := func(name string) bool {
		return refs[inputsPrefix+name] || derivedFrom[name] || templateInputs[name]
	}

	statements := make([]string, 0)
	statements, err = appendBindingStatements(statements, names, bound, used)
	if err != nil {
		return nil, nil, err
	}

	optional := make(map[string]bool)
	for _, name := range names {
		property, ok := derived[name]
		if !ok || !used(name) {
			continue
		}
		derivedStmts, isOptional, err := derivedStatement(name, property, schema)
		if err != nil {
			return nil, nil, err
		}
		statements = append(statements, derivedStmts...)
		if isOptional {
			optional[name] = true
		}
	}

	return statements, optional, nil
}

func classifyInputs(names []string, properties map[string]InputProperty) (map[string]string, map[string]InputProperty) {
	bound := make(map[string]string)
	derived := make(map[string]InputProperty)
	for _, name := range names {
		property := properties[name]
		switch {
		case property.XBinding != "":
			bound[name] = property.XBinding
		case property.XDerived != nil:
			derived[name] = property
		}
	}

	return bound, derived
}

func collectDerivedUsage(derived map[string]InputProperty) (map[string]bool, map[string]bool, error) {
	derivedFrom := make(map[string]bool)
	templateInputs := make(map[string]bool)
	for _, property := range derived {
		if property.XDerived.From != "" {
			derivedFrom[property.XDerived.From] = true
		}
		if property.XDerived.Template == "" {
			continue
		}
		parts, err := parseTemplate(property.XDerived.Template)
		if err != nil {
			return nil, nil, err
		}
		for _, part := range parts {
			if part.expr != "" {
				templateInputs[part.expr] = true
			}
		}
	}

	return derivedFrom, templateInputs, nil
}

func appendBindingStatements(statements []string, names []string, bound map[string]string, used func(string) bool) ([]string, error) {
	var err error
	for _, name := range names {
		kind, ok := bound[name]
		if !ok || !used(name) {
			continue
		}
		statements, err = appendBindingStatement(statements, name, kind)
		if err != nil {
			return nil, err
		}
	}

	return statements, nil
}

func appendBindingStatement(statements []string, name, kind string) ([]string, error) {
	ident := sanitizeIdent(name)
	switch kind {
	case "profile-org":
		return append(statements, ident+" := apitest.OrgID(t)"), nil
	case "profile-service":
		return append(statements, ident+" := apitest.Service()"), nil
	case "random":
		return append(statements, ident+" := apitest.RandomSuffix(t)"), nil
	default:
		return nil, fmt.Errorf("input %q declares unsupported x-binding %q", name, kind)
	}
}

// derivedStatement renders a derived input. Inputs derived from an input with
// an enum are only considered always set when every enum value has a mapping;
// partial mappings leave the input unset for unmapped services.
func derivedStatement(name string, property InputProperty, schema InputsSchema) ([]string, bool, error) {
	ident := sanitizeIdent(name)
	derived := property.XDerived

	switch {
	case derived.Template != "":
		parts, err := parseTemplate(derived.Template)
		if err != nil {
			return nil, false, err
		}
		expr, err := templateExpr(parts)
		if err != nil {
			return nil, false, fmt.Errorf("input %q: %w", name, err)
		}
		return []string{ident + " := " + expr}, false, nil
	case derived.From != "":
		return derivedFromStatement(ident, derived, schema)
	default:
		return nil, false, fmt.Errorf("input %q declares an empty x-derived", name)
	}
}

func derivedFromStatement(ident string, derived *Derived, schema InputsSchema) ([]string, bool, error) {
	from, ok := schema.Properties[derived.From]
	if !ok {
		return nil, false, fmt.Errorf("input %q derives from unknown input %q", ident, derived.From)
	}
	fromIdent := sanitizeIdent(derived.From)

	complete := true
	for _, value := range from.Enum {
		if _, mapped := derived.Values[value]; !mapped {
			complete = false
			break
		}
	}

	statements := []string{ident + ` := ""`}
	for _, service := range sortedStringKeys(derived.Values) {
		statements = append(statements,
			`if `+fromIdent+` == `+strconv.Quote(service)+` {`,
			"\t"+ident+" = "+strconv.Quote(derived.Values[service]),
			`}`,
		)
	}

	return statements, !complete, nil
}

func templateExpr(parts []templatePart) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("empty template")
	}
	if len(parts) == 1 && parts[0].literal != "" && parts[0].expr == "" {
		return strconv.Quote(parts[0].literal), nil
	}

	rendered := make([]string, 0, len(parts))
	for _, part := range parts {
		switch {
		case part.expr != "":
			rendered = append(rendered, part.expr)
		case part.literal != "":
			rendered = append(rendered, strconv.Quote(part.literal))
		}
	}

	return strings.Join(rendered, " + "), nil
}

func buildStep(
	step Step,
	index map[string]commandInfo,
	watches map[string]int,
	optionalInputs map[string]bool,
	outputUsage map[string]bool,
) (stepModel, error) {
	info, ok := index[step.OperationID]
	if !ok {
		return stepModel{}, fmt.Errorf("step %q: no `atlas api` command for operationId %q", step.StepID, step.OperationID)
	}

	model := stepModel{ID: step.StepID}
	for _, line := range strings.Split(strings.TrimSpace(step.Description), "\n") {
		model.Comment = append(model.Comment, "// "+strings.TrimSpace(line))
	}

	args, err := info.renderArgs(step)
	if err != nil {
		return stepModel{}, err
	}
	if timeout, watched := watches[step.StepID]; watched {
		args += `, "--watch", "--watchTimeout", ` + strconv.Quote(strconv.Itoa(timeout))
	}
	model.ArgsLiteral = args

	if err := buildStepCriteria(&model, step); err != nil {
		return stepModel{}, err
	}

	if step.RequestBody != nil {
		bodyLines, stdinExpr, err := renderBody(step, optionalInputs)
		if err != nil {
			return stepModel{}, err
		}
		model.BodyLines = bodyLines
		model.StdinExpr = stdinExpr
	} else {
		model.StdinExpr = `""`
	}

	if err := buildStepOutputs(&model, step, outputUsage); err != nil {
		return stepModel{}, err
	}

	model.NeedsResponse = len(model.OutputStmts) > 0 || len(model.DerivedStmts) > 0 || len(model.Assertions) > 0

	return model, nil
}

func buildStepCriteria(model *stepModel, step Step) error {
	expectError := false
	assertions := make([]string, 0)
	for _, criterion := range step.SuccessCriteria {
		parsed, err := parseCriterion(criterion.Condition)
		if err != nil {
			return fmt.Errorf("step %q: %w", step.StepID, err)
		}
		switch parsed.kind {
		case criterionStatus:
			if parsed.status < 200 || parsed.status > 299 {
				expectError = true
			}
		case criterionNotNull:
			assertions = append(assertions, notNullStatement(step.StepID, parsed.pointer))
		case criterionEqual:
			assertions = append(assertions, equalStatement(step.StepID, parsed.pointer, parsed.expected))
		}
	}
	model.ExpectError = expectError
	model.Assertions = assertions

	return nil
}

func buildStepOutputs(model *stepModel, step Step, outputUsage map[string]bool) error {
	for _, name := range sortedKeys(step.Outputs) {
		if !outputUsage[stepOutputRef(step.StepID, name)] {
			continue
		}
		pointer, err := responseBodyPointer(step.Outputs[name])
		if err != nil {
			return fmt.Errorf("step %q: output %q: %w", step.StepID, name, err)
		}
		model.OutputStmts = append(model.OutputStmts,
			sanitizeIdent(step.StepID+"_"+name)+` = apitest.JSONString(`+respVar(step.StepID)+`, `+strconv.Quote(pointer)+`)`)
	}

	for _, name := range sortedKeys(step.XDerivedOutputs) {
		if !outputUsage[stepOutputRef(step.StepID, name)] {
			continue
		}
		statements, err := derivedOutputStatements(step, name, step.XDerivedOutputs[name])
		if err != nil {
			return fmt.Errorf("step %q: derived output %q: %w", step.StepID, name, err)
		}
		model.DerivedStmts = append(model.DerivedStmts, statements...)
	}

	return nil
}

func renderBody(step Step, optionalInputs map[string]bool) ([]string, string, error) {
	payload, ok := step.RequestBody.Payload.(map[string]any)
	if !ok {
		return nil, "", fmt.Errorf("step %q: request body payload must be a mapping", step.StepID)
	}

	bodyVar := sanitizeIdent(step.StepID) + "Body"
	conditional := make(map[string]string)
	entries := make([]string, 0, len(payload))
	for _, key := range sortedKeys(payload) {
		if err := rejectNestedConditional(childrenOf(payload[key]), optionalInputs, "/"+key); err != nil {
			return nil, "", fmt.Errorf("step %q: %w", step.StepID, err)
		}
		rendered, err := renderGoValue(payload[key])
		if err != nil {
			return nil, "", fmt.Errorf("step %q: payload key %q: %w", step.StepID, key, err)
		}
		if optionalInputs[rendered] {
			conditional[key] = rendered
			continue
		}
		entries = append(entries, strconv.Quote(key)+": "+rendered)
	}

	lines := []string{bodyVar + ` := map[string]any{` + strings.Join(entries, ", ") + `}`}
	for _, key := range sortedKeys(conditional) {
		lines = append(lines,
			`if `+conditional[key]+` != "" {`,
			"\t"+bodyVar+"["+strconv.Quote(key)+"] = "+conditional[key],
			`}`,
		)
	}

	return lines, "apitest.Marshal(t, " + bodyVar + ")", nil
}

// childrenOf returns the child values of a payload container, so that
// conditionally derived inputs are only rejected below the top level.
func childrenOf(value any) map[string]any {
	switch typed := value.(type) {
	case map[string]any:
		return typed
	case []any:
		children := make(map[string]any, len(typed))
		for i, item := range typed {
			children[strconv.Itoa(i)] = item
		}
		return children
	default:
		return nil
	}
}

// rejectNestedConditional rejects conditionally derived inputs referenced
// below the top level of a payload, where they cannot be conditionally omitted.
func rejectNestedConditional(children map[string]any, optionalInputs map[string]bool, path string) error {
	for _, key := range sortedKeys(children) {
		child := children[key]
		childPath := path + "/" + key
		if typed, ok := child.(string); ok {
			if ident, resolved := resolveExpr(typed); resolved && optionalInputs[ident] {
				return fmt.Errorf("conditional input %q at %s must stay at the top level of the payload", ident, childPath)
			}
		}
		if err := rejectNestedConditional(childrenOf(child), optionalInputs, childPath); err != nil {
			return err
		}
	}

	return nil
}

func derivedOutputStatements(step Step, name string, derived Derived) ([]string, error) {
	ident := sanitizeIdent(step.StepID + "_" + name)
	parts, err := parseTemplate(derived.Template)
	if err != nil {
		return nil, err
	}
	if len(parts) != 1 || parts[0].expr == "" {
		return nil, fmt.Errorf("template %q must contain a single expression", derived.Template)
	}
	fromVar := parts[0].expr

	switch derived.Transform {
	case "":
		return []string{ident + ` = ` + fromVar}, nil
	case transformPreviousMajorVersion:
		local := sanitizeIdent(name)
		return []string{
			local + `, err := apitest.PreviousMajorVersion(` + fromVar + `)`,
			`require.NoError(t, err)`,
			ident + ` = ` + local,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported transform %q", derived.Transform)
	}
}

func responseBodyPointer(output string) (string, error) {
	pointer, found := strings.CutPrefix(output, responsePointer)
	if !found {
		return "", fmt.Errorf("unsupported output expression %q", output)
	}

	return pointer, nil
}

func notNullStatement(stepID, pointer string) string {
	return `require.NotNil(t, apitest.JSONValue(` + respVar(stepID) + `, ` + strconv.Quote(pointer) + `))`
}

func equalStatement(stepID, pointer, expected string) string {
	return `require.Equal(t, ` + expected + `, apitest.JSONString(` + respVar(stepID) + `, ` + strconv.Quote(pointer) + `))`
}

func respVar(stepID string) string {
	return sanitizeIdent(stepID) + "Resp"
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func sortedStringKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
