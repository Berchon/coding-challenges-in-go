# Diagonal Difference

**Link do Desafio:** [https://www.hackerrank.com/challenges/diagonal-difference/problem?isFullScreen=true](https://www.hackerrank.com/challenges/diagonal-difference/problem?isFullScreen=true)

---

## Problem Description

Given a square matrix, calculate the absolute difference between the
sums of its diagonals.

For example, the square matrix is:

1 2 3
4 5 6
9 8 9

The left-to-right diagonal is the primary diagonal.
The right-to-left diagonal is the secondary diagonal.

Their absolute difference is the result.

## Function Description

Complete the function with the following parameter:

- arr: a 2-D array of integers

## Return

- The absolute difference in sums along the diagonals.

## Input Format

The first line contains a single integer, n, the number of rows and
columns in the square matrix.

Each of the next n lines describes a row and consists of
space-separated integers.

## Sample Input

3
11 2 4
4 5 6
10 8 -12

## Sample Output

15

## Explanation

The primary diagonal is:

11
   5
     -12

Sum across the primary diagonal:

11 + 5 - 12 = 4

The secondary diagonal is:

     4
   5
10

Sum across the secondary diagonal:

4 + 5 + 10 = 19

Difference:

|4 - 19| = 15

Note: |x| is the absolute value of x.
