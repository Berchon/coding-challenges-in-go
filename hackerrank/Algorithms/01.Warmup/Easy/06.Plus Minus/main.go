package main

import "fmt"

// Minha solução
func plusMinus(arr []int32) {
	positives, negatives, zeros := 0, 0, 0

	for _, value := range arr {
		switch {
		case value > 0:
			positives++
		case value < 0:
			negatives++
		case value == 0:
			zeros++
		}
	}

	size := float64(len(arr))

	fmt.Printf("%.6f\n", float64(positives)/size)
	fmt.Printf("%.6f\n", float64(negatives)/size)
	fmt.Printf("%.6f\n", float64(zeros)/size)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(arr)
}
