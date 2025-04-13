package handlers

import (
	"calculate-bmi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Weight float64
var Height float64

func DataHandler(c *gin.Context) {
	var human models.Human

	if err := c.ShouldBindJSON(&human); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if human.Weight == 0 || human.Height == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing weight or height"})
		return
	}

	Weight = human.Weight
	Height = human.Height

	c.JSON(http.StatusOK, gin.H{
		"weight":  Weight,
		"height":  Height,
		"message": "All data has been saved successfully",
	})
}
