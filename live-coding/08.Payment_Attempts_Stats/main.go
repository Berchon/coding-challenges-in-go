package main

import (
	"cmp"
	"fmt"
	"slices"

	"coding-challenges-in-go/live-coding/08.Payment_Attempts_Stats/status"
)

type PaymentAttempt struct {
	AttemptID string
	SchoolID  string
	Amount    float64
	Status    status.Status
	Timestamp int64
}

type SchoolStats struct {
	SchoolID       string
	TotalProcessed int
	FailedAttempts int
	FailureRate    float64
}

type schoolAttemptSummary struct {
	successCount int
	failureCount int
}

func validatePaymentAttempt(attempt PaymentAttempt) error {
	switch {
	case attempt.Status.IsInvalidStatus():
		return fmt.Errorf("invalid status: %s", attempt.Status)
	case attempt.Amount <= 0:
		return fmt.Errorf("amount must be greater than zero: %f", attempt.Amount)
	case attempt.AttemptID == "":
		return fmt.Errorf("attempt ID cannot be empty")
	case attempt.SchoolID == "":
		return fmt.Errorf("school ID cannot be empty")
	default:
		return nil
	}
}

func aggregateAttemptsBySchool(attempts []PaymentAttempt) (map[string]schoolAttemptSummary, error) {
	attemptCountBySchool := make(map[string]schoolAttemptSummary)

	for _, attempt := range attempts {
		if err := validatePaymentAttempt(attempt); err != nil {
			return nil, err
		}

		if attempt.Status.IsPendingStatus() {
			continue
		}

		summary := attemptCountBySchool[attempt.SchoolID]

		if attempt.Status.IsFailedStatus() {
			summary.failureCount++
		} else if attempt.Status.IsSuccessStatus() {
			summary.successCount++
		}

		attemptCountBySchool[attempt.SchoolID] = summary
	}

	return attemptCountBySchool, nil
}

func shouldIncludeSchool(stats SchoolStats) bool {
	return stats.TotalProcessed >= 3 && stats.FailureRate > 0.5
}

func buildQualifiedSchoolStats(attemptCountBySchool map[string]schoolAttemptSummary) []SchoolStats {
	qualifiedSchools := make([]SchoolStats, 0, len(attemptCountBySchool))

	for schoolID, summary := range attemptCountBySchool {
		totalProcessed := summary.successCount + summary.failureCount
		failureRate := float64(summary.failureCount) / float64(totalProcessed)

		schoolStats := SchoolStats{
			SchoolID:       schoolID,
			TotalProcessed: totalProcessed,
			FailedAttempts: summary.failureCount,
			FailureRate:    failureRate,
		}

		if shouldIncludeSchool(schoolStats) {
			qualifiedSchools = append(qualifiedSchools, schoolStats)
		}
	}

	return qualifiedSchools
}

func sortSchoolsByFailureRate(stats []SchoolStats) {
	slices.SortFunc(stats, func(a, b SchoolStats) int {
		if comparison := cmp.Compare(b.FailureRate, a.FailureRate); comparison != 0 {
			return comparison
		}

		return cmp.Compare(a.SchoolID, b.SchoolID)
	})
}

func CalculateSchoolStats(attempts []PaymentAttempt) ([]SchoolStats, error) {
	attemptCountBySchool, err := aggregateAttemptsBySchool(attempts)
	if err != nil {
		return nil, err
	}

	qualifiedSchools := buildQualifiedSchoolStats(attemptCountBySchool)
	sortSchoolsByFailureRate(qualifiedSchools)

	return qualifiedSchools, nil
}

func main() {
}
