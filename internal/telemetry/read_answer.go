// Copyright 2022 MongoDB Inc
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

package telemetry

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrNotMapOrStruct = errors.New("invalid response: expected a struct or map[string]interface{}")
	ErrFieldNotFound  = errors.New("field not found")
)

type fieldValue struct {
	field reflect.StructField
	value reflect.Value
}

func listFields(target reflect.Value) []fieldValue {
	targetType := target.Type()
	numFields := targetType.NumField()
	ret := make([]fieldValue, numFields)
	for i := 0; i < numFields; i++ {
		ret[i] = fieldValue{
			field: targetType.Field(i),
			value: target.Field(i),
		}
	}
	return ret
}

func findAnswerInStruct(target reflect.Value, name string) (interface{}, error) { // based on https://pkg.go.dev/github.com/AlecAivazis/survey/v2/core#WriteAnswer
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}
	if target.Kind() != reflect.Struct {
		return nil, ErrNotMapOrStruct
	}

	toVisit := listFields(target)
	for len(toVisit) > 0 {
		top := toVisit[0]
		toVisit = toVisit[1:]

		if strings.EqualFold(top.field.Name, name) || top.field.Tag.Get("survey") == name {
			return top.value.Interface(), nil
		}

		fieldType := top.field.Type
		value := top.value
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
			value = value.Elem()
		}
		if fieldType.Kind() == reflect.Struct {
			toVisit = append(toVisit, listFields(value)...)
		}
	}

	return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, name)
}

func readAnswer(response interface{}, name string) (interface{}, error) {
	v, ok := response.(map[string]interface{})
	if ok {
		ret, ok := v[name]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, name)
		}
		return ret, nil
	}
	target := reflect.ValueOf(response)
	return findAnswerInStruct(target, name)
}
