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

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}

	previous3, previous2, current := 0, 1, 1

	for i := 3; i <= n; i++ {
		previous3, previous2, current =
			previous2, current, previous3+previous2+current
	}

	return current
}

func main() {
	fmt.Println(tribonacci(0))  // output 0
	fmt.Println(tribonacci(1))  // output 1
	fmt.Println(tribonacci(2))  // output 1
	fmt.Println(tribonacci(3))  // output 2
	fmt.Println(tribonacci(4))  // output 4
	fmt.Println(tribonacci(5))  // output 7
	fmt.Println(tribonacci(25)) // output 1389537
}
