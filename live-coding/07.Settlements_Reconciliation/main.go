package main

import (
	"cmp"
	"fmt"
	"slices"

	"coding-challenges-in-go/live-coding/07.Settlements_Reconciliation/status"
)

type Invoice struct {
	InvoiceID string
	SchoolID  int
	Amount    int
	DueDate   string // aqui eu poderia converter para time.Time
}

type Settlement struct {
	InvoiceID  string
	PaidAmount int
	PaidDate   string // aqui eu poderia converter para time.Time
	Status     status.Status
}

type SchoolSummary struct {
	SchoolID             int
	TotalCollected       int
	SettledInvoicesCount int
}

func validateInvoice(invoice Invoice) error {
	if invoice.InvoiceID == "" {
		return fmt.Errorf("InvoiceID não pode ser vazio")
	}
	if invoice.SchoolID <= 0 {
		return fmt.Errorf("SchoolID deve ser maior que zero")
	}
	if invoice.Amount <= 0 {
		return fmt.Errorf("amount deve ser maior que zero")
	}
	if invoice.DueDate == "" {
		return fmt.Errorf("DueDate não pode ser vazio")
	}

	return nil
}

func validateSettlement(settlement Settlement) error {
	if settlement.InvoiceID == "" {
		return fmt.Errorf("InvoiceID não pode ser vazio")
	}
	if settlement.PaidAmount <= 0 {
		return fmt.Errorf("PaidAmount deve ser maior que zero")
	}
	if settlement.PaidDate == "" {
		return fmt.Errorf("PaidDate não pode ser vazio")
	}
	if settlement.Status.IsInvalid() {
		return fmt.Errorf("status desconhecido: %s", settlement.Status)
	}

	return nil
}

func buildInvoiceMap(invoices []Invoice) (map[string]Invoice, error) {
	invoiceMap := make(map[string]Invoice, len(invoices))

	for _, invoice := range invoices {
		if err := validateInvoice(invoice); err != nil {
			return nil, err
		}

		if _, exists := invoiceMap[invoice.InvoiceID]; exists {
			return nil, fmt.Errorf("InvoiceID repetido: %s", invoice.InvoiceID)
		}

		invoiceMap[invoice.InvoiceID] = invoice
	}

	return invoiceMap, nil
}

func isFullySettled(settlement Settlement, invoice Invoice) bool {
	return settlement.Status == status.Confirmed &&
		settlement.PaidAmount == invoice.Amount
}

func addSettlementToSummary(summaries map[int]SchoolSummary, invoice Invoice, settlement Settlement) {
	summary := summaries[invoice.SchoolID]

	summary.SchoolID = invoice.SchoolID
	summary.TotalCollected += settlement.PaidAmount
	summary.SettledInvoicesCount++

	summaries[invoice.SchoolID] = summary
}

func buildSchoolSummaryMap(settlements []Settlement, invoiceMap map[string]Invoice) (map[int]SchoolSummary, error) {
	summaries := make(map[int]SchoolSummary)

	for _, settlement := range settlements {
		if err := validateSettlement(settlement); err != nil {
			return nil, err
		}

		invoice, exists := invoiceMap[settlement.InvoiceID]
		if !exists {
			continue
		}

		if !isFullySettled(settlement, invoice) {
			continue
		}

		addSettlementToSummary(summaries, invoice, settlement)

		// Uma invoice só pode ser liquidada uma vez.
		delete(invoiceMap, settlement.InvoiceID)
	}

	return summaries, nil
}

func summariesToSlice(summaries map[int]SchoolSummary) []SchoolSummary {
	result := make([]SchoolSummary, 0, len(summaries))

	for _, summary := range summaries {
		result = append(result, summary)
	}

	return result
}

func compareSchoolSummary(a, b SchoolSummary) int {
	if a.TotalCollected != b.TotalCollected {
		return cmp.Compare(b.TotalCollected, a.TotalCollected)
	}

	return cmp.Compare(a.SchoolID, b.SchoolID)
}

func buildReport(summaries map[int]SchoolSummary) []SchoolSummary {
	result := summariesToSlice(summaries)

	slices.SortFunc(result, compareSchoolSummary)

	return result
}

func ReconcileSettlements(invoices []Invoice, settlements []Settlement) ([]SchoolSummary, error) {
	invoiceMap, err := buildInvoiceMap(invoices)
	if err != nil {
		return nil, err
	}

	summaries, err := buildSchoolSummaryMap(settlements, invoiceMap)
	if err != nil {
		return nil, err
	}

	return buildReport(summaries), nil
}

func main() {
}
