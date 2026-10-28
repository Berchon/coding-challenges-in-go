package main

import "fmt"

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
