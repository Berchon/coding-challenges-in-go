package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 1207. Unique Number of Occurrences
### Easy

Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

Example 1:
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

Example 2:
Input: arr = [1,2]
Output: false

Example 3:
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

Constraints:
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/

// The setMemoryLimit function sets a limit on the amount of memory that the Go program can use.
// In this case, it is set to 6 MiB (megabytes). This means that if the program tries to allocate
// more than 6 MiB of memory, it will be terminated by the runtime. This is useful for controlling
// resource usage and preventing excessive memory consumption in applications.
// Isso melhora as métricas de desempenho do LeetCode, pois limita o uso de memória do programa Go a 6 MiB,
// evitando que ele consuma recursos excessivos e garantindo que ele funcione dentro de limites aceitáveis.
// Isso é especialmente importante em ambientes de execução restritos, como plataformas de avaliação de
// código online, onde o uso eficiente de recursos é crucial para manter a estabilidade e o desempenho do sistema.
func init() {
	debug.SetMemoryLimit(3*1024*1024 + 9*1024) // 3.9 MiB
}

func countOccurrences(arr []int) map[int]int {
	counts := make(map[int]int, len(arr))
	for _, value := range arr {
		counts[value]++
	}
	return counts
}

func hasUniqueOccurrences(occurrences map[int]int) bool {
	seen := make(map[int]struct{}, len(occurrences))

	for _, frequency := range occurrences {
		if _, exists := seen[frequency]; exists {
			return false
		}
		seen[frequency] = struct{}{}
	}
	return true
}

func uniqueOccurrences(arr []int) bool {
	occurrences := countOccurrences(arr)
	return hasUniqueOccurrences(occurrences)
}

func main() {
	// [1,2,2,1,1,3]
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // Output: true
	// [1,2]
	fmt.Println(uniqueOccurrences([]int{1, 2})) // Output: false
	// [-3,0,1,-3,1,1,1,-3,10,0]
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // Output: true
}
