// Package transactionstatus define os possíveis estados de uma transação.
package status

type Status string

const (
	Enrolled  Status = "ENROLLED"
	Dropped   Status = "DROPPED"
	Suspended Status = "SUSPENDED"
)

func (s Status) IsValid() bool {
	switch s {
	case Enrolled, Dropped, Suspended:
		return true
	default:
		return false
	}
}

func (s Status) IsInvalid() bool {
	return !s.IsValid()
}

func (s Status) IsEnrolled() bool {
	return s == Enrolled
}

func (s Status) IsDropped() bool {
	return s == Dropped
}

func (s Status) IsSuspended() bool {
	return s == Suspended
}
