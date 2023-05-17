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
	res := strings.ToLower(string(words[0][0])) + words[0][1:]
	for _, word := range words[1:] {
		res += strings.Title(word)
	}
	return res
}

func SnakeCase(str string) string {
	str = strings.TrimSpace(str)
	snake := regexp.MustCompile("(.)([A-Z][^A-Z]+)").ReplaceAllString(str, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ReplaceAll(snake, " ", "")
	snake = regexp.MustCompile(`[^_a-zA-Z0-9]+`).ReplaceAllString(snake, "_")
	return strings.ToLower(snake)
}

var FuncMap = template.FuncMap{
	"PascalCase": PascalCase,
	"SnakeCase":  SnakeCase,
	"CamelCase":  CamelCase,
}
