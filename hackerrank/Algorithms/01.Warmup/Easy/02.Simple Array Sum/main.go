package main

/*
Simple Array Sum

https://www.hackerrank.com/challenges/simple-array-sum/problem?isFullScreen=true

Problem Description:

Given an array of integers, find the sum of its elements.

For example, if the array ar = [1, 2, 3],
1 + 2 + 3 = 6, so return 6.

Function Description:

Complete the simpleArraySum function in the editor below.
It must return the sum of the array elements as an integer.

simpleArraySum has the following parameter(s):

- ar: an array of integers

Input Format:

The first line contains an integer, n, denoting the size of the array.
The second line contains n space-separated integers representing the
array's elements.

Output Format:

Print the sum of the array's elements as a single integer.

Constraints:

The HackerRank page currently does not expose the constraint values
in its rendered problem statement.

Sample Input:

6
1 2 3 4 10 11

Sample Output:

31

Explanation:

Print the sum of the array's elements:

1 + 2 + 3 + 4 + 10 + 11 = 31
*/

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0

	for _, element := range ar {
		sum += element
	}

	return sum
}

func main() {
	ar := []int32{1, 2, 3, 4, 10, 11}

	result := simpleArraySum(ar)

	println(result)
}
