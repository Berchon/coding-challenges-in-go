package main

import (
	"fmt"
	"sort"
)

// func isSubsequence(s string, t string) bool {
// 	subsequenceIndex := 0
// 	for _, char := range t {
// 		if subsequenceIndex == len(s) {
// 			break
// 		}

// 		if char == rune(s[subsequenceIndex]) {
// 			subsequenceIndex++
// 		}
// 	}

// 	return subsequenceIndex == len(s)
// }

// func isSubsequenceBatch(s []string, t string) []bool {
// 	result := make([]bool, len(s))
// 	for i, sequence := range s {
// 		result[i] = isSubsequence(sequence, t)
// 	}
// 	return result
// }

// func isSubsequence(s string, t string) bool {
// 	if len(s) > len(t) {
// 		return false
// 	}
// 	index := 0

// 	// etapa 1: Criar um mapa para armazenar as posições de cada caractere em t
// 	positions := make(map[byte][]int) // Used 'byte' instead of 'rune' for efficiency, since we are dealing with ASCII lowercase English letters

// 	for i := 0; i < len(t); i++ {
// 		key := t[i]
// 		positions[key] = append(positions[key], i)
// 	}

// 	// etapa 2: Verificar se s é uma subsequência de t
// 	for i := 0; i < len(s); i++ {
// 		// etapa 2.1: Verificar se o caractere atual de s existe em t
// 		if positions[s[i]] == nil {
// 			return false
// 		}

// 		if index >= positions[s[i]][0] {
// 			return false // The character in s is not in the correct order in t
// 		}

// 		index = positions[s[i]][0]            // Update index to the position of the current character in t
// 		positions[s[i]] = positions[s[i]][1:] // Remove the first occurrence of the character in positions[s[i]]

// 	}
// 	return true
// }

type SubsequenceIndex struct {
	positions map[byte][]int
}

func isSubsequence(s string, t string) bool {
	index := newSubsequenceIndex(t)

	return checkSubsequence(s, index)
}

func isSubsequenceBatch(sequences []string, t string) []bool {
	index := newSubsequenceIndex(t)

	return checkSubsequences(sequences, index)
}

func newSubsequenceIndex(t string) SubsequenceIndex {
	positions := make(map[byte][]int)

	for i := 0; i < len(t); i++ {
		char := t[i]
		positions[char] = append(positions[char], i)
	}

	return SubsequenceIndex{
		positions: positions,
	}
}

func checkSubsequences(sequences []string, index SubsequenceIndex) []bool {
	result := make([]bool, len(sequences))

	for i, sequence := range sequences {
		result[i] = checkSubsequence(sequence, index)
	}

	return result
}

func checkSubsequence(s string, index SubsequenceIndex) bool {
	currentPosition := -1

	for i := 0; i < len(s); i++ {
		char := s[i]

		positions, exists := index.positions[char]
		if !exists {
			return false
		}

		nextPosition, found := findNextPosition(positions, currentPosition)
		if !found {
			return false
		}

		currentPosition = nextPosition
	}

	return true
}

func findNextPosition(positions []int, currentPosition int) (int, bool) {
	nextIndex := sort.Search(len(positions), func(i int) bool {
		return positions[i] > currentPosition
	})

	if nextIndex == len(positions) {
		return 0, false
	}

	return positions[nextIndex], true
}

func main() {
	fmt.Println("===== isSubsequence =====")

	fmt.Println("Caminho feliz:")
	fmt.Println(isSubsequence("abc", "ahbgdc")) // esperado: true
	fmt.Println(isSubsequence("a", "a"))        // esperado: true
	fmt.Println(isSubsequence("ab", "abc"))     // esperado: true
	fmt.Println(isSubsequence("ace", "abcde"))  // esperado: true

	fmt.Println("\nCaminho triste:")
	fmt.Println(isSubsequence("axc", "ahbgdc")) // esperado: false
	fmt.Println(isSubsequence("aec", "abcde"))  // esperado: false
	fmt.Println(isSubsequence("abcd", "abc"))   // esperado: false
	fmt.Println(isSubsequence("z", "abc"))      // esperado: false

	fmt.Println("\nEdge cases:")
	fmt.Println(isSubsequence("", ""))       // esperado: true
	fmt.Println(isSubsequence("", "abc"))    // esperado: true
	fmt.Println(isSubsequence("a", ""))      // esperado: false
	fmt.Println(isSubsequence("abc", ""))    // esperado: false
	fmt.Println(isSubsequence("a", "a"))     // esperado: true
	fmt.Println(isSubsequence("aaa", "aaa")) // esperado: true

	fmt.Println("\n===== isSubsequenceBatch =====")

	fmt.Println("Caminho feliz:")
	fmt.Println(
		isSubsequenceBatch(
			[]string{"abc", "a", "ab", "ac", "bc"},
			"ahbgdc",
		),
	) // esperado: [true true true true true]

	fmt.Println("\nCaminho triste:")
	fmt.Println(
		isSubsequenceBatch(
			[]string{"axc", "aec", "xyz", "abcd"},
			"ahbgdc",
		),
	) // esperado: [false false false false]

	fmt.Println("\nEdge cases:")
	fmt.Println(
		isSubsequenceBatch(
			[]string{""},
			"abc",
		),
	) // esperado: [true]

	fmt.Println(
		isSubsequenceBatch(
			[]string{"", "a", "z"},
			"abc",
		),
	) // esperado: [true true false]

	fmt.Println(
		isSubsequenceBatch(
			[]string{},
			"abc",
		),
	) // esperado: []
}
