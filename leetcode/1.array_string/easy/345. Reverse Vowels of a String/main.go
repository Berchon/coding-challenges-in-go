package main

import "fmt"

func isVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func findNextVowel(runes []rune, start, limit int) int {
	for start < limit && !isVowel(runes[start]) {
		start++
	}

	return start
}

func findPreviousVowel(runes []rune, start, limit int) int {
	for start > limit && !isVowel(runes[start]) {
		start--
	}

	return start
}

func reverseVowels(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		left = findNextVowel(runes, left, right)
		right = findPreviousVowel(runes, right, left)

		if left >= right {
			break
		}

		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	s := "IceCreAm"
	result := reverseVowels(s)
	fmt.Println(result) // Output: "AceCreIm"
}
