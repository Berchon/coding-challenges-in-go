package main

import (
	"reflect"
	"testing"
)

func TestCalculateSchoolStats(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		attempts []PaymentAttempt
		want     []SchoolStats
		wantErr  bool
	}{
		// Caso 1: PaymentAttempt vazio deve retornar um slice vazio de SchoolStats e nenhum erro.
		{
			name:     "Caso 1: PaymentAttempt vazio",
			attempts: []PaymentAttempt{},
			want:     []SchoolStats{},
			wantErr:  false,
		},

		// todo: Caso 2: PaymentAttempt com status inválido deve retornar um erro.
		{
			name:     "Caso 2: PaymentAttempt com status inválido",
			attempts: []PaymentAttempt{{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "INVALID", Timestamp: 1633024800}},
			want:     []SchoolStats{},
			wantErr:  true,
		},

		// todo: Caso 3: PaymentAttempt com Amount <= 0 deve retornar um erro.
		{
			name:     "Caso 3: PaymentAttempt com Amount <= 0",
			attempts: []PaymentAttempt{{AttemptID: "1", SchoolID: "1", Amount: 0.0, Status: "SUCCESS", Timestamp: 1633024800}},
			want:     []SchoolStats{},
			wantErr:  true,
		},

		// todo: Caso 4: PaymentAttempt com AttemptID vazio deve retornar um erro.
		{
			name:     "Caso 4: PaymentAttempt com AttemptID vazio",
			attempts: []PaymentAttempt{{AttemptID: "", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800}},
			want:     []SchoolStats{},
			wantErr:  true,
		},

		// todo: Caso 5: PaymentAttempt com SchoolID vazio deve retornar um erro.
		{
			name:     "Caso 5: PaymentAttempt com SchoolID vazio",
			attempts: []PaymentAttempt{{AttemptID: "1", SchoolID: "", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800}},
			want:     []SchoolStats{},
			wantErr:  true,
		},

		// TODO: Caso 6: PaymentAttempt com status "PENDING" deve ser ignorado no cálculo de métricas.
		{
			name: "Caso 6: PaymentAttempt com status PENDING",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "PENDING", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 50.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 50.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want: []SchoolStats{
				{SchoolID: "1", TotalProcessed: 3, FailedAttempts: 2, FailureRate: float64(2) / float64(3)},
			},
			wantErr: false,
		},

		// TODO: Caso 7: Agrupar as tentativas válidas (SUCCESS e FAILED) por SchoolID.
		{
			name: "Caso 7: Agrupar as tentativas válidas por SchoolID",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "4", SchoolID: "2", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "5", SchoolID: "2", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "6", SchoolID: "2", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want: []SchoolStats{
				{SchoolID: "1", TotalProcessed: 3, FailedAttempts: 2, FailureRate: float64(2) / float64(3)},
				{SchoolID: "2", TotalProcessed: 3, FailedAttempts: 2, FailureRate: float64(2) / float64(3)},
			},
			wantErr: false,
		},

		// TODO: Caso 8: Considerar para o relatório de risco apenas escolas que possuam pelo menos 3 tentativas processadas (SUCCESS + FAILED >= 3).
		{
			name: "Caso 8: Considerar apenas escolas com pelo menos 3 tentativas processadas",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "4", SchoolID: "2", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "5", SchoolID: "2", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want: []SchoolStats{
				{SchoolID: "1", TotalProcessed: 3, FailedAttempts: 2, FailureRate: 0.6666666666666666},
			},
			wantErr: false,
		},

		// TODO: Caso 9: Para essas escolas, calcular a taxa de falhas (FailureRate = FailedAttempts / TotalProcessed).
		{
			name: "Caso 9: Calcular a taxa de falhas para escolas com pelo menos 3 tentativas processadas",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "4", SchoolID: "2", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "5", SchoolID: "2", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want: []SchoolStats{
				{SchoolID: "1", TotalProcessed: 3, FailedAttempts: 2, FailureRate: 0.6666666666666666},
			},
			wantErr: false,
		},

		// TODO: Caso 10: Incluir no relatório apenas as escolas cuja FailureRate seja estritamente maior que 50% (> 0.50).
		{
			name: "Caso 10: Incluir apenas as escolas com FailureRate > 0.50",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "3", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
				{AttemptID: "4", SchoolID: "2", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "5", SchoolID: "2", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want: []SchoolStats{
				{SchoolID: "1", TotalProcessed: 3, FailedAttempts: 2, FailureRate: 0.6666666666666666},
			},
			wantErr: false,
		},

		// todo: Caso 11: PaymentAttempt com sucesso e falha deve calcular a taxa de falhas corretamente.
		{
			name: "Caso 11: FailureRate menor ou igual à 0.5 não deve ser incluído no relatório",
			attempts: []PaymentAttempt{
				{AttemptID: "1", SchoolID: "1", Amount: 100.0, Status: "SUCCESS", Timestamp: 1633024800},
				{AttemptID: "2", SchoolID: "1", Amount: 100.0, Status: "FAILED", Timestamp: 1633024800},
			},
			want:    []SchoolStats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := CalculateSchoolStats(tt.attempts)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CalculateSchoolStats() failed: %v", gotErr)
				}
				return
			}

			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateSchoolStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
