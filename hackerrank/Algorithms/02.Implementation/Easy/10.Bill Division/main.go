package main

import "fmt"

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) string {
	var total int32

	for i, price := range bill {
		if int32(i) != k {
			total += price
		}
	}

	annaShouldPay := total / 2

	if annaShouldPay == b {
		fmt.Println("Bon Appetit")
		return "Bon Appetit"
	}

	fmt.Println(b - annaShouldPay)
	return fmt.Sprintf("%d", b-annaShouldPay)
}

func main() {
	fmt.Println(bonAppetit([]int32{3, 10, 2, 9}, 1, 12), bonAppetit([]int32{3, 10, 2, 9}, 1, 12) == "5")
	fmt.Println(bonAppetit([]int32{3, 10, 2, 9}, 1, 7), bonAppetit([]int32{3, 10, 2, 9}, 1, 7) == "Bon Appetit")
	fmt.Println(bonAppetit([]int32{2, 4, 6}, 2, 6), bonAppetit([]int32{2, 4, 6}, 2, 6) == "3")
	fmt.Println(bonAppetit([]int32{2, 4, 6}, 2, 3), bonAppetit([]int32{2, 4, 6}, 2, 3) == "Bon Appetit")
	fmt.Println(bonAppetit([]int32{10, 20}, 0, 20), bonAppetit([]int32{10, 20}, 0, 20) == "10")
	fmt.Println(bonAppetit([]int32{10, 20}, 1, 5), bonAppetit([]int32{10, 20}, 1, 5) == "Bon Appetit")
	fmt.Println(bonAppetit([]int32{1, 2, 3, 4, 5}, 0, 8), bonAppetit([]int32{1, 2, 3, 4, 5}, 0, 8) == "1")
	fmt.Println(bonAppetit([]int32{1, 2, 3, 4, 5}, 4, 5), bonAppetit([]int32{1, 2, 3, 4, 5}, 4, 5) == "0")
	fmt.Println(bonAppetit([]int32{100, 200, 300}, 1, 200), bonAppetit([]int32{100, 200, 300}, 1, 200) == "0")
	fmt.Println(bonAppetit([]int32{5, 5, 5, 5}, 2, 8), bonAppetit([]int32{5, 5, 5, 5}, 2, 8) == "1")
}
