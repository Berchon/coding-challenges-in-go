package main

/*
A Very Big Sum

https://www.hackerrank.com/challenges/a-very-big-sum/problem?isFullScreen=true

In this challenge, you need to calculate and print the sum of elements
in an array, considering that some integers may be very large.

Function Description

Complete the function with the following parameter(s):
- ar: an array of integers

Return
- the sum of the array elements

Input Format

The first line of the input consists of an integer n.
The next line contains space-separated integers contained in the array.

Output Format

Return the integer sum of the elements in the array.

Constraints

Sample Input

5
1000000001 1000000002 1000000003 1000000004 1000000005

Sample Output

5000000015

Note

The range of the 32-bit integer is limited.
When we add several integer values, the resulting sum might exceed
the above range. You might need to use long int in C/C++/Java
to store such sums.
*/

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	var sum int64 = 0

	for _, value := range ar {
		sum += value
	}

	return sum
}

func main() {
	ar := []int64{
		1000000001,
		1000000002,
		1000000003,
		1000000004,
		1000000005,
	}

	result := aVeryBigSum(ar)

	fmt.Println(result)
}
