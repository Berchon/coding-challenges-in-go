package main

import (
	"fmt"
	"strings"
)

// Minha solução
func staircase(n int32) {
	size := int(n)

	for i := 1; i <= size; i++ {
		fmt.Printf(
			"%s%s\n",
			strings.Repeat(" ", size-i),
			strings.Repeat("#", i),
		)
	}
}

func main() {
	staircase(6)
}