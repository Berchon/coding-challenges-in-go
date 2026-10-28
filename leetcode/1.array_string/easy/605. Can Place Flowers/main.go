package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] != 0 {
			continue
		}

		emptyLeft := i == 0 || flowerbed[i-1] == 0
		emptyRight := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if emptyLeft && emptyRight {
			flowerbed[i] = 1
			n--
		}
	}

	return n == 0
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1

	fmt.Println(canPlaceFlowers(flowerbed, n))
}
