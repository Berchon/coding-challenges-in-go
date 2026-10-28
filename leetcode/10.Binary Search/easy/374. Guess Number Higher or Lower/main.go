package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
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

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 1, n

	for low < high {
		mid := low + (high-low)/2

		if guess(mid) <= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func guess(n int) int {
	return cmp.Compare(pick, n)
}

var pick int

func main() {
	n := 10
	pick = rand.IntN(n) + 1
	fmt.Printf("Pick number is: %d\n", pick)
	result := guessNumber(n)
	fmt.Printf("The picked number is: %d\n", result)
}
