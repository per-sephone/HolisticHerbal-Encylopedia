package main

import (
	holistic "holistic-herbal-encyclopedia/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", holistic.Index)
	router.Run("localhost:8080")
}