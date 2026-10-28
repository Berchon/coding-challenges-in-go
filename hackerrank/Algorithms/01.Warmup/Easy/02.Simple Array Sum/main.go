package main

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
