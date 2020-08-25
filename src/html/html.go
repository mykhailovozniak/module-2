package html

import (
	"strings"
)

func RenderUL(items []string) string {
	var html strings.Builder

	html.WriteString("<ul>")
	for _, item := range items {
		html.WriteString("<li>" + item + "</li>")
	}
	html.WriteString("</ul>")

	return html.String()
}

func RenderHTML5(htmlElem string) string  {
	var html strings.Builder

	html.WriteString("<!DOCTYPE html><html></body>")
	html.WriteString(htmlElem)
	html.WriteString("</body></html>")

	return html.String()
}
