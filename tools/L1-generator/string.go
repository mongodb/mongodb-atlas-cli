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
