package main

import "fmt"

/*
# 283. Move Zeroes
# Easy

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]

Constraints:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/

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
