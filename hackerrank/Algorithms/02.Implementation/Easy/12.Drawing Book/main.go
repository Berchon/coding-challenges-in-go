package main

import "fmt"

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
	pagesFromStart := p / 2
	pagesFromEnd := n/2 - p/2

	return min(pagesFromStart, pagesFromEnd)
}

func main() {
	fmt.Println("Caminho feliz")

	fmt.Println(pageCount(6, 2), pageCount(6, 2) == 1)   // Esperado: 1, true
	fmt.Println(pageCount(5, 4), pageCount(5, 4) == 0)   // Esperado: 0, true
	fmt.Println(pageCount(5, 3), pageCount(5, 3) == 1)   // Esperado: 1, true
	fmt.Println(pageCount(10, 6), pageCount(10, 6) == 2) // Esperado: 2, true
	fmt.Println(pageCount(7, 3), pageCount(7, 3) == 1)   // Esperado: 1, true

	fmt.Println("\nCasos de borda")

	fmt.Println(pageCount(1, 1), pageCount(1, 1) == 0) // Esperado: 0, true
	fmt.Println(pageCount(2, 1), pageCount(2, 1) == 0) // Esperado: 0, true
	fmt.Println(pageCount(6, 6), pageCount(6, 6) == 0) // Esperado: 0, true
	fmt.Println(pageCount(6, 5), pageCount(6, 5) == 1) // Esperado: 1, true
	fmt.Println(pageCount(7, 4), pageCount(7, 4) == 1) // Esperado: 1, true

	fmt.Println("\nCasos de n=8")

	fmt.Println(pageCount(8, 1), pageCount(8, 1) == 0) // Esperado: 0, true
	fmt.Println(pageCount(8, 2), pageCount(8, 2) == 1) // Esperado: 1, true
	fmt.Println(pageCount(8, 3), pageCount(8, 3) == 1) // Esperado: 1, true
	fmt.Println(pageCount(8, 4), pageCount(8, 4) == 2) // Esperado: 2, true
	fmt.Println(pageCount(8, 5), pageCount(8, 5) == 2) // Esperado: 2, true
	fmt.Println(pageCount(8, 6), pageCount(8, 6) == 1) // Esperado: 1, true
	fmt.Println(pageCount(8, 7), pageCount(8, 7) == 1) // Esperado: 1, true
	fmt.Println(pageCount(8, 8), pageCount(8, 8) == 0) // Esperado: 0, true

	fmt.Println("\nCasos de n=9")

	fmt.Println(pageCount(9, 1), pageCount(9, 1) == 0) // Esperado: 0, true
	fmt.Println(pageCount(9, 2), pageCount(9, 2) == 1) // Esperado: 1, true
	fmt.Println(pageCount(9, 3), pageCount(9, 3) == 1) // Esperado: 1, true
	fmt.Println(pageCount(9, 4), pageCount(9, 4) == 2) // Esperado: 2, true
	fmt.Println(pageCount(9, 5), pageCount(9, 5) == 2) // Esperado: 2, true
	fmt.Println(pageCount(9, 6), pageCount(9, 6) == 1) // Esperado: 1, true
	fmt.Println(pageCount(9, 7), pageCount(9, 7) == 1) // Esperado: 1, true
	fmt.Println(pageCount(9, 8), pageCount(9, 8) == 0) // Esperado: 0, true
}
