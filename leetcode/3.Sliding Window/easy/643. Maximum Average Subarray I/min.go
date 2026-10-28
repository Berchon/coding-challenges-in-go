package main

import (
	"fmt"
)

// func findMaxAverage(nums []int, k int) float64 {
// 	windowSize := len(nums) - k + 1
// 	sums := make([]float64, windowSize)

// 	//
// 	for i := 0; i < windowSize; i++ {
// 		sum := 0
// 		for window := i; window < i+k && window < len(nums); window++ {
// 			sum += nums[window]
// 		}
// 		fmt.Println(i, sum)
// 		sums[i] = float64(sum) / float64(k)
// 	}

// 	return slices.Max(sums)
// }

func calculateWindowSum(nums []int, start, k int) int {
	sum := 0
	for i := start; i < start+k; i++ {
		sum += nums[i]
	}

	return sum
}

func calculateAverage(sum, windowSize int) float64 {
	return float64(sum) / float64(windowSize)
}

func maxAverageOf(currentMax, candidate float64) float64 {
	if candidate > currentMax {
		return candidate
	}

	return currentMax
}

func slideWindow(nums []int, windowSum, windowStart, k int) int {
	getOutNumber := nums[windowStart-1]
	getInNumber := nums[windowStart+k-1]

	return windowSum - getOutNumber + getInNumber
}

func findMaxAverage(nums []int, k int) float64 {
	windowSum := calculateWindowSum(nums, 0, k)
	maxAverage := calculateAverage(windowSum, k)

	for i := 1; i <= len(nums)-k; i++ {
		windowSum = slideWindow(nums, windowSum, i, k)
		currentAverage := calculateAverage(windowSum, k)
		maxAverage = maxAverageOf(maxAverage, currentAverage)
	}

	return maxAverage
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println(findMaxAverage(nums, k))
}
