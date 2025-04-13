package main

import (
	"calculate-bmi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/data", handlers.DataHandler)
	r.GET("/calculate", handlers.CalculateBMIHandler)

	r.Run(":8080")
}
