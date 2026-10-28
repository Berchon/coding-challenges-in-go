package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxCandies := slices.Max(candies)
	for index, candy := range candies {
		result[index] = (candy + extraCandies) >= maxCandies
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	kidsWithCandies(candies, extraCandies)

	fmt.Println(kidsWithCandies(candies, extraCandies))
}
