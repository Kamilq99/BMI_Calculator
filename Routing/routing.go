package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	staticPath := filepath.Join("..", "frontend")
	router.Static("/static", staticPath)

	router.Run(":8080")
}
