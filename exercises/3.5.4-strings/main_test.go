package __5_4_strings

import (
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestComma3(t *testing.T) {
	testcases := []TestCase{
		{
			input:    "123123123123123",
			expected: "123,123,123,123,123",
		},
		{
			input:    "123",
			expected: "123",
		},
		{
			input:    "+123",
			expected: "+123",
		},
		{
			input:    "-123",
			expected: "-123",
		},
		{
			input:    "12345",
			expected: "123,45",
		},
		{
			input:    "+12345",
			expected: "+123,45",
		},
		{
			input:    "-12345",
			expected: "-123,45",
		},
		{
			input:    "123.45",
			expected: "123.45",
		},
		{
			input:    "+123.45",
			expected: "+123.45",
		},
		{
			input:    "-123.45",
			expected: "-123.45",
		},
		{
			input:    "12345.45",
			expected: "123,45.45",
		},
		{
			input:    "+12345.45",
			expected: "+123,45.45",
		},
		{
			input:    "-12345.45",
			expected: "-123,45.45",
		},
		{
			input:    "123.34.45",
			expected: "not int or float number with sign",
		},
		{
			input:    "+123.34.45",
			expected: "not int or float number with sign",
		},
		{
			input:    "-123.34.45",
			expected: "not int or float number with sign",
		},
		{
			input:    "sdfasda",
			expected: "not int or float number with sign",
		},
		{
			input:    "+sdfasda",
			expected: "not int or float number with sign",
		},
		{
			input:    "-sdfasda",
			expected: "not int or float number with sign",
		},
		{
			input:    ".454123",
			expected: "not int or float number with sign",
		},
		{
			input:    "+.454123",
			expected: "not int or float number with sign",
		},
		{
			input:    "-.454123",
			expected: "not int or float number with sign",
		},
		{
			input:    "+12345.",
			expected: "not int or float number with sign",
		},
	}

	for _, tc := range testcases {
		actual := Comma3(tc.input)
		if tc.expected != actual {
			t.Fatalf("input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
		}
	}
}
