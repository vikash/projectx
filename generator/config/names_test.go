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
		"imageUrl":         "imageUrl",
		"imageURL":         "imageURL",
	}

	testNameFunctiont(t, cases, CamelCase)
}

func TestGetNameInPascalCase(t *testing.T) {
	cases := map[string]string{
		"brand ":           "Brand",
		"Brand":            "Brand",
		"Something Simple": "SomethingSimple",
		"she@something":    "SheSomething",
		"imageUrl":         "ImageUrl",
		"imageURL":         "ImageURL",
	}

	testNameFunctiont(t, cases, PascalCase)
}

func TestGetNameInSnakeCase(t *testing.T) {
	cases := map[string]string{
		"brand ":           "brand",
		"Brand":            "brand",
		"Something Simple": "something_simple",
		"she@something":    "she_something",
		"SomethingSimple":  "something_simple",
		"imageURL":         "image_url",
		"imageUrl":         "image_url",
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
