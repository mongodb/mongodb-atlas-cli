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
	"regexp"
	"strings"
)

func splitShortAndLongDescription(description string) (string, string) {
	// Split on periods that are followed by a space or end of string
	// This approach allows us to not accidentally split verion numbers like 8.0
	const maxSegments = 2
	split := regexp.MustCompile(`\.(?:\s+|$)`).Split(description, maxSegments)

	// Short description is the first sentence
	shortDescription := split[0]

	// Add the dot back, if needed
	if shortDescription != "" && !strings.HasSuffix(shortDescription, ".") && !strings.HasSuffix(shortDescription, ". ") {
		shortDescription += "."
	}

	// Long descriptions is everything after the first sentence
	longDescription := ""

	if len(split) > 1 {
		longDescription = strings.TrimSpace(split[1])
	}

	return shortDescription, longDescription
}
