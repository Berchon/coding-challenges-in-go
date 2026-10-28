package main

import "fmt"

// Minha solução

func roundGrade(grade int32) int32 {
	if grade < 38 {
		return grade
	}

	nextMultipleOfFive := grade + (5 - grade%5)
	if nextMultipleOfFive-grade < 3 {
		return nextMultipleOfFive
	}

	return grade
}

func gradingStudents(grades []int32) []int32 {
	finalGrades := make([]int32, len(grades))
	for i, grade := range grades {
		finalGrades[i] = roundGrade(grade)
	}

	return finalGrades
}

func main() {
	grades := []int32{73, 67, 38, 33}

	result := gradingStudents(grades)

	fmt.Println(result)
}
