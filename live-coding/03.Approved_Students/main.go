package main

import (
	"errors"
	"fmt"
)

type Record struct {
	StudentID int
	Subject   string
	Score     int
}

type StudentResult struct {
	StudentID int
	Average   float64
}

type studentScores struct {
	totalScore int
	count      int
}

func sumScores(records []Record) (map[int]studentScores, error) {
	scoresMap := map[int]studentScores{}
	for _, record := range records {
		if record.Score < 0 || record.Score > 100 {
			return nil, errors.New("invalid score. Score must be between 0 and 100")
		}

		curr := scoresMap[record.StudentID]
		curr.totalScore += record.Score
		curr.count++
		scoresMap[record.StudentID] = curr
	}

	return scoresMap, nil
}

func GetApprovedStudents(records []Record) ([]StudentResult, error) {
	scoresMap, err := sumScores(records)
	if err != nil {
		return nil, err
	}

	approvedStudents := make([]StudentResult, 0, len(scoresMap))
	for studentID, scores := range scoresMap {
		average := float64(scores.totalScore) / float64(scores.count)
		if average >= 70 {
			approvedStudents = append(approvedStudents, StudentResult{
				StudentID: studentID,
				Average:   average,
			})
		}
	}

	return approvedStudents, nil
}

func main() {
	records := []Record{
		{StudentID: 1, Subject: "Math", Score: 80},
		{StudentID: 1, Subject: "Science", Score: 90}, // Aprovado
		{StudentID: 2, Subject: "Math", Score: 70},
		{StudentID: 2, Subject: "Science", Score: 80},
		{StudentID: 3, Subject: "Math", Score: 60},
		{StudentID: 3, Subject: "Science", Score: 50}, // Reprovado
	}

	approvedStudents, err := GetApprovedStudents(records)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(approvedStudents) == 0 {
		fmt.Println("No approved students found.")
		return
	}

	for _, student := range approvedStudents {
		fmt.Printf("Student ID: %d, Average: %.2f\n", student.StudentID, student.Average)
	}
}
