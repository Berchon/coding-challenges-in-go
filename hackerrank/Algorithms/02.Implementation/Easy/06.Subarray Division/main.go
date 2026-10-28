package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	windowSize := int(m)
	
	if windowSize <= 0 || len(s) < windowSize {
		return 0
	}
	
	windowSum := sumWindow(s[:windowSize])
	targetSum := d

	var validSegments int32
	if windowSum == targetSum {
		validSegments++
	}

	for i := windowSize; i < len(s); i++ {
		windowSum += s[i] - s[i-windowSize]

		if windowSum == targetSum {
			validSegments++
		}
	}

	return validSegments
}

func sumWindow(window []int32) int32 {
	var sum int32

	for _, value := range window {
		sum += value
	}

	return sum
}

func main() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))    // Output: 2
	fmt.Println(birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 2)) // Output: 0
	fmt.Println(birthday([]int32{4}, 4, 1))                // Output: 1
}
