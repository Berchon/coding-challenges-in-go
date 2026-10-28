package main

import (
	"reflect"
	"testing"
)

func TestReconcileSettlements(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		invoices    []Invoice
		settlements []Settlement
		want        []SchoolSummary
		wantErr     bool
	}{
		// TODO: invoices e settlements vazios -> retorna lista vazia
		{
			name:        "invoices e settlements vazios deve retornar lista vazia",
			invoices:    []Invoice{},
			settlements: []Settlement{},
			want:        []SchoolSummary{},
			wantErr:     false,
		},
		// TODO: invoices com InvoiceID vazia -> retorna erro
		{
			name:        "invoices com InvoiceID vazia deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{},
			want:        nil,
			wantErr:     true,
		},
		// TODO: invoices com SchoolID menor ou igual a zero -> retorna erro
		{
			name:        "invoices com SchoolID menor ou igual a zero deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 0, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{},
			want:        nil,
			wantErr:     true,
		},
		// TODO: invoices com Amount menor ou igual a zero -> retorna erro
		{
			name:        "invoices com Amount menor ou igual a zero deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 0, DueDate: "2026-04-10"}},
			settlements: []Settlement{},
			want:        nil,
			wantErr:     true,
		},
		// TODO: invoices com DueDate vazia -> retorna erro
		{
			name:        "invoices com DueDate vazia deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: ""}},
			settlements: []Settlement{},
			want:        nil,
			wantErr:     true,
		},
		// TODO: invoices com InvoiceID repetidos -> retorna erro
		{
			name:        "invoices com InvoiceID repetidos deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}, {InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{},
			want:        nil,
			wantErr:     true,
		},
		// TODO: settlements com InvoiceID vazio -> retorna erro
		{
			name:        "settlements com InvoiceID vazio deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"}},
			want:        nil,
			wantErr:     true,
		},
		// TODO: settlements com PaidAmount menor ou igual a zero -> retorna erro
		{
			name:        "settlements com PaidAmount menor ou igual a zero deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-101", PaidAmount: 0, PaidDate: "2026-04-10", Status: "CONFIRMED"}},
			want:        nil,
			wantErr:     true,
		},
		// TODO: settlements com PaidDate vazia -> retorna erro
		{
			name:        "settlements com PaidDate vazia deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "", Status: "CONFIRMED"}},
			want:        nil,
			wantErr:     true,
		},
		// TODO: settlements com Status desconhecido -> retorna erro
		{
			name:        "settlements com Status desconhecido deve retornar erro",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "UNKNOWN"}},
			want:        nil,
			wantErr:     true,
		},
		// TODO: settlements com InvoiceID que nao existe na lista de invoices -> ignora esse settlement
		{
			name:        "settlements com InvoiceID que nao existe na lista de invoices deve ignorar esse settlement",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-999", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"}},
			want:        []SchoolSummary{},
			wantErr:     false,
		},
		// TODO: settlements com PaidAmount diferente do Amount da invoice correspondente -> ignora esse settlement
		{
			name:        "settlements com PaidAmount diferente do Amount da invoice correspondente deve ignorar esse settlement",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-101", PaidAmount: 40000, PaidDate: "2026-04-10", Status: "CONFIRMED"}},
			want:        []SchoolSummary{},
			wantErr:     false,
		},
		// TODO: settlements com Status diferente de "CONFIRMED" -> ignora esse settlement
		{
			name:        "settlements com Status diferente de 'CONFIRMED' deve ignorar esse settlement",
			invoices:    []Invoice{{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"}},
			settlements: []Settlement{{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "REJECTED"}},
			want:        []SchoolSummary{},
			wantErr:     false,
		},
		// TODO: settlements com InvoiceID repetidos -> considera apenas o mais antigo (o que vier primeiro na lista)
		{
			name: "settlements com InvoiceID repetidos deve considerar apenas o mais antigo (o que vier primeiro na lista)",
			invoices: []Invoice{
				{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"},
			},
			settlements: []Settlement{
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"},
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-11", Status: "CONFIRMED"},
			},
			want: []SchoolSummary{
				{SchoolID: 10, TotalCollected: 50000, SettledInvoicesCount: 1},
			},
			wantErr: false,
		},
		// TODO: invoices e settlements válidos -> retorna relatório correto
		{
			name: "invoices e settlements válidos deve retornar relatório correto",
			invoices: []Invoice{
				{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"},
				{InvoiceID: "INV-201", SchoolID: 20, Amount: 80000, DueDate: "2026-04-15"},
			},
			settlements: []Settlement{
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-15", Status: "CONFIRMED"},
			},
			want: []SchoolSummary{
				{SchoolID: 20, TotalCollected: 80000, SettledInvoicesCount: 1},
				{SchoolID: 10, TotalCollected: 50000, SettledInvoicesCount: 1},
			},
			wantErr: false,
		},
		// TODO: invoices e settlements válidos com múltiplos settlements para a mesma invoice -> retorna relatório correto considerando apenas o mais antigo
		{
			name: "invoices e settlements válidos com múltiplos settlements para a mesma invoice deve retornar relatório correto considerando apenas o mais antigo",
			invoices: []Invoice{
				{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"},
				{InvoiceID: "INV-201", SchoolID: 20, Amount: 80000, DueDate: "2026-04-15"},
			},
			settlements: []Settlement{
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"},
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-11", Status: "CONFIRMED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-15", Status: "CONFIRMED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-16", Status: "CONFIRMED"},
			},
			want: []SchoolSummary{
				{SchoolID: 20, TotalCollected: 80000, SettledInvoicesCount: 1},
				{SchoolID: 10, TotalCollected: 50000, SettledInvoicesCount: 1},
			},
			wantErr: false,
		},
		// TODO: invoices e settlements válidos com múltiplos settlements para a mesma invoice, mas com status diferente de "CONFIRMED" -> retorna relatório correto considerando apenas o mais antigo com status "CONFIRMED"
		{
			name: "invoices e settlements válidos com múltiplos settlements para a mesma invoice, mas com status diferente de 'CONFIRMED' deve retornar relatório correto considerando apenas o mais antigo com status 'CONFIRMED'",
			invoices: []Invoice{
				{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"},
				{InvoiceID: "INV-201", SchoolID: 20, Amount: 80000, DueDate: "2026-04-15"},
			},
			settlements: []Settlement{
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "REFUNDED"},
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-11", Status: "REJECTED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-15", Status: "REJECTED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-16", Status: "CONFIRMED"},
			},
			want: []SchoolSummary{
				{SchoolID: 20, TotalCollected: 80000, SettledInvoicesCount: 1},
			},
			wantErr: false,
		},
		// TODO: o relatório deve ser ordenado de forma decrescente pelo TotalCollected, e em caso de empate, de forma crescente pelo SchoolID
		{
			name: "o relatório deve ser ordenado de forma decrescente pelo TotalCollected, e em caso de empate, de forma crescente pelo SchoolID",
			invoices: []Invoice{
				{InvoiceID: "INV-101", SchoolID: 10, Amount: 50000, DueDate: "2026-04-10"},
				{InvoiceID: "INV-201", SchoolID: 20, Amount: 80000, DueDate: "2026-04-15"},
				{InvoiceID: "INV-301", SchoolID: 30, Amount: 50000, DueDate: "2026-04-20"},
			},
			settlements: []Settlement{
				{InvoiceID: "INV-101", PaidAmount: 50000, PaidDate: "2026-04-10", Status: "CONFIRMED"},
				{InvoiceID: "INV-201", PaidAmount: 80000, PaidDate: "2026-04-15", Status: "CONFIRMED"},
				{InvoiceID: "INV-301", PaidAmount: 50000, PaidDate: "2026-04-20", Status: "CONFIRMED"},
			},
			want: []SchoolSummary{
				{SchoolID: 20, TotalCollected: 80000, SettledInvoicesCount: 1},
				{SchoolID: 10, TotalCollected: 50000, SettledInvoicesCount: 1},
				{SchoolID: 30, TotalCollected: 50000, SettledInvoicesCount: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ReconcileSettlements(tt.invoices, tt.settlements)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ReconcileSettlements() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ReconcileSettlements() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconcileSettlements() = %v, want %v", got, tt.want)
			}
		})
	}
}
