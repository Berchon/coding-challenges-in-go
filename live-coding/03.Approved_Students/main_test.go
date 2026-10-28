package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestGetApprovedStudents(t *testing.T) {
	emptyRecords := []Record{}
	recordsLessThan0 := []Record{
		{StudentID: 1, Subject: "Math", Score: 80},
		{StudentID: 1, Subject: "Science", Score: 90},
		{StudentID: 2, Subject: "Math", Score: -700},
		{StudentID: 2, Subject: "Science", Score: 80},
	}
	recordsGreaterthan100 := []Record{
		{StudentID: 1, Subject: "Math", Score: 80},
		{StudentID: 1, Subject: "Science", Score: 90},
		{StudentID: 2, Subject: "Math", Score: 700},
		{StudentID: 2, Subject: "Science", Score: 80},
	}
	recordsLessThan70 := []Record{
		{StudentID: 1, Subject: "Math", Score: 80},
		{StudentID: 1, Subject: "Science", Score: 90},
		{StudentID: 2, Subject: "Math", Score: 60},
		{StudentID: 2, Subject: "Science", Score: 70},
		{StudentID: 2, Subject: "Science", Score: 79},
	}
	recordsEqual70 := []Record{
		{StudentID: 1, Subject: "Math", Score: 70},
		{StudentID: 1, Subject: "Science", Score: 70},
		{StudentID: 1, Subject: "Science", Score: 70},
		{StudentID: 2, Subject: "Math", Score: 60},
		{StudentID: 2, Subject: "Science", Score: 80},
	}
	recordsEqual100 := []Record{
		{StudentID: 1, Subject: "Math", Score: 100},
		{StudentID: 1, Subject: "Science", Score: 100},
		{StudentID: 1, Subject: "Science", Score: 100},
		{StudentID: 2, Subject: "Math", Score: 60},
		{StudentID: 2, Subject: "Science", Score: 80},
	}
	recordsOK := []Record{
		{StudentID: 1, Subject: "Math", Score: 80},
		{StudentID: 1, Subject: "Science", Score: 90},
		{StudentID: 2, Subject: "Math", Score: 70},
		{StudentID: 2, Subject: "Science", Score: 80},
		{StudentID: 2, Subject: "Science", Score: 80},
		{StudentID: 3, Subject: "Math", Score: 60},
		{StudentID: 3, Subject: "Science", Score: 50},
		{StudentID: 4, Subject: "Math", Score: 100},
		{StudentID: 4, Subject: "Science", Score: 90},
		{StudentID: 5, Subject: "Math", Score: 40},
		{StudentID: 5, Subject: "Science", Score: 60},
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		records []Record
		want    []StudentResult
		wantErr bool
	}{
		{"sem registros", emptyRecords, []StudentResult{}, false},
		{"com pontuação menor que 0", recordsLessThan0, nil, true},
		{"com pontuação maior que 100", recordsGreaterthan100, nil, true},
		{"com média menor que 70", recordsLessThan70, []StudentResult{{StudentID: 1, Average: 85}}, false},
		{"com média igual a 70", recordsEqual70, []StudentResult{{StudentID: 1, Average: 70}, {StudentID: 2, Average: 70}}, false},
		{"com média igual a 100", recordsEqual100, []StudentResult{{StudentID: 1, Average: 100}, {StudentID: 2, Average: 70}}, false},
		{"caso geral", recordsOK, []StudentResult{{StudentID: 1, Average: 85}, {StudentID: 2, Average: 76.66666666666667}, {StudentID: 4, Average: 95}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := GetApprovedStudents(tt.records)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetApprovedStudents() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetApprovedStudents() succeeded unexpectedly")
			}

			// Sort the slices by StudentID to ensure consistent order for comparison
			slices.SortFunc(got, func(a, b StudentResult) int {
				return a.StudentID - b.StudentID
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovedStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
