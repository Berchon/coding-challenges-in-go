package main

import "fmt"

/*
Between Two Sets

Link:
https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true

Descrição:

São fornecidos dois arrays de inteiros, a e b.

Precisamos encontrar quantos inteiros existem entre os dois arrays, seguindo
duas condições:

1. Todos os elementos do array a devem ser fatores (divisores) do número
   considerado.

2. O número considerado deve ser fator (divisor) de todos os elementos
   do array b.

Um número só deve ser contado quando satisfizer simultaneamente as duas
condições.

Exemplo:

a = [2, 4]
b = [16, 32, 96]

Os números que satisfazem as duas condições são:

4, 8 e 16

Portanto, o resultado é:

3

Explicação:

Para 4:
- 2 divide 4
- 4 divide 4
- 4 divide 16, 32 e 96

Para 8:
- 2 divide 8
- 4 divide 8
- 8 divide 16, 32 e 96

Para 16:
- 2 divide 16
- 4 divide 16
- 16 divide 16, 32 e 96

Logo, existem 3 números entre os dois conjuntos.

Entrada:

A primeira linha contém a quantidade de elementos dos arrays a e b.
A segunda linha contém os elementos do array a.
A terceira linha contém os elementos do array b.

Saída:

Retornar a quantidade de inteiros que satisfazem as duas condições.

Exemplo de entrada:

2 3
2 4
16 32 96

Exemplo de saída:

3
*/

func getTotalX(a []int32, b []int32) int32 {
	mmcA := mmcArray(a)
	mdcB := mdcArray(b)

	if mmcA > mdcB {
		return 0
	}

	var count int32

	for x := mmcA; x <= mdcB; x += mmcA {
		if mdcB%x == 0 {
			count++
		}
	}

	return count
}

func mdc(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func mmc(a, b int32) int32 {
	return a / mdc(a, b) * b
}

func mdcArray(values []int32) int32 {
	if len(values) == 0 {
		return 0
	}

	result := values[0]

	for _, value := range values[1:] {
		result = mdc(result, value)
	}

	return result
}

func mmcArray(values []int32) int32 {
	result := values[0]

	for _, value := range values[1:] {
		result = mmc(result, value)
	}

	return result
}

func main() {
	fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}) == 3) // expected: true
	fmt.Println(getTotalX([]int32{3, 4}, []int32{24, 48}) == 2)     // expected: true
}
