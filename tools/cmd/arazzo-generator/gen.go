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
	"fmt"
	"sort"
	"strconv"
	"strings"

	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

const (
	inputsPrefix    = "$inputs."
	stepsPrefix     = "$steps."
	outputsInfix    = ".outputs."
	responsePointer = "$response.body#/"
	statusExpr      = "$statusCode"

	retryAction = "retry"
)

// testModel is the fully resolved form of a workflow, ready to render.
type testModel struct {
	Package      string
	TestName     string
	Inputs       []string
	VarDecls     []string
	Steps        []stepModel
	NeedsRequire bool
}

type stepModel struct {
	ID            string
	Comment       []string
	BodyLines     []string
	ArgsLiteral   string
	StdinExpr     string
	ExpectError   bool
	NeedsResponse bool
	Assertions    []string
	OutputStmts   []string
	DerivedStmts  []string
	Cleanups      []cleanupModel
}

type cleanupModel struct {
	Kind string
	Args string
}

type criterionKind string

const (
	criterionStatus  criterionKind = "status"
	criterionNotNull criterionKind = "notnull"
	criterionEqual   criterionKind = "equal"
)

type criterion struct {
	kind        criterionKind
	status      int
	pointer     string
	expected    string
	expectedRaw string
}

type templatePart struct {
	literal string
	expr    string
}

// resolveExpr converts a runtime expression into the Go identifier that holds
// its resolved value.
func resolveExpr(expr string) (string, bool) {
	switch {
	case strings.HasPrefix(expr, inputsPrefix):
		return sanitizeIdent(expr[len(inputsPrefix):]), true
	case strings.HasPrefix(expr, stepsPrefix):
		rest := expr[len(stepsPrefix):]
		stepID, output, found := strings.Cut(rest, outputsInfix)
		if !found {
			return "", false
		}
		return sanitizeIdent(stepID + "_" + output), true
	}

	return "", false
}

func sanitizeIdent(name string) string {
	var builder strings.Builder
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '_':
			builder.WriteRune(r)
		default:
			builder.WriteRune('_')
		}
	}

	return builder.String()
}

// parseTemplate splits a derivation template into literal and expression parts.
func parseTemplate(template string) ([]templatePart, error) {
	var parts []templatePart
	for {
		open := strings.Index(template, "{")
		if open == -1 {
			if template != "" {
				parts = append(parts, templatePart{literal: template})
			}
			return parts, nil
		}
		end := strings.Index(template[open:], "}")
		if end == -1 {
			return nil, fmt.Errorf("unbalanced expression in template %q", template)
		}
		end += open
		if template[:open] != "" {
			parts = append(parts, templatePart{literal: template[:open]})
		}
		expr := template[open+1 : end]
		resolved, ok := resolveExpr(expr)
		if !ok {
			return nil, fmt.Errorf("unsupported expression %q in template", expr)
		}
		parts = append(parts, templatePart{expr: resolved})
		template = template[end+1:]
	}
}

func parseCriterion(condition string) (criterion, error) {
	condition = strings.TrimSpace(condition)

	if strings.HasPrefix(condition, statusExpr) {
		if !strings.HasPrefix(condition[len(statusExpr):], " == ") {
			return criterion{}, fmt.Errorf("unsupported status criterion %q", condition)
		}
		status, err := strconv.Atoi(strings.TrimSpace(condition[len(statusExpr)+len(" == "):]))
		if err != nil {
			return criterion{}, fmt.Errorf("unsupported status criterion %q", condition)
		}
		return criterion{kind: criterionStatus, status: status}, nil
	}

	if !strings.HasPrefix(condition, responsePointer) {
		return criterion{}, fmt.Errorf("unsupported criterion %q", condition)
	}

	rest := condition[len(responsePointer):]
	equalAt := strings.Index(rest, " == ")
	notEqualAt := strings.Index(rest, " != ")
	switch {
	case notEqualAt != -1 && (equalAt == -1 || notEqualAt < equalAt):
		pointer := rest[:notEqualAt]
		rhs := strings.TrimSpace(rest[notEqualAt+len(" != "):])
		if rhs != "null" {
			return criterion{}, fmt.Errorf("only comparisons against null are supported for !=, got %q", condition)
		}
		return criterion{kind: criterionNotNull, pointer: pointer}, nil
	case equalAt != -1:
		pointer := rest[:equalAt]
		rhs := strings.TrimSpace(rest[equalAt+len(" == "):])
		expected, err := expectedExpr(rhs)
		if err != nil {
			return criterion{}, err
		}
		return criterion{kind: criterionEqual, pointer: pointer, expected: expected, expectedRaw: rhs}, nil
	}

	return criterion{}, fmt.Errorf("unsupported criterion %q", condition)
}

// expectedExpr renders the right-hand side of an equality criterion as the Go
// expression compared against apitest.JSONString.
func expectedExpr(rhs string) (string, error) {
	switch {
	case strings.HasPrefix(rhs, "$"):
		resolved, ok := resolveExpr(rhs)
		if !ok {
			return "", fmt.Errorf("unsupported expression %q in criterion", rhs)
		}
		return resolved, nil
	case strings.HasPrefix(rhs, "'"), strings.HasPrefix(rhs, `"`):
		if len(rhs) < 2 || rhs[len(rhs)-1] != rhs[0] {
			return "", fmt.Errorf("unbalanced string literal %q in criterion", rhs)
		}
		return strconv.Quote(rhs[1 : len(rhs)-1]), nil
	case rhs == "true", rhs == "false":
		return strconv.Quote(rhs), nil
	default:
		if _, err := strconv.ParseFloat(rhs, 64); err != nil {
			return "", fmt.Errorf("unsupported literal %q in criterion", rhs)
		}
		return strconv.Quote(rhs), nil
	}
}

// renderGoValue converts a payload value into a Go expression. Expressions are
// resolved to the identifiers that hold them; everything else is rendered
// literally.
func renderGoValue(value any) (string, error) {
	switch typed := value.(type) {
	case string:
		if strings.HasPrefix(typed, "$") {
			resolved, ok := resolveExpr(typed)
			if !ok {
				return "", fmt.Errorf("unsupported expression %q in payload", typed)
			}
			return resolved, nil
		}
		return strconv.Quote(typed), nil
	case bool:
		return strconv.FormatBool(typed), nil
	case int:
		return strconv.Itoa(typed), nil
	case int64:
		return strconv.FormatInt(typed, 10), nil
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64), nil
	case []any:
		items := make([]string, 0, len(typed))
		for _, item := range typed {
			rendered, err := renderGoValue(item)
			if err != nil {
				return "", err
			}
			items = append(items, rendered)
		}
		return "[]any{" + strings.Join(items, ", ") + "}", nil
	case map[string]any:
		keys := make([]string, 0, len(typed))
		for key := range typed {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		entries := make([]string, 0, len(keys))
		for _, key := range keys {
			rendered, err := renderGoValue(typed[key])
			if err != nil {
				return "", err
			}
			entries = append(entries, strconv.Quote(key)+": "+rendered)
		}
		return "map[string]any{" + strings.Join(entries, ", ") + "}", nil
	default:
		return "", fmt.Errorf("unsupported payload value %v (%T)", value, value)
	}
}

// extractRefs walks a value collecting every runtime expression it contains.
func extractRefs(value any, refs map[string]bool) {
	switch typed := value.(type) {
	case string:
		if strings.HasPrefix(typed, "$") {
			refs[typed] = true
		}
	case []any:
		for _, item := range typed {
			extractRefs(item, refs)
		}
	case map[string]any:
		for _, item := range typed {
			extractRefs(item, refs)
		}
	}
}

func (s Step) refs(refs map[string]bool) {
	for _, parameter := range s.Parameters {
		extractRefs(parameter.Value, refs)
	}
	if s.RequestBody != nil {
		extractRefs(s.RequestBody.Payload, refs)
	}
	for _, derived := range s.XDerivedOutputs {
		for _, part := range strings.Split(derived.Template, "{") {
			if expr, found := strings.CutSuffix(part, "}"); found {
				refs[strings.TrimSpace(expr)] = true
			}
		}
	}
	for _, criterion := range s.SuccessCriteria {
		parsed, err := parseCriterion(criterion.Condition)
		if err != nil {
			continue
		}
		if parsed.kind == criterionEqual && strings.HasPrefix(parsed.expectedRaw, "$") {
			refs[parsed.expectedRaw] = true
		}
	}
}

// mergeWatchSteps folds polling steps into a --watch flag on the preceding
// step when the corresponding `atlas api` command defines a watcher for them.
// It returns the merged steps and the watch timeout in seconds per step ID.
func mergeWatchSteps(steps []Step, index map[string]commandInfo) ([]Step, map[string]int, error) {
	merged := make([]Step, 0, len(steps))
	watches := make(map[string]int)
	for _, step := range steps {
		if len(merged) > 0 && step.hasRetryOnFailure() && step.RequestBody == nil {
			previous := merged[len(merged)-1]
			previousInfo, ok := index[previous.OperationID]
			if ok && previousInfo.hasWatcher && previousInfo.watcherGetOperationID == step.OperationID {
				watches[previous.StepID] = watchTimeout(step)
				continue
			}
			return nil, nil, fmt.Errorf("step %q retries %q but the previous step does not watch it", step.StepID, step.OperationID)
		}
		merged = append(merged, step)
	}

	return merged, watches, nil
}

func watchTimeout(step Step) int {
	for _, action := range step.OnFailure {
		if action.Type != retryAction {
			continue
		}
		retryAfter := action.RetryAfter
		retryLimit := action.RetryLimit
		if retryAfter <= 0 {
			retryAfter = 1
		}
		if retryLimit <= 0 {
			retryLimit = 1
		}
		return retryAfter * retryLimit
	}

	return 0
}

func bodyVersion(contentType string) (string, error) {
	version, found := strings.CutPrefix(contentType, "application/vnd.atlas.")
	if !found {
		return "", fmt.Errorf("unsupported content type %q", contentType)
	}
	version, found = strings.CutSuffix(version, "+json")
	if !found {
		return "", fmt.Errorf("unsupported content type %q", contentType)
	}

	return version, nil
}

func (c commandInfo) renderArgs(step Step) (string, error) {
	args := []string{strconv.Quote(c.group), strconv.Quote(c.name)}

	for _, parameter := range step.Parameters {
		info, err := c.parameter(parameter.Name)
		if err != nil {
			return "", fmt.Errorf("step %q: %w", step.StepID, err)
		}

		switch {
		case info.Type.IsArray:
			items, ok := parameter.Value.([]any)
			if !ok {
				return "", fmt.Errorf("step %q: parameter %q expects a list", step.StepID, parameter.Name)
			}
			for _, item := range items {
				rendered, err := renderFlagValue(item)
				if err != nil {
					return "", fmt.Errorf("step %q: parameter %q: %w", step.StepID, parameter.Name, err)
				}
				args = append(args, strconv.Quote("--"+parameter.Name), rendered)
			}
		case info.Type.Type == shared_api.TypeBool:
			value, ok := parameter.Value.(bool)
			if !ok {
				return "", fmt.Errorf("step %q: parameter %q expects a boolean", step.StepID, parameter.Name)
			}
			if value {
				args = append(args, strconv.Quote("--"+parameter.Name))
			}
		default:
			rendered, err := renderFlagValue(parameter.Value)
			if err != nil {
				return "", fmt.Errorf("step %q: parameter %q: %w", step.StepID, parameter.Name, err)
			}
			args = append(args, strconv.Quote("--"+parameter.Name), rendered)
		}
	}

	if step.RequestBody != nil {
		version, err := bodyVersion(step.RequestBody.ContentType)
		if err != nil {
			return "", fmt.Errorf("step %q: %w", step.StepID, err)
		}
		if !c.supportsVersion(version) {
			return "", fmt.Errorf("step %q: version %q is not supported by %q", step.StepID, version, step.OperationID)
		}
		args = append(args, strconv.Quote("--version"), strconv.Quote(version))
	}

	return strings.Join(args, ", "), nil
}

func renderFlagValue(value any) (string, error) {
	switch typed := value.(type) {
	case string:
		if strings.HasPrefix(typed, "$") {
			resolved, ok := resolveExpr(typed)
			if !ok {
				return "", fmt.Errorf("unsupported expression %q", typed)
			}
			return resolved, nil
		}
		return strconv.Quote(typed), nil
	case bool:
		return strconv.Quote(strconv.FormatBool(typed)), nil
	case int:
		return strconv.Quote(strconv.Itoa(typed)), nil
	case int64:
		return strconv.Quote(strconv.FormatInt(typed, 10)), nil
	case float64:
		return strconv.Quote(strconv.FormatFloat(typed, 'f', -1, 64)), nil
	default:
		return "", fmt.Errorf("unsupported flag value %v (%T)", value, value)
	}
}
