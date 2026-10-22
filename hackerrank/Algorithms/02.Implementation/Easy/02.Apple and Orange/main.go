package main

import "fmt"

/*
Apple and Orange

Link:
https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true

Sam's house has an apple tree and an orange tree that yield an abundance
of fruit. Using the information given below, determine the number of
apples and oranges that land on Sam's house.

Sam's house is located on the inclusive range [s, t].

The apple tree is located at point a, and the orange tree is located
at point b.

When a fruit falls from its tree, it lands d units of distance from
its tree of origin along the x-axis.

A negative value of d means the fruit fell d units to the tree's left,
and a positive value of d means it falls d units to the tree's right.

Given the values of d for each apple and orange, determine how many
apples and oranges will fall on Sam's house.

Function Description

Complete the countApplesAndOranges function in the editor below.
It should print the number of apples and oranges that land on Sam's
house, each on a separate line.

countApplesAndOranges has the following parameters:

- s: integer, starting point of Sam's house location.
- t: integer, ending location of Sam's house location.
- a: integer, location of the Apple tree.
- b: integer, location of the Orange tree.
- apples: integer array, distances at which each apple falls from the tree.
- oranges: integer array, distances at which each orange falls from the tree.

Input Format

The first line contains two space-separated integers denoting the
respective values of s and t.

The second line contains two space-separated integers denoting the
respective values of a and b.

The third line contains two space-separated integers denoting the
respective values of m and n.

The fourth line contains space-separated integers denoting the
respective distances that each apple falls from point a.

The fifth line contains space-separated integers denoting the
respective distances that each orange falls from point b.

Output Format

Print two integers on two different lines:

1. The first integer: the number of apples that fall on Sam's house.
2. The second integer: the number of oranges that fall on Sam's house.

Sample Input:

7 11
5 15
3 2
-2 2 1
5 -6

Sample Output:

1
1

Explanation:

The first apple falls at position 3.
The second apple falls at position 7.
The third apple falls at position 6.

The first orange falls at position 20.
The second orange falls at position 9.

Only one apple and one orange fall within the region between 7 and 11.
*/

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
