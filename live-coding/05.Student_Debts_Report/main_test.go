package main

import (
	"reflect"
	"testing"

	"coding-challenges-in-go/live-coding/05.Student_Debts_Report/transactionstatus"
)

func Test_calculateStudentBalances(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		transactions []Transaction
		want         map[int]float64
		wantErr      bool
	}{
		{
			name: "Test case 1: Normal case",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 1, Amount: -50, Status: transactionstatus.Paid},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Pending},
			},
			want: map[int]float64{
				1: 50,
				2: 200,
			},
			wantErr: false,
		},
		{
			name: "Test case 2: Invalid transaction status",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 2, Amount: 200, Status: "INVALID"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:         "Test case 3: Empty transactions",
			transactions: []Transaction{},
			want:         map[int]float64{},
			wantErr:      false,
		},
		{
			name: "Test case 4: All transactions canceled",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Canceled},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Canceled},
			},
			want:    map[int]float64{},
			wantErr: false,
		},
		{
			name: "Test case 5: Mixed transactions",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 1, Amount: -50, Status: transactionstatus.Paid},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Canceled},
				{StudentID: 3, Amount: 300, Status: transactionstatus.Pending},
			},
			want: map[int]float64{
				1: 50,
				3: 300,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := calculateStudentBalances(tt.transactions)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("calculateStudentBalances() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("calculateStudentBalances() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateStudentBalances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterDebtors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		balances map[int]float64
		want     []DebtReport
	}{
		{
			name:     "Test case 1: Normal case",
			balances: map[int]float64{1: 50, 2: 200},
			want: []DebtReport{
				{StudentID: 2, OutstandingDebt: 200},
				{StudentID: 1, OutstandingDebt: 50},
			},
		},
		{
			name:     "Test case 2: No debtors",
			balances: map[int]float64{1: 0, 2: -100},
			want:     []DebtReport{},
		},
		{
			name:     "Test case 3: All debtors",
			balances: map[int]float64{1: 50, 2: 200, 3: 0},
			want: []DebtReport{
				{StudentID: 2, OutstandingDebt: 200},
				{StudentID: 1, OutstandingDebt: 50},
			},
		},
		{
			name:     "Test case 4: Empty balances",
			balances: map[int]float64{},
			want:     []DebtReport{},
		},
		{
			name:     "Test case 5: Mixed balances",
			balances: map[int]float64{1: 50, 2: -100, 3: 200, 4: 0},
			want: []DebtReport{
				{StudentID: 3, OutstandingDebt: 200},
				{StudentID: 1, OutstandingDebt: 50},
			},
		},
		{
			name:     "Test case 6: Single debtor",
			balances: map[int]float64{1: 50, 2: 0, 3: -100},
			want: []DebtReport{
				{StudentID: 1, OutstandingDebt: 50},
			},
		},
		{
			name:     "Test case 7: Single non-debtor",
			balances: map[int]float64{1: 0, 2: -100, 3: 0},
			want:     []DebtReport{},
		},
		{
			name:     "Test case 8: All zero balances",
			balances: map[int]float64{1: 0, 2: 0, 3: 0},
			want:     []DebtReport{},
		},
		{
			name:     "Test case 9: All negative balances",
			balances: map[int]float64{1: -50, 2: -100, 3: -200},
			want:     []DebtReport{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterDebtors(tt.balances)

			sortDebtReports(got)

			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterDebtors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortDebtReports(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		debtReports []DebtReport
		want        []DebtReport
	}{
		{
			name:        "Test case 1: Normal case",
			debtReports: []DebtReport{{StudentID: 1, OutstandingDebt: 50}, {StudentID: 2, OutstandingDebt: 200}},
			want:        []DebtReport{{StudentID: 2, OutstandingDebt: 200}, {StudentID: 1, OutstandingDebt: 50}},
		},
		{
			name:        "Test case 2: Empty slice",
			debtReports: []DebtReport{},
			want:        []DebtReport{},
		},
		{
			name:        "Test case 3: Already sorted",
			debtReports: []DebtReport{{StudentID: 2, OutstandingDebt: 200}, {StudentID: 1, OutstandingDebt: 50}},
			want:        []DebtReport{{StudentID: 2, OutstandingDebt: 200}, {StudentID: 1, OutstandingDebt: 50}},
		},
		{
			name:        "Test case 4: Same debt, different students",
			debtReports: []DebtReport{{StudentID: 2, OutstandingDebt: 100}, {StudentID: 1, OutstandingDebt: 100}},
			want:        []DebtReport{{StudentID: 1, OutstandingDebt: 100}, {StudentID: 2, OutstandingDebt: 100}},
		},
		{
			name:        "Test case 5: Mixed debts and students",
			debtReports: []DebtReport{{StudentID: 3, OutstandingDebt: 150}, {StudentID: 1, OutstandingDebt: 200}, {StudentID: 2, OutstandingDebt: 100}},
			want:        []DebtReport{{StudentID: 1, OutstandingDebt: 200}, {StudentID: 3, OutstandingDebt: 150}, {StudentID: 2, OutstandingDebt: 100}},
		},
		{
			name:        "Test case 6: Single element",
			debtReports: []DebtReport{{StudentID: 1, OutstandingDebt: 100}},
			want:        []DebtReport{{StudentID: 1, OutstandingDebt: 100}},
		},
		{
			name:        "Test case 7: Same debt, order by student ID",
			debtReports: []DebtReport{{StudentID: 2, OutstandingDebt: 100}, {StudentID: 1, OutstandingDebt: 100}},
			want:        []DebtReport{{StudentID: 1, OutstandingDebt: 100}, {StudentID: 2, OutstandingDebt: 100}},
		},
		{
			name:        "Test case 8: Negative debts",
			debtReports: []DebtReport{{StudentID: 1, OutstandingDebt: -50}, {StudentID: 2, OutstandingDebt: -100}},
			want:        []DebtReport{{StudentID: 1, OutstandingDebt: -50}, {StudentID: 2, OutstandingDebt: -100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortDebtReports(tt.debtReports)
			if !reflect.DeepEqual(tt.debtReports, tt.want) {
				t.Errorf("sortDebtReports() = %v, want %v", tt.debtReports, tt.want)
			}
		})
	}
}

func TestGenerateDebtReport(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		transactions []Transaction
		want         []DebtReport
		wantErr      bool
	}{
		{
			name:         "Test case 1: Invalid transaction status",
			transactions: []Transaction{{StudentID: 1, Amount: 100, Status: "invalid"}},
			wantErr:      true,
		},
		{
			name:         "Test case 2: Empty transactions",
			transactions: []Transaction{},
			want:         []DebtReport{},
		},
		{
			name:         "Test case 3: Single student with debt",
			transactions: []Transaction{{StudentID: 1, Amount: 100, Status: transactionstatus.Pending}},
			want:         []DebtReport{{StudentID: 1, OutstandingDebt: 100}},
		},
		{
			name: "Test case 4: Multiple students with debts",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Pending},
			},
			want: []DebtReport{
				{StudentID: 2, OutstandingDebt: 200},
				{StudentID: 1, OutstandingDebt: 100},
			},
		},
		{
			name: "Test case 5: Mixed transactions with some students having no debt",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Canceled},
				{StudentID: 3, Amount: 300, Status: transactionstatus.Pending},
			},
			want: []DebtReport{
				{StudentID: 3, OutstandingDebt: 300},
				{StudentID: 1, OutstandingDebt: 100},
			},
		},
		{
			name: "Test case 6: Students with positive balances",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 1, Amount: -50, Status: transactionstatus.Paid},
				{StudentID: 2, Amount: 200, Status: transactionstatus.Pending},
				{StudentID: 2, Amount: -100, Status: transactionstatus.Paid},
			},
			want: []DebtReport{
				{StudentID: 2, OutstandingDebt: 100},
				{StudentID: 1, OutstandingDebt: 50},
			},
		},
		{
			name: "Test case 7: Multiples transactions with some equal debts for different students should be sorted by student ID",
			transactions: []Transaction{
				{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 2, Amount: 100, Status: transactionstatus.Pending},
				{StudentID: 3, Amount: 200, Status: transactionstatus.Pending},
			},
			want: []DebtReport{
				{StudentID: 3, OutstandingDebt: 200},
				{StudentID: 1, OutstandingDebt: 100},
				{StudentID: 2, OutstandingDebt: 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := GenerateDebtReport(tt.transactions)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GenerateDebtReport() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GenerateDebtReport() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateDebtReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
