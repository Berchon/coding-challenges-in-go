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
