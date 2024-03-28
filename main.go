package main

import (
	"holistic-herbal-encyclopedia/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", internal.Index)
	router.POST("/create", internal.PostCreate)
	router.GET("/create", internal.GetCreate)
	router.GET("/search", internal.Search)
	router.GET("/nameSearch", internal.SearchByName)
	router.GET("/useSearch", internal.SearchByUse)
	router.Run(":8001")
}