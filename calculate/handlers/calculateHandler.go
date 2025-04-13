package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateBMIHandler(c *gin.Context) {
	if Height == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Height must be greater than 0"})
		return
	}

	bmi := Weight / (Height * Height)
	c.JSON(http.StatusOK, gin.H{"BMI": bmi})
}
