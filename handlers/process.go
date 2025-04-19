package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProcessImage handles image processing
func ProcessImage(c *gin.Context) {
	// Logic for image processing
	c.JSON(http.StatusOK, gin.H{"message": "Image processed successfully"})
}
