package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	builder.Grow(len(word1) + len(word2))

	i := 0
	for i = 0; i < len(word1) && i < len(word2); i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	builder.WriteString(word1[i:])
	builder.WriteString(word2[i:])

	return builder.String()
}

func main() {
	word1 := "abcde"
	word2 := "fghijk"

	fmt.Println(mergeAlternately(word1, word2))
}
