package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"coding-challenges-in-go/live-coding/06.Active_Enrollments/status"
)

type EnrollmentEvent struct {
	StudentID  int
	CourseCode string
	Status     status.Status
	Timestamp  int64
}

type StudentCoursesReport struct {
	StudentID     int
	ActiveCourses []string
}

type enrollmentKey struct {
	StudentID  int
	CourseCode string
}

func validateEnrollmentEvent(event EnrollmentEvent) error {
	if event.StudentID <= 0 {
		return errors.New("student ID must be a positive integer")
	}

	if event.Status.IsInvalid() {
		return errors.New("invalid enrollment status: " + string(event.Status))
	}

	if event.CourseCode == "" {
		return errors.New("course code cannot be empty")
	}

	if event.Timestamp <= 0 {
		return errors.New("timestamp must be a positive integer")
	}

	return nil
}

func getLatestEnrollmentEvents(events []EnrollmentEvent) (map[enrollmentKey]EnrollmentEvent, error) {
	enrollments := make(map[enrollmentKey]EnrollmentEvent, len(events))

	for _, event := range events {
		if err := validateEnrollmentEvent(event); err != nil {
			return nil, err
		}

		key := enrollmentKey{
			StudentID:  event.StudentID,
			CourseCode: event.CourseCode,
		}

		if existingEvent, exists := enrollments[key]; !exists ||
			event.Timestamp > existingEvent.Timestamp {
			enrollments[key] = event
		}
	}

	return enrollments, nil
}

func buildStudentCourseReports(
	latestEnrollments map[enrollmentKey]EnrollmentEvent,
) []StudentCoursesReport {
	activeCoursesByStudent := make(map[int][]string)

	for _, event := range latestEnrollments {
		if event.Status.IsEnrolled() {
			activeCoursesByStudent[event.StudentID] = append(
				activeCoursesByStudent[event.StudentID],
				event.CourseCode,
			)
		}
	}

	reports := make([]StudentCoursesReport, 0, len(activeCoursesByStudent))

	for studentID, courses := range activeCoursesByStudent {
		slices.Sort(courses)

		reports = append(reports, StudentCoursesReport{
			StudentID:     studentID,
			ActiveCourses: courses,
		})
	}

	sortReportsByStudentID(reports)

	return reports
}

func sortReportsByStudentID(reports []StudentCoursesReport) {
	slices.SortFunc(reports, func(a, b StudentCoursesReport) int {
		return cmp.Compare(a.StudentID, b.StudentID)
	})
}

func ConsolidateEnrollments(events []EnrollmentEvent) ([]StudentCoursesReport, error) {
	latestEnrollments, err := getLatestEnrollmentEvents(events)
	if err != nil {
		return nil, err
	}

	reports := buildStudentCourseReports(latestEnrollments)

	return reports, nil
}

func main() {
	events := []EnrollmentEvent{
		// Aluno 1: matriculou em MATH, depois cancelou (timestamp maior). Não deve ter MATH ativo.
		{StudentID: 1, CourseCode: "MATH", Status: "ENROLLED", Timestamp: 100},
		{StudentID: 1, CourseCode: "MATH", Status: "DROPPED", Timestamp: 200},

		// Aluno 1: matriculou em HIST (evento fora de ordem temporal, mas ativo).
		{StudentID: 1, CourseCode: "HIST", Status: "ENROLLED", Timestamp: 300},
		{StudentID: 1, CourseCode: "HIST", Status: "DROPPED", Timestamp: 150},

		// Aluno 2: matriculado em PHYS e BIOL (ambos ativos, testa ordenação alfabética dos cursos).
		{StudentID: 2, CourseCode: "PHYS", Status: "ENROLLED", Timestamp: 100},
		{StudentID: 2, CourseCode: "BIOL", Status: "ENROLLED", Timestamp: 120},

		// Aluno 3: suspenso no evento mais recente. Não deve aparecer no retorno.
		{StudentID: 3, CourseCode: "CHEM", Status: "ENROLLED", Timestamp: 100},
		{StudentID: 3, CourseCode: "CHEM", Status: "SUSPENDED", Timestamp: 250},
	}

	// Saída esperada (ordenada por StudentID crescente e cursos alfabeticamente):
	// [
	//   {StudentID: 1, ActiveCourses: ["HIST"]},
	//   {StudentID: 2, ActiveCourses: ["BIOL", "PHYS"]}
	// ]

	ConsolidatedReport, err := ConsolidateEnrollments(events)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Relatório consolidado de matrículas:")
	for _, report := range ConsolidatedReport {
		fmt.Printf("StudentID: %d, ActiveCourses: %v\n", report.StudentID, report.ActiveCourses)
	}
}
