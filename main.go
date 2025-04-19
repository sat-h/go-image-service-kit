package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sat-h/go-image-service-kit/handlers"
	"net/http"
)

func main() {
	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.POST("/upload", handlers.UploadImage)
	r.POST("/process", handlers.ProcessImage)
	r.GET("/healthcheck", healthCheck) // Healthcheck endpoint

	// Start server on defined port
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("startng server error")
	}
}

// healthCheck handles the health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
