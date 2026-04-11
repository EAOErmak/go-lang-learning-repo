package main

import (
	"go-learn/main/database"
	"go-learn/main/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Printf("diary database is unavailable: %v", err)
	}

	router := gin.Default()
	router.POST("/login", handlers.Login)

	protected := router.Group("/")
	protected.Use(handlers.AuthMiddleware())
	protected.GET("/dashboard", handlers.GetDashboard)
	protected.GET("/dictionary-items", handlers.GetAllDictionaryItems)
	protected.POST("/dictionary-items", handlers.CreateDictionaryItem)
	protected.GET("/dictionary-items/:id", handlers.GetDictionaryItemByID)
	protected.PUT("/dictionary-items/:id", handlers.UpdateDictionaryItem)
	protected.DELETE("/dictionary-items/:id", handlers.DeleteDictionaryItem)

	//Diary entries
	protected.GET("/diary-entries", handlers.GetAllDiaryEntries)
	protected.POST("/diary-entries", handlers.CreateDiaryEntry)
	protected.GET("/diary-entries/:id", handlers.GetDiaryEntryByID)
	protected.PUT("/diary-entries/:id", handlers.UpdateDiaryEntry)
	protected.DELETE("/diary-entries/:id", handlers.DeleteDiaryEntry)
	protected.POST("/diary-entries/:id/metrics", handlers.AddDiaryMetric)

	router.Run(":8080")
}
