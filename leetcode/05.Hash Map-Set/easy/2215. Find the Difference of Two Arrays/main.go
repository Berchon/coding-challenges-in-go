package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 2215. Find the Difference of Two Arrays
### Easy

Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

Example 1:
Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums1. Therefore, answer[1] = [4,6].

Example 2:
Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

Constraints:
1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000
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

func toSet(nums []int) map[int]struct{} {
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		set[num] = struct{}{}
	}

	return set
}

func difference(setA, setB map[int]struct{}) []int {
	result := make([]int, 0, len(setA))

	for num := range setA {
		if _, exists := setB[num]; !exists {
			result = append(result, num)
		}
	}

	return result
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := toSet(nums1)
	set2 := toSet(nums2)

	return [][]int{
		difference(set1, set2),
		difference(set2, set1),
	}
}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))       // Output: [[1,3],[4,6]]
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})) // Output: [[3],[]]
}
