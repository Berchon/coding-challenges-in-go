package main

import "fmt"

/*
Divisible Sum Pairs

Link:
https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true

Problem Description:

You are given an array of integers, ar, and a positive integer, k.
Find and print the number of (i, j) pairs where i < j
and ar[i] + ar[j] is evenly divisible by k.

Function Description:

Complete the divisibleSumPairs function.

Parameters:
- int32 n: the length of the array.
- int32 k: the integer divisor.
- []int32 ar: an array of integers.

Returns:
- int32: the number of valid pairs.

Input Format:

The first line contains two space-separated integers, n and k,
the length of the array and the integer divisor, respectively.

The second line contains n space-separated integers,
each a value of ar.

Constraints:
- 2 <= n <= 100
- 1 <= k <= 100
- 1 <= ar[i] <= 100

Sample Input:
6 3
1 3 2 6 1 2

Sample Output:
5

Explanation:

The five valid pairs are:

(0, 2): ar[0] + ar[2] = 1 + 2 = 3
(0, 5): ar[0] + ar[5] = 1 + 2 = 3
(1, 3): ar[1] + ar[3] = 3 + 6 = 9
(2, 4): ar[2] + ar[4] = 2 + 1 = 3
(4, 5): ar[4] + ar[5] = 1 + 2 = 3
*/

// func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
// 	remainderCounts := make([]int32, int(k))
// 	var pairCount int32

// 	for i := 0; i < int(n); i++ {
// 		remainder := ar[i] % k
// 		complement := (k - remainder) % k

// 		pairCount += remainderCounts[int(complement)]
// 		remainderCounts[int(remainder)]++
// 	}

// 	return pairCount
// }

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var pairCount int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if isDivisibleBy(ar[i]+ar[j], k) {
				pairCount++
			}
		}
	}
	return pairCount
}

func isDivisibleBy(value, divisor int32) bool {
	return value%divisor == 0
}

func main() {
	fmt.Println(divisibleSumPairs(int32(6), int32(3), []int32{1, 3, 2, 6, 1, 2})) // Output: 5
	fmt.Println(divisibleSumPairs(int32(6), int32(5), []int32{1, 2, 3, 4, 5, 6})) // Output: 3
}
