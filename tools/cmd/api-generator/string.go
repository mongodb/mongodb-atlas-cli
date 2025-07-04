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
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"golang.org/x/net/html"
)

func Clean(input string) (string, error) {
	htmlString, err := markdownToHTML(input)
	if err != nil {
		return "", err
	}

	cleaned := cleanHTML(htmlString)

	return cleaned, nil
}

func markdownToHTML(input string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(input), &buf); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func cleanHTML(input string) string {
	// Create a reader from the HTML string
	reader := strings.NewReader(input)

	// Parse HTML
	doc, err := html.Parse(reader)
	if err != nil {
		return input
	}

	var buf bytes.Buffer
	extractText(doc, &buf)
	return strings.TrimSpace(buf.String())
}

func extractText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}

	// Add newlines before block-level elements
	addNewLineForBlockLevelElements(n, buf)

	// Extract text
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, buf)
	}

	// Add newline after block-level elements
	addNewLineForBlockLevelElements(n, buf)
}

func addNewLineForBlockLevelElements(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "p", "div", "br", "h1", "h2", "h3", "h4", "h5", "h6", "table", "tr", "li":
			buf.WriteString("\n")
		}
	}
}

// Simplified port of:
// - https://github.com/Redocly/redoc/blob/59ee73fefa8e8edb398940076bdd721fc284caa3/src/utils/helpers.ts#L123
// - https://github.com/simov/slugify/blob/master/slugify.js
//
// Note: this does not handle special characters, which is fine for now as none of or tags or OperationIDs use special characters.
func safeSlugify(value string) string {
	// Trim spaces from both ends
	result := strings.TrimSpace(value)

	// Replace whitespace with single dash
	words := strings.Fields(result)
	result = strings.Join(words, "-")

	// Replace & with -and-
	result = strings.ReplaceAll(result, "&", "-and-")

	// Replace multiple consecutive dashes with single dash
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}

	// Trim dashes from start and end
	result = strings.Trim(result, "-")

	return result
}
