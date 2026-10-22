package main

import "fmt"

/*
Birthday Cake Candles

Link:
https://www.hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true

Problem Description:

You are in charge of the cake for a child's birthday. It will have one
candle for each year of their total age. They will only be able to blow
out the tallest of the candles. Your task is to count how many candles
are the tallest.

Example:

The tallest candles are 4 units high. There are 2 candles with this
height, so the function should return 2.

Function Description:

Complete the function with the following parameter(s):

- candles: the candle heights

Returns:

- The number of candles that are tallest.

Input Format:

The first line contains a single integer, n, the size of candles.

The second line contains space-separated integers, where each integer
describes the height of candles[i].

Constraints:

- 1 <= n <= 10^5
- 1 <= candles[i] <= 10^7

Sample Input 0:

4
3 2 1 3

Sample Output 0:

2

Explanation 0:

Candle heights are [3, 2, 1, 3]. The tallest candles are 3 units high,
and there are 2 of them.
*/

// Minha solução
func birthdayCakeCandles(candles []int32) int32 {
	maxHeight := candles[0]
	count := int32(1)

	for _, height := range candles[1:] {
		if height > maxHeight {
			maxHeight = height
			count = 1
		} else if height == maxHeight {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{1, 2, 3, 4, 5}))
}
