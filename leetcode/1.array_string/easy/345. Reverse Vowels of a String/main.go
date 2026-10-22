package main

import "fmt"

/*
# 345. Reverse Vowels of a String
### Easy

Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:
Input: s = "IceCreAm"
Output: "AceCreIm"

Explanation:
The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/

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
