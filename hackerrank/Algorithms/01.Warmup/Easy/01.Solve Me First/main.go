package main

/*
Problema: Solve Me First

Link:
https://www.hackerrank.com/challenges/solve-me-first/problem

Resumo do problema:
Complete a função `solveMeFirst` para retornar a soma de dois inteiros.

Parâmetros:
- a: primeiro inteiro
- b: segundo inteiro

Retorno:
- A soma de `a` e `b`

Exemplo:
Entrada:
a = 2
b = 3

Saída:
5

Observação:
O objetivo do desafio é praticar leitura de entrada, chamada de função
e escrita da saída no HackerRank.
*/

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
