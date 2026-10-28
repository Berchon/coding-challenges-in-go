package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	mdc := mdc(len(str1), len(str2))

	pattern := str1[:mdc]
	if divides(pattern, str1) && divides(pattern, str2) {
		return pattern
	}

	return ""
}

func divides(pattern, s string) bool {
	return strings.Repeat(pattern, len(s)/len(pattern)) == s
}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
