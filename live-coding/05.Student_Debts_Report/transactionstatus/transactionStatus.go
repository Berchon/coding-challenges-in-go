// Package transactionstatus define os possíveis estados de uma transação.
package transactionstatus

type TransactionStatus string

const (
	Canceled TransactionStatus = "CANCELED"
	Pending  TransactionStatus = "PENDING"
	Paid     TransactionStatus = "PAID"
)

func (s TransactionStatus) IsValid() bool {
	switch s {
	case Canceled, Pending, Paid:
		return true
	default:
		return false
	}
}

func (s TransactionStatus) IsInvalid() bool {
	return !s.IsValid()
}

func (s TransactionStatus) IsCanceled() bool {
	return s == Canceled
}

func (s TransactionStatus) IsPending() bool {
	return s == Pending
}

func (s TransactionStatus) IsPaid() bool {
	return s == Paid
}
