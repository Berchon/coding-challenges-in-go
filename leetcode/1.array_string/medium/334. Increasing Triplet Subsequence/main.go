package main

import "fmt"

func increasingTriplet(nums []int) bool {
	values := make([]int, 3)
	values[0], values[1], values[2] = nums[0], nums[0], nums[0]
	// ref := 1
	for _, num := range nums {
		// values[ref] = num
		if num < values[0] {
			values[0], values[1], values[2] = num, num, num
			continue
		}
		if num > values[0] 

		
		
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
