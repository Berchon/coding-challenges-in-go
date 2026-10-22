package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 136. Single Number
### Easy

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1

Constraints:
1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
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

func singleNumber(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}

func main() {
	// número único no início.
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}) == 4) // true

	// número único no final.
	fmt.Println(singleNumber([]int{2, 2, 1}) == 1) // true

	// array com um único elemento.
	fmt.Println(singleNumber([]int{7}) == 7) // true
	fmt.Println(singleNumber([]int{0}) == 0) // true

	// números negativos.
	fmt.Println(singleNumber([]int{-1, 3, -1}) == 3)   // true
	fmt.Println(singleNumber([]int{-1, -3, -1}) == -3) // true

	// número único negativo.
	fmt.Println(singleNumber([]int{5, 5, -8, 2, 2}) == -8) // true
}
