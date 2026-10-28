package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	best, worst  := scores[0], scores[0]
	bestBreaks, worstBreaks := int32(0), int32(0)

	for _, score := range scores[1:] {
		if score > best {
			best = score
			bestBreaks++
		}

		if score < worst {
			worst = score
			worstBreaks++
		}
	}

	return []int32{bestBreaks, worstBreaks}
}

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}

	result := breakingRecords(scores)

	fmt.Println(result)
	// Resultado esperado: [2 4]
}
