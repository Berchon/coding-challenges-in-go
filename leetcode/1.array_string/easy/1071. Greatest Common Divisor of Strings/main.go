package main

import "strings"

/*
# 1071. Greatest Common Divisor of Strings
### Easy

For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).
Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
Input: str1 = "LEET", str2 = "CODE"
Output: ""

Example 4:
Input: str1 = "AAAAAB", str2 = "AAA"
Output: ""​​​​​​​

Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.
*/

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
