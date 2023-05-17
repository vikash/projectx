package config

import (
	"regexp"
	"strings"
	"text/template"
)

func PascalCase(str string) string {
	str = strings.TrimSpace(str)
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, " ")
	str = strings.Title(str)
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func CamelCase(str string) string {
	str = strings.TrimSpace(str)
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, " ")
	words := strings.Split(str, " ")
	res := strings.ToLower(words[0])
	for _, word := range words[1:] {
		res += strings.Title(word)
	}
	return res
}

func SnakeCase(str string) string {
	str = strings.TrimSpace(str)
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, " ")
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "_")
	return str
}

var FuncMap = template.FuncMap{
	"PascalCase": PascalCase,
	"SnakeCase":  SnakeCase,
	"CamelCase":  CamelCase,
}
