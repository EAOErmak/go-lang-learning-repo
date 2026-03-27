package handlers

import (
	"errors"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}
var nextBookID = 1

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateBook(newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = nextBookID
	nextBookID++

	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var updatedBook models.Book

	err = c.ShouldBindJSON(&updatedBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	err = validateBook(updatedBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i].Title = updatedBook.Title
			books[i].AuthorID = updatedBook.AuthorID
			books[i].CategoryID = updatedBook.CategoryID
			books[i].Price = updatedBook.Price
			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = slices.Delete(books, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func validateBook(book models.Book) error {
	if book.Title == "" {
		return errors.New("book title is required")
	}

	if book.AuthorID < 1 {
		return errors.New("author is required")
	}

	if book.CategoryID < 1 {
		return errors.New("category is required")
	}

	if book.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return nil
}
