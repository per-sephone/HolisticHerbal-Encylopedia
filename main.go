package main

import (
	"holistic-herbal-encyclopedia/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", internal.Index)
	router.Run(":8080")
}