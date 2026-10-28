package main

import (
	"fmt"
	"slices"
	"sort"
)

func Ordering(r []rune) {
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
}

func IsAnagram(s, t string) bool {
	firstString := []rune(s)
	secondString := []rune(t)

	if len(firstString) != len(secondString) {
		return false
	}

	slices.Sort(firstString)
	// Ordering(firstString)
	Ordering(secondString)

	// return string(firstString) == string(secondString)
	return slices.Equal(firstString, secondString)
}

func main() {
	fmt.Println(IsAnagram("abc", "abcd"))        // tamanhos diferentes
	fmt.Println(IsAnagram("abc", "def"))         // mesmo tamanho, mas não é anagrama
	fmt.Println(IsAnagram("listen ", "sil ent")) // mesmo tamanho e é anagrama
	fmt.Println(IsAnagram("listen", "silent "))  // tamanhos diferentes
	fmt.Println(IsAnagram("", ""))               // mesmo tamanho e é anagrama
	fmt.Println(IsAnagram("", " "))              // tamanhos diferentes
	fmt.Println(IsAnagram("Abc", "abc"))         // mesmo tamanho, mas não é anagrama (case-sensitive)
	fmt.Println(IsAnagram("olá", "ola"))         // tamanhos diferentes (acentos)
	fmt.Println(IsAnagram("ol@", "#ol"))         // mesmo tamanho, mas não é anagrama (caracteres especiais)
	fmt.Println(IsAnagram("ol@", "@ol"))         // mesmo tamanho e é anagrama (caracteres especiais)
	fmt.Println(IsAnagram("listen", "silent"))   // mesmo tamanho e é anagrama
}
