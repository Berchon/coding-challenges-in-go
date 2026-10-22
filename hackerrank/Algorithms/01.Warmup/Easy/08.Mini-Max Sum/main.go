package main

import "fmt"

/*
Mini-Max Sum

Link:
https://www.hackerrank.com/challenges/mini-max-sum/problem?isFullScreen=true

Problem Description:

Given five positive integers, find the minimum and maximum values
that can be calculated by summing exactly four of the five integers.
Then print the respective minimum and maximum values as a single
line of two space-separated long integers.

Function Description:

Complete the miniMaxSum function.

miniMaxSum has the following parameter:
- arr: an array of integers

Print:

Print two space-separated integers on one line: the minimum sum
and the maximum sum of four of the five elements.

No value should be returned.

Input Format:

A single line of five space-separated integers.

Output Format:

Print two space-separated long integers denoting the respective
minimum and maximum values that can be calculated by summing
exactly four of the five integers.

The output can be greater than a 32-bit integer.

Explanation:

For the input:

1 2 3 4 5

The possible sums using four of the five integers are:

1. 2 + 3 + 4 + 5 = 14
2. 1 + 3 + 4 + 5 = 13
3. 1 + 2 + 4 + 5 = 12
4. 1 + 2 + 3 + 5 = 11
5. 1 + 2 + 3 + 4 = 10

Therefore:

Minimum sum = 10
Maximum sum = 14

Sample Input:

1 2 3 4 5

Sample Output:

10 14

Hint:

Beware of integer overflow. Use a 64-bit integer to store the sums.
*/

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
