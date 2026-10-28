package main

import "fmt"

func miniMaxSum(arr []int32) {
	sum := int64(0)
	min := int64(arr[0])
	max := int64(arr[0])

	for _, element := range arr {
		value := int64(element)
		sum += value

		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	fmt.Printf("%d %d", sum-max, sum-min)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}

	miniMaxSum(arr)
}
