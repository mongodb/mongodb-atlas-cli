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
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/spf13/cobra"
)

var (
	errUnsupportedType   = errors.New("unsupported parameter type")
	errFlagAlreadyExists = errors.New("parameter already exists")
)

func addFlag(cmd *cobra.Command, parameter api.Parameter) error {
	name := parameter.Name

	if cmd.Flag(name) != nil {
		// this should never happen, the api command generation tool should cover this
		return fmt.Errorf("%w: %s", errFlagAlreadyExists, name)
	}

	shortDescription := flagDescription(parameter.Description)

	if parameter.Type.IsArray {
		switch parameter.Type.Type {
		case api.TypeString:
			cmd.Flags().StringArrayP(name, parameter.Short, nil, shortDescription)
		case api.TypeInt:
			cmd.Flags().Int32SliceP(name, parameter.Short, nil, shortDescription)
		case api.TypeBool:
			cmd.Flags().BoolSliceP(name, parameter.Short, nil, shortDescription)
		default:
			return fmt.Errorf("%w: %s", errUnsupportedType, parameter.Type.Type)
		}
	} else {
		switch parameter.Type.Type {
		case api.TypeString:
			cmd.Flags().StringP(name, parameter.Short, "", shortDescription)
		case api.TypeInt:
			cmd.Flags().IntP(name, parameter.Short, 0, shortDescription)
		case api.TypeBool:
			cmd.Flags().BoolP(name, parameter.Short, false, shortDescription)
		default:
			return fmt.Errorf("%w: %s", errUnsupportedType, parameter.Type.Type)
		}
	}

	if parameter.Required {
		if err := cmd.MarkFlagRequired(name); err != nil {
			return err
		}
	}

	return nil
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
