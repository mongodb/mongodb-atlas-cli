package api

import (
	"regexp"
	"strings"
)

func splitShortAndLongDescription(description string) (string, string) {
	// Split on periods that are followed by a space or end of string
	// This approach allows us to not accidentally split verion numbers like 8.0
	split := regexp.MustCompile(`\.(?:\s+|$)`).Split(description, -1)

	// Short description is the first sentence
	shortDescription := split[0]

	// Add the dot back, if needed
	if shortDescription != "" && !strings.HasSuffix(shortDescription, ".") && !strings.HasSuffix(shortDescription, ". ") {
		shortDescription += "."
	}

	// Long descriptions is everything after the first sentence
	longDescription := ""

	if len(split) > 1 {
		// Remove all empty whitespace around sentences
		// This turns multi-line descriptions into single line
		for i, s := range split[1:] {
			split[i+1] = strings.TrimSpace(s)
		}

		// Add the ". " back
		longDescription = strings.Join(split[1:], ". ")
	}

	// Get rid of the last space after ". "
	longDescription = strings.TrimSpace(longDescription)

	return shortDescription, longDescription
}
