// Package status defines the possible statuses for payment attempts and provides utility methods to validate and check these statuses.
package status

type Status string

const (
	StatusSuccess Status = "SUCCESS"
	StatusFailed  Status = "FAILED"
	StatusPending Status = "PENDING"
)

func (s Status) IsValidStatus() bool {
	switch s {
	case StatusSuccess, StatusFailed, StatusPending:
		return true
	default:
		return false
	}
}

func (s Status) IsInvalidStatus() bool {
	return !s.IsValidStatus()
}

func (s Status) IsSuccessStatus() bool {
	return s == StatusSuccess
}

func (s Status) IsNotSuccessStatus() bool {
	return s != StatusSuccess
}

func (s Status) IsFailedStatus() bool {
	return s == StatusFailed
}

func (s Status) IsNotFailedStatus() bool {
	return s != StatusFailed
}

func (s Status) IsPendingStatus() bool {
	return s == StatusPending
}

func (s Status) IsNotPendingStatus() bool {
	return s != StatusPending
}
