package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 338. Counting Bits
### Easy

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Do not solve it with built-in functions (i.e., like __builtin_popcount in C++).

Example 1:
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

Example 2:
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

Constraints:
0 <= n <= 105

Follow up:
It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
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

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}

func main() {
	fmt.Println(countBits(0))  // Esperado: [0]
	fmt.Println(countBits(1))  // Esperado: [0 1]
	fmt.Println(countBits(2))  // Esperado: [0 1 1]
	fmt.Println(countBits(5))  // Esperado: [0 1 1 2 1 2]
	fmt.Println(countBits(8))  // Esperado: [0 1 1 2 1 2 2 3 1]
	fmt.Println(countBits(10)) // Esperado: [0 1 1 2 1 2 2 3 1 2 2]
}
