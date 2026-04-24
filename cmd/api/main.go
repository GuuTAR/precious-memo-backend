package main

import (
	"log"

	"github.com/GuuTAR/precious-memo-backend/internal/config"
	"github.com/GuuTAR/precious-memo-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	// Load ENV Configuration
	var cfg *config.Config
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Database Connection
	_, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Router Initializaion
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run server
	router.Run(":" + cfg.Port)
}
