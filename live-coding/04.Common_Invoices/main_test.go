package main

import (
	"reflect"
	"testing"
)

func Test_hasEmptyInvoices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		gatewayInvoices  []int
		internalInvoices []int
		want             bool
	}{
		{
			name:             "Test case 1: Both slices are empty",
			gatewayInvoices:  []int{},
			internalInvoices: []int{},
			want:             true,
		},
		{
			name:             "Test case 2: Gateway invoices slice is empty",
			gatewayInvoices:  []int{},
			internalInvoices: []int{1, 2, 3},
			want:             true,
		},
		{
			name:             "Test case 3: Internal invoices slice is empty",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{},
			want:             true,
		},
		{
			name:             "Test case 4: Both slices have elements",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{4, 5, 6},
			want:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasEmptyInvoices(tt.gatewayInvoices, tt.internalInvoices)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("hasEmptyInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createInvoiceSet(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		invoices []int
		want     map[int]struct{}
	}{
		{
			name:     "Test case 1: Empty slice",
			invoices: []int{},
			want:     map[int]struct{}{},
		},
		{
			name:     "Test case 2: Slice with one element",
			invoices: []int{1},
			want:     map[int]struct{}{1: {}},
		},
		{
			name:     "Test case 3: Slice with elements",
			invoices: []int{1, 2, 3},
			want:     map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Test case 4: Slice with duplicate elements",
			invoices: []int{1, 2, 2, 3},
			want:     map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Test case 5: Slice with multiple duplicate elements",
			invoices: []int{1, 1, 2, 2, 3, 3},
			want:     map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Test case 6: Slice with negative elements",
			invoices: []int{-1, -2, -3},
			want:     map[int]struct{}{-1: {}, -2: {}, -3: {}},
		},
		{
			name:     "Test case 7: Slice with mixed positive and negative elements",
			invoices: []int{-1, 0, 1},
			want:     map[int]struct{}{-1: {}, 0: {}, 1: {}},
		},
		{
			name:     "Test case 8: Slice with zero",
			invoices: []int{0},
			want:     map[int]struct{}{0: {}},
		},
		{
			name:     "Test case 9: Slice with multiple elements including zero without order",
			invoices: []int{5, 8, 3, 4, 1, 9, 2, 7, 6, 7, 0},
			want:     map[int]struct{}{5: {}, 8: {}, 3: {}, 4: {}, 1: {}, 2: {}, 7: {}, 6: {}, 9: {}, 0: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createInvoiceSet(tt.invoices)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("createInvoiceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCommonInvoices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		internalInvoices []int
		invoiceSet       map[int]struct{}
		want             []int
	}{
		{
			name:             "Test case 1: Empty slices",
			internalInvoices: []int{},
			invoiceSet:       map[int]struct{}{},
			want:             []int{},
		},
		{
			name:             "Test case 2: No common elements",
			internalInvoices: []int{1, 2, 3},
			invoiceSet:       map[int]struct{}{4: {}, 5: {}, 6: {}},
			want:             []int{},
		},
		{
			name:             "Test case 3: Some common elements",
			internalInvoices: []int{1, 2, 3},
			invoiceSet:       map[int]struct{}{2: {}, 3: {}, 4: {}},
			want:             []int{2, 3},
		},
		{
			name:             "Test case 4: All common elements",
			internalInvoices: []int{1, 2, 3},
			invoiceSet:       map[int]struct{}{1: {}, 2: {}, 3: {}},
			want:             []int{1, 2, 3},
		},
		{
			name:             "Test case 5: Duplicate elements in internalInvoices",
			internalInvoices: []int{1, 2, 2, 3},
			invoiceSet:       map[int]struct{}{2: {}, 3: {}, 4: {}},
			want:             []int{2, 3},
		},
		{
			name:             "Test case 7: Mixed common and non-common elements",
			internalInvoices: []int{1, 2, 3, 4, 5},
			invoiceSet:       map[int]struct{}{3: {}, 4: {}, 6: {}},
			want:             []int{3, 4},
		},
		{
			name:             "Test case 8: All elements common with duplicates",
			internalInvoices: []int{1, 2, 2, 3, 3},
			invoiceSet:       map[int]struct{}{1: {}, 2: {}, 3: {}},
			want:             []int{1, 2, 3},
		},
		{
			name:             "Test case 9: No common elements with duplicates",
			internalInvoices: []int{1, 1, 2, 2, 3, 3},
			invoiceSet:       map[int]struct{}{4: {}, 5: {}, 6: {}},
			want:             []int{},
		},
		{
			name:             "Test case 10: Mixed positive and negative elements",
			internalInvoices: []int{-1, 0, 1},
			invoiceSet:       map[int]struct{}{-1: {}, 2: {}, 3: {}},
			want:             []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCommonInvoices(tt.internalInvoices, tt.invoiceSet)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("findCommonInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCommonInvoices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		gatewayInvoices  []int
		internalInvoices []int
		want             []int
	}{
		{
			name:             "Test case 1: Empty slices should return nil",
			gatewayInvoices:  []int{},
			internalInvoices: []int{},
			want:             []int{},
		},
		{
			name:             "Test case 2: No common elements",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{4, 5, 6},
			want:             []int{},
		},
		{
			name:             "Test case 3: Some common elements",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{2, 3, 4},
			want:             []int{2, 3},
		},
		{
			name:             "Test case 4: All common elements",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{1, 2, 3},
			want:             []int{1, 2, 3},
		},
		{
			name:             "Test case 5: Duplicate elements in gatewayInvoices",
			gatewayInvoices:  []int{1, 2, 2, 3},
			internalInvoices: []int{2, 3, 4},
			want:             []int{2, 3},
		},
		{
			name:             "Test case 6: Duplicate elements in internalInvoices",
			gatewayInvoices:  []int{1, 2, 3},
			internalInvoices: []int{2, 2, 3, 4},
			want:             []int{2, 3},
		},
		{
			name:             "Test case 7: Mixed common and non-common elements",
			gatewayInvoices:  []int{1, 2, 3, 4, 5},
			internalInvoices: []int{3, 4, 6},
			want:             []int{3, 4},
		},
		{
			name:             "Test case 8: All elements common with duplicates",
			gatewayInvoices:  []int{1, 2, 2, 3, 3},
			internalInvoices: []int{1, 2, 3},
			want:             []int{1, 2, 3},
		},
		{
			name:             "Test case 9: No common elements with duplicates",
			gatewayInvoices:  []int{1, 1, 2, 2, 3, 3},
			internalInvoices: []int{4, 5, 6},
			want:             []int{},
		},
		{
			name:             "Test case 10: Mixed positive and negative elements",
			gatewayInvoices:  []int{-1, 0, 1},
			internalInvoices: []int{-1, -2, 3},
			want:             []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindCommonInvoices(tt.gatewayInvoices, tt.internalInvoices)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("FindCommonInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
