package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	mmcA := mmcArray(a)
	mdcB := mdcArray(b)

	if mmcA > mdcB {
		return 0
	}

	var count int32

	for x := mmcA; x <= mdcB; x += mmcA {
		if mdcB%x == 0 {
			count++
		}
	}

	return count
}

func mdc(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func mmc(a, b int32) int32 {
	return a / mdc(a, b) * b
}

func mdcArray(values []int32) int32 {
	if len(values) == 0 {
		return 0
	}

	result := values[0]

	for _, value := range values[1:] {
		result = mdc(result, value)
	}

	return result
}

func mmcArray(values []int32) int32 {
	result := values[0]

	for _, value := range values[1:] {
		result = mmc(result, value)
	}

	return result
}

func main() {
	fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}) == 3) // expected: true
	fmt.Println(getTotalX([]int32{3, 4}, []int32{24, 48}) == 2)     // expected: true
}
