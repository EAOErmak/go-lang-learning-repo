package handlers

import (
	"errors"
	"net/http"
	"strings"

	"go-learn/main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllDictionaryItems(c *gin.Context) {
	db := diaryDB(c)
	if db == nil {
		return
	}

	query := db.Order("type ASC").Order("label ASC").Order("id ASC")

	if itemTypeRaw := strings.TrimSpace(c.Query("type")); itemTypeRaw != "" {
		itemType, err := models.ParseDictionaryType(itemTypeRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query = query.Where("type = ?", itemType)
	}

	var items []models.DictionaryItem
	if err := query.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func CreateDictionaryItem(c *gin.Context) {
	var req dictionaryItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	itemType, err := models.ParseDictionaryType(req.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := models.NewDictionaryItem(itemType, req.Label)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	if err := ensureDictionaryItemUnique(db, 0, item.Type, item.Label); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(item).Error; err != nil {
		if isUniqueConstraintError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "dictionary item already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func GetDictionaryItemByID(c *gin.Context) {
	itemID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var item models.DictionaryItem
	if err := db.First(&item, itemID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func UpdateDictionaryItem(c *gin.Context) {
	itemID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req dictionaryItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	itemType, err := models.ParseDictionaryType(req.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var item models.DictionaryItem
	if err := db.First(&item, itemID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := item.UpdateType(itemType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := item.UpdateLabel(req.Label); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ensureDictionaryItemUnique(db, item.ID, item.Type, item.Label); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		if isUniqueConstraintError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "dictionary item already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func DeleteDictionaryItem(c *gin.Context) {
	itemID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := diaryDB(c)
	if db == nil {
		return
	}

	var item models.DictionaryItem
	if err := db.First(&item, itemID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inUse, err := dictionaryItemInUse(db, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if inUse {
		c.JSON(http.StatusConflict, gin.H{"error": "dictionary item is in use"})
		return
	}

	if err := db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}
