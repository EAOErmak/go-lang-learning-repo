package main

import (
	"go-learn/main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Books
	router.GET("/books", handlers.GetAllBooks)
	router.POST("/books", handlers.CreateBook)
	router.GET("/books/:id", handlers.GetBookByID)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	//Authors
	router.GET("/authors", handlers.GetAllAuthors)
	router.POST("/authors", handlers.CreateAuthor)
	router.GET("/authors/:id", handlers.GetAuthorByID)
	router.PUT("/authors/:id", handlers.UpdateAuthor)
	router.DELETE("/authors/:id", handlers.DeleteAuthor)

	//Category
	router.GET("/categories", handlers.GetAllCategories)
	router.POST("/categories", handlers.CreateCategory)
	router.GET("/categories/:id", handlers.GetCategoryByID)
	router.PUT("/categories/:id", handlers.UpdateCategory)
	router.DELETE("/categories/:id", handlers.DeleteCategory)

	router.Run(":8080")
}
