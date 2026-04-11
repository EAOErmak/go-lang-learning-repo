package handlers

import (
	"errors"
	"strings"

	"go-learn/main/models"

	"gorm.io/gorm"
)

func ensureDictionaryItemUnique(tx *gorm.DB, excludeID uint, itemType models.DictionaryType, label string) error {
	query := tx.Model(&models.DictionaryItem{}).Where("type = ? AND label = ?", itemType, label)
	if excludeID != 0 {
		query = query.Where("id <> ?", excludeID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("dictionary item already exists")
	}

	return nil
}

func dictionaryItemInUse(tx *gorm.DB, itemID uint) (bool, error) {
	var metricCount int64
	if err := tx.Model(&models.EntryMetric{}).Where("metric_type_id = ?", itemID).Count(&metricCount).Error; err != nil {
		return false, err
	}

	if metricCount > 0 {
		return true, nil
	}

	var valueCount int64
	if err := tx.Model(&models.EntryMetricValue{}).Where("unit_id = ?", itemID).Count(&valueCount).Error; err != nil {
		return false, err
	}

	return valueCount > 0, nil
}

func isUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	message := strings.ToLower(err.Error())
	return strings.Contains(message, "duplicate key") ||
		strings.Contains(message, "unique constraint") ||
		strings.Contains(message, "violates unique")
}
