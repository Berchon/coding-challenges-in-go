package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 724. Find Pivot Index
### Easy

Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.

Example 1:
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

Example 2:
Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.

Example 3:
Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0

Constraints:
1 <= nums.length <= 104
-1000 <= nums[i] <= 1000

Note: This question is the same as 1991: https://leetcode.com/problems/find-the-middle-index-in-array/
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

func pivotIndex(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	left := 0

	for i, num := range nums {
		right := sum - left - num

		if left == right {
			return i
		}

		left += num
	}

	return -1
}

func main() {
	// =========================================================
	// Casos oficiais do LeetCode
	// =========================================================

	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}), pivotIndex([]int{1, 7, 3, 6, 5, 6}) == 3) // expected: 3
	fmt.Println(pivotIndex([]int{1, 2, 3}), pivotIndex([]int{1, 2, 3}) == -1)                  // expected: -1
	fmt.Println(pivotIndex([]int{2, 1, -1}), pivotIndex([]int{2, 1, -1}) == 0)                 // expected: 0

	// =========================================================
	// Array com apenas um elemento
	// =========================================================

	fmt.Println(pivotIndex([]int{5}), pivotIndex([]int{5}) == 0)   // expected: 0
	fmt.Println(pivotIndex([]int{-5}), pivotIndex([]int{-5}) == 0) // expected: 0
	fmt.Println(pivotIndex([]int{0}), pivotIndex([]int{0}) == 0)   // expected: 0

	// =========================================================
	// Pivot no primeiro indice
	// =========================================================

	fmt.Println(pivotIndex([]int{2, 1, -1}), pivotIndex([]int{2, 1, -1}) == 0)      // expected: 0
	fmt.Println(pivotIndex([]int{0, 1, 2, 3}), pivotIndex([]int{0, 1, 2, 3}) == -1) // expected: -1

	// =========================================================
	// Pivot no ultimo indice
	// =========================================================

	fmt.Println(pivotIndex([]int{1, -1, 0}), pivotIndex([]int{1, -1, 0}) == 2) // expected: 2
	fmt.Println(pivotIndex([]int{1, -1}), pivotIndex([]int{1, -1}) == -1)      // expected: -1

	// =========================================================
	// Multiplos pivots
	// Deve retornar o menor indice
	// =========================================================

	fmt.Println(pivotIndex([]int{0, 0, 0}), pivotIndex([]int{0, 0, 0}) == 0)       // expected: 0
	fmt.Println(pivotIndex([]int{0, 0, 0, 0}), pivotIndex([]int{0, 0, 0, 0}) == 0) // expected: 0
	fmt.Println(pivotIndex([]int{0, 0}), pivotIndex([]int{0, 0}) == 0)             // expected: 0

	// =========================================================
	// Nenhum pivot
	// =========================================================

	fmt.Println(pivotIndex([]int{1, 2, 3}), pivotIndex([]int{1, 2, 3}) == -1)       // expected: -1
	fmt.Println(pivotIndex([]int{3, 2, 1}), pivotIndex([]int{3, 2, 1}) == -1)       // expected: -1
	fmt.Println(pivotIndex([]int{1, 1, 1, 1}), pivotIndex([]int{1, 1, 1, 1}) == -1) // expected: -1

	// =========================================================
	// Numeros negativos
	// =========================================================

	fmt.Println(pivotIndex([]int{-1, -1, 0, -1, -1}), pivotIndex([]int{-1, -1, 0, -1, -1}) == 2) // expected: 2
	fmt.Println(pivotIndex([]int{-1, 2, -1}), pivotIndex([]int{-1, 2, -1}) == 1)                 // expected: 1
	fmt.Println(pivotIndex([]int{-1, -2, -3}), pivotIndex([]int{-1, -2, -3}) == -1)              // expected: -1

	// =========================================================
	// Soma total zero
	// =========================================================

	fmt.Println(pivotIndex([]int{1, -1, 1, -1}), pivotIndex([]int{1, -1, 1, -1}) == -1) // expected: -1
	fmt.Println(pivotIndex([]int{1, -1, 0}), pivotIndex([]int{1, -1, 0}) == 2)          // expected: 2

	// =========================================================
	// Pivot com valor impar
	// Importante para testar a divisao inteira por 2
	// =========================================================

	fmt.Println(pivotIndex([]int{1, 2, 7, 1, 2}), pivotIndex([]int{1, 2, 7, 1, 2}) == 2) // expected: 2

	// =========================================================
	// Pivot no meio
	// =========================================================

	fmt.Println(pivotIndex([]int{1, 2, 3, 6, 5, 1}), pivotIndex([]int{1, 2, 3, 6, 5, 1}) == 3) // expected: 3
	fmt.Println(pivotIndex([]int{1, 2, 3, 7, 6, 0}), pivotIndex([]int{1, 2, 3, 7, 6, 0}) == 3) // expected: 3

	// =========================================================
	// Valores máximos permitidos pelo problema
	// =========================================================

	fmt.Println(pivotIndex([]int{1000, 1000, 1000}), pivotIndex([]int{1000, 1000, 1000}) == 1)       // expected: 1
	fmt.Println(pivotIndex([]int{-1000, -1000, -1000}), pivotIndex([]int{-1000, -1000, -1000}) == 1) // expected: 1

	// =========================================================
	// Valores extremos combinados
	// =========================================================

	fmt.Println(pivotIndex([]int{1000, -1000, 1000}), pivotIndex([]int{1000, -1000, 1000}) == 0)   // Mas 1 tbm é pivô
	fmt.Println(pivotIndex([]int{-1000, 1000, -1000}), pivotIndex([]int{-1000, 1000, -1000}) == 0) // Mas 1 tbm é pivô
}
