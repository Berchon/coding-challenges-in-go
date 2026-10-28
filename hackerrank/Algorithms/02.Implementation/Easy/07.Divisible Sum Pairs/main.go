package main

import "fmt"

// func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
// 	remainderCounts := make([]int32, int(k))
// 	var pairCount int32

// 	for i := 0; i < int(n); i++ {
// 		remainder := ar[i] % k
// 		complement := (k - remainder) % k

// 		pairCount += remainderCounts[int(complement)]
// 		remainderCounts[int(remainder)]++
// 	}

// 	return pairCount
// }

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var pairCount int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if isDivisibleBy(ar[i]+ar[j], k) {
				pairCount++
			}
		}
	}
	return pairCount
}

func isDivisibleBy(value, divisor int32) bool {
	return value%divisor == 0
}

func main() {
	fmt.Println(divisibleSumPairs(int32(6), int32(3), []int32{1, 3, 2, 6, 1, 2})) // Output: 5
	fmt.Println(divisibleSumPairs(int32(6), int32(5), []int32{1, 2, 3, 4, 5, 6})) // Output: 3
}
