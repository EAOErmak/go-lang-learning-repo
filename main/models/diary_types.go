package models

import (
	"fmt"
	"strings"
)

type DictionaryType string

const (
	DictionaryTypeMetricName DictionaryType = "METRIC_NAME"
	DictionaryTypeMetricUnit DictionaryType = "METRIC_UNIT"
)

func (t DictionaryType) IsValid() bool {
	switch t {
	case DictionaryTypeMetricName, DictionaryTypeMetricUnit:
		return true
	default:
		return false
	}
}

func ParseDictionaryType(value string) (DictionaryType, error) {
	itemType := DictionaryType(strings.ToUpper(strings.TrimSpace(value)))
	if !itemType.IsValid() {
		return "", fmt.Errorf("invalid dictionary type: %s", value)
	}

	return itemType, nil
}
