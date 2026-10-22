package main

import "fmt"

/*
Nome do problema:
Number Line Jumps

Link:
https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true

Descrição:

Você está preparando um espetáculo de circo com diversos animais. Em uma
das apresentações, há dois cangurus posicionados em uma linha numérica,
prontos para saltar na direção positiva, ou seja, em direção ao infinito
positivo.

O primeiro canguru começa na posição x1 e se desloca v1 metros a cada salto.

O segundo canguru começa na posição x2 e se desloca v2 metros a cada salto.

É necessário determinar se existe uma quantidade de saltos que faça com
que os dois cangurus estejam na mesma posição ao mesmo tempo.

Se for possível, a função deve retornar "YES". Caso contrário, deve
retornar "NO".

Exemplo:

x1 = 0
v1 = 3
x2 = 4
v2 = 2

Depois de um salto:

- Primeiro canguru: 0 + 3 = 3
- Segundo canguru: 4 + 2 = 6

Depois de dois saltos:

- Primeiro canguru: 0 + (2 × 3) = 6
- Segundo canguru: 4 + (2 × 2) = 8

Para os valores do exemplo acima, os cangurus se encontram na mesma
posição após 4? [A condição deve ser determinada pela equação geral.]

A posição de cada canguru após n saltos pode ser representada por:

    primeiro canguru = x1 + n × v1
    segundo canguru  = x2 + n × v2

Para que eles estejam na mesma posição:

    x1 + n × v1 = x2 + n × v2

A função deve determinar se existe um número inteiro e não negativo de
saltos que satisfaça essa igualdade.

Função:

Complete a função kangaroo.

Parâmetros:

- int32 x1: posição inicial do primeiro canguru
- int32 v1: distância percorrida pelo primeiro canguru a cada salto
- int32 x2: posição inicial do segundo canguru
- int32 v2: distância percorrida pelo segundo canguru a cada salto

Retorno:

- string: "YES" se os cangurus puderem se encontrar na mesma posição
  após a mesma quantidade de saltos.
- string: "NO" caso contrário.

Formato de entrada:

Uma única linha contendo quatro inteiros separados por espaços,
representando, respectivamente:

x1 v1 x2 v2

Restrições:

- 0 ≤ x1 < x2 ≤ 10000
- 1 ≤ v1 ≤ 10000
- 1 ≤ v2 ≤ 10000

Exemplo de entrada 0:

0 3 4 2

Exemplo de saída 0:

YES

Explicação:

Os dois cangurus saltam seguindo a sequência de posições determinada
pelas suas posições iniciais e respectivas velocidades.

O primeiro começa em 0 e avança 3 unidades por salto.
O segundo começa em 4 e avança 2 unidades por salto.

Após 4? saltos, ambos ocupam a mesma posição ao mesmo tempo, portanto
a resposta é "YES".

Exemplo de entrada 1:

0 2 5 3

Exemplo de saída 1:

NO

Explicação:

O segundo canguru começa à frente do primeiro:

    x2 > x1

Além disso, ele também se desloca mais rapidamente:

    v2 > v1

Como o segundo canguru já está à frente e continua aumentando sua
vantagem a cada salto, o primeiro nunca conseguirá alcançá-lo.

Portanto, a resposta é "NO".
*/

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 == x2 {
		return "YES"
	}

	if v1 == v2 {
		return "NO"
	}

	positionDifference := x2 - x1
	velocityDifference := v1 - v2

	isIntegerNumberOfJumps := positionDifference%velocityDifference == 0
	isPositiveNumberOfJumps := positionDifference/velocityDifference > 0

	if isIntegerNumberOfJumps && isPositiveNumberOfJumps {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2), kangaroo(0, 3, 4, 2) == "YES") // expected: YES
	fmt.Println(kangaroo(0, 2, 5, 3), kangaroo(0, 2, 5, 3) == "NO")  // expected: NO
}
