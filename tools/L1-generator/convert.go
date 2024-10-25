package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/autogeneration/L1"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

var (
	versionRegex = regexp.MustCompile(`^application/vnd\.atlas\.(?P<version>\d{4}-\d{2}-\d{2})\+(?P<contentType>[\w]+)$`)
)

func specToCommands(spec *openapi3.T) (L1.GroupedAndSortedCommands, error) {
	groups := make(map[string]*L1.Group, 0)

	for path, item := range spec.Paths.Map() {
		for verb, operation := range item.Operations() {
			command, err := operationToCommand(path, verb, *operation)
			if err != nil {
				return nil, fmt.Errorf("failed to convert operation to command: %w", err)
			}

			if len(operation.Tags) != 1 {
				return nil, fmt.Errorf("expect every operation to have exactly 1 tag, got: %v", len(operation.Tags))
			}

			tag := operation.Tags[0] // TODO: verify length
			if _, ok := groups[tag]; !ok {
				group, err := groupForTag(spec, tag)
				if err != nil {
					return nil, fmt.Errorf("failed to create group from tag: %w", err)
				}

				groups[tag] = group
			}

			groups[tag].Commands = append(groups[tag].Commands, *command)
		}
	}

	sortedGroups := make([]L1.Group, 0, len(groups))
	for _, group := range groups {
		sort.Slice(group.Commands, func(i, j int) bool {
			return group.Commands[i].OperationID < group.Commands[j].OperationID
		})

		sortedGroups = append(sortedGroups, *group)
	}

	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Name < sortedGroups[j].Name
	})

	return sortedGroups, nil
}

//nolint:gocyclo
func operationToCommand(path, verb string, operation openapi3.Operation) (*L1.Command, error) {
	httpVerb, err := L1.ToHTTPVerb(verb)
	if err != nil {
		return nil, err
	}

	queryParameters := make([]L1.Parameter, 0)
	urlParameters := make([]L1.Parameter, 0)

	for _, parameterRef := range operation.Parameters {
		parameter := parameterRef.Value

		l1Parameter := L1.Parameter{
			Name:        parameter.Name,
			Description: cleanString(parameter.Description),
			Required:    parameter.Required,
		}

		switch parameter.In {
		case "query":
			queryParameters = append(queryParameters, l1Parameter)
		case "path":
			urlParameters = append(urlParameters, l1Parameter)
		default:
			return nil, fmt.Errorf("invalid parameter 'in' location: %s", parameter.In)
		}
	}

	versionsMap := make(map[string]*L1.Version, 0)

	for statusString, responses := range operation.Responses.Map() {
		statusCode, err := strconv.Atoi(statusString)
		if err != nil {
			return nil, fmt.Errorf("http status code '%s' is not numeric: %w", statusString, err)
		}

		if statusCode < 200 || statusCode >= 300 {
			continue
		}

		// TODO: extract sunset data (x-sunset) from _ parameter
		for versionedContentType := range responses.Value.Content {
			version, contentType, err := extractVersionAndContentType(versionedContentType)
			if err != nil {
				return nil, fmt.Errorf("unsupported version '%s' error: %w", versionedContentType, err)
			}

			if _, ok := versionsMap[version]; !ok {
				versionsMap[version] = &L1.Version{
					Version:              version,
					RequestContentTypes:  []string{},
					ResponseContentTypes: []string{},
				}
			}

			versionsMap[version].ResponseContentTypes = append(versionsMap[version].ResponseContentTypes, contentType)
		}
	}

	// TODO: extract sunset data (x-sunset) from _ parameter
	if operation.RequestBody != nil {
		for versionedContentType := range operation.RequestBody.Value.Content {
			version, contentType, err := extractVersionAndContentType(versionedContentType)
			if err != nil {
				return nil, fmt.Errorf("unsupported version '%s' error: %w", versionedContentType, err)
			}

			if _, ok := versionsMap[version]; !ok {
				versionsMap[version] = &L1.Version{
					Version:              version,
					RequestContentTypes:  []string{},
					ResponseContentTypes: []string{},
				}
			}

			versionsMap[version].RequestContentTypes = append(versionsMap[version].RequestContentTypes, contentType)
		}
	}

	versions := make([]L1.Version, 0)
	for _, version := range versionsMap {
		sort.Slice(version.RequestContentTypes, func(i, j int) bool {
			return version.RequestContentTypes[i] < version.RequestContentTypes[j]
		})

		sort.Slice(version.ResponseContentTypes, func(i, j int) bool {
			return version.ResponseContentTypes[i] < version.ResponseContentTypes[j]
		})

		versions = append(versions, *version)
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Version < versions[j].Version
	})

	command := L1.Command{
		OperationID: operation.OperationID,
		Description: cleanString(operation.Description),
		RequestParameters: L1.RequestParameters{
			URL:             path,
			QueryParameters: queryParameters,
			URLParameters:   urlParameters,
			Verb:            httpVerb,
		},
		Versions: versions,
	}

	return &command, nil
}

func groupForTag(spec *openapi3.T, tag string) (*L1.Group, error) {
	specTag := spec.Tags.Get(tag)
	if specTag == nil {
		return nil, fmt.Errorf("tag '%s' not found", tag)
	}

	return &L1.Group{
		Name:        specTag.Name,
		Description: cleanString(specTag.Description),
		Commands:    []L1.Command{},
	}, nil
}

func cleanString(input string) string {
	inputBytes := []byte(input)

	md := goldmark.New()
	root := md.Parser().Parse(text.NewReader(inputBytes))

	// TODO: use non deprecated method

	plain := string(root.Text(inputBytes))
	cleaned := strings.TrimSpace(strings.ReplaceAll(plain, "`", "'"))

	return cleaned
}

func extractVersionAndContentType(input string) (version string, contentType string, err error) {
	matches := versionRegex.FindStringSubmatch(input)
	if matches == nil {
		return "", "", errors.New("invalid format")
	}

	// Get the named group indices
	versionIndex := versionRegex.SubexpIndex("version")
	contentTypeIndex := versionRegex.SubexpIndex("contentType")

	return matches[versionIndex], matches[contentTypeIndex], nil
}
