package main

import "fmt"

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countOccurrences(arr []int32) []int32 {
	counts := make([]int32, 5)

	for _, id := range arr {
		counts[id-1]++
	}

	return counts
}

func findMostFrequentTypeIndex(counts []int32) int32 {
	maxValueIndex := 0

	for index, count := range counts {
		if count > counts[maxValueIndex] {
			maxValueIndex = index
		}
	}

	return int32(maxValueIndex)
}

func migratoryBirds(arr []int32) int32 {
	counts := countOccurrences(arr)
	maxValueIndex := findMostFrequentTypeIndex(counts)

	return maxValueIndex + 1
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}), migratoryBirds([]int32{1, 4, 4, 4, 5, 3}) == 4)                               // Esperado: 4
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}), migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}) == 3) // Esperado: 3
	fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}), migratoryBirds([]int32{1, 1, 2, 2, 3}) == 1)                                     // Esperado: 1
	fmt.Println(migratoryBirds([]int32{5, 5, 5, 1, 1, 1, 3}), migratoryBirds([]int32{5, 5, 5, 1, 1, 1, 3}) == 1)                         // Esperado: 1
	fmt.Println(migratoryBirds([]int32{2, 2, 2, 3, 3, 3, 1}), migratoryBirds([]int32{2, 2, 2, 3, 3, 3, 1}) == 2)                         // Esperado: 2
}
