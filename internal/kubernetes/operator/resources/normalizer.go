package resources

import (
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"
)

// NormalizeAtlasResourceName normalizes the name to be used as a resource name in Kubernetes
// Run fuzzing test if you want to change this function!
func NormalizeAtlasResourceName(name string) string {
	dictionary := map[string]string{
		" ": "-",
		".": "dot",
		"@": "at",
		"(": "left-parenthesis",
		")": "right-parenthesis",
		"&": "and",
		"+": "plus",
		":": "colon",
		",": "comma",
		"'": "single-quote",
	}

	for k, v := range dictionary {
		name = strings.ReplaceAll(name, k, v)
	}

	restrictionForFirstAndLast := []string{"-", "_"}
	dictionaryForFirstAndLast := map[string]string{
		restrictionForFirstAndLast[0]: "dash",
		restrictionForFirstAndLast[1]: "underscore",
	}
	if len(name) > 0 {
		for _, v := range restrictionForFirstAndLast {
			if strings.HasPrefix(name, v) {
				name = dictionaryForFirstAndLast[v] + name[1:]
			}
			if strings.HasSuffix(name, v) {
				name = name[:len(name)-1] + dictionaryForFirstAndLast[v]
			}
		}
	}

	if len(name) > validation.DNS1123LabelMaxLength {
		name = name[:validation.DNS1123LabelMaxLength]
	}

	if len(name) > 0 && name[len(name)-1] == '-' {
		name = name[:len(name)-1] + "d"
	}

	return strings.ToLower(name)
}
