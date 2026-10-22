package main

import (
	"fmt"
	"strings"
)

/*
# 151. Reverse Words in a String

### Medium

Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:
1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

func reverseWords(s string) string {
	result := strings.Builder{}
	result.Grow(len(s))

	wordEnd := len(s) - 1

	for wordEnd >= 0 {
		wordEnd = skipSpaces(s, wordEnd)

		if wordEnd < 0 {
			break
		}

		wordStart := findWordStart(s, wordEnd)

		appendWord(&result, s, wordStart, wordEnd)

		wordEnd = wordStart - 1
	}

	return result.String()
}

func skipSpaces(s string, position int) int {
	for position >= 0 && s[position] == ' ' {
		position--
	}

	return position
}

func findWordStart(s string, wordEnd int) int {
	wordStart := wordEnd

	for wordStart >= 0 && s[wordStart] != ' ' {
		wordStart--
	}

	return wordStart + 1
}

func appendWord(result *strings.Builder, s string, wordStart, wordEnd int) {
	if result.Len() > 0 {
		result.WriteByte(' ')
	}

	result.WriteString(s[wordStart : wordEnd+1])
}

func main() {
	// Casos de teste baseados nos exemplos do problema
	tests := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
	}

	for _, s := range tests {
		fmt.Printf("Input:  %q\n", s)
		fmt.Printf("Output: %q\n\n", reverseWords(s))
	}
}
