package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"coding-challenges-in-go/live-coding/05.Student_Debts_Report/transactionstatus"
)

type Transaction struct {
	StudentID int
	Amount    float64
	Status    transactionstatus.TransactionStatus // string
}

type DebtReport struct {
	StudentID       int
	OutstandingDebt float64
}

func calculateStudentBalances(transactions []Transaction) (map[int]float64, error) {
	debts := make(map[int]float64)

	for _, transaction := range transactions {
		if transaction.Status.IsInvalid() {
			return nil, errors.New("invalid transaction status: " + string(transaction.Status))
		}

		if transaction.Status.IsPending() || transaction.Status.IsPaid() {
			debts[transaction.StudentID] += transaction.Amount
		}
	}

	return debts, nil
}

func filterDebtors(balances map[int]float64) []DebtReport {
	debtReports := make([]DebtReport, 0, len(balances))
	for studentID, balance := range balances {
		if balance > 0 {
			debtReports = append(debtReports, DebtReport{
				StudentID:       studentID,
				OutstandingDebt: balance,
			})
		}
	}
	return debtReports
}

func sortDebtReports(debtReports []DebtReport) {
	slices.SortFunc(debtReports, func(a, b DebtReport) int {
		if a.OutstandingDebt == b.OutstandingDebt {
			return cmp.Compare(a.StudentID, b.StudentID)
		}
		return cmp.Compare(b.OutstandingDebt, a.OutstandingDebt)
	})
}

func GenerateDebtReport(transactions []Transaction) ([]DebtReport, error) {
	studentBalances, err := calculateStudentBalances(transactions)
	if err != nil {
		return nil, err
	}

	debtReports := filterDebtors(studentBalances)

	sortDebtReports(debtReports)

	return debtReports, nil
}

func main() {
	// Caso 1: status desconhecido
	fmt.Println("=== Caso 1: Status desconhecido ===")

	transactions := []Transaction{
		{StudentID: 1, Amount: 100, Status: transactionstatus.Pending},
		{StudentID: 2, Amount: 200, Status: transactionstatus.Paid},
		{StudentID: 3, Amount: 50, Status: "UNKNOWN"},
	}

	reports, err := GenerateDebtReport(transactions)

	fmt.Println("Erro:", err)
	fmt.Println("Reports:", reports)
	fmt.Println()

	// Caso 2: pagamentos, pendências e cancelamentos
	fmt.Println("=== Caso 2: Pagamentos, pendências e cancelamentos ===")

	transactions = []Transaction{
		// Estudante 1:
		// Deve 500, pagou 200 => dívida de 300
		{StudentID: 1, Amount: 500, Status: transactionstatus.Pending},
		{StudentID: 1, Amount: -200, Status: transactionstatus.Paid},

		// Estudante 2:
		// Deve 300, pagou 100 => dívida de 200
		{StudentID: 2, Amount: 300, Status: transactionstatus.Pending},
		{StudentID: 2, Amount: -100, Status: transactionstatus.Paid},
		{StudentID: 2, Amount: 500, Status: transactionstatus.Canceled}, // Cancelado, não deve afetar a dívida

		// Estudante 3:
		// Deve 400, mas a transação de 400 foi cancelada.
		// Portanto, não deve aparecer no relatório.
		{StudentID: 3, Amount: 400, Status: transactionstatus.Canceled},

		// Estudante 4:
		// Pagou 100 e deve 100 => saldo 0.
		// Portanto, não deve aparecer no relatório.
		{StudentID: 4, Amount: 100, Status: transactionstatus.Pending},
		{StudentID: 4, Amount: -100, Status: transactionstatus.Paid},

		// Estudante 5:
		// Pagou 200 e deve 100 => saldo 100.
		// Portanto, não deve aparecer no relatório.
		{StudentID: 5, Amount: 100, Status: transactionstatus.Pending},
		{StudentID: 5, Amount: 100, Status: transactionstatus.Canceled},
		{StudentID: 5, Amount: -200, Status: transactionstatus.Paid},
	}

	reports, err = GenerateDebtReport(transactions)

	fmt.Println("Erro:", err)
	fmt.Println("Reports:", reports)
	fmt.Println()

	// Caso 3: empate nas dívidas.
	// Deve ordenar primeiro pela maior dívida e,
	// em caso de empate, pelo menor StudentID.
	fmt.Println("=== Caso 3: Empate nas dívidas ===")

	transactions = []Transaction{
		// Dívida: 500
		{StudentID: 10, Amount: 500, Status: transactionstatus.Pending},

		// Dívida: 300
		{StudentID: 20, Amount: 300, Status: transactionstatus.Pending},

		// Dívida: 500
		{StudentID: 5, Amount: 500, Status: transactionstatus.Pending},

		// Dívida: 300
		{StudentID: 15, Amount: 300, Status: transactionstatus.Pending},

		// Dívida: 700
		{StudentID: 30, Amount: 700, Status: transactionstatus.Pending},
	}

	reports, err = GenerateDebtReport(transactions)

	fmt.Println("Erro:", err)
	fmt.Println("Reports:")

	for _, report := range reports {
		fmt.Printf(
			"StudentID: %d | OutstandingDebt: %.2f\n",
			report.StudentID,
			report.OutstandingDebt,
		)
	}
}
