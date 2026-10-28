package main

import "fmt"

func moveZeroes2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	nextNonZero := 0

	for current := 0; current < len(nums); current++ {
		if nums[current] == 0 {
			continue
		}

		if current != nextNonZero {
			nums[nextNonZero], nums[current] = nums[current], nums[nextNonZero]
		}

		nextNonZero++
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
