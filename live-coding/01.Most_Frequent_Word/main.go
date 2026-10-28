package main

import (
	"errors"
	"fmt"
)

func MostFrequent(words []string) (string, error) {
	if len(words) == 0 {
		return "", errors.New("a lista de palavras não pode estar vazia")
	}

	frequency := make(map[string]int)
	mostFrequentWord := ""
	maxFrequency := 0
	for _, word := range words {
		frequency[word]++
		if frequency[word] > maxFrequency {
			maxFrequency = frequency[word]
			mostFrequentWord = word
		}
	}

	for word, count := range frequency {
		fmt.Printf("Palavra: %s, Frequência: %d\n", word, count)
	}

	return mostFrequentWord, nil
}

func main() {
	words := []string{"maçã", "banana", "maçã", "laranja", "banana", "maçã"}
	mostFrequentWord, err := MostFrequent(words)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("A palavra mais frequente é:", mostFrequentWord)
}
