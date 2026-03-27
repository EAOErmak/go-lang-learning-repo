package handlers

import (
	"errors"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Author
//Authors
//author
//authors

var authors = []models.Author{}

var nextAuthorID = 1

func GetAllAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author

	err := c.ShouldBindJSON(&newAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateAuthor(newAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAuthor.ID = nextAuthorID
	nextAuthorID++

	authors = append(authors, newAuthor)

	c.JSON(http.StatusCreated, newAuthor)
}

func GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Author id"})
		return
	}

	for _, author := range authors {
		if author.ID == id {
			c.JSON(http.StatusOK, author)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}

func UpdateAuthor(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Author id"})
		return
	}

	var updatedAuthor models.Author

	err = c.ShouldBindJSON(&updatedAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateAuthor(updatedAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, author := range authors {
		if author.ID == id {
			authors[i].Name = updatedAuthor.Name
			c.JSON(http.StatusOK, authors[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}

func DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Author id"})
		return
	}

	for i, author := range authors {
		if author.ID == id {
			authors = slices.Delete(authors, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}

func validateAuthor(author models.Author) error {
	if author.Name == "" {
		return errors.New("Author name is required")
	}

	return nil
}
