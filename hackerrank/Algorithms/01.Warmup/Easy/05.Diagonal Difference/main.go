package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var difference int32 = 0

	for i, line := range arr {
		difference += line[i] - line[len(line)-1-i]
	}

	return abs(difference)
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	result := diagonalDifference(arr)

	fmt.Println(result)
}
