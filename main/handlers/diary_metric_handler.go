package handlers

import (
	"go-learn/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDiaryMetric(c *gin.Context) {
	entryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req entryMetricRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var metric models.EntryMetric
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&models.DiaryEntry{}, entryID).Error; err != nil {
			return err
		}

		createdMetric, err := buildMetric(entryID, req)
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

		values := createdMetric.Values
		createdMetric.Values = nil

		if err := tx.Create(createdMetric).Error; err != nil {
			return err
		}

		for i := range values {
			values[i].EntryMetricID = createdMetric.ID
			if err := tx.Create(&values[i]).Error; err != nil {
				return err
			}
		}

		return preloadMetric(tx).First(&metric, createdMetric.ID).Error
	}); err != nil {
		writeDiaryError(c, err)
		return
	}

	c.JSON(http.StatusCreated, metric)
}
