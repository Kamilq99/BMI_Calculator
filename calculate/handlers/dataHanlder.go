package handlers

import (
	"calculate-bmi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DataHandler handles the POST request for creating a new Human.
func DataHandler(c *gin.Context) {
	var human models.Human

	// Bind the JSON data to the Human struct
	if err := c.ShouldBindJSON(&human); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if weight and height are provided
	if human.Weight == 0 || human.Height == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing weight or height"})
		return
	}

	// Return a success message with the received data
	c.JSON(http.StatusOK, gin.H{
		"weight":  human.Weight,
		"height":  human.Height,
		"message": "All data has been saved successfully",
	})
}
