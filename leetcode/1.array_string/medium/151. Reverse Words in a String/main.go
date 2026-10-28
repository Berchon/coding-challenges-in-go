package main

import (
	"fmt"
	"strings"
)

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
