package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	unpaired := make(map[int32]struct{})
	pairCount := int32(0)

	for _, color := range ar {
		if _, ok := unpaired[color]; ok {
			pairCount++
			delete(unpaired, color)
			continue
		}

		unpaired[color] = struct{}{}
	}

	return int32(pairCount)
}

func main() {
	n := int32(9)
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	result := sockMerchant(n, ar)

	fmt.Printf("Quantidade de pares: %d\n", result)
	// Resultado esperado: 3
}
