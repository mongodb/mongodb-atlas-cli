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

package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

var (
	contentTypeHeaderRegex = regexp.MustCompile(`^application/vnd\.atlas\.(?<version>[^+]+)\+(?P<contentType>[\w]+)$`)
)

func specToCommands(now time.Time, spec *openapi3.T) (api.GroupedAndSortedCommands, error) {
	groups := make(map[string]*api.Group, 0)

	for path, item := range spec.Paths.Map() {
		for verb, operation := range item.Operations() {
			command, err := operationToCommand(now, path, verb, operation)
			if err != nil {
				return nil, fmt.Errorf("failed to convert operation to command: %w", err)
			}
			if command == nil {
				continue
			}

			if len(operation.Tags) != 1 {
				return nil, fmt.Errorf("expect every operation to have exactly 1 tag, got: %v", len(operation.Tags))
			}

			tag := operation.Tags[0]
			if _, ok := groups[tag]; !ok {
				group, err := groupForTag(spec, tag)
				if err != nil {
					return nil, err
				}

				groups[tag] = group
			}

			groups[tag].Commands = append(groups[tag].Commands, *command)
		}
	}

	// Validate that the defined watchers:
	// - are pointing to existing commands+version combos
	// - are using parameters correctly
	if err := validateAllWatchers(groups); err != nil {
		return nil, err
	}

	// Sort commands inside of groups
	sortedGroups := make([]api.Group, 0, len(groups))
	for _, group := range groups {
		sort.Slice(group.Commands, func(i, j int) bool {
			return group.Commands[i].OperationID < group.Commands[j].OperationID
		})

		sortedGroups = append(sortedGroups, *group)
	}

	// Sort groups
	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Name < sortedGroups[j].Name
	})

	return sortedGroups, nil
}

func extractSunsetDate(extensions map[string]any) *time.Time {
	if sSunset, ok := extensions["x-sunset"].(string); ok && sSunset != "" {
		if sunset, err := time.Parse("2006-01-02", sSunset); err == nil {
			return &sunset
		}
	}

	return nil
}

type operationExtensions struct {
	skip             bool
	operationID      string
	shortOperationID string
	operationAliases []string
}

func extractExtensionsFromOperation(operation *openapi3.Operation) operationExtensions {
	ext := operationExtensions{
		skip:             false,
		operationID:      operation.OperationID,
		shortOperationID: "",
		operationAliases: []string{},
	}

	if shortOperationID, ok := operation.Extensions["x-xgen-operation-id-override"].(string); ok && shortOperationID != "" {
		ext.shortOperationID = shortOperationID
	}

	if extensions, okExtensions := operation.Extensions["x-xgen-atlascli"].(map[string]any); okExtensions && extensions != nil {
		if extSkip, okSkip := extensions["skip"].(bool); okSkip && extSkip {
			ext.skip = extSkip
		}

		if extAliases, okExtAliases := extensions["command-aliases"].([]any); okExtAliases && extAliases != nil {
			for _, alias := range extAliases {
				if sAlias, ok := alias.(string); ok && sAlias != "" {
					ext.operationAliases = append(ext.operationAliases, sAlias)
				}
			}
		}

		// OperationID override for x-xgen-atlascli. This takes priority over x-xgen-operation-id-override.
		if overrides := extractOverrides(operation.Extensions); overrides != nil {
			if overriddenOperationID, ok := overrides["operationId"].(string); ok && overriddenOperationID != "" {
				ext.operationID = overriddenOperationID
				ext.shortOperationID = ""
			}
		}
	}

	return ext
}

func operationToCommand(now time.Time, path, verb string, operation *openapi3.Operation) (*api.Command, error) {
	extensions := extractExtensionsFromOperation(operation)
	if extensions.skip {
		return nil, nil
	}

	operationID := extensions.operationID
	shortOperationID := extensions.shortOperationID
	aliases := extensions.operationAliases

	httpVerb, err := api.ToHTTPVerb(verb)
	if err != nil {
		return nil, err
	}

	parameters, err := extractParameters(operation.Parameters)
	if err != nil {
		return nil, err
	}

	versions, err := buildVersions(now, operation)
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return nil, nil
	}

	description, err := buildDescription(operation)
	if err != nil {
		return nil, fmt.Errorf("failed to clean description: %w", err)
	}

	if overrides := extractOverrides(operation.Extensions); overrides != nil {
		if overriddenOperationID, ok := overrides["operationId"].(string); ok && overriddenOperationID != "" {
			operationID = overriddenOperationID
		}
	}

	watcher, err := extractWatcherProperties(operation.Extensions)
	if err != nil {
		return nil, err
	}

	command := api.Command{
		OperationID:      operationID,
		ShortOperationID: shortOperationID,
		Aliases:          aliases,
		Description:      description,
		RequestParameters: api.RequestParameters{
			URL:             path,
			QueryParameters: parameters.query,
			URLParameters:   parameters.url,
			Verb:            httpVerb,
		},
		Versions: versions,
		Watcher:  watcher,
	}

	return &command, nil
}

func buildDescription(operation *openapi3.Operation) (string, error) {
	// Get the tag and build the documentation URL
	if len(operation.Tags) != 1 {
		return "", fmt.Errorf("expect exactly 1 tag, got: %v", len(operation.Tags))
	}

	inputDescription := operation.Description

	if overrides := extractOverrides(operation.Extensions); overrides != nil {
		if overriddenDescription, ok := overrides["description"].(string); ok && overriddenDescription != "" {
			inputDescription = overriddenDescription
		}
	}

	// Get the original description and clean it up
	description, err := Clean(inputDescription)
	if err != nil {
		return "", fmt.Errorf("failed to clean description: %w", err)
	}

	apiURL := "https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/operation/operation-" + safeSlugify(strings.ToLower(operation.OperationID))
	cmdURL := fmt.Sprintf("https://www.mongodb.com/docs/atlas/cli/current/command/atlas-api-%s-%s/", strcase.ToLowerCamel(operation.Tags[0]), operation.OperationID)

	// Add the documentation URL to the description
	description += fmt.Sprintf(`

This command is autogenerated and corresponds 1:1 with the Atlas API endpoint %s.

For more information and examples, see: %s`, apiURL, cmdURL)

	return description, nil
}

// Struct to hold both types of parameters.
type parameterSet struct {
	query []api.Parameter
	url   []api.Parameter
}

func extractOverrides(ext map[string]any) map[string]any {
	if extensions, okExtensions := ext["x-xgen-atlascli"].(map[string]any); okExtensions && extensions != nil {
		if overrides, okOverrides := extensions["override"].(map[string]any); okOverrides && overrides != nil {
			return overrides
		}
	}
	return nil
}

func extractParametersNameDescription(parameterRef *openapi3.ParameterRef) (string, string) {
	parameter := parameterRef.Value
	parameterName := parameter.Name
	parameterDescription := parameter.Description

	if overrides := extractOverrides(parameterRef.Extensions); overrides != nil {
		if overriddenDescription, ok := overrides["description"].(string); ok && overriddenDescription != "" {
			parameterDescription = overriddenDescription
		}
		if overriddenName, ok := overrides["name"].(string); ok && overriddenName != "" {
			parameterName = overriddenName
		}
	} else if overrides := extractOverrides(parameter.Extensions); overrides != nil {
		if overriddenDescription, ok := overrides["description"].(string); ok && overriddenDescription != "" {
			parameterDescription = overriddenDescription
		}
		if overriddenName, ok := overrides["name"].(string); ok && overriddenName != "" {
			parameterName = overriddenName
		}
	}

	return parameterName, parameterDescription
}

type parameterExtensions struct {
	aliases []string
	short   string
}

func extractParameterExtensions(parameterRef *openapi3.ParameterRef) parameterExtensions {
	ext := parameterExtensions{
		aliases: []string{},
		short:   "",
	}

	extractParameterExtensionsMap(&ext, parameterRef.Extensions)

	value := parameterRef.Value
	if value != nil {
		extractParameterExtensionsMap(&ext, value.Extensions)
	}

	return ext
}

func extractParameterExtensionsMap(ext *parameterExtensions, extensionsMap map[string]any) {
	if extensionsMap == nil {
		return
	}

	if extensions, okExtensions := extensionsMap["x-xgen-atlascli"].(map[string]any); okExtensions && extensions != nil {
		if rawParameterAliases, ok := extensions["aliases"].([]any); ok && rawParameterAliases != nil {
			for _, rawParameterAlias := range rawParameterAliases {
				if parameterAlias, ok := rawParameterAlias.(string); ok {
					ext.aliases = append(ext.aliases, parameterAlias)
				}
			}
		}

		if flagShort, okFlagShort := extensions["flag-short"].(string); okFlagShort && ext.short == "" {
			ext.short = flagShort
		}
	}
}

// Extract and categorize parameters.
func extractParameters(parameters openapi3.Parameters) (parameterSet, error) {
	parameterNames := make(map[string]struct{})
	queryParameters := make([]api.Parameter, 0)
	urlParameters := make([]api.Parameter, 0)

	for _, parameterRef := range parameters {
		parameter := parameterRef.Value
		parameterName, parameterDescription := extractParametersNameDescription(parameterRef)

		parameterExtensions := extractParameterExtensions(parameterRef)
		aliases := parameterExtensions.aliases
		parameterShort := parameterExtensions.short

		// Parameters are translated to flags, we don't want duplicates
		// Duplicates should be resolved by customization, in case they ever appeared
		if _, exists := parameterNames[parameterName]; exists {
			return parameterSet{}, fmt.Errorf("parameter with the name '%s' already exists", parameter.Name)
		}

		description, err := Clean(parameterDescription)
		if err != nil {
			return parameterSet{}, fmt.Errorf("failed to clean description: %w", err)
		}

		parameterType, err := getParameterType(parameter)
		if err != nil {
			return parameterSet{}, err
		}

		apiParameter := api.Parameter{
			Name:        parameterName,
			Short:       parameterShort,
			Description: description,
			Required:    parameter.Required,
			Type:        *parameterType,
			Aliases:     aliases,
		}

		switch parameter.In {
		case "query":
			queryParameters = append(queryParameters, apiParameter)
			parameterNames[parameterName] = struct{}{}
		case "path":
			urlParameters = append(urlParameters, apiParameter)
			parameterNames[parameterName] = struct{}{}
		default:
			return parameterSet{}, fmt.Errorf("invalid parameter 'in' location: %s", parameter.In)
		}
	}

	return parameterSet{
		query: queryParameters,
		url:   urlParameters,
	}, nil
}

// Build versions from responses and request body.
func buildVersions(now time.Time, operation *openapi3.Operation) ([]api.CommandVersion, error) {
	versionsMap := make(map[string]*api.CommandVersion)

	if err := processResponses(operation.Responses, versionsMap); err != nil {
		return nil, err
	}

	if err := processRequestBody(operation.RequestBody, versionsMap); err != nil {
		return nil, err
	}

	// filter sunsetted versions
	for key, version := range versionsMap {
		if version.Sunset != nil && now.After(*version.Sunset) {
			delete(versionsMap, key)
		}
	}

	return sortVersions(versionsMap), nil
}

// Process response content types.
func processResponses(responses *openapi3.Responses, versionsMap map[string]*api.CommandVersion) error {
	for statusString, responses := range responses.Map() {
		statusCode, err := strconv.Atoi(statusString)
		if err != nil {
			return fmt.Errorf("http status code '%s' is not numeric: %w", statusString, err)
		}

		if statusCode < 200 || statusCode >= 300 {
			continue
		}

		for versionedContentType, mediaType := range responses.Value.Content {
			if err := addContentTypeToVersion(versionedContentType, versionsMap, mediaType.Extensions, false); err != nil {
				return err
			}
		}
	}
	return nil
}

// Process request body content types.
func processRequestBody(requestBody *openapi3.RequestBodyRef, versionsMap map[string]*api.CommandVersion) error {
	if requestBody == nil {
		return nil
	}

	for versionedContentType, mediaType := range requestBody.Value.Content {
		if mediaType.Schema == nil || (mediaType.Schema.Ref == "" && mediaType.Schema.Value == nil) {
			continue
		}

		if err := addContentTypeToVersion(versionedContentType, versionsMap, mediaType.Extensions, true); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to add content type to version map.
func addContentTypeToVersion(versionedContentType string, versionsMap map[string]*api.CommandVersion, extensions map[string]any, isRequest bool) error {
	// Extract the version and content type from the versioned content type.
	version, contentType, err := extractVersionAndContentType(versionedContentType)
	if err != nil {
		return fmt.Errorf("unsupported version %q error: %w", versionedContentType, err)
	}

	// Extract the sunset date and private preview from the extensions.
	sunset := extractSunsetDate(extensions)
	publicPreview := extractPublicPreview(extensions)

	// Add the version to the versions map if it doesn't exist.
	versionString := version.String()
	if _, ok := versionsMap[versionString]; !ok {
		versionsMap[versionString] = &api.CommandVersion{
			Version:              version,
			Sunset:               sunset,
			ResponseContentTypes: []string{},
		}
	}

	// If the sunset date is set, update the sunset date if it's before the current sunset date.
	if sunset != nil {
		if versionsMap[versionString].Sunset == nil || sunset.Before(*versionsMap[versionString].Sunset) {
			versionsMap[versionString].Sunset = sunset
		}
	}

	// The default for public preview is false, override it if the extension says we're in a public preview.
	if publicPreview != nil && *publicPreview {
		versionsMap[versionString].PublicPreview = true
	}

	// If the versioned content type is a request, set the request content type.
	// If the versioned content type is a response, add the content type to the response content types.
	if isRequest {
		if versionsMap[versionString].RequestContentType != "" {
			return errors.New("multiple request content types is not supported")
		}

		versionsMap[versionString].RequestContentType = contentType
	} else {
		versionsMap[versionString].ResponseContentTypes = append(versionsMap[versionString].ResponseContentTypes, contentType)
	}

	return nil
}

// Extract public preview from extensions.
// Example yaml:
// ```yaml
// x-xgen-preview:
//
//	public: 'true'
//
// ```
//
// If the extension is present, return true if the preview is public, false if it's private.
// If the extension is not present, return nil.
func extractPublicPreview(extensions map[string]any) *bool {
	if extensions, ok := extensions["x-xgen-preview"].(map[string]any); ok && extensions != nil {
		if public, ok := extensions["public"].(string); ok {
			publicPreview := public == "true"
			return &publicPreview
		}
	}

	return nil
}

// Sort versions and their content types.
func sortVersions(versionsMap map[string]*api.CommandVersion) []api.CommandVersion {
	versions := make([]api.CommandVersion, 0)

	for _, version := range versionsMap {
		sort.Slice(version.ResponseContentTypes, func(i, j int) bool {
			return version.ResponseContentTypes[i] < version.ResponseContentTypes[j]
		})

		versions = append(versions, *version)
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Version.Less(versions[j].Version)
	})

	return versions
}

func groupForTag(spec *openapi3.T, tag string) (*api.Group, error) {
	description := ""

	if specTag := spec.Tags.Get(tag); specTag != nil {
		cleanDescription, err := Clean(specTag.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to clean description: %w", err)
		}
		description = cleanDescription
	}

	return &api.Group{
		Name:        tag,
		Description: description,
		Commands:    []api.Command{},
	}, nil
}

func extractVersionAndContentType(input string) (api.Version, string, error) {
	matches := contentTypeHeaderRegex.FindStringSubmatch(input)
	if len(matches) == 0 {
		return nil, "", fmt.Errorf("invalid content type header: %s", input)
	}

	versionIndex := contentTypeHeaderRegex.SubexpIndex("version")
	contentTypeIndex := contentTypeHeaderRegex.SubexpIndex("contentType")

	versionString := matches[versionIndex]
	contentType := matches[contentTypeIndex]

	if versionString == "" {
		return nil, "", errors.New("version is required")
	}

	if contentType == "" {
		return nil, "", errors.New("content type is required")
	}

	version, err := api.ParseVersion(versionString)

	if err != nil {
		return nil, "", fmt.Errorf("invalid version: %w", err)
	}

	return version, contentType, nil
}

func getParameterType(parameter *openapi3.Parameter) (*api.ParameterType, error) {
	if parameter.Schema == nil {
		return nil, errors.New("parameter schema is nil")
	}

	// Handle arrays first
	if parameter.Schema.Value.Type.Is("array") {
		if parameter.Schema.Value.Items == nil {
			return nil, errors.New("array items schema is nil")
		}

		// Get the type of array items
		itemType, err := resolveSchemaType(parameter.Schema.Value.Items.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve array item type: %w", err)
		}

		return &api.ParameterType{
			IsArray: true,
			Type:    itemType,
		}, nil
	}

	// Handle non-array types
	paramType, err := resolveSchemaType(parameter.Schema.Value)
	if err != nil {
		return nil, err
	}

	return &api.ParameterType{
		IsArray: false,
		Type:    paramType,
	}, nil
}

// resolveSchemaType handles the conversion of OpenAPI types to Go types.
func resolveSchemaType(schema *openapi3.Schema) (api.ParameterConcreteType, error) {
	// Handle oneOf
	if len(schema.OneOf) > 0 {
		return resolveOneOfAnyOf(schema.OneOf)
	}

	// Handle anyOf
	if len(schema.AnyOf) > 0 {
		return resolveOneOfAnyOf(schema.AnyOf)
	}

	// Handle basic types
	switch {
	case schema.Type.Is("string"):
		return api.TypeString, nil
	case schema.Type.Is("integer"):
		return api.TypeInt, nil
	case schema.Type.Is("boolean"):
		return api.TypeBool, nil
	default:
		return "", fmt.Errorf("unsupported type: %s", schema.Type)
	}
}

// resolveOneOfAnyOf recursively resolves oneOf/anyOf schemas and returns the first matching basic type.
func resolveOneOfAnyOf(schemas []*openapi3.SchemaRef) (api.ParameterConcreteType, error) {
	for _, schema := range schemas {
		if schema == nil || schema.Value == nil {
			continue
		}

		// Recursive handling of nested oneOf/anyOf
		if len(schema.Value.OneOf) > 0 {
			if t, err := resolveOneOfAnyOf(schema.Value.OneOf); err == nil {
				return t, nil
			}
			continue
		}
		if len(schema.Value.AnyOf) > 0 {
			if t, err := resolveOneOfAnyOf(schema.Value.AnyOf); err == nil {
				return t, nil
			}
			continue
		}

		// Try to resolve the current schema
		if t, err := resolveSchemaType(schema.Value); err == nil {
			return t, nil
		}
	}

	return "", errors.New("no valid basic type found in oneOf/anyOff")
}
