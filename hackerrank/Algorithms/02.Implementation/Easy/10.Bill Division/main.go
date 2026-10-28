package main

import "fmt"

/*

# Bon Appétit

Two friends Anna and Brian, are deciding how to split the bill at a dinner. Each will only pay for the items they consume. Brian gets the check and calculates Anna's portion. You must determine if his calculation is correct.

For example, assume the bill has the following prices: `bill = [2, 4, 6]`. Anna declines to eat item `k = 2`, which costs `bill[2] = 6`. If Brian calculates the bill correctly, Anna will pay `(2 + 4) / 2 = 3`. If he includes the cost of `bill[2]`, he will calculate `(2 + 4 + 6) / 2 = 6`. In the second case, he should refund `6 - 3 = 3` to Anna.

## Function Description

Complete the `bonAppetit` function in the editor below. It should print `Bon Appetit` if the bill is fairly split. Otherwise, it should print the integer amount of money that Brian owes Anna.

`bonAppetit` has the following parameter(s):

- `bill`: an array of integers representing the cost of each item ordered
- `k`: an integer representing the zero-based index of the item Anna doesn't eat
- `b`: the amount of money that Anna contributed to the bill

## Input Format

The first line contains two space-separated integers `n` and `k`, the number of items ordered and the 0-based index of the item that Anna did not eat.

The second line contains `n` space-separated integers `bill[i]` where `0 <= i < n`.

The third line contains an integer, `b`, the amount of money that Brian charged Anna for her share of the bill.

## Constraints

- `2 <= n <= 10^5`
- `0 <= k < n`
- `0 <= bill[i] <= 10^4`
- `0 <= b <= 10^4`

The amount of money due Anna will always be an integer.

## Output Format

If Brian did not overcharge Anna, print `Bon Appetit` on a new line; otherwise, print the difference (i.e., `b_charged - b_actual`) that Brian must refund to Anna. This will always be an integer.

## Sample Input 0

```text
4 1
3 10 2 9
12
```

## Sample Output 0

```text
5
```

## Explanation 0

Anna didn't eat item `bill[1] = 10`, but she shared the rest of the items with Brian. The total cost of the shared items is `3 + 2 + 9 = 14` and, split in half, the cost per person is `b_actual = 14 / 2 = 7`. Brian charged her `b_charged = 12` for her portion of the bill. We print the amount Anna was overcharged, `b_charged - b_actual = 12 - 7 = 5`, on a new line.

## Sample Input 1

```text
4 1
3 10 2 9
7
```

## Sample Output 1

```text
Bon Appetit
```

## Explanation 1

Anna didn't eat item `bill[1] = 10`, but she shared the rest of the items with Brian. The total cost of the shared items is `3 + 2 + 9 = 14` and, split in half, the cost per person is `b_actual = 14 / 2 = 7`. Because `b_actual = b_charged = 7`, we print `Bon Appetit` on a new line.
*/

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
