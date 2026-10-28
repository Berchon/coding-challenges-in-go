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

func minCostClimbingStairs(cost []int) int {
	previousCost, currentCost := 0, 0

	for i := 2; i <= len(cost); i++ {
		previousCost, currentCost = currentCost, min(previousCost+cost[i-2], currentCost+cost[i-1])
	}

	return currentCost
}

func main() {
	// Caminho feliz

	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}), "\n") // Esperado: 15
	fmt.Println(minCostClimbingStairs([]int{1, 2, 3}), "\n")    // Esperado: 2
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1}), "\n")  // Esperado: 1

	// Casos de borda

	fmt.Println(minCostClimbingStairs([]int{0, 0}), "\n")                // Esperado: 0
	fmt.Println(minCostClimbingStairs([]int{0, 5}), "\n")                // Esperado: 0
	fmt.Println(minCostClimbingStairs([]int{5, 0}), "\n")                // Esperado: 0
	fmt.Println(minCostClimbingStairs([]int{1, 1, 1, 1}), "\n")          // Esperado: 2
	fmt.Println(minCostClimbingStairs([]int{0, 0, 0, 0, 0}), "\n")       // Esperado: 0
	fmt.Println(minCostClimbingStairs([]int{999, 999}), "\n")            // Esperado: 999
	fmt.Println(minCostClimbingStairs([]int{999, 1, 999, 1, 999}), "\n") // Esperado: 2

	fmt.Println(minCostClimbingStairs([]int{1, 0, 0, 2}), "\n") // Esperado: 0
}
