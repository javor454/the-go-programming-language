package __5_4_strings

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// 3.10: write a non-recursive version of comma, using bytes.Buffer instead of string concatenation
func Comma2(s string) string {
	r := []rune(s)
	n := len(r)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := 1; i <= n; i++ {
		buf.WriteRune(r[i-1])

		if i%3 == 0 && i != n {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}

// 3.11: enhance comma so that it deals correctly with floating point numbers and an optional sign e.g. -12345.45 => -123,45.45
func Comma3(s string) string {
	reg, err := regexp.Compile("^[+-]?\\d+(\\.\\d+)?$")
	if err != nil {
		panic("bad regex")
	}
	if !reg.MatchString(s) {
		return "not int or float number with sign"
	}

	var buf bytes.Buffer
	var parts []string

	// handle optional sign
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteRune(rune(s[0]))
		parts = strings.Split(s[1:], ".")
	} else {
		parts = strings.Split(s, ".")
	}

	if len(parts) > 2 || len(parts) < 1 {
		return "invalid number"
	}

	r := []rune(parts[0])
	n := len(r)
	if n <= 3 {
		return s
	}

	for i := 1; i <= n; i++ {
		buf.WriteRune(r[i-1])

		if i%3 == 0 && i != n {
			buf.WriteRune(',')
		}
	}

	if len(parts) > 1 {
		buf.WriteString(fmt.Sprintf(".%s", parts[1]))
	}

	return buf.String()
}

// 3.12 write a function that reports whether two strings are anagrams of each other that is they contain the same letters in a different order
// TODO
