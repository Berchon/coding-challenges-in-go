// Package status define os possíveis estados de um pagamento.
package status

type Status string

const (
	Confirmed Status = "CONFIRMED"
	Rejected  Status = "REJECTED"
	Refunded  Status = "REFUNDED"
)

func (s Status) IsValid() bool {
	switch s {
	case Confirmed, Rejected, Refunded:
		return true
	default:
		return false
	}
}

func (s Status) IsInvalid() bool {
	return !s.IsValid()
}

func (s Status) IsConfirmed() bool {
	return s == Confirmed
}

func (s Status) IsNotConfirmed() bool {
	return !s.IsConfirmed()
}

func (s Status) IsRejected() bool {
	return s == Rejected
}

func (s Status) IsNotRejected() bool {
	return !s.IsRejected()
}

func (s Status) IsRefunded() bool {
	return s == Refunded
}

func (s Status) IsNotRefunded() bool {
	return !s.IsRefunded()
}
