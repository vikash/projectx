package config

import (
	"testing"
)

func TestGetNameInCamelCase(t *testing.T) {
	cases := map[string]string{
		"brand ":           "brand",
		"Brand":            "brand",
		"Something Simple": "somethingSimple",
		"she@something":    "sheSomething",
	}

	testNameFunctiont(t, cases, CamelCase)
}

func TestGetNameInPascalCase(t *testing.T) {
	cases := map[string]string{
		"brand ":           "Brand",
		"Brand":            "Brand",
		"Something Simple": "SomethingSimple",
		"she@something":    "SheSomething",
	}

	testNameFunctiont(t, cases, PascalCase)
}

func TestGetNameInSnakeCase(t *testing.T) {
	cases := map[string]string{
		"brand ":           "brand",
		"Brand":            "brand",
		"Something Simple": "something_simple",
		"she@something":    "she_something",
	}

	testNameFunctiont(t, cases, SnakeCase)
}

func testNameFunctiont(t *testing.T, testCases map[string]string, f func(string) string) {
	for c, expected := range testCases {
		name := f(c)
		if name != expected {
			t.Errorf("For testcase: '%s', Expected '%s' got '%s'", c, expected, name)
		}
	}
}
