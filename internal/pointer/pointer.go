// Copyright 2023 MongoDB Inc
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

package pointer

import (
	"time"

	customTime "github.com/mongodb/mongodb-atlas-cli/internal/time"
	"golang.org/x/exp/constraints"
)

func GetOrDefault[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

func Get[T any](val T) *T {
	return &val
}

func GetStringPointerIfNotEmpty(input string) *string {
	if input != "" {
		return &input
	}
	return nil
}

func GetArrayPointerIfNotEmpty(input []string) *[]string {
	if len(input) > 0 {
		return &input
	}
	return nil
}

func StringToTimePointer(value string) *time.Time {
	var result *time.Time
	if completedAfter, err := customTime.ParseTimestamp(value); err == nil {
		result = &completedAfter
	}
	return result
}

func GetNonZeroValue[T constraints.Integer](val T) *T {
	if val > T(0) {
		return Get(val)
	}

	return nil
}
