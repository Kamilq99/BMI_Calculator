package main

import (
	"calculate-bmi/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/data", handlers.DataHandler)
	r.GET("/calculate", handlers.CalculateBMIHandler)

	r.Run("0.0.0.0:8080")
}
