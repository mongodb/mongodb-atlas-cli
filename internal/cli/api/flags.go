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

package api

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/spf13/pflag"
)

func parameterToPFlag(parameter api.Parameter) *pflag.Flag {
	name := parameter.Name
	shortDescription := flagDescription(parameter.Description)

	if parameter.Type.IsArray {
		switch parameter.Type.Type {
		case api.TypeString:
			return createFlag(func(f *pflag.FlagSet) {
				f.StringArray(name, make([]string, 0), shortDescription)
			})
		case api.TypeInt:
			return createFlag(func(f *pflag.FlagSet) {
				f.Int32Slice(name, make([]int32, 0), shortDescription)
			})
		case api.TypeBool:
			return createFlag(func(f *pflag.FlagSet) {
				f.BoolSlice(name, make([]bool, 0), shortDescription)
			})
		}
	} else {
		switch parameter.Type.Type {
		case api.TypeString:
			return createFlag(func(f *pflag.FlagSet) {
				f.String(name, "", shortDescription)
			})
		case api.TypeInt:
			return createFlag(func(f *pflag.FlagSet) {
				f.Int(name, 0, shortDescription)
			})
		case api.TypeBool:
			return createFlag(func(f *pflag.FlagSet) {
				f.Bool(name, false, shortDescription)
			})
		}
	}

	// should never happen, can only happen if someone adds a new api.ParameterConcreteType
	panic(fmt.Sprintf("unsupported parameter type: %s", parameter.Type.Type))
}

func createFlag(factory func(f *pflag.FlagSet)) *pflag.Flag {
	flagSet := pflag.NewFlagSet("temp", pflag.ContinueOnError)
	factory(flagSet)
	var output *pflag.Flag
	flagSet.VisitAll(func(f *pflag.Flag) {
		output = f
	})
	return output
}

func flagDescription(description string) string {
	shortDescription, _ := splitShortAndLongDescription(description)
	if len(shortDescription) > 0 {
		shortDescription = strings.ToLower(shortDescription[:1]) + shortDescription[1:]
	}
	shortDescription = strings.TrimSuffix(shortDescription, ".")
	return shortDescription
}

type FlagValueProvider interface {
	ValueForFlag(flagName string) (*string, error)
}
