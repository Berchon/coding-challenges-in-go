package main

import "fmt"

/*
Nome do problema:
Breaking the Records

Link:
https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=true

Descrição:
Maria joga basquete e, durante uma temporada, registra a pontuação
obtida em cada partida.

A pontuação da primeira partida estabelece, simultaneamente, o melhor
e o pior recorde da temporada.

Para cada partida seguinte:

* Se a pontuação for estritamente maior que o melhor recorde atual,
  o melhor recorde é quebrado.
* Se a pontuação for estritamente menor que o pior recorde atual,
  o pior recorde é quebrado.
* Uma pontuação igual ao recorde atual não o quebra.

Dado um array contendo as pontuações das partidas na ordem em que
foram disputadas, determine quantas vezes Maria quebrou o recorde
de maior pontuação e quantas vezes quebrou o recorde de menor
pontuação.

Retorno:
Um array com dois elementos:

* índice 0: quantidade de vezes que o recorde de maior pontuação
  foi quebrado.
* índice 1: quantidade de vezes que o recorde de menor pontuação
  foi quebrado.

Exemplo:

scores = [10, 5, 20, 20, 4, 5, 2, 25, 1]

Resultado:
[2, 4]

Explicação resumida:

* 10 estabelece o melhor e o pior recorde.
* 5 quebra o pior.
* 20 quebra o melhor.
* 20 não quebra nenhum recorde, pois é igual ao melhor.
* 4 quebra o pior.
* 5 não quebra nenhum recorde.
* 2 quebra o pior.
* 25 quebra o melhor.
* 1 quebra o pior.
*/

func breakingRecords(scores []int32) []int32 {
	best, worst  := scores[0], scores[0]
	bestBreaks, worstBreaks := int32(0), int32(0)

	for _, score := range scores[1:] {
		if score > best {
			best = score
			bestBreaks++
		}

		if score < worst {
			worst = score
			worstBreaks++
		}
	}

	return []int32{bestBreaks, worstBreaks}
}

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}

	result := breakingRecords(scores)

	fmt.Println(result)
	// Resultado esperado: [2 4]
}
