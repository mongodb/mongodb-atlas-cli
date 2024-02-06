package templateparsing

import (
	"errors"
	"fmt"
	"go/types"
	"strings"
	"unicode"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/templates_test/astparsing"
	"golang.org/x/tools/go/packages"
)

type ValidationResult struct {
	isValid       bool
	errorMessages []string
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		isValid:       true,
		errorMessages: make([]string, 0),
	}
}

func (r *ValidationResult) AddErrorMessage(breadCrumb string, message string) {
	r.errorMessages = append(r.errorMessages, fmt.Sprintf("path: '%v', error: %v", breadCrumb, message))
	r.isValid = false
}

func (r *ValidationResult) IsValid() bool {
	return r.isValid
}

func (r *ValidationResult) ErrorMessages() []string {
	return r.errorMessages
}

func (c *TemplateCallTree) Validate(pkg *packages.Package, typeInfo types.Type) (*ValidationResult, error) {
	result := NewValidationResult()
	if err := c.validateInner(pkg, result, "", typeInfo); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (c *TemplateCallTree) validateInner(pkg *packages.Package, result *ValidationResult, breadCrumb string, typeInfo types.Type) error {
	switch typeInfo := typeInfo.(type) {
	case (*types.Basic):
		return c.validateBasic(result, breadCrumb, typeInfo)
	case (*types.Map):
		return c.validateMap(pkg, result, breadCrumb, typeInfo)
	case (*types.Named):
		return c.validateStruct(pkg, result, breadCrumb, typeInfo)
	case (*types.Pointer):
		return c.validateInner(pkg, result, breadCrumb, typeInfo.Elem())
	case (*types.Slice):
		return c.validateSlice(pkg, result, breadCrumb, typeInfo)

	default:
		return errors.New("unsupported type")
	}
}

func (c *TemplateCallTree) validateBasic(result *ValidationResult, breadCrumb string, sliceInfo *types.Basic) error {
	if c.listType != nil || (c.structType != nil && len(c.structType.fields) != 0) {
		result.AddErrorMessage(breadCrumb, "expecting a complext type, but received a basic type")
	}

	return nil
}

func (c *TemplateCallTree) validateMap(pkg *packages.Package, result *ValidationResult, breadCrumb string, sliceInfo *types.Map) error {
	// We expected a struct, did not get a struct
	if c.structType != nil {
		result.AddErrorMessage(breadCrumb, "expecting a struct, did not get a struct")
		return nil
	}

	// We don't need to drill any deeper, which means our list is valid
	if c.listType == nil {
		return nil
	}

	// Validate both keys and values recursively
	if err := c.listType.validateInner(pkg, result, breadCrumb+"[]", sliceInfo.Key()); err != nil {
		return err
	}

	if err := c.listType.validateInner(pkg, result, breadCrumb+"[]", sliceInfo.Elem()); err != nil {
		return err
	}

	return nil
}

func (c *TemplateCallTree) validateSlice(pkg *packages.Package, result *ValidationResult, breadCrumb string, sliceInfo *types.Slice) error {
	// We expected a list, got a struct
	if c.structType != nil {
		result.AddErrorMessage(breadCrumb, "expecting a list, got a struct instead")
		return nil
	}

	// We don't need to drill any deeper, which means our list is valid
	if c.listType == nil {
		return nil
	}

	// Validate recursively
	return c.listType.validateInner(pkg, result, breadCrumb+"[]", sliceInfo.Elem())
}

func (c *TemplateCallTree) validateStruct(pkg *packages.Package, result *ValidationResult, breadCrumb string, namedTypeInfo *types.Named) error {
	// We expected a struct, got a list
	if c.listType != nil {
		result.AddErrorMessage(breadCrumb, "expecting a struct, got a list instead")
		return nil
	}

	// We don't need to drill any deeper, which means our struct is valid
	if c.structType == nil || len(c.structType.fields) == 0 {
		return nil
	}

	// Get template fields on struct
	structStructure, err := getTemplateFields(pkg, namedTypeInfo)
	if err != nil {
		return err
	}

	// Check if all fields are valid
	for key, field := range c.structType.fields {
		// Templates ignore cases
		key := strings.ToLower(key)

		// Get the field on the struct
		structField := structStructure[key]

		// If we don't find the field the tempalate and struct don't match
		if structField == nil {
			result.AddErrorMessage(breadCrumb, fmt.Sprintf("field '%v' is missing", key))
			return nil
		}

		// In case we find it, validate the field
		err := field.validateInner(pkg, result, breadCrumb+"."+key, structField)
		if err != nil {
			return err
		}
	}

	return nil
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
