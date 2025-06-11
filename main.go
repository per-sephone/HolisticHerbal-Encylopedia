package main

import (
	"holistic-herbal-encyclopedia/internal"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	logger "holistic-herbal-encyclopedia/internal/log"
)

func main() {
	InitLogger()

	router := gin.Default()
	router.Use(logger.LogRequestContext())

	router.Static("static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", internal.Index)
	router.POST("/create", internal.PostCreate)
	router.GET("/create", internal.GetCreate)
	router.GET("/search", internal.Search)
	router.GET("/nameSearch", internal.SearchByName)
	router.GET("/useSearch", internal.SearchByUse)
	router.GET("/edit", internal.GetEdit)
	router.POST("/edit", internal.PostEdit)
	router.Run(":8080")
}

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}