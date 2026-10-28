package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 == x2 {
		return "YES"
	}

	if v1 == v2 {
		return "NO"
	}

	positionDifference := x2 - x1
	velocityDifference := v1 - v2

	isIntegerNumberOfJumps := positionDifference%velocityDifference == 0
	isPositiveNumberOfJumps := positionDifference/velocityDifference > 0

	if isIntegerNumberOfJumps && isPositiveNumberOfJumps {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2), kangaroo(0, 3, 4, 2) == "YES") // expected: YES
	fmt.Println(kangaroo(0, 2, 5, 3), kangaroo(0, 2, 5, 3) == "NO")  // expected: NO
}
