package main

import (
	"reflect"
	"testing"

	"coding-challenges-in-go/live-coding/06.Active_Enrollments/status"
)

func TestConsolidateEnrollments(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		events  []EnrollmentEvent
		want    []StudentCoursesReport
		wantErr bool
	}{
		{
			name: "Eventos fora de ordem e cancelamento",
			events: []EnrollmentEvent{
				{StudentID: 1, CourseCode: "MATH", Status: status.Enrolled, Timestamp: 100},
				{StudentID: 1, CourseCode: "MATH", Status: status.Dropped, Timestamp: 200},
				{StudentID: 1, CourseCode: "HIST", Status: status.Enrolled, Timestamp: 300},
			},
			want: []StudentCoursesReport{
				{StudentID: 1, ActiveCourses: []string{"HIST"}},
			},
			wantErr: false,
		},
		{
			name: "Ordenação de alunos e disciplinas",
			events: []EnrollmentEvent{
				{StudentID: 2, CourseCode: "PHYS", Status: status.Enrolled, Timestamp: 100},
				{StudentID: 2, CourseCode: "BIOL", Status: status.Enrolled, Timestamp: 100},
				{StudentID: 1, CourseCode: "MATH", Status: status.Enrolled, Timestamp: 100},
			},
			want: []StudentCoursesReport{
				{StudentID: 1, ActiveCourses: []string{"MATH"}},
				{StudentID: 2, ActiveCourses: []string{"BIOL", "PHYS"}},
			},
			wantErr: false,
		},
		{
			name: "Erro de validação em status inválido",
			events: []EnrollmentEvent{
				{StudentID: 1, CourseCode: "MATH", Status: "UNKNOWN", Timestamp: 100},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ConsolidateEnrollments(tt.events)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ConsolidateEnrollments() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ConsolidateEnrollments() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConsolidateEnrollments() = %v, want %v", got, tt.want)
			}
		})
	}
}
