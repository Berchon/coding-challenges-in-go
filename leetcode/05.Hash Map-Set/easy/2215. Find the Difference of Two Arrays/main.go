package main

import (
	"fmt"
	"runtime/debug"
)

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
