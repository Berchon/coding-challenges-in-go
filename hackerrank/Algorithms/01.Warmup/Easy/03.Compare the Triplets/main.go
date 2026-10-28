package main

// Imports
import "fmt"

// Minha solução
func compareTriplets(a []int32, b []int32) []int32 {
	score := []int32{0, 0}
	minLength := min(len(a), len(b))

	for i := 0; i < minLength; i++ {
		switch {
		case a[i] > b[i]:
			score[0]++
		case a[i] < b[i]:
			score[1]++
		}
	}

	return score
}

// Exemplo para execução local
func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}

	result := compareTriplets(a, b)

	fmt.Println(result)
}
