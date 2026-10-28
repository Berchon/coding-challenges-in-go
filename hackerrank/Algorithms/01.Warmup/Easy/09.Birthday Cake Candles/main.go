package main

import "fmt"

// Minha solução
func birthdayCakeCandles(candles []int32) int32 {
	maxHeight := candles[0]
	count := int32(1)

	for _, height := range candles[1:] {
		if height > maxHeight {
			maxHeight = height
			count = 1
		} else if height == maxHeight {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{1, 2, 3, 4, 5}))
}
