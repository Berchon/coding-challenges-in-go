package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	// Exemplo de execução local
	a := uint32(2)
	b := uint32(3)

	result := solveMeFirst(a, b)

	fmt.Printf("solveMeFirst(%d, %d) = %d\n", a, b, result)
}
