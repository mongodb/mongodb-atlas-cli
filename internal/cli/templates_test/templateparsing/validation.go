package templateparsing

import (
	"errors"
	"go/types"
	"strings"
	"unicode"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/astparsing"
	"golang.org/x/tools/go/packages"
)

func (c *TemplateCallTree) Validate(pkg *packages.Package, typeInfo types.Type) (bool, error) {
	switch typeInfo := typeInfo.(type) {
	case (*types.Basic):
		return c.validateBasic(typeInfo)
	case (*types.Map):
		return c.validateMap(pkg, typeInfo)
	case (*types.Named):
		return c.validateStruct(pkg, typeInfo)
	case (*types.Pointer):
		return c.Validate(pkg, typeInfo.Elem())
	case (*types.Slice):
		return c.validateSlice(pkg, typeInfo)

	default:
		return false, errors.New("unsupported type")
	}
}

func (c *TemplateCallTree) validateBasic(sliceInfo *types.Basic) (bool, error) {
	return c.listType == nil || c.structType == nil || (c.structType != nil && len(c.structType.fields) == 0), nil
}

func (c *TemplateCallTree) validateMap(pkg *packages.Package, sliceInfo *types.Map) (bool, error) {
	// We expected a list, got a struct
	if c.structType != nil {
		return false, nil
	}

	// We don't need to drill any deeper, which means our list is valid
	if c.listType == nil {
		return true, nil
	}

	// Validate both keys and values recursively
	if valid, err := c.listType.Validate(pkg, sliceInfo.Key()); !valid || err != nil {
		return valid, err
	}

	return c.listType.Validate(pkg, sliceInfo.Elem())
}

func (c *TemplateCallTree) validateSlice(pkg *packages.Package, sliceInfo *types.Slice) (bool, error) {
	// We expected a list, got a struct
	if c.structType != nil {
		return false, nil
	}

	// We don't need to drill any deeper, which means our list is valid
	if c.listType == nil {
		return true, nil
	}

	// Validate recursively
	return c.listType.Validate(pkg, sliceInfo.Elem())
}

func (c *TemplateCallTree) validateStruct(pkg *packages.Package, namedTypeInfo *types.Named) (bool, error) {
	// We expected a struct, got a list
	if c.listType != nil {
		return false, nil
	}

	// We don't need to drill any deeper, which means our struct is valid
	if c.structType == nil || len(c.structType.fields) == 0 {
		return true, nil
	}

	// Get template fields on struct
	structStructure, err := getTemplateFields(pkg, namedTypeInfo)
	if err != nil {
		return false, err
	}

	// Check if all fields are valid
	for key, field := range c.structType.fields {
		// Templates ignore cases
		key := strings.ToLower(key)

		// Get the field on the struct
		structField := structStructure[key]

		// If we don't find the field the tempalate and struct don't match
		if structField == nil {
			return false, nil
		}

		// In case we find it, validate the field
		valid, err := field.Validate(pkg, structField)
		if !valid || err != nil {
			return valid, err
		}
	}

	return true, nil
}

func getTemplateFields(pkg *packages.Package, namedTypeInfo *types.Named) (map[string]types.Type, error) {
	structInfo, ok := namedTypeInfo.Underlying().(*types.Struct)
	if !ok {
		return nil, errors.New("expecting underlying type to be a struct")
	}

	structStructure := make(map[string]types.Type)

	// Add fields and embedded fields
	for i := 0; i < structInfo.NumFields(); i++ {
		field := structInfo.Field(i)
		fieldName := field.Name()

		if field.Embedded() {
			ty, err := astparsing.GetUnderlyingStruct(field.Type())
			if err != nil {
				return nil, err
			}

			subStructure, err := getTemplateFields(pkg, ty.NamedStruct)
			if err != nil {
				return nil, err
			}

			for k, v := range subStructure {
				structStructure[k] = v
			}
		} else {
			// Only check public fields
			if firstRuneIsCapitalized(fieldName) {
				lowerCaseFieldName := strings.ToLower(fieldName)
				structStructure[lowerCaseFieldName] = field.Type()
			}
		}
	}

	// Add method names
	for i := 0; i < namedTypeInfo.NumMethods(); i++ {
		method := namedTypeInfo.Method(i)

		// Only check public fields
		methodName := method.Name()
		if !firstRuneIsCapitalized(methodName) {
			continue
		}

		methodAsField, err := methodToTemplateFieldType(method)
		if err != nil {
			continue
		}

		lowerCaseFieldName := strings.ToLower(methodName)
		structStructure[lowerCaseFieldName] = methodAsField
	}

	return structStructure, nil
}

func firstRuneIsCapitalized(str string) bool {
	for _, r := range str {
		return unicode.IsUpper(r)
	}

	return true
}

func methodToTemplateFieldType(method *types.Func) (types.Type, error) {
	signature, ok := method.Type().(*types.Signature)
	if !ok {
		return nil, errors.New("expected method to be of the type *types.Signature")
	}

	if signature.Params().Len() > 0 {
		return nil, errors.New("only parameterless methods are suppored")
	}

	res := signature.Results()
	if res.Len() != 1 {
		return nil, errors.New("only methods returning exactly one result")
	}

	return res.At(0).Type(), nil
}
