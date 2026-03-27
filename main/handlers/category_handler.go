package handlers

import (
	"errors"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{}
var nextCategoryID = 1

func GetAllCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var newCategory models.Category

	err := c.ShouldBindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateCategory(newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory.ID = nextCategoryID
	nextCategoryID++

	categories = append(categories, newCategory)

	c.JSON(http.StatusCreated, newCategory)
}

func GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category id"})
		return
	}

	for _, category := range categories {
		if category.ID == id {
			c.JSON(http.StatusOK, category)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category id"})
		return
	}

	var updatedCategory models.Category

	err = c.ShouldBindJSON(&updatedCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateCategory(updatedCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, category := range categories {
		if category.ID == id {
			categories[i].Name = updatedCategory.Name
			c.JSON(http.StatusOK, categories[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category id"})
		return
	}

	for i, category := range categories {
		if category.ID == id {
			categories = slices.Delete(categories, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

func validateCategory(category models.Category) error {
	if category.Name == "" {
		return errors.New("Category name is required")
	}
	return nil
}
