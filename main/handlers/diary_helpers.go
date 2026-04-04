package handlers

import (
	"errors"
	"fmt"
	"go-learn/main/database"
	"go-learn/main/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func diaryDB(c *gin.Context) *gorm.DB {
	if database.DB == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "database is not initialized; set DATABASE_URL first"})
		return nil
	}

	return database.DB
}

func preloadDiaryEntry(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Metrics.MetricType").
		Preload("Metrics.Values.Unit")
}

func preloadMetric(db *gorm.DB) *gorm.DB {
	return db.
		Preload("MetricType").
		Preload("Values.Unit")
}

func parseUintParam(c *gin.Context, name string) (uint, bool) {
	raw := strings.TrimSpace(c.Param(name))
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid %s", name)})
		return 0, false
	}

	return uint(id), true
}

func ensureDictionaryItemExists(tx *gorm.DB, itemID uint) error {
	return ensureRecordExists(tx, &models.DictionaryItem{}, itemID, "dictionary item")
}

func ensureRecordExists(tx *gorm.DB, model any, id uint, entityName string) error {
	var count int64
	if err := tx.Model(model).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("%s not found", entityName)
	}

	return nil
}

func createMetricRecords(tx *gorm.DB, entryID uint, requests []entryMetricRequest) error {
	for _, req := range requests {
		metric, err := buildMetric(entryID, req)
		if err != nil {
			return err
		}

		if err := ensureDictionaryItemExists(tx, req.MetricTypeID); err != nil {
			return err
		}

		for _, value := range req.Values {
			if err := ensureDictionaryItemExists(tx, value.UnitID); err != nil {
				return err
			}
		}

		values := metric.Values
		metric.Values = nil

		if err := tx.Create(metric).Error; err != nil {
			return err
		}

		for i := range values {
			values[i].EntryMetricID = metric.ID
			if err := tx.Create(&values[i]).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func buildMetric(entryID uint, req entryMetricRequest) (*models.EntryMetric, error) {
	metric, err := models.NewEntryMetric(entryID, req.MetricTypeID)
	if err != nil {
		return nil, err
	}

	for _, value := range req.Values {
		if err := metric.AddValue(value.UnitID, value.Value); err != nil {
			return nil, err
		}
	}

	return metric, nil
}

func writeDiaryError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
	case isValidationError(err):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func isValidationError(err error) bool {
	if err == nil {
		return false
	}

	message := strings.ToLower(err.Error())

	return strings.Contains(message, "required") ||
		strings.Contains(message, "invalid") ||
		strings.Contains(message, "cannot") ||
		strings.Contains(message, "already exists") ||
		strings.Contains(message, "must be") ||
		strings.Contains(message, "not found")
}
