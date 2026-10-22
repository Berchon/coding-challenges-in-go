package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 746. Min Cost Climbing Stairs

**Link:** https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=leetcode-75

**Difficulty:** Easy

## Description

You are given an integer array `cost` where `cost[i]` is the cost of `i`th step on a staircase.

Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the staircase, which is the position just past the last step (index `cost.length`).

## Examples

### Example 1

**Input:**

```text
cost = [10,15(*),20]
```

**Output:**

```text
15
```

**Explanation:**

You will start at index 1.

* Pay 15 and climb two steps to reach the top.

The total cost is 15.

### Example 2

**Input:**

```text
cost = [1(*), 100, 1(*), 1(*), 1(*), 100, 1(*), 1(*), 100, 1(*)]
```

**Output:**

```text
6
```

**Explanation:**

You will start at index 0.

* Pay 1 and climb two steps to reach index 2.
* Pay 1 and climb two steps to reach index 4.
* Pay 1 and climb two steps to reach index 6.
* Pay 1 and climb one step to reach index 7.
* Pay 1 and climb two steps to reach index 9.
* Pay 1 and climb one step to reach the top.

The total cost is 6.

## Constraints

* `2 <= cost.length <= 1000`
* `0 <= cost[i] <= 999`
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
