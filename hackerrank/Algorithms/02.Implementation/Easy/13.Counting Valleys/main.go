package main

import "fmt"

/*
Counting Valleys
Link: https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true

Descrição do problema:

Gary é um caminhante que registra cuidadosamente suas caminhadas.
Durante sua última caminhada, ele deu exatamente `steps` passos.
Para cada passo, registrou se foi uma subida (U) ou uma descida (D).

A caminhada sempre começa e termina no nível do mar. Cada passo
para cima ou para baixo representa uma mudança de uma unidade
na altitude.

Definições:

- Montanha: sequência de passos consecutivos acima do nível do mar,
  começando com uma subida a partir do nível do mar e terminando
  com uma descida de volta ao nível do mar.

- Vale: sequência de passos consecutivos abaixo do nível do mar,
  começando com uma descida a partir do nível do mar e terminando
  com uma subida de volta ao nível do mar.

Dada a sequência de subidas e descidas, determine quantos vales
foram percorridos durante a caminhada.

Descrição da função:

Complete a função countingValleys.

Parâmetros:
- steps (int32): quantidade de passos da caminhada.
- path (string): sequência de caracteres U e D que descreve
  os passos da caminhada.

Retorno:
- int32: quantidade de vales percorridos.

Formato de entrada:
- Primeira linha: inteiro steps, representando a quantidade de passos.
- Segunda linha: string path, contendo os passos da caminhada.

Restrições:
- 2 <= steps <= 10^6
- path[i] pertence ao conjunto {U, D}

Exemplo de entrada:
8
UDDDUDUU

Exemplo de saída:
1

Explicação:
A caminhada começa no nível do mar, desce abaixo dele, retorna ao
nível do mar e depois passa por uma montanha. O percurso contém
exatamente um vale.
*/

// Complete a função sem implementar a lógica.
func countingValleys(steps int32, path string) int32 {
	currentAltitude, valleyCount := int32(0), int32(0)

	for i := 0; i < len(path); i++ {
		previousAltitude := currentAltitude
		currentAltitude += altitudeChange(path[i])
		if currentAltitude == 0 && previousAltitude < 0 {
			valleyCount++
		}
	}
	return valleyCount
}

func altitudeChange(step byte) int32 {
	switch step {
	case 'U', 'u':
		return 1
	case 'D', 'd':
		return -1
	default:
		return 0
	}
}

func main() {
	fmt.Println("Caminho feliz:")
	fmt.Println("Teste 01 - Exemplo oficial:",
		countingValleys(8, "UDDDUDUU"),
		countingValleys(8, "UDDDUDUU") == 1) // Esperado: 1, true

	fmt.Println("Teste 02 - Um vale simples:",
		countingValleys(4, "DDUU"),
		countingValleys(4, "DDUU") == 1) // Esperado: 1, true

	fmt.Println("Teste 03 - Uma montanha:",
		countingValleys(4, "UUDD"),
		countingValleys(4, "UUDD") == 0) // Esperado: 0, true

	fmt.Println("Teste 04 - Dois vales:",
		countingValleys(8, "DDUUDDUU"),
		countingValleys(8, "DDUUDDUU") == 2) // Esperado: 2, true

	fmt.Println("Teste 05 - Vale e montanha:",
		countingValleys(8, "DDUUUUDD"),
		countingValleys(8, "DDUUUUDD") == 1) // Esperado: 1, true

	fmt.Println("Teste 06 - Caminhada alternada:",
		countingValleys(4, "DUDU"),
		countingValleys(4, "DUDU") == 2) // Esperado: 2, true

	fmt.Println("Edge cases:")
	fmt.Println("Teste 07 - Apenas um vale raso:",
		countingValleys(2, "DU"),
		countingValleys(2, "DU") == 1) // Esperado: 1, true

	fmt.Println("Teste 08 - Caminhada vazia (fora das restrições):",
		countingValleys(0, ""),
		countingValleys(0, "") == 0) // Esperado adicional: 0, true

	fmt.Println("Teste 09 - Apenas descidas (percurso incompleto):",
		countingValleys(4, "DDDD"),
		countingValleys(4, "DDDD") == 0) // Esperado adicional: 0, true

	fmt.Println("Teste 10 - Apenas subidas (percurso incompleto):",
		countingValleys(4, "UUUU"),
		countingValleys(4, "UUUU") == 0) // Esperado adicional: 0, true

	// Teste 11 - Cinco vales e três montanhas (32 passos)
	path11 := "DDUU" + "DDUU" + "DDUU" + "DDUU" + "DDUU" +
		"UUDD" + "UUDD" + "UUDD"

	fmt.Println(countingValleys(32, path11), countingValleys(32, path11) == 5) // Esperado: 5, true

	// Teste 12 - Três vales e cinco montanhas (32 passos)
	path12 := "DDUU" + "DDUU" + "DDUU" +
		"UUDD" + "UUDD" + "UUDD" + "UUDD" + "UUDD"

	fmt.Println(countingValleys(32, path12), countingValleys(32, path12) == 3) // Esperado: 3, true
}
