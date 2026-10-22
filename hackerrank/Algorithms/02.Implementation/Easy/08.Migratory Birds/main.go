package main

import "fmt"

/*
# Migratory Birds

**Plataforma:** [HackerRank — Migratory Birds](https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true)

**Dificuldade:** Easy

## Problem Description

You have been asked to help study the population of birds migrating across the continent. Each type of bird you are interested in will be identified by an integer value. Each time a particular kind of bird is spotted, its ID number will be added to your array of sightings.

You would like to find out which type of bird is most common given a list of sightings. Your task is to print the type number of that bird and, if two or more types of birds are equally common, choose the type with the smallest ID number.

For example, suppose the bird sightings are represented by the array:

`arr = [1, 1, 2, 2, 3]`

There are two sightings each of types `1` and `2`, and one sighting of type `3`. Since types `1` and `2` have the same highest frequency, choose the smaller ID: `1`.

## Function Description

Complete the `migratoryBirds` function.

### Parameters

- `int arr[n]`: an array of integers representing the types of birds sighted.

### Returns

- `int`: the lowest type ID among the most frequently sighted bird types.

## Input Format

- The first line contains an integer `n`, the size of the array `arr`.
- The second line contains `n` space-separated integers, each representing the type ID of a bird sighted.

## Constraints

- \(5 \leq n \leq 2 \times 10^5\)
- Each bird type ID is one of the following integers: `1`, `2`, `3`, `4`, or `5`.

## Sample Input 0

```text
6
1 4 4 4 5 3
```

## Sample Output 0

```text
4
```

## Explanation 0

The frequencies of the different bird types are:

| Bird type ID | Frequency |
|---|---:|
| 1 | 1 |
| 2 | 0 |
| 3 | 1 |
| 4 | 3 |
| 5 | 1 |

Bird type `4` has the highest frequency, appearing three times. Therefore, the answer is `4`.

## Sample Input 1

```text
11
1 2 3 4 5 4 3 2 1 3 4
```

## Sample Output 1

```text
3
```

## Explanation 1

The frequencies of the different bird types are:

| Bird type ID | Frequency |
|---|---:|
| 1 | 2 |
| 2 | 2 |
| 3 | 3 |
| 4 | 3 |
| 5 | 1 |

Bird types `3` and `4` are tied for the highest frequency, appearing three times each.

Since the problem requires choosing the smallest ID in case of a tie, the answer is `3`.

## Output Format

Return the type ID of the most frequently sighted bird. If multiple types have the same maximum frequency, return the smallest ID among them.
*/

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countOccurrences(arr []int32) []int32 {
	counts := make([]int32, 5)

	for _, id := range arr {
		counts[id-1]++
	}

	return counts
}

func findMostFrequentTypeIndex(counts []int32) int32 {
	maxValueIndex := 0

	for index, count := range counts {
		if count > counts[maxValueIndex] {
			maxValueIndex = index
		}
	}

	return int32(maxValueIndex)
}

func migratoryBirds(arr []int32) int32 {
	counts := countOccurrences(arr)
	maxValueIndex := findMostFrequentTypeIndex(counts)

	return maxValueIndex + 1
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}), migratoryBirds([]int32{1, 4, 4, 4, 5, 3}) == 4)                               // Esperado: 4
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}), migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}) == 3) // Esperado: 3
	fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}), migratoryBirds([]int32{1, 1, 2, 2, 3}) == 1)                                     // Esperado: 1
	fmt.Println(migratoryBirds([]int32{5, 5, 5, 1, 1, 1, 3}), migratoryBirds([]int32{5, 5, 5, 1, 1, 1, 3}) == 1)                         // Esperado: 1
	fmt.Println(migratoryBirds([]int32{2, 2, 2, 3, 3, 3, 1}), migratoryBirds([]int32{2, 2, 2, 3, 3, 3, 1}) == 2)                         // Esperado: 2
}
