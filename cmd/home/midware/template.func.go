package midware

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

func Indent(n int, html template.HTML) template.HTML {
	startOfLine := regexp.MustCompile(`(?m)^`)
	indentation := strings.Repeat(" ", n)
	return template.HTML(startOfLine.ReplaceAllLiteralString(string(html), indentation))
}

func NIndent(n int, html template.HTML) template.HTML {
	startOfLine := regexp.MustCompile(`(?m)^`)
	indentation := strings.Repeat(" ", n)
	text := startOfLine.ReplaceAllLiteralString(string(html), indentation)
	return template.HTML(fmt.Sprintf("\n%s", text))
}

func Mod(num int, m int, eq int) bool {
	return (num+1)%m == eq
}
