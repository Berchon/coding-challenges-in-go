package main

import (
	"fmt"
	"slices"
)

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	slices.Sort(keyboards)
	slices.Sort(drives)

	best := int32(-1)

	i, j := 0, len(drives)-1

	for i < len(keyboards) && j >= 0 {
		total := keyboards[i] + drives[j]

		switch {
		case total > b:
			j--

		default:
			best = max(best, total)
			i++
		}

		if best == b {
			return best
		}
	}

	return best
}

func main() {
	fmt.Println(getMoneySpent([]int32{40, 50, 60}, []int32{5, 8, 12}, 60))
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
	fmt.Println(getMoneySpent([]int32{4}, []int32{5}, 5))
}
