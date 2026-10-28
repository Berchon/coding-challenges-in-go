package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	applesInHouse, orangesInHouse := 0, 0

	for i := 0; i < max(len(apples), len(oranges)); i++ {
		if i < len(apples) {
			applePosition := a + apples[i]

			if applePosition >= s && applePosition <= t {
				applesInHouse++
			}
		}

		if i < len(oranges) {
			orangePosition := b + oranges[i]

			if orangePosition >= s && orangePosition <= t {
				orangesInHouse++
			}
		}
	}

	fmt.Println(applesInHouse)
	fmt.Println(orangesInHouse)
}
