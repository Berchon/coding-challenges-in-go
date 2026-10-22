package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		expected int
	}{
		// =========================================================
		// Casos oficiais do LeetCode
		// =========================================================

		{
			name:     "Oficial 1",
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "Oficial 2",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "Oficial 3",
			nums:     []int{2, 1, -1},
			expected: 0,
		},

		// =========================================================
		// Array com apenas um elemento
		// =========================================================

		{
			name:     "Um elemento positivo",
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "Um elemento negativo",
			nums:     []int{-5},
			expected: 0,
		},
		{
			name:     "Um elemento zero",
			nums:     []int{0},
			expected: 0,
		},

		// =========================================================
		// Pivot no primeiro indice
		// =========================================================

		{
			name:     "Pivot no primeiro indice",
			nums:     []int{2, 1, -1},
			expected: 0,
		},
		{
			name:     "Primeiro indice nao e pivot",
			nums:     []int{0, 1, 2, 3},
			expected: -1,
		},

		// =========================================================
		// Pivot no ultimo indice
		// =========================================================

		{
			name:     "Pivot no ultimo indice",
			nums:     []int{1, -1, 0},
			expected: 2,
		},
		{
			name:     "Ultimo indice nao e pivot",
			nums:     []int{1, -1},
			expected: -1,
		},

		// =========================================================
		// Multiplos pivots
		// Deve retornar o menor indice
		// =========================================================

		{
			name:     "Todos os indices sao pivots",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "Todos os indices sao pivots com quatro elementos",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Primeiro e ultimo sao pivots",
			nums:     []int{0, 0},
			expected: 0,
		},

		// =========================================================
		// Nenhum pivot
		// =========================================================

		{
			name:     "Nenhum pivot - crescente",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "Nenhum pivot - decrescente",
			nums:     []int{3, 2, 1},
			expected: -1,
		},
		{
			name:     "Nenhum pivot - valores positivos",
			nums:     []int{1, 1, 1, 1},
			expected: -1,
		},

		// =========================================================
		// Numeros negativos
		// =========================================================

		{
			name:     "Pivot com numeros negativos",
			nums:     []int{-1, -1, 0, -1, -1},
			expected: 2,
		},
		{
			name:     "Pivot com positivos e negativos",
			nums:     []int{-1, 2, -1},
			expected: 1,
		},
		{
			name:     "Nenhum pivot com negativos",
			nums:     []int{-1, -2, -3},
			expected: -1,
		},

		// =========================================================
		// Soma total zero
		// =========================================================

		{
			name:     "Soma total zero",
			nums:     []int{1, -1, 1, -1},
			expected: -1,
		},
		{
			name:     "Soma total zero com pivot no meio",
			nums:     []int{1, -1, 0},
			expected: 2,
		},

		// =========================================================
		// Pivot com valor impar
		// Importante para testar a divisao inteira por 2
		// =========================================================

		{
			name:     "Pivot impar",
			nums:     []int{1, 2, 7, 1, 2},
			expected: 2,
		},

		// =========================================================
		// Pivot no meio
		// =========================================================

		{
			name:     "Pivot no indice 3",
			nums:     []int{1, 2, 3, 6, 5, 1},
			expected: 3,
		},
		{
			name:     "Pivot no indice 3 com valor impar",
			nums:     []int{1, 2, 3, 7, 6, 0},
			expected: 3,
		},

		// =========================================================
		// Valores máximos permitidos pelo problema
		// =========================================================

		{
			name:     "Valores positivos maximos",
			nums:     []int{1000, 1000, 1000},
			expected: 1,
		},
		{
			name:     "Valores negativos minimos",
			nums:     []int{-1000, -1000, -1000},
			expected: 1,
		},

		// =========================================================
		// Valores extremos combinados
		// =========================================================

		{
			name:     "Positivo e negativo maximos",
			nums:     []int{1000, -1000, 1000},
			expected: 1,
		},
		{
			name:     "Negativo e positivo maximos",
			nums:     []int{-1000, 1000, -1000},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			// TODO: update the condition below to compare got with tt.expected.
			if got != tt.expected {
				t.Errorf("pivotIndex() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
