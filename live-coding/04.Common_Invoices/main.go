package main

import "fmt"

func hasEmptyInvoices(gatewayInvoices, internalInvoices []int) bool {
	if len(gatewayInvoices) == 0 || len(internalInvoices) == 0 {
		return true
	}

	return false
}

func createInvoiceSet(invoices []int) map[int]struct{} {
	invoiceSet := map[int]struct{}{}

	for _, invoice := range invoices {
		invoiceSet[invoice] = struct{}{}
	}

	return invoiceSet
}

func findCommonInvoices(internalInvoices []int, invoiceSet map[int]struct{}) []int {
	invoices := []int{}

	for _, invoice := range internalInvoices {
		if _, exists := invoiceSet[invoice]; exists {
			invoices = append(invoices, invoice)
			delete(invoiceSet, invoice)
		}
	}

	return invoices
}

func FindCommonInvoices(gatewayInvoices, internalInvoices []int) []int {
	if hasEmptyInvoices(gatewayInvoices, internalInvoices) {
		return []int{}
	}

	invoiceSet := createInvoiceSet(gatewayInvoices)

	return findCommonInvoices(internalInvoices, invoiceSet)
}

func main() {
	gateway := []int{4, 9, 5, 4}
	internal := []int{9, 4, 9, 8, 4}
	common := FindCommonInvoices(gateway, internal)
	fmt.Println("Common invoices:", common)
}
