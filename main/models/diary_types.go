package models

import (
	"fmt"
	"strings"
)

type EntryStatus string

const (
	EntryStatusScheduled EntryStatus = "SCHEDULED"
	EntryStatusActive    EntryStatus = "ACTIVE"
	EntryStatusFinished  EntryStatus = "FINISHED"
	EntryStatusFailed    EntryStatus = "FAILED"
	EntryStatusDeleted   EntryStatus = "DELETED"
)

func (s EntryStatus) IsValid() bool {
	switch s {
	case EntryStatusScheduled, EntryStatusActive, EntryStatusFinished, EntryStatusFailed, EntryStatusDeleted:
		return true
	default:
		return false
	}
}

func ParseEntryStatus(value string) (EntryStatus, error) {
	status := EntryStatus(strings.ToUpper(strings.TrimSpace(value)))
	if !status.IsValid() {
		return "", fmt.Errorf("invalid entry status: %s", value)
	}

	return status, nil
}

type DictionaryType string

const (
	DictionaryTypeMetricName DictionaryType = "METRIC_NAME"
	DictionaryTypeMetricUnit DictionaryType = "METRIC_UNIT"
)
